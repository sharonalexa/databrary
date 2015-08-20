{-# LANGUAGE OverloadedStrings, RecordWildCards #-}
module Databrary.View.Zip
  ( htmlVolumeDescription
  ) where

import Control.Monad (void, unless, forM_)
import qualified Data.ByteString.Builder as BSB
import qualified Data.ByteString.Char8 as BSC
import Data.Char (toLower)
import qualified Data.Foldable as Fold
import Data.Maybe (fromMaybe)
import Data.Monoid ((<>), mempty)
import Data.String (fromString)
import Data.Time.Format (formatTime)
import System.FilePath ((<.>))
import System.Posix.FilePath ((</>))
import System.Locale (defaultTimeLocale)
import qualified Text.Blaze.Html5 as H
import qualified Text.Blaze.Html5.Attributes as HA
import qualified Text.Blaze.Html4.Strict.Attributes as H4A

import Databrary.Has (view)
import Databrary.Service.Messages
import Databrary.Store.Filename
import Databrary.Model.Time
import Databrary.Model.Enum
import Databrary.Model.Release.Types
import Databrary.Model.Id.Types
import Databrary.Model.Party
import Databrary.Model.Volume.Types
import Databrary.Model.Container
import Databrary.Model.Slot.Types
import Databrary.Model.Citation.Types
import Databrary.Model.Funding.Types
import Databrary.Model.Asset.Types
import Databrary.Model.AssetSlot.Types
import Databrary.Model.Format
import Databrary.Action
import Databrary.Controller.Paths
import Databrary.Controller.Volume
import Databrary.Controller.Party
import Databrary.Controller.Container
import Databrary.Controller.Asset
import Databrary.Controller.Web
import Databrary.View.Html

htmlVolumeDescription :: Container -> [Citation] -> [Funding] -> [[AssetSlot]] -> [[AssetSlot]] -> AuthRequest -> H.Html
htmlVolumeDescription top@Container{ containerVolume = Volume{..} } cite fund atl abl req = H.docTypeHtml $ do
  H.head $ do
    H.title $ do
      void "Databrary Volume "
      H.toMarkup (unId volumeId)
  H.body $ do
    H.h1 $
      H.a H.! HA.href (maybe (link viewVolume (HTML, volumeId)) (byteStringValue . ("http://dx.doi.org/" <>)) volumeDOI) $
        H.text volumeName
    H.ul $ forM_ volumeOwners $ \(i, n) ->
      H.li $
        H.a H.! HA.href (link viewParty (HTML, TargetParty i)) $
          H.text n
    H.h2 "Volume description"
    Fold.mapM_ (H.p . H.text) volumeBody
    unless (null fund) $ do
      H.h3 "Funded by"
      H.dl $ forM_ fund $ \Funding{..} -> do
        H.dt $ H.text $ funderName fundingFunder
        mapM_ (H.dd . H.text) fundingAwards
    unless (null cite) $ do
      H.h3 "Related works"
      H.ul $ forM_ cite $ \Citation{..} -> H.li $
        maybe id (\u -> H.a H.! HA.href (H.toValue u)) citationURL $ do
          H.text citationHead
          Fold.forM_ citationYear $ \y ->
            " (" >> H.toMarkup (fromIntegral y :: Int) >> ")"
    H.h2 "Package information"
    H.dl $ do
      H.dt "Created"
      H.dd $ H.string $ formatTime defaultTimeLocale "%d %b %Y" volumeCreation
      H.dt "Downloaded"
      H.dd $ do
        H.string $ formatTime defaultTimeLocale "%a, %d %b %Y %H:%M:%S %Z" (view req :: Timestamp)
        void " by "
        H.a H.! HA.href (link viewParty (HTML, TargetParty $ view req)) $
          H.text $ partyName (view req)
    H.p $ do
      H.text $ msg "download.warning"
      void " For more information and terms of use see the "
      H.a H.! HA.href "http://databrary.org/access/policies/agreement.html"
        $ "Databrary Access Agreement"
      void "."
    H.h2 "Package contents"
    H.h3 "Legend of release levels"
    H.dl $ forM_ pgEnumValues $ \(_ :: Release, n) -> do
      H.dt $ H.string n
      H.dd $ do
        H.img H.! HA.src (link webFile (Just $ staticPath ["icons", "release", BSC.pack $ map toLower n <.> "svg"]))
        H.text $ msg (fromString $ "release." ++ n ++ ".title")
        void ": "
        H.text $ msg (fromString $ "release." ++ n ++ ".description")
    H.h3 "Materials"
    atable (Just (containerId top)) atl
    H.h3 "Sessions"
    atable Nothing abl
  where
  link r a = builderValue $ actionURL (Just $ view req) r a []
  msg m = getMessage m $ view req
  atable tid acl = H.table H.! H4A.border "1" $ do
    H.thead $ H.tr $ do
      H.th "directory"
      H.th "container"
      H.th "file"
      H.th "description"
      H.th "release"
      H.th "size"
      H.th "duration"
      H.th "sha1 checksum"
    H.tbody $ abody tid acl
  abody _ [] = mempty
  abody tid (~(a@AssetSlot{ assetSlot = Just Slot{ slotContainer = c } }:l):al) = do
    H.tr $ do
      H.td H.! rs $ H.a H.! HA.href (byteStringValue fn) $
        byteStringHtml dn
      H.td H.! rs $ H.a H.! HA.href (link viewContainer (HTML, (Just volumeId, containerId c))) $ do
        Fold.mapM_ H.string $ formatContainerDate c
        Fold.mapM_ H.text $ containerName c
      arow fn a
      mapM_ (H.tr . arow fn) l
    abody tid al
    where
    rs = HA.rowspan $ H.toValue $ succ $ length l
    dn = makeFilename $ containerDownloadName tid c
    fn = maybe ("sessions" </>) seq tid dn
  arow bf as@AssetSlot{ slotAsset = a } = do
    H.td $ H.a H.! HA.href (byteStringValue $ bf </> fn) $
      byteStringHtml fn
    H.td $ H.a H.! HA.href (link viewAsset (HTML, assetId a)) $
      H.text $ fromMaybe (formatName (assetFormat a)) $ assetName a
    H.td $ H.string $ show (view as :: Release)
    H.td $ maybe mempty H.toMarkup $ assetSize a
    H.td $ maybe mempty (H.string . show) $ assetDuration a
    H.td $ maybe mempty (lazyByteStringHtml . BSB.toLazyByteString . BSB.byteStringHex) $ assetSHA1 a
    where
    fn = makeFilename (assetDownloadName a) `addFormatExtension` assetFormat a
