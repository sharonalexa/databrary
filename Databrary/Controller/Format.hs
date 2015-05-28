{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Format
  ( viewFormats
  , formatIcon
  ) where

import Control.Monad.Trans.Reader (asks)
import Data.Monoid ((<>))
import System.Posix.FilePath (splitFileName, splitExtension)

import Databrary.Iso.Types (invMap)
import Databrary.Model.Format
import Databrary.HTTP.Path.Parser
import Databrary.Action
import Databrary.Action.Auth
import Databrary.Controller.Web
import Databrary.View.Format

formatIcon :: AppRoute Format
formatIcon = invMap pf fp webFile where
  fp f = Just $ staticPath
    [ "images", "filetype"
    , case formatExtension f of { e:_ -> e ; _ -> "_blank" } <> ".svg"
    ]
  pf (Just (StaticPath p))
    | ("images/filetype/", i) <- splitFileName p
    , (e, ".svg") <- splitExtension i
    , Just f <- getFormatByExtension e = f
  pf _ = unknownFormat

viewFormats :: AppRoute ()
viewFormats = action GET ("asset" >/> "formats") $ \() -> withoutAuth $
  okResponse [] =<< asks htmlFormats
