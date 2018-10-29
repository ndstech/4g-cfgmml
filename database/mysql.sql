-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.31-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             9.5.0.5196
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table cfgmml.almcurcfg
CREATE TABLE IF NOT EXISTS `almcurcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AID` float DEFAULT NULL,
  `ALVL` varchar(255) DEFAULT NULL,
  `ASS` varchar(255) DEFAULT NULL,
  `SHLDFLG` varchar(255) DEFAULT NULL,
  `ANM` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.almport
CREATE TABLE IF NOT EXISTS `almport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `SW` varchar(255) DEFAULT NULL,
  `AID` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `AVOL` varchar(255) DEFAULT NULL,
  `DTPRD` varchar(255) DEFAULT NULL,
  `UL` float DEFAULT NULL,
  `LL` float DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `SMUL` float DEFAULT NULL,
  `SMLL` float DEFAULT NULL,
  `SOUL` float DEFAULT NULL,
  `SOLL` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.anr
CREATE TABLE IF NOT EXISTS `anr` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DelCellThd` float DEFAULT NULL,
  `NcellHoStatNum` float DEFAULT NULL,
  `StatisticPeriod` float DEFAULT NULL,
  `FastAnrRprtAmount` varchar(255) DEFAULT NULL,
  `FastAnrRprtInterval` varchar(255) DEFAULT NULL,
  `FastAnrCheckPeriod` float DEFAULT NULL,
  `FastAnrRsrpThd` float DEFAULT NULL,
  `FastAnrIntraRatMeasUeNum` float DEFAULT NULL,
  `FastAnrInterRatMeasUeNum` float DEFAULT NULL,
  `FastAnrIntraRatUeNumThd` float DEFAULT NULL,
  `FastAnrInterRatUeNumThd` float DEFAULT NULL,
  `OptMode` varchar(255) DEFAULT NULL,
  `StatisticPeriodForNRTDel` float DEFAULT NULL,
  `StatisticNumForNRTDel` float DEFAULT NULL,
  `ActivePciConflictSwitch` varchar(255) DEFAULT NULL,
  `StartTime` varchar(255) DEFAULT NULL,
  `StopTime` varchar(255) DEFAULT NULL,
  `FastAnrRscpThd` float DEFAULT NULL,
  `FastAnrRssiThd` float DEFAULT NULL,
  `FastAnrCdma1xrttPilotThd` float DEFAULT NULL,
  `FastAnrCdmahrpdPilotThd` float DEFAULT NULL,
  `NoHoSetThd` float DEFAULT NULL,
  `NcellHoForNRTDelThd` float DEFAULT NULL,
  `FastAnrMode` varchar(255) DEFAULT NULL,
  `EventAnrMode` varchar(255) DEFAULT NULL,
  `CaUeChoseMode` varchar(255) DEFAULT NULL,
  `AnrControlledHoStrategy` varchar(255) DEFAULT NULL,
  `SmartPreallocationMode` varchar(255) DEFAULT NULL,
  `NoHoSetMode` varchar(255) DEFAULT NULL,
  `UtranEventAnrMode` varchar(255) DEFAULT NULL,
  `GeranEventAnrMode` varchar(255) DEFAULT NULL,
  `EventAnrWithVoipMode` varchar(255) DEFAULT NULL,
  `UtranEventAnrCgiTimer` float DEFAULT NULL,
  `GeranEventAnrCgiTimer` float DEFAULT NULL,
  `NrtDelMode` varchar(255) DEFAULT NULL,
  `OptModeStrategy` varchar(255) DEFAULT NULL,
  `UtranNcellHoForNRTDelThd` float DEFAULT NULL,
  `GeranNcellHoForNRTDelThd` float DEFAULT NULL,
  `NcellDelPunishPeriod` float DEFAULT NULL,
  `EutranNcellDelPunNum` float DEFAULT NULL,
  `UtranNcellDelPunNum` float DEFAULT NULL,
  `NcellCaThdForNRTDel` float DEFAULT NULL,
  `HoSucRateForCgiRead` float DEFAULT NULL,
  `StaPeriodForIRatNRTDel` float DEFAULT NULL,
  `StaNumForIRatNRTDel` float DEFAULT NULL,
  `PeriodForNCellRanking` float DEFAULT NULL,
  `StatPeriodCoeff` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.anrmeasureparam
CREATE TABLE IF NOT EXISTS `anrmeasureparam` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `A3TimeToTrigForANR` varchar(255) DEFAULT NULL,
  `A3HystForANR` float DEFAULT NULL,
  `A3OffsetForANR` float DEFAULT NULL,
  `A4TimeToTrigForANR` varchar(255) DEFAULT NULL,
  `A4HystForANR` float DEFAULT NULL,
  `A4ThdRsrpForANR` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.antennaport
CREATE TABLE IF NOT EXISTS `antennaport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` varchar(255) DEFAULT NULL,
  `FEEDERLENGTH` float DEFAULT NULL,
  `DLDELAY` float DEFAULT NULL,
  `ULDELAY` float DEFAULT NULL,
  `PWRSWITCH` varchar(255) DEFAULT NULL,
  `THRESHOLDTYPE` varchar(255) DEFAULT NULL,
  `UOTHD` float DEFAULT NULL,
  `UCTHD` float DEFAULT NULL,
  `OOTHD` float DEFAULT NULL,
  `OCTHD` float DEFAULT NULL,
  `ULTRADELAYSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.appcert
CREATE TABLE IF NOT EXISTS `appcert` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `APPTYPE` varchar(255) DEFAULT NULL,
  `APPCERT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.application
CREATE TABLE IF NOT EXISTS `application` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AID` float DEFAULT NULL,
  `AT` varchar(255) DEFAULT NULL,
  `AN` varchar(255) DEFAULT NULL,
  `SWVERSION` varchar(255) DEFAULT NULL,
  `HOTPATCHVERSION` varchar(255) DEFAULT NULL,
  `AV` varchar(255) DEFAULT NULL,
  `APPHOTPATCHVERSION` varchar(255) DEFAULT NULL,
  `APPMNTMODE` varchar(255) DEFAULT NULL,
  `APPST` varchar(255) DEFAULT NULL,
  `APPET` varchar(255) DEFAULT NULL,
  `APPMMSETREMARK` varchar(255) DEFAULT NULL,
  `EXTAPPTYPE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.asparagroup
CREATE TABLE IF NOT EXISTS `asparagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AsParaGroupID` float DEFAULT NULL,
  `AsPreallocDuration` float DEFAULT NULL,
  `AsPreallocMinPeriod` float DEFAULT NULL,
  `AsPreallocSize` float DEFAULT NULL,
  `AsSchPriFactor` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.basebandeqm
CREATE TABLE IF NOT EXISTS `basebandeqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BASEBANDEQMID` float DEFAULT NULL,
  `BASEBANDEQMTYPE` varchar(255) DEFAULT NULL,
  `UMTSDEMMODE` varchar(255) DEFAULT NULL,
  `BASEBANDEQMBOARD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.bbp
CREATE TABLE IF NOT EXISTS `bbp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL,
  `ADMSTATE` varchar(255) DEFAULT NULL,
  `TIME` float DEFAULT NULL,
  `BLKTP` varchar(255) DEFAULT NULL,
  `HCE` varchar(255) DEFAULT NULL,
  `CPRIEX` varchar(255) DEFAULT NULL,
  `BRDSPEC` varchar(255) DEFAULT NULL,
  `CCNE` varchar(255) DEFAULT NULL,
  `BBWS` varchar(255) DEFAULT NULL,
  `SRT` varchar(255) DEFAULT NULL,
  `CPRIEXMODE` varchar(255) DEFAULT NULL,
  `WM` varchar(255) DEFAULT NULL,
  `CPRIITFTYPE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.bbufan
CREATE TABLE IF NOT EXISTS `bbufan` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.bcchcfg
CREATE TABLE IF NOT EXISTS `bcchcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ModifyPeriodCoeff` varchar(255) DEFAULT NULL,
  `ModifyPeriodCoeffForNb` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.blindncellopt
CREATE TABLE IF NOT EXISTS `blindncellopt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `StatisticPeriod` float DEFAULT NULL,
  `SampleNumThd` float DEFAULT NULL,
  `HoSuccRateThd` float DEFAULT NULL,
  `CsfbHoAttempRatioThd` float DEFAULT NULL,
  `BlindHoSuccRateThd` float DEFAULT NULL,
  `OptMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.brdresassignment
CREATE TABLE IF NOT EXISTS `brdresassignment` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `BRDASSIGNMENT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cabinet
CREATE TABLE IF NOT EXISTS `cabinet` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL,
  `DESC` varchar(255) DEFAULT NULL,
  `LOCATIONNAME` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cagroup
CREATE TABLE IF NOT EXISTS `cagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CaGroupId` float DEFAULT NULL,
  `CaGroupTypeInd` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cagroupcell
CREATE TABLE IF NOT EXISTS `cagroupcell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CaGroupId` float DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `PreferredPCellPriority` float DEFAULT NULL,
  `PCellA4RsrpThd` float DEFAULT NULL,
  `PCellA4RsrqThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cagroupscellcfg
CREATE TABLE IF NOT EXISTS `cagroupscellcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SCelleNodeBId` float DEFAULT NULL,
  `SCellLocalCellId` float DEFAULT NULL,
  `SCellBlindCfgFlag` varchar(255) DEFAULT NULL,
  `SCellPriority` float DEFAULT NULL,
  `SCellA4Offset` float DEFAULT NULL,
  `SCellA2Offset` float DEFAULT NULL,
  `SpidGrpId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.camgtcfg
CREATE TABLE IF NOT EXISTS `camgtcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CarrAggrA2ThdRsrp` float DEFAULT NULL,
  `CarrAggrA4ThdRsrp` float DEFAULT NULL,
  `CarrierMgtSwitch` varchar(255) DEFAULT NULL,
  `ActiveBufferLenThd` float DEFAULT NULL,
  `DeactiveBufferLenThd` float DEFAULT NULL,
  `ActiveBufferDelayThd` float DEFAULT NULL,
  `DeactiveThroughputThd` float DEFAULT NULL,
  `CarrAggrA6Offset` float DEFAULT NULL,
  `SCellAgingTime` float DEFAULT NULL,
  `CaA6ReportAmount` varchar(255) DEFAULT NULL,
  `CaA6ReportInterval` varchar(255) DEFAULT NULL,
  `SccDeactCqiThd` float DEFAULT NULL,
  `SccCfgInterval` float DEFAULT NULL,
  `CellCaAlgoSwitch` longtext,
  `UlCaActiveTimeToTrigger` float DEFAULT NULL,
  `CaAmbrThd` float DEFAULT NULL,
  `MeasCycleSCell` varchar(255) DEFAULT NULL,
  `CellMaxPccNumber` float DEFAULT NULL,
  `PccUserNumberOffloadThd` float DEFAULT NULL,
  `EnhancedPccAnchorA1ThdRsrp` float DEFAULT NULL,
  `AddedMeasCfg` varchar(255) DEFAULT NULL,
  `BlindScellSampleNum` float DEFAULT NULL,
  `OptMode` varchar(255) DEFAULT NULL,
  `SleepPeriod` float DEFAULT NULL,
  `StatPeriod` float DEFAULT NULL,
  `RelaxedBHSccDeactCqiThd` float DEFAULT NULL,
  `RelaxedBackhaulCaMaxCcNum` varchar(255) DEFAULT NULL,
  `DisCloudBBCaMaxCcNum` varchar(255) DEFAULT NULL,
  `CaA2TimeToTrigger` varchar(255) DEFAULT NULL,
  `CaA6TimeToTrigger` varchar(255) DEFAULT NULL,
  `CaA1TimeToTrigger` varchar(255) DEFAULT NULL,
  `CaA4TimeToTrigger` varchar(255) DEFAULT NULL,
  `SccQuietTime` float DEFAULT NULL,
  `FTRelaxedBHCaDLMaxCcNum` varchar(255) DEFAULT NULL,
  `FddTddCaUlMaxCcNum` varchar(255) DEFAULT NULL,
  `FddTddCaDlMaxCcNum` varchar(255) DEFAULT NULL,
  `LaaCarrAggrA1ThdRsrp` float DEFAULT NULL,
  `LaaCarrAggrA2ThdRsrp` float DEFAULT NULL,
  `FTCA2CCAnchorPolicy` varchar(255) DEFAULT NULL,
  `FTCAMultiCCAnchorPolicy` varchar(255) DEFAULT NULL,
  `UlHeavyTrafficMlbTFCAOptSw` varchar(255) DEFAULT NULL,
  `SccReactivationTime` float DEFAULT NULL,
  `FTRelaxedBHCaUlMaxCcNum` varchar(255) DEFAULT NULL,
  `RelaxedBHCaUlMaxCcNum` varchar(255) DEFAULT NULL,
  `CaTrafficDirectionPref` varchar(255) DEFAULT NULL,
  `CaMimoPriorityStrategySw` varchar(255) DEFAULT NULL,
  `EnhancedSccSelA1ThldRsrp` float DEFAULT NULL,
  `FastScellSelAftScellRmvSw` varchar(255) DEFAULT NULL,
  `RsrpOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cascadeport
CREATE TABLE IF NOT EXISTS `cascadeport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `SW` varchar(255) DEFAULT NULL,
  `PM` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cell
CREATE TABLE IF NOT EXISTS `cell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellName` varchar(255) DEFAULT NULL,
  `CsgInd` varchar(255) DEFAULT NULL,
  `UlCyclicPrefix` varchar(255) DEFAULT NULL,
  `DlCyclicPrefix` varchar(255) DEFAULT NULL,
  `FreqBand` float DEFAULT NULL,
  `UlEarfcnCfgInd` varchar(255) DEFAULT NULL,
  `DlEarfcn` float DEFAULT NULL,
  `UlBandWidth` varchar(255) DEFAULT NULL,
  `DlBandWidth` varchar(255) DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `PhyCellId` float DEFAULT NULL,
  `AdditionalSpectrumEmission` float DEFAULT NULL,
  `CellActiveState` varchar(255) DEFAULT NULL,
  `CellAdminState` varchar(255) DEFAULT NULL,
  `FddTddInd` varchar(255) DEFAULT NULL,
  `CellSpecificOffset` varchar(255) DEFAULT NULL,
  `QoffsetFreq` varchar(255) DEFAULT NULL,
  `RootSequenceIdx` float DEFAULT NULL,
  `HighSpeedFlag` varchar(255) DEFAULT NULL,
  `PreambleFmt` float DEFAULT NULL,
  `CellRadius` float DEFAULT NULL,
  `CustomizedBandWidthCfgInd` varchar(255) DEFAULT NULL,
  `EmergencyAreaIdCfgInd` varchar(255) DEFAULT NULL,
  `UePowerMaxCfgInd` varchar(255) DEFAULT NULL,
  `MultiRruCellFlag` varchar(255) DEFAULT NULL,
  `CPRICompression` varchar(255) DEFAULT NULL,
  `AirCellFlag` varchar(255) DEFAULT NULL,
  `CrsPortNum` varchar(255) DEFAULT NULL,
  `TxRxMode` varchar(255) DEFAULT NULL,
  `UserLabel` varchar(255) DEFAULT NULL,
  `WorkMode` varchar(255) DEFAULT NULL,
  `EuCellStandbyMode` varchar(255) DEFAULT NULL,
  `SfnMasterCellLabel` varchar(255) DEFAULT NULL,
  `CellSlaveBand` varchar(255) DEFAULT NULL,
  `CnOpSharingGroupId` float DEFAULT NULL,
  `IntraFreqRanSharingInd` varchar(255) DEFAULT NULL,
  `IntraFreqAnrInd` varchar(255) DEFAULT NULL,
  `CellScaleInd` varchar(255) DEFAULT NULL,
  `FreqPriorityForAnr` float DEFAULT NULL,
  `CellRadiusStartLocation` float DEFAULT NULL,
  `SpecifiedCellFlag` varchar(255) DEFAULT NULL,
  `MultiCellShareMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `NbCellFlag` varchar(255) DEFAULT NULL,
  `SubframeAssignment` varchar(255) DEFAULT NULL,
  `SpecialSubframePatterns` varchar(255) DEFAULT NULL,
  `CrsPortMap` varchar(255) DEFAULT NULL,
  `CsiRsPeriod` varchar(255) DEFAULT NULL,
  `UlEarfcn` float DEFAULT NULL,
  `UePowerMax` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellacbar
CREATE TABLE IF NOT EXISTS `cellacbar` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `AcBarringInfoCfgInd` varchar(255) DEFAULT NULL,
  `VoLTEPreferCfgInd` varchar(255) DEFAULT NULL,
  `AcBarringForEmergency` varchar(255) DEFAULT NULL,
  `AcBarringForMoDataCfgInd` varchar(255) DEFAULT NULL,
  `AcBarringFactorForCall` varchar(255) DEFAULT NULL,
  `AcBarTimeForCall` varchar(255) DEFAULT NULL,
  `AC11BarforCall` varchar(255) DEFAULT NULL,
  `AC12BarforCall` varchar(255) DEFAULT NULL,
  `AC13BarforCall` varchar(255) DEFAULT NULL,
  `AC14BarforCall` varchar(255) DEFAULT NULL,
  `AC15BarforCall` varchar(255) DEFAULT NULL,
  `AcBarringForMoSigCfgInd` varchar(255) DEFAULT NULL,
  `AcBarringFactorForSig` varchar(255) DEFAULT NULL,
  `AcBarTimeForSig` varchar(255) DEFAULT NULL,
  `AC11BarForSig` varchar(255) DEFAULT NULL,
  `AC12BarForSig` varchar(255) DEFAULT NULL,
  `AC13BarForSig` varchar(255) DEFAULT NULL,
  `AC14BarForSig` varchar(255) DEFAULT NULL,
  `AC15BarForSig` varchar(255) DEFAULT NULL,
  `AcBarForMVoiceCfgInd` varchar(255) DEFAULT NULL,
  `AcBarForMVideoCfgInd` varchar(255) DEFAULT NULL,
  `AcBarForCsfbCfgInd` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellaccess
CREATE TABLE IF NOT EXISTS `cellaccess` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellBarred` varchar(255) DEFAULT NULL,
  `IntraFreqResel` varchar(255) DEFAULT NULL,
  `BroadcastMode` varchar(255) DEFAULT NULL,
  `RoundPeriod` float DEFAULT NULL,
  `RoundActionTime` varchar(255) DEFAULT NULL,
  `ReptSyncAvoidInd` varchar(255) DEFAULT NULL,
  `ReptSyncAvoidTime` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellalgoswitch
CREATE TABLE IF NOT EXISTS `cellalgoswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `VolteRedirectSwitch` longtext,
  `RachAlgoSwitch` longtext,
  `SrsAlgoSwitch` longtext,
  `PucchAlgoSwitch` longtext,
  `AqmAlgoSwitch` longtext,
  `CqiAdjAlgoSwitch` longtext,
  `DynAdjVoltSwitch` longtext,
  `RacAlgoSwitch` longtext,
  `MlbAlgoSwitch` longtext,
  `DlPcAlgoSwitch` longtext,
  `UlPcAlgoSwitch` longtext,
  `BfAlgoSwitch` longtext,
  `DlSchSwitch` longtext,
  `UlSchSwitch` longtext,
  `RanShareModeSwitch` longtext,
  `FreqPriorityHoSwitch` longtext,
  `MuBfAlgoSwitch` longtext,
  `DistBasedHoSwitch` longtext,
  `AcBarAlgoSwitch` longtext,
  `SfnUlSchSwitch` longtext,
  `SfnDlSchSwitch` longtext,
  `IrcSwitch` longtext,
  `DynDrxSwitch` longtext,
  `HighMobiTrigIdleModeSwitch` longtext,
  `AvoidInterfSwitch` longtext,
  `GLPwrShare` longtext,
  `EicicSwitch` longtext,
  `PUSCHMaxRBPUCCHAdjSwitch` longtext,
  `DlCompSwitch` longtext,
  `PsicSwitch` longtext,
  `MlbHoMode` longtext,
  `UplinkCompSwitch` longtext,
  `AntCalibrationAlgoSwitch` longtext,
  `DynSpectrumShareSwitch` longtext,
  `SfnLoadBasedAdptSwitch` longtext,
  `PuschIrcAlgoSwitch` longtext,
  `ReselecPriAdaptSwitch` longtext,
  `AnrFunctionSwitch` longtext,
  `SfnAlgoSwitch` longtext,
  `PrachIntrfRejSwitch` longtext,
  `EnhMIMOSwitch` longtext,
  `InterfRandSwitch` longtext,
  `RepeaterSwitch` longtext,
  `MultiFreqPriControlSwitch` longtext,
  `HarqAlgoSwitch` longtext,
  `CovBasedInterFreqHoMode` longtext,
  `LteUtcBroadcastSwitch` longtext,
  `CellSchStrategySwitch` longtext,
  `SsrdAlgoSwitch` longtext,
  `SfnUplinkCompSwitch` longtext,
  `LowSpeedInterFreqHoSwitch` longtext,
  `RelaySwitch` longtext,
  `InterFreqDirectHoSwitch` longtext,
  `PwrDeratSwitch` longtext,
  `DetectionAlgoSwitch` longtext,
  `PucchIRCEnhance` longtext,
  `AcBarAlgoforDynSwitch` longtext,
  `CreSwitch` longtext,
  `BackoffAlgoSwitch` longtext,
  `HoAllowedSwitch` longtext,
  `NCellClassMgtSw` longtext,
  `SpePCIBasedPolicySw` longtext,
  `CellPIMInterMitigSwitch` longtext,
  `PrachJointReceptionSwitch` longtext,
  `FeicicSwitch` longtext,
  `CamcSwitch` longtext,
  `RruUeMapSwitch` longtext,
  `HighSpeedSchOptSwitch` longtext,
  `ServiceDiffSwitch` longtext,
  `LtePttQoSSwitch` longtext,
  `SrsPucchEnhancedSwitch` longtext,
  `UEInactiveTimerQCI1Switch` longtext,
  `UlJRAntNumCombSw` longtext,
  `VamPhaseShiftAlgoSwitch` longtext,
  `AnrAlgoSwitch` longtext,
  `Dl256QamAlgoSwitch` longtext,
  `MlbAutoGroupSwitch` longtext,
  `CaAutoGroupSwitch` longtext,
  `AttachCellSelfCfgSwitch` longtext,
  `CellDlCoverEnhanceSwitch` longtext,
  `UlSchExtSwitch` longtext,
  `UdcAlgoSwitch` longtext,
  `VoLTESwitch` longtext,
  `VoipFailDecSelfRecSwitch` longtext,
  `DeprioritisationDeliverInd` longtext,
  `CmasSwitch` longtext,
  `MFBIAlgoSwitch` longtext,
  `PciConflictReportSwitch` longtext,
  `MroSwitch` longtext,
  `OpResourceGroupShareSwitch` longtext,
  `CellRecoverySwitch` longtext,
  `EnhancedMlbAlgoSwitch` longtext,
  `TrafficMlbSwitch` longtext,
  `MtcSwitch` longtext,
  `UlIcSwitch` longtext,
  `FcsMode` longtext,
  `ScVideoOptSwitch` longtext,
  `SpidSpecificHoSwitch` longtext,
  `DMIMOAlgoSwitch` longtext,
  `UlAmrcMode` longtext,
  `AmrcAlgoSwitch` longtext,
  `CrsIcSwitch` longtext,
  `MeasOptAlgoSwitch` longtext,
  `FreqLayerSwitch` longtext,
  `EmimoSwitch` longtext,
  `RohcSwitch` longtext,
  `McpttSwitch` longtext,
  `UlIcsRbRatio` float DEFAULT NULL,
  `UlIcsVoLTEPLTh` float DEFAULT NULL,
  `DacqSwitch` longtext,
  `RrcReestDataFwdSwitch` longtext,
  `MTCCongControlSwitch` varchar(255) DEFAULT NULL,
  `MTCPowerSavSwitch` varchar(255) DEFAULT NULL,
  `TcpCtrlSwitch` varchar(255) DEFAULT NULL,
  `TurboReceiverSwitch` varchar(255) DEFAULT NULL,
  `NaicsSwitch` varchar(255) DEFAULT NULL,
  `UplinkIcSwitch` varchar(255) DEFAULT NULL,
  `DacqEnhancementSwitch` varchar(255) DEFAULT NULL,
  `AsAlgoSwitch` varchar(255) DEFAULT NULL,
  `EnhChnCalSwitch` varchar(255) DEFAULT NULL,
  `NbCellAlgoSwitch` varchar(255) DEFAULT NULL,
  `MultiCnConnFreqPriSw` varchar(255) DEFAULT NULL,
  `NoSrsSccBfAlgoSwitch` varchar(255) DEFAULT NULL,
  `SpecUserAlgoSwitch` varchar(255) DEFAULT NULL,
  `MimoUePaAdaptSw` varchar(255) DEFAULT NULL,
  `UlSuMimoAlgoSwitch` varchar(255) DEFAULT NULL,
  `HighSpeedInterRatHoSwitch` varchar(255) DEFAULT NULL,
  `EmcSpsSchSwitch` varchar(255) DEFAULT NULL,
  `LbsSwitch` varchar(255) DEFAULT NULL,
  `DtxDetectionAlgoSwitch` varchar(255) DEFAULT NULL,
  `DlSchExtSwitch` varchar(255) DEFAULT NULL,
  `MPMUDetectSwitch` varchar(255) DEFAULT NULL,
  `VmsSwitch` varchar(255) DEFAULT NULL,
  `SmallBandOptSwitch` varchar(255) DEFAULT NULL,
  `IdleModeEdrxSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellbackoff
CREATE TABLE IF NOT EXISTS `cellbackoff` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterFreqBackoffPrbThd` float DEFAULT NULL,
  `InterFreqBackoffUeNumThd` float DEFAULT NULL,
  `RBPriMcsSelectTrigPrbThd` float DEFAULT NULL,
  `RBPriMcsSelectStopPrbThd` float DEFAULT NULL,
  `LowestEffDlMcsThd` float DEFAULT NULL,
  `LowEffDlMcsThd` float DEFAULT NULL,
  `LowEffDlFactorA` float DEFAULT NULL,
  `LowEffDlFactorK` float DEFAULT NULL,
  `LowestEffUlMcsThd` float DEFAULT NULL,
  `LowEffUlMcsThd` float DEFAULT NULL,
  `LowEffUlFactorA` float DEFAULT NULL,
  `LowEffUlFactorK` float DEFAULT NULL,
  `UlHeavyTrafficTtiProporThd` float DEFAULT NULL,
  `UlTrafficMlbUeNumThd` float DEFAULT NULL,
  `UlHeavyTrafficJudgePeriod` float DEFAULT NULL,
  `HoInRejectPrbThd` float DEFAULT NULL,
  `HoInRejectUeNumThd` float DEFAULT NULL,
  `RejectedHoInCause` varchar(255) DEFAULT NULL,
  `BackoffCAUserHOSw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellbf
CREATE TABLE IF NOT EXISTS `cellbf` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MaxBfRankPara` varchar(255) DEFAULT NULL,
  `DualLayerBFAlgType` varchar(255) DEFAULT NULL,
  `HighOrderMubfMaxLayer` varchar(255) DEFAULT NULL,
  `HighOrderMubfPairRule` varchar(255) DEFAULT NULL,
  `AntSelUEMubfPairMode` varchar(255) DEFAULT NULL,
  `NonAntSelUEMubfPairMode` varchar(255) DEFAULT NULL,
  `MovingUeMuBfScheme` varchar(255) DEFAULT NULL,
  `QualUEPortAvoidMode` varchar(255) DEFAULT NULL,
  `WaitPairingLayerThd` float DEFAULT NULL,
  `MassiveMIMOMubfPairRule` varchar(255) DEFAULT NULL,
  `MultiLayerThdSwitchToTM7` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellbfmimoparacfg
CREATE TABLE IF NOT EXISTS `cellbfmimoparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BfMimoAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `BfMimoAdapWithoutTm2` varchar(255) DEFAULT NULL,
  `BfSingleToDualThdOffset` float DEFAULT NULL,
  `TM3Rank2ToDualBfThdOffset` float DEFAULT NULL,
  `AsUeDualBfToTM3Rank2Offset` float DEFAULT NULL,
  `AsUeTM3Rank2ToDualBfOffset` float DEFAULT NULL,
  `AsUeBfSingleToDualOffset` float DEFAULT NULL,
  `DualBfToTM3Rank2Offset` float DEFAULT NULL,
  `OffsetOfInStat` float DEFAULT NULL,
  `OffsetOfOutStat` float DEFAULT NULL,
  `AntBasedBfMimoAlgoSelect` varchar(255) DEFAULT NULL,
  `Tm3DirectToDualBfSwitch` varchar(255) DEFAULT NULL,
  `SccBfMimoAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `HighLowSpeedUeThdOffset` float DEFAULT NULL,
  `TmAccelerationSwitch` varchar(255) DEFAULT NULL,
  `BfMimoAdapWithTm4Switch` varchar(255) DEFAULT NULL,
  `Tm3AndTm9ThdOffset` float DEFAULT NULL,
  `BfTo2LayerMubfThdOffset` float DEFAULT NULL,
  `BfTo4LayerMubfThdOffset` float DEFAULT NULL,
  `Tm3Rank2ToTm9Rank4Offset` float DEFAULT NULL,
  `Tm9Rank4ToTm3Rank2Offset` float DEFAULT NULL,
  `VolteMimoAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `HoBfThdAdjSwitch` varchar(255) DEFAULT NULL,
  `CaSccAddThldOffset` float DEFAULT NULL,
  `CaSccDelThldOffset` float DEFAULT NULL,
  `BfMimoAlgoOptSwitch` varchar(255) DEFAULT NULL,
  `InitialBfMimoType` varchar(255) DEFAULT NULL,
  `MultiLayerPairIsoThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcecfg
CREATE TABLE IF NOT EXISTS `cellcecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `RachRsrpFstThd` float DEFAULT NULL,
  `RachRsrpSndThd` float DEFAULT NULL,
  `RachRsrpTrdThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellceschcfg
CREATE TABLE IF NOT EXISTS `cellceschcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `SIB1RepNum` varchar(255) DEFAULT NULL,
  `PdschMaxNumRepModeA` varchar(255) DEFAULT NULL,
  `PdschMaxNumRepModeB` varchar(255) DEFAULT NULL,
  `PuschMaxNumRepModeA` varchar(255) DEFAULT NULL,
  `PuschMaxNumRepModeB` varchar(255) DEFAULT NULL,
  `MpdcchMaxNumRepPaging` varchar(255) DEFAULT NULL,
  `MpdcchMaxNumRepModeA` varchar(255) DEFAULT NULL,
  `MpdcchMaxNumRepModeB` varchar(255) DEFAULT NULL,
  `SiTransEcr` float DEFAULT NULL,
  `PagingGroupNum` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellchpwrcfg
CREATE TABLE IF NOT EXISTS `cellchpwrcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PcfichPwr` float DEFAULT NULL,
  `PbchPwr` float DEFAULT NULL,
  `SchPwr` float DEFAULT NULL,
  `DbchPwr` float DEFAULT NULL,
  `PchPwr` float DEFAULT NULL,
  `RaRspPwr` float DEFAULT NULL,
  `PrsPwr` float DEFAULT NULL,
  `AntOutputPwr` float DEFAULT NULL,
  `PmchPwrOffset` float DEFAULT NULL,
  `HoRarPwrEnhancedSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcounterparagroup
CREATE TABLE IF NOT EXISTS `cellcounterparagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellUserNumPdfThd` float DEFAULT NULL,
  `UlIblerOptSwitch` varchar(255) DEFAULT NULL,
  `CellCounterAlgoSwitch` varchar(255) DEFAULT NULL,
  `EdgeUserA3Offset` float DEFAULT NULL,
  `DlThrpPdfThd0` float DEFAULT NULL,
  `DlThrpPdfThd1` float DEFAULT NULL,
  `DlThrpPdfThd2` float DEFAULT NULL,
  `DlThrpPdfThd3` float DEFAULT NULL,
  `DlThrpPdfThd4` float DEFAULT NULL,
  `DlThrpPdfThd5` float DEFAULT NULL,
  `DlThrpPdfThd6` float DEFAULT NULL,
  `DlThrpPdfThd7` float DEFAULT NULL,
  `DlThrpPdfThd8` float DEFAULT NULL,
  `UlThrpPdfThd0` float DEFAULT NULL,
  `UlThrpPdfThd1` float DEFAULT NULL,
  `UlThrpPdfThd2` float DEFAULT NULL,
  `UlThrpPdfThd3` float DEFAULT NULL,
  `UlThrpPdfThd4` float DEFAULT NULL,
  `UlThrpPdfThd5` float DEFAULT NULL,
  `UlThrpPdfThd6` float DEFAULT NULL,
  `UlThrpPdfThd7` float DEFAULT NULL,
  `UlThrpPdfThd8` float DEFAULT NULL,
  `UeAbnormalJudgeThd` float DEFAULT NULL,
  `ThrPdfPolicy` varchar(255) DEFAULT NULL,
  `EdgeUserServRSRPThd` float DEFAULT NULL,
  `EdgeUserSRSOffset` float DEFAULT NULL,
  `ChrOutputModeSwitch` varchar(255) DEFAULT NULL,
  `DlUserXpUnsatiThd` float DEFAULT NULL,
  `DlTrafficVolumeThd` float DEFAULT NULL,
  `DlPktDelayMeasPolicy` varchar(255) DEFAULT NULL,
  `QCI1ContinuousPktLossThld` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcqiadaptivecfg
CREATE TABLE IF NOT EXISTS `cellcqiadaptivecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CqiPeriodAdaptive` varchar(255) DEFAULT NULL,
  `SimulAckNackAndCqiSwitch` varchar(255) DEFAULT NULL,
  `HoAperiodicCqiCfgSwitch` varchar(255) DEFAULT NULL,
  `MinCqiPeriod` varchar(255) DEFAULT NULL,
  `SccCqiRptEnhancedSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcqiadjalgo
CREATE TABLE IF NOT EXISTS `cellcqiadjalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InitDlIblerTarget` float DEFAULT NULL,
  `InitDeltaCqi` float DEFAULT NULL,
  `CqiAdjStep` float DEFAULT NULL,
  `CqiFilterCoefForMcs` float DEFAULT NULL,
  `VolteNackDeltaCqi` float DEFAULT NULL,
  `DlVolteCqiAdjOptSw` varchar(255) DEFAULT NULL,
  `CqiOptSwitch` varchar(255) DEFAULT NULL,
  `InitDlIblerTargetforQCI2` float DEFAULT NULL,
  `InitDlIblerTargetforVoLTE` float DEFAULT NULL,
  `CellDeltaCqiSampSelThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcsirsparacfg
CREATE TABLE IF NOT EXISTS `cellcsirsparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CsiRsSwitch` varchar(255) DEFAULT NULL,
  `CsiRsPeriod` varchar(255) DEFAULT NULL,
  `CsiRsConfigUserNumTh` float DEFAULT NULL,
  `CsiRsUnconfigUserNumTh` float DEFAULT NULL,
  `CsiRsConfigUserRatioTh` float DEFAULT NULL,
  `CsiRsUnconfigUserRatioTh` float DEFAULT NULL,
  `CsiRsSetJudgeHysTimer` float DEFAULT NULL,
  `CsiRsSetJudgeTimer` float DEFAULT NULL,
  `CsiRsPortNum` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellcspcpara
CREATE TABLE IF NOT EXISTS `cellcspcpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `eCspcPCAdjUeNumTh` float DEFAULT NULL,
  `CellCspcSwitch` varchar(255) DEFAULT NULL,
  `CspcRapidRptSwitch` varchar(255) DEFAULT NULL,
  `CspcUeSrsCfgRptPeriod` float DEFAULT NULL,
  `UlRsrpRptPeriod` float DEFAULT NULL,
  `CspcCqiFilterCoeff` float DEFAULT NULL,
  `UlRsrpPhyFilterCoeff` float DEFAULT NULL,
  `UlRsrpRrmFilterCoeff` float DEFAULT NULL,
  `eCspcA3Offset` float DEFAULT NULL,
  `CelleCspcSwitch` varchar(255) DEFAULT NULL,
  `eCspcPCAdjRange` varchar(255) DEFAULT NULL,
  `IntraEnbCspcSw` varchar(255) DEFAULT NULL,
  `ElasticCarrierSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldacqcfg
CREATE TABLE IF NOT EXISTS `celldacqcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlAmbrLimitFirstRank` float DEFAULT NULL,
  `DlAmbrLimitSecRank` float DEFAULT NULL,
  `UlAmbrLimitFirstRank` float DEFAULT NULL,
  `UlAmbrLimitSecRank` float DEFAULT NULL,
  `DlPrbUsageFirstRank` float DEFAULT NULL,
  `DlPrbUsageSecRank` float DEFAULT NULL,
  `UlPrbUsageFirstRank` float DEFAULT NULL,
  `UlPrbUsageSecRank` float DEFAULT NULL,
  `DacqExecutionDuration` float DEFAULT NULL,
  `DacqCongMonDur` float DEFAULT NULL,
  `AmbrLimitSchUserNum` float DEFAULT NULL,
  `AmbrProtectSchUserNum` float DEFAULT NULL,
  `DlAmbrLimitThirdRank` float DEFAULT NULL,
  `DlPrbUsageThirdRank` float DEFAULT NULL,
  `UlAmbrLimitThirdRank` float DEFAULT NULL,
  `UlPrbUsageThirdRank` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlcompalgo
CREATE TABLE IF NOT EXISTS `celldlcompalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlCompA3Offset` float DEFAULT NULL,
  `DpsCoCellUserRatioThd` float DEFAULT NULL,
  `DpsLoadDiffThd` float DEFAULT NULL,
  `DpsServingCellDlPrbThd` float DEFAULT NULL,
  `HetnetDpsCoCellRsrpThd` float DEFAULT NULL,
  `HomnetDpsCoCellRsrpThd` float DEFAULT NULL,
  `JtCoCellDlPrbThd` float DEFAULT NULL,
  `JtCoCellRsrpThd` float DEFAULT NULL,
  `JtServingCellDlPrbThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlicic
