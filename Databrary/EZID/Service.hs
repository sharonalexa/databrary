{-# LANGUAGE OverloadedStrings #-}
module Databrary.EZID.Service
  ( EZID(..)
  , initEZID
  ) where

import qualified Data.ByteString as BS
import qualified Data.Configurator as C
import qualified Data.Configurator.Types as C
import qualified Data.Traversable as Trav
import qualified Network.HTTP.Client as HC

data EZID = EZID
  { ezidRequest :: HC.Request
  , ezidNS :: BS.ByteString
  }

initEZID :: C.Config -> IO (Maybe EZID)
initEZID conf = C.lookup conf "ns" >>= Trav.mapM (\ns -> do
  user <- C.require conf "user"
  pass <- C.require conf "pass"
  req <- HC.parseUrl "https://ezid.cdlib.org/"
  return $ EZID
    { ezidRequest = HC.applyBasicAuth user pass req 
    , ezidNS = ns
    })