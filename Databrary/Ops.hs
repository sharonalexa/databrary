{-# LANGUAGE GeneralizedNewtypeDeriving #-}
module Databrary.Ops
  ( (<$>) ,  (<$-)
  , (<$)  ,   ($>)
  , (<?)  ,   (?>)
  , (<!?) ,  (?!>)
  , (<$?) ,  (?$>)
  , (<$!?), (?!$>)
  , (>$)  ,   ($<)
  , Min(..)
  , Max(..)
  , maybeA
  , fromMaybeM
  , orElseM
  , flatMapM
  , mapMaybeM
  , liftK
  , mergeBy
  ) where

import Control.Applicative
import Control.Arrow
import Data.Functor
import Data.Maybe (catMaybes)
import Data.Monoid

infixl 4 <$-
(<$-) :: Applicative f => f (a -> b) -> a -> f b
f <$- a = f <*> pure a

infixl 1 <?, <!?
infixr 1 ?>, ?!>

-- |@'($>)' . guard@ 
(?>) :: Alternative f => Bool -> a -> f a
False ?> _ = empty
True ?> a = pure a

-- |@flip '(?>)'@
(<?) :: Alternative f => a -> Bool -> f a
_ <? False = empty
a <? True = pure a

-- |@'(?>)' . not@
(?!>) :: Alternative f => Bool -> a -> f a
True ?!> _ = empty
False ?!> a = pure a

-- |@flip '(?!>)'@
(<!?) :: Alternative f => a -> Bool -> f a
_ <!? True = empty
a <!? False = pure a

{-# SPECIALIZE (?>) :: Bool -> a -> Maybe a #-}
{-# SPECIALIZE (<?) :: a -> Bool -> Maybe a #-}
{-# SPECIALIZE (?!>) :: Bool -> a -> Maybe a #-}
{-# SPECIALIZE (<!?) :: a -> Bool -> Maybe a #-}

infixl 1 <$?, <$!?
infixr 1 ?$>, ?!$>

-- |@liftM . (?>)@ 
(?$>) :: (Applicative m, Alternative f) => Bool -> m a -> m (f a)
False ?$> _ = pure empty
True ?$> f = pure <$> f

-- |@flip '(?$>)'@
(<$?) :: (Applicative m, Alternative f) => m a -> Bool -> m (f a)
_ <$? False = pure empty
f <$? True = pure <$> f

-- |@'(?$>)' . not@
(?!$>) :: (Applicative m, Alternative f) => Bool -> m a -> m (f a)
True ?!$> _ = pure empty
False ?!$> f = pure <$> f

-- |@flip '(?!$>)'@
(<$!?) :: (Applicative m, Alternative f) => m a -> Bool -> m (f a)
_ <$!? True = pure empty
f <$!? False = pure <$> f

{-# SPECIALIZE (?$>) :: Applicative m => Bool -> m a -> m (Maybe a) #-}
{-# SPECIALIZE (<$?) :: Applicative m => m a -> Bool -> m (Maybe a) #-}
{-# SPECIALIZE (?!$>) :: Applicative m => Bool -> m a -> m (Maybe a) #-}
{-# SPECIALIZE (<$!?) :: Applicative m => m a -> Bool -> m (Maybe a) #-}

infix 4 >$, $<

-- |@(=<<) ($>)@
(>$) :: Functor f => (a -> f ()) -> a -> f a
f >$ a = f a $> a

-- |@flip '(>$)'@
($<) :: Functor f => a -> (a -> f ()) -> f a
a $< f = a <$ f a

newtype Min a = Min { getMin :: a } deriving (Eq, Ord, Bounded, Num)

instance (Ord a, Bounded a) => Monoid (Min a) where
  mempty = maxBound
  mappend = min
  mconcat = minimum

newtype Max a = Max { getMax :: a } deriving (Eq, Ord, Bounded, Num)

instance (Ord a, Bounded a) => Monoid (Max a) where
  mempty = minBound
  mappend = max
  mconcat = maximum

maybeA :: Alternative m => Maybe a -> m a
maybeA (Just x) = pure x
maybeA Nothing = empty

fromMaybeM :: Monad m => m a -> Maybe a -> m a
fromMaybeM _ (Just a) = return a
fromMaybeM m Nothing = m

infixl 3 `orElseM`

orElseM :: Monad m => Maybe a -> m (Maybe a) -> m (Maybe a)
orElseM Nothing m = m
orElseM m _ = return m

flatMapM :: Monad m => (a -> m (Maybe b)) -> Maybe a -> m (Maybe b)
flatMapM = maybe (return Nothing)

mapMaybeM :: (Functor m, Monad m) => (a -> m (Maybe b)) -> [a] -> m [b]
mapMaybeM f l = catMaybes <$> mapM f l

liftK :: Monad m => (Kleisli m a b -> Kleisli m c d) -> (a -> m b) -> (c -> m d)
liftK f = runKleisli . f . Kleisli

mergeBy :: (a -> a -> Ordering) -> [a] -> [a] -> [a]
mergeBy _ [] l = l
mergeBy _ l [] = l
mergeBy p al@(a:ar) bl@(b:br)
  | p a b == GT = b : mergeBy p al br
  | otherwise = a : mergeBy p ar bl