CREATE TABLE IF NOT EXISTS `celldlicic` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BandMode` varchar(255) DEFAULT NULL,
  `DlIcicUserAttrGfactorThd` float DEFAULT NULL,
  `DlIcicNoiseUserRsrpThd` float DEFAULT NULL,
  `AIcIcPlusA3Offset` float DEFAULT NULL,
  `AIcicPlusPCAdjRange` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlpcpdcch
CREATE TABLE IF NOT EXISTS `celldlpcpdcch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DediDciPwrOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlpcpdsch
CREATE TABLE IF NOT EXISTS `celldlpcpdsch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CcuPa` varchar(255) DEFAULT NULL,
  `CeuPa` varchar(255) DEFAULT NULL,
  `CellEdgLoadEvalPrd` float DEFAULT NULL,
  `NeighCellEdgLoadThd` float DEFAULT NULL,
  `ExceedNCellEdgLoadThdRatio` float DEFAULT NULL,
  `BFUserPwrBoostPrd` float DEFAULT NULL,
  `BFUserAdptPwrA3Offset` float DEFAULT NULL,
  `BFUserAdptPwrA6Offset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlpcpdschpa
CREATE TABLE IF NOT EXISTS `celldlpcpdschpa` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PdschPaAdjSwitch` varchar(255) DEFAULT NULL,
  `PaPcOff` varchar(255) DEFAULT NULL,
  `NomPdschRsEpreOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlpcphich
CREATE TABLE IF NOT EXISTS `celldlpcphich` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PwrOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldlschalgo
CREATE TABLE IF NOT EXISTS `celldlschalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlschStrategy` varchar(255) DEFAULT NULL,
  `FreeUserDlRbUsedRatio` float DEFAULT NULL,
  `MaxMimoRankPara` varchar(255) DEFAULT NULL,
  `CaSchStrategy` varchar(255) DEFAULT NULL,
  `NonGbrResourceRatio` float DEFAULT NULL,
  `DlEpfCapacityFactor` varchar(255) DEFAULT NULL,
  `RBPriMcsSelectRatioThd` float DEFAULT NULL,
  `CellEdgeUserRbAllocMode` varchar(255) DEFAULT NULL,
  `DlIcicSchMode` varchar(255) DEFAULT NULL,
  `RbgAllocStrategy` varchar(255) DEFAULT NULL,
  `MidUserMcsThreshold` float DEFAULT NULL,
  `DlLowLoadSdmaThdOffset` float DEFAULT NULL,
  `DlHarqMaxTxNum` float DEFAULT NULL,
  `NonGbrSchDelayWeight` float DEFAULT NULL,
  `RBDamageNearPointIblerTh` float DEFAULT NULL,
  `BreathingPilotAlgoSwitch` varchar(255) DEFAULT NULL,
  `HoStaticMcsTimer` float DEFAULT NULL,
  `DlRankOptimizeSwitch` varchar(255) DEFAULT NULL,
  `CaSccDopMeas` varchar(255) DEFAULT NULL,
  `DlSpsInterval` varchar(255) DEFAULT NULL,
  `RBPriMcsSelectStrategy` varchar(255) DEFAULT NULL,
  `DlRankDetectSwitch` varchar(255) DEFAULT NULL,
  `MbsfnSfCfg` varchar(255) DEFAULT NULL,
  `TddSpecSubfSchSwitch` varchar(255) DEFAULT NULL,
  `DataThdInPdcchPdschBal` float DEFAULT NULL,
  `UeNumThdInPdcchPdschBal` float DEFAULT NULL,
  `EnbInterfRandMod` varchar(255) DEFAULT NULL,
  `TxdDci1aSwitch` varchar(255) DEFAULT NULL,
  `DlSchAbnUeThd` float DEFAULT NULL,
  `RarAndPagingCR` float DEFAULT NULL,
  `CqiAdjInitialStep` float DEFAULT NULL,
  `CsiRsSfSchStrSwitch` varchar(255) DEFAULT NULL,
  `SfnDlLoadPeriod` float DEFAULT NULL,
  `FreqSelJudgePeriod` float DEFAULT NULL,
  `DlEnhancedVoipSchSw` varchar(255) DEFAULT NULL,
  `FDUEEnhAperCQITrigPeriod` varchar(255) DEFAULT NULL,
  `SfnDlHighLoadThd` float DEFAULT NULL,
  `SfnDlLowLoadThd` float DEFAULT NULL,
  `DlHighLoadSdmaThdOffset` float DEFAULT NULL,
  `LbtSwitch` varchar(255) DEFAULT NULL,
  `LbtNiHighThreshold` float DEFAULT NULL,
  `LbtNiLowThreshold` float DEFAULT NULL,
  `EnAperiodicCqiTrigStrategy` varchar(255) DEFAULT NULL,
  `DlCaUeCapAllocStrategy` varchar(255) DEFAULT NULL,
  `DlCaUserSchSwitch` varchar(255) DEFAULT NULL,
  `Dl256QamCqiTblAdpPeriod` float DEFAULT NULL,
  `Dl256QamCqiTblCfgStrategy` varchar(255) DEFAULT NULL,
  `BenefitMcsThreshold` float DEFAULT NULL,
  `DlSpsMcsDecreaseIblerThd` float DEFAULT NULL,
  `DlAfcMaxFreq` float DEFAULT NULL,
  `SrsRsrpFilterCoefForDlAfc` float DEFAULT NULL,
  `RlcMacJointSchSw` varchar(255) DEFAULT NULL,
  `NoSchStopACqiThd` float DEFAULT NULL,
  `InterPoleRruIdThld` float DEFAULT NULL,
  `IntraPoleRruIdDuration` float DEFAULT NULL,
  `IntraPoleRRUIdPeriod` float DEFAULT NULL,
  `NgbRruIdThld` float DEFAULT NULL,
  `UserSpeedDlSchPriWeight` float DEFAULT NULL,
  `EmimoCellRbRatio` float DEFAULT NULL,
  `EmimoJointSchThd` float DEFAULT NULL,
  `EmimoIndependentSchThd` float DEFAULT NULL,
  `EmimoJointIndepThdOffset` float DEFAULT NULL,
  `StartRankDetectEffThd` float DEFAULT NULL,
  `StartRankDetectCntThd` float DEFAULT NULL,
  `RankDetectSuccessCntThd` float DEFAULT NULL,
  `RankReverseDetectCntThd` float DEFAULT NULL,
  `AmbrCtrlTcycle` varchar(255) DEFAULT NULL,
  `CaSchWeight` float DEFAULT NULL,
  `PosNegFoRatioThd` float DEFAULT NULL,
  `UserFoThdForDlAfc` float DEFAULT NULL,
  `ZoneUserNumRatioThd` float DEFAULT NULL,
  `CrsFullBwPostODTtiNum` float DEFAULT NULL,
  `CrsFullBwPreODTtiNum` float DEFAULT NULL,
  `CrsFullBwPreSibPagingTtiNum` float DEFAULT NULL,
  `UeAttJudgePeriod` varchar(255) DEFAULT NULL,
  `UeAttJudgeRsrpHyst` float DEFAULT NULL,
  `BPUeOdAlignThd` float DEFAULT NULL,
  `PrbEnableThldVideoResCtrl` float DEFAULT NULL,
  `IblerPdcchSinrOffset` float DEFAULT NULL,
  `FSUEAperCQITrigPeriod` float DEFAULT NULL,
  `FSUESbCQIValidityPeriod` float DEFAULT NULL,
  `InterfRandPciOffset` float DEFAULT NULL,
  `PilotOffStrategy` varchar(255) DEFAULT NULL,
  `CrsCompensatePeriod` float DEFAULT NULL,
  `DlFirstHarqTxTbsIncNum` float DEFAULT NULL,
  `NbDlHarqMaxTxCount` float DEFAULT NULL,
  `CongestMaxVideoRate` float DEFAULT NULL,
  `UserEnableThldVideoResCtrl` float DEFAULT NULL,
  `PilotOffDrxParaGroupId` float DEFAULT NULL,
  `MccsCqiPeriodThld` varchar(255) DEFAULT NULL,
  `MccsPrbRateThld` float DEFAULT NULL,
  `MultiCarrierCoSchAlgoSw` varchar(255) DEFAULT NULL,
  `OverlapRsrpIsolationThd` float DEFAULT NULL,
  `CbtcDlSchedulingMcsUpLimit` float DEFAULT NULL,
  `DlEnhMuCellRbRatioThld` float DEFAULT NULL,
  `DlEnhMuCellAvgCQIThd` float DEFAULT NULL,
  `DlEnhMuPrbPairRatioThld` float DEFAULT NULL,
  `DlEnhMuOnOffTimer` float DEFAULT NULL,
  `HighIblerTargetTbsIdxThld` float DEFAULT NULL,
  `LowIblerTargetTbsIdxThld` float DEFAULT NULL,
  `DlCandPairUeCntThld` float DEFAULT NULL,
  `PdschMaxCodeRate` varchar(255) DEFAULT NULL,
  `IndependentBeamPmiThld` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldrxpara
CREATE TABLE IF NOT EXISTS `celldrxpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `FddEnterDrxThd` float DEFAULT NULL,
  `FddExitDrxThd` float DEFAULT NULL,
  `TddEnterDrxThdUl` float DEFAULT NULL,
  `TddExitDrxThdUl` float DEFAULT NULL,
  `TddEnterDrxThdDl` float DEFAULT NULL,
  `TddExitDrxThdDl` float DEFAULT NULL,
  `DataAmountStatTimer` float DEFAULT NULL,
  `TddPowerSaveMeasSwitch` varchar(255) DEFAULT NULL,
  `TddPowerSaveStatTimer` float DEFAULT NULL,
  `TddPowerSavingEnterDrxThd` float DEFAULT NULL,
  `TddPowerSavingExitDrxThd` float DEFAULT NULL,
  `LongDrxCycleUnsync` varchar(255) DEFAULT NULL,
  `CqiMask` varchar(255) DEFAULT NULL,
  `OndurationTimerUnsync` varchar(255) DEFAULT NULL,
  `DrxInactivityTimerUnsync` varchar(255) DEFAULT NULL,
  `DrxPolicyMode` varchar(255) DEFAULT NULL,
  `DrxStartOffsetOptSwitch` varchar(255) DEFAULT NULL,
  `DrxRcvDtxProSwitch` varchar(255) DEFAULT NULL,
  `DrxForMeasSwitch` varchar(255) DEFAULT NULL,
  `LongDrxCycleForMeas` varchar(255) DEFAULT NULL,
  `OnDurTimerForMeas` varchar(255) DEFAULT NULL,
  `DrxInactTimerForMeas` varchar(255) DEFAULT NULL,
  `DrxReTxTimerForMeas` varchar(255) DEFAULT NULL,
  `ShortDrxSwForMeas` varchar(255) DEFAULT NULL,
  `ShortDrxCycleForMeas` varchar(255) DEFAULT NULL,
  `ShortCycleTimerForMeas` float DEFAULT NULL,
  `DrxSrDetectOptSwitch` varchar(255) DEFAULT NULL,
  `DrxStartoffsetAdjustSW` varchar(255) DEFAULT NULL,
  `MeasDrxSpecSchSwitch` varchar(255) DEFAULT NULL,
  `CovGsmMeasDrxCfgSwitch` varchar(255) DEFAULT NULL,
  `OnDurationUnextendSwitch` varchar(255) DEFAULT NULL,
  `DrxStateDuringUlHarqRetx` varchar(255) DEFAULT NULL,
  `DrxAlgSwitch` varchar(255) DEFAULT NULL,
  `ShortDrxCycleSwitch` varchar(255) DEFAULT NULL,
  `VolteGapDrxExclusiveSwitch` varchar(255) DEFAULT NULL,
  `DrxStopSrPendingSw` varchar(255) DEFAULT NULL,
  `GapDrxExclusiveSwitch` varchar(255) DEFAULT NULL,
  `NbDrxInactivityTimer` varchar(255) DEFAULT NULL,
  `NbDrxReTxTimer` varchar(255) DEFAULT NULL,
  `NbDrxUlReTxTimer` varchar(255) DEFAULT NULL,
  `NbLongDrxCycle` varchar(255) DEFAULT NULL,
  `NbOnDurationTimer` varchar(255) DEFAULT NULL,
  `DrxOdPreSchSwitch` varchar(255) DEFAULT NULL,
  `DrxUeSrsOptSwitch` varchar(255) DEFAULT NULL,
  `SinrThldForVolteDrxCtrl` float DEFAULT NULL,
  `VoltePlrThldForExitingDrx` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldrxspecialpara
CREATE TABLE IF NOT EXISTS `celldrxspecialpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellDrxSpecialParaValid` varchar(255) DEFAULT NULL,
  `LongDrxCycleSpecial` varchar(255) DEFAULT NULL,
  `OnDurationTimerSpecial` varchar(255) DEFAULT NULL,
  `DrxInactivityTimerSpecial` varchar(255) DEFAULT NULL,
  `ShortDrxCycleSwitchSpecial` varchar(255) DEFAULT NULL,
  `LongDrxCycleForIntraRatAnr` varchar(255) DEFAULT NULL,
  `LongDrxCycleForInterRatAnr` varchar(255) DEFAULT NULL,
  `FddAnrDrxInactivityTimer` varchar(255) DEFAULT NULL,
  `TddAnrDrxInactivityTimer` varchar(255) DEFAULT NULL,
  `BfSpecificDrxParaSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldss
CREATE TABLE IF NOT EXISTS `celldss` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `HighFreqShareRbNum` float DEFAULT NULL,
  `LowFreqShareRbNum` float DEFAULT NULL,
  `ReMuteSwitch` varchar(255) DEFAULT NULL,
  `UlInterfRestrictionMode` varchar(255) DEFAULT NULL,
  `A3Offset` float DEFAULT NULL,
  `A6Offset` float DEFAULT NULL,
  `NearAreaSinrThd` float DEFAULT NULL,
  `MiddleAreaSinrThd` float DEFAULT NULL,
  `FarAreaSinrThd` float DEFAULT NULL,
  `InterfNCellConfigMode` varchar(255) DEFAULT NULL,
  `SpecShrPfmOptSwitch` varchar(255) DEFAULT NULL,
  `SinrThdWithoutGsmInterf` float DEFAULT NULL,
  `GsmInterfINThd` float DEFAULT NULL,
  `NearPtUserRsrpThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celldynacbaralgopara
CREATE TABLE IF NOT EXISTS `celldynacbaralgopara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DynAcBarStatPeriod` float DEFAULT NULL,
  `DynAcBarTriggerThd` float DEFAULT NULL,
  `DynAcBarCancelThd` float DEFAULT NULL,
  `SsacTriggerCondSatiPeriods` float DEFAULT NULL,
  `SsacCancelCondSatiPeriods` float DEFAULT NULL,
  `DisasterReferenceInd` varchar(255) DEFAULT NULL,
  `DisasterDuration` float DEFAULT NULL,
  `MoTriggerCondSatiPeriods` float DEFAULT NULL,
  `MoCancelCondSatiPeriods` float DEFAULT NULL,
  `MoFactorAdjStep` float DEFAULT NULL,
  `MoFactorRetreatStep` float DEFAULT NULL,
  `SsacFactorAdjStep` float DEFAULT NULL,
  `SsacFactorRetreatStep` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celleabalgopara
CREATE TABLE IF NOT EXISTS `celleabalgopara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `EABTriggerThd` float DEFAULT NULL,
  `EABStatPeriod` float DEFAULT NULL,
  `EABCategory` varchar(255) DEFAULT NULL,
  `EABCancelThd` float DEFAULT NULL,
  `EABCancelCondSatiPeriod` float DEFAULT NULL,
  `ABForExceptionData` varchar(255) DEFAULT NULL,
  `ABForSpecialAC` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celleicic
CREATE TABLE IF NOT EXISTS `celleicic` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `AbsPattern` varchar(255) DEFAULT NULL,
  `EicicMeasureEnable` varchar(255) DEFAULT NULL,
  `A3Offset` float DEFAULT NULL,
  `AbsAdjPeriod` float DEFAULT NULL,
  `SampleNumTarget` float DEFAULT NULL,
  `AttachCellAddThd` float DEFAULT NULL,
  `WorkMode` varchar(255) DEFAULT NULL,
  `StatPeriod` float DEFAULT NULL,
  `CreAdjPeriod` float DEFAULT NULL,
  `EicicConfigUserRatioThd` float DEFAULT NULL,
  `EicicUnConfigUserRatioThd` float DEFAULT NULL,
  `LargeCreMicroNumThd` float DEFAULT NULL,
  `EicicAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `NaicNCellStatPeriod` float DEFAULT NULL,
  `HighCoorCCEThd` float DEFAULT NULL,
  `HighCoorDTXThd` float DEFAULT NULL,
  `HighSpeedEnhancedOptSwitch` varchar(255) DEFAULT NULL,
  `HSHandinUserThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellemtcalgo
CREATE TABLE IF NOT EXISTS `cellemtcalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `EmtcUlRbTargetRatio` float DEFAULT NULL,
  `EmtcAperCqiTrigPrd` varchar(255) DEFAULT NULL,
  `EmtcDlRbTargetRatio` float DEFAULT NULL,
  `EmtcAlgoSwitch` varchar(255) DEFAULT NULL,
  `DlLteRsvNbCount` float DEFAULT NULL,
  `UlLteRsvNbCount` float DEFAULT NULL,
  `EmtcVolteSupportSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellhoparacfg
CREATE TABLE IF NOT EXISTS `cellhoparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BlindHoA1A2ThdRsrp` float DEFAULT NULL,
  `BlindHoA1A2ThdRsrq` float DEFAULT NULL,
  `HighSpeedThreshold` float DEFAULT NULL,
  `HoModeSwitch` varchar(255) DEFAULT NULL,
  `SrvccHoOptSwitch` varchar(255) DEFAULT NULL,
  `UlBadQualHoMcsThd` float DEFAULT NULL,
  `UlBadQualHoIblerThd` float DEFAULT NULL,
  `SpeedEvaluatedPeriod` varchar(255) DEFAULT NULL,
  `SpeedEvaluatedNum` float DEFAULT NULL,
  `L2UCsfbMRProMode` varchar(255) DEFAULT NULL,
  `CsfbMRWaitingTimer` float DEFAULT NULL,
  `CellHoAlgoSwitch` varchar(255) DEFAULT NULL,
  `SpeedEvaluatedNumForFastUe` float DEFAULT NULL,
  `HoUseInactiveTimerSwitch` varchar(255) DEFAULT NULL,
  `HSCellSleepUserNum` float DEFAULT NULL,
  `LowSpeedUsersSelProTimer` float DEFAULT NULL,
  `HSCellHoInProtectTimer` float DEFAULT NULL,
  `FlashSrvccSwitch` varchar(255) DEFAULT NULL,
  `UlPoorCoverPathLossThd` float DEFAULT NULL,
  `UlPoorCoverSinrThd` float DEFAULT NULL,
  `HoMrDelayTimerQci1` float DEFAULT NULL,
  `VoLTEQualityHoAlgoSwitch` varchar(255) DEFAULT NULL,
  `Qci1PlrThldForVolteQualHo` float DEFAULT NULL,
  `VolteQualPktLossPeriod` float DEFAULT NULL,
  `CovBasedIntraRatMeasTime` float DEFAULT NULL,
  `HighSpeedUeJudgeMode` varchar(255) DEFAULT NULL,
  `RingingMsgCheckSw` varchar(255) DEFAULT NULL,
  `SrvccMrDelayTimer` float DEFAULT NULL,
  `VolteQualIfHoQci1PlrThld` float DEFAULT NULL,
  `VolteQualRecoveryQci1PlrThld` float DEFAULT NULL,
  `InterRatUuHoFailRetryTimes` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellidprdupt
CREATE TABLE IF NOT EXISTS `cellidprdupt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PrdUptSwitch` varchar(255) DEFAULT NULL,
  `ActionTime` varchar(255) DEFAULT NULL,
  `Period` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celliicspara
CREATE TABLE IF NOT EXISTS `celliicspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `IicsUserAttrA3Offset` float DEFAULT NULL,
  `IicsUserAttrThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celllowpower
CREATE TABLE IF NOT EXISTS `celllowpower` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `LowPwrSwitch` varchar(255) DEFAULT NULL,
  `CellUsedPwrReduceTimeLen` float DEFAULT NULL,
  `RsPwrReduceTimeLen` float DEFAULT NULL,
  `RfShutDownTimeLen` float DEFAULT NULL,
  `CellUsedPwrRatio` float DEFAULT NULL,
  `RsPwrAdjOffset` float DEFAULT NULL,
  `EnterTimeLen` float DEFAULT NULL,
  `BakPwrSavPolicy` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celllteflexbw
CREATE TABLE IF NOT EXISTS `celllteflexbw` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `LteFlexBwSwitch` varchar(255) DEFAULT NULL,
  `DlCustStartPrbIndex` float DEFAULT NULL,
  `DlCustEndPrbIndex` float DEFAULT NULL,
  `UlCustStartPrbIndex` float DEFAULT NULL,
  `UlCustEndPrbIndex` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celllteflexbwitfcfg
CREATE TABLE IF NOT EXISTS `celllteflexbwitfcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `HighItfGsmArfcn` float DEFAULT NULL,
  `GsmCarrierFreq` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmbfcs
CREATE TABLE IF NOT EXISTS `cellmbfcs` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellSyntheticalRateHoHyst` float DEFAULT NULL,
  `CellTrafficLoadThd` float DEFAULT NULL,
  `InterFreqMeasLoadRatioThd` float DEFAULT NULL,
  `InterFreqMeasMcsThd` float DEFAULT NULL,
  `InterFreqMeasServiceThd` float DEFAULT NULL,
  `MeasInfoUpdInterval` varchar(255) DEFAULT NULL,
  `UeSpectralEffHoHyst` float DEFAULT NULL,
  `LoadDiffRatioThld` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmbmscfg
CREATE TABLE IF NOT EXISTS `cellmbmscfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MBSFNSyncAreaId` float DEFAULT NULL,
  `MBMSSwitch` varchar(255) DEFAULT NULL,
  `ServiceAreaIdList` varchar(255) DEFAULT NULL,
  `MBMSServiceSwitch` varchar(255) DEFAULT NULL,
  `NCellMbmsCfgSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmcpara
CREATE TABLE IF NOT EXISTS `cellmcpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `A3Offset` float DEFAULT NULL,
  `Hysteresis` float DEFAULT NULL,
  `TimetoTrigger` varchar(255) DEFAULT NULL,
  `MaxReportCells` float DEFAULT NULL,
  `ReportAmount` varchar(255) DEFAULT NULL,
  `ReportInterval` varchar(255) DEFAULT NULL,
  `ReportQuantity` varchar(255) DEFAULT NULL,
  `TriggerQuantity` varchar(255) DEFAULT NULL,
  `A6Offset` float DEFAULT NULL,
  `IntraFreqMaxReportCells` float DEFAULT NULL,
  `IntraFreqTriggerQuantity` varchar(255) DEFAULT NULL,
  `IntraFreqReportQuantity` varchar(255) DEFAULT NULL,
  `InterFreqMaxReportCells` float DEFAULT NULL,
  `InterFreqTriggerQuantity` varchar(255) DEFAULT NULL,
  `InterFreqReportQuantity` varchar(255) DEFAULT NULL,
  `A3MCAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `AoaDetectThd` float DEFAULT NULL,
  `EcidReportAmount` varchar(255) DEFAULT NULL,
  `EcidReportInterval` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmimoparacfg
CREATE TABLE IF NOT EXISTS `cellmimoparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MimoAdaptiveSwitch` varchar(255) DEFAULT NULL,
  `FixedMimoMode` varchar(255) DEFAULT NULL,
  `InitialMimoType` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmlb
CREATE TABLE IF NOT EXISTS `cellmlb` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterFreqMlbThd` float DEFAULT NULL,
  `InterRatMlbThd` float DEFAULT NULL,
  `LoadOffset` float DEFAULT NULL,
  `LoadDiffThd` float DEFAULT NULL,
  `InterRatMlbUeNumThd` float DEFAULT NULL,
  `InitValidPeriod` float DEFAULT NULL,
  `LoadTransferFactor` float DEFAULT NULL,
  `MlbTriggerMode` varchar(255) DEFAULT NULL,
  `InterFreqMlbUeNumThd` float DEFAULT NULL,
  `MlbUeNumOffset` float DEFAULT NULL,
  `MlbMaxUeNum` float DEFAULT NULL,
  `MlbUeSelectPrbThd` float DEFAULT NULL,
  `DlDataMlbMode` varchar(255) DEFAULT NULL,
  `InterFreqMLBRanShareMode` varchar(255) DEFAULT NULL,
  `UeNumDiffThd` float DEFAULT NULL,
  `HotSpotUeMode` varchar(255) DEFAULT NULL,
  `InterFreqIdleMlbMode` varchar(255) DEFAULT NULL,
  `MlbMinUeNumThd` float DEFAULT NULL,
  `MlbMinUeNumOffset` float DEFAULT NULL,
  `InterFreqUeTrsfType` varchar(255) DEFAULT NULL,
  `InterFreqIdleMlbUeNumThd` float DEFAULT NULL,
  `InterRatIdleMlbUeNumThd` float DEFAULT NULL,
  `InterFreqLoadEvalPrd` float DEFAULT NULL,
  `InterRatLoadEvalPrd` float DEFAULT NULL,
  `FreqSelectStrategy` varchar(255) DEFAULT NULL,
  `LoadBalanceNCellScope` varchar(255) DEFAULT NULL,
  `InterRatMlbUeNumOffset` float DEFAULT NULL,
  `IdleUeSelFreqScope` varchar(255) DEFAULT NULL,
  `InterRatMlbUeSelStrategy` varchar(255) DEFAULT NULL,
  `InterRatMlbUeSelPrbThd` float DEFAULT NULL,
  `PrbLoadCalcMethod` varchar(255) DEFAULT NULL,
  `MlbUeSelectPunishTimer` float DEFAULT NULL,
  `MlbHoCellSelectStrategy` varchar(255) DEFAULT NULL,
  `InterRatMlbTriggerMode` varchar(255) DEFAULT NULL,
  `InterRatMlbUeNumModeThd` float DEFAULT NULL,
  `PunishJudgePrdNum` float DEFAULT NULL,
  `FreqPunishPrdNum` float DEFAULT NULL,
  `CellPunishPrdNum` float DEFAULT NULL,
  `MultiRruMode` varchar(255) DEFAULT NULL,
  `InterFreqOffloadOffset` float DEFAULT NULL,
  `InterFrqUeNumOffloadOffset` float DEFAULT NULL,
  `CellCapacityScaleFactor` float DEFAULT NULL,
  `InterRatMlbMaxUeNum` float DEFAULT NULL,
  `InterRatMlbHoFailPunish` varchar(255) DEFAULT NULL,
  `EutranLoadJudgeStrategy` varchar(255) DEFAULT NULL,
  `MlbTrigJudgePeriod` float DEFAULT NULL,
  `InterFreqMlbStrategy` varchar(255) DEFAULT NULL,
  `MaxSpectralEfficiencyValue` float DEFAULT NULL,
  `MinSpectralEfficiencyValue` float DEFAULT NULL,
  `SpectralEffAdjustMaxStep` float DEFAULT NULL,
  `UeNumDiffOffsetTransCa` float DEFAULT NULL,
  `MlbIdleUeNumAdjFactor` float DEFAULT NULL,
  `IdleMlbUeNumDiffThd` float DEFAULT NULL,
  `L2USmartOffloadThd` float DEFAULT NULL,
  `L2USmartOffloadOffset` float DEFAULT NULL,
  `InterFreqMlbUlThd` float DEFAULT NULL,
  `UeDlPrbLowThdOffset` float DEFAULT NULL,
  `UeUlPrbHighThdOffset` float DEFAULT NULL,
  `UeUlPrbLowThdOffset` float DEFAULT NULL,
  `VideoLoadHighThd` float DEFAULT NULL,
  `VideoLoadLowThd` float DEFAULT NULL,
  `VideoDlPrbThd` float DEFAULT NULL,
  `InterFIdleUeNumOffloadOfs` float DEFAULT NULL,
  `UlExperienceDiffThd` float DEFAULT NULL,
  `UlExperienceEvalPrd` float DEFAULT NULL,
  `UlExperienceMaxUeNum` float DEFAULT NULL,
  `UlExperienceOffloadThd` float DEFAULT NULL,
  `UlExperienceOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmlbautogroup
CREATE TABLE IF NOT EXISTS `cellmlbautogroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `OverlapIndAddThd` float DEFAULT NULL,
  `OverlapIndDelThd` float DEFAULT NULL,
  `StatPeriod` float DEFAULT NULL,
  `SleepPeriod` float DEFAULT NULL,
  `AddedMeasCfg` varchar(255) DEFAULT NULL,
  `OverlapIndAutoSampleNum` float DEFAULT NULL,
  `OptMode` varchar(255) DEFAULT NULL,
  `MicroOverlapIndAddThd` float DEFAULT NULL,
  `MicroOverlapIndDelThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmlbho
CREATE TABLE IF NOT EXISTS `cellmlbho` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `IdleUeSelFreqStrategy` varchar(255) DEFAULT NULL,
  `MlbHoInProtectMode` varchar(255) DEFAULT NULL,
  `MlbHoInProtectTimer` float DEFAULT NULL,
  `InterFreqMlbHoInA1ThdRsrp` float DEFAULT NULL,
  `InterFreqMlbHoInA1ThdRsrq` float DEFAULT NULL,
  `InterFreqMlbHoInA2ThdRsrp` float DEFAULT NULL,
  `InterFreqMlbHoInA2ThdRsrq` float DEFAULT NULL,
  `InterRatMlbStrategy` varchar(255) DEFAULT NULL,
  `MlbMatchOtherFeatureMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmlbuesel
CREATE TABLE IF NOT EXISTS `cellmlbuesel` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UeSelectPrbPrio` float DEFAULT NULL,
  `UeSelectQciPrio` float DEFAULT NULL,
  `UeSelectArpPrio` float DEFAULT NULL,
  `UeSelectDlMcsPrio` float DEFAULT NULL,
  `InterFreqMlbUeArpThd` float DEFAULT NULL,
  `InterFreqMlbUeDlMcsThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmmalgo
CREATE TABLE IF NOT EXISTS `cellmmalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BeamGainLimitSwitch` varchar(255) DEFAULT NULL,
  `MMAlgoOptSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellmro
CREATE TABLE IF NOT EXISTS `cellmro` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CioAdjLimitCfgInd` varchar(255) DEFAULT NULL,
  `CioAdjUpperLimit` varchar(255) DEFAULT NULL,
  `CioAdjLowerLimit` varchar(255) DEFAULT NULL,
  `InterFreqA2RsrpUpLimit` float DEFAULT NULL,
  `InterFreqA2RsrpLowLimit` float DEFAULT NULL,
  `A3InterFreqA2RsrpUpLimit` float DEFAULT NULL,
  `A3InterFreqA2RsrpLowLimit` float DEFAULT NULL,
  `InterFreqMroAdjParaSel` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellnoaccessalmpara
CREATE TABLE IF NOT EXISTS `cellnoaccessalmpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `NoAccessStartDetectTime` varchar(255) DEFAULT NULL,
  `NoAccessStopDetectTime` varchar(255) DEFAULT NULL,
  `NoAccessDetectTimer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellop
CREATE TABLE IF NOT EXISTS `cellop` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TrackingAreaId` float DEFAULT NULL,
  `CellReservedForOp` varchar(255) DEFAULT NULL,
  `OpUlRbUsedRatio` float DEFAULT NULL,
  `OpDlRbUsedRatio` float DEFAULT NULL,
  `MMECfgNum` varchar(255) DEFAULT NULL,
  `OpUeNumRatio` float DEFAULT NULL,
  `OpPttUlRbHighThd` float DEFAULT NULL,
  `OpPttUlRbLowThd` float DEFAULT NULL,
  `OpPttDlRbHighThd` float DEFAULT NULL,
  `OpPttDlRbLowThd` float DEFAULT NULL,
  `OperatorPlmnRoundPeriod` float DEFAULT NULL,
  `OpResourceGroupIndex` float DEFAULT NULL,
  `OpNonGbrResourceRatio` float DEFAULT NULL,
  `RatFreqPriorityGroupId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellpcalgo
CREATE TABLE IF NOT EXISTS `cellpcalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SrsPcStrategy` varchar(255) DEFAULT NULL,
  `PucchCloseLoopPcType` varchar(255) DEFAULT NULL,
  `PucchPcPeriod` float DEFAULT NULL,
  `PucchPcTargetSinrOffset` float DEFAULT NULL,
  `PuschRsrpHighThd` float DEFAULT NULL,
  `PuschIoTCtrlA3Offset` float DEFAULT NULL,
  `IoTNearPointOptSwitch` varchar(255) DEFAULT NULL,
  `IoTNearPointPLThreshold` varchar(255) DEFAULT NULL,
  `SrsPcSinrTarget` float DEFAULT NULL,
  `SrsPcRsrpTarget` float DEFAULT NULL,
  `PUSCHPsdCtrlTarget` varchar(255) DEFAULT NULL,
  `CloseLoopOptPUSCHType` varchar(255) DEFAULT NULL,
  `PUSCHOptIBLERThreshold` float DEFAULT NULL,
  `PUSCHPsdCtrlTargetForUs` varchar(255) DEFAULT NULL,
  `IoTCtrlINCorrectSwitch` varchar(255) DEFAULT NULL,
  `IoTCtrlEUPLThreshold` float DEFAULT NULL,
  `IoTCtrlNIThreshold` float DEFAULT NULL,
  `NearPointUePuschType` varchar(255) DEFAULT NULL,
  `VoLteSinrHighThdOffset` float DEFAULT NULL,
  `VoltePUSCHPowerOffset` float DEFAULT NULL,
  `PuschRsrpHighThdOffsetVoIP` float DEFAULT NULL,
  `DMSrsPcSinrOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellpdcchalgo
CREATE TABLE IF NOT EXISTS `cellpdcchalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ComSigCongregLv` varchar(255) DEFAULT NULL,
  `CceRatioAdjSwitch` varchar(255) DEFAULT NULL,
  `SfnPdcchDcsThd` float DEFAULT NULL,
  `InitPdcchSymNum` float DEFAULT NULL,
  `VirtualLoadPro` float DEFAULT NULL,
  `PdcchInitialCceAdjustValue` float DEFAULT NULL,
  `PdcchSymNumSwitch` varchar(255) DEFAULT NULL,
  `CceUseRatio` float DEFAULT NULL,
  `PdcchAggLvlCLAdjustSwitch` varchar(255) DEFAULT NULL,
  `DPDVirtualLoadSwitch` varchar(255) DEFAULT NULL,
  `DPDVirtualLoadType` varchar(255) DEFAULT NULL,
  `AggLvlSelStrageForDualCW` varchar(255) DEFAULT NULL,
  `PdcchCapacityImproveSwitch` varchar(255) DEFAULT NULL,
  `PdcchMaxCodeRate` float DEFAULT NULL,
  `ULDLPdcchSymNum` varchar(255) DEFAULT NULL,
  `PDCCHAggLvlAdaptStrage` varchar(255) DEFAULT NULL,
  `HysForCfiBasedPreSch` float DEFAULT NULL,
  `SfnPdcchSdmaThd` float DEFAULT NULL,
  `UlPdcchAllocImproveSwitch` varchar(255) DEFAULT NULL,
  `CceMaxInitialRatio` varchar(255) DEFAULT NULL,
  `PdcchPowerEnhancedSwitch` varchar(255) DEFAULT NULL,
  `PdcchBlerTarget` float DEFAULT NULL,
  `HLNetAccSigAggLvlSelEnhSw` varchar(255) DEFAULT NULL,
  `EpdcchAlgoSwitch` varchar(255) DEFAULT NULL,
  `CceThdtoEnableEpdcch` float DEFAULT NULL,
  `RbThdtoEnableEpdcch` float DEFAULT NULL,
  `CceThdtoDisableEpdcch` float DEFAULT NULL,
  `EcceThdtoDisableEpdcch` float DEFAULT NULL,
  `RbThdtoDisableEpdcch` float DEFAULT NULL,
  `HLNetAccSigAttempt` varchar(255) DEFAULT NULL,
  `EpdcchSfCfg` varchar(255) DEFAULT NULL,
  `UlDlEcceInitialRatioAdjSw` varchar(255) DEFAULT NULL,
  `CapacityBfOpt` varchar(255) DEFAULT NULL,
  `CCEThdEnableFlowCtrl` float DEFAULT NULL,
  `CCEThdDisableFlowCtrl` float DEFAULT NULL,
  `PDCCHPwrUpperLimitOptSw` varchar(255) DEFAULT NULL,
  `VoltePdcchSinrOffset` float DEFAULT NULL,
  `PdcchSparePowerAllocStrage` varchar(255) DEFAULT NULL,
  `PdcchDlAggLvlSlcEhnSwitch` varchar(255) DEFAULT NULL,
  `PdcchOutLoopAdjBaseStep` float DEFAULT NULL,
  `PdcchOutLoopAdjLowerLimit` float DEFAULT NULL,
  `PdcchAdjAlgoSwitch` varchar(255) DEFAULT NULL,
  `SplitBeamPdcchSdmaThd` float DEFAULT NULL,
  `NackDtxRatioThd` float DEFAULT NULL,
  `SplitBeamPdcchSdmaSw` varchar(255) DEFAULT NULL,
  `PdcchPwrCtrlUserNumThd` float DEFAULT NULL,
  `PdcchBfGainOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellpdcchcecfg
CREATE TABLE IF NOT EXISTS `cellpdcchcecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `PdcchMaxRepetitionCnt` varchar(255) DEFAULT NULL,
  `PdcchPeriodFactor` varchar(255) DEFAULT NULL,
  `PdcchTransRptCntFactor` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellprbvalmlb
