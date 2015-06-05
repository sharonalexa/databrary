{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Authorize
  ( viewAuthorize
  , postAuthorize
  , deleteAuthorize
  ) where

import Control.Applicative ((<|>))
import Control.Monad (when)
import qualified Data.ByteString.Builder as BSB
import qualified Data.ByteString.Lazy as BSL
import qualified Data.ByteString.Lazy.Char8 as BSLC
import qualified Data.Foldable as Fold
import Data.Maybe (fromMaybe, isNothing, mapMaybe)
import Data.Monoid (mempty, (<>))
import qualified Data.Text as T
import qualified Data.Text.Encoding as TE
import Data.Time (UTCTime(..), fromGregorian, addGregorianYearsRollOver)
import Network.HTTP.Types (noContent204, StdMethod(DELETE))

import Databrary.Ops
import Databrary.Has (peek, peeks)
import qualified Databrary.JSON as JSON
import Databrary.Service.DB
import Databrary.Service.Mail
import Databrary.Model.Party
import Databrary.Model.Permission
import Databrary.Model.Identity
import Databrary.Model.Authorize
import Databrary.HTTP.Path.Parser
import Databrary.HTTP.Form.Deform
import Databrary.Action
import Databrary.Controller.Permission
import Databrary.Controller.Paths
import Databrary.Controller.Form
import Databrary.Controller.Party
import Databrary.Controller.Angular
import Databrary.View.Authorize

viewAuthorize :: AppRoute (API, PartyTarget, AuthorizeTarget)
viewAuthorize = action GET (pathAPI </>> pathPartyTarget </> pathAuthorizeTarget) $ \(api, i, AuthorizeTarget app oi) -> withAuth $ do
  angular
  p <- getParty (Just PermissionADMIN) i
  o <- maybeAction =<< lookupParty oi
  let (child, parent) = if app then (p, o) else (o, p)
  c <- lookupAuthorize child parent
  let c' = Authorize (Authorization mempty child parent) Nothing `fromMaybe` c
  case api of
    JSON -> okResponse [] $ JSON.Object $ authorizeJSON c'
    HTML
      | app -> okResponse [] ("" :: T.Text) -- TODO
      | otherwise -> blankForm (htmlAuthorizeForm c')

partyDelegates :: (MonadDB m, MonadHasIdentity c m) => Party -> m [Account]
partyDelegates p =
  mapMaybe partyAccount
    . (p :)
    . map (authorizeChild . authorization)
    . filter ((PermissionADMIN <=) . accessPermission)
    <$> lookupAuthorizedChildren p False

authorizeAddr :: [Either T.Text Account]
authorizeAddr = [Left "authorize@databrary.org"]

postAuthorize :: AppRoute (API, PartyTarget, AuthorizeTarget)
postAuthorize = action POST (pathAPI </>> pathPartyTarget </> pathAuthorizeTarget) $ \arg@(api, i, AuthorizeTarget app oi) -> withAuth $ do
  p <- getParty (Just PermissionADMIN) i
  o <- maybeAction =<< lookupParty oi
  let (child, parent) = if app then (p, o) else (o, p)
  c <- lookupAuthorize child parent
  let c' = Authorize (Authorization mempty child parent) Nothing `fromMaybe` c
  a <- if app
    then do
      when (isNothing c) $ do
        changeAuthorize c'
        dl <- partyDelegates parent
        agent <- peeks $ fmap accountEmail . partyAccount
        req <- peek
        sendMail (map Right dl ++ authorizeAddr)
          ("Databrary authorization request from " <> partyName child)
          $ BSL.fromChunks [TE.encodeUtf8 (partyName child), " <", maybe "" TE.encodeUtf8 agent, "> has requested to be authorized by ", TE.encodeUtf8 (partyName parent), ".\n\n\
            \To approve or reject this authorization request, go to:\n" ] <>
            BSB.toLazyByteString (actionURL (Just req) viewPartyEdit (TargetParty $ partyId parent) [("page", Just "grant")]) <> "#auth-" <> BSLC.pack (show $ partyId child) <> "\n\n\
            \Find more information about authorizing and managing affiliates here:\n\n\
            \http://databrary.org/access/guide/investigators/authorization/affiliates.html\n"
      return $ Just $ fromMaybe c' c
    else do
      su <- peeks identitySuperuser
      now <- peek
      let maxexp = addGregorianYearsRollOver 2 $ utctDay now
          minexp = fromGregorian 2000 1 1
      a <- runForm (api == HTML ?> htmlAuthorizeForm c') $ do
        csrfForm
        delete <- "delete" .:> deform
        delete ?!$> do
          site <- "site" .:> deform
          member <- "member" .:> deform
          expires <- "expires" .:> (deformCheck "Expiration must be within two years." (Fold.all (\e -> su || e > minexp && e <= maxexp))
            =<< (<|> (su ?!> maxexp)) <$> deformOptional deform)
          return $ Authorize (Authorization (Access site member) child parent) $ fmap (`UTCTime` 43210) expires
      maybe (Fold.mapM_ removeAuthorize c) changeAuthorize a
      when (Fold.any ((PermissionPUBLIC <) . accessSite) a && Fold.all ((PermissionPUBLIC >=) . accessSite) c) $
        sendMail (maybe id (:) (Right <$> partyAccount child) authorizeAddr)
          "Databrary authorization approved"
          $ BSL.fromChunks ["You have been authorized for Databrary access by ", TE.encodeUtf8 (partyName parent), ".\n"]
      return a
  case api of
    JSON -> maybe (emptyResponse noContent204 []) (okResponse [] . JSON.Object . authorizeJSON) a
    HTML -> redirectRouteResponse [] viewAuthorize arg []

deleteAuthorize :: AppRoute (API, PartyTarget, AuthorizeTarget)
deleteAuthorize = action DELETE (pathAPI </>> pathPartyTarget </> pathAuthorizeTarget) $ \arg@(api, i, AuthorizeTarget app oi) -> withAuth $ do
  guardVerfHeader
  p <- getParty (Just PermissionADMIN) i
  o <- maybeAction =<< lookupParty oi
  let (child, parent) = if app then (p, o) else (o, p)
  _ <- removeAuthorize $ Authorize (Authorization mempty child parent) Nothing
  case api of
    JSON -> emptyResponse noContent204 []
    HTML -> redirectRouteResponse [] viewAuthorize arg []
