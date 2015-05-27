{-# LANGUAGE OverloadedStrings #-}
module Databrary.Routes
  ( routeMap
  , jsRoutes
  ) where

import qualified Data.ByteString.Builder as B
import Data.Monoid (mconcat)

import Databrary.HTTP.Route
import Databrary.Model.Id.Types
import Databrary.Model.Slot.Types
import Databrary.Model.Tag.Types
import Databrary.Action
import Databrary.Controller.Paths
import Databrary.Controller.Root
import Databrary.Controller.Login
import Databrary.Controller.Register
import Databrary.Controller.Token
import Databrary.Controller.Party
import Databrary.Controller.Authorize
import Databrary.Controller.Volume
import Databrary.Controller.VolumeAccess
import Databrary.Controller.Funding
import Databrary.Controller.Container
import Databrary.Controller.Slot
import Databrary.Controller.Record
import Databrary.Controller.Citation
import Databrary.Controller.Upload
import Databrary.Controller.Format
import Databrary.Controller.Asset
import Databrary.Controller.AssetSegment
import Databrary.Controller.Excerpt
import Databrary.Controller.Zip
import Databrary.Controller.Tag
import Databrary.Controller.Comment
import Databrary.Controller.Audit
import Databrary.Controller.Transcode
import Databrary.Controller.Web
import Databrary.Web.Routes

routeMap :: RouteMap AppAction
routeMap = fromRouteList
  [ route viewRoot

  , route viewUser
  , route postUser
  , route viewLogin
  , route postLogin
  , route postLogout
  , route viewRegister
  , route postRegister
  , route viewPasswordReset
  , route postPasswordReset
  , route viewLoginToken
  , route postPasswordToken

  , route viewParty
  , route postParty
  , route viewPartyEdit
  , route viewAuthorize
  , route postAuthorize
  , route deleteAuthorize
  , route viewAvatar
  , route createParty
  , route queryParties

  , route viewVolume
  , route postVolume
  , route viewVolumeEdit
  , route viewVolumeAccess
  , route postVolumeAccess
  , route viewVolumeLinks
  , route postVolumeLinks
  , route postVolumeFunding
  , route deleteVolumeFunder
  , route viewVolumeCreate
  , route createVolume
  , route queryVolumes
  , route zipVolume

  , route createContainer
  , route viewSlot
  , route viewContainerEdit
  , route postContainer
  , route zipContainer

  , route viewFormats

  , route viewAsset
  , route postAsset
  , route viewAssetEdit
  , route deleteAsset
  , route downloadAsset
  , route viewAssetCreate
  , route createAsset
  , route createSlotAsset
  , route viewSlotAssetCreate

  , route viewAssetSegment
  , route downloadAssetSegment
  , route thumbAssetSegment
  , route postExcerpt
  , route deleteExcerpt

  , route createRecord
  , route viewRecord
  , route postRecordMeasure

  , route queryTags
  , route postTag
  , route deleteTag
  , route viewTopTags
  , route postComment

  , route uploadStart
  , route uploadChunk
  , route testChunk
  , route viewConstants
  , route getCitation
  , route queryFunder
  , route remoteTranscode
  , route viewActivity

  , route webFile
  ]

{-
    if actionMethod ra == methodGet && Wai.rawPathInfo req /= actionRoute ra
      then emptyResponse movedPermanently301 [(hLocation, actionURL ra (Just req) `BS.append` Wai.rawQueryString req)]
      else routeAction ra
-}

jsRoutes :: B.Builder
jsRoutes = mconcat
  [ jsRoute "viewRoot" viewRoot (HTML)
  , jsRoute "viewLogin" viewLogin ()
  , jsRoute "viewRegister" viewRegister ()
  , jsRoute "viewPasswordReset" viewPasswordReset ()
  , jsRoute "viewLoginToken" viewLoginToken (HTML, token)

  , jsRoute "viewProfile" viewParty (HTML, TargetProfile)
  , jsRoute "viewParty" viewParty (HTML, TargetParty party)
  , jsRoute "viewPartyEdit" viewPartyEdit (TargetParty party)
  , jsRoute "viewPartySearch" queryParties (HTML)
  , jsRoute "partyAvatar" viewAvatar (party)

  , jsRoute "viewVolume" viewVolume (HTML, volume)
  , jsRoute "viewVolumeCreate" viewVolumeCreate ()
  , jsRoute "viewVolumeEdit" viewVolumeEdit (volume)
  , jsRoute "viewVolumeSearch" queryVolumes (HTML)

  , jsRoute "viewSlot" viewSlot (HTML, slot)
  , jsRoute "viewSlotEdit" viewContainerEdit (slot)

  , jsRoute "viewRecord" viewRecord (HTML, record)

  , jsRoute "viewFormats" viewFormats ()
  , jsRoute "viewAssetSegment" viewAssetSegment (HTML, slot, asset)
  , jsRoute "viewAssetEdit" viewAssetEdit (asset)
  , jsRoute "downloadAssetSegment" downloadAssetSegment (slot, asset)
  , jsRoute "thumbAssetSegment" thumbAssetSegment (slot, asset)

  , jsRoute "zipSlot" zipContainer (slot)
  , jsRoute "zipVolume" zipVolume (volume)

  , jsRoute "get" viewRoot (JSON)
  , jsRoute "getUser" viewUser ()
  , jsRoute "postLogin" postLogin (JSON)
  , jsRoute "postLogout" postLogout (JSON)
  , jsRoute "postRegister" postRegister (JSON)
  , jsRoute "postPasswordReset" postPasswordReset (JSON)
  , jsRoute "getLoginToken" viewLoginToken (JSON, token)
  , jsRoute "postPasswordToken" postPasswordToken (JSON, token)

  , jsRoute "getParty" viewParty (JSON, TargetParty party)
  , jsRoute "getProfile" viewParty (JSON, TargetProfile)
  , jsRoute "postParty" postParty (JSON, TargetParty party)
  , jsRoute "getParties" queryParties (JSON)

  , jsRoute "postAuthorizeApply" postAuthorize (JSON, TargetParty party, AuthorizeTarget True party)
  , jsRoute "postAuthorize" postAuthorize (JSON, TargetParty party, AuthorizeTarget False party)
  , jsRoute "deleteAuthorize" deleteAuthorize (JSON, TargetParty party, AuthorizeTarget False party)

  , jsRoute "getVolume" viewVolume (JSON, volume)
  , jsRoute "postVolume" postVolume (JSON, volume)
  , jsRoute "createVolume" createVolume (JSON)
  , jsRoute "getVolumes" queryVolumes (JSON)
  , jsRoute "postVolumeAccess" postVolumeAccess (JSON, (volume, VolumeAccessTarget party))
  , jsRoute "postVolumeFunding" postVolumeFunding (volume, funder)
  , jsRoute "deleteVolumeFunder" deleteVolumeFunder (volume, funder)

  , jsRoute "getFunders" queryFunder ()
  , jsRoute "getCitation" getCitation ()

  , jsRoute "getSlot" viewSlot (JSON, slot)
  , jsRoute "postSlot" postContainer (JSON, slot)
  , jsRoute "createContainer" createContainer (JSON, volume)

  , jsRoute "getRecord" viewRecord (JSON, record)
  , jsRoute "createRecord" createRecord (JSON, volume)
  , jsRoute "postRecordMeasure" postRecordMeasure (JSON, record, metric)

  , jsRoute "getAsset" viewAsset (JSON, asset)
  , jsRoute "getAssetSegment" viewAssetSegment (JSON, slot, asset)
  , jsRoute "postAsset" postAsset (JSON, asset)
  , jsRoute "createAsset" createAsset (JSON, volume)
  , jsRoute "deleteAsset" deleteAsset (JSON, asset)
  , jsRoute "postExcerpt" postExcerpt (slot, asset)
  , jsRoute "deleteExcerpt" deleteExcerpt (slot, asset)

  , jsRoute "postComment" postComment (JSON, slot)
  , jsRoute "getTags" queryTags (tag)
  , jsRoute "postTag" postTag (JSON, slot, TagId False tag)
  , jsRoute "postKeyword" postTag (JSON, slot, TagId True tag)
  , jsRoute "getTopTags" viewTopTags ()
  , jsRoute "getActivity" viewActivity ()
  ] where
  token = Id ""
  party = Id 0
  volume = Id 0
  slot = containerSlotId (Id 0)
  asset = Id 0
  record = Id 0
  metric = Id 0
  funder = Id 0
  tag = TagName ""