CREATE TABLE IF NOT EXISTS `cellprbvalmlb` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PrbValMlbTrigThd` float DEFAULT NULL,
  `PrbValMlbAdmitThd` float DEFAULT NULL,
  `MlbUeSelectPrbValThd` float DEFAULT NULL,
  `PrbValFilterFactor` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellpucchalgo
CREATE TABLE IF NOT EXISTS `cellpucchalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SriLowLoadThd` float DEFAULT NULL,
  `SriReCfgInd` varchar(255) DEFAULT NULL,
  `SriAlgoSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellqcipara
CREATE TABLE IF NOT EXISTS `cellqcipara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `InterFreqHoGroupId` float DEFAULT NULL,
  `InterRatHoCdma1xRttGroupId` float DEFAULT NULL,
  `InterRatHoCdmaHrpdGroupId` float DEFAULT NULL,
  `InterRatHoCommGroupId` float DEFAULT NULL,
  `InterRatHoGeranGroupId` float DEFAULT NULL,
  `InterRatHoUtranGroupId` float DEFAULT NULL,
  `IntraFreqHoGroupId` float DEFAULT NULL,
  `DrxParaGroupId` float DEFAULT NULL,
  `SriPeriod` varchar(255) DEFAULT NULL,
  `RlfTimerConstCfgInd` varchar(255) DEFAULT NULL,
  `TrafficRelDelay` float DEFAULT NULL,
  `QciPriorityForHo` float DEFAULT NULL,
  `PreallocationParaGroupId` float DEFAULT NULL,
  `QciPriorityForDrx` float DEFAULT NULL,
  `QciAlgoSwitch` varchar(255) DEFAULT NULL,
  `QciEutranFreqRelationId` float DEFAULT NULL,
  `QciUtranFreqRelationId` float DEFAULT NULL,
  `QciGeranFreqRelationId` float DEFAULT NULL,
  `MlbQciGroupId` float DEFAULT NULL,
  `EmtcModeADrxParaGroupId` float DEFAULT NULL,
  `EmtcModeBDrxParaGroupId` float DEFAULT NULL,
  `ServiceFlag` varchar(255) DEFAULT NULL,
  `DlAmbrLimitFactor` float DEFAULT NULL,
  `QciAmbrLimitFlag` varchar(255) DEFAULT NULL,
  `UlAmbrLimitFactor` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellrachalgo
CREATE TABLE IF NOT EXISTS `cellrachalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PrachFalseAlarmDetRadThd` float DEFAULT NULL,
  `RachThdBoostRatio` float DEFAULT NULL,
  `PrachInterfPeriod` float DEFAULT NULL,
  `PrachInterfThreshold` float DEFAULT NULL,
  `PrachInterfHysteresis` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellrachcecfg
CREATE TABLE IF NOT EXISTS `cellrachcecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `ContentionResolutionTimer` varchar(255) DEFAULT NULL,
  `PrachTransmissionPeriod` varchar(255) DEFAULT NULL,
  `PrachSubcarrierOffset` varchar(255) DEFAULT NULL,
  `PrachRepetitionCount` varchar(255) DEFAULT NULL,
  `MaxNumPreambleAttempt` varchar(255) DEFAULT NULL,
  `PrachDetectionThld` varchar(255) DEFAULT NULL,
  `PrachStartTime` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellracthd
CREATE TABLE IF NOT EXISTS `cellracthd` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `GoldServiceArpThd` float DEFAULT NULL,
  `SilverServiceArpThd` float DEFAULT NULL,
  `Qci1HoThd` float DEFAULT NULL,
  `Qci2HoThd` float DEFAULT NULL,
  `Qci3HoThd` float DEFAULT NULL,
  `Qci4HoThd` float DEFAULT NULL,
  `NewGoldServiceOffset` float DEFAULT NULL,
  `NewSilverServiceOffset` float DEFAULT NULL,
  `NewCopperServiceOffset` float DEFAULT NULL,
  `Qci1CongThd` float DEFAULT NULL,
  `Qci2CongThd` float DEFAULT NULL,
  `Qci3CongThd` float DEFAULT NULL,
  `Qci4CongThd` float DEFAULT NULL,
  `CongRelOffset` float DEFAULT NULL,
  `UlRbHighThd` float DEFAULT NULL,
  `UlRbLowThd` float DEFAULT NULL,
  `AcReservedUserNumber` float DEFAULT NULL,
  `VolteReservedNumber` float DEFAULT NULL,
  `VoltePrefAdmissionTimer` float DEFAULT NULL,
  `CellCapacityMode` varchar(255) DEFAULT NULL,
  `LoadHoAdmitOffset` float DEFAULT NULL,
  `VoipOverAdmitOffset` float DEFAULT NULL,
  `AcUserNumber` float DEFAULT NULL,
  `MoSigArpOverride` float DEFAULT NULL,
  `UlSatisThdforVolteLoadAmrc` float DEFAULT NULL,
  `CceUsageThd` float DEFAULT NULL,
  `PreResNeedTuningFactor` float DEFAULT NULL,
  `CceAlFailHighThd` float DEFAULT NULL,
  `CqiFarThd` float DEFAULT NULL,
  `DlExperienceThd` float DEFAULT NULL,
  `RbCongHighThd` float DEFAULT NULL,
  `UlExperienceThd` float DEFAULT NULL,
  `VolteArpOverride` float DEFAULT NULL,
  `CceThdforVolteLoadAmrc` float DEFAULT NULL,
  `UlRbThdforVolteLoadAmrc` float DEFAULT NULL,
  `CongBearRelPeriod` float DEFAULT NULL,
  `PreemptableBearerNum` float DEFAULT NULL,
  `Qci65HoThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellresel
CREATE TABLE IF NOT EXISTS `cellresel` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Qhyst` varchar(255) DEFAULT NULL,
  `SpeedDepReselCfgInd` varchar(255) DEFAULT NULL,
  `SNonIntraSearchCfgInd` varchar(255) DEFAULT NULL,
  `SNonIntraSearch` float DEFAULT NULL,
  `ThrshServLow` float DEFAULT NULL,
  `CellReselPriority` float DEFAULT NULL,
  `QRxLevMin` float DEFAULT NULL,
  `PMaxCfgInd` varchar(255) DEFAULT NULL,
  `SIntraSearchCfgInd` varchar(255) DEFAULT NULL,
  `SIntraSearch` float DEFAULT NULL,
  `MeasBandWidthCfgInd` varchar(255) DEFAULT NULL,
  `TReselEutran` float DEFAULT NULL,
  `SpeedStateSfCfgInd` varchar(255) DEFAULT NULL,
  `TReselEutranSfMedium` varchar(255) DEFAULT NULL,
  `TReselEutranSfHigh` varchar(255) DEFAULT NULL,
  `NeighCellConfig` varchar(255) DEFAULT NULL,
  `PresenceAntennaPort1` varchar(255) DEFAULT NULL,
  `SIntraSearchQ` float DEFAULT NULL,
  `SNonIntraSearchQ` float DEFAULT NULL,
  `ThrshServLowQCfgInd` varchar(255) DEFAULT NULL,
  `QQualMinCfgInd` varchar(255) DEFAULT NULL,
  `QQualMin` float DEFAULT NULL,
  `TReselForNb` varchar(255) DEFAULT NULL,
  `TReselInterFreqForNb` varchar(255) DEFAULT NULL,
  `TReselEutranCE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellreselgeran
CREATE TABLE IF NOT EXISTS `cellreselgeran` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TReselGeran` float DEFAULT NULL,
  `SpeedStateSfCfgInd` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellreselutran
CREATE TABLE IF NOT EXISTS `cellreselutran` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TReselUtran` float DEFAULT NULL,
  `SpeedStateSfCfgInd` varchar(255) DEFAULT NULL,
  `TReselUtranSfMedium` varchar(255) DEFAULT NULL,
  `TReselUtranSfHigh` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellrfshutdown
CREATE TABLE IF NOT EXISTS `cellrfshutdown` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RfShutdownSwitch` varchar(255) DEFAULT NULL,
  `StartTime` varchar(255) DEFAULT NULL,
  `StopTime` varchar(255) DEFAULT NULL,
  `RsPwrAdjOffset` float DEFAULT NULL,
  `DlPrbThd` float DEFAULT NULL,
  `UlPrbThd` float DEFAULT NULL,
  `DlPrbOffset` float DEFAULT NULL,
  `UlPrbOffset` float DEFAULT NULL,
  `UENumThd` float DEFAULT NULL,
  `RfShutdownJudgingPolicy` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellricalgo
CREATE TABLE IF NOT EXISTS `cellricalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RicAlgoSwitch` varchar(255) DEFAULT NULL,
  `MuteUpPTSSymNum` float DEFAULT NULL,
  `MuteULSym` varchar(255) DEFAULT NULL,
  `UlInterferenceThd` float DEFAULT NULL,
  `DuctingRemoteInfThd` float DEFAULT NULL,
  `InfAvoidDetPeriodNum` float DEFAULT NULL,
  `InfAvoidRecDetPeriodNum` float DEFAULT NULL,
  `RemoteInfAdpAvoidSwitch` varchar(255) DEFAULT NULL,
  `UlInterferenceSymbMaxDiff` float DEFAULT NULL,
  `DuctDLSubfrmShutoffSwitch` varchar(255) DEFAULT NULL,
  `UlInterferenceDuration` float DEFAULT NULL,
  `DuctDLSubfrmShutoffInfThld` float DEFAULT NULL,
  `DuctDLSubfrmShutoffOptSw` varchar(255) DEFAULT NULL,
  `DuctInfDetPeriodCount` float DEFAULT NULL,
  `DuctInfPwrDiffInThld` float DEFAULT NULL,
  `DuctInfPwrDiffOutThld` float DEFAULT NULL,
  `DuctInfRateThld` float DEFAULT NULL,
  `DuctSubfrmShutoffDepth` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellsel
CREATE TABLE IF NOT EXISTS `cellsel` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `QRxLevMin` float DEFAULT NULL,
  `QRxLevMinOffset` float DEFAULT NULL,
  `QQualMin` float DEFAULT NULL,
  `QQualMinOffsetCfgInd` varchar(255) DEFAULT NULL,
  `QRxLevMinCE` float DEFAULT NULL,
  `QQualMinCE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellservicediffcfg
CREATE TABLE IF NOT EXISTS `cellservicediffcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ServiceDiffSwitch` varchar(255) DEFAULT NULL,
  `CellConStatePrbThd` float DEFAULT NULL,
  `CellOverloadStatePrbThd` float DEFAULT NULL,
  `CellConStateUeNumThd` float DEFAULT NULL,
  `CellOverloadStateUeNumThd` float DEFAULT NULL,
  `QueueServiceWeight0` float DEFAULT NULL,
  `QueueServiceWeight1` float DEFAULT NULL,
  `QueueServiceWeight2` float DEFAULT NULL,
  `QueueServiceWeight3` float DEFAULT NULL,
  `QueueServiceWeight4` float DEFAULT NULL,
  `QueueServiceWeight5` float DEFAULT NULL,
  `QueueServiceWeight6` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellshutdown
CREATE TABLE IF NOT EXISTS `cellshutdown` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellShutdownSwitch` varchar(255) DEFAULT NULL,
  `StartTime` varchar(255) DEFAULT NULL,
  `StopTime` varchar(255) DEFAULT NULL,
  `DlPrbThd` float DEFAULT NULL,
  `DlPrbOffset` float DEFAULT NULL,
  `UlPrbThd` float DEFAULT NULL,
  `UlPrbOffset` float DEFAULT NULL,
  `ForceShutdownUENumThd` float DEFAULT NULL,
  `PunishTime` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellsimap
CREATE TABLE IF NOT EXISTS `cellsimap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SiMapSwitch` varchar(255) DEFAULT NULL,
  `Sib2Period` varchar(255) DEFAULT NULL,
  `Sib3Period` varchar(255) DEFAULT NULL,
  `Sib4Period` varchar(255) DEFAULT NULL,
  `Sib5Period` varchar(255) DEFAULT NULL,
  `Sib6Period` varchar(255) DEFAULT NULL,
  `Sib7Period` varchar(255) DEFAULT NULL,
  `Sib8Period` varchar(255) DEFAULT NULL,
  `Sib10Period` varchar(255) DEFAULT NULL,
  `Sib11Period` varchar(255) DEFAULT NULL,
  `EtwsPnDuration` float DEFAULT NULL,
  `EtwsSnOverlapPolicy` varchar(255) DEFAULT NULL,
  `EtwsPnOverlapPolicy` varchar(255) DEFAULT NULL,
  `Sib12Period` varchar(255) DEFAULT NULL,
  `SiTransEcr` float DEFAULT NULL,
  `Sib13Period` varchar(255) DEFAULT NULL,
  `Sib15Period` varchar(255) DEFAULT NULL,
  `Sib16Period` varchar(255) DEFAULT NULL,
  `SibTransCtrlSwitch` varchar(255) DEFAULT NULL,
  `SiSwitch` varchar(255) DEFAULT NULL,
  `SiSchResRatio` float DEFAULT NULL,
  `SibUpdOptSwitch` varchar(255) DEFAULT NULL,
  `Sib14Period` varchar(255) DEFAULT NULL,
  `NbSib1RepetitionNum` varchar(255) DEFAULT NULL,
  `NbSib2Period` varchar(255) DEFAULT NULL,
  `NbSib3Period` varchar(255) DEFAULT NULL,
  `NbSib4Period` varchar(255) DEFAULT NULL,
  `NbSib5Period` varchar(255) DEFAULT NULL,
  `NbSib14Period` varchar(255) DEFAULT NULL,
  `NbSib16Period` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellsrlte
CREATE TABLE IF NOT EXISTS `cellsrlte` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SrlteSwitch` varchar(255) DEFAULT NULL,
  `SrlteDtxThrd` float DEFAULT NULL,
  `SrlteSuspendTime` float DEFAULT NULL,
  `SrlteDiscardAlgoSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellsrsadaptivecfg
CREATE TABLE IF NOT EXISTS `cellsrsadaptivecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SrsPeriodAdaptive` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellstandardqci
CREATE TABLE IF NOT EXISTS `cellstandardqci` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Qci` varchar(255) DEFAULT NULL,
  `InterFreqHoGroupId` float DEFAULT NULL,
  `InterRatHoCdma1xRttGroupId` float DEFAULT NULL,
  `InterRatHoCdmaHrpdGroupId` float DEFAULT NULL,
  `InterRatHoCommGroupId` float DEFAULT NULL,
  `InterRatHoGeranGroupId` float DEFAULT NULL,
  `InterRatHoUtranGroupId` float DEFAULT NULL,
  `IntraFreqHoGroupId` float DEFAULT NULL,
  `DrxParaGroupId` float DEFAULT NULL,
  `SriPeriod` varchar(255) DEFAULT NULL,
  `RlfTimerConstCfgInd` varchar(255) DEFAULT NULL,
  `RlfTimerConstGroupId` float DEFAULT NULL,
  `TrafficRelDelay` float DEFAULT NULL,
  `QciPriorityForHo` float DEFAULT NULL,
  `MlbQciGroupId` float DEFAULT NULL,
  `PreallocationParaGroupId` float DEFAULT NULL,
  `QciPriorityForDrx` float DEFAULT NULL,
  `QciAlgoSwitch` varchar(255) DEFAULT NULL,
  `QciEutranFreqRelationId` float DEFAULT NULL,
  `QciUtranFreqRelationId` float DEFAULT NULL,
  `QciGeranFreqRelationId` float DEFAULT NULL,
  `QciAmbrLimitFlag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellttibundlingalgo
CREATE TABLE IF NOT EXISTS `cellttibundlingalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SinrThdToTrigTtib` float DEFAULT NULL,
  `SinrThdToTrigVideoTtib` float DEFAULT NULL,
  `TtiBundlingAlgoSw` varchar(255) DEFAULT NULL,
  `R12TtiBHarqMaxTxNum` varchar(255) DEFAULT NULL,
  `R12TtiBundlingSwitch` varchar(255) DEFAULT NULL,
  `SinrThdToTrigR12TtiB` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellucionpuschpara
CREATE TABLE IF NOT EXISTS `cellucionpuschpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CellUciOnPuschParaValid` varchar(255) DEFAULT NULL,
  `DeltaOffsetCqiIndex` float DEFAULT NULL,
  `DeltaOffsetRiIndex` float DEFAULT NULL,
  `DeltaOffsetAckIndex` float DEFAULT NULL,
  `DeltaOffsetAckIndexForTtiB` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.celluemeascontrolcfg
CREATE TABLE IF NOT EXISTS `celluemeascontrolcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MaxNonIntraMeasObjNum` float DEFAULT NULL,
  `MaxEutranFddMeasFreqNum` float DEFAULT NULL,
  `MaxEutranTddMeasFreqNum` float DEFAULT NULL,
  `MaxUtranFddMeasFreqNum` float DEFAULT NULL,
  `MaxUtranTddMeasFreqNum` float DEFAULT NULL,
  `MaxGeranMeasFreqNum` float DEFAULT NULL,
  `MaxUeTddMeasFreqNum` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulcompalgo
CREATE TABLE IF NOT EXISTS `cellulcompalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlCompA3Offset` float DEFAULT NULL,
  `SfnUlCompThd` float DEFAULT NULL,
  `UlCompA3OffsetForRelaxedBH` float DEFAULT NULL,
  `UlCompUlRsrpOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulicalgo
CREATE TABLE IF NOT EXISTS `cellulicalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlIcA3Offset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulicic
CREATE TABLE IF NOT EXISTS `cellulicic` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BandMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulicicmcpara
CREATE TABLE IF NOT EXISTS `cellulicicmcpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `A3Offset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulmimoparacfg
CREATE TABLE IF NOT EXISTS `cellulmimoparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlSuMimoRankPara` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulpccomm
CREATE TABLE IF NOT EXISTS `cellulpccomm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PassLossCoeff` varchar(255) DEFAULT NULL,
  `P0NominalPUCCH` float DEFAULT NULL,
  `P0NominalPUSCH` float DEFAULT NULL,
  `DeltaFPUCCHFormat1` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat1b` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat2` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat2a` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat2b` varchar(255) DEFAULT NULL,
  `DeltaPreambleMsg3` float DEFAULT NULL,
  `DeltaMsg2` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat1bcs` varchar(255) DEFAULT NULL,
  `DeltaFPUCCHFormat3` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulpcdedic
CREATE TABLE IF NOT EXISTS `cellulpcdedic` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DeltaMcsEnabled` varchar(255) DEFAULT NULL,
  `PSrsOffsetDeltaMcsDisable` float DEFAULT NULL,
  `PSrsOffsetDeltaMcsEnable` float DEFAULT NULL,
  `FilterRsrp` varchar(255) DEFAULT NULL,
  `PathlossReferenceLink` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulschalgo
CREATE TABLE IF NOT EXISTS `cellulschalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlschStrategy` varchar(255) DEFAULT NULL,
  `AdaptHarqSwitch` varchar(255) DEFAULT NULL,
  `SinrAdjustTargetIbler` float DEFAULT NULL,
  `PreAllocationBandwidthRatio` float DEFAULT NULL,
  `PreAllocationMinPeriod` float DEFAULT NULL,
  `PreAllocationSize` float DEFAULT NULL,
  `UlHoppingType` varchar(255) DEFAULT NULL,
  `FreeUserUlRbUsedRatio` float DEFAULT NULL,
  `UlSrSchDateLen` float DEFAULT NULL,
  `SpsRelThd` float DEFAULT NULL,
  `SmartPreAllocationDuration` float DEFAULT NULL,
  `UlEpfCapacityFactor` varchar(255) DEFAULT NULL,
  `UlRbAllocationStrategy` varchar(255) DEFAULT NULL,
  `DopMeasLevel` varchar(255) DEFAULT NULL,
  `UlHarqMaxTxNum` float DEFAULT NULL,
  `SmartPreAllocDuraForSparse` float DEFAULT NULL,
  `UlSpsInterval` varchar(255) DEFAULT NULL,
  `SriFalseDetThdSwitch` varchar(255) DEFAULT NULL,
  `NoUlSchTtiNumAffterGap` float DEFAULT NULL,
  `UlDelaySchStrategy` varchar(255) DEFAULT NULL,
  `UlSchAbnUeThd` float DEFAULT NULL,
  `SpecificPktSizeThd` float DEFAULT NULL,
  `SrMaskSwitch` varchar(255) DEFAULT NULL,
  `PuschDtxSchStrategy` varchar(255) DEFAULT NULL,
  `UlVoipRlcMaxSegNum` float DEFAULT NULL,
  `UlEnhencedVoipSchSw` longtext,
  `UlInBasedFssSinrThld` float DEFAULT NULL,
  `UlCamcDlRsrpOffset` float DEFAULT NULL,
  `StatisticNumThdForTtibTrig` varchar(255) DEFAULT NULL,
  `StatisticNumThdForTtibExit` varchar(255) DEFAULT NULL,
  `HystToExitTtiBundling` float DEFAULT NULL,
  `TtiBundlingRlcMaxSegNum` float DEFAULT NULL,
  `TtiBundlingHarqMaxTxNum` varchar(255) DEFAULT NULL,
  `TtiBundlingTriggerStrategy` varchar(255) DEFAULT NULL,
  `DopAlgoSwitch` varchar(255) DEFAULT NULL,
  `EnhancedVmimoSwitch` varchar(255) DEFAULT NULL,
  `UeNumThdInPdcchPuschBal` float DEFAULT NULL,
  `DataThdInPdcchPuschBal` float DEFAULT NULL,
  `HeadOverheadForUlSch` varchar(255) DEFAULT NULL,
  `PreAllocMinPeriodForSparse` varchar(255) DEFAULT NULL,
  `PreallocationSizeForSparse` varchar(255) DEFAULT NULL,
  `UlInterfRandomMode` varchar(255) DEFAULT NULL,
  `SinrAdjTargetIblerforVoLTE` float DEFAULT NULL,
  `SfnUlLoadPeriod` float DEFAULT NULL,
  `MaxLayerHOVMIMO` varchar(255) DEFAULT NULL,
  `UlCompenSchPeriodinSpurt` varchar(255) DEFAULT NULL,
  `UlCompenSchPeriodinSilence` varchar(255) DEFAULT NULL,
  `UlTargetIBlerAdaptType` varchar(255) DEFAULT NULL,
  `AperiodicCsiUlTxMode` varchar(255) DEFAULT NULL,
  `UlVoipRsvRbStart` float DEFAULT NULL,
  `UlVoipRsvRbNum` varchar(255) DEFAULT NULL,
  `UlIBlerAdaptBigTrafficSw` varchar(255) DEFAULT NULL,
  `VmimoOptAlgoSwitch` varchar(255) DEFAULT NULL,
  `UlCamcInterfTh` float DEFAULT NULL,
  `UlIcsA3Offset` float DEFAULT NULL,
  `UlMcsLowThreshold` varchar(255) DEFAULT NULL,
  `UserSpeedUlSchPriWeight` float DEFAULT NULL,
  `SinrAdjTargetIblerforPtt` float DEFAULT NULL,
  `TarRruSelRsrpOffsetThd` float DEFAULT NULL,
  `MaxUlSchRbNum` float DEFAULT NULL,
  `UlExtVolteSchSw` varchar(255) DEFAULT NULL,
  `UlVolteDeltaSinrForNack` float DEFAULT NULL,
  `InitUlSinrAdjust` float DEFAULT NULL,
  `UlSinrAdjustStep` float DEFAULT NULL,
  `UlSinrFilterCoef` float DEFAULT NULL,
  `SfnUlPairRsrpThd` float DEFAULT NULL,
  `UlSrSchDataVolAdptOptUpThd` float DEFAULT NULL,
  `TtiBundlingRetxStrategy` varchar(255) DEFAULT NULL,
  `UlVoLTERetransSchStrategy` varchar(255) DEFAULT NULL,
  `NbUlHarqMaxTxCount` float DEFAULT NULL,
  `PrealloSystemBwCoeff` float DEFAULT NULL,
  `SmartPrealloDurationCoeff` float DEFAULT NULL,
  `HighSpeedSdmaIsolationThld` float DEFAULT NULL,
  `VmimoPairingStrategy` varchar(255) DEFAULT NULL,
  `VMIMOEgdeResRatio` float DEFAULT NULL,
  `MaxLayerMMVMIMO` varchar(255) DEFAULT NULL,
  `UlSrSchDateLenForVoLTE` float DEFAULT NULL,
  `AmrcDecreasingPeriod` float DEFAULT NULL,
  `SinrThdForVoLTERateCtrl` float DEFAULT NULL,
  `RateCtrlCmrProcessStrategy` varchar(255) DEFAULT NULL,
  `SinrAdjTargetIblerforQCI2` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellulunifiedolc
CREATE TABLE IF NOT EXISTS `cellulunifiedolc` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RrcRejectRateHighThd` float DEFAULT NULL,
  `RrcRejectRateLowThd` float DEFAULT NULL,
  `RrcReqNumHighThd` float DEFAULT NULL,
  `RrcReqNumLowThd` float DEFAULT NULL,
  `UeNumHighThd` float DEFAULT NULL,
  `UeNumLowThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellusparacfg
CREATE TABLE IF NOT EXISTS `cellusparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UsAlgoSwitch` longtext,
  `UsVoIPPreAllocMinPeriod` float DEFAULT NULL,
  `UsDataPreAllocMinPeriod` float DEFAULT NULL,
  `UsVoIPPreAllocSize` float DEFAULT NULL,
  `UsDataPreAllocSize` float DEFAULT NULL,
  `UsVoIPSmartPreallocDura` float DEFAULT NULL,
  `UsDataSmartPreallocDura` float DEFAULT NULL,
  `UlUsVoIPIblerTarget` float DEFAULT NULL,
  `NsNumThdForUsHo` float DEFAULT NULL,
  `UsPucchSinrTargetOffset` float DEFAULT NULL,
  `UsPucchRsrpHighThdOffset` float DEFAULT NULL,
  `UsDataRatFreqPriGroupId` float DEFAULT NULL,
  `UsGuaranteeDurTimer` float DEFAULT NULL,
  `UsGapMeasurementPeriod` float DEFAULT NULL,
  `UsHoA2ThdRsrp` float DEFAULT NULL,
  `UsHoA2ThdRsrq` float DEFAULT NULL,
  `UsVoIPRatFreqPriGroupId` float DEFAULT NULL,
  `NsRatFreqPriGroupId` float DEFAULT NULL,
  `UsPaPcOff` varchar(255) DEFAULT NULL,
  `UsPuschSinrHighThdOffset` float DEFAULT NULL,
  `NsRbRestrictRatio` float DEFAULT NULL,
  `UsDetectTimer` float DEFAULT NULL,
  `UsDlMinGbr` float DEFAULT NULL,
  `UsUlMinGbr` float DEFAULT NULL,
  `UsNbInterfCtrl` float DEFAULT NULL,
  `UlUsVoipRsvStartRb` float DEFAULT NULL,
  `UlUsVoipRsvEndRb` float DEFAULT NULL,
  `SfCtrlPrbThd` varchar(255) DEFAULT NULL,
  `UsVoIPKeepOnLen` float DEFAULT NULL,
  `NsIdleRatFreqPriGroupId` float DEFAULT NULL,
  `UsDlPriorityRatio` float DEFAULT NULL,
  `UsUlPriorityRatio` float DEFAULT NULL,
  `UsA3HoA2ThdOffset` float DEFAULT NULL,
  `UsA4A5HoA2ThdOffset` float DEFAULT NULL,
  `UlUsDataRatFreqPriGroupId` float DEFAULT NULL,
  `UsNbUlSinrHighThdOffset` float DEFAULT NULL,
  `UsNbUlRbRestrictRat` float DEFAULT NULL,
  `UsA3HoA1ThdOffset` float DEFAULT NULL,
  `UsA4A5HoA1ThdOffset` float DEFAULT NULL,
  `UsDataPdcchSinrOffset` float DEFAULT NULL,
  `UsVoipPdcchSinrOffset` float DEFAULT NULL,
  `UsVoipInitDlIblerTarget` float DEFAULT NULL,
  `UsVoipCompenSchPeriod` varchar(255) DEFAULT NULL,
  `CrsShutDownRsrpOffset` float DEFAULT NULL,
  `CrsShutDownStrategy` varchar(255) DEFAULT NULL,
  `VolteSilDelayHoRsrpThld` float DEFAULT NULL,
  `UsCellSchOptSwitch` varchar(255) DEFAULT NULL,
  `MMPowerRatio` float DEFAULT NULL,
  `NsDlFirstRetxRbRatio` float DEFAULT NULL,
  `NsMaxNumAllowSchPerTTI` float DEFAULT NULL,
  `NsNotAllowSchSubframe` varchar(255) DEFAULT NULL,
  `SchGuaranteeUsNum` float DEFAULT NULL,
  `UsWithGapMode` varchar(255) DEFAULT NULL,
  `LargePacketDlDataThld` float DEFAULT NULL,
  `UsNotSchRbgRsrpOffset` float DEFAULT NULL,
  `UsNotSchRbgRsrpThld` float DEFAULT NULL,
  `UsAlgoExSwitch` varchar(255) DEFAULT NULL,
  `UsSrsGuaranteeSwitch` varchar(255) DEFAULT NULL,
  `UeMobilIdentifyStrategy` varchar(255) DEFAULT NULL,
  `UsSmoothHoCompThldOfs` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellvms
CREATE TABLE IF NOT EXISTS `cellvms` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `VmsHoUeNumTh` float DEFAULT NULL,
  `VmsPrbDiffTh` float DEFAULT NULL,
  `VmsPrbLoadTh` float DEFAULT NULL,
  `VmsA3Offset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cellwttxparacfg
CREATE TABLE IF NOT EXISTS `cellwttxparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MbbUserDlPrbUpLimit` float DEFAULT NULL,
  `MbbUserUlPrbUpLimit` float DEFAULT NULL,
  `PrbUpLimitCtrlMode` varchar(255) DEFAULT NULL,
  `WbbOrMbbUserDefMode` varchar(255) DEFAULT NULL,
  `WbbUserDlPrbUpLimit` float DEFAULT NULL,
  `WbbUserUlPrbUpLimit` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cepucchcfg
CREATE TABLE IF NOT EXISTS `cepucchcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `PucchRepNum` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cerachcfg
CREATE TABLE IF NOT EXISTS `cerachcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `ContentionResolutionTimer` varchar(255) DEFAULT NULL,
  `MaxNumPrbAttempt` varchar(255) DEFAULT NULL,
  `PreambleRatio` float DEFAULT NULL,
  `PreambleRepetitionNum` varchar(255) DEFAULT NULL,
  `RandomPreambleRatio` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.certcfg
CREATE TABLE IF NOT EXISTS `certcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IKECHECKSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.certchktsk
CREATE TABLE IF NOT EXISTS `certchktsk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ISENABLE` varchar(255) DEFAULT NULL,
  `PERIOD` float DEFAULT NULL,
  `ALMRNG` float DEFAULT NULL,
  `UPDATEMETHOD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.certdeploy
CREATE TABLE IF NOT EXISTS `certdeploy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DEPLOYTYPE` varchar(255) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.certmk
CREATE TABLE IF NOT EXISTS `certmk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `APPCERT` varchar(255) DEFAULT NULL,
  `CASW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.certreq
CREATE TABLE IF NOT EXISTS `certreq` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `COMMNAME` varchar(255) DEFAULT NULL,
  `USERADDINFO` varchar(255) DEFAULT NULL,
  `COUNTRY` varchar(255) DEFAULT NULL,
  `ORG` varchar(255) DEFAULT NULL,
  `ORGUNIT` varchar(255) DEFAULT NULL,
  `STATEPROVINCENAME` varchar(255) DEFAULT NULL,
  `LOCALITY` varchar(255) DEFAULT NULL,
  `KEYUSAGE` varchar(255) DEFAULT NULL,
  `SIGNALG` varchar(255) DEFAULT NULL,
  `KEYSIZE` varchar(255) DEFAULT NULL,
  `LOCALNAME` varchar(255) DEFAULT NULL,
  `LOCALIP` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.chk
CREATE TABLE IF NOT EXISTS `chk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ENABLEFLAG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.clkdetect
CREATE TABLE IF NOT EXISTS `clkdetect` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ClkAsyncDetectSwitch` varchar(255) DEFAULT NULL,
  `ClkAsyncInterfRptThld` float DEFAULT NULL,
  `ClkAsyncPrbUsageThld` float DEFAULT NULL,
  `ClkAsyncSilDetInfDifThld` float DEFAULT NULL,
  `ClkAsyncSilDetInfThld` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.clzerobufferzone
CREATE TABLE IF NOT EXISTS `clzerobufferzone` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ClZeroBufzoneDlSharedInd` varchar(255) DEFAULT NULL,
  `ClZeroBufzoneUlSharedInd` varchar(255) DEFAULT NULL,
  `ClSharedFreqStartRb1` float DEFAULT NULL,
  `ClSharedFreqEndRb1` float DEFAULT NULL,
  `ClSharedFreqStartRb2` float DEFAULT NULL,
  `ClSharedFreqEndRb2` float DEFAULT NULL,
  `ClZeroBufferZoneUlPrbThd` float DEFAULT NULL,
  `ClZeroBufferZoneUlPrbOst` float DEFAULT NULL,
  `UlNearPtUserRsrpThd` float DEFAULT NULL,
  `DlNearPtUserRsrpThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cnoperator
CREATE TABLE IF NOT EXISTS `cnoperator` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `CnOperatorName` varchar(255) DEFAULT NULL,
  `CnOperatorType` varchar(255) DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `PlmnMode` varchar(255) DEFAULT NULL,
  `OperatorFunSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cnoperatorhocfg
CREATE TABLE IF NOT EXISTS `cnoperatorhocfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `FirstRatPri` varchar(255) DEFAULT NULL,
  `SecondRatPri` varchar(255) DEFAULT NULL,
  `TddIfHoA2ThdRsrpOffset` float DEFAULT NULL,
  `FddIfHoA2ThdRsrpOffset` float DEFAULT NULL,
  `UtranA2ThdRsrpOffset` float DEFAULT NULL,
  `GeranA2ThdRsrpOffset` float DEFAULT NULL,
  `SrvccWithPsBasedCellCapSw` varchar(255) DEFAULT NULL,
  `PsInterRatSecondPri` varchar(255) DEFAULT NULL,
  `PsInterRatHighestPri` varchar(255) DEFAULT NULL,
  `PsInterRatLowestPri` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cnoperatorqcipara
CREATE TABLE IF NOT EXISTS `cnoperatorqcipara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `ServiceIrHoCfgGroupId` float DEFAULT NULL,
  `ServiceIfHoCfgGroupId` float DEFAULT NULL,
  `ServiceHoBearerPolicy` varchar(255) DEFAULT NULL,
  `LocalQciArp` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cnoperatorstandardqci
CREATE TABLE IF NOT EXISTS `cnoperatorstandardqci` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `Qci` varchar(255) DEFAULT NULL,
  `ServiceIrHoCfgGroupId` float DEFAULT NULL,
  `ServiceIfHoCfgGroupId` float DEFAULT NULL,
  `ServiceHoBearerPolicy` varchar(255) DEFAULT NULL,
  `LocalQciArp` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cnoperatorta
