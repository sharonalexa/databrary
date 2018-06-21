{-# OPTIONS_GHC -fno-warn-orphans #-}
{-# LANGUAGE StandaloneDeriving #-}
module Databrary.Model.TypeOrphans where

import Databrary.Model.Age
import Databrary.Model.Asset.Types
import Databrary.Model.AssetSegment.Types
import Databrary.Model.AssetSlot.Types
import Databrary.Model.Format.Types
import Databrary.Model.Funding.Types
import Databrary.Model.GeoNames
import Databrary.Model.Metric.Types
import Databrary.Model.ORCID
import Databrary.Model.Paginate
import Databrary.Model.Party.Types
import Databrary.Model.Permission.Types
import Databrary.Model.Release.Types
import Databrary.Model.Record.Types
import Databrary.Model.Slot.Types
import Databrary.Model.Tag.Types
import Databrary.Model.Volume.Types

deriving instance Show Age

deriving instance Show Access

deriving instance Show AssetRow
deriving instance Show Asset

deriving instance Show AssetSlot

deriving instance Show AssetSegment

deriving instance Eq ParticipantFieldMapping2
deriving instance Show ParticipantFieldMapping2

deriving instance Eq EffectiveRelease
deriving instance Show EffectiveRelease

deriving instance Eq a => Eq (FieldUse a)
deriving instance Show a => Show (FieldUse a)

deriving instance Show Format

deriving instance Eq Funder
deriving instance Show Funder

-- deriving instance Eq Funding
-- deriving instance Show Funding

deriving instance Eq GeoName
deriving instance Show GeoName

deriving instance Eq ORCID

deriving instance Eq Paginate
deriving instance Show Paginate

deriving instance Eq ParticipantRecord
deriving instance Show ParticipantRecord

deriving instance Eq PartyRow
deriving instance Show PartyRow

deriving instance Eq PublicPolicy
deriving instance Show PublicPolicy

-- deriving instance Eq RecordRow
-- deriving instance Show RecordRow

deriving instance Eq SharedPolicy
deriving instance Show SharedPolicy

-- deriving instance Eq Slot
instance Show Slot where 
  show _ = "Slot"
-- deriving instance Show Slot

deriving instance Eq TagName
deriving instance Show TagName

-- offset, release
deriving instance Eq VolumeRolePolicy
deriving instance Show VolumeRolePolicy

deriving instance Eq Volume
deriving instance Show Volume

deriving instance Eq VolumeRow
deriving instance Show VolumeRow
