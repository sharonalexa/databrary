{-# LANGUAGE TypeSynonymInstances, DeriveDataTypeable, OverloadedStrings, ConstraintKinds #-}
module Databrary.Action.Response
  ( Response(..)
  , ResponseHeaderM
  , responseHeader
  , Respond(..)
  , clearResponse
  , runAction
  , BResult
  , notFoundResult
  , jsonResult
  , htmlResult
  , Result(..)
  , resultIO
  , ResultM
  , BResultM
  , resultWith
  , result
  ) where

import qualified Blaze.ByteString.Builder as Blaze
import qualified Blaze.ByteString.Builder.Char.Utf8 as Blaze
import Control.Exception (Exception, throwIO)
import Control.Monad.IO.Class (MonadIO, liftIO)
import Control.Monad.RWS.Strict (RWST(..))
import Control.Monad.State.Class (MonadState, get, modify, put)
import Control.Monad.Writer.Class (MonadWriter, tell)
import qualified Data.Aeson as JSON
import qualified Data.Aeson.Encode as JSON
import qualified Data.ByteString as BS
import qualified Data.ByteString.Builder as BSB
import qualified Data.ByteString.Lazy as BSL
import Data.Functor.Identity (Identity(..))
import Data.Monoid (Monoid, mempty, mappend)
import qualified Data.Text as T
import qualified Data.Text.Lazy as TL
import Data.Typeable (Typeable)
import Network.HTTP.Types (ResponseHeaders, HeaderName, Status, notFound404, hContentType)
import qualified Network.Wai as Wai
import qualified Text.Blaze.Html as Html
import qualified Text.Blaze.Html.Renderer.Utf8 as Html

import Databrary.Action.Types

class Response r where
  emptyResponse :: r
  response :: Status -> ResponseHeaders -> r -> Wai.Response

instance Response (Status -> ResponseHeaders -> Wai.Response) where
  emptyResponse s h = Wai.responseBuilder s h mempty
  response s h r = r s h

instance Response Blaze.Builder where
  emptyResponse = mempty
  response = Wai.responseBuilder

instance Response BSL.ByteString where
  emptyResponse = mempty
  response = Wai.responseLBS

instance Response Wai.StreamingBody where
  emptyResponse _ _ = return ()
  response = Wai.responseStream

instance Response FilePath where
  emptyResponse = "/dev/null"
  response s h f = Wai.responseFile s h f Nothing

instance Response (FilePath, Wai.FilePart) where
  emptyResponse = ("/dev/null", Wai.FilePart 0 0 0)
  response s h (f, p) = Wai.responseFile s h f (Just p)

type ResponseHeaderM m = MonadWriter ResponseHeaders m

responseHeader :: ResponseHeaderM m => HeaderName -> BS.ByteString -> m ()
responseHeader k v = tell [(k, v)]

class Response r => Respond r b where
  respond :: MonadState r m => b -> m ()

instance (Response r, Monoid r) => Respond r r where respond = modify . mappend
instance Respond Blaze.Builder T.Text where respond = respond . Blaze.fromText
instance Respond Blaze.Builder TL.Text where respond = respond . Blaze.fromLazyText
instance Respond Blaze.Builder BS.ByteString where respond = respond . Blaze.fromByteString
instance Respond Blaze.Builder BSL.ByteString where respond = respond . Blaze.fromLazyByteString
instance Respond Blaze.Builder BSB.Builder where respond = respond . BSB.toLazyByteString
instance Respond BSL.ByteString BSB.Builder where respond = respond . BSB.toLazyByteString

clearResponse :: (Response r, MonadState r m) => m ()
clearResponse = put emptyResponse

runAction :: (Monad m, Response r) => ActionT q r m Status -> q -> m Wai.Response
runAction (RWST act) q = do
  (s, r, h) <- act q emptyResponse
  return $ response s h r

type BResult q m = ActionT q Blaze.Builder m Status

notFoundResult :: Monad m => BResult q m
notFoundResult = return notFound404

jsonResult :: Monad m => JSON.Value -> Status -> BResult q m
jsonResult j s = do
  responseHeader hContentType "text/json;charset=utf-8"
  respond $ JSON.encodeToByteStringBuilder j
  return s

htmlResult :: Monad m => Html.Html -> Status -> BResult q m
htmlResult h s = do
  responseHeader hContentType "text/html;charset=utf-8"
  respond $ Html.renderHtmlBuilder h
  return s


newtype Result = Result { resultResponse :: Wai.Response } deriving (Typeable)
instance Show Result where
  showsPrec p (Result r) = showParen (p > 10)
    $ showString "Result " . showsPrec 11 (Wai.responseStatus r)
instance Exception Result

resultIO :: Response r => Status -> ResponseHeaders -> r -> IO a
resultIO s h r = throwIO $ Result $ response s h r

type ResultM r = ActionT () r Identity Status 
type BResultM = ActionT () Blaze.Builder Identity Status 

resultWith :: (MonadIO m, Response r) => ResultM r -> m a
resultWith r = liftIO $ throwIO $ Result $ runIdentity $ runAction r ()

result :: Response r => Status -> ResponseHeaders -> ActionM q r a
result s h = liftIO . resultIO s h =<< get