CREATE TABLE IF NOT EXISTS `cnoperatorta` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TrackingAreaId` float DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `Tac` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `NbIotTaFlag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.countercheckpara
CREATE TABLE IF NOT EXISTS `countercheckpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CounterCheckTimer` float DEFAULT NULL,
  `CounterCheckCountNum` float DEFAULT NULL,
  `CounterCheckUserRelSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cpbearer
CREATE TABLE IF NOT EXISTS `cpbearer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CPBEARID` float DEFAULT NULL,
  `FLAG` varchar(255) DEFAULT NULL,
  `BEARTYPE` varchar(255) DEFAULT NULL,
  `LINKNO` float DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cpriport
CREATE TABLE IF NOT EXISTS `cpriport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `OPTN` float DEFAULT NULL,
  `ADMINISTRATIVESTATE` varchar(255) DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `SPN` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cpswitch
CREATE TABLE IF NOT EXISTS `cpswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ES` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cqiadaptivecfg
CREATE TABLE IF NOT EXISTS `cqiadaptivecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CqiPeriodAdaptive` varchar(255) DEFAULT NULL,
  `SimulAckNackAndCqiSwitch` varchar(255) DEFAULT NULL,
  `PucchPeriodicCqiOptSwitch` varchar(255) DEFAULT NULL,
  `HoAperiodicCqiCfgSwitch` varchar(255) DEFAULT NULL,
  `SimulAckNackAndCqiFmt3Sw` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.crlpolicy
CREATE TABLE IF NOT EXISTS `crlpolicy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CRLPOLICY` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.csfallbackblindhocfg
CREATE TABLE IF NOT EXISTS `csfallbackblindhocfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `InterRatHighestPri` varchar(255) DEFAULT NULL,
  `InterRatSecondPri` varchar(255) DEFAULT NULL,
  `InterRatLowestPri` varchar(255) DEFAULT NULL,
  `UtranLcsCap` varchar(255) DEFAULT NULL,
  `GeranLcsCap` varchar(255) DEFAULT NULL,
  `CdmaLcsCap` varchar(255) DEFAULT NULL,
  `IdleCsfbHighestPri` varchar(255) DEFAULT NULL,
  `IdleCsfbSecondPri` varchar(255) DEFAULT NULL,
  `IdleCsfbLowestPri` varchar(255) DEFAULT NULL,
  `UtranCsfbBlindRedirRrSw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.csfallbackho
CREATE TABLE IF NOT EXISTS `csfallbackho` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CsfbHoUtranTimeToTrig` varchar(255) DEFAULT NULL,
  `CsfbHoGeranTimeToTrig` varchar(255) DEFAULT NULL,
  `CsfbHoCdmaTimeToTrig` varchar(255) DEFAULT NULL,
  `CsfbHoUtranB1ThdRscp` float DEFAULT NULL,
  `CsfbHoUtranB1ThdEcn0` float DEFAULT NULL,
  `CsfbHoGeranB1Thd` float DEFAULT NULL,
  `CsfbHoCdmaB1ThdPs` float DEFAULT NULL,
  `BlindHoA1ThdRsrp` float DEFAULT NULL,
  `CsfbProtectionTimer` float DEFAULT NULL,
  `CsfbProtectTimer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.csfallbackpolicycfg
CREATE TABLE IF NOT EXISTS `csfallbackpolicycfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CsfbHoPolicyCfg` varchar(255) DEFAULT NULL,
  `IdleModeCsfbHoPolicyCfg` varchar(255) DEFAULT NULL,
  `CsfbUserArpCfgSwitch` varchar(255) DEFAULT NULL,
  `NormalCsfbUserArp` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.cspcalgopara
CREATE TABLE IF NOT EXISTS `cspcalgopara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CspcAlgoSwitch` varchar(255) DEFAULT NULL,
  `CspcRsrpMeasMode` varchar(255) DEFAULT NULL,
  `CspcClusterMode` varchar(255) DEFAULT NULL,
  `CspcClusterPartPeriod` float DEFAULT NULL,
  `CspcComputeSwitch` varchar(255) DEFAULT NULL,
  `CspcFullPowerSubframeRatio` float DEFAULT NULL,
  `CspcPowerConfigDelay` float DEFAULT NULL,
  `CspcScheduleUeSpec` float DEFAULT NULL,
  `CspcEnableDlPrbRatioThd` float DEFAULT NULL,
  `TddCspcAlgoSwitch` varchar(255) DEFAULT NULL,
  `CspcCapacityFactor` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.devip
CREATE TABLE IF NOT EXISTS `devip` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `IP` varchar(255) DEFAULT NULL,
  `MASK` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.dhcprelayswitch
CREATE TABLE IF NOT EXISTS `dhcprelayswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ES` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.dhcpsvrip
CREATE TABLE IF NOT EXISTS `dhcpsvrip` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DHCPSVRIP` varchar(255) DEFAULT NULL,
  `DHCPRELAYIPSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.dhcpsw
CREATE TABLE IF NOT EXISTS `dhcpsw` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL,
  `VLANSCANSW` varchar(255) DEFAULT NULL,
  `DELAYSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.difpri
CREATE TABLE IF NOT EXISTS `difpri` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PRIRULE` varchar(255) DEFAULT NULL,
  `SIGPRI` float DEFAULT NULL,
  `OMHIGHPRI` float DEFAULT NULL,
  `OMLOWPRI` float DEFAULT NULL,
  `IPCLKPRI` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.distbasedho
CREATE TABLE IF NOT EXISTS `distbasedho` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DistBasedMeasObjType` varchar(255) DEFAULT NULL,
  `DistBasedHOThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.dlflowctrlpara
CREATE TABLE IF NOT EXISTS `dlflowctrlpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `BEAR` varchar(255) DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL,
  `TD` float DEFAULT NULL,
  `DR` float DEFAULT NULL,
  `ITM` varchar(255) DEFAULT NULL,
  `BWPRTSWITCH` varchar(255) DEFAULT NULL,
  `FAIRSWITCH` varchar(255) DEFAULT NULL,
  `FAIRRATIO` float DEFAULT NULL,
  `DLIUBMINBW` float DEFAULT NULL,
  `DLFLOWCTRLENHSW` varchar(255) DEFAULT NULL,
  `DLIUBBWMINPRORATIO` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.drx
CREATE TABLE IF NOT EXISTS `drx` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DrxAlgSwitch` varchar(255) DEFAULT NULL,
  `ShortDrxSwitch` varchar(255) DEFAULT NULL,
  `LongDrxCycleSpecial` varchar(255) DEFAULT NULL,
  `OnDurationTimerSpecial` varchar(255) DEFAULT NULL,
  `DrxInactivityTimerSpecial` varchar(255) DEFAULT NULL,
  `SupportShortDrxSpecial` varchar(255) DEFAULT NULL,
  `LongDrxCycleForAnr` varchar(255) DEFAULT NULL,
  `LongDRXCycleforIRatAnr` varchar(255) DEFAULT NULL,
  `DrxInactivityTimerForAnr` varchar(255) DEFAULT NULL,
  `TddAnrDrxInactivityTimer` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `DrxFlexCtrlSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.drxparagroup
CREATE TABLE IF NOT EXISTS `drxparagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DrxParaGroupId` float DEFAULT NULL,
  `EnterDrxSwitch` varchar(255) DEFAULT NULL,
  `OnDurationTimer` varchar(255) DEFAULT NULL,
  `DrxInactivityTimer` varchar(255) DEFAULT NULL,
  `DrxReTxTimer` varchar(255) DEFAULT NULL,
  `LongDrxCycle` varchar(255) DEFAULT NULL,
  `SupportShortDrx` varchar(255) DEFAULT NULL,
  `ShortDrxCycle` varchar(255) DEFAULT NULL,
  `DrxShortCycleTimer` float DEFAULT NULL,
  `DrxUlReTxTimer` varchar(255) DEFAULT NULL,
  `ExtendLongDrxCycleSwitch` varchar(255) DEFAULT NULL,
  `CatType` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.dscpmap
CREATE TABLE IF NOT EXISTS `dscpmap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL,
  `VLANPRIO` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.e1t1
CREATE TABLE IF NOT EXISTS `e1t1` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PS` varchar(255) DEFAULT NULL,
  `WORKMODE` varchar(255) DEFAULT NULL,
  `FRAME` varchar(255) DEFAULT NULL,
  `LNCODE` varchar(255) DEFAULT NULL,
  `CLKM` varchar(255) DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.e1t1bear
CREATE TABLE IF NOT EXISTS `e1t1bear` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.e1t1ber
CREATE TABLE IF NOT EXISTS `e1t1ber` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BTL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.emc
CREATE TABLE IF NOT EXISTS `emc` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `EmcEnable` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.emu
CREATE TABLE IF NOT EXISTS `emu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MCN` float DEFAULT NULL,
  `MSRN` float DEFAULT NULL,
  `MPN` float DEFAULT NULL,
  `ADDR` float DEFAULT NULL,
  `TLTHD` float DEFAULT NULL,
  `TUTHD` float DEFAULT NULL,
  `HLTHD` float DEFAULT NULL,
  `HUTHD` float DEFAULT NULL,
  `SAAF` varchar(255) DEFAULT NULL,
  `SBAF` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbcelloprsvdpara
CREATE TABLE IF NOT EXISTS `enbcelloprsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TrackingAreaId` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara7` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara14` float DEFAULT NULL,
  `RsvdPara15` float DEFAULT NULL,
  `RsvdPara16` float DEFAULT NULL,
  `RsvdPara17` float DEFAULT NULL,
  `RsvdPara18` float DEFAULT NULL,
  `RsvdPara19` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbcellqcirsvdpara
CREATE TABLE IF NOT EXISTS `enbcellqcirsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara7` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara14` float DEFAULT NULL,
  `RsvdPara15` float DEFAULT NULL,
  `RsvdPara16` float DEFAULT NULL,
  `RsvdPara17` float DEFAULT NULL,
  `RsvdPara18` float DEFAULT NULL,
  `RsvdPara19` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbcellrsvdpara
CREATE TABLE IF NOT EXISTS `enbcellrsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdSwPara3` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara22` float DEFAULT NULL,
  `RsvdPara23` float DEFAULT NULL,
  `RsvdPara24` float DEFAULT NULL,
  `RsvdPara25` float DEFAULT NULL,
  `RsvdPara26` float DEFAULT NULL,
  `RsvdPara27` float DEFAULT NULL,
  `RsvdPara28` float DEFAULT NULL,
  `RsvdPara29` float DEFAULT NULL,
  `RsvdPara30` float DEFAULT NULL,
  `RsvdPara31` float DEFAULT NULL,
  `RsvdPara32` float DEFAULT NULL,
  `RsvdPara33` float DEFAULT NULL,
  `RsvdPara34` float DEFAULT NULL,
  `RsvdPara35` float DEFAULT NULL,
  `RsvdPara36` float DEFAULT NULL,
  `RsvdPara37` float DEFAULT NULL,
  `RsvdPara38` float DEFAULT NULL,
  `RsvdPara39` float DEFAULT NULL,
  `RsvdPara40` float DEFAULT NULL,
  `RsvdPara41` float DEFAULT NULL,
  `RsvdPara42` float DEFAULT NULL,
  `RsvdPara43` float DEFAULT NULL,
  `RsvdPara44` float DEFAULT NULL,
  `RsvdPara45` float DEFAULT NULL,
  `RsvdPara46` float DEFAULT NULL,
  `RsvdPara47` float DEFAULT NULL,
  `RsvdPara48` float DEFAULT NULL,
  `RsvdPara49` float DEFAULT NULL,
  `RsvdPara50` float DEFAULT NULL,
  `RsvdPara51` float DEFAULT NULL,
  `RsvdPara52` float DEFAULT NULL,
  `RsvdPara53` float DEFAULT NULL,
  `RsvdPara54` float DEFAULT NULL,
  `RsvdPara55` float DEFAULT NULL,
  `RsvdPara56` float DEFAULT NULL,
  `RsvdPara57` float DEFAULT NULL,
  `RsvdPara58` float DEFAULT NULL,
  `RsvdPara59` float DEFAULT NULL,
  `RsvdPara60` float DEFAULT NULL,
  `RsvdPara61` float DEFAULT NULL,
  `RsvdPara62` float DEFAULT NULL,
  `RsvdPara63` float DEFAULT NULL,
  `RsvdPara64` float DEFAULT NULL,
  `RsvdPara65` float DEFAULT NULL,
  `RsvdPara66` float DEFAULT NULL,
  `RsvdPara67` float DEFAULT NULL,
  `RsvdPara68` float DEFAULT NULL,
  `RsvdPara69` float DEFAULT NULL,
  `RsvdU32Para1` float DEFAULT NULL,
  `RsvdU32Para2` float DEFAULT NULL,
  `RsvdU32Para3` float DEFAULT NULL,
  `RsvdU32Para4` float DEFAULT NULL,
  `RsvdU32Para5` float DEFAULT NULL,
  `RsvdU32Para6` float DEFAULT NULL,
  `RsvdU32Para7` float DEFAULT NULL,
  `RsvdU32Para8` float DEFAULT NULL,
  `RsvdU16Para1` float DEFAULT NULL,
  `RsvdU16Para2` float DEFAULT NULL,
  `RsvdU16Para3` float DEFAULT NULL,
  `RsvdU16Para4` float DEFAULT NULL,
  `RsvdU16Para5` float DEFAULT NULL,
  `RsvdU16Para6` float DEFAULT NULL,
  `RsvdU16Para7` float DEFAULT NULL,
  `RsvdU16Para8` float DEFAULT NULL,
  `RsvdU16Para9` float DEFAULT NULL,
  `RsvdU16Para10` float DEFAULT NULL,
  `RsvdU16Para11` float DEFAULT NULL,
  `RsvdU16Para12` float DEFAULT NULL,
  `RsvdU16Para13` float DEFAULT NULL,
  `RsvdU16Para14` float DEFAULT NULL,
  `RsvdU16Para15` float DEFAULT NULL,
  `RsvdU16Para16` float DEFAULT NULL,
  `RsvdU8Para1` float DEFAULT NULL,
  `RsvdU8Para2` float DEFAULT NULL,
  `RsvdU8Para3` float DEFAULT NULL,
  `RsvdU8Para4` float DEFAULT NULL,
  `RsvdU8Para5` float DEFAULT NULL,
  `RsvdU8Para6` float DEFAULT NULL,
  `RsvdU8Para7` float DEFAULT NULL,
  `RsvdU8Para8` float DEFAULT NULL,
  `RsvdU8Para9` float DEFAULT NULL,
  `RsvdU8Para10` float DEFAULT NULL,
  `RsvdU8Para11` float DEFAULT NULL,
  `RsvdU8Para12` float DEFAULT NULL,
  `RsvdU8Para13` float DEFAULT NULL,
  `RsvdU8Para14` float DEFAULT NULL,
  `RsvdU8Para15` float DEFAULT NULL,
  `RsvdU8Para16` float DEFAULT NULL,
  `RsvdU8Para17` float DEFAULT NULL,
  `RsvdU8Para18` float DEFAULT NULL,
  `RsvdU8Para19` float DEFAULT NULL,
  `RsvdU8Para20` float DEFAULT NULL,
  `RsvdU8Para21` float DEFAULT NULL,
  `RsvdU8Para22` float DEFAULT NULL,
  `RsvdU8Para23` float DEFAULT NULL,
  `RsvdU8Para24` float DEFAULT NULL,
  `RsvdU8Para25` float DEFAULT NULL,
  `RsvdU8Para26` float DEFAULT NULL,
  `RsvdU8Para27` float DEFAULT NULL,
  `RsvdU8Para28` float DEFAULT NULL,
  `RsvdU8Para29` float DEFAULT NULL,
  `RsvdU8Para30` float DEFAULT NULL,
  `RsvdU8Para31` float DEFAULT NULL,
  `RsvdU8Para32` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbcnopqcirsvdpara
CREATE TABLE IF NOT EXISTS `enbcnopqcirsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara7` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara14` float DEFAULT NULL,
  `RsvdPara15` float DEFAULT NULL,
  `RsvdPara16` float DEFAULT NULL,
  `RsvdPara17` float DEFAULT NULL,
  `RsvdPara18` float DEFAULT NULL,
  `RsvdPara19` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbcnoprsvdpara
CREATE TABLE IF NOT EXISTS `enbcnoprsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara7` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara14` float DEFAULT NULL,
  `RsvdPara15` float DEFAULT NULL,
  `RsvdPara16` float DEFAULT NULL,
  `RsvdPara17` float DEFAULT NULL,
  `RsvdPara18` float DEFAULT NULL,
  `RsvdPara19` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enblicensealmthd
CREATE TABLE IF NOT EXISTS `enblicensealmthd` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OPRD` float DEFAULT NULL,
  `OTHD` float DEFAULT NULL,
  `RPRD` float DEFAULT NULL,
  `RTHD` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbqcirsvdpara
CREATE TABLE IF NOT EXISTS `enbqcirsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdPara2` float DEFAULT NULL,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara7` float DEFAULT NULL,
  `RsvdPara8` float DEFAULT NULL,
  `RsvdPara9` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara14` float DEFAULT NULL,
  `RsvdPara15` float DEFAULT NULL,
  `RsvdPara16` float DEFAULT NULL,
  `RsvdPara17` float DEFAULT NULL,
  `RsvdPara18` float DEFAULT NULL,
  `RsvdPara19` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enbrsvdpara
CREATE TABLE IF NOT EXISTS `enbrsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RsvdSwPara0` longtext,
  `RsvdSwPara1` longtext,
  `RsvdSwPara2` longtext,
  `RsvdSwPara3` longtext,
  `RsvdPara3` float DEFAULT NULL,
  `RsvdPara4` float DEFAULT NULL,
  `RsvdPara5` float DEFAULT NULL,
  `RsvdPara6` float DEFAULT NULL,
  `RsvdPara10` float DEFAULT NULL,
  `RsvdPara11` float DEFAULT NULL,
  `RsvdPara12` float DEFAULT NULL,
  `RsvdPara13` float DEFAULT NULL,
  `RsvdPara20` float DEFAULT NULL,
  `RsvdPara21` float DEFAULT NULL,
  `RsvdPara22` float DEFAULT NULL,
  `RsvdPara23` float DEFAULT NULL,
  `RsvdPara24` float DEFAULT NULL,
  `RsvdPara25` float DEFAULT NULL,
  `RsvdPara26` float DEFAULT NULL,
  `RsvdPara27` float DEFAULT NULL,
  `RsvdPara28` float DEFAULT NULL,
  `RsvdPara29` float DEFAULT NULL,
  `RsvdPara31` float DEFAULT NULL,
  `RsvdPara32` float DEFAULT NULL,
  `RsvdPara33` float DEFAULT NULL,
  `RsvdPara34` float DEFAULT NULL,
  `RsvdPara35` float DEFAULT NULL,
  `RsvdPara36` float DEFAULT NULL,
  `RsvdPara37` float DEFAULT NULL,
  `RsvdPara38` float DEFAULT NULL,
  `RsvdPara39` float DEFAULT NULL,
  `RsvdPara40` float DEFAULT NULL,
  `RsvdPara41` float DEFAULT NULL,
  `RsvdPara42` float DEFAULT NULL,
  `RsvdPara43` float DEFAULT NULL,
  `RsvdPara44` float DEFAULT NULL,
  `RsvdPara45` float DEFAULT NULL,
  `RsvdPara46` float DEFAULT NULL,
  `RsvdPara47` float DEFAULT NULL,
  `RsvdPara48` float DEFAULT NULL,
  `RsvdPara49` float DEFAULT NULL,
  `RsvdPara50` float DEFAULT NULL,
  `RsvdPara51` float DEFAULT NULL,
  `RsvdPara52` float DEFAULT NULL,
  `RsvdPara53` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.energycon
CREATE TABLE IF NOT EXISTS `energycon` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `MP` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebalgoswitch
CREATE TABLE IF NOT EXISTS `enodebalgoswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `HoAlgoSwitch` longtext,
  `HoModeSwitch` longtext,
  `DlIcicSwitch` longtext,
  `AnrSwitch` longtext,
  `RedirectSwitch` longtext,
  `MroSwitch` longtext,
  `MacAssemblyOptSwitch` longtext,
  `TpeSwitch` longtext,
  `SpidSelectPlmnAlgoSwitch` longtext,
  `UlIcicFreqSwitch` longtext,
  `LcsSwitch` longtext,
  `TrmSwitch` longtext,
  `PciConflictAlmSwitch` longtext,
  `PowerSaveSwitch` longtext,
  `RimSwitch` longtext,
  `RanSharingAnrSwitch` longtext,
  `FreqLayerSwtich` longtext,
  `CmasSwitch` longtext,
  `VQMAlgoSwitch` longtext,
  `UeNumPreemptSwitch` longtext,
  `MultiOpCtrlSwitch` longtext,
  `OverBBUsSwitch` longtext,
  `OperatorSpecificAlgoSwitch` longtext,
  `HoSignalingOptSwitch` longtext,
  `CompatibilityCtrlSwitch` longtext,
  `BlindNcellOptSwitch` longtext,
  `RimOnEcoSwitch` longtext,
  `EutranVoipSupportSwitch` longtext,
  `CaAlgoSwitch` longtext,
  `MlbAlgoSwitch` longtext,
  `HoCommOptSwitch` longtext,
  `HighLoadNetOptSwitch` longtext,
  `SchOptSwitch` longtext,
  `PrachTimeStagSwitch` longtext,
  `HighSpeedRootSeqCSSwitch` longtext,
  `DropPktsStatSwitch` longtext,
  `NCellRankingSwitch` longtext,
  `EUCompactRANSwitch` longtext,
  `BbpCollaborateSwitch` longtext,
  `RootSeqConflictDetSwitch` longtext,
  `IOptAlgoSwitch` longtext,
  `ServiceHoMultiTargetFreqSw` longtext,
  `SFSSwitch` longtext,
  `PciConflictDetectSwitch` longtext,
  `CaLbAlgoSwitch` longtext,
  `UlSchOptSwitch` longtext,
  `CompactRANMultiAPN` longtext,
  `ClkJumpCellReStartSwitch` longtext,
  `E2EVQIAlgoSwitch` longtext,
  `TlFreqFrameOffsetSwitch` longtext,
  `FastEnhancePccAnchorSwitch` longtext,
  `HoWithSccCfgAddBlindSwitch` longtext,
  `objId` float DEFAULT NULL,
  `AntCalibrationTimeSwitch` varchar(255) DEFAULT NULL,
  `PimSwitch` varchar(255) DEFAULT NULL,
  `FltrRptRrcConReqExtdSwitch` varchar(255) DEFAULT NULL,
  `SSRDAlgoSwitch` varchar(255) DEFAULT NULL,
  `IoTClkJumpCellResetSw` varchar(255) DEFAULT NULL,
  `CaAlgoExtSwitch` varchar(255) DEFAULT NULL,
  `UlResManageOptSw` varchar(255) DEFAULT NULL,
  `LTEPreemptNbSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebalmcfg
CREATE TABLE IF NOT EXISTS `enodebalmcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TrafficFaultyDetectPeriod` float DEFAULT NULL,
  `SrvIntfAlmCfgSw` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebautopoweroff
CREATE TABLE IF NOT EXISTS `enodebautopoweroff` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AutoPowerOffSwitch` varchar(255) DEFAULT NULL,
  `PowerOffTime` varchar(255) DEFAULT NULL,
  `PowerOnTime` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebchroutputctrl
CREATE TABLE IF NOT EXISTS `enodebchroutputctrl` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OutPutMode` varchar(255) DEFAULT NULL,
  `CallSampleRate` float DEFAULT NULL,
  `MaxStoreCall` float DEFAULT NULL,
  `CHRUpLoadingTimeSwitch` varchar(255) DEFAULT NULL,
  `SIGOutputMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebciphercap
CREATE TABLE IF NOT EXISTS `enodebciphercap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PrimaryCipherAlgo` varchar(255) DEFAULT NULL,
  `SecondCipherAlgo` varchar(255) DEFAULT NULL,
  `ThirdCipherAlgo` varchar(255) DEFAULT NULL,
  `FourthCipherAlgo` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebconnstatetimer
CREATE TABLE IF NOT EXISTS `enodebconnstatetimer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `S1MessageWaitingTimer` float DEFAULT NULL,
  `X2MessageWaitingTimer` float DEFAULT NULL,
  `UuMessageWaitingTimer` float DEFAULT NULL,
  `Cdma1XrttHoUuPrepareTimer` float DEFAULT NULL,
  `Cdma1XrttHoS1WaitingTimer` float DEFAULT NULL,
  `Cdma1XrttHoCompleteTimer` float DEFAULT NULL,
  `CdmaHrpdHoCompleteTimer` float DEFAULT NULL,
  `CdmaHrpdHoS1WaitingTimer` float DEFAULT NULL,
  `CdmaHrpdHoUuPrepareTimer` float DEFAULT NULL,
  `WaitRrcConnSetupCmpTimer` float DEFAULT NULL,
  `SecCmpWaitingTimer` float DEFAULT NULL,
  `UpUeCapInfoWaitingTimer` float DEFAULT NULL,
  `FirstForwardPacketTimer` float DEFAULT NULL,
  `EndMarkerTimer` float DEFAULT NULL,
  `S1MsgWaitingTimerQci1` float DEFAULT NULL,
  `X2MessageWaitingTimerQci1` float DEFAULT NULL,
  `UuMessageWaitingTimerQci1` float DEFAULT NULL,
  `BearerActivityThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `RrcReestActivityThld` float DEFAULT NULL,
  `EcidWaitingTimer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebfddbbres
CREATE TABLE IF NOT EXISTS `enodebfddbbres` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `eNodeBFddBbResId` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `Capacity` varchar(255) DEFAULT NULL,
  `AdmState` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebflowctrlpara
CREATE TABLE IF NOT EXISTS `enodebflowctrlpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AdaptUnsyncTimerLen` float DEFAULT NULL,
  `AdaptUnsyncUserNumThd` float DEFAULT NULL,
  `DynAcBarPolicyMode` varchar(255) DEFAULT NULL,
  `CpuLoadThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `McpttMsgCntSwitch` varchar(255) DEFAULT NULL,
  `DynAcBarringTrigAllCellSw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebframeoffset
CREATE TABLE IF NOT EXISTS `enodebframeoffset` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TddFrameOffset` float DEFAULT NULL,
  `FddFrameOffset` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebfunction
CREATE TABLE IF NOT EXISTS `enodebfunction` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `eNodeBFunctionName` varchar(255) DEFAULT NULL,
  `ApplicationRef` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `UserLabel` varchar(255) DEFAULT NULL,
  `NermVersion` varchar(255) DEFAULT NULL,
  `ProductVersion` varchar(255) DEFAULT NULL,
  `ProductInterfaceID` varchar(255) DEFAULT NULL,
  `LmtVersion` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebintegritycap
CREATE TABLE IF NOT EXISTS `enodebintegritycap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PrimaryIntegrityAlgo` varchar(255) DEFAULT NULL,
  `SecondIntegrityAlgo` varchar(255) DEFAULT NULL,
  `ThirdIntegrityAlgo` varchar(255) DEFAULT NULL,
  `NullAlgo` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebmlb
CREATE TABLE IF NOT EXISTS `enodebmlb` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `InterFreqIdleMlbMode` varchar(255) DEFAULT NULL,
  `InterFreqIdleMlbInterval` float DEFAULT NULL,
  `InterFreqIdleMlbStaThd` float DEFAULT NULL,
  `DlCaLbMaxCCNum` varchar(255) DEFAULT NULL,
  `CaUeMinDataVolRatio` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebnbpara
CREATE TABLE IF NOT EXISTS `enodebnbpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `NbRsvMinUserNumRatio` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebresmodealgo
CREATE TABLE IF NOT EXISTS `enodebresmodealgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BbpResAutoConfigSw` varchar(255) DEFAULT NULL,
  `TddBbpResDeployAlgo` varchar(255) DEFAULT NULL,
  `FddBbpResDeployAlgo` varchar(255) DEFAULT NULL,
  `BbpResAutoRecfgTimer` float DEFAULT NULL,
  `eNBCellNumCapabilityMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `BbpCpriSharingSwitch` varchar(255) DEFAULT NULL,
  `EnbCellDstDeploySw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebsharingmode
CREATE TABLE IF NOT EXISTS `enodebsharingmode` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ENodeBSharingMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebsondbcfg
CREATE TABLE IF NOT EXISTS `enodebsondbcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `StartTime` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebtddbbres
CREATE TABLE IF NOT EXISTS `enodebtddbbres` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `eNodeBTddBbResId` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `Capacity` varchar(255) DEFAULT NULL,
  `AdmState` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebtgalgcfg
CREATE TABLE IF NOT EXISTS `enodebtgalgcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `FilterCoeffTG` float DEFAULT NULL,
  `StatPeriodTG` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.enodebusparacfg
CREATE TABLE IF NOT EXISTS `enodebusparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UsStrategy` varchar(255) DEFAULT NULL,
  `UsSPIDConfig` float DEFAULT NULL,
  `UsUeInactiveTimer` float DEFAULT NULL,
  `IPDetectInterval` float DEFAULT NULL,
  `UsUeDlPriorityFactor` float DEFAULT NULL,
  `UsLPktUlschPriFactor` float DEFAULT NULL,
  `UsSPktUlschPriFactor` float DEFAULT NULL,
  `UsUESTMSIKeepAliveTimer` float DEFAULT NULL,
  `BigPackageIdentify` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `UlDlUsUserDetectPeriod` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.epgroup
CREATE TABLE IF NOT EXISTS `epgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `EPGROUPID` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `SCTPHOSTREFS` varchar(255) DEFAULT NULL,
  `SCTPPEERREFS` longtext,
  `USERPLANEHOSTREFS` varchar(255) DEFAULT NULL,
  `USERPLANEPEERREFS` longtext,
  `PACKETFILTERSWITCH` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `TYPEFLAG` varchar(255) DEFAULT NULL,
  `LNKPFMSW` varchar(255) DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.equipment
CREATE TABLE IF NOT EXISTS `equipment` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BATIMS` varchar(255) DEFAULT NULL,
  `PAE` varchar(255) DEFAULT NULL,
  `MNTMODE` varchar(255) DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `ET` varchar(255) DEFAULT NULL,
  `MMSETREMARK` varchar(255) DEFAULT NULL,
  `DVAS` varchar(255) DEFAULT NULL,
  `DVAH` float DEFAULT NULL,
  `DVAL` float DEFAULT NULL,
  `ODIID` float DEFAULT NULL,
  `DCALMSW` varchar(255) DEFAULT NULL,
  `EQUIPMENTTY` varchar(255) DEFAULT NULL,
  `PSUFP` varchar(255) DEFAULT NULL,
  `PROTOCOL` varchar(255) DEFAULT NULL,
  `SMARTTRX` varchar(255) DEFAULT NULL,
  `ESN` varchar(255) DEFAULT NULL,
  `SDBBLSD` varchar(255) DEFAULT NULL,
  `APST` varchar(255) DEFAULT NULL,
  `NPST` varchar(255) DEFAULT NULL,
  `SDRCONNSW` varchar(255) DEFAULT NULL,
  `PWRFAILOMCONNENHSW` varchar(255) DEFAULT NULL,
  `OMCONNENHSWCTLTIME` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ethport
CREATE TABLE IF NOT EXISTS `ethport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `PA` varchar(255) DEFAULT NULL,
  `MTU` float DEFAULT NULL,
  `SPEED` varchar(255) DEFAULT NULL,
  `DUPLEX` varchar(255) DEFAULT NULL,
  `ARPPROXY` varchar(255) DEFAULT NULL,
  `FC` varchar(255) DEFAULT NULL,
  `FERAT` float DEFAULT NULL,
  `FERDT` float DEFAULT NULL,
  `RXBCPKTALMOCRTHD` float DEFAULT NULL,
  `RXBCPKTALMCLRTHD` float DEFAULT NULL,
  `FIBERSPEEDMATCH` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucellalmcfg
CREATE TABLE IF NOT EXISTS `eucellalmcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RxNoisePwrThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucellsectoreqm
CREATE TABLE IF NOT EXISTS `eucellsectoreqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SectorEqmId` float DEFAULT NULL,
  `ReferenceSignalPwr` float DEFAULT NULL,
  `BaseBandEqmId` float DEFAULT NULL,
  `ReferenceSignalPwrMargin` float DEFAULT NULL,
  `SectorCpriCompression` varchar(255) DEFAULT NULL,
  `AutoCfgFlag` varchar(255) DEFAULT NULL,
  `WeightNO` float DEFAULT NULL,
  `VisualCellId` float DEFAULT NULL,
  `PrbIdList` varchar(255) DEFAULT NULL,
  `SectorEqmCombineGrpId` float DEFAULT NULL,
  `CellBeamMode` varchar(255) DEFAULT NULL,
  `BeamId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucommcellsectoreqm
CREATE TABLE IF NOT EXISTS `eucommcellsectoreqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `SectorEqmId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucoschcfg
CREATE TABLE IF NOT EXISTS `eucoschcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PrtNodeBaseBandEqmId` float DEFAULT NULL,
  `SchNodeBaseBandEqmId` float DEFAULT NULL,
  `WorkMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucoschdlcompcfg
CREATE TABLE IF NOT EXISTS `eucoschdlcompcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CordInfoEffDelay` float DEFAULT NULL,
  `InterEnbDlCompSwitch` varchar(255) DEFAULT NULL,
  `EuCoSchUeSpec` float DEFAULT NULL,
  `CordInfoEffDelayMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eucoschulicscfg
CREATE TABLE IF NOT EXISTS `eucoschulicscfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UlIcsAlgoSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.euepcsec
CREATE TABLE IF NOT EXISTS `euepcsec` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CiphAlgo` varchar(255) DEFAULT NULL,
  `IntAlgo` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eutranexternalcell
CREATE TABLE IF NOT EXISTS `eutranexternalcell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `DlEarfcn` float DEFAULT NULL,
  `UlEarfcnCfgInd` varchar(255) DEFAULT NULL,
  `PhyCellId` float DEFAULT NULL,
  `Tac` float DEFAULT NULL,
  `CellName` varchar(255) DEFAULT NULL,
  `EutranExternalCellSlaveBand` varchar(255) DEFAULT NULL,
  `RoamingAreaHoInd` varchar(255) DEFAULT NULL,
  `SpecifiedCellFlag` varchar(255) DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `HighSpeedFlag` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `DlFreqOffset` varchar(255) DEFAULT NULL,
  `NbCellFlag` varchar(255) DEFAULT NULL,
  `NclUpdateMode` varchar(255) DEFAULT NULL,
  `SupportEmtcFlag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eutraninterfreqncell
CREATE TABLE IF NOT EXISTS `eutraninterfreqncell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `CellIndividualOffset` varchar(255) DEFAULT NULL,
  `CellQoffset` varchar(255) DEFAULT NULL,
  `NoHoFlag` varchar(255) DEFAULT NULL,
  `NoRmvFlag` varchar(255) DEFAULT NULL,
  `BlindHoPriority` float DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `LocalCellName` varchar(255) DEFAULT NULL,
  `NeighbourCellName` varchar(255) DEFAULT NULL,
  `CellMeasPriority` varchar(255) DEFAULT NULL,
  `OverlapInd` varchar(255) DEFAULT NULL,
  `OverlapRange` varchar(255) DEFAULT NULL,
  `NCellClassLabel` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `AggregationProperty` varchar(255) DEFAULT NULL,
  `OverlapIndicatorExtension` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eutraninternfreq
