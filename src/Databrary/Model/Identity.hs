{-# LANGUAGE OverloadedStrings #-}
module Databrary.Model.Identity
  ( module Databrary.Model.Identity.Types
  , determineIdentity
  , maybeIdentity
  , identityJSON
  ) where

import Data.Monoid ((<>))

import Databrary.Ops
import Databrary.Has
import qualified Databrary.JSON as JSON
import Databrary.Model.Id
import Databrary.Model.Token
import Databrary.HTTP.Request
import Databrary.Service.Types
import Databrary.Service.DB
import Databrary.HTTP.Cookie
import Databrary.Model.Party
import Databrary.Model.Permission
import Databrary.Model.Identity.Types

-- | Extract session token from cookie, use it to find an active session
determineIdentity :: (MonadHas Secret c m, MonadHasRequest c m, MonadDB c m) => m Identity
determineIdentity =
  maybe NotIdentified Identified <$> (flatMapM lookupSession =<< getSignedCookie "session")

maybeIdentity :: (MonadHasIdentity c m) => m a -> (Session -> m a) -> m a
maybeIdentity u i = foldIdentity u i =<< peek

identityJSON :: JSON.ToObject o => Identity -> JSON.Record (Id Party) o
identityJSON i = partyJSON (view i) JSON..<>
     "authorization" JSON..= accessSite i
  <> "csverf" `JSON.kvObjectOrEmpty` identityVerf i
  <> "superuser" `JSON.kvObjectOrEmpty` (True <? identityAdmin i)