CREATE TABLE IF NOT EXISTS `eutraninternfreq` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlEarfcn` float DEFAULT NULL,
  `UlEarfcnCfgInd` varchar(255) DEFAULT NULL,
  `CellReselPriorityCfgInd` varchar(255) DEFAULT NULL,
  `CellReselPriority` float DEFAULT NULL,
  `EutranReselTime` float DEFAULT NULL,
  `SpeedDependSPCfgInd` varchar(255) DEFAULT NULL,
  `MeasBandWidth` varchar(255) DEFAULT NULL,
  `QoffsetFreq` varchar(255) DEFAULT NULL,
  `ThreshXhigh` float DEFAULT NULL,
  `ThreshXlow` float DEFAULT NULL,
  `QRxLevMin` float DEFAULT NULL,
  `PmaxCfgInd` varchar(255) DEFAULT NULL,
  `NeighCellConfig` varchar(255) DEFAULT NULL,
  `PresenceAntennaPort1` varchar(255) DEFAULT NULL,
  `InterFreqHoEventType` varchar(255) DEFAULT NULL,
  `ThreshXhighQ` float DEFAULT NULL,
  `ThreshXlowQ` float DEFAULT NULL,
  `QqualMinCfgInd` varchar(255) DEFAULT NULL,
  `ConnFreqPriority` float DEFAULT NULL,
  `MlbTargetInd` varchar(255) DEFAULT NULL,
  `FreqPriBasedHoMeasFlag` varchar(255) DEFAULT NULL,
  `IdleMlbUEReleaseRatio` float DEFAULT NULL,
  `MlbFreqPriority` float DEFAULT NULL,
  `QoffsetFreqConn` varchar(255) DEFAULT NULL,
  `MeasFreqPriority` float DEFAULT NULL,
  `IfHoThdRsrpOffset` float DEFAULT NULL,
  `IfMlbThdRsrpOffset` float DEFAULT NULL,
  `MasterBandFlag` varchar(255) DEFAULT NULL,
  `InterFreqRanSharingInd` varchar(255) DEFAULT NULL,
  `InterFreqHighSpeedFlag` varchar(255) DEFAULT NULL,
  `AnrInd` varchar(255) DEFAULT NULL,
  `VoipPriority` float DEFAULT NULL,
  `PsPriority` float DEFAULT NULL,
  `VolteHoTargetInd` varchar(255) DEFAULT NULL,
  `FreqPriorityForAnr` float DEFAULT NULL,
  `BackoffTargetInd` varchar(255) DEFAULT NULL,
  `MlbInterFreqHoEventType` varchar(255) DEFAULT NULL,
  `MobilityTargetInd` varchar(255) DEFAULT NULL,
  `MlbInterFreqEffiRatio` float DEFAULT NULL,
  `SnrBasedUeSelectionMode` varchar(255) DEFAULT NULL,
  `UlTrafficMlbTargetInd` varchar(255) DEFAULT NULL,
  `UlTrafficMlbPriority` float DEFAULT NULL,
  `MlbInterFreqHoA3Offset` float DEFAULT NULL,
  `IfSrvHoThdRsrpOffset` float DEFAULT NULL,
  `IfSrvHoThdRsrqOffset` float DEFAULT NULL,
  `MlbFreqUlPriority` float DEFAULT NULL,
  `InterFreqMlbDlPrbOffset` float DEFAULT NULL,
  `InterFreqMlbUlPrbOffset` float DEFAULT NULL,
  `NcellNumForAnr` float DEFAULT NULL,
  `MeasPerformanceDemand` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `DlFreqOffset` varchar(255) DEFAULT NULL,
  `IfBackoffThdRsrpOffset` float DEFAULT NULL,
  `IfBackoffThdRsrqOffset` float DEFAULT NULL,
  `VoLTEQualityIfHoTargetInd` varchar(255) DEFAULT NULL,
  `IdleMlbeMtcUEReleaseRatio` float DEFAULT NULL,
  `InterFreqCioAdjLimitCfgInd` varchar(255) DEFAULT NULL,
  `InterFreq4TInd` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.eutranintrafreqncell
CREATE TABLE IF NOT EXISTS `eutranintrafreqncell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `eNodeBId` float DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `CellIndividualOffset` varchar(255) DEFAULT NULL,
  `CellQoffset` varchar(255) DEFAULT NULL,
  `NoHoFlag` varchar(255) DEFAULT NULL,
  `NoRmvFlag` varchar(255) DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `LocalCellName` varchar(255) DEFAULT NULL,
  `NeighbourCellName` varchar(255) DEFAULT NULL,
  `CellMeasPriority` varchar(255) DEFAULT NULL,
  `CellRangeExpansion` float DEFAULT NULL,
  `NCellClassLabel` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `AttachCellSwitch` varchar(255) DEFAULT NULL,
  `HighSpeedCellIndOffset` varchar(255) DEFAULT NULL,
  `VectorCellFlag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.euulcoschcfg
CREATE TABLE IF NOT EXISTS `euulcoschcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UlInterEnbCamcSw` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.extendedqci
CREATE TABLE IF NOT EXISTS `extendedqci` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ExtendedQci` float DEFAULT NULL,
  `ServiceType` varchar(255) DEFAULT NULL,
  `UlschPriorityFactor` float DEFAULT NULL,
  `DlschPriorityFactor` float DEFAULT NULL,
  `UlMinGbr` varchar(255) DEFAULT NULL,
  `DlMinGbr` varchar(255) DEFAULT NULL,
  `PreAllocationWeight` float DEFAULT NULL,
  `InterRatPolicyCfgGroupId` float DEFAULT NULL,
  `PrioritisedBitRate` varchar(255) DEFAULT NULL,
  `LogicalChannelPriority` float DEFAULT NULL,
  `RlcPdcpParaGroupId` float DEFAULT NULL,
  `FreeUserFlag` varchar(255) DEFAULT NULL,
  `FlowCtrlType` varchar(255) DEFAULT NULL,
  `LtePttDedicatedExtendedQci` varchar(255) DEFAULT NULL,
  `LtePttDownlinkPriority` varchar(255) DEFAULT NULL,
  `LtePttUplinkPriority` varchar(255) DEFAULT NULL,
  `ExtQciCounterIndex` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.fddresmode
CREATE TABLE IF NOT EXISTS `fddresmode` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BbCapabilityMode` varchar(255) DEFAULT NULL,
  `SfnCapabilityMode` varchar(255) DEFAULT NULL,
  `BbResExclusiveSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.fltcorrenablecfg
CREATE TABLE IF NOT EXISTS `fltcorrenablecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ENABLERELA` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.fmu
CREATE TABLE IF NOT EXISTS `fmu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MCN` float DEFAULT NULL,
  `MSRN` float DEFAULT NULL,
  `MPN` float DEFAULT NULL,
  `ADDR` float DEFAULT NULL,
  `SBAF` varchar(255) DEFAULT NULL,
  `STC` varchar(255) DEFAULT NULL,
  `TCMODE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ftpclt
CREATE TABLE IF NOT EXISTS `ftpclt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ENCRYMODE` varchar(255) DEFAULT NULL,
  `SPTSTATEFWL` varchar(255) DEFAULT NULL,
  `SSLCERTAUTH` varchar(255) DEFAULT NULL,
  `MINDHLEN` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ftpcltport
CREATE TABLE IF NOT EXISTS `ftpcltport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SERVERIP` varchar(255) DEFAULT NULL,
  `PORT` float DEFAULT NULL,
  `STARTDATAPORT` float DEFAULT NULL,
  `ENDDATAPORT` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtsabiscp
CREATE TABLE IF NOT EXISTS `gbtsabiscp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ABISCPID` float DEFAULT NULL,
  `CPBEARID` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtsbbres
CREATE TABLE IF NOT EXISTS `gbtsbbres` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GBTSBBRESID` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `CAPACITY` varchar(255) DEFAULT NULL,
  `ADMSTATE` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtsenergymgtpara
CREATE TABLE IF NOT EXISTS `gbtsenergymgtpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `BAKPWRSAVMETHOD` varchar(255) DEFAULT NULL,
  `PAADJVOL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtsfunction
CREATE TABLE IF NOT EXISTS `gbtsfunction` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `GBTSFUNCTIONNAME` varchar(255) DEFAULT NULL,
  `APPLICATIONREF` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `NERMVERSION` varchar(255) DEFAULT NULL,
  `PRODUCTVERSION` varchar(255) DEFAULT NULL,
  `interfaceId` varchar(255) DEFAULT NULL,
  `lmtVersion` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtsglobalpara
CREATE TABLE IF NOT EXISTS `gbtsglobalpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `MCSTANDARD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gbtspath
CREATE TABLE IF NOT EXISTS `gbtspath` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PATHID` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.geranexternalcell
CREATE TABLE IF NOT EXISTS `geranexternalcell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `GeranCellId` float DEFAULT NULL,
  `Lac` float DEFAULT NULL,
  `RacCfgInd` varchar(255) DEFAULT NULL,
  `Rac` float DEFAULT NULL,
  `BandIndicator` varchar(255) DEFAULT NULL,
  `GeranArfcn` float DEFAULT NULL,
  `NetworkColourCode` float DEFAULT NULL,
  `BaseStationColourCode` float DEFAULT NULL,
  `DtmInd` varchar(255) DEFAULT NULL,
  `CellName` varchar(255) DEFAULT NULL,
  `CsPsHOInd` varchar(255) DEFAULT NULL,
  `UltraFlashCsfbInd` varchar(255) DEFAULT NULL,
  `RoamingAreaHoInd` varchar(255) DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.geraninterfcfg
CREATE TABLE IF NOT EXISTS `geraninterfcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlGeranIntefRbNum` float DEFAULT NULL,
  `DlRbAvailSinrThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.geranncell
CREATE TABLE IF NOT EXISTS `geranncell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `Lac` float DEFAULT NULL,
  `GeranCellId` float DEFAULT NULL,
  `NoRmvFlag` varchar(255) DEFAULT NULL,
  `NoHoFlag` varchar(255) DEFAULT NULL,
  `BlindHoPriority` float DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `LocalCellName` varchar(255) DEFAULT NULL,
  `NeighbourCellName` varchar(255) DEFAULT NULL,
  `OverlapInd` varchar(255) DEFAULT NULL,
  `NCellMeasPriority` float DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gerannfreqgroup
CREATE TABLE IF NOT EXISTS `gerannfreqgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BcchGroupId` float DEFAULT NULL,
  `GeranVersion` varchar(255) DEFAULT NULL,
  `StartingArfcn` float DEFAULT NULL,
  `BandIndicator` varchar(255) DEFAULT NULL,
  `CellReselPriorityCfgInd` varchar(255) DEFAULT NULL,
  `CellReselPriority` float DEFAULT NULL,
  `PmaxGeranCfgInd` varchar(255) DEFAULT NULL,
  `QRxLevMin` float DEFAULT NULL,
  `ThreshXHigh` float DEFAULT NULL,
  `ThreshXLow` float DEFAULT NULL,
  `OffsetFreq` float DEFAULT NULL,
  `NccPermitted` float DEFAULT NULL,
  `ConnFreqPriority` float DEFAULT NULL,
  `ContinuCoverageIndication` varchar(255) DEFAULT NULL,
  `GeranRanSharingInd` varchar(255) DEFAULT NULL,
  `AnrInd` varchar(255) DEFAULT NULL,
  `FreqPriorityForAnr` float DEFAULT NULL,
  `PmaxGeran` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gerannfreqgrouparfcn
CREATE TABLE IF NOT EXISTS `gerannfreqgrouparfcn` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `BcchGroupId` float DEFAULT NULL,
  `GeranArfcn` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.globalprocswitch
CREATE TABLE IF NOT EXISTS `globalprocswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `X2SonSetupSwitch` longtext,
  `X2SonLinkSetupType` longtext,
  `SriAdaptiveSwitch` longtext,
  `LcgProfile` longtext,
  `RncPoolHoRimSwitch` longtext,
  `UuMsgSimulSendSwitch` longtext,
  `X2BasedUptENodeBCfgSwitch` longtext,
  `TargetOpBasedX2Switch` longtext,
  `X2SonTNLSelectMode` longtext,
  `UtranLoadTransChan` longtext,
  `S1HoInDataFwdSwitch` longtext,
  `RrcReestProtectThd` float DEFAULT NULL,
  `X2SonDeleteTimer` float DEFAULT NULL,
  `RimCodingPolicy` longtext,
  `L2GUHoWithLCapSwitch` longtext,
  `DiffOpWithSameMmecSwitch` longtext,
  `PeerReqBasedX2DelSwitch` longtext,
  `UlPdcpSduRcvStatSendSwitch` longtext,
  `ProtocolMsgOptSwitch` longtext,
  `X2BasedDelNcellCfgSwitch` longtext,
  `X2ServedCellType` longtext,
  `EnbTrigMmeLoadRebalSwitch` longtext,
  `UeRelReSynTimes` longtext,
  `UeRelChkLostSwitch` longtext,
  `UeLinkAbnormalDetectSwitch` longtext,
  `S1DefaultPagingDrxSelect` longtext,
  `EnhancedPhrRptCtrlSw` longtext,
  `EutranLoadTransSwitch` longtext,
  `ProtocolSupportSwitch` longtext,
  `CellTrafficTraceMsgSwitch` longtext,
  `PreferSigExtend` longtext,
  `MmeSelectProcSwitch` longtext,
  `S1OfflineCommitSwitch` longtext,
  `ProtocolCompatibilitySw` longtext,
  `X2SctpEstType` longtext,
  `X2SonDeleteSwitch` longtext,
  `X2SonDeleteTimerforX2Fault` float DEFAULT NULL,
  `X2SonDeleteTimerforX2Usage` float DEFAULT NULL,
  `X2SonDeleteHoInNumThd` float DEFAULT NULL,
  `X2SonDeleteHoOutNumThd` float DEFAULT NULL,
  `VoipWithGapMode` longtext,
  `IntraEnodebHoStaticSw` longtext,
  `MaxSyncUserNumPerBbi` longtext,
  `X2SetupFailureTimetoWait` longtext,
  `X2SetupFailureNumThd` float DEFAULT NULL,
  `X2SetupFailureNumTimer` float DEFAULT NULL,
  `eX2AutoDeleteSwitch` longtext,
  `eX2AutoDeleteTimerforFault` float DEFAULT NULL,
  `X2SonDeleteMode` longtext,
  `X2SonSetupNumThd` float DEFAULT NULL,
  `X2SonSetupTimer` float DEFAULT NULL,
  `SecKeyRecfgSwitch` longtext,
  `X2BasedUptENodeBPolicy` longtext,
  `MMEDomNameMode` longtext,
  `RrcConnPunishThd` float DEFAULT NULL,
  `X2BasedUptNcellCfgSwitch` longtext,
  `HoProcCtrlSwitch` longtext,
  `RrcReestOptSwitch` longtext,
  `S1MMESonSwitch` longtext,
  `PrivateMdtUeSelSwitch` longtext,
  `QciUpdParaCheckSwitch` longtext,
  `UeCompatSwitch` longtext,
  `S1MmePrivFeatureInd` longtext,
  `SctpAbortSmoothSwitch` longtext,
  `MaxUserNumPerCell` longtext,
  `CsfbFlowOptSwitch` longtext,
  `X2InitFailDelSwitch` longtext,
  `X2DynBlacklistAgingTimer` float DEFAULT NULL,
  `eX2AutoSetupSwitch` longtext,
  `eX2DynBlackListSwitch` longtext,
  `eX2DynBlackListAgingTimer` float DEFAULT NULL,
  `PRSMutingCtrlSwitch` longtext,
  `RrcConnReqStatSwitch` longtext,
  `QciParaEffectFlag` longtext,
  `RimFirstMultiReportSwitch` longtext,
  `UuMsgTimeOutRelCause` longtext,
  `SigPoolOptPolicy` longtext,
  `EnhancedRRCReestProtectThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `S1FaultSelfRecoverySwitch` longtext,
  `ItfTypeForNonIdealModeServ` longtext,
  `S1DefaultPagingDrxForNb` longtext,
  `X2SonSecMode` longtext,
  `ProcTypeForNonIdealServData` longtext,
  `Ex2DeleteTimerForEx2Usage` float DEFAULT NULL,
  `SinglePlmnHostSplitSwitch` longtext,
  `InitMsgProtectThld` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocell
CREATE TABLE IF NOT EXISTS `glocell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `BASEBANDPOLICY` varchar(255) DEFAULT NULL,
  `BASEBANDEQMID` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `LOCELLTYPE` varchar(255) DEFAULT NULL,
  `BBEQMPOLICY` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocellalgpara
CREATE TABLE IF NOT EXISTS `glocellalgpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `GUDEGRATEPWRCTRL` varchar(255) DEFAULT NULL,
  `GMSKDELAY` float DEFAULT NULL,
  `DIVERT8PSKDELAY` float DEFAULT NULL,
  `DIVERT16QAMDELAY` float DEFAULT NULL,
  `DIVERT32QAMDELAY` float DEFAULT NULL,
  `GMSKDELAYDYNADJSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocellenergymgtpara
CREATE TABLE IF NOT EXISTS `glocellenergymgtpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `MAINBCCHPWRDTEN` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocellothpara
CREATE TABLE IF NOT EXISTS `glocellothpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `MULTRXBRDSTABMONITUNNSW` varchar(255) DEFAULT NULL,
  `MULTRXBRDUNSTABMONITUNNSW` varchar(255) DEFAULT NULL,
  `IMMOCCUPYPCHOPTSW` varchar(255) DEFAULT NULL,
  `PCHFORBIDRPTLOADSW` varchar(255) DEFAULT NULL,
  `PTCCHPOWEROPTSW` varchar(255) DEFAULT NULL,
  `FIRSTMROPTSW` varchar(255) DEFAULT NULL,
  `SPEECHDELAYOPTSW` varchar(255) DEFAULT NULL,
  `SPEECHBADFRAMEOPTSW` varchar(255) DEFAULT NULL,
  `PAGINGOVLDPROCOPTSW` varchar(255) DEFAULT NULL,
  `RFMODULESTDETECTSW` varchar(255) DEFAULT NULL,
  `NOTRAFFICSELFHEALSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocellrlalmpara
CREATE TABLE IF NOT EXISTS `glocellrlalmpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `ARELWTHD` float DEFAULT NULL,
  `AREDISTHD` float DEFAULT NULL,
  `NOTRASP` float DEFAULT NULL,
  `WLASTIME` float DEFAULT NULL,
  `WLAETIME` float DEFAULT NULL,
  `BWTHD` float DEFAULT NULL,
  `GUNFAULTDECTSW` varchar(255) DEFAULT NULL,
  `GUNCHDROPDECTTHRD` float DEFAULT NULL,
  `GUNCHDROPCLRTHRD` float DEFAULT NULL,
  `GUNBADQUALDETCTTHRD` float DEFAULT NULL,
  `GUNBADQUALCLRTHRD` float DEFAULT NULL,
  `RSSIUNBALANCEDTHRD` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.glocellrsvdpara
CREATE TABLE IF NOT EXISTS `glocellrsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `GLOCELLRSVDPARA1` float DEFAULT NULL,
  `GLOCELLRSVDPARA2` float DEFAULT NULL,
  `GLOCELLRSVDPARA3` float DEFAULT NULL,
  `GLOCELLRSVDPARA4` float DEFAULT NULL,
  `GLOCELLRSVDPARA5` float DEFAULT NULL,
  `GLOCELLRSVDPARA6` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gps
CREATE TABLE IF NOT EXISTS `gps` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GN` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `WPOS` varchar(255) DEFAULT NULL,
  `LONG` float DEFAULT NULL,
  `LAT` float DEFAULT NULL,
  `ALT` float DEFAULT NULL,
  `AGL` float DEFAULT NULL,
  `DURATION` float DEFAULT NULL,
  `PRECISION` float DEFAULT NULL,
  `CABLE_LEN` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `PRI` float DEFAULT NULL,
  `POSCHECKSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gtpu
CREATE TABLE IF NOT EXISTS `gtpu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TIMEOUTTH` float DEFAULT NULL,
  `TIMEOUTCNT` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL,
  `STATICCHK` varchar(255) DEFAULT NULL,
  `ONLYUPIPCHK` varchar(255) DEFAULT NULL,
  `ULGTPUSNPADSW` varchar(255) DEFAULT NULL,
  `DLGTPUSNCHKSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gtranspara
CREATE TABLE IF NOT EXISTS `gtranspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LPSCHSW` varchar(255) DEFAULT NULL,
  `RATECFGTYPE` varchar(255) DEFAULT NULL,
  `SBTIME` float DEFAULT NULL,
  `ARPAGINGTIME` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `OAMTYPE` varchar(255) DEFAULT NULL,
  `BYPASSSWITCH` varchar(255) DEFAULT NULL,
  `STATUS` varchar(255) DEFAULT NULL,
  `OMCHSWITCHBACK` varchar(255) DEFAULT NULL,
  `CPLBSW` varchar(255) DEFAULT NULL,
  `CPLISTENINGMODE` varchar(255) DEFAULT NULL,
  `SCTPRTRPTSW` varchar(255) DEFAULT NULL,
  `IPERRFRMSW` varchar(255) DEFAULT NULL,
  `VLANID` float DEFAULT NULL,
  `FORWARDMODE` varchar(255) DEFAULT NULL,
  `SAMEPORTFORWARDSW` varchar(255) DEFAULT NULL,
  `ONLYOMIP` varchar(255) DEFAULT NULL,
  `DNSPRD` float DEFAULT NULL,
  `EPOBJAUTOCFGSW` varchar(255) DEFAULT NULL,
  `EPCFGMODE` varchar(255) DEFAULT NULL,
  `PMAUTOSW` varchar(255) DEFAULT NULL,
  `DIRECTIPSECPRIMATCHSW` varchar(255) DEFAULT NULL,
  `MODELSELECT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gtransparagw
CREATE TABLE IF NOT EXISTS `gtransparagw` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PRIRULE` varchar(255) DEFAULT NULL,
  `LPSCHSW` varchar(255) DEFAULT NULL,
  `RATECFGTYPE` varchar(255) DEFAULT NULL,
  `SBTIME` float DEFAULT NULL,
  `ARPAGINGTIME` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `OAMTYPE` varchar(255) DEFAULT NULL,
  `IPERRFRMSW` varchar(255) DEFAULT NULL,
  `FORWARDMODE` varchar(255) DEFAULT NULL,
  `SAMEPORTFORWARDSW` varchar(255) DEFAULT NULL,
  `PMAUTOSW` varchar(255) DEFAULT NULL,
  `IPCLKPRI` float DEFAULT NULL,
  `PORTULOBSW` varchar(255) DEFAULT NULL,
  `PORTDLOBSW` varchar(255) DEFAULT NULL,
  `PORTULCACSW` varchar(255) DEFAULT NULL,
  `PORTDLCACSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gtrxgroup
CREATE TABLE IF NOT EXISTS `gtrxgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GTRXGROUPID` float DEFAULT NULL,
  `GLOCELLID` float DEFAULT NULL,
  `SNDMD` varchar(255) DEFAULT NULL,
  `RCVMD` varchar(255) DEFAULT NULL,
  `WORKMODE` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.gtrxgroupsectoreqm
CREATE TABLE IF NOT EXISTS `gtrxgroupsectoreqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GTRXGROUPID` float DEFAULT NULL,
  `SECTOREQMID` float DEFAULT NULL,
  `MAXPOWER` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.highspdadaptionpara
CREATE TABLE IF NOT EXISTS `highspdadaptionpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UserLowSpeedJudgeNum` float DEFAULT NULL,
  `DopplerFactor` float DEFAULT NULL,
  `HighSpeedUserNumThd` float DEFAULT NULL,
  `HoHisJudgeHighThd` float DEFAULT NULL,
  `InterfAvoidCellNum` float DEFAULT NULL,
  `InterfAvoidMode` varchar(255) DEFAULT NULL,
  `InterfBasedPowerOff` float DEFAULT NULL,
  `InterfBasedRbRatio` float DEFAULT NULL,
  `HighspdVectorUpdatePeriod` float DEFAULT NULL,
  `HighspdVectorUpdateSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.homeascomm
CREATE TABLE IF NOT EXISTS `homeascomm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `GapPatternType` varchar(255) DEFAULT NULL,
  `EutranFilterCoeffRsrp` varchar(255) DEFAULT NULL,
  `EutranFilterCoeffRsrq` varchar(255) DEFAULT NULL,
  `GeranFilterCoeff` varchar(255) DEFAULT NULL,
  `SMeasureInd` varchar(255) DEFAULT NULL,
  `SpeedDepParaInd` varchar(255) DEFAULT NULL,
  `UtranFilterCoeffRscp` varchar(255) DEFAULT NULL,
  `UtranFilterCoeffEcn0` varchar(255) DEFAULT NULL,
  `OptHoPreFailPunishTimer` float DEFAULT NULL,
  `ResHoPreFailPunishTimer` float DEFAULT NULL,
  `NonResHoPreFailPunishTimes` float DEFAULT NULL,
  `NonResHoPreFailRetryTimes` float DEFAULT NULL,
  `DedicatedGapPatternType` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.htcdpa
CREATE TABLE IF NOT EXISTS `htcdpa` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `LTCP` float DEFAULT NULL,
  `HTCP` float DEFAULT NULL,
  `TLT` float DEFAULT NULL,
  `DBD` float DEFAULT NULL,
  `NTDI` float DEFAULT NULL,
  `NTDO` float DEFAULT NULL,
  `HTDO` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ikecfg
CREATE TABLE IF NOT EXISTS `ikecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IKELNM` varchar(255) DEFAULT NULL,
  `IKEKLI` float DEFAULT NULL,
  `IKEKLT` float DEFAULT NULL,
  `NATKLI` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL,
  `IPSECSWITCHBACK` varchar(255) DEFAULT NULL,
  `IPSECSBWAITTIME` float DEFAULT NULL,
  `IPSECSBRANDTIME` float DEFAULT NULL,
  `IPSECSOWAITTIME` float DEFAULT NULL,
  `IPSECSORANDTIME` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.imagrp
CREATE TABLE IF NOT EXISTS `imagrp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `IMAGRPN` float DEFAULT NULL,
  `VER` varchar(255) DEFAULT NULL,
  `CLKM` varchar(255) DEFAULT NULL,
  `FRMLEN` varchar(255) DEFAULT NULL,
  `MINLNK` float DEFAULT NULL,
  `DELAY` float DEFAULT NULL,
  `SCRAM` varchar(255) DEFAULT NULL,
  `TS16` varchar(255) DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.imalnk
CREATE TABLE IF NOT EXISTS `imalnk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `IMALNKN` float DEFAULT NULL,
  `IMAGRPN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `IMAGRPSBT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ingchktsk
CREATE TABLE IF NOT EXISTS `ingchktsk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `FLINGCHKSW` varchar(255) DEFAULT NULL,
  `FLINGCHKTM` float DEFAULT NULL,
  `FLINGCHKITEM` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interfreqhogroup
CREATE TABLE IF NOT EXISTS `interfreqhogroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterFreqHoGroupId` float DEFAULT NULL,
  `InterFreqHoA1A2Hyst` float DEFAULT NULL,
  `InterFreqHoA1A2TimeToTrig` varchar(255) DEFAULT NULL,
  `InterFreqHoA1ThdRsrp` float DEFAULT NULL,
  `InterFreqHoA1ThdRsrq` float DEFAULT NULL,
  `InterFreqHoA2ThdRsrp` float DEFAULT NULL,
  `InterFreqHoA2ThdRsrq` float DEFAULT NULL,
  `InterFreqHoA4Hyst` float DEFAULT NULL,
  `InterFreqHoA4ThdRsrp` float DEFAULT NULL,
  `InterFreqHoA4ThdRsrq` float DEFAULT NULL,
  `InterFreqHoA4TimeToTrig` varchar(255) DEFAULT NULL,
  `InterFreqLoadBasedHoA4ThdRsrp` float DEFAULT NULL,
  `InterFreqLoadBasedHoA4ThdRsrq` float DEFAULT NULL,
  `FreqPriInterFreqHoA1ThdRsrp` float DEFAULT NULL,
  `FreqPriInterFreqHoA1ThdRsrq` float DEFAULT NULL,
  `InterFreqHoA3Offset` float DEFAULT NULL,
  `A3InterFreqHoA1ThdRsrp` float DEFAULT NULL,
  `A3InterFreqHoA2ThdRsrp` float DEFAULT NULL,
  `FreqPriInterFreqHoA2ThdRsrp` float DEFAULT NULL,
  `FreqPriInterFreqHoA2ThdRsrq` float DEFAULT NULL,
  `InterFreqMlbA1A2ThdRsrp` float DEFAULT NULL,
  `InterFreqHoA5Thd1Rsrp` float DEFAULT NULL,
  `InterFreqHoA5Thd1Rsrq` float DEFAULT NULL,
  `SrvReqHoA4ThdRsrp` float DEFAULT NULL,
  `SrvReqHoA4ThdRsrq` float DEFAULT NULL,
  `MlbInterFreqHoA5Thd1Rsrp` float DEFAULT NULL,
  `MlbInterFreqHoA5Thd1Rsrq` float DEFAULT NULL,
  `UlBadQualHoA4Offset` float DEFAULT NULL,
  `A3InterFreqHoA1ThdRsrq` float DEFAULT NULL,
  `A3InterFreqHoA2ThdRsrq` float DEFAULT NULL,
  `UlHeavyTrafficMlbA4ThdRsrp` float DEFAULT NULL,
  `UlHeavyTrafficMlbA4ThdRsrq` float DEFAULT NULL,
  `InterFreqHoA1ThdRsrpCE` float DEFAULT NULL,
  `InterFreqHoA2ThdRsrpCE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interratcellshutdown
CREATE TABLE IF NOT EXISTS `interratcellshutdown` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ForceShutdownSwitch` varchar(255) DEFAULT NULL,
  `StartTime` varchar(255) DEFAULT NULL,
  `StopTime` varchar(255) DEFAULT NULL,
  `BearNumThd` float DEFAULT NULL,
  `DlPrbThd` float DEFAULT NULL,
  `UlPrbThd` float DEFAULT NULL,
  `ShutDownType` varchar(255) DEFAULT NULL,
  `HighPriArpThd` float DEFAULT NULL,
  `UtranCellDlLoadThd` float DEFAULT NULL,
  `UtranCellDlLoadOffset` float DEFAULT NULL,
  `UtranCellUlLoadThd` float DEFAULT NULL,
  `UtranCellUlLoadOffset` float DEFAULT NULL,
  `GeranCellLoadThd` float DEFAULT NULL,
  `GeranCellLoadOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathocdma1xrttgroup
CREATE TABLE IF NOT EXISTS `interrathocdma1xrttgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterRatHoCdma1xRttGroupId` float DEFAULT NULL,
  `InterRatHoCdmaB1Hyst` float DEFAULT NULL,
  `InterRatHoCdmaB1ThdPs` float DEFAULT NULL,
  `InterRatHoCdmaB1TimeToTrig` varchar(255) DEFAULT NULL,
  `LdSvBasedHoCdmaB1ThdPs` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathocdmahrpdgroup
CREATE TABLE IF NOT EXISTS `interrathocdmahrpdgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterRatHoCdmaHrpdGroupId` float DEFAULT NULL,
  `InterRatHoCdmaB1Hyst` float DEFAULT NULL,
  `InterRatHoCdmaB1ThdPs` float DEFAULT NULL,
  `InterRatHoCdmaB1TimeToTrig` varchar(255) DEFAULT NULL,
  `LdSvBasedHoCdmaB1ThdPs` float DEFAULT NULL,
  `Cdma2000HrpdB2Thd1Rsrp` float DEFAULT NULL,
  `Cdma2000HrpdB2Thd1Rsrq` float DEFAULT NULL,
  `Cdma2000HrpdNonB2ThdRsrp` float DEFAULT NULL,
  `Cdma2000HrpdNonB2ThdRsrq` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathocomm
CREATE TABLE IF NOT EXISTS `interrathocomm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `InterRatHoMaxRprtCell` float DEFAULT NULL,
  `InterRatHoRprtAmount` varchar(255) DEFAULT NULL,
  `InterRatHoGeranRprtInterval` varchar(255) DEFAULT NULL,
  `InterRatHoUtranB1MeasQuan` varchar(255) DEFAULT NULL,
  `InterRatHoUtranRprtInterval` varchar(255) DEFAULT NULL,
  `InterRatHoCdma1xRttB1MeasQuan` varchar(255) DEFAULT NULL,
  `InterRatCdma1xRttRprtInterval` varchar(255) DEFAULT NULL,
  `InterRatHoCdmaHrpdB1MeasQuan` varchar(255) DEFAULT NULL,
  `InterRatCdmaHrpdRprtInterval` varchar(255) DEFAULT NULL,
  `InterRatHoA1A2TrigQuan` varchar(255) DEFAULT NULL,
  `InterRatHoEventType` varchar(255) DEFAULT NULL,
  `Cdma20001XrttMeasTimer` float DEFAULT NULL,
  `Cdma20001XrttJudgePnNum` float DEFAULT NULL,
  `Cdma2000HrpdFreqSelMode` varchar(255) DEFAULT NULL,
  `Cdma20001XrttFreqSelMode` varchar(255) DEFAULT NULL,
  `CdmaEcsfbPsConcurrentMode` varchar(255) DEFAULT NULL,
  `Cdma20001XrttMeasMode` varchar(255) DEFAULT NULL,
  `Cdma1XrttSectorIdSelMode` varchar(255) DEFAULT NULL,
  `CellInfoMaxUtranCellNum` float DEFAULT NULL,
  `CellInfoMaxGeranCellNum` float DEFAULT NULL,
  `GeranCellNumForEmcRedirect` float DEFAULT NULL,
  `UtranCellNumForEmcRedirect` float DEFAULT NULL,
  `CdmaHrpdSectorIdSelMode` varchar(255) DEFAULT NULL,
  `IRatBlindRedirPlmnCfgSimSw` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathocommgroup
CREATE TABLE IF NOT EXISTS `interrathocommgroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterRatHoCommGroupId` float DEFAULT NULL,
  `InterRatHoA1A2Hyst` float DEFAULT NULL,
  `InterRatHoA1A2TimeToTrig` varchar(255) DEFAULT NULL,
  `InterRatHoA1ThdRsrp` float DEFAULT NULL,
  `InterRatHoA1ThdRsrq` float DEFAULT NULL,
  `InterRatHoA2ThdRsrp` float DEFAULT NULL,
  `InterRatHoA2ThdRsrq` float DEFAULT NULL,
  `GeranB2Thd1Rsrp` float DEFAULT NULL,
  `GeranB2Thd1Rsrq` float DEFAULT NULL,
  `UtranB2Thd1Rsrp` float DEFAULT NULL,
  `UtranB2Thd1Rsrq` float DEFAULT NULL,
  `DelIfMeasA2ThdRsrpOffset` float DEFAULT NULL,
  `DelIfMeasA2ThdRsrqOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathogerangroup
CREATE TABLE IF NOT EXISTS `interrathogerangroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterRatHoGeranGroupId` float DEFAULT NULL,
  `InterRatHoGeranB1Hyst` float DEFAULT NULL,
  `InterRatHoGeranB1Thd` float DEFAULT NULL,
  `InterRatHoGeranB1TimeToTrig` varchar(255) DEFAULT NULL,
  `LdSvBasedHoGeranB1Thd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interrathoutrangroup
CREATE TABLE IF NOT EXISTS `interrathoutrangroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `InterRatHoUtranGroupId` float DEFAULT NULL,
  `InterRatHoUtranB1ThdEcn0` float DEFAULT NULL,
  `InterRatHoUtranB1ThdRscp` float DEFAULT NULL,
  `InterRatHoUtranB1Hyst` float DEFAULT NULL,
  `InterRatHoUtranB1TimeToTrig` varchar(255) DEFAULT NULL,
  `LdSvBasedHoUtranB1ThdEcn0` float DEFAULT NULL,
  `LdSvBasedHoUtranB1ThdRscp` float DEFAULT NULL,
  `HSInterRatHoUtranB1TimeTrig` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.interratpolicycfggroup
CREATE TABLE IF NOT EXISTS `interratpolicycfggroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `InterRatPolicyCfgGroupId` float DEFAULT NULL,
  `UtranHoCfg` varchar(255) DEFAULT NULL,
  `GeranGsmHoCfg` varchar(255) DEFAULT NULL,
  `GeranGprsEdgeHoCfg` varchar(255) DEFAULT NULL,
  `Cdma1xRttHoCfg` varchar(255) DEFAULT NULL,
  `CdmaHrpdHoCfg` varchar(255) DEFAULT NULL,
  `NoHoFlag` varchar(255) DEFAULT NULL,
  `NoFastAnrFlag` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.intrafreqhogroup
CREATE TABLE IF NOT EXISTS `intrafreqhogroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `IntraFreqHoGroupId` float DEFAULT NULL,
  `IntraFreqHoA3Hyst` float DEFAULT NULL,
  `IntraFreqHoA3Offset` float DEFAULT NULL,
  `IntraFreqHoA3TimeToTrig` varchar(255) DEFAULT NULL,
  `HighSpeedA3TimeToTrig` varchar(255) DEFAULT NULL,
  `IntraFreqHoA2ThldRsrpCE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.intrarathocomm
CREATE TABLE IF NOT EXISTS `intrarathocomm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IntraRatHoMaxRprtCell` float DEFAULT NULL,
  `IntraRatHoRprtAmount` varchar(255) DEFAULT NULL,
  `IntraFreqHoA3TrigQuan` varchar(255) DEFAULT NULL,
  `IntraFreqHoA3RprtQuan` varchar(255) DEFAULT NULL,
  `IntraFreqHoRprtInterval` varchar(255) DEFAULT NULL,
  `InterFreqHoA4RprtQuan` varchar(255) DEFAULT NULL,
  `InterFreqHoRprtInterval` varchar(255) DEFAULT NULL,
  `InterFreqHoA1A2TrigQuan` varchar(255) DEFAULT NULL,
  `FreqPriInterFreqHoA1TrigQuan` varchar(255) DEFAULT NULL,
  `InterFreqHoA4TrigQuan` varchar(255) DEFAULT NULL,
  `CovBasedIfHoWaitingTimer` float DEFAULT NULL,
  `A3InterFreqHoA1A2TrigQuan` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `FreqPriHoCandidateUeSelPer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.iopscfg
CREATE TABLE IF NOT EXISTS `iopscfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IopsSwitch` varchar(255) DEFAULT NULL,
  `EnterIopsThd` float DEFAULT NULL,
  `ExitIopsThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ipclklnk
CREATE TABLE IF NOT EXISTS `ipclklnk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LN` float DEFAULT NULL,
  `ICPT` varchar(255) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `CIP` varchar(255) DEFAULT NULL,
  `SIP` varchar(255) DEFAULT NULL,
  `DM` float DEFAULT NULL,
  `DELAYTYPE` varchar(255) DEFAULT NULL,
  `CLASSIDENTIFY` varchar(255) DEFAULT NULL,
  `CLASS0` float DEFAULT NULL,
  `CLASS1` float DEFAULT NULL,
  `CLASS2` float DEFAULT NULL,
  `CLASS3` float DEFAULT NULL,
  `CNM` varchar(255) DEFAULT NULL,
  `CMPST` float DEFAULT NULL,
  `PRI` float DEFAULT NULL,
  `ANNFREQ` varchar(255) DEFAULT NULL,
  `SYNCFREQ` varchar(255) DEFAULT NULL,
  `MASTERPRIO` float DEFAULT NULL,
  `PROFILETYPE` varchar(255) DEFAULT NULL,
  `NEGDURATION` float DEFAULT NULL,
  `IPMODE` varchar(255) DEFAULT NULL,
  `PREHOPIP` varchar(255) DEFAULT NULL,
  `PDELAYREQFREQ` varchar(255) DEFAULT NULL,
  `DEVTYPE` varchar(255) DEFAULT NULL,
  `DSTMLTMACTYPE` varchar(255) DEFAULT NULL,
  `DELAYFREQ` varchar(255) DEFAULT NULL,
  `IPSYNCMODE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ipguard
CREATE TABLE IF NOT EXISTS `ipguard` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ARPSPOOFCHKSW` varchar(255) DEFAULT NULL,
  `ARPLRNSTRICTSW` varchar(255) DEFAULT NULL,
  `INVALIDPKTCHKSW` varchar(255) DEFAULT NULL,
  `IPSECREPLAYCHKSW` varchar(255) DEFAULT NULL,
  `REDIRECTDOSCHKSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ippath
CREATE TABLE IF NOT EXISTS `ippath` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PATHID` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `PATHTYPE` varchar(255) DEFAULT NULL,
  `DSCP` float DEFAULT NULL,
  `LOCALIP` varchar(255) DEFAULT NULL,
  `PEERIP` varchar(255) DEFAULT NULL,
  `QT` varchar(255) DEFAULT NULL,
  `IPMUXSWITCH` varchar(255) DEFAULT NULL,
  `DESCRI` varchar(255) DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `JNRSCGRP` varchar(255) DEFAULT NULL,
  `STATICCHK` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ippmsession
CREATE TABLE IF NOT EXISTS `ippmsession` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IPPMSN` float DEFAULT NULL,
  `LOCALIP` varchar(255) DEFAULT NULL,
  `PEERIP` varchar(255) DEFAULT NULL,
  `IPPMDSCP` float DEFAULT NULL,
  `DIR` varchar(255) DEFAULT NULL,
  `IPPMTYPE` varchar(255) DEFAULT NULL,
  `BINDPATH` varchar(255) DEFAULT NULL,
  `ASSOCIATEUPPATHPLR` varchar(255) DEFAULT NULL,
  `PATHID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.iprt
CREATE TABLE IF NOT EXISTS `iprt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RTIDX` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `RTTYPE` varchar(255) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `DSTIP` varchar(255) DEFAULT NULL,
  `DSTMASK` varchar(255) DEFAULT NULL,
  `NEXTHOP` varchar(255) DEFAULT NULL,
  `MTUSWITCH` varchar(255) DEFAULT NULL,
  `PREF` float DEFAULT NULL,
  `DESCRI` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.iratncellclassmgt
CREATE TABLE IF NOT EXISTS `iratncellclassmgt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RatType` varchar(255) DEFAULT NULL,
  `StatPeriodForNcellClass` float DEFAULT NULL,
  `NCellMeasNumThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.iub
CREATE TABLE IF NOT EXISTS `iub` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IUBID` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `UPEPGROUPID` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `STATICCHKSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.iubcp
CREATE TABLE IF NOT EXISTS `iubcp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CPPT` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `CPPN` float DEFAULT NULL,
  `CPBEARID` float DEFAULT NULL,
  `BELONG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lansw
CREATE TABLE IF NOT EXISTS `lansw` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `BCPKTRATETHD` float DEFAULT NULL,
  `MACAGINGTIME` float DEFAULT NULL,
  `DYNAMICMACLRNSW` varchar(255) DEFAULT NULL,
  `SRCMACATTACKCHKSW` varchar(255) DEFAULT NULL,
  `SRCMACATTACKALMTHD` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lanswitchport
CREATE TABLE IF NOT EXISTS `lanswitchport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PORTIDX` float DEFAULT NULL,
  `PORTTYPE` varchar(255) DEFAULT NULL,
  `USERLABLE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.licensecontrolstrategy
CREATE TABLE IF NOT EXISTS `licensecontrolstrategy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SmoothUpgradeTime` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.licratio
CREATE TABLE IF NOT EXISTS `licratio` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UpLicRatio` float DEFAULT NULL,
  `TrafficSharingType` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lineclk
CREATE TABLE IF NOT EXISTS `lineclk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LN` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `PRI` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lioptatomrule
CREATE TABLE IF NOT EXISTS `lioptatomrule` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AtomRuleID` float DEFAULT NULL,
  `MeasureObjType` varchar(255) DEFAULT NULL,
  `ConditionType` varchar(255) DEFAULT NULL,
  `ThresholdforNumPara` float DEFAULT NULL,
  `MeasureObject` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lioptfeature
CREATE TABLE IF NOT EXISTS `lioptfeature` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IOptFeatureID` float DEFAULT NULL,
  `IOptFeatureName` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lioptfunction
CREATE TABLE IF NOT EXISTS `lioptfunction` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IOptFunctionID` float DEFAULT NULL,
  `IOptFunctionName` varchar(255) DEFAULT NULL,
  `IOptFeatureID` float DEFAULT NULL,
  `MeasureObjType` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lioptrule
CREATE TABLE IF NOT EXISTS `lioptrule` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RuleID` float DEFAULT NULL,
  `ActionType` varchar(255) DEFAULT NULL,
  `AtomRuleRelationType` varchar(255) DEFAULT NULL,
  `Period` float DEFAULT NULL,
  `Action` varchar(255) DEFAULT NULL,
  `IOptFunctionID` float DEFAULT NULL,
  `PenaltyTime` float DEFAULT NULL,
  `AdaptiveRAT` varchar(255) DEFAULT NULL,
  `ActiveStatus` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lioptrulemember
CREATE TABLE IF NOT EXISTS `lioptrulemember` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RuleID` float DEFAULT NULL,
  `AtomRuleID` float DEFAULT NULL,
  `ActiveStatus` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lldpglobal
CREATE TABLE IF NOT EXISTS `lldpglobal` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TXINTVAL` float DEFAULT NULL,
  `HOLDMULTI` float DEFAULT NULL,
  `REINITDELAY` float DEFAULT NULL,
  `DELAY` float DEFAULT NULL,
  `NOTIFYSW` varchar(255) DEFAULT NULL,
  `NOTIFYINTERVAL` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.localethport
CREATE TABLE IF NOT EXISTS `localethport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL,
  `GRATUITOUSARPSW` varchar(255) DEFAULT NULL,
  `IP6SW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.localip
CREATE TABLE IF NOT EXISTS `localip` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IP` varchar(255) DEFAULT NULL,
  `MASK` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.localip6
CREATE TABLE IF NOT EXISTS `localip6` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IP6` varchar(255) DEFAULT NULL,
  `PFXLEN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.localwap
CREATE TABLE IF NOT EXISTS `localwap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL,
  `SSID` varchar(255) DEFAULT NULL,
  `SHOWSSID` varchar(255) DEFAULT NULL,
  `AUTODISABLETIME` float DEFAULT NULL,
  `REGIONCODE` varchar(255) DEFAULT NULL,
  `CHANNEL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.location
CREATE TABLE IF NOT EXISTS `location` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ADDRESS` varchar(255) DEFAULT NULL,
  `ALTITUDE` float DEFAULT NULL,
  `CITY` varchar(255) DEFAULT NULL,
  `CONTACT` varchar(255) DEFAULT NULL,
  `GCDF` varchar(255) DEFAULT NULL,
  `LATITUDEDEGFORMAT` float DEFAULT NULL,
  `LOCATIONID` float DEFAULT NULL,
  `LOCATIONNAME` varchar(255) DEFAULT NULL,
  `LONGITUDEDEGFORMAT` float DEFAULT NULL,
  `OFFICE` varchar(255) DEFAULT NULL,
  `RANGE` float DEFAULT NULL,
  `REGION` varchar(255) DEFAULT NULL,
  `TELEPHONE` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `PRECISE` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `LOCATIONTYPE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.lwamgtcfg
CREATE TABLE IF NOT EXISTS `lwamgtcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `LteWlanAggrW1ThdRssi` float DEFAULT NULL,
  `LteWlanAggrW2Thd1Rssi` float DEFAULT NULL,
  `LteWlanAggrW3ThdRssi` float DEFAULT NULL,
  `LwaW1TimeToTrigger` varchar(255) DEFAULT NULL,
  `LwaW2TimeToTrigger` varchar(255) DEFAULT NULL,
  `LwaW3TimeToTrigger` varchar(255) DEFAULT NULL,
  `WlanMeasHyst` float DEFAULT NULL,
  `CellLwaAlgoSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mainsalarmbind
CREATE TABLE IF NOT EXISTS `mainsalarmbind` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ISDSWITCH` varchar(255) DEFAULT NULL,
  `NMSACN` float DEFAULT NULL,
  `NMSASRN` float DEFAULT NULL,
  `NMSASN` float DEFAULT NULL,
  `NMSAPN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.manresalmrpt
CREATE TABLE IF NOT EXISTS `manresalmrpt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mbmspara
CREATE TABLE IF NOT EXISTS `mbmspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SyncPeriod` float DEFAULT NULL,
  `MbmsCtrlSwitch` varchar(255) DEFAULT NULL,
  `M2DisconnProtectTime` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mmefeaturecfg
CREATE TABLE IF NOT EXISTS `mmefeaturecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `S1InterfaceId` float DEFAULT NULL,
  `MdtEnable` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mpt
CREATE TABLE IF NOT EXISTS `mpt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL,
  `BRDSPEC` varchar(255) DEFAULT NULL,
  `MPTWORKMODE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mptresassignment
CREATE TABLE IF NOT EXISTS `mptresassignment` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `MasterMptAssignmentMode` varchar(255) DEFAULT NULL,
  `SlaveMptAssignmentMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.mro
CREATE TABLE IF NOT EXISTS `mro` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OptPeriod` float DEFAULT NULL,
  `NcellOptThd` float DEFAULT NULL,
  `StatNumThd` float DEFAULT NULL,
  `PingpongTimeThd` float DEFAULT NULL,
  `PingpongRatioThd` float DEFAULT NULL,
  `CoverAbnormalThd` float DEFAULT NULL,
  `ServingRsrpThd` float DEFAULT NULL,
  `NeighborRsrpThd` float DEFAULT NULL,
  `UePingPongNumThd` float DEFAULT NULL,
  `InterFreqMeasTooLateHoThd` float DEFAULT NULL,
  `InterFreqA2RollBackThd` float DEFAULT NULL,
  `MroOptMode` varchar(255) DEFAULT NULL,
  `UnnecInterRatHoOptThd` float DEFAULT NULL,
  `UnnecInterRatHoRsrpThd` float DEFAULT NULL,
  `InterRatA2RsrpMinThd` float DEFAULT NULL,
  `InterRatStatNumThd` float DEFAULT NULL,
  `IntraRatTooEarlyHoRatioThd` float DEFAULT NULL,
  `IntraRatTooLateHoRatioThd` float DEFAULT NULL,
  `IntraRatAbnormalRatioThd` float DEFAULT NULL,
  `InterFreqA2RollBackPeriod` float DEFAULT NULL,
  `IntraRatHoTooEarlyTimeThd` float DEFAULT NULL,
  `InterRatAbnormalHoRatioThd` float DEFAULT NULL,
  `InterRatMeasTooLateHoThd` float DEFAULT NULL,
  `UnnecInterRatHoRatioThd` float DEFAULT NULL,
  `UnnecInterRatHoMeasTime` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nbcelldlschcealgo
CREATE TABLE IF NOT EXISTS `nbcelldlschcealgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `DlInitialTransRptCount` varchar(255) DEFAULT NULL,
  `DlInitialMcs` varchar(255) DEFAULT NULL,
  `UuMessageWaitingTimer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nbcellulschcealgo
CREATE TABLE IF NOT EXISTS `nbcellulschcealgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `CoverageLevel` varchar(255) DEFAULT NULL,
  `UlInitialMcs` varchar(255) DEFAULT NULL,
  `UlInitialTransRptCount` varchar(255) DEFAULT NULL,
  `AckNackTransRptCount` varchar(255) DEFAULT NULL,
  `AckNackTransRptCountMsg4` varchar(255) DEFAULT NULL,
  `NbLogicChSrProhibitTimer` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ncellclassmgt
CREATE TABLE IF NOT EXISTS `ncellclassmgt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `StatPeriodForNcellClass` float DEFAULT NULL,
  `HoAttemptThd` float DEFAULT NULL,
  `CaSCellCfgThd` float DEFAULT NULL,
  `HoSuccThd` float DEFAULT NULL,
  `IntraRatNcellMgtMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ncelldlrsrpmeaspara
CREATE TABLE IF NOT EXISTS `ncelldlrsrpmeaspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlRsrpAutoNCellMeasSwitch` varchar(255) DEFAULT NULL,
  `NCellDlRsrpMeasA3Offset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ncellparacfg
CREATE TABLE IF NOT EXISTS `ncellparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `RatType` varchar(255) DEFAULT NULL,
  `NCellOdDisThd` float DEFAULT NULL,
  `HoStatThd` float DEFAULT NULL,
  `HoSuccThd` float DEFAULT NULL,
  `NcellNumForAnr` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ncellsrsmeaspara
CREATE TABLE IF NOT EXISTS `ncellsrsmeaspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SrsAutoNCellMeasSwitch` varchar(255) DEFAULT NULL,
  `NCellSrsMeasA3Offset` float DEFAULT NULL,
  `NCellMeasSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ne
CREATE TABLE IF NOT EXISTS `ne` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LOCATION` varchar(255) DEFAULT NULL,
  `SWVERSION` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `NERMVERSION` varchar(255) DEFAULT NULL,
  `INTERFACEID` varchar(255) DEFAULT NULL,
  `PRODUCTVERSION` varchar(255) DEFAULT NULL,
  `LMTVERSION` varchar(255) DEFAULT NULL,
  `DID` varchar(255) DEFAULT NULL,
  `SITENAME` varchar(255) DEFAULT NULL,
  `NENAME` varchar(255) DEFAULT NULL,
  `HOTPATCHVERSION` varchar(255) DEFAULT NULL,
  `CLOUDBBID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nemnt
CREATE TABLE IF NOT EXISTS `nemnt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `MNTMODE` varchar(255) DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `ET` varchar(255) DEFAULT NULL,
  `MMSETREMARK` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.node
CREATE TABLE IF NOT EXISTS `node` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PRODUCTTYPE` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `NERMVERSION` varchar(255) DEFAULT NULL,
  `NODEID` float DEFAULT NULL,
  `NODENAME` varchar(255) DEFAULT NULL,
  `WM` varchar(255) DEFAULT NULL,
  `SWVERSION` varchar(255) DEFAULT NULL,
  `HOTPATCHVERSION` varchar(255) DEFAULT NULL,
  `PRODUCTVERSION` varchar(255) DEFAULT NULL,
  `INTERFACEID` varchar(255) DEFAULT NULL,
  `LMTVERSION` varchar(255) DEFAULT NULL,
  `APPVERSION` varchar(255) DEFAULT NULL,
  `APPHOTPATCHVERSION` varchar(255) DEFAULT NULL,
  `LTEMODE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebalgpara
CREATE TABLE IF NOT EXISTS `nodebalgpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `CERSVFOR2MSUSER` float DEFAULT NULL,
  `CCPICPHASESWCTRL` varchar(255) DEFAULT NULL,
  `HSDPASCHPOOLSW` varchar(255) DEFAULT NULL,
  `PWRCHKSW` varchar(255) DEFAULT NULL,
  `CCPICPH2SIRTARMODSW` varchar(255) DEFAULT NULL,
  `ADPRETRANSSW` varchar(255) DEFAULT NULL,
  `FRTRJT` float DEFAULT NULL,
  `INTERBOARDICSW` varchar(255) DEFAULT NULL,
  `SLEEPINGCELLDETECTSW` varchar(255) DEFAULT NULL,
  `UTRPCEXTCNBAPCAP` varchar(255) DEFAULT NULL,
  `RTWPSCALEPRECISIONOPT` varchar(255) DEFAULT NULL,
  `CEDETECTPERIODLEVEL` varchar(255) DEFAULT NULL,
  `ULRESWORKMODE` varchar(255) DEFAULT NULL,
  `CEIMPROVEMENT2MSSW` varchar(255) DEFAULT NULL,
  `INITSINGLEHARQSW` varchar(255) DEFAULT NULL,
  `ULHITHPMACESW` varchar(255) DEFAULT NULL,
  `ULINNERPCABNRISECTRLSW` varchar(255) DEFAULT NULL,
  `KPIFAULTDETECTSW` varchar(255) DEFAULT NULL,
  `KPIFAULTAUTORCVRMTHD` varchar(255) DEFAULT NULL,
  `FIXEDCERSVFOR2MS` float DEFAULT NULL,
  `DLTXPWRMEASPERIOD` varchar(255) DEFAULT NULL,
  `PWRPREDICTSWFORCAC` varchar(255) DEFAULT NULL,
  `DLFPWINADJSW` varchar(255) DEFAULT NULL,
  `HSPAMCINTERBRDSCHSW` varchar(255) DEFAULT NULL,
  `R99FAULTTHD` float DEFAULT NULL,
  `UPAFAULTTHD` float DEFAULT NULL,
  `PROCUNITFAULTTHD` float DEFAULT NULL,
  `PROTECTIONTHD` float DEFAULT NULL,
  `PROCUNITFAULTCLEARTHD` float DEFAULT NULL,
  `ALMOCCURPRD` float DEFAULT NULL,
  `ALMCLEARPRD` float DEFAULT NULL,
  `PDMEASMSW` varchar(255) DEFAULT NULL,
  `ULCCHOLPCMINSIRTARGET` float DEFAULT NULL,
  `HSUPASCHPOOLSW` varchar(255) DEFAULT NULL,
  `ULCOVEXPCFGMOD` varchar(255) DEFAULT NULL,
  `PWRCFGPROTECT` varchar(255) DEFAULT NULL,
  `UPCEIMPROVEFORAMRWBSW` varchar(255) DEFAULT NULL,
  `NBISADAPTIVESW` varchar(255) DEFAULT NULL,
  `UX2SONSETUPSW` varchar(255) DEFAULT NULL,
  `SLEEPINGCELLALARMSW` varchar(255) DEFAULT NULL,
  `SLEEPINGCELLDETECTOPTION` varchar(255) DEFAULT NULL,
  `RFFAULTCTRSW` varchar(255) DEFAULT NULL,
  `CPCFDPCHDTXSW` varchar(255) DEFAULT NULL,
  `IUBCONGESTKPIIMPROVESW` varchar(255) DEFAULT NULL,
  `HSDPASERCELLCHOPTSW` varchar(255) DEFAULT NULL,
  `CELLRTWPEXCESSALMSW` varchar(255) DEFAULT NULL,
  `CELLRTWPEXCESSALMTHD` float DEFAULT NULL,
  `VAMPWRSAVSW` varchar(255) DEFAULT NULL,
  `VAMPWRSAVSTARTTIME` varchar(255) DEFAULT NULL,
  `VAMPWRSAVSTOPTIME` varchar(255) DEFAULT NULL,
  `VAMPWRSAVTXPWRLOWTHD` float DEFAULT NULL,
  `VAMPWRSAVTXPWRHIGHTHD` float DEFAULT NULL,
  `DLAUTORESHUFFLESW` varchar(255) DEFAULT NULL,
  `ADPETADJFLAG` varchar(255) DEFAULT NULL,
  `AMRTRANSBWCOMP` varchar(255) DEFAULT NULL,
  `SFSSWITCH` varchar(255) DEFAULT NULL,
  `KPIFAULTDETECT` varchar(255) DEFAULT NULL,
  `RRUINTELDORFEQBAND` varchar(255) DEFAULT NULL,
  `OPERATORSHARECEMODE` varchar(255) DEFAULT NULL,
  `OPERATORINDEXSELECTION` varchar(255) DEFAULT NULL,
  `NBISRESELPERIOD` float DEFAULT NULL,
  `RTWPIMBALANCEALMTHD` float DEFAULT NULL,
  `NBISINIACTTHD` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebbbres
CREATE TABLE IF NOT EXISTS `nodebbbres` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `NODEBBBRESID` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `CAPACITY` varchar(255) DEFAULT NULL,
  `ADMSTATE` varchar(255) DEFAULT NULL,
  `HCE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebchrlevel
CREATE TABLE IF NOT EXISTS `nodebchrlevel` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `COMMCHREVENTSW` varchar(255) DEFAULT NULL,
  `USERCHREVENTSW` varchar(255) DEFAULT NULL,
  `COMMCHRPRDEVENTSW` varchar(255) DEFAULT NULL,
  `USERCHRPRDEVENTSW` longtext,
  `HSDPAKQIDIAGTHPTHD` varchar(255) DEFAULT NULL,
  `HSUPAKQIDIAGTHPTHD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebclspatimer
CREATE TABLE IF NOT EXISTS `nodebclspatimer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LOWERLIMIT` float DEFAULT NULL,
  `UPPERLIMIT` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebfunction
CREATE TABLE IF NOT EXISTS `nodebfunction` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `NODEBFUNCTIONNAME` varchar(255) DEFAULT NULL,
  `APPLICATIONREF` float DEFAULT NULL,
  `NERMVERSION` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `NODEBID` float DEFAULT NULL,
  `PRODUCTVERSION` varchar(255) DEFAULT NULL,
  `INTERFACEID` varchar(255) DEFAULT NULL,
  `LMTVERSION` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodeblicensealmthd
CREATE TABLE IF NOT EXISTS `nodeblicensealmthd` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `OTHD` float DEFAULT NULL,
  `OPRD` float DEFAULT NULL,
  `RTHD` float DEFAULT NULL,
  `RPRD` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebmulticellgrp
CREATE TABLE IF NOT EXISTS `nodebmulticellgrp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `MULTICELLGRPID` float DEFAULT NULL,
  `MULTICELLGRPTYPE` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `ULOCELLREF` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodeboptdynadjpara
CREATE TABLE IF NOT EXISTS `nodeboptdynadjpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DYNADJSW` varchar(255) DEFAULT NULL,
  `DYNADJSTARTTIME` float DEFAULT NULL,
  `DYNADJENDTIME` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebpoweroutage
CREATE TABLE IF NOT EXISTS `nodebpoweroutage` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ISDSW` varchar(255) DEFAULT NULL,
  `BAKPWRSAVPOLICY` varchar(255) DEFAULT NULL,
  `PHASE1PERIOD` float DEFAULT NULL,
  `PHASE2PERIOD` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebresallocrule
CREATE TABLE IF NOT EXISTS `nodebresallocrule` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RULE` varchar(255) DEFAULT NULL,
  `SW` varchar(255) DEFAULT NULL,
  `CSUSERNUM` float DEFAULT NULL,
  `PSUSERNUM` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `RESREAVESW` varchar(255) DEFAULT NULL,
  `MULTICELLREBUILDSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebrsvdpara
CREATE TABLE IF NOT EXISTS `nodebrsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `NODEBRSVDPARA1` longtext,
  `NODEBRSVDPARA2` longtext,
  `NODEBRSVDPARA3` float DEFAULT NULL,
  `NODEBRSVDPARA4` float DEFAULT NULL,
  `NODEBRSVDPARA5` float DEFAULT NULL,
  `NODEBRSVDPARA6` float DEFAULT NULL,
  `NODEBRSVDPARA7` float DEFAULT NULL,
  `NODEBRSVDPARA8` float DEFAULT NULL,
  `NODEBRSVDPARA9` float DEFAULT NULL,
  `NODEBRSVDPARA10` float DEFAULT NULL,
  `NODEBRSVDPARA11` float DEFAULT NULL,
  `NODEBRSVDPARA12` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `NODEBRSVDPARA13` longtext,
  `NODEBRSVDPARA14` longtext,
  `NODEBRSVDPARA15` float DEFAULT NULL,
  `NODEBRSVDPARA16` float DEFAULT NULL,
  `NODEBRSVDPARA17` float DEFAULT NULL,
  `NODEBRSVDPARA18` float DEFAULT NULL,
  `NODEBRSVDPARA19` float DEFAULT NULL,
  `NODEBRSVDPARA20` float DEFAULT NULL,
  `NODEBRSVDPARA21` float DEFAULT NULL,
  `NODEBRSVDPARA22` float DEFAULT NULL,
  `NODEBRSVDPARA23` float DEFAULT NULL,
  `NODEBRSVDPARA24` float DEFAULT NULL,
  `NODEBRSVDPARA25` float DEFAULT NULL,
  `NODEBRSVDPARA26` float DEFAULT NULL,
  `NODEBRSVDPARA27` float DEFAULT NULL,
  `NODEBRSVDPARA28` float DEFAULT NULL,
  `NODEBRSVDPARA29` float DEFAULT NULL,
  `NODEBRSVDPARA30` float DEFAULT NULL,
  `NODEBRSVDPARA31` float DEFAULT NULL,
  `NODEBRSVDPARA32` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebruleactionpara
CREATE TABLE IF NOT EXISTS `nodebruleactionpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RULEID` float DEFAULT NULL,
  `MOCNAME` varchar(255) DEFAULT NULL,
  `PARANAME` varchar(255) DEFAULT NULL,
  `ADMINSTATE` varchar(255) DEFAULT NULL,
  `PARAINTERVALVALUE` varchar(255) DEFAULT NULL,
  `PARAENUMVALUE` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebsmthpwr
CREATE TABLE IF NOT EXISTS `nodebsmthpwr` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SMTHPWRSWITCH` varchar(255) DEFAULT NULL,
  `STEP` float DEFAULT NULL,
  `PERIOD` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebtrfoverloadthd
CREATE TABLE IF NOT EXISTS `nodebtrfoverloadthd` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TRAFFICOVERLOADOTHD` float DEFAULT NULL,
  `TRAFFICOVERLOADOPRD` float DEFAULT NULL,
  `TRAFFICOVERLOADRTHD` float DEFAULT NULL,
  `TRAFFICOVERLOADRPRD` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.nodebueqosenhance
CREATE TABLE IF NOT EXISTS `nodebueqosenhance` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULGOLDUEBPS` float DEFAULT NULL,
  `ULSILVERUEBPS` float DEFAULT NULL,
  `ULCOPPERUEBPS` float DEFAULT NULL,
  `DLGOLDUEBPS` float DEFAULT NULL,
  `DLSILVERUEBPS` float DEFAULT NULL,
  `DLCOPPERUEBPS` float DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `LOWTRAFFENHUSERNUMTHD` float DEFAULT NULL,
  `TBSIZEFOR10MSLOWTRAFFENH` float DEFAULT NULL,
  `TBSIZEFOR2MSLOWTRAFFENH` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ntpcp
CREATE TABLE IF NOT EXISTS `ntpcp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `IP` varchar(255) DEFAULT NULL,
  `SYNCCYCLE` float DEFAULT NULL,
  `PORT` float DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `AUTHMODE` varchar(255) DEFAULT NULL,
  `KEY` varchar(255) DEFAULT NULL,
  `KEYID` float DEFAULT NULL,
  `MASTERFLAG` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.omch
CREATE TABLE IF NOT EXISTS `omch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `FLAG` varchar(255) DEFAULT NULL,
  `IP` varchar(255) DEFAULT NULL,
  `MASK` varchar(255) DEFAULT NULL,
  `PEERIP` varchar(255) DEFAULT NULL,
  `PEERMASK` varchar(255) DEFAULT NULL,
  `BEAR` varchar(255) DEFAULT NULL,
  `BRT` varchar(255) DEFAULT NULL,
  `CHECKTYPE` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `RTIDX` float DEFAULT NULL,
  `BINDSECONDARYRT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.op
CREATE TABLE IF NOT EXISTS `op` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SWOP` varchar(255) DEFAULT NULL,
  `LOCKST` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.outport
CREATE TABLE IF NOT EXISTS `outport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `NAME` varchar(255) DEFAULT NULL,
  `SW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.paraautooptcfg
CREATE TABLE IF NOT EXISTS `paraautooptcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PUSCHRsrpHighThd4AutoOpt` float DEFAULT NULL,
  `PUCCHPcSINROffset4AutoOpt` float DEFAULT NULL,
  `P0NominalPUCCH4AutoOpt` float DEFAULT NULL,
  `HOTimesThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pccfreqcfg
CREATE TABLE IF NOT EXISTS `pccfreqcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PccDlEarfcn` float DEFAULT NULL,
  `PreferredPccPriority` float DEFAULT NULL,
  `PccA4RsrpThd` float DEFAULT NULL,
  `PccA4RsrqThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pcchcfg
CREATE TABLE IF NOT EXISTS `pcchcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DefaultPagingCycle` varchar(255) DEFAULT NULL,
  `Nb` varchar(255) DEFAULT NULL,
  `PagingSentNum` float DEFAULT NULL,
  `MaxPagingRecordsNum` float DEFAULT NULL,
  `PagingStrategy` varchar(255) DEFAULT NULL,
  `EnhPagingCR` float DEFAULT NULL,
  `EnhPagingSwitch` varchar(255) DEFAULT NULL,
  `DefaultPagingCycleForNb` varchar(255) DEFAULT NULL,
  `NbForNbIoT` varchar(255) DEFAULT NULL,
  `MaxNumRepetitionForPaging` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pdcprohcpara
CREATE TABLE IF NOT EXISTS `pdcprohcpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RohcSwitch` varchar(255) DEFAULT NULL,
  `HighestMode` varchar(255) DEFAULT NULL,
  `Profiles` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pdschcfg
CREATE TABLE IF NOT EXISTS `pdschcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ReferenceSignalPwr` float DEFAULT NULL,
  `Pb` float DEFAULT NULL,
  `ReferenceSignalPwrMargin` float DEFAULT NULL,
  `TxPowerOffsetAnt0` float DEFAULT NULL,
  `TxPowerOffsetAnt1` float DEFAULT NULL,
  `TxPowerOffsetAnt2` float DEFAULT NULL,
  `TxPowerOffsetAnt3` float DEFAULT NULL,
  `TxChnPowerCfgSw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.peerclk
CREATE TABLE IF NOT EXISTS `peerclk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `PS` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.peu
CREATE TABLE IF NOT EXISTS `peu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.phichcfg
CREATE TABLE IF NOT EXISTS `phichcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PhichDuration` varchar(255) DEFAULT NULL,
  `PhichResource` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.phyport
CREATE TABLE IF NOT EXISTS `phyport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `ADMINISTRATIVESTATE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.plrthreshold
CREATE TABLE IF NOT EXISTS `plrthreshold` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PLRAT` float DEFAULT NULL,
  `PLRDT` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pmtucfg
CREATE TABLE IF NOT EXISTS `pmtucfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `UDPPORT` float DEFAULT NULL,
  `DETECTPERIOD` float DEFAULT NULL,
  `AGINGPERIOD` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.portpolicy
CREATE TABLE IF NOT EXISTS `portpolicy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `USBACCESSSECPOLICY` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pri2que
CREATE TABLE IF NOT EXISTS `pri2que` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `PRI0` float DEFAULT NULL,
  `PRI1` float DEFAULT NULL,
  `PRI2` float DEFAULT NULL,
  `PRI3` float DEFAULT NULL,
  `PRI4` float DEFAULT NULL,
  `PRI5` float DEFAULT NULL,
  `PRI6` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.privatecabandcomb
CREATE TABLE IF NOT EXISTS `privatecabandcomb` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PrivateCaCombId` float DEFAULT NULL,
  `MaxAggregatedBw` float DEFAULT NULL,
  `BwCombSetId` float DEFAULT NULL,
  `CombBand1Id` float DEFAULT NULL,
  `CombBand2Id` float DEFAULT NULL,
  `CombBand3Id` float DEFAULT NULL,
  `CombBand4Id` float DEFAULT NULL,
  `CombBand1Bw` varchar(255) DEFAULT NULL,
  `CombBand2Bw` varchar(255) DEFAULT NULL,
  `CombBand3Bw` varchar(255) DEFAULT NULL,
  `CombBand4Bw` varchar(255) DEFAULT NULL,
  `CombBand5Id` float DEFAULT NULL,
  `CombBand5Bw` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.psuis
CREATE TABLE IF NOT EXISTS `psuis` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PSUISS` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pucchcfg
CREATE TABLE IF NOT EXISTS `pucchcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DeltaShift` varchar(255) DEFAULT NULL,
  `NaSriChNum` float DEFAULT NULL,
  `CqiRbNum` float DEFAULT NULL,
  `PucchExtendedRBNum` float DEFAULT NULL,
  `Format1ChAllocMode` varchar(255) DEFAULT NULL,
  `SriPeriodAdaptive` varchar(255) DEFAULT NULL,
  `Format2ChAllocMode` varchar(255) DEFAULT NULL,
  `PucchAllocPolicy` varchar(255) DEFAULT NULL,
  `Format3RbNum` float DEFAULT NULL,
  `Max2CCAckChNum` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.puschcfg
CREATE TABLE IF NOT EXISTS `puschcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Nsb` float DEFAULT NULL,
  `HoppingMode` varchar(255) DEFAULT NULL,
  `HoppingOffset` float DEFAULT NULL,
  `GroupHoppingEnabled` varchar(255) DEFAULT NULL,
  `GroupAssignPUSCH` float DEFAULT NULL,
  `SeqHoppingEnabled` varchar(255) DEFAULT NULL,
  `CyclicShift` float DEFAULT NULL,
  `Qam64Enabled` varchar(255) DEFAULT NULL,
  `R12Qam64Enabled` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.puschparam
CREATE TABLE IF NOT EXISTS `puschparam` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DeltaOffsetCqiIndex` float DEFAULT NULL,
  `DeltaOffsetRiIndex` float DEFAULT NULL,
  `DeltaOffsetAckIndex` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.pwdpolicy
CREATE TABLE IF NOT EXISTS `pwdpolicy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PWDMINLEN` float DEFAULT NULL,
  `COMPLICACY` varchar(255) DEFAULT NULL,
  `PASSREPLMT` float DEFAULT NULL,
  `MAXPERIOD` float DEFAULT NULL,
  `MINPERIOD` float DEFAULT NULL,
  `MAXMISSTIMES` float DEFAULT NULL,
  `AUTOUNLOCKTIME` float DEFAULT NULL,
  `RESETINTERVAL` float DEFAULT NULL,
  `PWDEXPRT` float DEFAULT NULL,
  `FirstLoginMustModPWD` varchar(255) DEFAULT NULL,
  `MAXREPEATCHARTIMES` float DEFAULT NULL,
  `DICTCHKPWD` varchar(255) DEFAULT NULL,
  `MAXCCUN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.qcipara
CREATE TABLE IF NOT EXISTS `qcipara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Qci` float DEFAULT NULL,
  `PreAllocationWeight` float DEFAULT NULL,
  `InterRatPolicyCfgGroupId` float DEFAULT NULL,
  `LogicalChannelPriority` float DEFAULT NULL,
  `RlcPdcpParaGroupId` float DEFAULT NULL,
  `UeInactiveTimerForQci` float DEFAULT NULL,
  `UlSynTimerForQci` float DEFAULT NULL,
  `UeInactivityTimerDynDrxQci` float DEFAULT NULL,
  `UeInactiveTimerPri` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `PrioritisedBitRate` varchar(255) DEFAULT NULL,
  `FlowCtrlType` varchar(255) DEFAULT NULL,
  `UlschPriorityFactor` float DEFAULT NULL,
  `DlschPriorityFactor` float DEFAULT NULL,
  `UlMinGbr` varchar(255) DEFAULT NULL,
  `DlMinGbr` varchar(255) DEFAULT NULL,
  `EmtcModeARlcParaGroupId` float DEFAULT NULL,
  `EmtcModeBRlcParaGroupId` float DEFAULT NULL,
  `NbRlcPdcpParaGroupId` float DEFAULT NULL,
  `CiotUeInactiveTimer` float DEFAULT NULL,
  `ServiceType` varchar(255) DEFAULT NULL,
  `FreeUserFlag` varchar(255) DEFAULT NULL,
  `LtePttDedicatedExtendedQci` varchar(255) DEFAULT NULL,
  `ExtQciCounterIndex` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.qoehocommoncfg
CREATE TABLE IF NOT EXISTS `qoehocommoncfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `QoEBasedHandoverStat` float DEFAULT NULL,
  `QoEBasedHandoverLast` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rachcfg
CREATE TABLE IF NOT EXISTS `rachcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `PwrRampingStep` varchar(255) DEFAULT NULL,
  `PreambInitRcvTargetPwr` varchar(255) DEFAULT NULL,
  `MessageSizeGroupA` varchar(255) DEFAULT NULL,
  `PrachFreqOffset` float DEFAULT NULL,
  `PrachConfigIndexCfgInd` varchar(255) DEFAULT NULL,
  `PreambleTransMax` varchar(255) DEFAULT NULL,
  `ContentionResolutionTimer` varchar(255) DEFAULT NULL,
  `MaxHarqMsg3Tx` float DEFAULT NULL,
  `RandomPreambleRatio` varchar(255) DEFAULT NULL,
  `RaPreambleGrpARatio` float DEFAULT NULL,
  `PrachFreqOffsetStrategy` varchar(255) DEFAULT NULL,
  `PrachStartTimeCfgInd` varchar(255) DEFAULT NULL,
  `NbCyclicPrefixLength` varchar(255) DEFAULT NULL,
  `NbRsrpFirstThreshold` float DEFAULT NULL,
  `NbRsrpSecondThreshold` float DEFAULT NULL,
  `UeRaInfoReportRatioThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ret
CREATE TABLE IF NOT EXISTS `ret` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DEVICENO` float DEFAULT NULL,
  `DEVICENAME` varchar(255) DEFAULT NULL,
  `CTRLCN` float DEFAULT NULL,
  `CTRLSRN` float DEFAULT NULL,
  `CTRLSN` float DEFAULT NULL,
  `VENDORCODE` varchar(255) DEFAULT NULL,
  `SERIALNO` varchar(255) DEFAULT NULL,
  `RETTYPE` varchar(255) DEFAULT NULL,
  `POLARTYPE` varchar(255) DEFAULT NULL,
  `SCENARIO` varchar(255) DEFAULT NULL,
  `SUBUNITNUM` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.retdevicedata
CREATE TABLE IF NOT EXISTS `retdevicedata` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DEVICENO` float DEFAULT NULL,
  `SUBUNITNO` float DEFAULT NULL,
  `BEARING` float DEFAULT NULL,
  `MODELNO` varchar(255) DEFAULT NULL,
  `BSID` varchar(255) DEFAULT NULL,
  `BEAMWIDTH1` float DEFAULT NULL,
  `BEAMWIDTH2` float DEFAULT NULL,
  `BEAMWIDTH3` float DEFAULT NULL,
  `BEAMWIDTH4` float DEFAULT NULL,
  `GAIN1` float DEFAULT NULL,
  `GAIN2` float DEFAULT NULL,
  `GAIN3` float DEFAULT NULL,
  `GAIN4` float DEFAULT NULL,
  `DATE` varchar(255) DEFAULT NULL,
  `TILT` float DEFAULT NULL,
  `INSTALLERID` varchar(255) DEFAULT NULL,
  `BAND1` varchar(255) DEFAULT NULL,
  `BAND2` varchar(255) DEFAULT NULL,
  `BAND3` varchar(255) DEFAULT NULL,
  `BAND4` varchar(255) DEFAULT NULL,
  `SECTORID` varchar(255) DEFAULT NULL,
  `SERIALNO` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.retport
CREATE TABLE IF NOT EXISTS `retport` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` varchar(255) DEFAULT NULL,
  `PWRSWITCH` varchar(255) DEFAULT NULL,
  `THRESHOLDTYPE` varchar(255) DEFAULT NULL,
  `UOTHD` float DEFAULT NULL,
  `UCTHD` float DEFAULT NULL,
  `OOTHD` float DEFAULT NULL,
  `OCTHD` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.retsubunit
CREATE TABLE IF NOT EXISTS `retsubunit` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `DEVICENO` float DEFAULT NULL,
  `SUBUNITNO` float DEFAULT NULL,
  `CONNCN1` float DEFAULT NULL,
  `CONNSRN1` float DEFAULT NULL,
  `CONNSN1` float DEFAULT NULL,
  `CONNPN1` varchar(255) DEFAULT NULL,
  `CONNCN2` float DEFAULT NULL,
  `CONNSRN2` float DEFAULT NULL,
  `CONNSN2` float DEFAULT NULL,
  `CONNPN2` varchar(255) DEFAULT NULL,
  `TILT` varchar(255) DEFAULT NULL,
  `AER` float DEFAULT NULL,
  `SUBNAME` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rfu
CREATE TABLE IF NOT EXISTS `rfu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `ADMSTATE` varchar(255) DEFAULT NULL,
  `IFFREQ` float DEFAULT NULL,
  `ALMPROCSW` varchar(255) DEFAULT NULL,
  `ALMPROCTHRHLD` float DEFAULT NULL,
  `ALMTHRHLD` float DEFAULT NULL,
  `PS` float DEFAULT NULL,
  `RCN` float DEFAULT NULL,
  `TP` varchar(255) DEFAULT NULL,
  `RS` varchar(255) DEFAULT NULL,
  `RXNUM` float DEFAULT NULL,
  `TXNUM` float DEFAULT NULL,
  `RFTXSIGNDETECTSW` varchar(255) DEFAULT NULL,
  `RFTXSIGNDETECTPERIOD` varchar(255) DEFAULT NULL,
  `RFTXSIGNDETECTTHLD` varchar(255) DEFAULT NULL,
  `RT` varchar(255) DEFAULT NULL,
  `IFOFFSET` float DEFAULT NULL,
  `RFDS` float DEFAULT NULL,
  `FMBWH` varchar(255) DEFAULT NULL,
  `LCPSW` varchar(255) DEFAULT NULL,
  `FLAG` float DEFAULT NULL,
  `RUSPEC` varchar(255) DEFAULT NULL,
  `RFCONNCN2` varchar(255) DEFAULT NULL,
  `RFCONNSN2` varchar(255) DEFAULT NULL,
  `RFCONNSRN2` varchar(255) DEFAULT NULL,
  `RFCONNTYPE` varchar(255) DEFAULT NULL,
  `PAEFFSWITCH` varchar(255) DEFAULT NULL,
  `SCR` varchar(255) DEFAULT NULL,
  `RXFREQBANDMUTUALSW` varchar(255) DEFAULT NULL,
  `MNTMODE` varchar(255) DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `ET` varchar(255) DEFAULT NULL,
  `MMSETREMARK` varchar(255) DEFAULT NULL,
  `LEDSW` varchar(255) DEFAULT NULL,
  `CUSTOMEDRFSPECSW` varchar(255) DEFAULT NULL,
  `PSGID` float DEFAULT NULL,
  `WIFISW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rlcpdcpparagroup
CREATE TABLE IF NOT EXISTS `rlcpdcpparagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RlcPdcpParaGroupId` float DEFAULT NULL,
  `DiscardTimer` varchar(255) DEFAULT NULL,
  `RlcMode` varchar(255) DEFAULT NULL,
  `PdcpSnSize` varchar(255) DEFAULT NULL,
  `UlRlcSnSize` varchar(255) DEFAULT NULL,
  `DlRlcSnSize` varchar(255) DEFAULT NULL,
  `UeUmReorderingTimer` varchar(255) DEFAULT NULL,
  `ENodeBUmReorderingTimer` varchar(255) DEFAULT NULL,
  `NonsptUmUeAdaptSwitch` varchar(255) DEFAULT NULL,
  `UlDlDiscardtimerSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `UeMaxRetxThreshold` varchar(255) DEFAULT NULL,
  `ENodeBMaxRetxThreshold` varchar(255) DEFAULT NULL,
  `UePollByte` varchar(255) DEFAULT NULL,
  `ENodeBPollByte` varchar(255) DEFAULT NULL,
  `UePollPdu` varchar(255) DEFAULT NULL,
  `ENodeBPollPdu` varchar(255) DEFAULT NULL,
  `UePollRetransmitTimer` varchar(255) DEFAULT NULL,
  `ENodeBPollRetransmitTimer` varchar(255) DEFAULT NULL,
  `UeStatusProhibitTimer` varchar(255) DEFAULT NULL,
  `ENodeBStatusProhibitTimer` varchar(255) DEFAULT NULL,
  `UeAmReorderingTimer` varchar(255) DEFAULT NULL,
  `ENodeBAmReorderingTimer` varchar(255) DEFAULT NULL,
  `PdcpStatusRptReq` varchar(255) DEFAULT NULL,
  `RlcParaAdaptSwitch` varchar(255) DEFAULT NULL,
  `eNodeBPollRtrTimerPreset` varchar(255) DEFAULT NULL,
  `eNodeBStatProhTimerPreset` varchar(255) DEFAULT NULL,
  `UePollRtrTimerPreset` varchar(255) DEFAULT NULL,
  `UeStatProhTimerPreset` varchar(255) DEFAULT NULL,
  `CaUeRlcParaAdptiveThd` float DEFAULT NULL,
  `CaUeReorderingTimer` varchar(255) DEFAULT NULL,
  `CaUeStatProhTimer` varchar(255) DEFAULT NULL,
  `AmPdcpSnSize` varchar(255) DEFAULT NULL,
  `ENodeBReorderingTimerAdapt` varchar(255) DEFAULT NULL,
  `UePdcpReorderingTimer` varchar(255) DEFAULT NULL,
  `CatType` varchar(255) DEFAULT NULL,
  `NbPdcpDiscardTimer` varchar(255) DEFAULT NULL,
  `NbDlPdcpDiscardTimer` varchar(255) DEFAULT NULL,
  `NbEnodebPollRetxTimer` varchar(255) DEFAULT NULL,
  `NbUePollRetxTimer` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rrcconnstatetimer
CREATE TABLE IF NOT EXISTS `rrcconnstatetimer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `T302` float DEFAULT NULL,
  `T320ForLoadBalance` varchar(255) DEFAULT NULL,
  `T304ForEutran` varchar(255) DEFAULT NULL,
  `T304ForGeran` varchar(255) DEFAULT NULL,
  `T320ForOther` varchar(255) DEFAULT NULL,
  `UeInactiveTimer` float DEFAULT NULL,
  `UlSynTimer` float DEFAULT NULL,
  `FilterReptRrcConnReqTimer` float DEFAULT NULL,
  `UeInactivityTimerDynDrx` float DEFAULT NULL,
  `UlSynTimerDynDrx` float DEFAULT NULL,
  `UeInactiveTimerQci1` float DEFAULT NULL,
  `UeInactTimerDynDrxQci1` float DEFAULT NULL,
  `RrcConnRelTimer` float DEFAULT NULL,
  `DRXRrcConnRelTimerOffset` float DEFAULT NULL,
  `SRLTERrcConnRelTimerOffset` float DEFAULT NULL,
  `ReptRrcReestProtectTimer` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `NBUeInactiveTimer` float DEFAULT NULL,
  `ExtendedWaitTime` float DEFAULT NULL,
  `MTCUeInactiveTimer` float DEFAULT NULL,
  `PowerPrefIndicationTimer` varchar(255) DEFAULT NULL,
  `OutOfSyncRelTimerOffset` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rru
CREATE TABLE IF NOT EXISTS `rru` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `ADMSTATE` varchar(255) DEFAULT NULL,
  `IFFREQ` float DEFAULT NULL,
  `ALMPROCSW` varchar(255) DEFAULT NULL,
  `ALMPROCTHRHLD` float DEFAULT NULL,
  `ALMTHRHLD` float DEFAULT NULL,
  `PS` float DEFAULT NULL,
  `RCN` float DEFAULT NULL,
  `TP` varchar(255) DEFAULT NULL,
  `RS` varchar(255) DEFAULT NULL,
  `RXNUM` float DEFAULT NULL,
  `TXNUM` float DEFAULT NULL,
  `RFTXSIGNDETECTSW` varchar(255) DEFAULT NULL,
  `RFTXSIGNDETECTPERIOD` varchar(255) DEFAULT NULL,
  `RFTXSIGNDETECTTHLD` varchar(255) DEFAULT NULL,
  `RN` varchar(255) DEFAULT NULL,
  `RT` varchar(255) DEFAULT NULL,
  `IFOFFSET` float DEFAULT NULL,
  `RFDS` float DEFAULT NULL,
  `FMBWH` varchar(255) DEFAULT NULL,
  `LCPSW` varchar(255) DEFAULT NULL,
  `FLAG` float DEFAULT NULL,
  `RUSPEC` varchar(255) DEFAULT NULL,
  `RFCONNCN2` varchar(255) DEFAULT NULL,
  `RFCONNSN2` varchar(255) DEFAULT NULL,
  `RFCONNSRN2` varchar(255) DEFAULT NULL,
  `RFCONNTYPE` varchar(255) DEFAULT NULL,
  `PAEFFSWITCH` varchar(255) DEFAULT NULL,
  `SCR` varchar(255) DEFAULT NULL,
  `RXFREQBANDMUTUALSW` varchar(255) DEFAULT NULL,
  `REMOTEFLAG` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `RFDCPWROFFALMDETECTSW` varchar(255) DEFAULT NULL,
  `BATTOUTPUNDERVOLTTHLD` float DEFAULT NULL,
  `MNTMODE` varchar(255) DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `ET` varchar(255) DEFAULT NULL,
  `MMSETREMARK` varchar(255) DEFAULT NULL,
  `LEDSW` varchar(255) DEFAULT NULL,
  `CUSTOMEDRFSPECSW` varchar(255) DEFAULT NULL,
  `PSGID` float DEFAULT NULL,
  `WIFISW` varchar(255) DEFAULT NULL,
  `LOCATIONNAME` varchar(255) DEFAULT NULL,
  `CIRCUITBREAKERVALUE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rruchain
CREATE TABLE IF NOT EXISTS `rruchain` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RCN` float DEFAULT NULL,
  `TT` varchar(255) DEFAULT NULL,
  `BM` varchar(255) DEFAULT NULL,
  `HCN` float DEFAULT NULL,
  `HSRN` float DEFAULT NULL,
  `HSN` float DEFAULT NULL,
  `HPN` float DEFAULT NULL,
  `BRKPOS1` varchar(255) DEFAULT NULL,
  `BRKPOS2` varchar(255) DEFAULT NULL,
  `AT` varchar(255) DEFAULT NULL,
  `CR` varchar(255) DEFAULT NULL,
  `LSN` float DEFAULT NULL,
  `PROTOCOL` varchar(255) DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `CONNPORTNUM` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rrujointcalparacfg
CREATE TABLE IF NOT EXISTS `rrujointcalparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TxChnCalSwitch` varchar(255) DEFAULT NULL,
  `TxChnCalTime` varchar(255) DEFAULT NULL,
  `TxChnCalPeriod` float DEFAULT NULL,
  `TxChnAntSpacing` float DEFAULT NULL,
  `AauPassivePortCalibPeriod` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rscgrp
CREATE TABLE IF NOT EXISTS `rscgrp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `BEAR` varchar(255) DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `RSCGRPID` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `RU` varchar(255) DEFAULT NULL,
  `TXBW` float DEFAULT NULL,
  `RXBW` float DEFAULT NULL,
  `TXCBS` float DEFAULT NULL,
  `TXEBS` float DEFAULT NULL,
  `OID` float DEFAULT NULL,
  `WEIGHT` float DEFAULT NULL,
  `TXCIR` float DEFAULT NULL,
  `RXCIR` float DEFAULT NULL,
  `TXPIR` float DEFAULT NULL,
  `RXPIR` float DEFAULT NULL,
  `TXPBS` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rscgrpalg
CREATE TABLE IF NOT EXISTS `rscgrpalg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `BEAR` varchar(255) DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `RSCGRPID` varchar(255) DEFAULT NULL,
  `TXSSW` varchar(255) DEFAULT NULL,
  `TXBWASW` varchar(255) DEFAULT NULL,
  `RXBWASW` varchar(255) DEFAULT NULL,
  `PLRDTH` float DEFAULT NULL,
  `DDTH` float DEFAULT NULL,
  `TXBWAMIN` float DEFAULT NULL,
  `RXBWAMIN` float DEFAULT NULL,
  `TCSW` varchar(255) DEFAULT NULL,
  `PQN` float DEFAULT NULL,
  `CTTH` float DEFAULT NULL,
  `CCTTH` float DEFAULT NULL,
  `TXRSVBW` float DEFAULT NULL,
  `RXRSVBW` float DEFAULT NULL,
  `DROPPKTNUM` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.rxbranch
CREATE TABLE IF NOT EXISTS `rxbranch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `RXNO` float DEFAULT NULL,
  `RXSW` varchar(255) DEFAULT NULL,
  `ATTEN` float DEFAULT NULL,
  `RTWPINITADJ0` float DEFAULT NULL,
  `RTWPINITADJ1` float DEFAULT NULL,
  `RTWPINITADJ2` float DEFAULT NULL,
  `RTWPINITADJ3` float DEFAULT NULL,
  `RTWPINITADJ4` float DEFAULT NULL,
  `RTWPINITADJ5` float DEFAULT NULL,
  `RTWPINITADJ6` float DEFAULT NULL,
  `RTWPINITADJ7` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.s1
CREATE TABLE IF NOT EXISTS `s1` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `S1Id` float DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `UpEpGroupId` float DEFAULT NULL,
  `MmeRelease` varchar(255) DEFAULT NULL,
  `UserLabel` varchar(255) DEFAULT NULL,
  `EpGroupCfgFlag` varchar(255) DEFAULT NULL,
  `Priority` float DEFAULT NULL,
  `CnOpSharingGroupId` float DEFAULT NULL,
  `MdtEnable` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.s1interface
CREATE TABLE IF NOT EXISTS `s1interface` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `S1InterfaceId` float DEFAULT NULL,
  `S1CpBearerId` float DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `MmeRelease` varchar(255) DEFAULT NULL,
  `S1InterfaceIsBlock` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `AutoCfgFlag` varchar(255) DEFAULT NULL,
  `Priority` float DEFAULT NULL,
  `CnOpSharingGroupId` float DEFAULT NULL,
  `ServedGummeis` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.s1reesttimer
CREATE TABLE IF NOT EXISTS `s1reesttimer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `S1ReestMinTimer` float DEFAULT NULL,
  `S1ReestMaxTimer` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.saallnk
CREATE TABLE IF NOT EXISTS `saallnk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SAALNO` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `JNRSCGRP` varchar(255) DEFAULT NULL,
  `VPI` float DEFAULT NULL,
  `VCI` float DEFAULT NULL,
  `ST` varchar(255) DEFAULT NULL,
  `PCR` float DEFAULT NULL,
  `CCTM` float DEFAULT NULL,
  `POLL` float DEFAULT NULL,
  `IDLE` float DEFAULT NULL,
  `NRTM` float DEFAULT NULL,
  `KATM` float DEFAULT NULL,
  `MAXCC` float DEFAULT NULL,
  `MAXPD` float DEFAULT NULL,
  `MAXLE` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `RU` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.scappparacfg
CREATE TABLE IF NOT EXISTS `scappparacfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AppDnsId` float DEFAULT NULL,
  `MatchRule` varchar(255) DEFAULT NULL,
  `AppName` varchar(255) DEFAULT NULL,
  `AppIdentType` varchar(255) DEFAULT NULL,
  `AppIpv4` varchar(255) DEFAULT NULL,
  `AppCfgTargetInd` varchar(255) DEFAULT NULL,
  `MainAppDnsFlag` varchar(255) DEFAULT NULL,
  `AsParaGroupID` float DEFAULT NULL,
  `MainAppDnsId` float DEFAULT NULL,
  `ServiceEndTimeThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sccfreqcfg
CREATE TABLE IF NOT EXISTS `sccfreqcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PccDlEarfcn` float DEFAULT NULL,
  `SccDlEarfcn` float DEFAULT NULL,
  `SccPriority` float DEFAULT NULL,
  `SccA2Offset` float DEFAULT NULL,
  `SccA4Offset` float DEFAULT NULL,
  `BlindScellAddThd` float DEFAULT NULL,
  `BlindScellDelThd` float DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `SpidGrpId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.scpolicy
CREATE TABLE IF NOT EXISTS `scpolicy` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ScAlgoSwitch` varchar(255) DEFAULT NULL,
  `VideoInitBufTime` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.scserviceqos
CREATE TABLE IF NOT EXISTS `scserviceqos` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ScQosId` float DEFAULT NULL,
  `AppDns` varchar(255) DEFAULT NULL,
  `DlSgbr` float DEFAULT NULL,
  `AppIdentType` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sctphost
CREATE TABLE IF NOT EXISTS `sctphost` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SCTPHOSTID` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `IPVERSION` varchar(255) DEFAULT NULL,
  `SIGIP1V4` varchar(255) DEFAULT NULL,
  `SIGIP1SECSWITCH` varchar(255) DEFAULT NULL,
  `SIGIP2V4` varchar(255) DEFAULT NULL,
  `SIGIP2SECSWITCH` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `SCTPTEMPLATEID` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `SIMPLEMODESWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sctplnk
CREATE TABLE IF NOT EXISTS `sctplnk` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SCTPNO` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MAXSTREAM` float DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL,
  `LOCIP` varchar(255) DEFAULT NULL,
  `SECLOCIP` varchar(255) DEFAULT NULL,
  `LOCPORT` float DEFAULT NULL,
  `PEERIP` varchar(255) DEFAULT NULL,
  `SECPEERIP` varchar(255) DEFAULT NULL,
  `PEERPORT` float DEFAULT NULL,
  `RTOMIN` float DEFAULT NULL,
  `RTOMAX` float DEFAULT NULL,
  `RTOINIT` float DEFAULT NULL,
  `RTOALPHA` float DEFAULT NULL,
  `RTOBETA` float DEFAULT NULL,
  `HBINTER` float DEFAULT NULL,
  `MAXASSOCRETR` float DEFAULT NULL,
  `MAXPATHRETR` float DEFAULT NULL,
  `CHKSUMTYPE` varchar(255) DEFAULT NULL,
  `AUTOSWITCH` varchar(255) DEFAULT NULL,
  `SWITCHBACKHBNUM` float DEFAULT NULL,
  `BLKFLAG` varchar(255) DEFAULT NULL,
  `TSACK` float DEFAULT NULL,
  `DESCRI` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sctppeer
CREATE TABLE IF NOT EXISTS `sctppeer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SCTPPEERID` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `IPVERSION` varchar(255) DEFAULT NULL,
  `SIGIP1V4` varchar(255) DEFAULT NULL,
  `SIGIP1SECSWITCH` varchar(255) DEFAULT NULL,
  `SIGIP2V4` varchar(255) DEFAULT NULL,
  `SIGIP2SECSWITCH` varchar(255) DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `REMOTEID` varchar(255) DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `BLKFLAG` varchar(255) DEFAULT NULL,
  `SIMPLEMODESWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sctptemplate
CREATE TABLE IF NOT EXISTS `sctptemplate` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SCTPTEMPLATEID` float DEFAULT NULL,
  `RTOMIN` float DEFAULT NULL,
  `RTOMAX` float DEFAULT NULL,
  `RTOINIT` float DEFAULT NULL,
  `RTOALPHA` float DEFAULT NULL,
  `RTOBETA` float DEFAULT NULL,
  `HBINTER` float DEFAULT NULL,
  `MAXASSOCRETR` float DEFAULT NULL,
  `MAXPATHRETR` float DEFAULT NULL,
  `SWITCHBACKFLAG` varchar(255) DEFAULT NULL,
  `SWITCHBACHHBNUM` float DEFAULT NULL,
  `TSACK` float DEFAULT NULL,
  `CHKSUMTYPE` varchar(255) DEFAULT NULL,
  `MAXSTREAM` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sector
CREATE TABLE IF NOT EXISTS `sector` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SECTORID` float DEFAULT NULL,
  `SECNAME` varchar(255) DEFAULT NULL,
  `LOCATIONNAME` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `ANTAZIMUTH` float DEFAULT NULL,
  `OLDSECTORID` float DEFAULT NULL,
  `SECTORIDFORCONVERSION` float DEFAULT NULL,
  `SECTORTYPEUMTS` varchar(255) DEFAULT NULL,
  `RXANTNUM` float DEFAULT NULL,
  `DIVMODE` varchar(255) DEFAULT NULL,
  `COVERTYPE` varchar(255) DEFAULT NULL,
  `RFCONNMODE` varchar(255) DEFAULT NULL,
  `SECTORMODELTE` varchar(255) DEFAULT NULL,
  `ANTENNAMODE` varchar(255) DEFAULT NULL,
  `SECTORCOMBIND` varchar(255) DEFAULT NULL,
  `OMNIFLAG` varchar(255) DEFAULT NULL,
  `ORIENTOFMAJORAXIS` float DEFAULT NULL,
  `CONFIDENCE` float DEFAULT NULL,
  `UNCERTSEMIMAJOR` float DEFAULT NULL,
  `UNCERTSEMIMINOR` float DEFAULT NULL,
  `UNCERTALTITUDE` float DEFAULT NULL,
  `SECTORANTENNA` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sectoreqm
CREATE TABLE IF NOT EXISTS `sectoreqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SECTOREQMID` float DEFAULT NULL,
  `SECTORID` float DEFAULT NULL,
  `SECTOREQMANTENNA` varchar(255) DEFAULT NULL,
  `ANTCFGMODE` varchar(255) DEFAULT NULL,
  `RRUCN` float DEFAULT NULL,
  `RRUSRN` float DEFAULT NULL,
  `RRUSN` float DEFAULT NULL,
  `BEAMSHAPE` varchar(255) DEFAULT NULL,
  `BEAMLAYERSPLIT` varchar(255) DEFAULT NULL,
  `BEAMAZIMUTHOFFSET` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.servicediffsetting
CREATE TABLE IF NOT EXISTS `servicediffsetting` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `QueueWeight0` float DEFAULT NULL,
  `QueueWeight1` float DEFAULT NULL,
  `QueueWeight2` float DEFAULT NULL,
  `QueueWeight3` float DEFAULT NULL,
  `QueueWeight4` float DEFAULT NULL,
  `QueueWeight5` float DEFAULT NULL,
  `QueueWeight6` float DEFAULT NULL,
  `QueueWeight7` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.serviceidentifypara
CREATE TABLE IF NOT EXISTS `serviceidentifypara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `HeartbeatPacketLengthThld` float DEFAULT NULL,
  `HeartbeatPacketNumberThld` float DEFAULT NULL,
  `MassFlowBigPacketRateThld` float DEFAULT NULL,
  `MassFlowDuration` float DEFAULT NULL,
  `MassFlowPacketNumberThld` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.serviceifdlearfcngrp
CREATE TABLE IF NOT EXISTS `serviceifdlearfcngrp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `ServiceIfHoCfgGroupId` float DEFAULT NULL,
  `DlEarfcnIndex` float DEFAULT NULL,
  `DlEarfcn` float DEFAULT NULL,
  `ServiceHoFreqPriority` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.serviceifhocfggroup
CREATE TABLE IF NOT EXISTS `serviceifhocfggroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `ServiceIfHoCfgGroupId` float DEFAULT NULL,
  `InterFreqHoState` varchar(255) DEFAULT NULL,
  `A4RptWaitingTimer` float DEFAULT NULL,
  `A4TimeToTriger` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.serviceirhocfggroup
CREATE TABLE IF NOT EXISTS `serviceirhocfggroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `ServiceIrHoCfgGroupId` float DEFAULT NULL,
  `InterRatHoState` varchar(255) DEFAULT NULL,
  `ServiceIrMeasMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.sfp
CREATE TABLE IF NOT EXISTS `sfp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MODULEID` float DEFAULT NULL,
  `PT` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.simuload
CREATE TABLE IF NOT EXISTS `simuload` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SimuLoadCfgIndex` float DEFAULT NULL,
  `SimuLoadRbThd` float DEFAULT NULL,
  `SimuLoadPwrThd` float DEFAULT NULL,
  `SimuLoadReportPeriod` float DEFAULT NULL,
  `SimuLoadStatPeriod` float DEFAULT NULL,
  `FreqSelSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.singleipswitch
CREATE TABLE IF NOT EXISTS `singleipswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SINGLEIPSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.srbcfg
CREATE TABLE IF NOT EXISTS `srbcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SrbRlcParaAdaptSwitch` varchar(255) DEFAULT NULL,
  `SrbPollTimerAdjustStep` float DEFAULT NULL,
  `SrbPollTimerMaxAdjustValue` varchar(255) DEFAULT NULL,
  `SrbPollTimerAdjUserNumThd` float DEFAULT NULL,
  `RrcConnRelMaxRetxThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `SrbPollTimerPreset` varchar(255) DEFAULT NULL,
  `HOCmdRRCRlsMsgRetxAdpSw` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.srbrlcpdcpcfg
CREATE TABLE IF NOT EXISTS `srbrlcpdcpcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SrbId` varchar(255) DEFAULT NULL,
  `UeMaxRetxThreshold` varchar(255) DEFAULT NULL,
  `ENodeBMaxRetxThreshold` varchar(255) DEFAULT NULL,
  `UePollByte` varchar(255) DEFAULT NULL,
  `ENodeBPollByte` varchar(255) DEFAULT NULL,
  `UePollPdu` varchar(255) DEFAULT NULL,
  `ENodeBPollPdu` varchar(255) DEFAULT NULL,
  `UePollRetransmitTimer` varchar(255) DEFAULT NULL,
  `ENodeBPollRetransmitTimer` varchar(255) DEFAULT NULL,
  `UeStatusProhibitTimer` varchar(255) DEFAULT NULL,
  `ENodeBStatusProhibitTimer` varchar(255) DEFAULT NULL,
  `UeAmReorderingTimer` varchar(255) DEFAULT NULL,
  `ENodeBAmReorderingTimer` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `NbENodeBPollRetransTimer` varchar(255) DEFAULT NULL,
  `NbUePollRetransTimer` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.srciprt
CREATE TABLE IF NOT EXISTS `srciprt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SRCRTIDX` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SBT` varchar(255) DEFAULT NULL,
  `SRCIP` varchar(255) DEFAULT NULL,
  `RTTYPE` varchar(255) DEFAULT NULL,
  `NEXTHOP` varchar(255) DEFAULT NULL,
  `PREF` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.srsadaptivecfg
CREATE TABLE IF NOT EXISTS `srsadaptivecfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SrsPeriodAdaptive` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.srscfg
CREATE TABLE IF NOT EXISTS `srscfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `SrsSubframeCfg` varchar(255) DEFAULT NULL,
  `AnSrsSimuTrans` varchar(255) DEFAULT NULL,
  `SrsCfgInd` varchar(255) DEFAULT NULL,
  `TddSrsCfgMode` varchar(255) DEFAULT NULL,
  `FddSrsCfgMode` varchar(255) DEFAULT NULL,
  `SrsAlgoOptSwitch` varchar(255) DEFAULT NULL,
  `SrsCfgPolicySwitch` varchar(255) DEFAULT NULL,
  `SrsResOptSwitch` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ssl
CREATE TABLE IF NOT EXISTS `ssl` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CONNTYPE` varchar(255) DEFAULT NULL,
  `AUTHMODE` varchar(255) DEFAULT NULL,
  `RENEGO` varchar(255) DEFAULT NULL,
  `RENEGOINTERVAL` float DEFAULT NULL,
  `LOWESTCSLEVEL` varchar(255) DEFAULT NULL,
  `VERSION` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.standardqci
CREATE TABLE IF NOT EXISTS `standardqci` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Qci` varchar(255) DEFAULT NULL,
  `PreAllocationWeight` float DEFAULT NULL,
  `InterRatPolicyCfgGroupId` float DEFAULT NULL,
  `RlcPdcpParaGroupId` float DEFAULT NULL,
  `LogicalChannelPriority` float DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `PrioritisedBitRate` varchar(255) DEFAULT NULL,
  `FlowCtrlType` varchar(255) DEFAULT NULL,
  `UlschPriorityFactor` float DEFAULT NULL,
  `DlschPriorityFactor` float DEFAULT NULL,
  `UlMinGbr` varchar(255) DEFAULT NULL,
  `DlMinGbr` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.subrack
CREATE TABLE IF NOT EXISTS `subrack` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL,
  `DESC` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.synceth
CREATE TABLE IF NOT EXISTS `synceth` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LN` float DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `QL` varchar(255) DEFAULT NULL,
  `SSM` varchar(255) DEFAULT NULL,
  `PRI` float DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tacalg
CREATE TABLE IF NOT EXISTS `tacalg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RSCGRPULCACSWITCH` varchar(255) DEFAULT NULL,
  `RSCGRPDLCACSWITCH` varchar(255) DEFAULT NULL,
  `TRMULHOCACTH` float DEFAULT NULL,
  `TRMDLHOCACTH` float DEFAULT NULL,
  `TRMULGOLDCACTH` float DEFAULT NULL,
  `TRMDLGOLDCACTH` float DEFAULT NULL,
  `TRMULSILVERCACTH` float DEFAULT NULL,
  `TRMDLSILVERCACTH` float DEFAULT NULL,
  `TRMULBRONZECACTH` float DEFAULT NULL,
  `TRMDLBRONZECACTH` float DEFAULT NULL,
  `TRMULGBRCACTH` float DEFAULT NULL,
  `TRMDLGBRCACTH` float DEFAULT NULL,
  `TRMULPRESW` varchar(255) DEFAULT NULL,
  `TRMDLPRESW` varchar(255) DEFAULT NULL,
  `PORTULOBSW` varchar(255) DEFAULT NULL,
  `PORTDLOBSW` varchar(255) DEFAULT NULL,
  `PORTULCACSW` varchar(255) DEFAULT NULL,
  `PORTDLCACSW` varchar(255) DEFAULT NULL,
  `EMCTACPSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tasm
CREATE TABLE IF NOT EXISTS `tasm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `SRCNO` varchar(255) DEFAULT NULL,
  `CLKSRC` varchar(255) DEFAULT NULL,
  `MODE` varchar(255) DEFAULT NULL,
  `ALGORITHM` varchar(255) DEFAULT NULL,
  `FREECOUNT` float DEFAULT NULL,
  `SEARCHCOUNT` float DEFAULT NULL,
  `HOLDCOUNT` float DEFAULT NULL,
  `LOCKCOUNT` float DEFAULT NULL,
  `SAMPLETIME` float DEFAULT NULL,
  `SYNMODE` varchar(255) DEFAULT NULL,
  `PERIOD` float DEFAULT NULL,
  `TM` varchar(255) DEFAULT NULL,
  `CLKSYNCMODE` varchar(255) DEFAULT NULL,
  `QL` varchar(255) DEFAULT NULL,
  `FREQUENCE` varchar(255) DEFAULT NULL,
  `NETMODE` varchar(255) DEFAULT NULL,
  `FLAG` varchar(255) DEFAULT NULL,
  `DAY` float DEFAULT NULL,
  `GSM_FBOFFSET` float DEFAULT NULL,
  `SYSCLKSRC` varchar(255) DEFAULT NULL,
  `CLOUDSRC` varchar(255) DEFAULT NULL,
  `FRAMESYNCSW` varchar(255) DEFAULT NULL,
  `GSMFRAMESYNCSW` varchar(255) DEFAULT NULL,
  `STANDARD` varchar(255) DEFAULT NULL,
  `ENSYSCLKCHKSW` varchar(255) DEFAULT NULL,
  `ATRSW` varchar(255) DEFAULT NULL,
  `FNSYNCSW` varchar(255) DEFAULT NULL,
  `DATE` varchar(255) DEFAULT NULL,
  `TIME` varchar(255) DEFAULT NULL,
  `LEAPSECONDSCHGDATE` varchar(255) DEFAULT NULL,
  `LEAPSECONDSCHGTIME` varchar(255) DEFAULT NULL,
  `CRTGPSTOUTCLEAPSECONDS` float DEFAULT NULL,
  `NEXTGPSTOUTCLEAPSECONDS` float DEFAULT NULL,
  `LPFNSYNCSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tbdspinfo
CREATE TABLE IF NOT EXISTS `tbdspinfo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LANNO` varchar(255) DEFAULT NULL,
  `VERSION` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tblangno
CREATE TABLE IF NOT EXISTS `tblangno` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LANNO` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tcpackctrlalgo
CREATE TABLE IF NOT EXISTS `tcpackctrlalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `AckCtrlSwitch` varchar(255) DEFAULT NULL,
  `DlMaxThroughput` float DEFAULT NULL,
  `CtrlTimerLength` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tcpacklimitalg
CREATE TABLE IF NOT EXISTS `tcpacklimitalg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TCPACKLIMITSWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tcpmssctrl
CREATE TABLE IF NOT EXISTS `tcpmssctrl` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TcpMssCtrlSwitch` varchar(255) DEFAULT NULL,
  `TcpMssThd` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tcu
CREATE TABLE IF NOT EXISTS `tcu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `MCN` float DEFAULT NULL,
  `MSRN` float DEFAULT NULL,
  `MPN` float DEFAULT NULL,
  `ADDR` float DEFAULT NULL,
  `TLTHD` float DEFAULT NULL,
  `TUTHD` float DEFAULT NULL,
  `SBAF` varchar(255) DEFAULT NULL,
  `TCMODE` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tddframeoffset
CREATE TABLE IF NOT EXISTS `tddframeoffset` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TimeOffset` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tddresmodeswitch
CREATE TABLE IF NOT EXISTS `tddresmodeswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ClkResExcludeSwitch` varchar(255) DEFAULT NULL,
  `BbResExclusiveSwitch` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.timealignmenttimer
CREATE TABLE IF NOT EXISTS `timealignmenttimer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TimeAlignmentTimer` varchar(255) DEFAULT NULL,
  `TimingAdvCmdOptSwitch` varchar(255) DEFAULT NULL,
  `TimingMeasMode` varchar(255) DEFAULT NULL,
  `TACmdSendPeriod` varchar(255) DEFAULT NULL,
  `TimingResOptSwitch` varchar(255) DEFAULT NULL,
  `PucchTimeOffsetCompVal` float DEFAULT NULL,
  `HighSpeedTaCmdSendPeriod` varchar(255) DEFAULT NULL,
  `TaEnhance` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.timesrc
CREATE TABLE IF NOT EXISTS `timesrc` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TIMESRC` varchar(255) DEFAULT NULL,
  `AUTOSWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.timethrd
CREATE TABLE IF NOT EXISTS `timethrd` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `THRD` float DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tldralg
CREATE TABLE IF NOT EXISTS `tldralg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TRMULLDRTRGTH` float DEFAULT NULL,
  `TRMDLLDRTRGTH` float DEFAULT NULL,
  `TRMULLDRCLRTH` float DEFAULT NULL,
  `TRMDLLDRCLRTH` float DEFAULT NULL,
  `TRMULMLDTRGTH` float DEFAULT NULL,
  `TRMDLMLDTRGTH` float DEFAULT NULL,
  `TRMULMLDCLRTH` float DEFAULT NULL,
  `TRMDLMLDCLRTH` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tlfrswitch
CREATE TABLE IF NOT EXISTS `tlfrswitch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TLFRSWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tolcalg
CREATE TABLE IF NOT EXISTS `tolcalg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TRMULOLCSWITCH` varchar(255) DEFAULT NULL,
  `TRMDLOLCSWITCH` varchar(255) DEFAULT NULL,
  `TRMULOLCTRIGTH` float DEFAULT NULL,
  `TRMULOLCRELTH` float DEFAULT NULL,
  `TRMOLCRELBEARERNUM` float DEFAULT NULL,
  `TRMDLOLCTRIGTH` float DEFAULT NULL,
  `TRMDLOLCRELTH` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tpealgo
CREATE TABLE IF NOT EXISTS `tpealgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `PortListNum` float DEFAULT NULL,
  `Port1` float DEFAULT NULL,
  `Port2` float DEFAULT NULL,
  `Port3` float DEFAULT NULL,
  `Port4` float DEFAULT NULL,
  `Port5` float DEFAULT NULL,
  `Port6` float DEFAULT NULL,
  `Port7` float DEFAULT NULL,
  `Port8` float DEFAULT NULL,
  `Port9` float DEFAULT NULL,
  `Port10` float DEFAULT NULL,
  `Port11` float DEFAULT NULL,
  `Port12` float DEFAULT NULL,
  `Port13` float DEFAULT NULL,
  `Port14` float DEFAULT NULL,
  `Port15` float DEFAULT NULL,
  `Port16` float DEFAULT NULL,
  `Port17` float DEFAULT NULL,
  `Port18` float DEFAULT NULL,
  `Port19` float DEFAULT NULL,
  `Port20` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.trp
CREATE TABLE IF NOT EXISTS `trp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `SUBTYPE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.trustcert
CREATE TABLE IF NOT EXISTS `trustcert` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CERTNAME` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.twampresponder
CREATE TABLE IF NOT EXISTS `twampresponder` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `RESPONDERID` float DEFAULT NULL,
  `LOCALIP` varchar(255) DEFAULT NULL,
  `VRFINDEX` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL,
  `REFWAIT` float DEFAULT NULL,
  `SERVWAIT` float DEFAULT NULL,
  `LOCALPORT` float DEFAULT NULL,
  `TWAMPARCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.txbranch
CREATE TABLE IF NOT EXISTS `txbranch` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `TXNO` float DEFAULT NULL,
  `TXSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.typdrbbsr
CREATE TABLE IF NOT EXISTS `typdrbbsr` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Qci` varchar(255) DEFAULT NULL,
  `TPerodicBSRTimer` varchar(255) DEFAULT NULL,
  `RetxBsrTimer` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.tz
CREATE TABLE IF NOT EXISTS `tz` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ZONET` varchar(255) DEFAULT NULL,
  `DST` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.udpping
CREATE TABLE IF NOT EXISTS `udpping` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `TIMEOUTTHD` float DEFAULT NULL,
  `TIMEOUTCNT` float DEFAULT NULL,
  `DSCP` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.udt
CREATE TABLE IF NOT EXISTS `udt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UDTNO` float DEFAULT NULL,
  `UDTPARAGRPID` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.udtparagrp
CREATE TABLE IF NOT EXISTS `udtparagrp` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UDTPARAGRPID` float DEFAULT NULL,
  `PRIRULE` varchar(255) DEFAULT NULL,
  `PRI` float DEFAULT NULL,
  `ACTFACTOR` float DEFAULT NULL,
  `PRIMTRANRSCTYPE` varchar(255) DEFAULT NULL,
  `PRIMPTLOADTH` float DEFAULT NULL,
  `PRIM2SECPTLOADRATH` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.uecooperationpara
CREATE TABLE IF NOT EXISTS `uecooperationpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `MacCeMsgLCID` float DEFAULT NULL,
  `A2Offset` varchar(255) DEFAULT NULL,
  `A3Offset` varchar(255) DEFAULT NULL,
  `UEAwarePowerSavingSwitch` varchar(255) DEFAULT NULL,
  `UEForbiddenMsgThd` float DEFAULT NULL,
  `SpecUserCooperationSwitch` varchar(255) DEFAULT NULL,
  `DsdsCooperationRptSwitch` varchar(255) DEFAULT NULL,
  `DsdsLcid` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ueiu
CREATE TABLE IF NOT EXISTS `ueiu` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.uespecdrxparagroup
CREATE TABLE IF NOT EXISTS `uespecdrxparagroup` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DrxParaGroupIndex` float DEFAULT NULL,
  `DrxParaGroupStat` varchar(255) DEFAULT NULL,
  `OnDurationTimer` varchar(255) DEFAULT NULL,
  `DrxInactivityTimer` varchar(255) DEFAULT NULL,
  `DrxReTxTimer` varchar(255) DEFAULT NULL,
  `LongDrxCycle` varchar(255) DEFAULT NULL,
  `SupportShortDrx` varchar(255) DEFAULT NULL,
  `ShortDrxCycle` varchar(255) DEFAULT NULL,
  `DrxShortCycleTimer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.uetimerconst
CREATE TABLE IF NOT EXISTS `uetimerconst` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `T300` varchar(255) DEFAULT NULL,
  `T301` varchar(255) DEFAULT NULL,
  `T310` varchar(255) DEFAULT NULL,
  `T311` varchar(255) DEFAULT NULL,
  `N311` varchar(255) DEFAULT NULL,
  `N310` varchar(255) DEFAULT NULL,
  `T325` varchar(255) DEFAULT NULL,
  `N313` varchar(255) DEFAULT NULL,
  `N314` varchar(255) DEFAULT NULL,
  `T307` varchar(255) DEFAULT NULL,
  `T313` varchar(255) DEFAULT NULL,
  `T300ForNb` varchar(255) DEFAULT NULL,
  `T310ForNb` varchar(255) DEFAULT NULL,
  `T301ForNb` varchar(255) DEFAULT NULL,
  `T311ForNb` varchar(255) DEFAULT NULL,
  `T300CE` varchar(255) DEFAULT NULL,
  `T301CE` varchar(255) DEFAULT NULL,
  `T310CE` varchar(255) DEFAULT NULL,
  `T311CE` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulcsalgopara
CREATE TABLE IF NOT EXISTS `ulcsalgopara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlCoPcInThd` float DEFAULT NULL,
  `UlCoPcRbkRsrpThd` float DEFAULT NULL,
  `UlCoPcUserNumThd` float DEFAULT NULL,
  `UlCoResAllocBandMode` varchar(255) DEFAULT NULL,
  `UlCoResAllocRbLoadThld` float DEFAULT NULL,
  `UlCsA3Offset` float DEFAULT NULL,
  `UlCsSw` varchar(255) DEFAULT NULL,
  `UlCraInThld` float DEFAULT NULL,
  `UlCraUserNumThld` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulinterfsuppresscfg
CREATE TABLE IF NOT EXISTS `ulinterfsuppresscfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UlInterfsuppressionThd` float DEFAULT NULL,
  `P0UePuschOffset` float DEFAULT NULL,
  `RASignalMCSOffset` float DEFAULT NULL,
  `P0UePucchOffset` float DEFAULT NULL,
  `DeltaMSG2Offset` varchar(255) DEFAULT NULL,
  `StrongInfUserThdRatio` float DEFAULT NULL,
  `ULInfStatisticPeriod` float DEFAULT NULL,
  `ULInfStatisticPeriodNum` float DEFAULT NULL,
  `VolteMCSOffset` float DEFAULT NULL,
  `VoltePucchPcTarSinrOffset` float DEFAULT NULL,
  `VoltePuschPsdCtrlTarget` float DEFAULT NULL,
  `RemoteInfULEnhanceSw` varchar(255) DEFAULT NULL,
  `PuschEnhDeltaSinrThd` float DEFAULT NULL,
  `RASignalPcSwitch` varchar(255) DEFAULT NULL,
  `UlInterPwrDiffThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocell
CREATE TABLE IF NOT EXISTS `ulocell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `VIRTUALCPC` varchar(255) DEFAULT NULL,
  `RADIUS` float DEFAULT NULL,
  `BOOST` varchar(255) DEFAULT NULL,
  `HORAD` float DEFAULT NULL,
  `TIMELIMIT` varchar(255) DEFAULT NULL,
  `ULCOVEXPANSION` varchar(255) DEFAULT NULL,
  `ULBASEBANDEQMID` float DEFAULT NULL,
  `PRECISECHEST2MS` varchar(255) DEFAULT NULL,
  `DLBASEBANDEQMID` float DEFAULT NULL,
  `INTERNODEBULCOMP` varchar(255) DEFAULT NULL,
  `ULFREQIND` varchar(255) DEFAULT NULL,
  `SPEEDBASEDDEMODSW` varchar(255) DEFAULT NULL,
  `OBJID` float DEFAULT NULL,
  `NBIS` varchar(255) DEFAULT NULL,
  `FREQUENCYBANDWIDTH` varchar(255) DEFAULT NULL,
  `HSDPA4C` varchar(255) DEFAULT NULL,
  `DL64QAM` varchar(255) DEFAULT NULL,
  `BBUNITREF` varchar(255) DEFAULT NULL,
  `MULTIRRUSTATICDESSW` varchar(255) DEFAULT NULL,
  `AIR` varchar(255) DEFAULT NULL,
  `TTW` varchar(255) DEFAULT NULL,
  `LOCELLPRI` float DEFAULT NULL,
  `LOCELLTYPE` varchar(255) DEFAULT NULL,
  `TURBOIC` varchar(255) DEFAULT NULL,
  `ERACHHSDPCCH` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `ULFREQ` float DEFAULT NULL,
  `DLFREQ` float DEFAULT NULL,
  `MAXPWR` float DEFAULT NULL,
  `DEFPWRLVL` float DEFAULT NULL,
  `DLRESMODE` varchar(255) DEFAULT NULL,
  `DI` float DEFAULT NULL,
  `HISPM` varchar(255) DEFAULT NULL,
  `RMTCM` varchar(255) DEFAULT NULL,
  `UL16QAM` varchar(255) DEFAULT NULL,
  `FDEMODE` varchar(255) DEFAULT NULL,
  `ULL2PLUS` varchar(255) DEFAULT NULL,
  `ICMODE` varchar(255) DEFAULT NULL,
  `ERACH` varchar(255) DEFAULT NULL,
  `RSV` varchar(255) DEFAULT NULL,
  `GUPOWERSHARE` varchar(255) DEFAULT NULL,
  `HSPAUSERNUMEXT` varchar(255) DEFAULT NULL,
  `IRCSW` varchar(255) DEFAULT NULL,
  `TURBOICPHASE2` varchar(255) DEFAULT NULL,
  `MFHSDPASW` varchar(255) DEFAULT NULL,
  `ULCOMPSW` varchar(255) DEFAULT NULL,
  `VAM` varchar(255) DEFAULT NULL,
  `CELLSCALEIND` varchar(255) DEFAULT NULL,
  `INTERNODEBHSDPCCHCOMP` varchar(255) DEFAULT NULL,
  `ULSUPERNARROWBANDFILTER` varchar(255) DEFAULT NULL,
  `DLADAPTIVEBPFILTER` varchar(255) DEFAULT NULL,
  `GSMQLBASEDROTDYNACTRL` varchar(255) DEFAULT NULL,
  `INTEL2TVAM` varchar(255) DEFAULT NULL,
  `DLASYMFILTER` varchar(255) DEFAULT NULL,
  `DLASYMFILTERNFCELLID` varchar(255) DEFAULT NULL,
  `DLASYMLEFTBANDWIDTH` varchar(255) DEFAULT NULL,
  `DLASYMRIGHTBANDWIDTH` varchar(255) DEFAULT NULL,
  `NBISADAPTIVEFORBIDSW` varchar(255) DEFAULT NULL,
  `PRECDWSRSW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellalgpara
CREATE TABLE IF NOT EXISTS `ulocellalgpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `RTWPSIRTGTADJSW` varchar(255) DEFAULT NULL,
  `SIB7RTWPOPTSW` varchar(255) DEFAULT NULL,
  `ANTIANTENNAIMBALANCESW` varchar(255) DEFAULT NULL,
  `R99FLOWCTRLSW` varchar(255) DEFAULT NULL,
  `CQISAMPLEPERIOD` varchar(255) DEFAULT NULL,
  `CQIUSERNUM` float DEFAULT NULL,
  `HIGHSPEEDRAKESW` varchar(255) DEFAULT NULL,
  `DPCHMAXTXPWRRESTRSW` varchar(255) DEFAULT NULL,
  `DPCHMAXPWRRTRLOADSTAT` varchar(255) DEFAULT NULL,
  `SIB7MAXRTWP` float DEFAULT NULL,
  `ULCCHOLPCSW` varchar(255) DEFAULT NULL,
  `ULCCHOLPCLOADHIGHTHD` float DEFAULT NULL,
  `ULCCHOLPCLOADLOWTHD` float DEFAULT NULL,
  `RADIOQUALITYDPCHPCSW` varchar(255) DEFAULT NULL,
  `RADIOQUALITYDPCHPCLOADSTAT` varchar(255) DEFAULT NULL,
  `RADIOQUALITYDPCHPCPO` float DEFAULT NULL,
  `HTOHPWSHSW` varchar(255) DEFAULT NULL,
  `SHARINGMARGIN` float DEFAULT NULL,
  `MAXSHARINGRATIO` float DEFAULT NULL,
  `SRBOVERHSDPAOPTSW` varchar(255) DEFAULT NULL,
  `DLPWRLOADTHD` float DEFAULT NULL,
  `SINGLEHARQACTTHRD` float DEFAULT NULL,
  `R99TOHPWSHSW` varchar(255) DEFAULT NULL,
  `HTOHRTPWSHSW` varchar(255) DEFAULT NULL,
  `DPCHTPCPOADJSW` varchar(255) DEFAULT NULL,
  `HSDPACELLTHPTHD` float DEFAULT NULL,
  `HSDPAPWRMGNCANCELSW` varchar(255) DEFAULT NULL,
  `RADIOQUALITYDPCHPCSHOSW` varchar(255) DEFAULT NULL,
  `ULCARLEVELSCHSW` varchar(255) DEFAULT NULL,
  `RADIOQUALITYDPCHPCENHSW` varchar(255) DEFAULT NULL,
  `RADIOQUALITYFDPCHPCPO` float DEFAULT NULL,
  `COVERIMPBASEDONADVDEM` varchar(255) DEFAULT NULL,
  `SRBOVERHSDPAMMTOPTSW` varchar(255) DEFAULT NULL,
  `TURBOICENHANCEDSW` varchar(255) DEFAULT NULL,
  `AMRIMPBASEDONPLVASW` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellmacepara
CREATE TABLE IF NOT EXISTS `ulocellmacepara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `SCHEDULEPARA` varchar(255) DEFAULT NULL,
  `HSUPAOLSCHSW` varchar(255) DEFAULT NULL,
  `MAXDELTAOFTARGETROT` float DEFAULT NULL,
  `EAGCHPCPARA` varchar(255) DEFAULT NULL,
  `EAGCHPCMOD` varchar(255) DEFAULT NULL,
  `EAGCHPWROFFSET` float DEFAULT NULL,
  `EAGCHPOWER` float DEFAULT NULL,
  `MAXAGCHPOWER` float DEFAULT NULL,
  `MINAGCHPOWER` float DEFAULT NULL,
  `SERGCHPCPARA` varchar(255) DEFAULT NULL,
  `SERGCHPCMOD` varchar(255) DEFAULT NULL,
  `SERGCHPWROFFSET` float DEFAULT NULL,
  `SERGCHPOWER` float DEFAULT NULL,
  `NSERGCHPCPARA` varchar(255) DEFAULT NULL,
  `NSERGCHPCMOD` varchar(255) DEFAULT NULL,
  `NSERGCHPWROFFSET` float DEFAULT NULL,
  `NSERGCHPOWER` float DEFAULT NULL,
  `SEHICHPCPARA` varchar(255) DEFAULT NULL,
  `SEHICHPWROFFSET` float DEFAULT NULL,
  `SEHICHPOWER` float DEFAULT NULL,
  `SRLEHICHPWROFFSET` float DEFAULT NULL,
  `SINGLERLEHICHPOWER` float DEFAULT NULL,
  `NSEHICHPCPARA` varchar(255) DEFAULT NULL,
  `NSEHICHPWROFFSET` float DEFAULT NULL,
  `NSEHICHPOWER` float DEFAULT NULL,
  `OUTERSYSINTERSCHSW` varchar(255) DEFAULT NULL,
  `OWNCELLULLOADRATIO` float DEFAULT NULL,
  `RTWPMEAOPTIMSW` varchar(255) DEFAULT NULL,
  `LOADTHRESH4MINULCOV` float DEFAULT NULL,
  `HSUPATDSCHSW` varchar(255) DEFAULT NULL,
  `HSUPATDALIGNUENUM` float DEFAULT NULL,
  `HSUPATDSCHENSW` varchar(255) DEFAULT NULL,
  `RTWPSINGLEHARQSCHSW` varchar(255) DEFAULT NULL,
  `ULSECCELLACTOPTSW` varchar(255) DEFAULT NULL,
  `ULSECCELLDEACTCOVTHR` float DEFAULT NULL,
  `HSUPATOTALTHRGUARANTEESW` varchar(255) DEFAULT NULL,
  `HSUPATOTALTHRTHD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellmachspara
CREATE TABLE IF NOT EXISTS `ulocellmachspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `RSCALLOCM` varchar(255) DEFAULT NULL,
  `SM` varchar(255) DEFAULT NULL,
  `PWRMGN` float DEFAULT NULL,
  `HSSCCHPWRCMINDCH` varchar(255) DEFAULT NULL,
  `SCCHPWR` float DEFAULT NULL,
  `HSSCCHFERTRGTINDCH` float DEFAULT NULL,
  `RSCLMSW` varchar(255) DEFAULT NULL,
  `DYNCODESW` varchar(255) DEFAULT NULL,
  `CME16QAMSW` varchar(255) DEFAULT NULL,
  `MXPWRPHUSR` float DEFAULT NULL,
  `CQIFA` float DEFAULT NULL,
  `MAXEFACHHARQRT` float DEFAULT NULL,
  `MAXNONCONVERHARQRT` float DEFAULT NULL,
  `HSSCCHPWRCMINEFACH` varchar(255) DEFAULT NULL,
  `CQIADJALGOFCON` varchar(255) DEFAULT NULL,
  `CQIADJALGOFNONCON` varchar(255) DEFAULT NULL,
  `RBLERTARGET` float DEFAULT NULL,
  `IBLERTARGET` float DEFAULT NULL,
  `SECCELLACTDEASW` varchar(255) DEFAULT NULL,
  `IICSW` varchar(255) DEFAULT NULL,
  `DTXALGOSW` varchar(255) DEFAULT NULL,
  `MIMOPRIMESW` varchar(255) DEFAULT NULL,
  `LOCWEIGHT` varchar(255) DEFAULT NULL,
  `EXTRAPOWER` float DEFAULT NULL,
  `MAXEFACHHSHARQRT` float DEFAULT NULL,
  `BURSTBLEROPTSW` varchar(255) DEFAULT NULL,
  `DLBESCHWHENULURLBSW` varchar(255) DEFAULT NULL,
  `EXTRAPOWERFORSRB` float DEFAULT NULL,
  `CODEOPTSW` varchar(255) DEFAULT NULL,
  `PREALLOCBWOPTISW` varchar(255) DEFAULT NULL,
  `SFDCMACEHSDSCTMR` varchar(255) DEFAULT NULL,
  `DF3CMACEHSDSCTMR` varchar(255) DEFAULT NULL,
  `HSDPAL2OPTSW` varchar(255) DEFAULT NULL,
  `PWRMGNDELTAFORPT` float DEFAULT NULL,
  `VIDEOCQITHD` float DEFAULT NULL,
  `HSDPAFULLBUFDATATHD` float DEFAULT NULL,
  `HSDPAFULLBUFTHROUTHD` float DEFAULT NULL,
  `HSSCCHCQIPWRINDCHOPTSW` varchar(255) DEFAULT NULL,
  `MAXEXTRAPWRFORHSSCCHINDCH` float DEFAULT NULL,
  `SERVICECQITHD` float DEFAULT NULL,
  `EFACHDTCHRSCALLOCMODE` varchar(255) DEFAULT NULL,
  `CQIADJBYIBLERADJSW` varchar(255) DEFAULT NULL,
  `CQIADJBYDYNBLERADJSW` varchar(255) DEFAULT NULL,
  `ACTSECCELLTHPTHD` float DEFAULT NULL,
  `DEASECCELLTHPTHD` float DEFAULT NULL,
  `ACTSECCELLDATABUFTHD` float DEFAULT NULL,
  `DEASECCELLDATABUFTHD` float DEFAULT NULL,
  `INDEPTCARSCHSW` varchar(255) DEFAULT NULL,
  `BUFFERSMSW` varchar(255) DEFAULT NULL,
  `DYNCODEOPTSW` varchar(255) DEFAULT NULL,
  `HARQRTSCHOPTSW` varchar(255) DEFAULT NULL,
  `MCDELTACQITHD` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellnoaccesspara
CREATE TABLE IF NOT EXISTS `ulocellnoaccesspara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `NOUETIMER` float DEFAULT NULL,
  `NOFSTRLTIMER` float DEFAULT NULL,
  `AUTORCVRMTHD` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellr99algpara
CREATE TABLE IF NOT EXISTS `ulocellr99algpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `PWRCNGCACSW` varchar(255) DEFAULT NULL,
  `SHORTTH` float DEFAULT NULL,
  `LONGTH` float DEFAULT NULL,
  `SHORTTHFORRTRRC` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellrsclmtpara
CREATE TABLE IF NOT EXISTS `ulocellrsclmtpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `CME8KRSCLMT` float DEFAULT NULL,
  `CME16KRSCLMT` float DEFAULT NULL,
  `CME32KRSCLMT` float DEFAULT NULL,
  `CME64KRSCLMT` float DEFAULT NULL,
  `CME128KRSCLMT` float DEFAULT NULL,
  `CME256KRSCLMT` float DEFAULT NULL,
  `CME384KRSCLMT` float DEFAULT NULL,
  `CME512KRSCLMT` float DEFAULT NULL,
  `CME768KRSCLMT` float DEFAULT NULL,
  `CME1024KRSCLMT` float DEFAULT NULL,
  `CME1536KRSCLMT` float DEFAULT NULL,
  `CME1800KRSCLMT` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellrsvdpara
CREATE TABLE IF NOT EXISTS `ulocellrsvdpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `ULOCELLRSVDPARA1` longtext,
  `ULOCELLRSVDPARA2` longtext,
  `ULOCELLRSVDPARA3` float DEFAULT NULL,
  `ULOCELLRSVDPARA4` float DEFAULT NULL,
  `ULOCELLRSVDPARA5` float DEFAULT NULL,
  `ULOCELLRSVDPARA6` float DEFAULT NULL,
  `ULOCELLRSVDPARA7` float DEFAULT NULL,
  `ULOCELLRSVDPARA8` float DEFAULT NULL,
  `ULOCELLRSVDPARA9` float DEFAULT NULL,
  `ULOCELLRSVDPARA10` float DEFAULT NULL,
  `ULOCELLRSVDPARA11` float DEFAULT NULL,
  `ULOCELLRSVDPARA12` float DEFAULT NULL,
  `ULOCELLRSVDPARA13` longtext,
  `ULOCELLRSVDPARA14` longtext,
  `ULOCELLRSVDPARA15` float DEFAULT NULL,
  `ULOCELLRSVDPARA16` float DEFAULT NULL,
  `ULOCELLRSVDPARA17` float DEFAULT NULL,
  `ULOCELLRSVDPARA18` float DEFAULT NULL,
  `ULOCELLRSVDPARA19` float DEFAULT NULL,
  `ULOCELLRSVDPARA20` float DEFAULT NULL,
  `ULOCELLRSVDPARA21` float DEFAULT NULL,
  `ULOCELLRSVDPARA22` float DEFAULT NULL,
  `ULOCELLRSVDPARA23` float DEFAULT NULL,
  `ULOCELLRSVDPARA24` float DEFAULT NULL,
  `ULOCELLRSVDPARA25` float DEFAULT NULL,
  `ULOCELLRSVDPARA26` float DEFAULT NULL,
  `ULOCELLRSVDPARA27` float DEFAULT NULL,
  `ULOCELLRSVDPARA28` float DEFAULT NULL,
  `ULOCELLRSVDPARA29` float DEFAULT NULL,
  `ULOCELLRSVDPARA30` float DEFAULT NULL,
  `ULOCELLRSVDPARA31` float DEFAULT NULL,
  `ULOCELLRSVDPARA32` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulocellsectoreqm
CREATE TABLE IF NOT EXISTS `ulocellsectoreqm` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULOCELLID` float DEFAULT NULL,
  `SECTOREQMID` float DEFAULT NULL,
  `MAXPWR` float DEFAULT NULL,
  `SECTOREQMPROPERTY` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.ulzerobufferzone
CREATE TABLE IF NOT EXISTS `ulzerobufferzone` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `ULSharedFreqStartRB` float DEFAULT NULL,
  `ULSharedFreqEndRB` float DEFAULT NULL,
  `ULZeroBufZonePRBThd` float DEFAULT NULL,
  `ULZeroBufZonePRBOffset` float DEFAULT NULL,
  `ULZeroBufZoneRsrpThd` float DEFAULT NULL,
  `ULZeroBufZoneRsrpOffset` float DEFAULT NULL,
  `ULZeroBufZoneB1RmvOffset` float DEFAULT NULL,
  `ULZeroBufZoneUtranRscpThd` float DEFAULT NULL,
  `ULZeroBufZoInterRatB1Timer` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.upptsinterfcfg
CREATE TABLE IF NOT EXISTS `upptsinterfcfg` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TestPeriod` float DEFAULT NULL,
  `TestThreshold` float DEFAULT NULL,
  `TestHysteresis` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.usb
CREATE TABLE IF NOT EXISTS `usb` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `CN` float DEFAULT NULL,
  `SRN` float DEFAULT NULL,
  `SN` float DEFAULT NULL,
  `PN` float DEFAULT NULL,
  `SWITCH` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.userplanehost
CREATE TABLE IF NOT EXISTS `userplanehost` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UPHOSTID` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `IPVERSION` varchar(255) DEFAULT NULL,
  `LOCIPV4` varchar(255) DEFAULT NULL,
  `IPSECSWITCH` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.userplanepeer
CREATE TABLE IF NOT EXISTS `userplanepeer` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `UPPEERID` float DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `IPVERSION` varchar(255) DEFAULT NULL,
  `PEERIPV4` varchar(255) DEFAULT NULL,
  `IPSECSWITCH` varchar(255) DEFAULT NULL,
  `REMOTEID` varchar(255) DEFAULT NULL,
  `CTRLMODE` varchar(255) DEFAULT NULL,
  `AUTOCFGFLAG` varchar(255) DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL,
  `STATICCHK` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.userpriority
CREATE TABLE IF NOT EXISTS `userpriority` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ArpSchSwitch` varchar(255) DEFAULT NULL,
  `Arp1Priority` varchar(255) DEFAULT NULL,
  `Arp2Priority` varchar(255) DEFAULT NULL,
  `Arp3Priority` varchar(255) DEFAULT NULL,
  `Arp4Priority` varchar(255) DEFAULT NULL,
  `Arp5Priority` varchar(255) DEFAULT NULL,
  `Arp6Priority` varchar(255) DEFAULT NULL,
  `Arp7Priority` varchar(255) DEFAULT NULL,
  `Arp8Priority` varchar(255) DEFAULT NULL,
  `Arp9Priority` varchar(255) DEFAULT NULL,
  `Arp10Priority` varchar(255) DEFAULT NULL,
  `Arp11Priority` varchar(255) DEFAULT NULL,
  `Arp12Priority` varchar(255) DEFAULT NULL,
  `Arp13Priority` varchar(255) DEFAULT NULL,
  `Arp14Priority` varchar(255) DEFAULT NULL,
  `Arp15Priority` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.userqcipriority
CREATE TABLE IF NOT EXISTS `userqcipriority` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Qci` varchar(255) DEFAULT NULL,
  `GoldUlSchPriorityFactor` float DEFAULT NULL,
  `GoldDlSchPriorityFactor` float DEFAULT NULL,
  `SilverUlSchPriorityFactor` float DEFAULT NULL,
  `SilverDlSchPriorityFactor` float DEFAULT NULL,
  `BronzeUlSchPriorityFactor` float DEFAULT NULL,
  `BronzeDlSchPriorityFactor` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.uservpfmpara
CREATE TABLE IF NOT EXISTS `uservpfmpara` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `DlUserThruputThd0` float DEFAULT NULL,
  `DlUserThruputThd1` float DEFAULT NULL,
  `DlUserThruputThd2` float DEFAULT NULL,
  `DlUserThruputThd3` float DEFAULT NULL,
  `DlUserThruputThd4` float DEFAULT NULL,
  `UlAccessDelayGoodThd` float DEFAULT NULL,
  `UlAccessDelayBadThd` float DEFAULT NULL,
  `DlAccessDelayGoodThd` float DEFAULT NULL,
  `DlAccessDelayBadThd` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.utranexternalcell
CREATE TABLE IF NOT EXISTS `utranexternalcell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `RncId` float DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `UtranDlArfcn` float DEFAULT NULL,
  `UtranUlArfcnCfgInd` varchar(255) DEFAULT NULL,
  `UtranFddTddType` varchar(255) DEFAULT NULL,
  `RacCfgInd` varchar(255) DEFAULT NULL,
  `Rac` float DEFAULT NULL,
  `PScrambCode` float DEFAULT NULL,
  `Lac` float DEFAULT NULL,
  `CellName` varchar(255) DEFAULT NULL,
  `CsPsHOInd` varchar(255) DEFAULT NULL,
  `UtranExternalCellSlaveBand` varchar(255) DEFAULT NULL,
  `RoamingAreaHoInd` varchar(255) DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL,
  `UtranUlArfcn` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.utranncell
CREATE TABLE IF NOT EXISTS `utranncell` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `RncId` float DEFAULT NULL,
  `CellId` float DEFAULT NULL,
  `NoHoFlag` varchar(255) DEFAULT NULL,
  `NoRmvFlag` varchar(255) DEFAULT NULL,
  `AnrFlag` varchar(255) DEFAULT NULL,
  `BlindHoPriority` float DEFAULT NULL,
  `CellMeasPriority` varchar(255) DEFAULT NULL,
  `OverlapInd` varchar(255) DEFAULT NULL,
  `NCellMeasPriority` float DEFAULT NULL,
  `LocalCellName` varchar(255) DEFAULT NULL,
  `NeighbourCellName` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.utrannfreq
CREATE TABLE IF NOT EXISTS `utrannfreq` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UtranDlArfcn` float DEFAULT NULL,
  `UtranVersion` varchar(255) DEFAULT NULL,
  `UtranFddTddType` varchar(255) DEFAULT NULL,
  `UtranUlArfcnCfgInd` varchar(255) DEFAULT NULL,
  `CellReselPriorityCfgInd` varchar(255) DEFAULT NULL,
  `CellReselPriority` float DEFAULT NULL,
  `PmaxUtran` float DEFAULT NULL,
  `OffsetFreq` float DEFAULT NULL,
  `Qqualmin` float DEFAULT NULL,
  `QRxLevMin` float DEFAULT NULL,
  `ThreshXHigh` float DEFAULT NULL,
  `ThreshXLow` float DEFAULT NULL,
  `ThreshXHighQ` float DEFAULT NULL,
  `ThreshXLowQ` float DEFAULT NULL,
  `PsPriority` varchar(255) DEFAULT NULL,
  `CsPriority` varchar(255) DEFAULT NULL,
  `ConnFreqPriority` float DEFAULT NULL,
  `CsPsMixedPriority` varchar(255) DEFAULT NULL,
  `ContinuCoverageIndication` varchar(255) DEFAULT NULL,
  `MlbFreqPriority` float DEFAULT NULL,
  `MasterBandFlag` varchar(255) DEFAULT NULL,
  `UtranRanSharingInd` varchar(255) DEFAULT NULL,
  `AnrInd` varchar(255) DEFAULT NULL,
  `SrvccPriority` varchar(255) DEFAULT NULL,
  `ULSharedFreqInd` varchar(255) DEFAULT NULL,
  `FreqPriorityForAnr` float DEFAULT NULL,
  `MlbTargetInd` varchar(255) DEFAULT NULL,
  `UtranFreqHighSpeedFlag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.utranranshare
CREATE TABLE IF NOT EXISTS `utranranshare` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `UtranDlArfcn` float DEFAULT NULL,
  `Mcc` float DEFAULT NULL,
  `Mnc` float DEFAULT NULL,
  `CellReselPriorityCfgInd` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.vlanclass
CREATE TABLE IF NOT EXISTS `vlanclass` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `VLANGROUPNO` float DEFAULT NULL,
  `TRAFFIC` varchar(255) DEFAULT NULL,
  `SRVPRIO` float DEFAULT NULL,
  `VLANID` float DEFAULT NULL,
  `VLANPRIO` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.vlanmap
CREATE TABLE IF NOT EXISTS `vlanmap` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `NEXTHOPIP` varchar(255) DEFAULT NULL,
  `MASK` varchar(255) DEFAULT NULL,
  `VLANMODE` varchar(255) DEFAULT NULL,
  `VLANID` float DEFAULT NULL,
  `VLANPRIO` float DEFAULT NULL,
  `SETPRIO` varchar(255) DEFAULT NULL,
  `VLANGROUPNO` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.vqmalgo
CREATE TABLE IF NOT EXISTS `vqmalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `ULDelayJitter` float DEFAULT NULL,
  `VqiGoodThd` float DEFAULT NULL,
  `VqiPoorThd` float DEFAULT NULL,
  `VqiBadThd` float DEFAULT NULL,
  `AmrWbRadioLowQualRelThd` float DEFAULT NULL,
  `AmrNbRadioLowQualRelThd` float DEFAULT NULL,
  `AmrSilentPeriodNum` float DEFAULT NULL,
  `AmrLowQualRelPeriodNum` float DEFAULT NULL,
  `VqiExcellentThd` float DEFAULT NULL,
  `AmrNbSilentThd` float DEFAULT NULL,
  `AmrWbRcvLowQualRelThd` float DEFAULT NULL,
  `AmrNbRcvLowQualRelThd` float DEFAULT NULL,
  `AmrWbSilentThd` float DEFAULT NULL,
  `VQMAlgoPeriod` varchar(255) DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.vrf
CREATE TABLE IF NOT EXISTS `vrf` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `VRFIDX` float DEFAULT NULL,
  `USERLABEL` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.weblmt
CREATE TABLE IF NOT EXISTS `weblmt` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `POLICY` varchar(255) DEFAULT NULL,
  `SSLVER` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.wtcpproxyalgo
CREATE TABLE IF NOT EXISTS `wtcpproxyalgo` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `LocalCellId` float DEFAULT NULL,
  `TcpAccelerationSwitch` varchar(255) DEFAULT NULL,
  `TCPStatisticsSwitch` varchar(255) DEFAULT NULL,
  `MaxRttStatisticsThd` float DEFAULT NULL,
  `MaxProxyPktNum` float DEFAULT NULL,
  `AckSplitCount` float DEFAULT NULL,
  `TcpAsSchWeightFactor` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.x2
CREATE TABLE IF NOT EXISTS `x2` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `X2Id` float DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `CpEpGroupId` float DEFAULT NULL,
  `UpEpGroupId` float DEFAULT NULL,
  `TargetENodeBRelease` varchar(255) DEFAULT NULL,
  `UserLabel` varchar(255) DEFAULT NULL,
  `EpGroupCfgFlag` varchar(255) DEFAULT NULL,
  `CnOpSharingGroupId` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
-- Dumping structure for table cfgmml.x2interface
CREATE TABLE IF NOT EXISTS `x2interface` (
  `NE NAME` varchar(100) DEFAULT NULL,
  `X2InterfaceId` float DEFAULT NULL,
  `X2CpBearerId` float DEFAULT NULL,
  `CnOperatorId` float DEFAULT NULL,
  `TargetENodeBRelease` varchar(255) DEFAULT NULL,
  `CtrlMode` varchar(255) DEFAULT NULL,
  `AutoCfgFlag` varchar(255) DEFAULT NULL,
  `CnOpSharingGroupId` float DEFAULT NULL,
  `objId` float DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
