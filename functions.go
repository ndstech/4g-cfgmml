package main

import (
	"fmt"
	"github.com/yogawa/4g-cfgmml/model"
)

func insertAlmcurcfg(eNodeBId string, baseName string, data []model.Almcurcfg) {
	fmt.Println("[+] Processing Almcurcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `almcurcfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AID,
			data[j].ATTRIBUTES.ALVL,
			data[j].ATTRIBUTES.ASS,
			data[j].ATTRIBUTES.SHLDFLG,
			data[j].ATTRIBUTES.ANM,	

		)
		checkErr(err)
		//fmt.Println("[+] Almcurcfg data has been saved")
	}
	tx.Commit()
}


func insertAlmport(eNodeBId string, baseName string, data []model.Almport) {
	fmt.Println("[+] Processing Almport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `almport` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.SW,
			data[j].ATTRIBUTES.AID,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.AVOL,
			data[j].ATTRIBUTES.DTPRD,
			data[j].ATTRIBUTES.UL,
			data[j].ATTRIBUTES.LL,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.SMUL,
			data[j].ATTRIBUTES.SMLL,
			data[j].ATTRIBUTES.SOUL,
			data[j].ATTRIBUTES.SOLL,	

		)
		checkErr(err)
		//fmt.Println("[+] Almport data has been saved")
	}
	tx.Commit()
}


func insertAnr(eNodeBId string, baseName string, data []model.Anr) {
	fmt.Println("[+] Processing Anr data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `anr` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DelCellThd,
			data[j].ATTRIBUTES.NcellHoStatNum,
			data[j].ATTRIBUTES.StatisticPeriod,
			data[j].ATTRIBUTES.FastAnrRprtAmount,
			data[j].ATTRIBUTES.FastAnrRprtInterval,
			data[j].ATTRIBUTES.FastAnrCheckPeriod,
			data[j].ATTRIBUTES.FastAnrRsrpThd,
			data[j].ATTRIBUTES.FastAnrIntraRatMeasUeNum,
			data[j].ATTRIBUTES.FastAnrInterRatMeasUeNum,
			data[j].ATTRIBUTES.FastAnrIntraRatUeNumThd,
			data[j].ATTRIBUTES.FastAnrInterRatUeNumThd,
			data[j].ATTRIBUTES.OptMode,
			data[j].ATTRIBUTES.StatisticPeriodForNRTDel,
			data[j].ATTRIBUTES.StatisticNumForNRTDel,
			data[j].ATTRIBUTES.ActivePciConflictSwitch,
			data[j].ATTRIBUTES.StartTime,
			data[j].ATTRIBUTES.StopTime,
			data[j].ATTRIBUTES.FastAnrRscpThd,
			data[j].ATTRIBUTES.FastAnrRssiThd,
			data[j].ATTRIBUTES.FastAnrCdma1xrttPilotThd,
			data[j].ATTRIBUTES.FastAnrCdmahrpdPilotThd,
			data[j].ATTRIBUTES.NoHoSetThd,
			data[j].ATTRIBUTES.NcellHoForNRTDelThd,
			data[j].ATTRIBUTES.FastAnrMode,
			data[j].ATTRIBUTES.EventAnrMode,
			data[j].ATTRIBUTES.CaUeChoseMode,
			data[j].ATTRIBUTES.AnrControlledHoStrategy,
			data[j].ATTRIBUTES.SmartPreallocationMode,
			data[j].ATTRIBUTES.NoHoSetMode,
			data[j].ATTRIBUTES.UtranEventAnrMode,
			data[j].ATTRIBUTES.GeranEventAnrMode,
			data[j].ATTRIBUTES.EventAnrWithVoipMode,
			data[j].ATTRIBUTES.UtranEventAnrCgiTimer,
			data[j].ATTRIBUTES.GeranEventAnrCgiTimer,
			data[j].ATTRIBUTES.NrtDelMode,
			data[j].ATTRIBUTES.OptModeStrategy,
			data[j].ATTRIBUTES.UtranNcellHoForNRTDelThd,
			data[j].ATTRIBUTES.GeranNcellHoForNRTDelThd,
			data[j].ATTRIBUTES.NcellDelPunishPeriod,
			data[j].ATTRIBUTES.EutranNcellDelPunNum,
			data[j].ATTRIBUTES.UtranNcellDelPunNum,
			data[j].ATTRIBUTES.NcellCaThdForNRTDel,
			data[j].ATTRIBUTES.HoSucRateForCgiRead,
			data[j].ATTRIBUTES.StaPeriodForIRatNRTDel,
			data[j].ATTRIBUTES.StaNumForIRatNRTDel,
			data[j].ATTRIBUTES.PeriodForNCellRanking,
			data[j].ATTRIBUTES.StatPeriodCoeff,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Anr data has been saved")
	}
	tx.Commit()
}


func insertAnrmeasureparam(eNodeBId string, baseName string, data []model.Anrmeasureparam) {
	fmt.Println("[+] Processing Anrmeasureparam data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `anrmeasureparam` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.A3TimeToTrigForANR,
			data[j].ATTRIBUTES.A3HystForANR,
			data[j].ATTRIBUTES.A3OffsetForANR,
			data[j].ATTRIBUTES.A4TimeToTrigForANR,
			data[j].ATTRIBUTES.A4HystForANR,
			data[j].ATTRIBUTES.A4ThdRsrpForANR,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Anrmeasureparam data has been saved")
	}
	tx.Commit()
}


func insertAntennaport(eNodeBId string, baseName string, data []model.Antennaport) {
	fmt.Println("[+] Processing Antennaport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `antennaport` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.FEEDERLENGTH,
			data[j].ATTRIBUTES.DLDELAY,
			data[j].ATTRIBUTES.ULDELAY,
			data[j].ATTRIBUTES.PWRSWITCH,
			data[j].ATTRIBUTES.THRESHOLDTYPE,
			data[j].ATTRIBUTES.UOTHD,
			data[j].ATTRIBUTES.UCTHD,
			data[j].ATTRIBUTES.OOTHD,
			data[j].ATTRIBUTES.OCTHD,
			data[j].ATTRIBUTES.ULTRADELAYSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Antennaport data has been saved")
	}
	tx.Commit()
}


func insertAppcert(eNodeBId string, baseName string, data []model.Appcert) {
	fmt.Println("[+] Processing Appcert data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `appcert` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.APPTYPE,
			data[j].ATTRIBUTES.APPCERT,	

		)
		checkErr(err)
		//fmt.Println("[+] Appcert data has been saved")
	}
	tx.Commit()
}


func insertApplication(eNodeBId string, baseName string, data []model.Application) {
	fmt.Println("[+] Processing Application data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `application` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AID,
			data[j].ATTRIBUTES.AT,
			data[j].ATTRIBUTES.AN,
			data[j].ATTRIBUTES.SWVERSION,
			data[j].ATTRIBUTES.HOTPATCHVERSION,
			data[j].ATTRIBUTES.AV,
			data[j].ATTRIBUTES.APPHOTPATCHVERSION,
			data[j].ATTRIBUTES.APPMNTMODE,
			data[j].ATTRIBUTES.APPST,
			data[j].ATTRIBUTES.APPET,
			data[j].ATTRIBUTES.APPMMSETREMARK,
			data[j].ATTRIBUTES.EXTAPPTYPE,	

		)
		checkErr(err)
		//fmt.Println("[+] Application data has been saved")
	}
	tx.Commit()
}


func insertAsparagroup(eNodeBId string, baseName string, data []model.Asparagroup) {
	fmt.Println("[+] Processing Asparagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `asparagroup` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AsParaGroupID,
			data[j].ATTRIBUTES.AsPreallocDuration,
			data[j].ATTRIBUTES.AsPreallocMinPeriod,
			data[j].ATTRIBUTES.AsPreallocSize,
			data[j].ATTRIBUTES.AsSchPriFactor,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Asparagroup data has been saved")
	}
	tx.Commit()
}


func insertBasebandeqm(eNodeBId string, baseName string, data []model.Basebandeqm) {
	fmt.Println("[+] Processing Basebandeqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `basebandeqm` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BASEBANDEQMID,
			data[j].ATTRIBUTES.BASEBANDEQMTYPE,
			data[j].ATTRIBUTES.UMTSDEMMODE,
			data[j].ATTRIBUTES.BASEBANDEQMBOARD,	

		)
		checkErr(err)
		//fmt.Println("[+] Basebandeqm data has been saved")
	}
	tx.Commit()
}


func insertBbp(eNodeBId string, baseName string, data []model.Bbp) {
	fmt.Println("[+] Processing Bbp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `bbp` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.TYPE,
			data[j].ATTRIBUTES.ADMSTATE,
			data[j].ATTRIBUTES.TIME,
			data[j].ATTRIBUTES.BLKTP,
			data[j].ATTRIBUTES.HCE,
			data[j].ATTRIBUTES.CPRIEX,
			data[j].ATTRIBUTES.BRDSPEC,
			data[j].ATTRIBUTES.CCNE,
			data[j].ATTRIBUTES.BBWS,
			data[j].ATTRIBUTES.SRT,
			data[j].ATTRIBUTES.CPRIEXMODE,
			data[j].ATTRIBUTES.WM,
			data[j].ATTRIBUTES.CPRIITFTYPE,	

		)
		checkErr(err)
		//fmt.Println("[+] Bbp data has been saved")
	}
	tx.Commit()
}


func insertBbufan(eNodeBId string, baseName string, data []model.Bbufan) {
	fmt.Println("[+] Processing Bbufan data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `bbufan` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,	

		)
		checkErr(err)
		//fmt.Println("[+] Bbufan data has been saved")
	}
	tx.Commit()
}


func insertBcchcfg(eNodeBId string, baseName string, data []model.Bcchcfg) {
	fmt.Println("[+] Processing Bcchcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `bcchcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ModifyPeriodCoeff,
			data[j].ATTRIBUTES.ModifyPeriodCoeffForNb,	

		)
		checkErr(err)
		//fmt.Println("[+] Bcchcfg data has been saved")
	}
	tx.Commit()
}


func insertBlindncellopt(eNodeBId string, baseName string, data []model.Blindncellopt) {
	fmt.Println("[+] Processing Blindncellopt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `blindncellopt` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.StatisticPeriod,
			data[j].ATTRIBUTES.SampleNumThd,
			data[j].ATTRIBUTES.HoSuccRateThd,
			data[j].ATTRIBUTES.CsfbHoAttempRatioThd,
			data[j].ATTRIBUTES.BlindHoSuccRateThd,
			data[j].ATTRIBUTES.OptMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Blindncellopt data has been saved")
	}
	tx.Commit()
}


func insertBrdresassignment(eNodeBId string, baseName string, data []model.Brdresassignment) {
	fmt.Println("[+] Processing Brdresassignment data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `brdresassignment` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.BRDASSIGNMENT,	

		)
		checkErr(err)
		//fmt.Println("[+] Brdresassignment data has been saved")
	}
	tx.Commit()
}


func insertCabinet(eNodeBId string, baseName string, data []model.Cabinet) {
	fmt.Println("[+] Processing Cabinet data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cabinet` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.TYPE,
			data[j].ATTRIBUTES.DESC,
			data[j].ATTRIBUTES.LOCATIONNAME,	

		)
		checkErr(err)
		//fmt.Println("[+] Cabinet data has been saved")
	}
	tx.Commit()
}


func insertCagroup(eNodeBId string, baseName string, data []model.Cagroup) {
	fmt.Println("[+] Processing Cagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cagroup` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CaGroupId,
			data[j].ATTRIBUTES.CaGroupTypeInd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cagroup data has been saved")
	}
	tx.Commit()
}


func insertCagroupcell(eNodeBId string, baseName string, data []model.Cagroupcell) {
	fmt.Println("[+] Processing Cagroupcell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cagroupcell` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CaGroupId,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.PreferredPCellPriority,
			data[j].ATTRIBUTES.PCellA4RsrpThd,
			data[j].ATTRIBUTES.PCellA4RsrqThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cagroupcell data has been saved")
	}
	tx.Commit()
}


func insertCagroupscellcfg(eNodeBId string, baseName string, data []model.Cagroupscellcfg) {
	fmt.Println("[+] Processing Cagroupscellcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cagroupscellcfg` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SCellENODEB_ID,
			data[j].ATTRIBUTES.SCellLocalCellId,
			data[j].ATTRIBUTES.SCellBlindCfgFlag,
			data[j].ATTRIBUTES.SCellPriority,
			data[j].ATTRIBUTES.SCellA4Offset,
			data[j].ATTRIBUTES.SCellA2Offset,
			data[j].ATTRIBUTES.SpidGrpId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cagroupscellcfg data has been saved")
	}
	tx.Commit()
}


func insertCamgtcfg(eNodeBId string, baseName string, data []model.Camgtcfg) {
	fmt.Println("[+] Processing Camgtcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `camgtcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CarrAggrA2ThdRsrp,
			data[j].ATTRIBUTES.CarrAggrA4ThdRsrp,
			data[j].ATTRIBUTES.CarrierMgtSwitch,
			data[j].ATTRIBUTES.ActiveBufferLenThd,
			data[j].ATTRIBUTES.DeactiveBufferLenThd,
			data[j].ATTRIBUTES.ActiveBufferDelayThd,
			data[j].ATTRIBUTES.DeactiveThroughputThd,
			data[j].ATTRIBUTES.CarrAggrA6Offset,
			data[j].ATTRIBUTES.SCellAgingTime,
			data[j].ATTRIBUTES.CaA6ReportAmount,
			data[j].ATTRIBUTES.CaA6ReportInterval,
			data[j].ATTRIBUTES.SccDeactCqiThd,
			data[j].ATTRIBUTES.SccCfgInterval,
			data[j].ATTRIBUTES.CellCaAlgoSwitch,
			data[j].ATTRIBUTES.UlCaActiveTimeToTrigger,
			data[j].ATTRIBUTES.CaAmbrThd,
			data[j].ATTRIBUTES.MeasCycleSCell,
			data[j].ATTRIBUTES.CellMaxPccNumber,
			data[j].ATTRIBUTES.PccUserNumberOffloadThd,
			data[j].ATTRIBUTES.EnhancedPccAnchorA1ThdRsrp,
			data[j].ATTRIBUTES.AddedMeasCfg,
			data[j].ATTRIBUTES.BlindScellSampleNum,
			data[j].ATTRIBUTES.OptMode,
			data[j].ATTRIBUTES.SleepPeriod,
			data[j].ATTRIBUTES.StatPeriod,
			data[j].ATTRIBUTES.RelaxedBHSccDeactCqiThd,
			data[j].ATTRIBUTES.RelaxedBackhaulCaMaxCcNum,
			data[j].ATTRIBUTES.DisCloudBBCaMaxCcNum,
			data[j].ATTRIBUTES.CaA2TimeToTrigger,
			data[j].ATTRIBUTES.CaA6TimeToTrigger,
			data[j].ATTRIBUTES.CaA1TimeToTrigger,
			data[j].ATTRIBUTES.CaA4TimeToTrigger,
			data[j].ATTRIBUTES.SccQuietTime,
			data[j].ATTRIBUTES.FTRelaxedBHCaDLMaxCcNum,
			data[j].ATTRIBUTES.FddTddCaUlMaxCcNum,
			data[j].ATTRIBUTES.FddTddCaDlMaxCcNum,
			data[j].ATTRIBUTES.LaaCarrAggrA1ThdRsrp,
			data[j].ATTRIBUTES.LaaCarrAggrA2ThdRsrp,
			data[j].ATTRIBUTES.FTCA2CCAnchorPolicy,
			data[j].ATTRIBUTES.FTCAMultiCCAnchorPolicy,
			data[j].ATTRIBUTES.UlHeavyTrafficMlbTFCAOptSw,
			data[j].ATTRIBUTES.SccReactivationTime,
			data[j].ATTRIBUTES.FTRelaxedBHCaUlMaxCcNum,
			data[j].ATTRIBUTES.RelaxedBHCaUlMaxCcNum,
			data[j].ATTRIBUTES.CaTrafficDirectionPref,
			data[j].ATTRIBUTES.CaMimoPriorityStrategySw,
			data[j].ATTRIBUTES.EnhancedSccSelA1ThldRsrp,
			data[j].ATTRIBUTES.FastScellSelAftScellRmvSw,
			data[j].ATTRIBUTES.RsrpOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Camgtcfg data has been saved")
	}
	tx.Commit()
}


func insertCascadeport(eNodeBId string, baseName string, data []model.Cascadeport) {
	fmt.Println("[+] Processing Cascadeport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cascadeport` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.SW,
			data[j].ATTRIBUTES.PM,	

		)
		checkErr(err)
		//fmt.Println("[+] Cascadeport data has been saved")
	}
	tx.Commit()
}


func insertCell(eNodeBId string, baseName string, data []model.Cell) {
	fmt.Println("[+] Processing Cell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellName,
			data[j].ATTRIBUTES.CsgInd,
			data[j].ATTRIBUTES.UlCyclicPrefix,
			data[j].ATTRIBUTES.DlCyclicPrefix,
			data[j].ATTRIBUTES.FreqBand,
			data[j].ATTRIBUTES.UlEarfcnCfgInd,
			data[j].ATTRIBUTES.DlEarfcn,
			data[j].ATTRIBUTES.UlBandWidth,
			data[j].ATTRIBUTES.DlBandWidth,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.PhyCellId,
			data[j].ATTRIBUTES.AdditionalSpectrumEmission,
			data[j].ATTRIBUTES.CellActiveState,
			data[j].ATTRIBUTES.CellAdminState,
			data[j].ATTRIBUTES.FddTddInd,
			data[j].ATTRIBUTES.CellSpecificOffset,
			data[j].ATTRIBUTES.QoffsetFreq,
			data[j].ATTRIBUTES.RootSequenceIdx,
			data[j].ATTRIBUTES.HighSpeedFlag,
			data[j].ATTRIBUTES.PreambleFmt,
			data[j].ATTRIBUTES.CellRadius,
			data[j].ATTRIBUTES.CustomizedBandWidthCfgInd,
			data[j].ATTRIBUTES.EmergencyAreaIdCfgInd,
			data[j].ATTRIBUTES.UePowerMaxCfgInd,
			data[j].ATTRIBUTES.MultiRruCellFlag,
			data[j].ATTRIBUTES.CPRICompression,
			data[j].ATTRIBUTES.AirCellFlag,
			data[j].ATTRIBUTES.CrsPortNum,
			data[j].ATTRIBUTES.TxRxMode,
			data[j].ATTRIBUTES.UserLabel,
			data[j].ATTRIBUTES.WorkMode,
			data[j].ATTRIBUTES.EuCellStandbyMode,
			data[j].ATTRIBUTES.SfnMasterCellLabel,
			data[j].ATTRIBUTES.CellSlaveBand,
			data[j].ATTRIBUTES.CnOpSharingGroupId,
			data[j].ATTRIBUTES.IntraFreqRanSharingInd,
			data[j].ATTRIBUTES.IntraFreqAnrInd,
			data[j].ATTRIBUTES.CellScaleInd,
			data[j].ATTRIBUTES.FreqPriorityForAnr,
			data[j].ATTRIBUTES.CellRadiusStartLocation,
			data[j].ATTRIBUTES.SpecifiedCellFlag,
			data[j].ATTRIBUTES.MultiCellShareMode,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.NbCellFlag,
			data[j].ATTRIBUTES.SubframeAssignment,
			data[j].ATTRIBUTES.SpecialSubframePatterns,
			data[j].ATTRIBUTES.CrsPortMap,
			data[j].ATTRIBUTES.CsiRsPeriod,
			data[j].ATTRIBUTES.UlEarfcn,
			data[j].ATTRIBUTES.UePowerMax,	

		)
		checkErr(err)
		//fmt.Println("[+] Cell data has been saved")
	}
	tx.Commit()
}


func insertCellacbar(eNodeBId string, baseName string, data []model.Cellacbar) {
	fmt.Println("[+] Processing Cellacbar data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellacbar` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.AcBarringInfoCfgInd,
			data[j].ATTRIBUTES.VoLTEPreferCfgInd,
			data[j].ATTRIBUTES.AcBarringForEmergency,
			data[j].ATTRIBUTES.AcBarringForMoDataCfgInd,
			data[j].ATTRIBUTES.AcBarringFactorForCall,
			data[j].ATTRIBUTES.AcBarTimeForCall,
			data[j].ATTRIBUTES.AC11BarforCall,
			data[j].ATTRIBUTES.AC12BarforCall,
			data[j].ATTRIBUTES.AC13BarforCall,
			data[j].ATTRIBUTES.AC14BarforCall,
			data[j].ATTRIBUTES.AC15BarforCall,
			data[j].ATTRIBUTES.AcBarringForMoSigCfgInd,
			data[j].ATTRIBUTES.AcBarringFactorForSig,
			data[j].ATTRIBUTES.AcBarTimeForSig,
			data[j].ATTRIBUTES.AC11BarForSig,
			data[j].ATTRIBUTES.AC12BarForSig,
			data[j].ATTRIBUTES.AC13BarForSig,
			data[j].ATTRIBUTES.AC14BarForSig,
			data[j].ATTRIBUTES.AC15BarForSig,
			data[j].ATTRIBUTES.AcBarForMVoiceCfgInd,
			data[j].ATTRIBUTES.AcBarForMVideoCfgInd,
			data[j].ATTRIBUTES.AcBarForCsfbCfgInd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellacbar data has been saved")
	}
	tx.Commit()
}


func insertCellaccess(eNodeBId string, baseName string, data []model.Cellaccess) {
	fmt.Println("[+] Processing Cellaccess data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellaccess` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellBarred,
			data[j].ATTRIBUTES.IntraFreqResel,
			data[j].ATTRIBUTES.BroadcastMode,
			data[j].ATTRIBUTES.RoundPeriod,
			data[j].ATTRIBUTES.RoundActionTime,
			data[j].ATTRIBUTES.ReptSyncAvoidInd,
			data[j].ATTRIBUTES.ReptSyncAvoidTime,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellaccess data has been saved")
	}
	tx.Commit()
}


func insertCellalgoswitch(eNodeBId string, baseName string, data []model.Cellalgoswitch) {
	fmt.Println("[+] Processing Cellalgoswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellalgoswitch` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.VolteRedirectSwitch,
			data[j].ATTRIBUTES.RachAlgoSwitch,
			data[j].ATTRIBUTES.SrsAlgoSwitch,
			data[j].ATTRIBUTES.PucchAlgoSwitch,
			data[j].ATTRIBUTES.AqmAlgoSwitch,
			data[j].ATTRIBUTES.CqiAdjAlgoSwitch,
			data[j].ATTRIBUTES.DynAdjVoltSwitch,
			data[j].ATTRIBUTES.RacAlgoSwitch,
			data[j].ATTRIBUTES.MlbAlgoSwitch,
			data[j].ATTRIBUTES.DlPcAlgoSwitch,
			data[j].ATTRIBUTES.UlPcAlgoSwitch,
			data[j].ATTRIBUTES.BfAlgoSwitch,
			data[j].ATTRIBUTES.DlSchSwitch,
			data[j].ATTRIBUTES.UlSchSwitch,
			data[j].ATTRIBUTES.RanShareModeSwitch,
			data[j].ATTRIBUTES.FreqPriorityHoSwitch,
			data[j].ATTRIBUTES.MuBfAlgoSwitch,
			data[j].ATTRIBUTES.DistBasedHoSwitch,
			data[j].ATTRIBUTES.AcBarAlgoSwitch,
			data[j].ATTRIBUTES.SfnUlSchSwitch,
			data[j].ATTRIBUTES.SfnDlSchSwitch,
			data[j].ATTRIBUTES.IrcSwitch,
			data[j].ATTRIBUTES.DynDrxSwitch,
			data[j].ATTRIBUTES.HighMobiTrigIdleModeSwitch,
			data[j].ATTRIBUTES.AvoidInterfSwitch,
			data[j].ATTRIBUTES.GLPwrShare,
			data[j].ATTRIBUTES.EicicSwitch,
			data[j].ATTRIBUTES.PUSCHMaxRBPUCCHAdjSwitch,
			data[j].ATTRIBUTES.DlCompSwitch,
			data[j].ATTRIBUTES.PsicSwitch,
			data[j].ATTRIBUTES.MlbHoMode,
			data[j].ATTRIBUTES.UplinkCompSwitch,
			data[j].ATTRIBUTES.AntCalibrationAlgoSwitch,
			data[j].ATTRIBUTES.DynSpectrumShareSwitch,
			data[j].ATTRIBUTES.SfnLoadBasedAdptSwitch,
			data[j].ATTRIBUTES.PuschIrcAlgoSwitch,
			data[j].ATTRIBUTES.ReselecPriAdaptSwitch,
			data[j].ATTRIBUTES.AnrFunctionSwitch,
			data[j].ATTRIBUTES.SfnAlgoSwitch,
			data[j].ATTRIBUTES.PrachIntrfRejSwitch,
			data[j].ATTRIBUTES.EnhMIMOSwitch,
			data[j].ATTRIBUTES.InterfRandSwitch,
			data[j].ATTRIBUTES.RepeaterSwitch,
			data[j].ATTRIBUTES.MultiFreqPriControlSwitch,
			data[j].ATTRIBUTES.HarqAlgoSwitch,
			data[j].ATTRIBUTES.CovBasedInterFreqHoMode,
			data[j].ATTRIBUTES.LteUtcBroadcastSwitch,
			data[j].ATTRIBUTES.CellSchStrategySwitch,
			data[j].ATTRIBUTES.SsrdAlgoSwitch,
			data[j].ATTRIBUTES.SfnUplinkCompSwitch,
			data[j].ATTRIBUTES.LowSpeedInterFreqHoSwitch,
			data[j].ATTRIBUTES.RelaySwitch,
			data[j].ATTRIBUTES.InterFreqDirectHoSwitch,
			data[j].ATTRIBUTES.PwrDeratSwitch,
			data[j].ATTRIBUTES.DetectionAlgoSwitch,
			data[j].ATTRIBUTES.PucchIRCEnhance,
			data[j].ATTRIBUTES.AcBarAlgoforDynSwitch,
			data[j].ATTRIBUTES.CreSwitch,
			data[j].ATTRIBUTES.BackoffAlgoSwitch,
			data[j].ATTRIBUTES.HoAllowedSwitch,
			data[j].ATTRIBUTES.NCellClassMgtSw,
			data[j].ATTRIBUTES.SpePCIBasedPolicySw,
			data[j].ATTRIBUTES.CellPIMInterMitigSwitch,
			data[j].ATTRIBUTES.PrachJointReceptionSwitch,
			data[j].ATTRIBUTES.FeicicSwitch,
			data[j].ATTRIBUTES.CamcSwitch,
			data[j].ATTRIBUTES.RruUeMapSwitch,
			data[j].ATTRIBUTES.HighSpeedSchOptSwitch,
			data[j].ATTRIBUTES.ServiceDiffSwitch,
			data[j].ATTRIBUTES.LtePttQoSSwitch,
			data[j].ATTRIBUTES.SrsPucchEnhancedSwitch,
			data[j].ATTRIBUTES.UEInactiveTimerQCI1Switch,
			data[j].ATTRIBUTES.UlJRAntNumCombSw,
			data[j].ATTRIBUTES.VamPhaseShiftAlgoSwitch,
			data[j].ATTRIBUTES.AnrAlgoSwitch,
			data[j].ATTRIBUTES.Dl256QamAlgoSwitch,
			data[j].ATTRIBUTES.MlbAutoGroupSwitch,
			data[j].ATTRIBUTES.CaAutoGroupSwitch,
			data[j].ATTRIBUTES.AttachCellSelfCfgSwitch,
			data[j].ATTRIBUTES.CellDlCoverEnhanceSwitch,
			data[j].ATTRIBUTES.UlSchExtSwitch,
			data[j].ATTRIBUTES.UdcAlgoSwitch,
			data[j].ATTRIBUTES.VoLTESwitch,
			data[j].ATTRIBUTES.VoipFailDecSelfRecSwitch,
			data[j].ATTRIBUTES.DeprioritisationDeliverInd,
			data[j].ATTRIBUTES.CmasSwitch,
			data[j].ATTRIBUTES.MFBIAlgoSwitch,
			data[j].ATTRIBUTES.PciConflictReportSwitch,
			data[j].ATTRIBUTES.MroSwitch,
			data[j].ATTRIBUTES.OpResourceGroupShareSwitch,
			data[j].ATTRIBUTES.CellRecoverySwitch,
			data[j].ATTRIBUTES.EnhancedMlbAlgoSwitch,
			data[j].ATTRIBUTES.TrafficMlbSwitch,
			data[j].ATTRIBUTES.MtcSwitch,
			data[j].ATTRIBUTES.UlIcSwitch,
			data[j].ATTRIBUTES.FcsMode,
			data[j].ATTRIBUTES.ScVideoOptSwitch,
			data[j].ATTRIBUTES.SpidSpecificHoSwitch,
			data[j].ATTRIBUTES.DMIMOAlgoSwitch,
			data[j].ATTRIBUTES.UlAmrcMode,
			data[j].ATTRIBUTES.AmrcAlgoSwitch,
			data[j].ATTRIBUTES.CrsIcSwitch,
			data[j].ATTRIBUTES.MeasOptAlgoSwitch,
			data[j].ATTRIBUTES.FreqLayerSwitch,
			data[j].ATTRIBUTES.EmimoSwitch,
			data[j].ATTRIBUTES.RohcSwitch,
			data[j].ATTRIBUTES.McpttSwitch,
			data[j].ATTRIBUTES.UlIcsRbRatio,
			data[j].ATTRIBUTES.UlIcsVoLTEPLTh,
			data[j].ATTRIBUTES.DacqSwitch,
			data[j].ATTRIBUTES.RrcReestDataFwdSwitch,
			data[j].ATTRIBUTES.MTCCongControlSwitch,
			data[j].ATTRIBUTES.MTCPowerSavSwitch,
			data[j].ATTRIBUTES.TcpCtrlSwitch,
			data[j].ATTRIBUTES.TurboReceiverSwitch,
			data[j].ATTRIBUTES.NaicsSwitch,
			data[j].ATTRIBUTES.UplinkIcSwitch,
			data[j].ATTRIBUTES.DacqEnhancementSwitch,
			data[j].ATTRIBUTES.AsAlgoSwitch,
			data[j].ATTRIBUTES.EnhChnCalSwitch,
			data[j].ATTRIBUTES.NbCellAlgoSwitch,
			data[j].ATTRIBUTES.MultiCnConnFreqPriSw,
			data[j].ATTRIBUTES.NoSrsSccBfAlgoSwitch,
			data[j].ATTRIBUTES.SpecUserAlgoSwitch,
			data[j].ATTRIBUTES.MimoUePaAdaptSw,
			data[j].ATTRIBUTES.UlSuMimoAlgoSwitch,
			data[j].ATTRIBUTES.HighSpeedInterRatHoSwitch,
			data[j].ATTRIBUTES.EmcSpsSchSwitch,
			data[j].ATTRIBUTES.LbsSwitch,
			data[j].ATTRIBUTES.DtxDetectionAlgoSwitch,
			data[j].ATTRIBUTES.DlSchExtSwitch,
			data[j].ATTRIBUTES.MPMUDetectSwitch,
			data[j].ATTRIBUTES.VmsSwitch,
			data[j].ATTRIBUTES.SmallBandOptSwitch,
			data[j].ATTRIBUTES.IdleModeEdrxSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellalgoswitch data has been saved")
	}
	tx.Commit()
}


func insertCellbackoff(eNodeBId string, baseName string, data []model.Cellbackoff) {
	fmt.Println("[+] Processing Cellbackoff data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellbackoff` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterFreqBackoffPrbThd,
			data[j].ATTRIBUTES.InterFreqBackoffUeNumThd,
			data[j].ATTRIBUTES.RBPriMcsSelectTrigPrbThd,
			data[j].ATTRIBUTES.RBPriMcsSelectStopPrbThd,
			data[j].ATTRIBUTES.LowestEffDlMcsThd,
			data[j].ATTRIBUTES.LowEffDlMcsThd,
			data[j].ATTRIBUTES.LowEffDlFactorA,
			data[j].ATTRIBUTES.LowEffDlFactorK,
			data[j].ATTRIBUTES.LowestEffUlMcsThd,
			data[j].ATTRIBUTES.LowEffUlMcsThd,
			data[j].ATTRIBUTES.LowEffUlFactorA,
			data[j].ATTRIBUTES.LowEffUlFactorK,
			data[j].ATTRIBUTES.UlHeavyTrafficTtiProporThd,
			data[j].ATTRIBUTES.UlTrafficMlbUeNumThd,
			data[j].ATTRIBUTES.UlHeavyTrafficJudgePeriod,
			data[j].ATTRIBUTES.HoInRejectPrbThd,
			data[j].ATTRIBUTES.HoInRejectUeNumThd,
			data[j].ATTRIBUTES.RejectedHoInCause,
			data[j].ATTRIBUTES.BackoffCAUserHOSw,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellbackoff data has been saved")
	}
	tx.Commit()
}


func insertCellbf(eNodeBId string, baseName string, data []model.Cellbf) {
	fmt.Println("[+] Processing Cellbf data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellbf` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MaxBfRankPara,
			data[j].ATTRIBUTES.DualLayerBFAlgType,
			data[j].ATTRIBUTES.HighOrderMubfMaxLayer,
			data[j].ATTRIBUTES.HighOrderMubfPairRule,
			data[j].ATTRIBUTES.AntSelUEMubfPairMode,
			data[j].ATTRIBUTES.NonAntSelUEMubfPairMode,
			data[j].ATTRIBUTES.MovingUeMuBfScheme,
			data[j].ATTRIBUTES.QualUEPortAvoidMode,
			data[j].ATTRIBUTES.WaitPairingLayerThd,
			data[j].ATTRIBUTES.MassiveMIMOMubfPairRule,
			data[j].ATTRIBUTES.MultiLayerThdSwitchToTM7,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellbf data has been saved")
	}
	tx.Commit()
}


func insertCellbfmimoparacfg(eNodeBId string, baseName string, data []model.Cellbfmimoparacfg) {
	fmt.Println("[+] Processing Cellbfmimoparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellbfmimoparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BfMimoAdaptiveSwitch,
			data[j].ATTRIBUTES.BfMimoAdapWithoutTm2,
			data[j].ATTRIBUTES.BfSingleToDualThdOffset,
			data[j].ATTRIBUTES.TM3Rank2ToDualBfThdOffset,
			data[j].ATTRIBUTES.AsUeDualBfToTM3Rank2Offset,
			data[j].ATTRIBUTES.AsUeTM3Rank2ToDualBfOffset,
			data[j].ATTRIBUTES.AsUeBfSingleToDualOffset,
			data[j].ATTRIBUTES.DualBfToTM3Rank2Offset,
			data[j].ATTRIBUTES.OffsetOfInStat,
			data[j].ATTRIBUTES.OffsetOfOutStat,
			data[j].ATTRIBUTES.AntBasedBfMimoAlgoSelect,
			data[j].ATTRIBUTES.Tm3DirectToDualBfSwitch,
			data[j].ATTRIBUTES.SccBfMimoAdaptiveSwitch,
			data[j].ATTRIBUTES.HighLowSpeedUeThdOffset,
			data[j].ATTRIBUTES.TmAccelerationSwitch,
			data[j].ATTRIBUTES.BfMimoAdapWithTm4Switch,
			data[j].ATTRIBUTES.Tm3AndTm9ThdOffset,
			data[j].ATTRIBUTES.BfTo2LayerMubfThdOffset,
			data[j].ATTRIBUTES.BfTo4LayerMubfThdOffset,
			data[j].ATTRIBUTES.Tm3Rank2ToTm9Rank4Offset,
			data[j].ATTRIBUTES.Tm9Rank4ToTm3Rank2Offset,
			data[j].ATTRIBUTES.VolteMimoAdaptiveSwitch,
			data[j].ATTRIBUTES.HoBfThdAdjSwitch,
			data[j].ATTRIBUTES.CaSccAddThldOffset,
			data[j].ATTRIBUTES.CaSccDelThldOffset,
			data[j].ATTRIBUTES.BfMimoAlgoOptSwitch,
			data[j].ATTRIBUTES.InitialBfMimoType,
			data[j].ATTRIBUTES.MultiLayerPairIsoThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellbfmimoparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellcecfg(eNodeBId string, baseName string, data []model.Cellcecfg) {
	fmt.Println("[+] Processing Cellcecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcecfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.RachRsrpFstThd,
			data[j].ATTRIBUTES.RachRsrpSndThd,
			data[j].ATTRIBUTES.RachRsrpTrdThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcecfg data has been saved")
	}
	tx.Commit()
}


func insertCellceschcfg(eNodeBId string, baseName string, data []model.Cellceschcfg) {
	fmt.Println("[+] Processing Cellceschcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellceschcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.SIB1RepNum,
			data[j].ATTRIBUTES.PdschMaxNumRepModeA,
			data[j].ATTRIBUTES.PdschMaxNumRepModeB,
			data[j].ATTRIBUTES.PuschMaxNumRepModeA,
			data[j].ATTRIBUTES.PuschMaxNumRepModeB,
			data[j].ATTRIBUTES.MpdcchMaxNumRepPaging,
			data[j].ATTRIBUTES.MpdcchMaxNumRepModeA,
			data[j].ATTRIBUTES.MpdcchMaxNumRepModeB,
			data[j].ATTRIBUTES.SiTransEcr,
			data[j].ATTRIBUTES.PagingGroupNum,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellceschcfg data has been saved")
	}
	tx.Commit()
}


func insertCellchpwrcfg(eNodeBId string, baseName string, data []model.Cellchpwrcfg) {
	fmt.Println("[+] Processing Cellchpwrcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellchpwrcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PcfichPwr,
			data[j].ATTRIBUTES.PbchPwr,
			data[j].ATTRIBUTES.SchPwr,
			data[j].ATTRIBUTES.DbchPwr,
			data[j].ATTRIBUTES.PchPwr,
			data[j].ATTRIBUTES.RaRspPwr,
			data[j].ATTRIBUTES.PrsPwr,
			data[j].ATTRIBUTES.AntOutputPwr,
			data[j].ATTRIBUTES.PmchPwrOffset,
			data[j].ATTRIBUTES.HoRarPwrEnhancedSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellchpwrcfg data has been saved")
	}
	tx.Commit()
}


func insertCellcounterparagroup(eNodeBId string, baseName string, data []model.Cellcounterparagroup) {
	fmt.Println("[+] Processing Cellcounterparagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcounterparagroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellUserNumPdfThd,
			data[j].ATTRIBUTES.UlIblerOptSwitch,
			data[j].ATTRIBUTES.CellCounterAlgoSwitch,
			data[j].ATTRIBUTES.EdgeUserA3Offset,
			data[j].ATTRIBUTES.DlThrpPdfThd0,
			data[j].ATTRIBUTES.DlThrpPdfThd1,
			data[j].ATTRIBUTES.DlThrpPdfThd2,
			data[j].ATTRIBUTES.DlThrpPdfThd3,
			data[j].ATTRIBUTES.DlThrpPdfThd4,
			data[j].ATTRIBUTES.DlThrpPdfThd5,
			data[j].ATTRIBUTES.DlThrpPdfThd6,
			data[j].ATTRIBUTES.DlThrpPdfThd7,
			data[j].ATTRIBUTES.DlThrpPdfThd8,
			data[j].ATTRIBUTES.UlThrpPdfThd0,
			data[j].ATTRIBUTES.UlThrpPdfThd1,
			data[j].ATTRIBUTES.UlThrpPdfThd2,
			data[j].ATTRIBUTES.UlThrpPdfThd3,
			data[j].ATTRIBUTES.UlThrpPdfThd4,
			data[j].ATTRIBUTES.UlThrpPdfThd5,
			data[j].ATTRIBUTES.UlThrpPdfThd6,
			data[j].ATTRIBUTES.UlThrpPdfThd7,
			data[j].ATTRIBUTES.UlThrpPdfThd8,
			data[j].ATTRIBUTES.UeAbnormalJudgeThd,
			data[j].ATTRIBUTES.ThrPdfPolicy,
			data[j].ATTRIBUTES.EdgeUserServRSRPThd,
			data[j].ATTRIBUTES.EdgeUserSRSOffset,
			data[j].ATTRIBUTES.ChrOutputModeSwitch,
			data[j].ATTRIBUTES.DlUserXpUnsatiThd,
			data[j].ATTRIBUTES.DlTrafficVolumeThd,
			data[j].ATTRIBUTES.DlPktDelayMeasPolicy,
			data[j].ATTRIBUTES.QCI1ContinuousPktLossThld,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcounterparagroup data has been saved")
	}
	tx.Commit()
}


func insertCellcqiadaptivecfg(eNodeBId string, baseName string, data []model.Cellcqiadaptivecfg) {
	fmt.Println("[+] Processing Cellcqiadaptivecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcqiadaptivecfg` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CqiPeriodAdaptive,
			data[j].ATTRIBUTES.SimulAckNackAndCqiSwitch,
			data[j].ATTRIBUTES.HoAperiodicCqiCfgSwitch,
			data[j].ATTRIBUTES.MinCqiPeriod,
			data[j].ATTRIBUTES.SccCqiRptEnhancedSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcqiadaptivecfg data has been saved")
	}
	tx.Commit()
}


func insertCellcqiadjalgo(eNodeBId string, baseName string, data []model.Cellcqiadjalgo) {
	fmt.Println("[+] Processing Cellcqiadjalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcqiadjalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InitDlIblerTarget,
			data[j].ATTRIBUTES.InitDeltaCqi,
			data[j].ATTRIBUTES.CqiAdjStep,
			data[j].ATTRIBUTES.CqiFilterCoefForMcs,
			data[j].ATTRIBUTES.VolteNackDeltaCqi,
			data[j].ATTRIBUTES.DlVolteCqiAdjOptSw,
			data[j].ATTRIBUTES.CqiOptSwitch,
			data[j].ATTRIBUTES.InitDlIblerTargetforQCI2,
			data[j].ATTRIBUTES.InitDlIblerTargetforVoLTE,
			data[j].ATTRIBUTES.CellDeltaCqiSampSelThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcqiadjalgo data has been saved")
	}
	tx.Commit()
}


func insertCellcsirsparacfg(eNodeBId string, baseName string, data []model.Cellcsirsparacfg) {
	fmt.Println("[+] Processing Cellcsirsparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcsirsparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CsiRsSwitch,
			data[j].ATTRIBUTES.CsiRsPeriod,
			data[j].ATTRIBUTES.CsiRsConfigUserNumTh,
			data[j].ATTRIBUTES.CsiRsUnconfigUserNumTh,
			data[j].ATTRIBUTES.CsiRsConfigUserRatioTh,
			data[j].ATTRIBUTES.CsiRsUnconfigUserRatioTh,
			data[j].ATTRIBUTES.CsiRsSetJudgeHysTimer,
			data[j].ATTRIBUTES.CsiRsSetJudgeTimer,
			data[j].ATTRIBUTES.CsiRsPortNum,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcsirsparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellcspcpara(eNodeBId string, baseName string, data []model.Cellcspcpara) {
	fmt.Println("[+] Processing Cellcspcpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellcspcpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ECspcPCAdjUeNumTh,
			data[j].ATTRIBUTES.CellCspcSwitch,
			data[j].ATTRIBUTES.CspcRapidRptSwitch,
			data[j].ATTRIBUTES.CspcUeSrsCfgRptPeriod,
			data[j].ATTRIBUTES.UlRsrpRptPeriod,
			data[j].ATTRIBUTES.CspcCqiFilterCoeff,
			data[j].ATTRIBUTES.UlRsrpPhyFilterCoeff,
			data[j].ATTRIBUTES.UlRsrpRrmFilterCoeff,
			data[j].ATTRIBUTES.ECspcA3Offset,
			data[j].ATTRIBUTES.CelleCspcSwitch,
			data[j].ATTRIBUTES.ECspcPCAdjRange,
			data[j].ATTRIBUTES.IntraEnbCspcSw,
			data[j].ATTRIBUTES.ElasticCarrierSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellcspcpara data has been saved")
	}
	tx.Commit()
}


func insertCelldacqcfg(eNodeBId string, baseName string, data []model.Celldacqcfg) {
	fmt.Println("[+] Processing Celldacqcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldacqcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlAmbrLimitFirstRank,
			data[j].ATTRIBUTES.DlAmbrLimitSecRank,
			data[j].ATTRIBUTES.UlAmbrLimitFirstRank,
			data[j].ATTRIBUTES.UlAmbrLimitSecRank,
			data[j].ATTRIBUTES.DlPrbUsageFirstRank,
			data[j].ATTRIBUTES.DlPrbUsageSecRank,
			data[j].ATTRIBUTES.UlPrbUsageFirstRank,
			data[j].ATTRIBUTES.UlPrbUsageSecRank,
			data[j].ATTRIBUTES.DacqExecutionDuration,
			data[j].ATTRIBUTES.DacqCongMonDur,
			data[j].ATTRIBUTES.AmbrLimitSchUserNum,
			data[j].ATTRIBUTES.AmbrProtectSchUserNum,
			data[j].ATTRIBUTES.DlAmbrLimitThirdRank,
			data[j].ATTRIBUTES.DlPrbUsageThirdRank,
			data[j].ATTRIBUTES.UlAmbrLimitThirdRank,
			data[j].ATTRIBUTES.UlPrbUsageThirdRank,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldacqcfg data has been saved")
	}
	tx.Commit()
}


func insertCelldlcompalgo(eNodeBId string, baseName string, data []model.Celldlcompalgo) {
	fmt.Println("[+] Processing Celldlcompalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlcompalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlCompA3Offset,
			data[j].ATTRIBUTES.DpsCoCellUserRatioThd,
			data[j].ATTRIBUTES.DpsLoadDiffThd,
			data[j].ATTRIBUTES.DpsServingCellDlPrbThd,
			data[j].ATTRIBUTES.HetnetDpsCoCellRsrpThd,
			data[j].ATTRIBUTES.HomnetDpsCoCellRsrpThd,
			data[j].ATTRIBUTES.JtCoCellDlPrbThd,
			data[j].ATTRIBUTES.JtCoCellRsrpThd,
			data[j].ATTRIBUTES.JtServingCellDlPrbThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlcompalgo data has been saved")
	}
	tx.Commit()
}


func insertCelldlicic(eNodeBId string, baseName string, data []model.Celldlicic) {
	fmt.Println("[+] Processing Celldlicic data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlicic` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BandMode,
			data[j].ATTRIBUTES.DlIcicUserAttrGfactorThd,
			data[j].ATTRIBUTES.DlIcicNoiseUserRsrpThd,
			data[j].ATTRIBUTES.AIcIcPlusA3Offset,
			data[j].ATTRIBUTES.AIcicPlusPCAdjRange,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlicic data has been saved")
	}
	tx.Commit()
}


func insertCelldlpcpdcch(eNodeBId string, baseName string, data []model.Celldlpcpdcch) {
	fmt.Println("[+] Processing Celldlpcpdcch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlpcpdcch` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DediDciPwrOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlpcpdcch data has been saved")
	}
	tx.Commit()
}


func insertCelldlpcpdsch(eNodeBId string, baseName string, data []model.Celldlpcpdsch) {
	fmt.Println("[+] Processing Celldlpcpdsch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlpcpdsch` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CcuPa,
			data[j].ATTRIBUTES.CeuPa,
			data[j].ATTRIBUTES.CellEdgLoadEvalPrd,
			data[j].ATTRIBUTES.NeighCellEdgLoadThd,
			data[j].ATTRIBUTES.ExceedNCellEdgLoadThdRatio,
			data[j].ATTRIBUTES.BFUserPwrBoostPrd,
			data[j].ATTRIBUTES.BFUserAdptPwrA3Offset,
			data[j].ATTRIBUTES.BFUserAdptPwrA6Offset,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlpcpdsch data has been saved")
	}
	tx.Commit()
}


func insertCelldlpcpdschpa(eNodeBId string, baseName string, data []model.Celldlpcpdschpa) {
	fmt.Println("[+] Processing Celldlpcpdschpa data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlpcpdschpa` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PdschPaAdjSwitch,
			data[j].ATTRIBUTES.PaPcOff,
			data[j].ATTRIBUTES.NomPdschRsEpreOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlpcpdschpa data has been saved")
	}
	tx.Commit()
}


func insertCelldlpcphich(eNodeBId string, baseName string, data []model.Celldlpcphich) {
	fmt.Println("[+] Processing Celldlpcphich data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlpcphich` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PwrOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlpcphich data has been saved")
	}
	tx.Commit()
}


func insertCelldlschalgo(eNodeBId string, baseName string, data []model.Celldlschalgo) {
	fmt.Println("[+] Processing Celldlschalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldlschalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlschStrategy,
			data[j].ATTRIBUTES.FreeUserDlRbUsedRatio,
			data[j].ATTRIBUTES.MaxMimoRankPara,
			data[j].ATTRIBUTES.CaSchStrategy,
			data[j].ATTRIBUTES.NonGbrResourceRatio,
			data[j].ATTRIBUTES.DlEpfCapacityFactor,
			data[j].ATTRIBUTES.RBPriMcsSelectRatioThd,
			data[j].ATTRIBUTES.CellEdgeUserRbAllocMode,
			data[j].ATTRIBUTES.DlIcicSchMode,
			data[j].ATTRIBUTES.RbgAllocStrategy,
			data[j].ATTRIBUTES.MidUserMcsThreshold,
			data[j].ATTRIBUTES.DlLowLoadSdmaThdOffset,
			data[j].ATTRIBUTES.DlHarqMaxTxNum,
			data[j].ATTRIBUTES.NonGbrSchDelayWeight,
			data[j].ATTRIBUTES.RBDamageNearPointIblerTh,
			data[j].ATTRIBUTES.BreathingPilotAlgoSwitch,
			data[j].ATTRIBUTES.HoStaticMcsTimer,
			data[j].ATTRIBUTES.DlRankOptimizeSwitch,
			data[j].ATTRIBUTES.CaSccDopMeas,
			data[j].ATTRIBUTES.DlSpsInterval,
			data[j].ATTRIBUTES.RBPriMcsSelectStrategy,
			data[j].ATTRIBUTES.DlRankDetectSwitch,
			data[j].ATTRIBUTES.MbsfnSfCfg,
			data[j].ATTRIBUTES.TddSpecSubfSchSwitch,
			data[j].ATTRIBUTES.DataThdInPdcchPdschBal,
			data[j].ATTRIBUTES.UeNumThdInPdcchPdschBal,
			data[j].ATTRIBUTES.EnbInterfRandMod,
			data[j].ATTRIBUTES.TxdDci1aSwitch,
			data[j].ATTRIBUTES.DlSchAbnUeThd,
			data[j].ATTRIBUTES.RarAndPagingCR,
			data[j].ATTRIBUTES.CqiAdjInitialStep,
			data[j].ATTRIBUTES.CsiRsSfSchStrSwitch,
			data[j].ATTRIBUTES.SfnDlLoadPeriod,
			data[j].ATTRIBUTES.FreqSelJudgePeriod,
			data[j].ATTRIBUTES.DlEnhancedVoipSchSw,
			data[j].ATTRIBUTES.FDUEEnhAperCQITrigPeriod,
			data[j].ATTRIBUTES.SfnDlHighLoadThd,
			data[j].ATTRIBUTES.SfnDlLowLoadThd,
			data[j].ATTRIBUTES.DlHighLoadSdmaThdOffset,
			data[j].ATTRIBUTES.LbtSwitch,
			data[j].ATTRIBUTES.LbtNiHighThreshold,
			data[j].ATTRIBUTES.LbtNiLowThreshold,
			data[j].ATTRIBUTES.EnAperiodicCqiTrigStrategy,
			data[j].ATTRIBUTES.DlCaUeCapAllocStrategy,
			data[j].ATTRIBUTES.DlCaUserSchSwitch,
			data[j].ATTRIBUTES.Dl256QamCqiTblAdpPeriod,
			data[j].ATTRIBUTES.Dl256QamCqiTblCfgStrategy,
			data[j].ATTRIBUTES.BenefitMcsThreshold,
			data[j].ATTRIBUTES.DlSpsMcsDecreaseIblerThd,
			data[j].ATTRIBUTES.DlAfcMaxFreq,
			data[j].ATTRIBUTES.SrsRsrpFilterCoefForDlAfc,
			data[j].ATTRIBUTES.RlcMacJointSchSw,
			data[j].ATTRIBUTES.NoSchStopACqiThd,
			data[j].ATTRIBUTES.InterPoleRruIdThld,
			data[j].ATTRIBUTES.IntraPoleRruIdDuration,
			data[j].ATTRIBUTES.IntraPoleRRUIdPeriod,
			data[j].ATTRIBUTES.NgbRruIdThld,
			data[j].ATTRIBUTES.UserSpeedDlSchPriWeight,
			data[j].ATTRIBUTES.EmimoCellRbRatio,
			data[j].ATTRIBUTES.EmimoJointSchThd,
			data[j].ATTRIBUTES.EmimoIndependentSchThd,
			data[j].ATTRIBUTES.EmimoJointIndepThdOffset,
			data[j].ATTRIBUTES.StartRankDetectEffThd,
			data[j].ATTRIBUTES.StartRankDetectCntThd,
			data[j].ATTRIBUTES.RankDetectSuccessCntThd,
			data[j].ATTRIBUTES.RankReverseDetectCntThd,
			data[j].ATTRIBUTES.AmbrCtrlTcycle,
			data[j].ATTRIBUTES.CaSchWeight,
			data[j].ATTRIBUTES.PosNegFoRatioThd,
			data[j].ATTRIBUTES.UserFoThdForDlAfc,
			data[j].ATTRIBUTES.ZoneUserNumRatioThd,
			data[j].ATTRIBUTES.CrsFullBwPostODTtiNum,
			data[j].ATTRIBUTES.CrsFullBwPreODTtiNum,
			data[j].ATTRIBUTES.CrsFullBwPreSibPagingTtiNum,
			data[j].ATTRIBUTES.UeAttJudgePeriod,
			data[j].ATTRIBUTES.UeAttJudgeRsrpHyst,
			data[j].ATTRIBUTES.BPUeOdAlignThd,
			data[j].ATTRIBUTES.PrbEnableThldVideoResCtrl,
			data[j].ATTRIBUTES.IblerPdcchSinrOffset,
			data[j].ATTRIBUTES.FSUEAperCQITrigPeriod,
			data[j].ATTRIBUTES.FSUESbCQIValidityPeriod,
			data[j].ATTRIBUTES.InterfRandPciOffset,
			data[j].ATTRIBUTES.PilotOffStrategy,
			data[j].ATTRIBUTES.CrsCompensatePeriod,
			data[j].ATTRIBUTES.DlFirstHarqTxTbsIncNum,
			data[j].ATTRIBUTES.NbDlHarqMaxTxCount,
			data[j].ATTRIBUTES.CongestMaxVideoRate,
			data[j].ATTRIBUTES.UserEnableThldVideoResCtrl,
			data[j].ATTRIBUTES.PilotOffDrxParaGroupId,
			data[j].ATTRIBUTES.MccsCqiPeriodThld,
			data[j].ATTRIBUTES.MccsPrbRateThld,
			data[j].ATTRIBUTES.MultiCarrierCoSchAlgoSw,
			data[j].ATTRIBUTES.OverlapRsrpIsolationThd,
			data[j].ATTRIBUTES.CbtcDlSchedulingMcsUpLimit,
			data[j].ATTRIBUTES.DlEnhMuCellRbRatioThld,
			data[j].ATTRIBUTES.DlEnhMuCellAvgCQIThd,
			data[j].ATTRIBUTES.DlEnhMuPrbPairRatioThld,
			data[j].ATTRIBUTES.DlEnhMuOnOffTimer,
			data[j].ATTRIBUTES.HighIblerTargetTbsIdxThld,
			data[j].ATTRIBUTES.LowIblerTargetTbsIdxThld,
			data[j].ATTRIBUTES.DlCandPairUeCntThld,
			data[j].ATTRIBUTES.PdschMaxCodeRate,
			data[j].ATTRIBUTES.IndependentBeamPmiThld,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldlschalgo data has been saved")
	}
	tx.Commit()
}


func insertCelldrxpara(eNodeBId string, baseName string, data []model.Celldrxpara) {
	fmt.Println("[+] Processing Celldrxpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldrxpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.FddEnterDrxThd,
			data[j].ATTRIBUTES.FddExitDrxThd,
			data[j].ATTRIBUTES.TddEnterDrxThdUl,
			data[j].ATTRIBUTES.TddExitDrxThdUl,
			data[j].ATTRIBUTES.TddEnterDrxThdDl,
			data[j].ATTRIBUTES.TddExitDrxThdDl,
			data[j].ATTRIBUTES.DataAmountStatTimer,
			data[j].ATTRIBUTES.TddPowerSaveMeasSwitch,
			data[j].ATTRIBUTES.TddPowerSaveStatTimer,
			data[j].ATTRIBUTES.TddPowerSavingEnterDrxThd,
			data[j].ATTRIBUTES.TddPowerSavingExitDrxThd,
			data[j].ATTRIBUTES.LongDrxCycleUnsync,
			data[j].ATTRIBUTES.CqiMask,
			data[j].ATTRIBUTES.OndurationTimerUnsync,
			data[j].ATTRIBUTES.DrxInactivityTimerUnsync,
			data[j].ATTRIBUTES.DrxPolicyMode,
			data[j].ATTRIBUTES.DrxStartOffsetOptSwitch,
			data[j].ATTRIBUTES.DrxRcvDtxProSwitch,
			data[j].ATTRIBUTES.DrxForMeasSwitch,
			data[j].ATTRIBUTES.LongDrxCycleForMeas,
			data[j].ATTRIBUTES.OnDurTimerForMeas,
			data[j].ATTRIBUTES.DrxInactTimerForMeas,
			data[j].ATTRIBUTES.DrxReTxTimerForMeas,
			data[j].ATTRIBUTES.ShortDrxSwForMeas,
			data[j].ATTRIBUTES.ShortDrxCycleForMeas,
			data[j].ATTRIBUTES.ShortCycleTimerForMeas,
			data[j].ATTRIBUTES.DrxSrDetectOptSwitch,
			data[j].ATTRIBUTES.DrxStartoffsetAdjustSW,
			data[j].ATTRIBUTES.MeasDrxSpecSchSwitch,
			data[j].ATTRIBUTES.CovGsmMeasDrxCfgSwitch,
			data[j].ATTRIBUTES.OnDurationUnextendSwitch,
			data[j].ATTRIBUTES.DrxStateDuringUlHarqRetx,
			data[j].ATTRIBUTES.DrxAlgSwitch,
			data[j].ATTRIBUTES.ShortDrxCycleSwitch,
			data[j].ATTRIBUTES.VolteGapDrxExclusiveSwitch,
			data[j].ATTRIBUTES.DrxStopSrPendingSw,
			data[j].ATTRIBUTES.GapDrxExclusiveSwitch,
			data[j].ATTRIBUTES.NbDrxInactivityTimer,
			data[j].ATTRIBUTES.NbDrxReTxTimer,
			data[j].ATTRIBUTES.NbDrxUlReTxTimer,
			data[j].ATTRIBUTES.NbLongDrxCycle,
			data[j].ATTRIBUTES.NbOnDurationTimer,
			data[j].ATTRIBUTES.DrxOdPreSchSwitch,
			data[j].ATTRIBUTES.DrxUeSrsOptSwitch,
			data[j].ATTRIBUTES.SinrThldForVolteDrxCtrl,
			data[j].ATTRIBUTES.VoltePlrThldForExitingDrx,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldrxpara data has been saved")
	}
	tx.Commit()
}


func insertCelldrxspecialpara(eNodeBId string, baseName string, data []model.Celldrxspecialpara) {
	fmt.Println("[+] Processing Celldrxspecialpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldrxspecialpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellDrxSpecialParaValid,
			data[j].ATTRIBUTES.LongDrxCycleSpecial,
			data[j].ATTRIBUTES.OnDurationTimerSpecial,
			data[j].ATTRIBUTES.DrxInactivityTimerSpecial,
			data[j].ATTRIBUTES.ShortDrxCycleSwitchSpecial,
			data[j].ATTRIBUTES.LongDrxCycleForIntraRatAnr,
			data[j].ATTRIBUTES.LongDrxCycleForInterRatAnr,
			data[j].ATTRIBUTES.FddAnrDrxInactivityTimer,
			data[j].ATTRIBUTES.TddAnrDrxInactivityTimer,
			data[j].ATTRIBUTES.BfSpecificDrxParaSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldrxspecialpara data has been saved")
	}
	tx.Commit()
}


func insertCelldss(eNodeBId string, baseName string, data []model.Celldss) {
	fmt.Println("[+] Processing Celldss data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldss` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.HighFreqShareRbNum,
			data[j].ATTRIBUTES.LowFreqShareRbNum,
			data[j].ATTRIBUTES.ReMuteSwitch,
			data[j].ATTRIBUTES.UlInterfRestrictionMode,
			data[j].ATTRIBUTES.A3Offset,
			data[j].ATTRIBUTES.A6Offset,
			data[j].ATTRIBUTES.NearAreaSinrThd,
			data[j].ATTRIBUTES.MiddleAreaSinrThd,
			data[j].ATTRIBUTES.FarAreaSinrThd,
			data[j].ATTRIBUTES.InterfNCellConfigMode,
			data[j].ATTRIBUTES.SpecShrPfmOptSwitch,
			data[j].ATTRIBUTES.SinrThdWithoutGsmInterf,
			data[j].ATTRIBUTES.GsmInterfINThd,
			data[j].ATTRIBUTES.NearPtUserRsrpThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldss data has been saved")
	}
	tx.Commit()
}


func insertCelldynacbaralgopara(eNodeBId string, baseName string, data []model.Celldynacbaralgopara) {
	fmt.Println("[+] Processing Celldynacbaralgopara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celldynacbaralgopara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DynAcBarStatPeriod,
			data[j].ATTRIBUTES.DynAcBarTriggerThd,
			data[j].ATTRIBUTES.DynAcBarCancelThd,
			data[j].ATTRIBUTES.SsacTriggerCondSatiPeriods,
			data[j].ATTRIBUTES.SsacCancelCondSatiPeriods,
			data[j].ATTRIBUTES.DisasterReferenceInd,
			data[j].ATTRIBUTES.DisasterDuration,
			data[j].ATTRIBUTES.MoTriggerCondSatiPeriods,
			data[j].ATTRIBUTES.MoCancelCondSatiPeriods,
			data[j].ATTRIBUTES.MoFactorAdjStep,
			data[j].ATTRIBUTES.MoFactorRetreatStep,
			data[j].ATTRIBUTES.SsacFactorAdjStep,
			data[j].ATTRIBUTES.SsacFactorRetreatStep,	

		)
		checkErr(err)
		//fmt.Println("[+] Celldynacbaralgopara data has been saved")
	}
	tx.Commit()
}


func insertCelleabalgopara(eNodeBId string, baseName string, data []model.Celleabalgopara) {
	fmt.Println("[+] Processing Celleabalgopara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celleabalgopara` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.EABTriggerThd,
			data[j].ATTRIBUTES.EABStatPeriod,
			data[j].ATTRIBUTES.EABCategory,
			data[j].ATTRIBUTES.EABCancelThd,
			data[j].ATTRIBUTES.EABCancelCondSatiPeriod,
			data[j].ATTRIBUTES.ABForExceptionData,
			data[j].ATTRIBUTES.ABForSpecialAC,	

		)
		checkErr(err)
		//fmt.Println("[+] Celleabalgopara data has been saved")
	}
	tx.Commit()
}


func insertCelleicic(eNodeBId string, baseName string, data []model.Celleicic) {
	fmt.Println("[+] Processing Celleicic data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celleicic` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.AbsPattern,
			data[j].ATTRIBUTES.EicicMeasureEnable,
			data[j].ATTRIBUTES.A3Offset,
			data[j].ATTRIBUTES.AbsAdjPeriod,
			data[j].ATTRIBUTES.SampleNumTarget,
			data[j].ATTRIBUTES.AttachCellAddThd,
			data[j].ATTRIBUTES.WorkMode,
			data[j].ATTRIBUTES.StatPeriod,
			data[j].ATTRIBUTES.CreAdjPeriod,
			data[j].ATTRIBUTES.EicicConfigUserRatioThd,
			data[j].ATTRIBUTES.EicicUnConfigUserRatioThd,
			data[j].ATTRIBUTES.LargeCreMicroNumThd,
			data[j].ATTRIBUTES.EicicAdaptiveSwitch,
			data[j].ATTRIBUTES.NaicNCellStatPeriod,
			data[j].ATTRIBUTES.HighCoorCCEThd,
			data[j].ATTRIBUTES.HighCoorDTXThd,
			data[j].ATTRIBUTES.HighSpeedEnhancedOptSwitch,
			data[j].ATTRIBUTES.HSHandinUserThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Celleicic data has been saved")
	}
	tx.Commit()
}


func insertCellemtcalgo(eNodeBId string, baseName string, data []model.Cellemtcalgo) {
	fmt.Println("[+] Processing Cellemtcalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellemtcalgo` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.EmtcUlRbTargetRatio,
			data[j].ATTRIBUTES.EmtcAperCqiTrigPrd,
			data[j].ATTRIBUTES.EmtcDlRbTargetRatio,
			data[j].ATTRIBUTES.EmtcAlgoSwitch,
			data[j].ATTRIBUTES.DlLteRsvNbCount,
			data[j].ATTRIBUTES.UlLteRsvNbCount,
			data[j].ATTRIBUTES.EmtcVolteSupportSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellemtcalgo data has been saved")
	}
	tx.Commit()
}


func insertCellhoparacfg(eNodeBId string, baseName string, data []model.Cellhoparacfg) {
	fmt.Println("[+] Processing Cellhoparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellhoparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BlindHoA1A2ThdRsrp,
			data[j].ATTRIBUTES.BlindHoA1A2ThdRsrq,
			data[j].ATTRIBUTES.HighSpeedThreshold,
			data[j].ATTRIBUTES.HoModeSwitch,
			data[j].ATTRIBUTES.SrvccHoOptSwitch,
			data[j].ATTRIBUTES.UlBadQualHoMcsThd,
			data[j].ATTRIBUTES.UlBadQualHoIblerThd,
			data[j].ATTRIBUTES.SpeedEvaluatedPeriod,
			data[j].ATTRIBUTES.SpeedEvaluatedNum,
			data[j].ATTRIBUTES.L2UCsfbMRProMode,
			data[j].ATTRIBUTES.CsfbMRWaitingTimer,
			data[j].ATTRIBUTES.CellHoAlgoSwitch,
			data[j].ATTRIBUTES.SpeedEvaluatedNumForFastUe,
			data[j].ATTRIBUTES.HoUseInactiveTimerSwitch,
			data[j].ATTRIBUTES.HSCellSleepUserNum,
			data[j].ATTRIBUTES.LowSpeedUsersSelProTimer,
			data[j].ATTRIBUTES.HSCellHoInProtectTimer,
			data[j].ATTRIBUTES.FlashSrvccSwitch,
			data[j].ATTRIBUTES.UlPoorCoverPathLossThd,
			data[j].ATTRIBUTES.UlPoorCoverSinrThd,
			data[j].ATTRIBUTES.HoMrDelayTimerQci1,
			data[j].ATTRIBUTES.VoLTEQualityHoAlgoSwitch,
			data[j].ATTRIBUTES.Qci1PlrThldForVolteQualHo,
			data[j].ATTRIBUTES.VolteQualPktLossPeriod,
			data[j].ATTRIBUTES.CovBasedIntraRatMeasTime,
			data[j].ATTRIBUTES.HighSpeedUeJudgeMode,
			data[j].ATTRIBUTES.RingingMsgCheckSw,
			data[j].ATTRIBUTES.SrvccMrDelayTimer,
			data[j].ATTRIBUTES.VolteQualIfHoQci1PlrThld,
			data[j].ATTRIBUTES.VolteQualRecoveryQci1PlrThld,
			data[j].ATTRIBUTES.InterRatUuHoFailRetryTimes,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellhoparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellidprdupt(eNodeBId string, baseName string, data []model.Cellidprdupt) {
	fmt.Println("[+] Processing Cellidprdupt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellidprdupt` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PrdUptSwitch,
			data[j].ATTRIBUTES.ActionTime,
			data[j].ATTRIBUTES.Period,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellidprdupt data has been saved")
	}
	tx.Commit()
}


func insertCelliicspara(eNodeBId string, baseName string, data []model.Celliicspara) {
	fmt.Println("[+] Processing Celliicspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celliicspara` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.IicsUserAttrA3Offset,
			data[j].ATTRIBUTES.IicsUserAttrThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Celliicspara data has been saved")
	}
	tx.Commit()
}


func insertCelllowpower(eNodeBId string, baseName string, data []model.Celllowpower) {
	fmt.Println("[+] Processing Celllowpower data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celllowpower` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.LowPwrSwitch,
			data[j].ATTRIBUTES.CellUsedPwrReduceTimeLen,
			data[j].ATTRIBUTES.RsPwrReduceTimeLen,
			data[j].ATTRIBUTES.RfShutDownTimeLen,
			data[j].ATTRIBUTES.CellUsedPwrRatio,
			data[j].ATTRIBUTES.RsPwrAdjOffset,
			data[j].ATTRIBUTES.EnterTimeLen,
			data[j].ATTRIBUTES.BakPwrSavPolicy,	

		)
		checkErr(err)
		//fmt.Println("[+] Celllowpower data has been saved")
	}
	tx.Commit()
}


func insertCelllteflexbw(eNodeBId string, baseName string, data []model.Celllteflexbw) {
	fmt.Println("[+] Processing Celllteflexbw data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celllteflexbw` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.LteFlexBwSwitch,
			data[j].ATTRIBUTES.DlCustStartPrbIndex,
			data[j].ATTRIBUTES.DlCustEndPrbIndex,
			data[j].ATTRIBUTES.UlCustStartPrbIndex,
			data[j].ATTRIBUTES.UlCustEndPrbIndex,	

		)
		checkErr(err)
		//fmt.Println("[+] Celllteflexbw data has been saved")
	}
	tx.Commit()
}


func insertCelllteflexbwitfcfg(eNodeBId string, baseName string, data []model.Celllteflexbwitfcfg) {
	fmt.Println("[+] Processing Celllteflexbwitfcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celllteflexbwitfcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.HighItfGsmArfcn,
			data[j].ATTRIBUTES.GsmCarrierFreq,	

		)
		checkErr(err)
		//fmt.Println("[+] Celllteflexbwitfcfg data has been saved")
	}
	tx.Commit()
}


func insertCellmbfcs(eNodeBId string, baseName string, data []model.Cellmbfcs) {
	fmt.Println("[+] Processing Cellmbfcs data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmbfcs` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellSyntheticalRateHoHyst,
			data[j].ATTRIBUTES.CellTrafficLoadThd,
			data[j].ATTRIBUTES.InterFreqMeasLoadRatioThd,
			data[j].ATTRIBUTES.InterFreqMeasMcsThd,
			data[j].ATTRIBUTES.InterFreqMeasServiceThd,
			data[j].ATTRIBUTES.MeasInfoUpdInterval,
			data[j].ATTRIBUTES.UeSpectralEffHoHyst,
			data[j].ATTRIBUTES.LoadDiffRatioThld,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmbfcs data has been saved")
	}
	tx.Commit()
}


func insertCellmbmscfg(eNodeBId string, baseName string, data []model.Cellmbmscfg) {
	fmt.Println("[+] Processing Cellmbmscfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmbmscfg` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MBSFNSyncAreaId,
			data[j].ATTRIBUTES.MBMSSwitch,
			data[j].ATTRIBUTES.ServiceAreaIdList,
			data[j].ATTRIBUTES.MBMSServiceSwitch,
			data[j].ATTRIBUTES.NCellMbmsCfgSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmbmscfg data has been saved")
	}
	tx.Commit()
}


func insertCellmcpara(eNodeBId string, baseName string, data []model.Cellmcpara) {
	fmt.Println("[+] Processing Cellmcpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmcpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.A3Offset,
			data[j].ATTRIBUTES.Hysteresis,
			data[j].ATTRIBUTES.TimetoTrigger,
			data[j].ATTRIBUTES.MaxReportCells,
			data[j].ATTRIBUTES.ReportAmount,
			data[j].ATTRIBUTES.ReportInterval,
			data[j].ATTRIBUTES.ReportQuantity,
			data[j].ATTRIBUTES.TriggerQuantity,
			data[j].ATTRIBUTES.A6Offset,
			data[j].ATTRIBUTES.IntraFreqMaxReportCells,
			data[j].ATTRIBUTES.IntraFreqTriggerQuantity,
			data[j].ATTRIBUTES.IntraFreqReportQuantity,
			data[j].ATTRIBUTES.InterFreqMaxReportCells,
			data[j].ATTRIBUTES.InterFreqTriggerQuantity,
			data[j].ATTRIBUTES.InterFreqReportQuantity,
			data[j].ATTRIBUTES.A3MCAdaptiveSwitch,
			data[j].ATTRIBUTES.AoaDetectThd,
			data[j].ATTRIBUTES.EcidReportAmount,
			data[j].ATTRIBUTES.EcidReportInterval,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmcpara data has been saved")
	}
	tx.Commit()
}


func insertCellmimoparacfg(eNodeBId string, baseName string, data []model.Cellmimoparacfg) {
	fmt.Println("[+] Processing Cellmimoparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmimoparacfg` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MimoAdaptiveSwitch,
			data[j].ATTRIBUTES.FixedMimoMode,
			data[j].ATTRIBUTES.InitialMimoType,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmimoparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellmlb(eNodeBId string, baseName string, data []model.Cellmlb) {
	fmt.Println("[+] Processing Cellmlb data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmlb` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterFreqMlbThd,
			data[j].ATTRIBUTES.InterRatMlbThd,
			data[j].ATTRIBUTES.LoadOffset,
			data[j].ATTRIBUTES.LoadDiffThd,
			data[j].ATTRIBUTES.InterRatMlbUeNumThd,
			data[j].ATTRIBUTES.InitValidPeriod,
			data[j].ATTRIBUTES.LoadTransferFactor,
			data[j].ATTRIBUTES.MlbTriggerMode,
			data[j].ATTRIBUTES.InterFreqMlbUeNumThd,
			data[j].ATTRIBUTES.MlbUeNumOffset,
			data[j].ATTRIBUTES.MlbMaxUeNum,
			data[j].ATTRIBUTES.MlbUeSelectPrbThd,
			data[j].ATTRIBUTES.DlDataMlbMode,
			data[j].ATTRIBUTES.InterFreqMLBRanShareMode,
			data[j].ATTRIBUTES.UeNumDiffThd,
			data[j].ATTRIBUTES.HotSpotUeMode,
			data[j].ATTRIBUTES.InterFreqIdleMlbMode,
			data[j].ATTRIBUTES.MlbMinUeNumThd,
			data[j].ATTRIBUTES.MlbMinUeNumOffset,
			data[j].ATTRIBUTES.InterFreqUeTrsfType,
			data[j].ATTRIBUTES.InterFreqIdleMlbUeNumThd,
			data[j].ATTRIBUTES.InterRatIdleMlbUeNumThd,
			data[j].ATTRIBUTES.InterFreqLoadEvalPrd,
			data[j].ATTRIBUTES.InterRatLoadEvalPrd,
			data[j].ATTRIBUTES.FreqSelectStrategy,
			data[j].ATTRIBUTES.LoadBalanceNCellScope,
			data[j].ATTRIBUTES.InterRatMlbUeNumOffset,
			data[j].ATTRIBUTES.IdleUeSelFreqScope,
			data[j].ATTRIBUTES.InterRatMlbUeSelStrategy,
			data[j].ATTRIBUTES.InterRatMlbUeSelPrbThd,
			data[j].ATTRIBUTES.PrbLoadCalcMethod,
			data[j].ATTRIBUTES.MlbUeSelectPunishTimer,
			data[j].ATTRIBUTES.MlbHoCellSelectStrategy,
			data[j].ATTRIBUTES.InterRatMlbTriggerMode,
			data[j].ATTRIBUTES.InterRatMlbUeNumModeThd,
			data[j].ATTRIBUTES.PunishJudgePrdNum,
			data[j].ATTRIBUTES.FreqPunishPrdNum,
			data[j].ATTRIBUTES.CellPunishPrdNum,
			data[j].ATTRIBUTES.MultiRruMode,
			data[j].ATTRIBUTES.InterFreqOffloadOffset,
			data[j].ATTRIBUTES.InterFrqUeNumOffloadOffset,
			data[j].ATTRIBUTES.CellCapacityScaleFactor,
			data[j].ATTRIBUTES.InterRatMlbMaxUeNum,
			data[j].ATTRIBUTES.InterRatMlbHoFailPunish,
			data[j].ATTRIBUTES.EutranLoadJudgeStrategy,
			data[j].ATTRIBUTES.MlbTrigJudgePeriod,
			data[j].ATTRIBUTES.InterFreqMlbStrategy,
			data[j].ATTRIBUTES.MaxSpectralEfficiencyValue,
			data[j].ATTRIBUTES.MinSpectralEfficiencyValue,
			data[j].ATTRIBUTES.SpectralEffAdjustMaxStep,
			data[j].ATTRIBUTES.UeNumDiffOffsetTransCa,
			data[j].ATTRIBUTES.MlbIdleUeNumAdjFactor,
			data[j].ATTRIBUTES.IdleMlbUeNumDiffThd,
			data[j].ATTRIBUTES.L2USmartOffloadThd,
			data[j].ATTRIBUTES.L2USmartOffloadOffset,
			data[j].ATTRIBUTES.InterFreqMlbUlThd,
			data[j].ATTRIBUTES.UeDlPrbLowThdOffset,
			data[j].ATTRIBUTES.UeUlPrbHighThdOffset,
			data[j].ATTRIBUTES.UeUlPrbLowThdOffset,
			data[j].ATTRIBUTES.VideoLoadHighThd,
			data[j].ATTRIBUTES.VideoLoadLowThd,
			data[j].ATTRIBUTES.VideoDlPrbThd,
			data[j].ATTRIBUTES.InterFIdleUeNumOffloadOfs,
			data[j].ATTRIBUTES.UlExperienceDiffThd,
			data[j].ATTRIBUTES.UlExperienceEvalPrd,
			data[j].ATTRIBUTES.UlExperienceMaxUeNum,
			data[j].ATTRIBUTES.UlExperienceOffloadThd,
			data[j].ATTRIBUTES.UlExperienceOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmlb data has been saved")
	}
	tx.Commit()
}


func insertCellmlbautogroup(eNodeBId string, baseName string, data []model.Cellmlbautogroup) {
	fmt.Println("[+] Processing Cellmlbautogroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmlbautogroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.OverlapIndAddThd,
			data[j].ATTRIBUTES.OverlapIndDelThd,
			data[j].ATTRIBUTES.StatPeriod,
			data[j].ATTRIBUTES.SleepPeriod,
			data[j].ATTRIBUTES.AddedMeasCfg,
			data[j].ATTRIBUTES.OverlapIndAutoSampleNum,
			data[j].ATTRIBUTES.OptMode,
			data[j].ATTRIBUTES.MicroOverlapIndAddThd,
			data[j].ATTRIBUTES.MicroOverlapIndDelThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmlbautogroup data has been saved")
	}
	tx.Commit()
}


func insertCellmlbho(eNodeBId string, baseName string, data []model.Cellmlbho) {
	fmt.Println("[+] Processing Cellmlbho data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmlbho` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.IdleUeSelFreqStrategy,
			data[j].ATTRIBUTES.MlbHoInProtectMode,
			data[j].ATTRIBUTES.MlbHoInProtectTimer,
			data[j].ATTRIBUTES.InterFreqMlbHoInA1ThdRsrp,
			data[j].ATTRIBUTES.InterFreqMlbHoInA1ThdRsrq,
			data[j].ATTRIBUTES.InterFreqMlbHoInA2ThdRsrp,
			data[j].ATTRIBUTES.InterFreqMlbHoInA2ThdRsrq,
			data[j].ATTRIBUTES.InterRatMlbStrategy,
			data[j].ATTRIBUTES.MlbMatchOtherFeatureMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmlbho data has been saved")
	}
	tx.Commit()
}


func insertCellmlbuesel(eNodeBId string, baseName string, data []model.Cellmlbuesel) {
	fmt.Println("[+] Processing Cellmlbuesel data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmlbuesel` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UeSelectPrbPrio,
			data[j].ATTRIBUTES.UeSelectQciPrio,
			data[j].ATTRIBUTES.UeSelectArpPrio,
			data[j].ATTRIBUTES.UeSelectDlMcsPrio,
			data[j].ATTRIBUTES.InterFreqMlbUeArpThd,
			data[j].ATTRIBUTES.InterFreqMlbUeDlMcsThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmlbuesel data has been saved")
	}
	tx.Commit()
}


func insertCellmmalgo(eNodeBId string, baseName string, data []model.Cellmmalgo) {
	fmt.Println("[+] Processing Cellmmalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmmalgo` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BeamGainLimitSwitch,
			data[j].ATTRIBUTES.MMAlgoOptSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmmalgo data has been saved")
	}
	tx.Commit()
}


func insertCellmro(eNodeBId string, baseName string, data []model.Cellmro) {
	fmt.Println("[+] Processing Cellmro data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellmro` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CioAdjLimitCfgInd,
			data[j].ATTRIBUTES.CioAdjUpperLimit,
			data[j].ATTRIBUTES.CioAdjLowerLimit,
			data[j].ATTRIBUTES.InterFreqA2RsrpUpLimit,
			data[j].ATTRIBUTES.InterFreqA2RsrpLowLimit,
			data[j].ATTRIBUTES.A3InterFreqA2RsrpUpLimit,
			data[j].ATTRIBUTES.A3InterFreqA2RsrpLowLimit,
			data[j].ATTRIBUTES.InterFreqMroAdjParaSel,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellmro data has been saved")
	}
	tx.Commit()
}


func insertCellnoaccessalmpara(eNodeBId string, baseName string, data []model.Cellnoaccessalmpara) {
	fmt.Println("[+] Processing Cellnoaccessalmpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellnoaccessalmpara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.NoAccessStartDetectTime,
			data[j].ATTRIBUTES.NoAccessStopDetectTime,
			data[j].ATTRIBUTES.NoAccessDetectTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellnoaccessalmpara data has been saved")
	}
	tx.Commit()
}


func insertCellop(eNodeBId string, baseName string, data []model.Cellop) {
	fmt.Println("[+] Processing Cellop data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellop` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TrackingAreaId,
			data[j].ATTRIBUTES.CellReservedForOp,
			data[j].ATTRIBUTES.OpUlRbUsedRatio,
			data[j].ATTRIBUTES.OpDlRbUsedRatio,
			data[j].ATTRIBUTES.MMECfgNum,
			data[j].ATTRIBUTES.OpUeNumRatio,
			data[j].ATTRIBUTES.OpPttUlRbHighThd,
			data[j].ATTRIBUTES.OpPttUlRbLowThd,
			data[j].ATTRIBUTES.OpPttDlRbHighThd,
			data[j].ATTRIBUTES.OpPttDlRbLowThd,
			data[j].ATTRIBUTES.OperatorPlmnRoundPeriod,
			data[j].ATTRIBUTES.OpResourceGroupIndex,
			data[j].ATTRIBUTES.OpNonGbrResourceRatio,
			data[j].ATTRIBUTES.RatFreqPriorityGroupId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellop data has been saved")
	}
	tx.Commit()
}


func insertCellpcalgo(eNodeBId string, baseName string, data []model.Cellpcalgo) {
	fmt.Println("[+] Processing Cellpcalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellpcalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SrsPcStrategy,
			data[j].ATTRIBUTES.PucchCloseLoopPcType,
			data[j].ATTRIBUTES.PucchPcPeriod,
			data[j].ATTRIBUTES.PucchPcTargetSinrOffset,
			data[j].ATTRIBUTES.PuschRsrpHighThd,
			data[j].ATTRIBUTES.PuschIoTCtrlA3Offset,
			data[j].ATTRIBUTES.IoTNearPointOptSwitch,
			data[j].ATTRIBUTES.IoTNearPointPLThreshold,
			data[j].ATTRIBUTES.SrsPcSinrTarget,
			data[j].ATTRIBUTES.SrsPcRsrpTarget,
			data[j].ATTRIBUTES.PUSCHPsdCtrlTarget,
			data[j].ATTRIBUTES.CloseLoopOptPUSCHType,
			data[j].ATTRIBUTES.PUSCHOptIBLERThreshold,
			data[j].ATTRIBUTES.PUSCHPsdCtrlTargetForUs,
			data[j].ATTRIBUTES.IoTCtrlINCorrectSwitch,
			data[j].ATTRIBUTES.IoTCtrlEUPLThreshold,
			data[j].ATTRIBUTES.IoTCtrlNIThreshold,
			data[j].ATTRIBUTES.NearPointUePuschType,
			data[j].ATTRIBUTES.VoLteSinrHighThdOffset,
			data[j].ATTRIBUTES.VoltePUSCHPowerOffset,
			data[j].ATTRIBUTES.PuschRsrpHighThdOffsetVoIP,
			data[j].ATTRIBUTES.DMSrsPcSinrOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellpcalgo data has been saved")
	}
	tx.Commit()
}


func insertCellpdcchalgo(eNodeBId string, baseName string, data []model.Cellpdcchalgo) {
	fmt.Println("[+] Processing Cellpdcchalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellpdcchalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ComSigCongregLv,
			data[j].ATTRIBUTES.CceRatioAdjSwitch,
			data[j].ATTRIBUTES.SfnPdcchDcsThd,
			data[j].ATTRIBUTES.InitPdcchSymNum,
			data[j].ATTRIBUTES.VirtualLoadPro,
			data[j].ATTRIBUTES.PdcchInitialCceAdjustValue,
			data[j].ATTRIBUTES.PdcchSymNumSwitch,
			data[j].ATTRIBUTES.CceUseRatio,
			data[j].ATTRIBUTES.PdcchAggLvlCLAdjustSwitch,
			data[j].ATTRIBUTES.DPDVirtualLoadSwitch,
			data[j].ATTRIBUTES.DPDVirtualLoadType,
			data[j].ATTRIBUTES.AggLvlSelStrageForDualCW,
			data[j].ATTRIBUTES.PdcchCapacityImproveSwitch,
			data[j].ATTRIBUTES.PdcchMaxCodeRate,
			data[j].ATTRIBUTES.ULDLPdcchSymNum,
			data[j].ATTRIBUTES.PDCCHAggLvlAdaptStrage,
			data[j].ATTRIBUTES.HysForCfiBasedPreSch,
			data[j].ATTRIBUTES.SfnPdcchSdmaThd,
			data[j].ATTRIBUTES.UlPdcchAllocImproveSwitch,
			data[j].ATTRIBUTES.CceMaxInitialRatio,
			data[j].ATTRIBUTES.PdcchPowerEnhancedSwitch,
			data[j].ATTRIBUTES.PdcchBlerTarget,
			data[j].ATTRIBUTES.HLNetAccSigAggLvlSelEnhSw,
			data[j].ATTRIBUTES.EpdcchAlgoSwitch,
			data[j].ATTRIBUTES.CceThdtoEnableEpdcch,
			data[j].ATTRIBUTES.RbThdtoEnableEpdcch,
			data[j].ATTRIBUTES.CceThdtoDisableEpdcch,
			data[j].ATTRIBUTES.EcceThdtoDisableEpdcch,
			data[j].ATTRIBUTES.RbThdtoDisableEpdcch,
			data[j].ATTRIBUTES.HLNetAccSigAttempt,
			data[j].ATTRIBUTES.EpdcchSfCfg,
			data[j].ATTRIBUTES.UlDlEcceInitialRatioAdjSw,
			data[j].ATTRIBUTES.CapacityBfOpt,
			data[j].ATTRIBUTES.CCEThdEnableFlowCtrl,
			data[j].ATTRIBUTES.CCEThdDisableFlowCtrl,
			data[j].ATTRIBUTES.PDCCHPwrUpperLimitOptSw,
			data[j].ATTRIBUTES.VoltePdcchSinrOffset,
			data[j].ATTRIBUTES.PdcchSparePowerAllocStrage,
			data[j].ATTRIBUTES.PdcchDlAggLvlSlcEhnSwitch,
			data[j].ATTRIBUTES.PdcchOutLoopAdjBaseStep,
			data[j].ATTRIBUTES.PdcchOutLoopAdjLowerLimit,
			data[j].ATTRIBUTES.PdcchAdjAlgoSwitch,
			data[j].ATTRIBUTES.SplitBeamPdcchSdmaThd,
			data[j].ATTRIBUTES.NackDtxRatioThd,
			data[j].ATTRIBUTES.SplitBeamPdcchSdmaSw,
			data[j].ATTRIBUTES.PdcchPwrCtrlUserNumThd,
			data[j].ATTRIBUTES.PdcchBfGainOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellpdcchalgo data has been saved")
	}
	tx.Commit()
}


func insertCellpdcchcecfg(eNodeBId string, baseName string, data []model.Cellpdcchcecfg) {
	fmt.Println("[+] Processing Cellpdcchcecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellpdcchcecfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.PdcchMaxRepetitionCnt,
			data[j].ATTRIBUTES.PdcchPeriodFactor,
			data[j].ATTRIBUTES.PdcchTransRptCntFactor,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellpdcchcecfg data has been saved")
	}
	tx.Commit()
}


func insertCellprbvalmlb(eNodeBId string, baseName string, data []model.Cellprbvalmlb) {
	fmt.Println("[+] Processing Cellprbvalmlb data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellprbvalmlb` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PrbValMlbTrigThd,
			data[j].ATTRIBUTES.PrbValMlbAdmitThd,
			data[j].ATTRIBUTES.MlbUeSelectPrbValThd,
			data[j].ATTRIBUTES.PrbValFilterFactor,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellprbvalmlb data has been saved")
	}
	tx.Commit()
}


func insertCellpucchalgo(eNodeBId string, baseName string, data []model.Cellpucchalgo) {
	fmt.Println("[+] Processing Cellpucchalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellpucchalgo` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SriLowLoadThd,
			data[j].ATTRIBUTES.SriReCfgInd,
			data[j].ATTRIBUTES.SriAlgoSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellpucchalgo data has been saved")
	}
	tx.Commit()
}


func insertCellqcipara(eNodeBId string, baseName string, data []model.Cellqcipara) {
	fmt.Println("[+] Processing Cellqcipara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellqcipara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.InterFreqHoGroupId,
			data[j].ATTRIBUTES.InterRatHoCdma1xRttGroupId,
			data[j].ATTRIBUTES.InterRatHoCdmaHrpdGroupId,
			data[j].ATTRIBUTES.InterRatHoCommGroupId,
			data[j].ATTRIBUTES.InterRatHoGeranGroupId,
			data[j].ATTRIBUTES.InterRatHoUtranGroupId,
			data[j].ATTRIBUTES.IntraFreqHoGroupId,
			data[j].ATTRIBUTES.DrxParaGroupId,
			data[j].ATTRIBUTES.SriPeriod,
			data[j].ATTRIBUTES.RlfTimerConstCfgInd,
			data[j].ATTRIBUTES.TrafficRelDelay,
			data[j].ATTRIBUTES.QciPriorityForHo,
			data[j].ATTRIBUTES.PreallocationParaGroupId,
			data[j].ATTRIBUTES.QciPriorityForDrx,
			data[j].ATTRIBUTES.QciAlgoSwitch,
			data[j].ATTRIBUTES.QciEutranFreqRelationId,
			data[j].ATTRIBUTES.QciUtranFreqRelationId,
			data[j].ATTRIBUTES.QciGeranFreqRelationId,
			data[j].ATTRIBUTES.MlbQciGroupId,
			data[j].ATTRIBUTES.EmtcModeADrxParaGroupId,
			data[j].ATTRIBUTES.EmtcModeBDrxParaGroupId,
			data[j].ATTRIBUTES.ServiceFlag,
			data[j].ATTRIBUTES.DlAmbrLimitFactor,
			data[j].ATTRIBUTES.QciAmbrLimitFlag,
			data[j].ATTRIBUTES.UlAmbrLimitFactor,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellqcipara data has been saved")
	}
	tx.Commit()
}


func insertCellrachalgo(eNodeBId string, baseName string, data []model.Cellrachalgo) {
	fmt.Println("[+] Processing Cellrachalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellrachalgo` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PrachFalseAlarmDetRadThd,
			data[j].ATTRIBUTES.RachThdBoostRatio,
			data[j].ATTRIBUTES.PrachInterfPeriod,
			data[j].ATTRIBUTES.PrachInterfThreshold,
			data[j].ATTRIBUTES.PrachInterfHysteresis,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellrachalgo data has been saved")
	}
	tx.Commit()
}


func insertCellrachcecfg(eNodeBId string, baseName string, data []model.Cellrachcecfg) {
	fmt.Println("[+] Processing Cellrachcecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellrachcecfg` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.ContentionResolutionTimer,
			data[j].ATTRIBUTES.PrachTransmissionPeriod,
			data[j].ATTRIBUTES.PrachSubcarrierOffset,
			data[j].ATTRIBUTES.PrachRepetitionCount,
			data[j].ATTRIBUTES.MaxNumPreambleAttempt,
			data[j].ATTRIBUTES.PrachDetectionThld,
			data[j].ATTRIBUTES.PrachStartTime,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellrachcecfg data has been saved")
	}
	tx.Commit()
}


func insertCellracthd(eNodeBId string, baseName string, data []model.Cellracthd) {
	fmt.Println("[+] Processing Cellracthd data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellracthd` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.GoldServiceArpThd,
			data[j].ATTRIBUTES.SilverServiceArpThd,
			data[j].ATTRIBUTES.Qci1HoThd,
			data[j].ATTRIBUTES.Qci2HoThd,
			data[j].ATTRIBUTES.Qci3HoThd,
			data[j].ATTRIBUTES.Qci4HoThd,
			data[j].ATTRIBUTES.NewGoldServiceOffset,
			data[j].ATTRIBUTES.NewSilverServiceOffset,
			data[j].ATTRIBUTES.NewCopperServiceOffset,
			data[j].ATTRIBUTES.Qci1CongThd,
			data[j].ATTRIBUTES.Qci2CongThd,
			data[j].ATTRIBUTES.Qci3CongThd,
			data[j].ATTRIBUTES.Qci4CongThd,
			data[j].ATTRIBUTES.CongRelOffset,
			data[j].ATTRIBUTES.UlRbHighThd,
			data[j].ATTRIBUTES.UlRbLowThd,
			data[j].ATTRIBUTES.AcReservedUserNumber,
			data[j].ATTRIBUTES.VolteReservedNumber,
			data[j].ATTRIBUTES.VoltePrefAdmissionTimer,
			data[j].ATTRIBUTES.CellCapacityMode,
			data[j].ATTRIBUTES.LoadHoAdmitOffset,
			data[j].ATTRIBUTES.VoipOverAdmitOffset,
			data[j].ATTRIBUTES.AcUserNumber,
			data[j].ATTRIBUTES.MoSigArpOverride,
			data[j].ATTRIBUTES.UlSatisThdforVolteLoadAmrc,
			data[j].ATTRIBUTES.CceUsageThd,
			data[j].ATTRIBUTES.PreResNeedTuningFactor,
			data[j].ATTRIBUTES.CceAlFailHighThd,
			data[j].ATTRIBUTES.CqiFarThd,
			data[j].ATTRIBUTES.DlExperienceThd,
			data[j].ATTRIBUTES.RbCongHighThd,
			data[j].ATTRIBUTES.UlExperienceThd,
			data[j].ATTRIBUTES.VolteArpOverride,
			data[j].ATTRIBUTES.CceThdforVolteLoadAmrc,
			data[j].ATTRIBUTES.UlRbThdforVolteLoadAmrc,
			data[j].ATTRIBUTES.CongBearRelPeriod,
			data[j].ATTRIBUTES.PreemptableBearerNum,
			data[j].ATTRIBUTES.Qci65HoThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellracthd data has been saved")
	}
	tx.Commit()
}


func insertCellresel(eNodeBId string, baseName string, data []model.Cellresel) {
	fmt.Println("[+] Processing Cellresel data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellresel` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Qhyst,
			data[j].ATTRIBUTES.SpeedDepReselCfgInd,
			data[j].ATTRIBUTES.SNonIntraSearchCfgInd,
			data[j].ATTRIBUTES.SNonIntraSearch,
			data[j].ATTRIBUTES.ThrshServLow,
			data[j].ATTRIBUTES.CellReselPriority,
			data[j].ATTRIBUTES.QRxLevMin,
			data[j].ATTRIBUTES.PMaxCfgInd,
			data[j].ATTRIBUTES.SIntraSearchCfgInd,
			data[j].ATTRIBUTES.SIntraSearch,
			data[j].ATTRIBUTES.MeasBandWidthCfgInd,
			data[j].ATTRIBUTES.TReselEutran,
			data[j].ATTRIBUTES.SpeedStateSfCfgInd,
			data[j].ATTRIBUTES.TReselEutranSfMedium,
			data[j].ATTRIBUTES.TReselEutranSfHigh,
			data[j].ATTRIBUTES.NeighCellConfig,
			data[j].ATTRIBUTES.PresenceAntennaPort1,
			data[j].ATTRIBUTES.SIntraSearchQ,
			data[j].ATTRIBUTES.SNonIntraSearchQ,
			data[j].ATTRIBUTES.ThrshServLowQCfgInd,
			data[j].ATTRIBUTES.QQualMinCfgInd,
			data[j].ATTRIBUTES.QQualMin,
			data[j].ATTRIBUTES.TReselForNb,
			data[j].ATTRIBUTES.TReselInterFreqForNb,
			data[j].ATTRIBUTES.TReselEutranCE,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellresel data has been saved")
	}
	tx.Commit()
}


func insertCellreselgeran(eNodeBId string, baseName string, data []model.Cellreselgeran) {
	fmt.Println("[+] Processing Cellreselgeran data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellreselgeran` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TReselGeran,
			data[j].ATTRIBUTES.SpeedStateSfCfgInd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellreselgeran data has been saved")
	}
	tx.Commit()
}


func insertCellreselutran(eNodeBId string, baseName string, data []model.Cellreselutran) {
	fmt.Println("[+] Processing Cellreselutran data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellreselutran` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TReselUtran,
			data[j].ATTRIBUTES.SpeedStateSfCfgInd,
			data[j].ATTRIBUTES.TReselUtranSfMedium,
			data[j].ATTRIBUTES.TReselUtranSfHigh,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellreselutran data has been saved")
	}
	tx.Commit()
}


func insertCellrfshutdown(eNodeBId string, baseName string, data []model.Cellrfshutdown) {
	fmt.Println("[+] Processing Cellrfshutdown data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellrfshutdown` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RfShutdownSwitch,
			data[j].ATTRIBUTES.StartTime,
			data[j].ATTRIBUTES.StopTime,
			data[j].ATTRIBUTES.RsPwrAdjOffset,
			data[j].ATTRIBUTES.DlPrbThd,
			data[j].ATTRIBUTES.UlPrbThd,
			data[j].ATTRIBUTES.DlPrbOffset,
			data[j].ATTRIBUTES.UlPrbOffset,
			data[j].ATTRIBUTES.UENumThd,
			data[j].ATTRIBUTES.RfShutdownJudgingPolicy,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellrfshutdown data has been saved")
	}
	tx.Commit()
}


func insertCellricalgo(eNodeBId string, baseName string, data []model.Cellricalgo) {
	fmt.Println("[+] Processing Cellricalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellricalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RicAlgoSwitch,
			data[j].ATTRIBUTES.MuteUpPTSSymNum,
			data[j].ATTRIBUTES.MuteULSym,
			data[j].ATTRIBUTES.UlInterferenceThd,
			data[j].ATTRIBUTES.DuctingRemoteInfThd,
			data[j].ATTRIBUTES.InfAvoidDetPeriodNum,
			data[j].ATTRIBUTES.InfAvoidRecDetPeriodNum,
			data[j].ATTRIBUTES.RemoteInfAdpAvoidSwitch,
			data[j].ATTRIBUTES.UlInterferenceSymbMaxDiff,
			data[j].ATTRIBUTES.DuctDLSubfrmShutoffSwitch,
			data[j].ATTRIBUTES.UlInterferenceDuration,
			data[j].ATTRIBUTES.DuctDLSubfrmShutoffInfThld,
			data[j].ATTRIBUTES.DuctDLSubfrmShutoffOptSw,
			data[j].ATTRIBUTES.DuctInfDetPeriodCount,
			data[j].ATTRIBUTES.DuctInfPwrDiffInThld,
			data[j].ATTRIBUTES.DuctInfPwrDiffOutThld,
			data[j].ATTRIBUTES.DuctInfRateThld,
			data[j].ATTRIBUTES.DuctSubfrmShutoffDepth,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellricalgo data has been saved")
	}
	tx.Commit()
}


func insertCellsel(eNodeBId string, baseName string, data []model.Cellsel) {
	fmt.Println("[+] Processing Cellsel data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellsel` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.QRxLevMin,
			data[j].ATTRIBUTES.QRxLevMinOffset,
			data[j].ATTRIBUTES.QQualMin,
			data[j].ATTRIBUTES.QQualMinOffsetCfgInd,
			data[j].ATTRIBUTES.QRxLevMinCE,
			data[j].ATTRIBUTES.QQualMinCE,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellsel data has been saved")
	}
	tx.Commit()
}


func insertCellservicediffcfg(eNodeBId string, baseName string, data []model.Cellservicediffcfg) {
	fmt.Println("[+] Processing Cellservicediffcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellservicediffcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ServiceDiffSwitch,
			data[j].ATTRIBUTES.CellConStatePrbThd,
			data[j].ATTRIBUTES.CellOverloadStatePrbThd,
			data[j].ATTRIBUTES.CellConStateUeNumThd,
			data[j].ATTRIBUTES.CellOverloadStateUeNumThd,
			data[j].ATTRIBUTES.QueueServiceWeight0,
			data[j].ATTRIBUTES.QueueServiceWeight1,
			data[j].ATTRIBUTES.QueueServiceWeight2,
			data[j].ATTRIBUTES.QueueServiceWeight3,
			data[j].ATTRIBUTES.QueueServiceWeight4,
			data[j].ATTRIBUTES.QueueServiceWeight5,
			data[j].ATTRIBUTES.QueueServiceWeight6,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellservicediffcfg data has been saved")
	}
	tx.Commit()
}


func insertCellshutdown(eNodeBId string, baseName string, data []model.Cellshutdown) {
	fmt.Println("[+] Processing Cellshutdown data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellshutdown` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellShutdownSwitch,
			data[j].ATTRIBUTES.StartTime,
			data[j].ATTRIBUTES.StopTime,
			data[j].ATTRIBUTES.DlPrbThd,
			data[j].ATTRIBUTES.DlPrbOffset,
			data[j].ATTRIBUTES.UlPrbThd,
			data[j].ATTRIBUTES.UlPrbOffset,
			data[j].ATTRIBUTES.ForceShutdownUENumThd,
			data[j].ATTRIBUTES.PunishTime,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellshutdown data has been saved")
	}
	tx.Commit()
}


func insertCellsimap(eNodeBId string, baseName string, data []model.Cellsimap) {
	fmt.Println("[+] Processing Cellsimap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellsimap` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SiMapSwitch,
			data[j].ATTRIBUTES.Sib2Period,
			data[j].ATTRIBUTES.Sib3Period,
			data[j].ATTRIBUTES.Sib4Period,
			data[j].ATTRIBUTES.Sib5Period,
			data[j].ATTRIBUTES.Sib6Period,
			data[j].ATTRIBUTES.Sib7Period,
			data[j].ATTRIBUTES.Sib8Period,
			data[j].ATTRIBUTES.Sib10Period,
			data[j].ATTRIBUTES.Sib11Period,
			data[j].ATTRIBUTES.EtwsPnDuration,
			data[j].ATTRIBUTES.EtwsSnOverlapPolicy,
			data[j].ATTRIBUTES.EtwsPnOverlapPolicy,
			data[j].ATTRIBUTES.Sib12Period,
			data[j].ATTRIBUTES.SiTransEcr,
			data[j].ATTRIBUTES.Sib13Period,
			data[j].ATTRIBUTES.Sib15Period,
			data[j].ATTRIBUTES.Sib16Period,
			data[j].ATTRIBUTES.SibTransCtrlSwitch,
			data[j].ATTRIBUTES.SiSwitch,
			data[j].ATTRIBUTES.SiSchResRatio,
			data[j].ATTRIBUTES.SibUpdOptSwitch,
			data[j].ATTRIBUTES.Sib14Period,
			data[j].ATTRIBUTES.NbSib1RepetitionNum,
			data[j].ATTRIBUTES.NbSib2Period,
			data[j].ATTRIBUTES.NbSib3Period,
			data[j].ATTRIBUTES.NbSib4Period,
			data[j].ATTRIBUTES.NbSib5Period,
			data[j].ATTRIBUTES.NbSib14Period,
			data[j].ATTRIBUTES.NbSib16Period,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellsimap data has been saved")
	}
	tx.Commit()
}


func insertCellsrlte(eNodeBId string, baseName string, data []model.Cellsrlte) {
	fmt.Println("[+] Processing Cellsrlte data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellsrlte` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SrlteSwitch,
			data[j].ATTRIBUTES.SrlteDtxThrd,
			data[j].ATTRIBUTES.SrlteSuspendTime,
			data[j].ATTRIBUTES.SrlteDiscardAlgoSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellsrlte data has been saved")
	}
	tx.Commit()
}


func insertCellsrsadaptivecfg(eNodeBId string, baseName string, data []model.Cellsrsadaptivecfg) {
	fmt.Println("[+] Processing Cellsrsadaptivecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellsrsadaptivecfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SrsPeriodAdaptive,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellsrsadaptivecfg data has been saved")
	}
	tx.Commit()
}


func insertCellstandardqci(eNodeBId string, baseName string, data []model.Cellstandardqci) {
	fmt.Println("[+] Processing Cellstandardqci data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellstandardqci` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.InterFreqHoGroupId,
			data[j].ATTRIBUTES.InterRatHoCdma1xRttGroupId,
			data[j].ATTRIBUTES.InterRatHoCdmaHrpdGroupId,
			data[j].ATTRIBUTES.InterRatHoCommGroupId,
			data[j].ATTRIBUTES.InterRatHoGeranGroupId,
			data[j].ATTRIBUTES.InterRatHoUtranGroupId,
			data[j].ATTRIBUTES.IntraFreqHoGroupId,
			data[j].ATTRIBUTES.DrxParaGroupId,
			data[j].ATTRIBUTES.SriPeriod,
			data[j].ATTRIBUTES.RlfTimerConstCfgInd,
			data[j].ATTRIBUTES.RlfTimerConstGroupId,
			data[j].ATTRIBUTES.TrafficRelDelay,
			data[j].ATTRIBUTES.QciPriorityForHo,
			data[j].ATTRIBUTES.MlbQciGroupId,
			data[j].ATTRIBUTES.PreallocationParaGroupId,
			data[j].ATTRIBUTES.QciPriorityForDrx,
			data[j].ATTRIBUTES.QciAlgoSwitch,
			data[j].ATTRIBUTES.QciEutranFreqRelationId,
			data[j].ATTRIBUTES.QciUtranFreqRelationId,
			data[j].ATTRIBUTES.QciGeranFreqRelationId,
			data[j].ATTRIBUTES.QciAmbrLimitFlag,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellstandardqci data has been saved")
	}
	tx.Commit()
}


func insertCellttibundlingalgo(eNodeBId string, baseName string, data []model.Cellttibundlingalgo) {
	fmt.Println("[+] Processing Cellttibundlingalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellttibundlingalgo` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SinrThdToTrigTtib,
			data[j].ATTRIBUTES.SinrThdToTrigVideoTtib,
			data[j].ATTRIBUTES.TtiBundlingAlgoSw,
			data[j].ATTRIBUTES.R12TtiBHarqMaxTxNum,
			data[j].ATTRIBUTES.R12TtiBundlingSwitch,
			data[j].ATTRIBUTES.SinrThdToTrigR12TtiB,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellttibundlingalgo data has been saved")
	}
	tx.Commit()
}


func insertCellucionpuschpara(eNodeBId string, baseName string, data []model.Cellucionpuschpara) {
	fmt.Println("[+] Processing Cellucionpuschpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellucionpuschpara` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CellUciOnPuschParaValid,
			data[j].ATTRIBUTES.DeltaOffsetCqiIndex,
			data[j].ATTRIBUTES.DeltaOffsetRiIndex,
			data[j].ATTRIBUTES.DeltaOffsetAckIndex,
			data[j].ATTRIBUTES.DeltaOffsetAckIndexForTtiB,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellucionpuschpara data has been saved")
	}
	tx.Commit()
}


func insertCelluemeascontrolcfg(eNodeBId string, baseName string, data []model.Celluemeascontrolcfg) {
	fmt.Println("[+] Processing Celluemeascontrolcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `celluemeascontrolcfg` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MaxNonIntraMeasObjNum,
			data[j].ATTRIBUTES.MaxEutranFddMeasFreqNum,
			data[j].ATTRIBUTES.MaxEutranTddMeasFreqNum,
			data[j].ATTRIBUTES.MaxUtranFddMeasFreqNum,
			data[j].ATTRIBUTES.MaxUtranTddMeasFreqNum,
			data[j].ATTRIBUTES.MaxGeranMeasFreqNum,
			data[j].ATTRIBUTES.MaxUeTddMeasFreqNum,	

		)
		checkErr(err)
		//fmt.Println("[+] Celluemeascontrolcfg data has been saved")
	}
	tx.Commit()
}


func insertCellulcompalgo(eNodeBId string, baseName string, data []model.Cellulcompalgo) {
	fmt.Println("[+] Processing Cellulcompalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulcompalgo` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlCompA3Offset,
			data[j].ATTRIBUTES.SfnUlCompThd,
			data[j].ATTRIBUTES.UlCompA3OffsetForRelaxedBH,
			data[j].ATTRIBUTES.UlCompUlRsrpOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulcompalgo data has been saved")
	}
	tx.Commit()
}


func insertCellulicalgo(eNodeBId string, baseName string, data []model.Cellulicalgo) {
	fmt.Println("[+] Processing Cellulicalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulicalgo` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlIcA3Offset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulicalgo data has been saved")
	}
	tx.Commit()
}


func insertCellulicic(eNodeBId string, baseName string, data []model.Cellulicic) {
	fmt.Println("[+] Processing Cellulicic data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulicic` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BandMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulicic data has been saved")
	}
	tx.Commit()
}


func insertCellulicicmcpara(eNodeBId string, baseName string, data []model.Cellulicicmcpara) {
	fmt.Println("[+] Processing Cellulicicmcpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulicicmcpara` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.A3Offset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulicicmcpara data has been saved")
	}
	tx.Commit()
}


func insertCellulmimoparacfg(eNodeBId string, baseName string, data []model.Cellulmimoparacfg) {
	fmt.Println("[+] Processing Cellulmimoparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulmimoparacfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlSuMimoRankPara,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulmimoparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellulpccomm(eNodeBId string, baseName string, data []model.Cellulpccomm) {
	fmt.Println("[+] Processing Cellulpccomm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulpccomm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PassLossCoeff,
			data[j].ATTRIBUTES.P0NominalPUCCH,
			data[j].ATTRIBUTES.P0NominalPUSCH,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat1,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat1b,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat2,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat2a,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat2b,
			data[j].ATTRIBUTES.DeltaPreambleMsg3,
			data[j].ATTRIBUTES.DeltaMsg2,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat1bcs,
			data[j].ATTRIBUTES.DeltaFPUCCHFormat3,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulpccomm data has been saved")
	}
	tx.Commit()
}


func insertCellulpcdedic(eNodeBId string, baseName string, data []model.Cellulpcdedic) {
	fmt.Println("[+] Processing Cellulpcdedic data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulpcdedic` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DeltaMcsEnabled,
			data[j].ATTRIBUTES.PSrsOffsetDeltaMcsDisable,
			data[j].ATTRIBUTES.PSrsOffsetDeltaMcsEnable,
			data[j].ATTRIBUTES.FilterRsrp,
			data[j].ATTRIBUTES.PathlossReferenceLink,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulpcdedic data has been saved")
	}
	tx.Commit()
}


func insertCellulschalgo(eNodeBId string, baseName string, data []model.Cellulschalgo) {
	fmt.Println("[+] Processing Cellulschalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulschalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlschStrategy,
			data[j].ATTRIBUTES.AdaptHarqSwitch,
			data[j].ATTRIBUTES.SinrAdjustTargetIbler,
			data[j].ATTRIBUTES.PreAllocationBandwidthRatio,
			data[j].ATTRIBUTES.PreAllocationMinPeriod,
			data[j].ATTRIBUTES.PreAllocationSize,
			data[j].ATTRIBUTES.UlHoppingType,
			data[j].ATTRIBUTES.FreeUserUlRbUsedRatio,
			data[j].ATTRIBUTES.UlSrSchDateLen,
			data[j].ATTRIBUTES.SpsRelThd,
			data[j].ATTRIBUTES.SmartPreAllocationDuration,
			data[j].ATTRIBUTES.UlEpfCapacityFactor,
			data[j].ATTRIBUTES.UlRbAllocationStrategy,
			data[j].ATTRIBUTES.DopMeasLevel,
			data[j].ATTRIBUTES.UlHarqMaxTxNum,
			data[j].ATTRIBUTES.SmartPreAllocDuraForSparse,
			data[j].ATTRIBUTES.UlSpsInterval,
			data[j].ATTRIBUTES.SriFalseDetThdSwitch,
			data[j].ATTRIBUTES.NoUlSchTtiNumAffterGap,
			data[j].ATTRIBUTES.UlDelaySchStrategy,
			data[j].ATTRIBUTES.UlSchAbnUeThd,
			data[j].ATTRIBUTES.SpecificPktSizeThd,
			data[j].ATTRIBUTES.SrMaskSwitch,
			data[j].ATTRIBUTES.PuschDtxSchStrategy,
			data[j].ATTRIBUTES.UlVoipRlcMaxSegNum,
			data[j].ATTRIBUTES.UlEnhencedVoipSchSw,
			data[j].ATTRIBUTES.UlInBasedFssSinrThld,
			data[j].ATTRIBUTES.UlCamcDlRsrpOffset,
			data[j].ATTRIBUTES.StatisticNumThdForTtibTrig,
			data[j].ATTRIBUTES.StatisticNumThdForTtibExit,
			data[j].ATTRIBUTES.HystToExitTtiBundling,
			data[j].ATTRIBUTES.TtiBundlingRlcMaxSegNum,
			data[j].ATTRIBUTES.TtiBundlingHarqMaxTxNum,
			data[j].ATTRIBUTES.TtiBundlingTriggerStrategy,
			data[j].ATTRIBUTES.DopAlgoSwitch,
			data[j].ATTRIBUTES.EnhancedVmimoSwitch,
			data[j].ATTRIBUTES.UeNumThdInPdcchPuschBal,
			data[j].ATTRIBUTES.DataThdInPdcchPuschBal,
			data[j].ATTRIBUTES.HeadOverheadForUlSch,
			data[j].ATTRIBUTES.PreAllocMinPeriodForSparse,
			data[j].ATTRIBUTES.PreallocationSizeForSparse,
			data[j].ATTRIBUTES.UlInterfRandomMode,
			data[j].ATTRIBUTES.SinrAdjTargetIblerforVoLTE,
			data[j].ATTRIBUTES.SfnUlLoadPeriod,
			data[j].ATTRIBUTES.MaxLayerHOVMIMO,
			data[j].ATTRIBUTES.UlCompenSchPeriodinSpurt,
			data[j].ATTRIBUTES.UlCompenSchPeriodinSilence,
			data[j].ATTRIBUTES.UlTargetIBlerAdaptType,
			data[j].ATTRIBUTES.AperiodicCsiUlTxMode,
			data[j].ATTRIBUTES.UlVoipRsvRbStart,
			data[j].ATTRIBUTES.UlVoipRsvRbNum,
			data[j].ATTRIBUTES.UlIBlerAdaptBigTrafficSw,
			data[j].ATTRIBUTES.VmimoOptAlgoSwitch,
			data[j].ATTRIBUTES.UlCamcInterfTh,
			data[j].ATTRIBUTES.UlIcsA3Offset,
			data[j].ATTRIBUTES.UlMcsLowThreshold,
			data[j].ATTRIBUTES.UserSpeedUlSchPriWeight,
			data[j].ATTRIBUTES.SinrAdjTargetIblerforPtt,
			data[j].ATTRIBUTES.TarRruSelRsrpOffsetThd,
			data[j].ATTRIBUTES.MaxUlSchRbNum,
			data[j].ATTRIBUTES.UlExtVolteSchSw,
			data[j].ATTRIBUTES.UlVolteDeltaSinrForNack,
			data[j].ATTRIBUTES.InitUlSinrAdjust,
			data[j].ATTRIBUTES.UlSinrAdjustStep,
			data[j].ATTRIBUTES.UlSinrFilterCoef,
			data[j].ATTRIBUTES.SfnUlPairRsrpThd,
			data[j].ATTRIBUTES.UlSrSchDataVolAdptOptUpThd,
			data[j].ATTRIBUTES.TtiBundlingRetxStrategy,
			data[j].ATTRIBUTES.UlVoLTERetransSchStrategy,
			data[j].ATTRIBUTES.NbUlHarqMaxTxCount,
			data[j].ATTRIBUTES.PrealloSystemBwCoeff,
			data[j].ATTRIBUTES.SmartPrealloDurationCoeff,
			data[j].ATTRIBUTES.HighSpeedSdmaIsolationThld,
			data[j].ATTRIBUTES.VmimoPairingStrategy,
			data[j].ATTRIBUTES.VMIMOEgdeResRatio,
			data[j].ATTRIBUTES.MaxLayerMMVMIMO,
			data[j].ATTRIBUTES.UlSrSchDateLenForVoLTE,
			data[j].ATTRIBUTES.AmrcDecreasingPeriod,
			data[j].ATTRIBUTES.SinrThdForVoLTERateCtrl,
			data[j].ATTRIBUTES.RateCtrlCmrProcessStrategy,
			data[j].ATTRIBUTES.SinrAdjTargetIblerforQCI2,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulschalgo data has been saved")
	}
	tx.Commit()
}


func insertCellulunifiedolc(eNodeBId string, baseName string, data []model.Cellulunifiedolc) {
	fmt.Println("[+] Processing Cellulunifiedolc data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellulunifiedolc` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RrcRejectRateHighThd,
			data[j].ATTRIBUTES.RrcRejectRateLowThd,
			data[j].ATTRIBUTES.RrcReqNumHighThd,
			data[j].ATTRIBUTES.RrcReqNumLowThd,
			data[j].ATTRIBUTES.UeNumHighThd,
			data[j].ATTRIBUTES.UeNumLowThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellulunifiedolc data has been saved")
	}
	tx.Commit()
}


func insertCellusparacfg(eNodeBId string, baseName string, data []model.Cellusparacfg) {
	fmt.Println("[+] Processing Cellusparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellusparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UsAlgoSwitch,
			data[j].ATTRIBUTES.UsVoIPPreAllocMinPeriod,
			data[j].ATTRIBUTES.UsDataPreAllocMinPeriod,
			data[j].ATTRIBUTES.UsVoIPPreAllocSize,
			data[j].ATTRIBUTES.UsDataPreAllocSize,
			data[j].ATTRIBUTES.UsVoIPSmartPreallocDura,
			data[j].ATTRIBUTES.UsDataSmartPreallocDura,
			data[j].ATTRIBUTES.UlUsVoIPIblerTarget,
			data[j].ATTRIBUTES.NsNumThdForUsHo,
			data[j].ATTRIBUTES.UsPucchSinrTargetOffset,
			data[j].ATTRIBUTES.UsPucchRsrpHighThdOffset,
			data[j].ATTRIBUTES.UsDataRatFreqPriGroupId,
			data[j].ATTRIBUTES.UsGuaranteeDurTimer,
			data[j].ATTRIBUTES.UsGapMeasurementPeriod,
			data[j].ATTRIBUTES.UsHoA2ThdRsrp,
			data[j].ATTRIBUTES.UsHoA2ThdRsrq,
			data[j].ATTRIBUTES.UsVoIPRatFreqPriGroupId,
			data[j].ATTRIBUTES.NsRatFreqPriGroupId,
			data[j].ATTRIBUTES.UsPaPcOff,
			data[j].ATTRIBUTES.UsPuschSinrHighThdOffset,
			data[j].ATTRIBUTES.NsRbRestrictRatio,
			data[j].ATTRIBUTES.UsDetectTimer,
			data[j].ATTRIBUTES.UsDlMinGbr,
			data[j].ATTRIBUTES.UsUlMinGbr,
			data[j].ATTRIBUTES.UsNbInterfCtrl,
			data[j].ATTRIBUTES.UlUsVoipRsvStartRb,
			data[j].ATTRIBUTES.UlUsVoipRsvEndRb,
			data[j].ATTRIBUTES.SfCtrlPrbThd,
			data[j].ATTRIBUTES.UsVoIPKeepOnLen,
			data[j].ATTRIBUTES.NsIdleRatFreqPriGroupId,
			data[j].ATTRIBUTES.UsDlPriorityRatio,
			data[j].ATTRIBUTES.UsUlPriorityRatio,
			data[j].ATTRIBUTES.UsA3HoA2ThdOffset,
			data[j].ATTRIBUTES.UsA4A5HoA2ThdOffset,
			data[j].ATTRIBUTES.UlUsDataRatFreqPriGroupId,
			data[j].ATTRIBUTES.UsNbUlSinrHighThdOffset,
			data[j].ATTRIBUTES.UsNbUlRbRestrictRat,
			data[j].ATTRIBUTES.UsA3HoA1ThdOffset,
			data[j].ATTRIBUTES.UsA4A5HoA1ThdOffset,
			data[j].ATTRIBUTES.UsDataPdcchSinrOffset,
			data[j].ATTRIBUTES.UsVoipPdcchSinrOffset,
			data[j].ATTRIBUTES.UsVoipInitDlIblerTarget,
			data[j].ATTRIBUTES.UsVoipCompenSchPeriod,
			data[j].ATTRIBUTES.CrsShutDownRsrpOffset,
			data[j].ATTRIBUTES.CrsShutDownStrategy,
			data[j].ATTRIBUTES.VolteSilDelayHoRsrpThld,
			data[j].ATTRIBUTES.UsCellSchOptSwitch,
			data[j].ATTRIBUTES.MMPowerRatio,
			data[j].ATTRIBUTES.NsDlFirstRetxRbRatio,
			data[j].ATTRIBUTES.NsMaxNumAllowSchPerTTI,
			data[j].ATTRIBUTES.NsNotAllowSchSubframe,
			data[j].ATTRIBUTES.SchGuaranteeUsNum,
			data[j].ATTRIBUTES.UsWithGapMode,
			data[j].ATTRIBUTES.LargePacketDlDataThld,
			data[j].ATTRIBUTES.UsNotSchRbgRsrpOffset,
			data[j].ATTRIBUTES.UsNotSchRbgRsrpThld,
			data[j].ATTRIBUTES.UsAlgoExSwitch,
			data[j].ATTRIBUTES.UsSrsGuaranteeSwitch,
			data[j].ATTRIBUTES.UeMobilIdentifyStrategy,
			data[j].ATTRIBUTES.UsSmoothHoCompThldOfs,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellusparacfg data has been saved")
	}
	tx.Commit()
}


func insertCellvms(eNodeBId string, baseName string, data []model.Cellvms) {
	fmt.Println("[+] Processing Cellvms data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellvms` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.VmsHoUeNumTh,
			data[j].ATTRIBUTES.VmsPrbDiffTh,
			data[j].ATTRIBUTES.VmsPrbLoadTh,
			data[j].ATTRIBUTES.VmsA3Offset,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellvms data has been saved")
	}
	tx.Commit()
}


func insertCellwttxparacfg(eNodeBId string, baseName string, data []model.Cellwttxparacfg) {
	fmt.Println("[+] Processing Cellwttxparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cellwttxparacfg` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MbbUserDlPrbUpLimit,
			data[j].ATTRIBUTES.MbbUserUlPrbUpLimit,
			data[j].ATTRIBUTES.PrbUpLimitCtrlMode,
			data[j].ATTRIBUTES.WbbOrMbbUserDefMode,
			data[j].ATTRIBUTES.WbbUserDlPrbUpLimit,
			data[j].ATTRIBUTES.WbbUserUlPrbUpLimit,	

		)
		checkErr(err)
		//fmt.Println("[+] Cellwttxparacfg data has been saved")
	}
	tx.Commit()
}


func insertCepucchcfg(eNodeBId string, baseName string, data []model.Cepucchcfg) {
	fmt.Println("[+] Processing Cepucchcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cepucchcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.PucchRepNum,	

		)
		checkErr(err)
		//fmt.Println("[+] Cepucchcfg data has been saved")
	}
	tx.Commit()
}


func insertCerachcfg(eNodeBId string, baseName string, data []model.Cerachcfg) {
	fmt.Println("[+] Processing Cerachcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cerachcfg` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.ContentionResolutionTimer,
			data[j].ATTRIBUTES.MaxNumPrbAttempt,
			data[j].ATTRIBUTES.PreambleRatio,
			data[j].ATTRIBUTES.PreambleRepetitionNum,
			data[j].ATTRIBUTES.RandomPreambleRatio,	

		)
		checkErr(err)
		//fmt.Println("[+] Cerachcfg data has been saved")
	}
	tx.Commit()
}


func insertCertcfg(eNodeBId string, baseName string, data []model.Certcfg) {
	fmt.Println("[+] Processing Certcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `certcfg` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IKECHECKSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Certcfg data has been saved")
	}
	tx.Commit()
}


func insertCertchktsk(eNodeBId string, baseName string, data []model.Certchktsk) {
	fmt.Println("[+] Processing Certchktsk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `certchktsk` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ISENABLE,
			data[j].ATTRIBUTES.PERIOD,
			data[j].ATTRIBUTES.ALMRNG,
			data[j].ATTRIBUTES.UPDATEMETHOD,	

		)
		checkErr(err)
		//fmt.Println("[+] Certchktsk data has been saved")
	}
	tx.Commit()
}


func insertCertdeploy(eNodeBId string, baseName string, data []model.Certdeploy) {
	fmt.Println("[+] Processing Certdeploy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `certdeploy` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DEPLOYTYPE,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,	

		)
		checkErr(err)
		//fmt.Println("[+] Certdeploy data has been saved")
	}
	tx.Commit()
}


func insertCertmk(eNodeBId string, baseName string, data []model.Certmk) {
	fmt.Println("[+] Processing Certmk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `certmk` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.APPCERT,
			data[j].ATTRIBUTES.CASW,	

		)
		checkErr(err)
		//fmt.Println("[+] Certmk data has been saved")
	}
	tx.Commit()
}


func insertCertreq(eNodeBId string, baseName string, data []model.Certreq) {
	fmt.Println("[+] Processing Certreq data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `certreq` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.COMMNAME,
			data[j].ATTRIBUTES.USERADDINFO,
			data[j].ATTRIBUTES.COUNTRY,
			data[j].ATTRIBUTES.ORG,
			data[j].ATTRIBUTES.ORGUNIT,
			data[j].ATTRIBUTES.STATEPROVINCENAME,
			data[j].ATTRIBUTES.LOCALITY,
			data[j].ATTRIBUTES.KEYUSAGE,
			data[j].ATTRIBUTES.SIGNALG,
			data[j].ATTRIBUTES.KEYSIZE,
			data[j].ATTRIBUTES.LOCALNAME,
			data[j].ATTRIBUTES.LOCALIP,	

		)
		checkErr(err)
		//fmt.Println("[+] Certreq data has been saved")
	}
	tx.Commit()
}


func insertChk(eNodeBId string, baseName string, data []model.Chk) {
	fmt.Println("[+] Processing Chk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `chk` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENABLEFLAG,	

		)
		checkErr(err)
		//fmt.Println("[+] Chk data has been saved")
	}
	tx.Commit()
}


func insertClkdetect(eNodeBId string, baseName string, data []model.Clkdetect) {
	fmt.Println("[+] Processing Clkdetect data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `clkdetect` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ClkAsyncDetectSwitch,
			data[j].ATTRIBUTES.ClkAsyncInterfRptThld,
			data[j].ATTRIBUTES.ClkAsyncPrbUsageThld,
			data[j].ATTRIBUTES.ClkAsyncSilDetInfDifThld,
			data[j].ATTRIBUTES.ClkAsyncSilDetInfThld,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Clkdetect data has been saved")
	}
	tx.Commit()
}


func insertClzerobufferzone(eNodeBId string, baseName string, data []model.Clzerobufferzone) {
	fmt.Println("[+] Processing Clzerobufferzone data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `clzerobufferzone` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ClZeroBufzoneDlSharedInd,
			data[j].ATTRIBUTES.ClZeroBufzoneUlSharedInd,
			data[j].ATTRIBUTES.ClSharedFreqStartRb1,
			data[j].ATTRIBUTES.ClSharedFreqEndRb1,
			data[j].ATTRIBUTES.ClSharedFreqStartRb2,
			data[j].ATTRIBUTES.ClSharedFreqEndRb2,
			data[j].ATTRIBUTES.ClZeroBufferZoneUlPrbThd,
			data[j].ATTRIBUTES.ClZeroBufferZoneUlPrbOst,
			data[j].ATTRIBUTES.UlNearPtUserRsrpThd,
			data[j].ATTRIBUTES.DlNearPtUserRsrpThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Clzerobufferzone data has been saved")
	}
	tx.Commit()
}


func insertCnoperator(eNodeBId string, baseName string, data []model.Cnoperator) {
	fmt.Println("[+] Processing Cnoperator data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cnoperator` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.CnOperatorName,
			data[j].ATTRIBUTES.CnOperatorType,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.PlmnMode,
			data[j].ATTRIBUTES.OperatorFunSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Cnoperator data has been saved")
	}
	tx.Commit()
}


func insertCnoperatorhocfg(eNodeBId string, baseName string, data []model.Cnoperatorhocfg) {
	fmt.Println("[+] Processing Cnoperatorhocfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cnoperatorhocfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.FirstRatPri,
			data[j].ATTRIBUTES.SecondRatPri,
			data[j].ATTRIBUTES.TddIfHoA2ThdRsrpOffset,
			data[j].ATTRIBUTES.FddIfHoA2ThdRsrpOffset,
			data[j].ATTRIBUTES.UtranA2ThdRsrpOffset,
			data[j].ATTRIBUTES.GeranA2ThdRsrpOffset,
			data[j].ATTRIBUTES.SrvccWithPsBasedCellCapSw,
			data[j].ATTRIBUTES.PsInterRatSecondPri,
			data[j].ATTRIBUTES.PsInterRatHighestPri,
			data[j].ATTRIBUTES.PsInterRatLowestPri,	

		)
		checkErr(err)
		//fmt.Println("[+] Cnoperatorhocfg data has been saved")
	}
	tx.Commit()
}


func insertCnoperatorqcipara(eNodeBId string, baseName string, data []model.Cnoperatorqcipara) {
	fmt.Println("[+] Processing Cnoperatorqcipara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cnoperatorqcipara` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.ServiceIrHoCfgGroupId,
			data[j].ATTRIBUTES.ServiceIfHoCfgGroupId,
			data[j].ATTRIBUTES.ServiceHoBearerPolicy,
			data[j].ATTRIBUTES.LocalQciArp,	

		)
		checkErr(err)
		//fmt.Println("[+] Cnoperatorqcipara data has been saved")
	}
	tx.Commit()
}


func insertCnoperatorstandardqci(eNodeBId string, baseName string, data []model.Cnoperatorstandardqci) {
	fmt.Println("[+] Processing Cnoperatorstandardqci data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cnoperatorstandardqci` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.ServiceIrHoCfgGroupId,
			data[j].ATTRIBUTES.ServiceIfHoCfgGroupId,
			data[j].ATTRIBUTES.ServiceHoBearerPolicy,
			data[j].ATTRIBUTES.LocalQciArp,	

		)
		checkErr(err)
		//fmt.Println("[+] Cnoperatorstandardqci data has been saved")
	}
	tx.Commit()
}


func insertCnoperatorta(eNodeBId string, baseName string, data []model.Cnoperatorta) {
	fmt.Println("[+] Processing Cnoperatorta data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cnoperatorta` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TrackingAreaId,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.Tac,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.NbIotTaFlag,	

		)
		checkErr(err)
		//fmt.Println("[+] Cnoperatorta data has been saved")
	}
	tx.Commit()
}


func insertCountercheckpara(eNodeBId string, baseName string, data []model.Countercheckpara) {
	fmt.Println("[+] Processing Countercheckpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `countercheckpara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CounterCheckTimer,
			data[j].ATTRIBUTES.CounterCheckCountNum,
			data[j].ATTRIBUTES.CounterCheckUserRelSwitch,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Countercheckpara data has been saved")
	}
	tx.Commit()
}


func insertCpbearer(eNodeBId string, baseName string, data []model.Cpbearer) {
	fmt.Println("[+] Processing Cpbearer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cpbearer` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CPBEARID,
			data[j].ATTRIBUTES.FLAG,
			data[j].ATTRIBUTES.BEARTYPE,
			data[j].ATTRIBUTES.LINKNO,
			data[j].ATTRIBUTES.CTRLMODE,
			data[j].ATTRIBUTES.AUTOCFGFLAG,	

		)
		checkErr(err)
		//fmt.Println("[+] Cpbearer data has been saved")
	}
	tx.Commit()
}


func insertCpriport(eNodeBId string, baseName string, data []model.Cpriport) {
	fmt.Println("[+] Processing Cpriport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cpriport` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.OPTN,
			data[j].ATTRIBUTES.ADMINISTRATIVESTATE,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.SPN,	

		)
		checkErr(err)
		//fmt.Println("[+] Cpriport data has been saved")
	}
	tx.Commit()
}


func insertCpswitch(eNodeBId string, baseName string, data []model.Cpswitch) {
	fmt.Println("[+] Processing Cpswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cpswitch` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ES,	

		)
		checkErr(err)
		//fmt.Println("[+] Cpswitch data has been saved")
	}
	tx.Commit()
}


func insertCqiadaptivecfg(eNodeBId string, baseName string, data []model.Cqiadaptivecfg) {
	fmt.Println("[+] Processing Cqiadaptivecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cqiadaptivecfg` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CqiPeriodAdaptive,
			data[j].ATTRIBUTES.SimulAckNackAndCqiSwitch,
			data[j].ATTRIBUTES.PucchPeriodicCqiOptSwitch,
			data[j].ATTRIBUTES.HoAperiodicCqiCfgSwitch,
			data[j].ATTRIBUTES.SimulAckNackAndCqiFmt3Sw,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cqiadaptivecfg data has been saved")
	}
	tx.Commit()
}


func insertCrlpolicy(eNodeBId string, baseName string, data []model.Crlpolicy) {
	fmt.Println("[+] Processing Crlpolicy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `crlpolicy` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CRLPOLICY,	

		)
		checkErr(err)
		//fmt.Println("[+] Crlpolicy data has been saved")
	}
	tx.Commit()
}


func insertCsfallbackblindhocfg(eNodeBId string, baseName string, data []model.Csfallbackblindhocfg) {
	fmt.Println("[+] Processing Csfallbackblindhocfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `csfallbackblindhocfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.InterRatHighestPri,
			data[j].ATTRIBUTES.InterRatSecondPri,
			data[j].ATTRIBUTES.InterRatLowestPri,
			data[j].ATTRIBUTES.UtranLcsCap,
			data[j].ATTRIBUTES.GeranLcsCap,
			data[j].ATTRIBUTES.CdmaLcsCap,
			data[j].ATTRIBUTES.IdleCsfbHighestPri,
			data[j].ATTRIBUTES.IdleCsfbSecondPri,
			data[j].ATTRIBUTES.IdleCsfbLowestPri,
			data[j].ATTRIBUTES.UtranCsfbBlindRedirRrSw,	

		)
		checkErr(err)
		//fmt.Println("[+] Csfallbackblindhocfg data has been saved")
	}
	tx.Commit()
}


func insertCsfallbackho(eNodeBId string, baseName string, data []model.Csfallbackho) {
	fmt.Println("[+] Processing Csfallbackho data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `csfallbackho` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CsfbHoUtranTimeToTrig,
			data[j].ATTRIBUTES.CsfbHoGeranTimeToTrig,
			data[j].ATTRIBUTES.CsfbHoCdmaTimeToTrig,
			data[j].ATTRIBUTES.CsfbHoUtranB1ThdRscp,
			data[j].ATTRIBUTES.CsfbHoUtranB1ThdEcn0,
			data[j].ATTRIBUTES.CsfbHoGeranB1Thd,
			data[j].ATTRIBUTES.CsfbHoCdmaB1ThdPs,
			data[j].ATTRIBUTES.BlindHoA1ThdRsrp,
			data[j].ATTRIBUTES.CsfbProtectionTimer,
			data[j].ATTRIBUTES.CsfbProtectTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Csfallbackho data has been saved")
	}
	tx.Commit()
}


func insertCsfallbackpolicycfg(eNodeBId string, baseName string, data []model.Csfallbackpolicycfg) {
	fmt.Println("[+] Processing Csfallbackpolicycfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `csfallbackpolicycfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CsfbHoPolicyCfg,
			data[j].ATTRIBUTES.IdleModeCsfbHoPolicyCfg,
			data[j].ATTRIBUTES.CsfbUserArpCfgSwitch,
			data[j].ATTRIBUTES.NormalCsfbUserArp,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Csfallbackpolicycfg data has been saved")
	}
	tx.Commit()
}


func insertCspcalgopara(eNodeBId string, baseName string, data []model.Cspcalgopara) {
	fmt.Println("[+] Processing Cspcalgopara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `cspcalgopara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CspcAlgoSwitch,
			data[j].ATTRIBUTES.CspcRsrpMeasMode,
			data[j].ATTRIBUTES.CspcClusterMode,
			data[j].ATTRIBUTES.CspcClusterPartPeriod,
			data[j].ATTRIBUTES.CspcComputeSwitch,
			data[j].ATTRIBUTES.CspcFullPowerSubframeRatio,
			data[j].ATTRIBUTES.CspcPowerConfigDelay,
			data[j].ATTRIBUTES.CspcScheduleUeSpec,
			data[j].ATTRIBUTES.CspcEnableDlPrbRatioThd,
			data[j].ATTRIBUTES.TddCspcAlgoSwitch,
			data[j].ATTRIBUTES.CspcCapacityFactor,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Cspcalgopara data has been saved")
	}
	tx.Commit()
}


func insertDevip(eNodeBId string, baseName string, data []model.Devip) {
	fmt.Println("[+] Processing Devip data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `devip` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.IP,
			data[j].ATTRIBUTES.MASK,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.CTRLMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Devip data has been saved")
	}
	tx.Commit()
}


func insertDhcprelayswitch(eNodeBId string, baseName string, data []model.Dhcprelayswitch) {
	fmt.Println("[+] Processing Dhcprelayswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `dhcprelayswitch` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ES,	

		)
		checkErr(err)
		//fmt.Println("[+] Dhcprelayswitch data has been saved")
	}
	tx.Commit()
}


func insertDhcpsvrip(eNodeBId string, baseName string, data []model.Dhcpsvrip) {
	fmt.Println("[+] Processing Dhcpsvrip data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `dhcpsvrip` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DHCPSVRIP,
			data[j].ATTRIBUTES.DHCPRELAYIPSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Dhcpsvrip data has been saved")
	}
	tx.Commit()
}


func insertDhcpsw(eNodeBId string, baseName string, data []model.Dhcpsw) {
	fmt.Println("[+] Processing Dhcpsw data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `dhcpsw` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SWITCH,
			data[j].ATTRIBUTES.VLANSCANSW,
			data[j].ATTRIBUTES.DELAYSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Dhcpsw data has been saved")
	}
	tx.Commit()
}


func insertDifpri(eNodeBId string, baseName string, data []model.Difpri) {
	fmt.Println("[+] Processing Difpri data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `difpri` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PRIRULE,
			data[j].ATTRIBUTES.SIGPRI,
			data[j].ATTRIBUTES.OMHIGHPRI,
			data[j].ATTRIBUTES.OMLOWPRI,
			data[j].ATTRIBUTES.IPCLKPRI,	

		)
		checkErr(err)
		//fmt.Println("[+] Difpri data has been saved")
	}
	tx.Commit()
}


func insertDistbasedho(eNodeBId string, baseName string, data []model.Distbasedho) {
	fmt.Println("[+] Processing Distbasedho data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `distbasedho` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DistBasedMeasObjType,
			data[j].ATTRIBUTES.DistBasedHOThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Distbasedho data has been saved")
	}
	tx.Commit()
}


func insertDlflowctrlpara(eNodeBId string, baseName string, data []model.Dlflowctrlpara) {
	fmt.Println("[+] Processing Dlflowctrlpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `dlflowctrlpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.BEAR,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.SWITCH,
			data[j].ATTRIBUTES.TD,
			data[j].ATTRIBUTES.DR,
			data[j].ATTRIBUTES.ITM,
			data[j].ATTRIBUTES.BWPRTSWITCH,
			data[j].ATTRIBUTES.FAIRSWITCH,
			data[j].ATTRIBUTES.FAIRRATIO,
			data[j].ATTRIBUTES.DLIUBMINBW,
			data[j].ATTRIBUTES.DLFLOWCTRLENHSW,
			data[j].ATTRIBUTES.DLIUBBWMINPRORATIO,	

		)
		checkErr(err)
		//fmt.Println("[+] Dlflowctrlpara data has been saved")
	}
	tx.Commit()
}


func insertDrx(eNodeBId string, baseName string, data []model.Drx) {
	fmt.Println("[+] Processing Drx data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `drx` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DrxAlgSwitch,
			data[j].ATTRIBUTES.ShortDrxSwitch,
			data[j].ATTRIBUTES.LongDrxCycleSpecial,
			data[j].ATTRIBUTES.OnDurationTimerSpecial,
			data[j].ATTRIBUTES.DrxInactivityTimerSpecial,
			data[j].ATTRIBUTES.SupportShortDrxSpecial,
			data[j].ATTRIBUTES.LongDrxCycleForAnr,
			data[j].ATTRIBUTES.LongDRXCycleforIRatAnr,
			data[j].ATTRIBUTES.DrxInactivityTimerForAnr,
			data[j].ATTRIBUTES.TddAnrDrxInactivityTimer,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.DrxFlexCtrlSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Drx data has been saved")
	}
	tx.Commit()
}


func insertDrxparagroup(eNodeBId string, baseName string, data []model.Drxparagroup) {
	fmt.Println("[+] Processing Drxparagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `drxparagroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DrxParaGroupId,
			data[j].ATTRIBUTES.EnterDrxSwitch,
			data[j].ATTRIBUTES.OnDurationTimer,
			data[j].ATTRIBUTES.DrxInactivityTimer,
			data[j].ATTRIBUTES.DrxReTxTimer,
			data[j].ATTRIBUTES.LongDrxCycle,
			data[j].ATTRIBUTES.SupportShortDrx,
			data[j].ATTRIBUTES.ShortDrxCycle,
			data[j].ATTRIBUTES.DrxShortCycleTimer,
			data[j].ATTRIBUTES.DrxUlReTxTimer,
			data[j].ATTRIBUTES.ExtendLongDrxCycleSwitch,
			data[j].ATTRIBUTES.CatType,	

		)
		checkErr(err)
		//fmt.Println("[+] Drxparagroup data has been saved")
	}
	tx.Commit()
}


func insertDscpmap(eNodeBId string, baseName string, data []model.Dscpmap) {
	fmt.Println("[+] Processing Dscpmap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `dscpmap` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.DSCP,
			data[j].ATTRIBUTES.VLANPRIO,	

		)
		checkErr(err)
		//fmt.Println("[+] Dscpmap data has been saved")
	}
	tx.Commit()
}


func insertE1t1(eNodeBId string, baseName string, data []model.E1t1) {
	fmt.Println("[+] Processing E1t1 data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `e1t1` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PS,
			data[j].ATTRIBUTES.WORKMODE,
			data[j].ATTRIBUTES.FRAME,
			data[j].ATTRIBUTES.LNCODE,
			data[j].ATTRIBUTES.CLKM,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.PN,	

		)
		checkErr(err)
		//fmt.Println("[+] E1t1 data has been saved")
	}
	tx.Commit()
}


func insertE1t1bear(eNodeBId string, baseName string, data []model.E1t1bear) {
	fmt.Println("[+] Processing E1t1bear data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `e1t1bear` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MODE,	

		)
		checkErr(err)
		//fmt.Println("[+] E1t1bear data has been saved")
	}
	tx.Commit()
}


func insertE1t1ber(eNodeBId string, baseName string, data []model.E1t1ber) {
	fmt.Println("[+] Processing E1t1ber data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `e1t1ber` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BTL,	

		)
		checkErr(err)
		//fmt.Println("[+] E1t1ber data has been saved")
	}
	tx.Commit()
}


func insertEmc(eNodeBId string, baseName string, data []model.Emc) {
	fmt.Println("[+] Processing Emc data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `emc` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.EmcEnable,	

		)
		checkErr(err)
		//fmt.Println("[+] Emc data has been saved")
	}
	tx.Commit()
}


func insertEmu(eNodeBId string, baseName string, data []model.Emu) {
	fmt.Println("[+] Processing Emu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `emu` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MCN,
			data[j].ATTRIBUTES.MSRN,
			data[j].ATTRIBUTES.MPN,
			data[j].ATTRIBUTES.ADDR,
			data[j].ATTRIBUTES.TLTHD,
			data[j].ATTRIBUTES.TUTHD,
			data[j].ATTRIBUTES.HLTHD,
			data[j].ATTRIBUTES.HUTHD,
			data[j].ATTRIBUTES.SAAF,
			data[j].ATTRIBUTES.SBAF,	

		)
		checkErr(err)
		//fmt.Println("[+] Emu data has been saved")
	}
	tx.Commit()
}


func insertEnbcelloprsvdpara(eNodeBId string, baseName string, data []model.Enbcelloprsvdpara) {
	fmt.Println("[+] Processing Enbcelloprsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbcelloprsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TrackingAreaId,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara7,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara14,
			data[j].ATTRIBUTES.RsvdPara15,
			data[j].ATTRIBUTES.RsvdPara16,
			data[j].ATTRIBUTES.RsvdPara17,
			data[j].ATTRIBUTES.RsvdPara18,
			data[j].ATTRIBUTES.RsvdPara19,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbcelloprsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnbcellqcirsvdpara(eNodeBId string, baseName string, data []model.Enbcellqcirsvdpara) {
	fmt.Println("[+] Processing Enbcellqcirsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbcellqcirsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara7,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara14,
			data[j].ATTRIBUTES.RsvdPara15,
			data[j].ATTRIBUTES.RsvdPara16,
			data[j].ATTRIBUTES.RsvdPara17,
			data[j].ATTRIBUTES.RsvdPara18,
			data[j].ATTRIBUTES.RsvdPara19,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbcellqcirsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnbcellrsvdpara(eNodeBId string, baseName string, data []model.Enbcellrsvdpara) {
	fmt.Println("[+] Processing Enbcellrsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbcellrsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdSwPara3,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara22,
			data[j].ATTRIBUTES.RsvdPara23,
			data[j].ATTRIBUTES.RsvdPara24,
			data[j].ATTRIBUTES.RsvdPara25,
			data[j].ATTRIBUTES.RsvdPara26,
			data[j].ATTRIBUTES.RsvdPara27,
			data[j].ATTRIBUTES.RsvdPara28,
			data[j].ATTRIBUTES.RsvdPara29,
			data[j].ATTRIBUTES.RsvdPara30,
			data[j].ATTRIBUTES.RsvdPara31,
			data[j].ATTRIBUTES.RsvdPara32,
			data[j].ATTRIBUTES.RsvdPara33,
			data[j].ATTRIBUTES.RsvdPara34,
			data[j].ATTRIBUTES.RsvdPara35,
			data[j].ATTRIBUTES.RsvdPara36,
			data[j].ATTRIBUTES.RsvdPara37,
			data[j].ATTRIBUTES.RsvdPara38,
			data[j].ATTRIBUTES.RsvdPara39,
			data[j].ATTRIBUTES.RsvdPara40,
			data[j].ATTRIBUTES.RsvdPara41,
			data[j].ATTRIBUTES.RsvdPara42,
			data[j].ATTRIBUTES.RsvdPara43,
			data[j].ATTRIBUTES.RsvdPara44,
			data[j].ATTRIBUTES.RsvdPara45,
			data[j].ATTRIBUTES.RsvdPara46,
			data[j].ATTRIBUTES.RsvdPara47,
			data[j].ATTRIBUTES.RsvdPara48,
			data[j].ATTRIBUTES.RsvdPara49,
			data[j].ATTRIBUTES.RsvdPara50,
			data[j].ATTRIBUTES.RsvdPara51,
			data[j].ATTRIBUTES.RsvdPara52,
			data[j].ATTRIBUTES.RsvdPara53,
			data[j].ATTRIBUTES.RsvdPara54,
			data[j].ATTRIBUTES.RsvdPara55,
			data[j].ATTRIBUTES.RsvdPara56,
			data[j].ATTRIBUTES.RsvdPara57,
			data[j].ATTRIBUTES.RsvdPara58,
			data[j].ATTRIBUTES.RsvdPara59,
			data[j].ATTRIBUTES.RsvdPara60,
			data[j].ATTRIBUTES.RsvdPara61,
			data[j].ATTRIBUTES.RsvdPara62,
			data[j].ATTRIBUTES.RsvdPara63,
			data[j].ATTRIBUTES.RsvdPara64,
			data[j].ATTRIBUTES.RsvdPara65,
			data[j].ATTRIBUTES.RsvdPara66,
			data[j].ATTRIBUTES.RsvdPara67,
			data[j].ATTRIBUTES.RsvdPara68,
			data[j].ATTRIBUTES.RsvdPara69,
			data[j].ATTRIBUTES.RsvdU32Para1,
			data[j].ATTRIBUTES.RsvdU32Para2,
			data[j].ATTRIBUTES.RsvdU32Para3,
			data[j].ATTRIBUTES.RsvdU32Para4,
			data[j].ATTRIBUTES.RsvdU32Para5,
			data[j].ATTRIBUTES.RsvdU32Para6,
			data[j].ATTRIBUTES.RsvdU32Para7,
			data[j].ATTRIBUTES.RsvdU32Para8,
			data[j].ATTRIBUTES.RsvdU16Para1,
			data[j].ATTRIBUTES.RsvdU16Para2,
			data[j].ATTRIBUTES.RsvdU16Para3,
			data[j].ATTRIBUTES.RsvdU16Para4,
			data[j].ATTRIBUTES.RsvdU16Para5,
			data[j].ATTRIBUTES.RsvdU16Para6,
			data[j].ATTRIBUTES.RsvdU16Para7,
			data[j].ATTRIBUTES.RsvdU16Para8,
			data[j].ATTRIBUTES.RsvdU16Para9,
			data[j].ATTRIBUTES.RsvdU16Para10,
			data[j].ATTRIBUTES.RsvdU16Para11,
			data[j].ATTRIBUTES.RsvdU16Para12,
			data[j].ATTRIBUTES.RsvdU16Para13,
			data[j].ATTRIBUTES.RsvdU16Para14,
			data[j].ATTRIBUTES.RsvdU16Para15,
			data[j].ATTRIBUTES.RsvdU16Para16,
			data[j].ATTRIBUTES.RsvdU8Para1,
			data[j].ATTRIBUTES.RsvdU8Para2,
			data[j].ATTRIBUTES.RsvdU8Para3,
			data[j].ATTRIBUTES.RsvdU8Para4,
			data[j].ATTRIBUTES.RsvdU8Para5,
			data[j].ATTRIBUTES.RsvdU8Para6,
			data[j].ATTRIBUTES.RsvdU8Para7,
			data[j].ATTRIBUTES.RsvdU8Para8,
			data[j].ATTRIBUTES.RsvdU8Para9,
			data[j].ATTRIBUTES.RsvdU8Para10,
			data[j].ATTRIBUTES.RsvdU8Para11,
			data[j].ATTRIBUTES.RsvdU8Para12,
			data[j].ATTRIBUTES.RsvdU8Para13,
			data[j].ATTRIBUTES.RsvdU8Para14,
			data[j].ATTRIBUTES.RsvdU8Para15,
			data[j].ATTRIBUTES.RsvdU8Para16,
			data[j].ATTRIBUTES.RsvdU8Para17,
			data[j].ATTRIBUTES.RsvdU8Para18,
			data[j].ATTRIBUTES.RsvdU8Para19,
			data[j].ATTRIBUTES.RsvdU8Para20,
			data[j].ATTRIBUTES.RsvdU8Para21,
			data[j].ATTRIBUTES.RsvdU8Para22,
			data[j].ATTRIBUTES.RsvdU8Para23,
			data[j].ATTRIBUTES.RsvdU8Para24,
			data[j].ATTRIBUTES.RsvdU8Para25,
			data[j].ATTRIBUTES.RsvdU8Para26,
			data[j].ATTRIBUTES.RsvdU8Para27,
			data[j].ATTRIBUTES.RsvdU8Para28,
			data[j].ATTRIBUTES.RsvdU8Para29,
			data[j].ATTRIBUTES.RsvdU8Para30,
			data[j].ATTRIBUTES.RsvdU8Para31,
			data[j].ATTRIBUTES.RsvdU8Para32,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbcellrsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnbcnopqcirsvdpara(eNodeBId string, baseName string, data []model.Enbcnopqcirsvdpara) {
	fmt.Println("[+] Processing Enbcnopqcirsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbcnopqcirsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara7,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara14,
			data[j].ATTRIBUTES.RsvdPara15,
			data[j].ATTRIBUTES.RsvdPara16,
			data[j].ATTRIBUTES.RsvdPara17,
			data[j].ATTRIBUTES.RsvdPara18,
			data[j].ATTRIBUTES.RsvdPara19,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbcnopqcirsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnbcnoprsvdpara(eNodeBId string, baseName string, data []model.Enbcnoprsvdpara) {
	fmt.Println("[+] Processing Enbcnoprsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbcnoprsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara7,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara14,
			data[j].ATTRIBUTES.RsvdPara15,
			data[j].ATTRIBUTES.RsvdPara16,
			data[j].ATTRIBUTES.RsvdPara17,
			data[j].ATTRIBUTES.RsvdPara18,
			data[j].ATTRIBUTES.RsvdPara19,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbcnoprsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnblicensealmthd(eNodeBId string, baseName string, data []model.Enblicensealmthd) {
	fmt.Println("[+] Processing Enblicensealmthd data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enblicensealmthd` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OPRD,
			data[j].ATTRIBUTES.OTHD,
			data[j].ATTRIBUTES.RPRD,
			data[j].ATTRIBUTES.RTHD,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enblicensealmthd data has been saved")
	}
	tx.Commit()
}


func insertEnbqcirsvdpara(eNodeBId string, baseName string, data []model.Enbqcirsvdpara) {
	fmt.Println("[+] Processing Enbqcirsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbqcirsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdPara2,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara7,
			data[j].ATTRIBUTES.RsvdPara8,
			data[j].ATTRIBUTES.RsvdPara9,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara14,
			data[j].ATTRIBUTES.RsvdPara15,
			data[j].ATTRIBUTES.RsvdPara16,
			data[j].ATTRIBUTES.RsvdPara17,
			data[j].ATTRIBUTES.RsvdPara18,
			data[j].ATTRIBUTES.RsvdPara19,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbqcirsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnbrsvdpara(eNodeBId string, baseName string, data []model.Enbrsvdpara) {
	fmt.Println("[+] Processing Enbrsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enbrsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RsvdSwPara0,
			data[j].ATTRIBUTES.RsvdSwPara1,
			data[j].ATTRIBUTES.RsvdSwPara2,
			data[j].ATTRIBUTES.RsvdSwPara3,
			data[j].ATTRIBUTES.RsvdPara3,
			data[j].ATTRIBUTES.RsvdPara4,
			data[j].ATTRIBUTES.RsvdPara5,
			data[j].ATTRIBUTES.RsvdPara6,
			data[j].ATTRIBUTES.RsvdPara10,
			data[j].ATTRIBUTES.RsvdPara11,
			data[j].ATTRIBUTES.RsvdPara12,
			data[j].ATTRIBUTES.RsvdPara13,
			data[j].ATTRIBUTES.RsvdPara20,
			data[j].ATTRIBUTES.RsvdPara21,
			data[j].ATTRIBUTES.RsvdPara22,
			data[j].ATTRIBUTES.RsvdPara23,
			data[j].ATTRIBUTES.RsvdPara24,
			data[j].ATTRIBUTES.RsvdPara25,
			data[j].ATTRIBUTES.RsvdPara26,
			data[j].ATTRIBUTES.RsvdPara27,
			data[j].ATTRIBUTES.RsvdPara28,
			data[j].ATTRIBUTES.RsvdPara29,
			data[j].ATTRIBUTES.RsvdPara31,
			data[j].ATTRIBUTES.RsvdPara32,
			data[j].ATTRIBUTES.RsvdPara33,
			data[j].ATTRIBUTES.RsvdPara34,
			data[j].ATTRIBUTES.RsvdPara35,
			data[j].ATTRIBUTES.RsvdPara36,
			data[j].ATTRIBUTES.RsvdPara37,
			data[j].ATTRIBUTES.RsvdPara38,
			data[j].ATTRIBUTES.RsvdPara39,
			data[j].ATTRIBUTES.RsvdPara40,
			data[j].ATTRIBUTES.RsvdPara41,
			data[j].ATTRIBUTES.RsvdPara42,
			data[j].ATTRIBUTES.RsvdPara43,
			data[j].ATTRIBUTES.RsvdPara44,
			data[j].ATTRIBUTES.RsvdPara45,
			data[j].ATTRIBUTES.RsvdPara46,
			data[j].ATTRIBUTES.RsvdPara47,
			data[j].ATTRIBUTES.RsvdPara48,
			data[j].ATTRIBUTES.RsvdPara49,
			data[j].ATTRIBUTES.RsvdPara50,
			data[j].ATTRIBUTES.RsvdPara51,
			data[j].ATTRIBUTES.RsvdPara52,
			data[j].ATTRIBUTES.RsvdPara53,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enbrsvdpara data has been saved")
	}
	tx.Commit()
}


func insertEnergycon(eNodeBId string, baseName string, data []model.Energycon) {
	fmt.Println("[+] Processing Energycon data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `energycon` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.MP,	

		)
		checkErr(err)
		//fmt.Println("[+] Energycon data has been saved")
	}
	tx.Commit()
}


func insertEnodebalgoswitch(eNodeBId string, baseName string, data []model.Enodebalgoswitch) {
	fmt.Println("[+] Processing Enodebalgoswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebalgoswitch` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.HoAlgoSwitch,
			data[j].ATTRIBUTES.HoModeSwitch,
			data[j].ATTRIBUTES.DlIcicSwitch,
			data[j].ATTRIBUTES.AnrSwitch,
			data[j].ATTRIBUTES.RedirectSwitch,
			data[j].ATTRIBUTES.MroSwitch,
			data[j].ATTRIBUTES.MacAssemblyOptSwitch,
			data[j].ATTRIBUTES.TpeSwitch,
			data[j].ATTRIBUTES.SpidSelectPlmnAlgoSwitch,
			data[j].ATTRIBUTES.UlIcicFreqSwitch,
			data[j].ATTRIBUTES.LcsSwitch,
			data[j].ATTRIBUTES.TrmSwitch,
			data[j].ATTRIBUTES.PciConflictAlmSwitch,
			data[j].ATTRIBUTES.PowerSaveSwitch,
			data[j].ATTRIBUTES.RimSwitch,
			data[j].ATTRIBUTES.RanSharingAnrSwitch,
			data[j].ATTRIBUTES.FreqLayerSwtich,
			data[j].ATTRIBUTES.CmasSwitch,
			data[j].ATTRIBUTES.VQMAlgoSwitch,
			data[j].ATTRIBUTES.UeNumPreemptSwitch,
			data[j].ATTRIBUTES.MultiOpCtrlSwitch,
			data[j].ATTRIBUTES.OverBBUsSwitch,
			data[j].ATTRIBUTES.OperatorSpecificAlgoSwitch,
			data[j].ATTRIBUTES.HoSignalingOptSwitch,
			data[j].ATTRIBUTES.CompatibilityCtrlSwitch,
			data[j].ATTRIBUTES.BlindNcellOptSwitch,
			data[j].ATTRIBUTES.RimOnEcoSwitch,
			data[j].ATTRIBUTES.EutranVoipSupportSwitch,
			data[j].ATTRIBUTES.CaAlgoSwitch,
			data[j].ATTRIBUTES.MlbAlgoSwitch,
			data[j].ATTRIBUTES.HoCommOptSwitch,
			data[j].ATTRIBUTES.HighLoadNetOptSwitch,
			data[j].ATTRIBUTES.SchOptSwitch,
			data[j].ATTRIBUTES.PrachTimeStagSwitch,
			data[j].ATTRIBUTES.HighSpeedRootSeqCSSwitch,
			data[j].ATTRIBUTES.DropPktsStatSwitch,
			data[j].ATTRIBUTES.NCellRankingSwitch,
			data[j].ATTRIBUTES.EUCompactRANSwitch,
			data[j].ATTRIBUTES.BbpCollaborateSwitch,
			data[j].ATTRIBUTES.RootSeqConflictDetSwitch,
			data[j].ATTRIBUTES.IOptAlgoSwitch,
			data[j].ATTRIBUTES.ServiceHoMultiTargetFreqSw,
			data[j].ATTRIBUTES.SFSSwitch,
			data[j].ATTRIBUTES.PciConflictDetectSwitch,
			data[j].ATTRIBUTES.CaLbAlgoSwitch,
			data[j].ATTRIBUTES.UlSchOptSwitch,
			data[j].ATTRIBUTES.CompactRANMultiAPN,
			data[j].ATTRIBUTES.ClkJumpCellReStartSwitch,
			data[j].ATTRIBUTES.E2EVQIAlgoSwitch,
			data[j].ATTRIBUTES.TlFreqFrameOffsetSwitch,
			data[j].ATTRIBUTES.FastEnhancePccAnchorSwitch,
			data[j].ATTRIBUTES.HoWithSccCfgAddBlindSwitch,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.AntCalibrationTimeSwitch,
			data[j].ATTRIBUTES.PimSwitch,
			data[j].ATTRIBUTES.FltrRptRrcConReqExtdSwitch,
			data[j].ATTRIBUTES.SSRDAlgoSwitch,
			data[j].ATTRIBUTES.IoTClkJumpCellResetSw,
			data[j].ATTRIBUTES.CaAlgoExtSwitch,
			data[j].ATTRIBUTES.UlResManageOptSw,
			data[j].ATTRIBUTES.LTEPreemptNbSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebalgoswitch data has been saved")
	}
	tx.Commit()
}


func insertEnodebalmcfg(eNodeBId string, baseName string, data []model.Enodebalmcfg) {
	fmt.Println("[+] Processing Enodebalmcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebalmcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TrafficFaultyDetectPeriod,
			data[j].ATTRIBUTES.SrvIntfAlmCfgSw,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebalmcfg data has been saved")
	}
	tx.Commit()
}


func insertEnodebautopoweroff(eNodeBId string, baseName string, data []model.Enodebautopoweroff) {
	fmt.Println("[+] Processing Enodebautopoweroff data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebautopoweroff` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AutoPowerOffSwitch,
			data[j].ATTRIBUTES.PowerOffTime,
			data[j].ATTRIBUTES.PowerOnTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebautopoweroff data has been saved")
	}
	tx.Commit()
}


func insertEnodebchroutputctrl(eNodeBId string, baseName string, data []model.Enodebchroutputctrl) {
	fmt.Println("[+] Processing Enodebchroutputctrl data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebchroutputctrl` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OutPutMode,
			data[j].ATTRIBUTES.CallSampleRate,
			data[j].ATTRIBUTES.MaxStoreCall,
			data[j].ATTRIBUTES.CHRUpLoadingTimeSwitch,
			data[j].ATTRIBUTES.SIGOutputMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebchroutputctrl data has been saved")
	}
	tx.Commit()
}


func insertEnodebciphercap(eNodeBId string, baseName string, data []model.Enodebciphercap) {
	fmt.Println("[+] Processing Enodebciphercap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebciphercap` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PrimaryCipherAlgo,
			data[j].ATTRIBUTES.SecondCipherAlgo,
			data[j].ATTRIBUTES.ThirdCipherAlgo,
			data[j].ATTRIBUTES.FourthCipherAlgo,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebciphercap data has been saved")
	}
	tx.Commit()
}


func insertEnodebconnstatetimer(eNodeBId string, baseName string, data []model.Enodebconnstatetimer) {
	fmt.Println("[+] Processing Enodebconnstatetimer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebconnstatetimer` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.S1MessageWaitingTimer,
			data[j].ATTRIBUTES.X2MessageWaitingTimer,
			data[j].ATTRIBUTES.UuMessageWaitingTimer,
			data[j].ATTRIBUTES.Cdma1XrttHoUuPrepareTimer,
			data[j].ATTRIBUTES.Cdma1XrttHoS1WaitingTimer,
			data[j].ATTRIBUTES.Cdma1XrttHoCompleteTimer,
			data[j].ATTRIBUTES.CdmaHrpdHoCompleteTimer,
			data[j].ATTRIBUTES.CdmaHrpdHoS1WaitingTimer,
			data[j].ATTRIBUTES.CdmaHrpdHoUuPrepareTimer,
			data[j].ATTRIBUTES.WaitRrcConnSetupCmpTimer,
			data[j].ATTRIBUTES.SecCmpWaitingTimer,
			data[j].ATTRIBUTES.UpUeCapInfoWaitingTimer,
			data[j].ATTRIBUTES.FirstForwardPacketTimer,
			data[j].ATTRIBUTES.EndMarkerTimer,
			data[j].ATTRIBUTES.S1MsgWaitingTimerQci1,
			data[j].ATTRIBUTES.X2MessageWaitingTimerQci1,
			data[j].ATTRIBUTES.UuMessageWaitingTimerQci1,
			data[j].ATTRIBUTES.BearerActivityThd,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.RrcReestActivityThld,
			data[j].ATTRIBUTES.EcidWaitingTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebconnstatetimer data has been saved")
	}
	tx.Commit()
}


func insertEnodebfddbbres(eNodeBId string, baseName string, data []model.Enodebfddbbres) {
	fmt.Println("[+] Processing Enodebfddbbres data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebfddbbres` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENodeBFddBbResId,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.Capacity,
			data[j].ATTRIBUTES.AdmState,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebfddbbres data has been saved")
	}
	tx.Commit()
}


func insertEnodebflowctrlpara(eNodeBId string, baseName string, data []model.Enodebflowctrlpara) {
	fmt.Println("[+] Processing Enodebflowctrlpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebflowctrlpara` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AdaptUnsyncTimerLen,
			data[j].ATTRIBUTES.AdaptUnsyncUserNumThd,
			data[j].ATTRIBUTES.DynAcBarPolicyMode,
			data[j].ATTRIBUTES.CpuLoadThd,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.McpttMsgCntSwitch,
			data[j].ATTRIBUTES.DynAcBarringTrigAllCellSw,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebflowctrlpara data has been saved")
	}
	tx.Commit()
}


func insertEnodebframeoffset(eNodeBId string, baseName string, data []model.Enodebframeoffset) {
	fmt.Println("[+] Processing Enodebframeoffset data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebframeoffset` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TddFrameOffset,
			data[j].ATTRIBUTES.FddFrameOffset,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebframeoffset data has been saved")
	}
	tx.Commit()
}


func insertEnodebfunction(eNodeBId string, baseName string, data []model.Enodebfunction) {
	fmt.Println("[+] Processing Enodebfunction data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebfunction` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.ENodeBFunctionName,
			data[j].ATTRIBUTES.ApplicationRef,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.UserLabel,
			data[j].ATTRIBUTES.NermVersion,
			data[j].ATTRIBUTES.ProductVersion,
			data[j].ATTRIBUTES.ProductInterfaceID,
			data[j].ATTRIBUTES.LmtVersion,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebfunction data has been saved")
	}
	tx.Commit()
}


func insertEnodebintegritycap(eNodeBId string, baseName string, data []model.Enodebintegritycap) {
	fmt.Println("[+] Processing Enodebintegritycap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebintegritycap` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PrimaryIntegrityAlgo,
			data[j].ATTRIBUTES.SecondIntegrityAlgo,
			data[j].ATTRIBUTES.ThirdIntegrityAlgo,
			data[j].ATTRIBUTES.NullAlgo,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebintegritycap data has been saved")
	}
	tx.Commit()
}


func insertEnodebmlb(eNodeBId string, baseName string, data []model.Enodebmlb) {
	fmt.Println("[+] Processing Enodebmlb data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebmlb` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.InterFreqIdleMlbMode,
			data[j].ATTRIBUTES.InterFreqIdleMlbInterval,
			data[j].ATTRIBUTES.InterFreqIdleMlbStaThd,
			data[j].ATTRIBUTES.DlCaLbMaxCCNum,
			data[j].ATTRIBUTES.CaUeMinDataVolRatio,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebmlb data has been saved")
	}
	tx.Commit()
}


func insertEnodebnbpara(eNodeBId string, baseName string, data []model.Enodebnbpara) {
	fmt.Println("[+] Processing Enodebnbpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebnbpara` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.NbRsvMinUserNumRatio,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebnbpara data has been saved")
	}
	tx.Commit()
}


func insertEnodebresmodealgo(eNodeBId string, baseName string, data []model.Enodebresmodealgo) {
	fmt.Println("[+] Processing Enodebresmodealgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebresmodealgo` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BbpResAutoConfigSw,
			data[j].ATTRIBUTES.TddBbpResDeployAlgo,
			data[j].ATTRIBUTES.FddBbpResDeployAlgo,
			data[j].ATTRIBUTES.BbpResAutoRecfgTimer,
			data[j].ATTRIBUTES.ENBCellNumCapabilityMode,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.BbpCpriSharingSwitch,
			data[j].ATTRIBUTES.EnbCellDstDeploySw,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebresmodealgo data has been saved")
	}
	tx.Commit()
}


func insertEnodebsharingmode(eNodeBId string, baseName string, data []model.Enodebsharingmode) {
	fmt.Println("[+] Processing Enodebsharingmode data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebsharingmode` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENodeBSharingMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebsharingmode data has been saved")
	}
	tx.Commit()
}


func insertEnodebsondbcfg(eNodeBId string, baseName string, data []model.Enodebsondbcfg) {
	fmt.Println("[+] Processing Enodebsondbcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebsondbcfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.StartTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebsondbcfg data has been saved")
	}
	tx.Commit()
}


func insertEnodebtddbbres(eNodeBId string, baseName string, data []model.Enodebtddbbres) {
	fmt.Println("[+] Processing Enodebtddbbres data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebtddbbres` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENodeBTddBbResId,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.Capacity,
			data[j].ATTRIBUTES.AdmState,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebtddbbres data has been saved")
	}
	tx.Commit()
}


func insertEnodebtgalgcfg(eNodeBId string, baseName string, data []model.Enodebtgalgcfg) {
	fmt.Println("[+] Processing Enodebtgalgcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebtgalgcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.FilterCoeffTG,
			data[j].ATTRIBUTES.StatPeriodTG,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebtgalgcfg data has been saved")
	}
	tx.Commit()
}


func insertEnodebusparacfg(eNodeBId string, baseName string, data []model.Enodebusparacfg) {
	fmt.Println("[+] Processing Enodebusparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `enodebusparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UsStrategy,
			data[j].ATTRIBUTES.UsSPIDConfig,
			data[j].ATTRIBUTES.UsUeInactiveTimer,
			data[j].ATTRIBUTES.IPDetectInterval,
			data[j].ATTRIBUTES.UsUeDlPriorityFactor,
			data[j].ATTRIBUTES.UsLPktUlschPriFactor,
			data[j].ATTRIBUTES.UsSPktUlschPriFactor,
			data[j].ATTRIBUTES.UsUESTMSIKeepAliveTimer,
			data[j].ATTRIBUTES.BigPackageIdentify,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.UlDlUsUserDetectPeriod,	

		)
		checkErr(err)
		//fmt.Println("[+] Enodebusparacfg data has been saved")
	}
	tx.Commit()
}


func insertEpgroup(eNodeBId string, baseName string, data []model.Epgroup) {
	fmt.Println("[+] Processing Epgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `epgroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.EPGROUPID,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.SCTPHOSTREFS,
			data[j].ATTRIBUTES.SCTPPEERREFS,
			data[j].ATTRIBUTES.USERPLANEHOSTREFS,
			data[j].ATTRIBUTES.USERPLANEPEERREFS,
			data[j].ATTRIBUTES.PACKETFILTERSWITCH,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.TYPEFLAG,
			data[j].ATTRIBUTES.LNKPFMSW,
			data[j].ATTRIBUTES.CTRLMODE,
			data[j].ATTRIBUTES.AUTOCFGFLAG,	

		)
		checkErr(err)
		//fmt.Println("[+] Epgroup data has been saved")
	}
	tx.Commit()
}


func insertEquipment(eNodeBId string, baseName string, data []model.Equipment) {
	fmt.Println("[+] Processing Equipment data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `equipment` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BATIMS,
			data[j].ATTRIBUTES.PAE,
			data[j].ATTRIBUTES.MNTMODE,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.ET,
			data[j].ATTRIBUTES.MMSETREMARK,
			data[j].ATTRIBUTES.DVAS,
			data[j].ATTRIBUTES.DVAH,
			data[j].ATTRIBUTES.DVAL,
			data[j].ATTRIBUTES.ODIID,
			data[j].ATTRIBUTES.DCALMSW,
			data[j].ATTRIBUTES.EQUIPMENTTY,
			data[j].ATTRIBUTES.PSUFP,
			data[j].ATTRIBUTES.PROTOCOL,
			data[j].ATTRIBUTES.SMARTTRX,
			data[j].ATTRIBUTES.ESN,
			data[j].ATTRIBUTES.SDBBLSD,
			data[j].ATTRIBUTES.APST,
			data[j].ATTRIBUTES.NPST,
			data[j].ATTRIBUTES.SDRCONNSW,
			data[j].ATTRIBUTES.PWRFAILOMCONNENHSW,
			data[j].ATTRIBUTES.OMCONNENHSWCTLTIME,	

		)
		checkErr(err)
		//fmt.Println("[+] Equipment data has been saved")
	}
	tx.Commit()
}


func insertEthport(eNodeBId string, baseName string, data []model.Ethport) {
	fmt.Println("[+] Processing Ethport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ethport` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PA,
			data[j].ATTRIBUTES.MTU,
			data[j].ATTRIBUTES.SPEED,
			data[j].ATTRIBUTES.DUPLEX,
			data[j].ATTRIBUTES.ARPPROXY,
			data[j].ATTRIBUTES.FC,
			data[j].ATTRIBUTES.FERAT,
			data[j].ATTRIBUTES.FERDT,
			data[j].ATTRIBUTES.RXBCPKTALMOCRTHD,
			data[j].ATTRIBUTES.RXBCPKTALMCLRTHD,
			data[j].ATTRIBUTES.FIBERSPEEDMATCH,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.AUTOCFGFLAG,	

		)
		checkErr(err)
		//fmt.Println("[+] Ethport data has been saved")
	}
	tx.Commit()
}


func insertEucellalmcfg(eNodeBId string, baseName string, data []model.Eucellalmcfg) {
	fmt.Println("[+] Processing Eucellalmcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucellalmcfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RxNoisePwrThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucellalmcfg data has been saved")
	}
	tx.Commit()
}


func insertEucellsectoreqm(eNodeBId string, baseName string, data []model.Eucellsectoreqm) {
	fmt.Println("[+] Processing Eucellsectoreqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucellsectoreqm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SectorEqmId,
			data[j].ATTRIBUTES.ReferenceSignalPwr,
			data[j].ATTRIBUTES.BaseBandEqmId,
			data[j].ATTRIBUTES.ReferenceSignalPwrMargin,
			data[j].ATTRIBUTES.SectorCpriCompression,
			data[j].ATTRIBUTES.AutoCfgFlag,
			data[j].ATTRIBUTES.WeightNO,
			data[j].ATTRIBUTES.VisualCellId,
			data[j].ATTRIBUTES.PrbIdList,
			data[j].ATTRIBUTES.SectorEqmCombineGrpId,
			data[j].ATTRIBUTES.CellBeamMode,
			data[j].ATTRIBUTES.BeamId,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucellsectoreqm data has been saved")
	}
	tx.Commit()
}


func insertEucommcellsectoreqm(eNodeBId string, baseName string, data []model.Eucommcellsectoreqm) {
	fmt.Println("[+] Processing Eucommcellsectoreqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucommcellsectoreqm` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.SectorEqmId,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucommcellsectoreqm data has been saved")
	}
	tx.Commit()
}


func insertEucoschcfg(eNodeBId string, baseName string, data []model.Eucoschcfg) {
	fmt.Println("[+] Processing Eucoschcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucoschcfg` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PrtNodeBaseBandEqmId,
			data[j].ATTRIBUTES.SchNodeBaseBandEqmId,
			data[j].ATTRIBUTES.WorkMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucoschcfg data has been saved")
	}
	tx.Commit()
}


func insertEucoschdlcompcfg(eNodeBId string, baseName string, data []model.Eucoschdlcompcfg) {
	fmt.Println("[+] Processing Eucoschdlcompcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucoschdlcompcfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CordInfoEffDelay,
			data[j].ATTRIBUTES.InterEnbDlCompSwitch,
			data[j].ATTRIBUTES.EuCoSchUeSpec,
			data[j].ATTRIBUTES.CordInfoEffDelayMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucoschdlcompcfg data has been saved")
	}
	tx.Commit()
}


func insertEucoschulicscfg(eNodeBId string, baseName string, data []model.Eucoschulicscfg) {
	fmt.Println("[+] Processing Eucoschulicscfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eucoschulicscfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UlIcsAlgoSwitch,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Eucoschulicscfg data has been saved")
	}
	tx.Commit()
}


func insertEuepcsec(eNodeBId string, baseName string, data []model.Euepcsec) {
	fmt.Println("[+] Processing Euepcsec data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `euepcsec` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CiphAlgo,
			data[j].ATTRIBUTES.IntAlgo,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Euepcsec data has been saved")
	}
	tx.Commit()
}


func insertEutranexternalcell(eNodeBId string, baseName string, data []model.Eutranexternalcell) {
	fmt.Println("[+] Processing Eutranexternalcell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eutranexternalcell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.DlEarfcn,
			data[j].ATTRIBUTES.UlEarfcnCfgInd,
			data[j].ATTRIBUTES.PhyCellId,
			data[j].ATTRIBUTES.Tac,
			data[j].ATTRIBUTES.CellName,
			data[j].ATTRIBUTES.EutranExternalCellSlaveBand,
			data[j].ATTRIBUTES.RoamingAreaHoInd,
			data[j].ATTRIBUTES.SpecifiedCellFlag,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.HighSpeedFlag,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.DlFreqOffset,
			data[j].ATTRIBUTES.NbCellFlag,
			data[j].ATTRIBUTES.NclUpdateMode,
			data[j].ATTRIBUTES.SupportEmtcFlag,	

		)
		checkErr(err)
		//fmt.Println("[+] Eutranexternalcell data has been saved")
	}
	tx.Commit()
}


func insertEutraninterfreqncell(eNodeBId string, baseName string, data []model.Eutraninterfreqncell) {
	fmt.Println("[+] Processing Eutraninterfreqncell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eutraninterfreqncell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.CellIndividualOffset,
			data[j].ATTRIBUTES.CellQoffset,
			data[j].ATTRIBUTES.NoHoFlag,
			data[j].ATTRIBUTES.NoRmvFlag,
			data[j].ATTRIBUTES.BlindHoPriority,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.LocalCellName,
			data[j].ATTRIBUTES.NeighbourCellName,
			data[j].ATTRIBUTES.CellMeasPriority,
			data[j].ATTRIBUTES.OverlapInd,
			data[j].ATTRIBUTES.OverlapRange,
			data[j].ATTRIBUTES.NCellClassLabel,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.AggregationProperty,
			data[j].ATTRIBUTES.OverlapIndicatorExtension,	

		)
		checkErr(err)
		//fmt.Println("[+] Eutraninterfreqncell data has been saved")
	}
	tx.Commit()
}


func insertEutraninternfreq(eNodeBId string, baseName string, data []model.Eutraninternfreq) {
	fmt.Println("[+] Processing Eutraninternfreq data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eutraninternfreq` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlEarfcn,
			data[j].ATTRIBUTES.UlEarfcnCfgInd,
			data[j].ATTRIBUTES.CellReselPriorityCfgInd,
			data[j].ATTRIBUTES.CellReselPriority,
			data[j].ATTRIBUTES.EutranReselTime,
			data[j].ATTRIBUTES.SpeedDependSPCfgInd,
			data[j].ATTRIBUTES.MeasBandWidth,
			data[j].ATTRIBUTES.QoffsetFreq,
			data[j].ATTRIBUTES.ThreshXhigh,
			data[j].ATTRIBUTES.ThreshXlow,
			data[j].ATTRIBUTES.QRxLevMin,
			data[j].ATTRIBUTES.PmaxCfgInd,
			data[j].ATTRIBUTES.NeighCellConfig,
			data[j].ATTRIBUTES.PresenceAntennaPort1,
			data[j].ATTRIBUTES.InterFreqHoEventType,
			data[j].ATTRIBUTES.ThreshXhighQ,
			data[j].ATTRIBUTES.ThreshXlowQ,
			data[j].ATTRIBUTES.QqualMinCfgInd,
			data[j].ATTRIBUTES.ConnFreqPriority,
			data[j].ATTRIBUTES.MlbTargetInd,
			data[j].ATTRIBUTES.FreqPriBasedHoMeasFlag,
			data[j].ATTRIBUTES.IdleMlbUEReleaseRatio,
			data[j].ATTRIBUTES.MlbFreqPriority,
			data[j].ATTRIBUTES.QoffsetFreqConn,
			data[j].ATTRIBUTES.MeasFreqPriority,
			data[j].ATTRIBUTES.IfHoThdRsrpOffset,
			data[j].ATTRIBUTES.IfMlbThdRsrpOffset,
			data[j].ATTRIBUTES.MasterBandFlag,
			data[j].ATTRIBUTES.InterFreqRanSharingInd,
			data[j].ATTRIBUTES.InterFreqHighSpeedFlag,
			data[j].ATTRIBUTES.AnrInd,
			data[j].ATTRIBUTES.VoipPriority,
			data[j].ATTRIBUTES.PsPriority,
			data[j].ATTRIBUTES.VolteHoTargetInd,
			data[j].ATTRIBUTES.FreqPriorityForAnr,
			data[j].ATTRIBUTES.BackoffTargetInd,
			data[j].ATTRIBUTES.MlbInterFreqHoEventType,
			data[j].ATTRIBUTES.MobilityTargetInd,
			data[j].ATTRIBUTES.MlbInterFreqEffiRatio,
			data[j].ATTRIBUTES.SnrBasedUeSelectionMode,
			data[j].ATTRIBUTES.UlTrafficMlbTargetInd,
			data[j].ATTRIBUTES.UlTrafficMlbPriority,
			data[j].ATTRIBUTES.MlbInterFreqHoA3Offset,
			data[j].ATTRIBUTES.IfSrvHoThdRsrpOffset,
			data[j].ATTRIBUTES.IfSrvHoThdRsrqOffset,
			data[j].ATTRIBUTES.MlbFreqUlPriority,
			data[j].ATTRIBUTES.InterFreqMlbDlPrbOffset,
			data[j].ATTRIBUTES.InterFreqMlbUlPrbOffset,
			data[j].ATTRIBUTES.NcellNumForAnr,
			data[j].ATTRIBUTES.MeasPerformanceDemand,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.DlFreqOffset,
			data[j].ATTRIBUTES.IfBackoffThdRsrpOffset,
			data[j].ATTRIBUTES.IfBackoffThdRsrqOffset,
			data[j].ATTRIBUTES.VoLTEQualityIfHoTargetInd,
			data[j].ATTRIBUTES.IdleMlbeMtcUEReleaseRatio,
			data[j].ATTRIBUTES.InterFreqCioAdjLimitCfgInd,
			data[j].ATTRIBUTES.InterFreq4TInd,	

		)
		checkErr(err)
		//fmt.Println("[+] Eutraninternfreq data has been saved")
	}
	tx.Commit()
}


func insertEutranintrafreqncell(eNodeBId string, baseName string, data []model.Eutranintrafreqncell) {
	fmt.Println("[+] Processing Eutranintrafreqncell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `eutranintrafreqncell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.ENodeBId,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.CellIndividualOffset,
			data[j].ATTRIBUTES.CellQoffset,
			data[j].ATTRIBUTES.NoHoFlag,
			data[j].ATTRIBUTES.NoRmvFlag,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.LocalCellName,
			data[j].ATTRIBUTES.NeighbourCellName,
			data[j].ATTRIBUTES.CellMeasPriority,
			data[j].ATTRIBUTES.CellRangeExpansion,
			data[j].ATTRIBUTES.NCellClassLabel,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.AttachCellSwitch,
			data[j].ATTRIBUTES.HighSpeedCellIndOffset,
			data[j].ATTRIBUTES.VectorCellFlag,	

		)
		checkErr(err)
		//fmt.Println("[+] Eutranintrafreqncell data has been saved")
	}
	tx.Commit()
}


func insertEuulcoschcfg(eNodeBId string, baseName string, data []model.Euulcoschcfg) {
	fmt.Println("[+] Processing Euulcoschcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `euulcoschcfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UlInterEnbCamcSw,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Euulcoschcfg data has been saved")
	}
	tx.Commit()
}


func insertExtendedqci(eNodeBId string, baseName string, data []model.Extendedqci) {
	fmt.Println("[+] Processing Extendedqci data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `extendedqci` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ExtendedQci,
			data[j].ATTRIBUTES.ServiceType,
			data[j].ATTRIBUTES.UlschPriorityFactor,
			data[j].ATTRIBUTES.DlschPriorityFactor,
			data[j].ATTRIBUTES.UlMinGbr,
			data[j].ATTRIBUTES.DlMinGbr,
			data[j].ATTRIBUTES.PreAllocationWeight,
			data[j].ATTRIBUTES.InterRatPolicyCfgGroupId,
			data[j].ATTRIBUTES.PrioritisedBitRate,
			data[j].ATTRIBUTES.LogicalChannelPriority,
			data[j].ATTRIBUTES.RlcPdcpParaGroupId,
			data[j].ATTRIBUTES.FreeUserFlag,
			data[j].ATTRIBUTES.FlowCtrlType,
			data[j].ATTRIBUTES.LtePttDedicatedExtendedQci,
			data[j].ATTRIBUTES.LtePttDownlinkPriority,
			data[j].ATTRIBUTES.LtePttUplinkPriority,
			data[j].ATTRIBUTES.ExtQciCounterIndex,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Extendedqci data has been saved")
	}
	tx.Commit()
}


func insertFddresmode(eNodeBId string, baseName string, data []model.Fddresmode) {
	fmt.Println("[+] Processing Fddresmode data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `fddresmode` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BbCapabilityMode,
			data[j].ATTRIBUTES.SfnCapabilityMode,
			data[j].ATTRIBUTES.BbResExclusiveSwitch,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Fddresmode data has been saved")
	}
	tx.Commit()
}


func insertFltcorrenablecfg(eNodeBId string, baseName string, data []model.Fltcorrenablecfg) {
	fmt.Println("[+] Processing Fltcorrenablecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `fltcorrenablecfg` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENABLERELA,	

		)
		checkErr(err)
		//fmt.Println("[+] Fltcorrenablecfg data has been saved")
	}
	tx.Commit()
}


func insertFmu(eNodeBId string, baseName string, data []model.Fmu) {
	fmt.Println("[+] Processing Fmu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `fmu` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MCN,
			data[j].ATTRIBUTES.MSRN,
			data[j].ATTRIBUTES.MPN,
			data[j].ATTRIBUTES.ADDR,
			data[j].ATTRIBUTES.SBAF,
			data[j].ATTRIBUTES.STC,
			data[j].ATTRIBUTES.TCMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Fmu data has been saved")
	}
	tx.Commit()
}


func insertFtpclt(eNodeBId string, baseName string, data []model.Ftpclt) {
	fmt.Println("[+] Processing Ftpclt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ftpclt` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ENCRYMODE,
			data[j].ATTRIBUTES.SPTSTATEFWL,
			data[j].ATTRIBUTES.SSLCERTAUTH,
			data[j].ATTRIBUTES.MINDHLEN,	

		)
		checkErr(err)
		//fmt.Println("[+] Ftpclt data has been saved")
	}
	tx.Commit()
}


func insertFtpcltport(eNodeBId string, baseName string, data []model.Ftpcltport) {
	fmt.Println("[+] Processing Ftpcltport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ftpcltport` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SERVERIP,
			data[j].ATTRIBUTES.PORT,
			data[j].ATTRIBUTES.STARTDATAPORT,
			data[j].ATTRIBUTES.ENDDATAPORT,	

		)
		checkErr(err)
		//fmt.Println("[+] Ftpcltport data has been saved")
	}
	tx.Commit()
}


func insertGbtsabiscp(eNodeBId string, baseName string, data []model.Gbtsabiscp) {
	fmt.Println("[+] Processing Gbtsabiscp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtsabiscp` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ABISCPID,
			data[j].ATTRIBUTES.CPBEARID,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtsabiscp data has been saved")
	}
	tx.Commit()
}


func insertGbtsbbres(eNodeBId string, baseName string, data []model.Gbtsbbres) {
	fmt.Println("[+] Processing Gbtsbbres data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtsbbres` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GBTSBBRESID,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.CAPACITY,
			data[j].ATTRIBUTES.ADMSTATE,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtsbbres data has been saved")
	}
	tx.Commit()
}


func insertGbtsenergymgtpara(eNodeBId string, baseName string, data []model.Gbtsenergymgtpara) {
	fmt.Println("[+] Processing Gbtsenergymgtpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtsenergymgtpara` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.BAKPWRSAVMETHOD,
			data[j].ATTRIBUTES.PAADJVOL,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtsenergymgtpara data has been saved")
	}
	tx.Commit()
}


func insertGbtsfunction(eNodeBId string, baseName string, data []model.Gbtsfunction) {
	fmt.Println("[+] Processing Gbtsfunction data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtsfunction` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.GBTSFUNCTIONNAME,
			data[j].ATTRIBUTES.APPLICATIONREF,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.NERMVERSION,
			data[j].ATTRIBUTES.PRODUCTVERSION,
			data[j].ATTRIBUTES.InterfaceId,
			data[j].ATTRIBUTES.LmtVersion,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtsfunction data has been saved")
	}
	tx.Commit()
}


func insertGbtsglobalpara(eNodeBId string, baseName string, data []model.Gbtsglobalpara) {
	fmt.Println("[+] Processing Gbtsglobalpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtsglobalpara` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.MCSTANDARD,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtsglobalpara data has been saved")
	}
	tx.Commit()
}


func insertGbtspath(eNodeBId string, baseName string, data []model.Gbtspath) {
	fmt.Println("[+] Processing Gbtspath data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gbtspath` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PATHID,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Gbtspath data has been saved")
	}
	tx.Commit()
}


func insertGeranexternalcell(eNodeBId string, baseName string, data []model.Geranexternalcell) {
	fmt.Println("[+] Processing Geranexternalcell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `geranexternalcell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.GeranCellId,
			data[j].ATTRIBUTES.Lac,
			data[j].ATTRIBUTES.RacCfgInd,
			data[j].ATTRIBUTES.Rac,
			data[j].ATTRIBUTES.BandIndicator,
			data[j].ATTRIBUTES.GeranArfcn,
			data[j].ATTRIBUTES.NetworkColourCode,
			data[j].ATTRIBUTES.BaseStationColourCode,
			data[j].ATTRIBUTES.DtmInd,
			data[j].ATTRIBUTES.CellName,
			data[j].ATTRIBUTES.CsPsHOInd,
			data[j].ATTRIBUTES.UltraFlashCsfbInd,
			data[j].ATTRIBUTES.RoamingAreaHoInd,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Geranexternalcell data has been saved")
	}
	tx.Commit()
}


func insertGeraninterfcfg(eNodeBId string, baseName string, data []model.Geraninterfcfg) {
	fmt.Println("[+] Processing Geraninterfcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `geraninterfcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlGeranIntefRbNum,
			data[j].ATTRIBUTES.DlRbAvailSinrThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Geraninterfcfg data has been saved")
	}
	tx.Commit()
}


func insertGeranncell(eNodeBId string, baseName string, data []model.Geranncell) {
	fmt.Println("[+] Processing Geranncell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `geranncell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.Lac,
			data[j].ATTRIBUTES.GeranCellId,
			data[j].ATTRIBUTES.NoRmvFlag,
			data[j].ATTRIBUTES.NoHoFlag,
			data[j].ATTRIBUTES.BlindHoPriority,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.LocalCellName,
			data[j].ATTRIBUTES.NeighbourCellName,
			data[j].ATTRIBUTES.OverlapInd,
			data[j].ATTRIBUTES.NCellMeasPriority,
			data[j].ATTRIBUTES.CtrlMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Geranncell data has been saved")
	}
	tx.Commit()
}


func insertGerannfreqgroup(eNodeBId string, baseName string, data []model.Gerannfreqgroup) {
	fmt.Println("[+] Processing Gerannfreqgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gerannfreqgroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BcchGroupId,
			data[j].ATTRIBUTES.GeranVersion,
			data[j].ATTRIBUTES.StartingArfcn,
			data[j].ATTRIBUTES.BandIndicator,
			data[j].ATTRIBUTES.CellReselPriorityCfgInd,
			data[j].ATTRIBUTES.CellReselPriority,
			data[j].ATTRIBUTES.PmaxGeranCfgInd,
			data[j].ATTRIBUTES.QRxLevMin,
			data[j].ATTRIBUTES.ThreshXHigh,
			data[j].ATTRIBUTES.ThreshXLow,
			data[j].ATTRIBUTES.OffsetFreq,
			data[j].ATTRIBUTES.NccPermitted,
			data[j].ATTRIBUTES.ConnFreqPriority,
			data[j].ATTRIBUTES.ContinuCoverageIndication,
			data[j].ATTRIBUTES.GeranRanSharingInd,
			data[j].ATTRIBUTES.AnrInd,
			data[j].ATTRIBUTES.FreqPriorityForAnr,
			data[j].ATTRIBUTES.PmaxGeran,	

		)
		checkErr(err)
		//fmt.Println("[+] Gerannfreqgroup data has been saved")
	}
	tx.Commit()
}


func insertGerannfreqgrouparfcn(eNodeBId string, baseName string, data []model.Gerannfreqgrouparfcn) {
	fmt.Println("[+] Processing Gerannfreqgrouparfcn data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gerannfreqgrouparfcn` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.BcchGroupId,
			data[j].ATTRIBUTES.GeranArfcn,	

		)
		checkErr(err)
		//fmt.Println("[+] Gerannfreqgrouparfcn data has been saved")
	}
	tx.Commit()
}


func insertGlobalprocswitch(eNodeBId string, baseName string, data []model.Globalprocswitch) {
	fmt.Println("[+] Processing Globalprocswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `globalprocswitch` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.X2SonSetupSwitch,
			data[j].ATTRIBUTES.X2SonLinkSetupType,
			data[j].ATTRIBUTES.SriAdaptiveSwitch,
			data[j].ATTRIBUTES.LcgProfile,
			data[j].ATTRIBUTES.RncPoolHoRimSwitch,
			data[j].ATTRIBUTES.UuMsgSimulSendSwitch,
			data[j].ATTRIBUTES.X2BasedUptENodeBCfgSwitch,
			data[j].ATTRIBUTES.TargetOpBasedX2Switch,
			data[j].ATTRIBUTES.X2SonTNLSelectMode,
			data[j].ATTRIBUTES.UtranLoadTransChan,
			data[j].ATTRIBUTES.S1HoInDataFwdSwitch,
			data[j].ATTRIBUTES.RrcReestProtectThd,
			data[j].ATTRIBUTES.X2SonDeleteTimer,
			data[j].ATTRIBUTES.RimCodingPolicy,
			data[j].ATTRIBUTES.L2GUHoWithLCapSwitch,
			data[j].ATTRIBUTES.DiffOpWithSameMmecSwitch,
			data[j].ATTRIBUTES.PeerReqBasedX2DelSwitch,
			data[j].ATTRIBUTES.UlPdcpSduRcvStatSendSwitch,
			data[j].ATTRIBUTES.ProtocolMsgOptSwitch,
			data[j].ATTRIBUTES.X2BasedDelNcellCfgSwitch,
			data[j].ATTRIBUTES.X2ServedCellType,
			data[j].ATTRIBUTES.EnbTrigMmeLoadRebalSwitch,
			data[j].ATTRIBUTES.UeRelReSynTimes,
			data[j].ATTRIBUTES.UeRelChkLostSwitch,
			data[j].ATTRIBUTES.UeLinkAbnormalDetectSwitch,
			data[j].ATTRIBUTES.S1DefaultPagingDrxSelect,
			data[j].ATTRIBUTES.EnhancedPhrRptCtrlSw,
			data[j].ATTRIBUTES.EutranLoadTransSwitch,
			data[j].ATTRIBUTES.ProtocolSupportSwitch,
			data[j].ATTRIBUTES.CellTrafficTraceMsgSwitch,
			data[j].ATTRIBUTES.PreferSigExtend,
			data[j].ATTRIBUTES.MmeSelectProcSwitch,
			data[j].ATTRIBUTES.S1OfflineCommitSwitch,
			data[j].ATTRIBUTES.ProtocolCompatibilitySw,
			data[j].ATTRIBUTES.X2SctpEstType,
			data[j].ATTRIBUTES.X2SonDeleteSwitch,
			data[j].ATTRIBUTES.X2SonDeleteTimerforX2Fault,
			data[j].ATTRIBUTES.X2SonDeleteTimerforX2Usage,
			data[j].ATTRIBUTES.X2SonDeleteHoInNumThd,
			data[j].ATTRIBUTES.X2SonDeleteHoOutNumThd,
			data[j].ATTRIBUTES.VoipWithGapMode,
			data[j].ATTRIBUTES.IntraEnodebHoStaticSw,
			data[j].ATTRIBUTES.MaxSyncUserNumPerBbi,
			data[j].ATTRIBUTES.X2SetupFailureTimetoWait,
			data[j].ATTRIBUTES.X2SetupFailureNumThd,
			data[j].ATTRIBUTES.X2SetupFailureNumTimer,
			data[j].ATTRIBUTES.EX2AutoDeleteSwitch,
			data[j].ATTRIBUTES.EX2AutoDeleteTimerforFault,
			data[j].ATTRIBUTES.X2SonDeleteMode,
			data[j].ATTRIBUTES.X2SonSetupNumThd,
			data[j].ATTRIBUTES.X2SonSetupTimer,
			data[j].ATTRIBUTES.SecKeyRecfgSwitch,
			data[j].ATTRIBUTES.X2BasedUptENodeBPolicy,
			data[j].ATTRIBUTES.MMEDomNameMode,
			data[j].ATTRIBUTES.RrcConnPunishThd,
			data[j].ATTRIBUTES.X2BasedUptNcellCfgSwitch,
			data[j].ATTRIBUTES.HoProcCtrlSwitch,
			data[j].ATTRIBUTES.RrcReestOptSwitch,
			data[j].ATTRIBUTES.S1MMESonSwitch,
			data[j].ATTRIBUTES.PrivateMdtUeSelSwitch,
			data[j].ATTRIBUTES.QciUpdParaCheckSwitch,
			data[j].ATTRIBUTES.UeCompatSwitch,
			data[j].ATTRIBUTES.S1MmePrivFeatureInd,
			data[j].ATTRIBUTES.SctpAbortSmoothSwitch,
			data[j].ATTRIBUTES.MaxUserNumPerCell,
			data[j].ATTRIBUTES.CsfbFlowOptSwitch,
			data[j].ATTRIBUTES.X2InitFailDelSwitch,
			data[j].ATTRIBUTES.X2DynBlacklistAgingTimer,
			data[j].ATTRIBUTES.EX2AutoSetupSwitch,
			data[j].ATTRIBUTES.EX2DynBlackListSwitch,
			data[j].ATTRIBUTES.EX2DynBlackListAgingTimer,
			data[j].ATTRIBUTES.PRSMutingCtrlSwitch,
			data[j].ATTRIBUTES.RrcConnReqStatSwitch,
			data[j].ATTRIBUTES.QciParaEffectFlag,
			data[j].ATTRIBUTES.RimFirstMultiReportSwitch,
			data[j].ATTRIBUTES.UuMsgTimeOutRelCause,
			data[j].ATTRIBUTES.SigPoolOptPolicy,
			data[j].ATTRIBUTES.EnhancedRRCReestProtectThd,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.S1FaultSelfRecoverySwitch,
			data[j].ATTRIBUTES.ItfTypeForNonIdealModeServ,
			data[j].ATTRIBUTES.S1DefaultPagingDrxForNb,
			data[j].ATTRIBUTES.X2SonSecMode,
			data[j].ATTRIBUTES.ProcTypeForNonIdealServData,
			data[j].ATTRIBUTES.Ex2DeleteTimerForEx2Usage,
			data[j].ATTRIBUTES.SinglePlmnHostSplitSwitch,
			data[j].ATTRIBUTES.InitMsgProtectThld,	

		)
		checkErr(err)
		//fmt.Println("[+] Globalprocswitch data has been saved")
	}
	tx.Commit()
}


func insertGlocell(eNodeBId string, baseName string, data []model.Glocell) {
	fmt.Println("[+] Processing Glocell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocell` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.BASEBANDPOLICY,
			data[j].ATTRIBUTES.BASEBANDEQMID,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.LOCELLTYPE,
			data[j].ATTRIBUTES.BBEQMPOLICY,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocell data has been saved")
	}
	tx.Commit()
}


func insertGlocellalgpara(eNodeBId string, baseName string, data []model.Glocellalgpara) {
	fmt.Println("[+] Processing Glocellalgpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocellalgpara` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.GUDEGRATEPWRCTRL,
			data[j].ATTRIBUTES.GMSKDELAY,
			data[j].ATTRIBUTES.DIVERT8PSKDELAY,
			data[j].ATTRIBUTES.DIVERT16QAMDELAY,
			data[j].ATTRIBUTES.DIVERT32QAMDELAY,
			data[j].ATTRIBUTES.GMSKDELAYDYNADJSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocellalgpara data has been saved")
	}
	tx.Commit()
}


func insertGlocellenergymgtpara(eNodeBId string, baseName string, data []model.Glocellenergymgtpara) {
	fmt.Println("[+] Processing Glocellenergymgtpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocellenergymgtpara` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.MAINBCCHPWRDTEN,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocellenergymgtpara data has been saved")
	}
	tx.Commit()
}


func insertGlocellothpara(eNodeBId string, baseName string, data []model.Glocellothpara) {
	fmt.Println("[+] Processing Glocellothpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocellothpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.MULTRXBRDSTABMONITUNNSW,
			data[j].ATTRIBUTES.MULTRXBRDUNSTABMONITUNNSW,
			data[j].ATTRIBUTES.IMMOCCUPYPCHOPTSW,
			data[j].ATTRIBUTES.PCHFORBIDRPTLOADSW,
			data[j].ATTRIBUTES.PTCCHPOWEROPTSW,
			data[j].ATTRIBUTES.FIRSTMROPTSW,
			data[j].ATTRIBUTES.SPEECHDELAYOPTSW,
			data[j].ATTRIBUTES.SPEECHBADFRAMEOPTSW,
			data[j].ATTRIBUTES.PAGINGOVLDPROCOPTSW,
			data[j].ATTRIBUTES.RFMODULESTDETECTSW,
			data[j].ATTRIBUTES.NOTRAFFICSELFHEALSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocellothpara data has been saved")
	}
	tx.Commit()
}


func insertGlocellrlalmpara(eNodeBId string, baseName string, data []model.Glocellrlalmpara) {
	fmt.Println("[+] Processing Glocellrlalmpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocellrlalmpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.ARELWTHD,
			data[j].ATTRIBUTES.AREDISTHD,
			data[j].ATTRIBUTES.NOTRASP,
			data[j].ATTRIBUTES.WLASTIME,
			data[j].ATTRIBUTES.WLAETIME,
			data[j].ATTRIBUTES.BWTHD,
			data[j].ATTRIBUTES.GUNFAULTDECTSW,
			data[j].ATTRIBUTES.GUNCHDROPDECTTHRD,
			data[j].ATTRIBUTES.GUNCHDROPCLRTHRD,
			data[j].ATTRIBUTES.GUNBADQUALDETCTTHRD,
			data[j].ATTRIBUTES.GUNBADQUALCLRTHRD,
			data[j].ATTRIBUTES.RSSIUNBALANCEDTHRD,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocellrlalmpara data has been saved")
	}
	tx.Commit()
}


func insertGlocellrsvdpara(eNodeBId string, baseName string, data []model.Glocellrsvdpara) {
	fmt.Println("[+] Processing Glocellrsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `glocellrsvdpara` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA1,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA2,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA3,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA4,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA5,
			data[j].ATTRIBUTES.GLOCELLRSVDPARA6,	

		)
		checkErr(err)
		//fmt.Println("[+] Glocellrsvdpara data has been saved")
	}
	tx.Commit()
}


func insertGps(eNodeBId string, baseName string, data []model.Gps) {
	fmt.Println("[+] Processing Gps data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gps` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GN,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.WPOS,
			data[j].ATTRIBUTES.LONG,
			data[j].ATTRIBUTES.LAT,
			data[j].ATTRIBUTES.ALT,
			data[j].ATTRIBUTES.AGL,
			data[j].ATTRIBUTES.DURATION,
			data[j].ATTRIBUTES.PRECISION,
			data[j].ATTRIBUTES.CABLE_LEN,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.PRI,
			data[j].ATTRIBUTES.POSCHECKSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Gps data has been saved")
	}
	tx.Commit()
}


func insertGtpu(eNodeBId string, baseName string, data []model.Gtpu) {
	fmt.Println("[+] Processing Gtpu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gtpu` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TIMEOUTTH,
			data[j].ATTRIBUTES.TIMEOUTCNT,
			data[j].ATTRIBUTES.DSCP,
			data[j].ATTRIBUTES.STATICCHK,
			data[j].ATTRIBUTES.ONLYUPIPCHK,
			data[j].ATTRIBUTES.ULGTPUSNPADSW,
			data[j].ATTRIBUTES.DLGTPUSNCHKSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Gtpu data has been saved")
	}
	tx.Commit()
}


func insertGtranspara(eNodeBId string, baseName string, data []model.Gtranspara) {
	fmt.Println("[+] Processing Gtranspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gtranspara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LPSCHSW,
			data[j].ATTRIBUTES.RATECFGTYPE,
			data[j].ATTRIBUTES.SBTIME,
			data[j].ATTRIBUTES.ARPAGINGTIME,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.OAMTYPE,
			data[j].ATTRIBUTES.BYPASSSWITCH,
			data[j].ATTRIBUTES.STATUS,
			data[j].ATTRIBUTES.OMCHSWITCHBACK,
			data[j].ATTRIBUTES.CPLBSW,
			data[j].ATTRIBUTES.CPLISTENINGMODE,
			data[j].ATTRIBUTES.SCTPRTRPTSW,
			data[j].ATTRIBUTES.IPERRFRMSW,
			data[j].ATTRIBUTES.VLANID,
			data[j].ATTRIBUTES.FORWARDMODE,
			data[j].ATTRIBUTES.SAMEPORTFORWARDSW,
			data[j].ATTRIBUTES.ONLYOMIP,
			data[j].ATTRIBUTES.DNSPRD,
			data[j].ATTRIBUTES.EPOBJAUTOCFGSW,
			data[j].ATTRIBUTES.EPCFGMODE,
			data[j].ATTRIBUTES.PMAUTOSW,
			data[j].ATTRIBUTES.DIRECTIPSECPRIMATCHSW,
			data[j].ATTRIBUTES.MODELSELECT,	

		)
		checkErr(err)
		//fmt.Println("[+] Gtranspara data has been saved")
	}
	tx.Commit()
}


func insertGtransparagw(eNodeBId string, baseName string, data []model.Gtransparagw) {
	fmt.Println("[+] Processing Gtransparagw data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gtransparagw` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PRIRULE,
			data[j].ATTRIBUTES.LPSCHSW,
			data[j].ATTRIBUTES.RATECFGTYPE,
			data[j].ATTRIBUTES.SBTIME,
			data[j].ATTRIBUTES.ARPAGINGTIME,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.OAMTYPE,
			data[j].ATTRIBUTES.IPERRFRMSW,
			data[j].ATTRIBUTES.FORWARDMODE,
			data[j].ATTRIBUTES.SAMEPORTFORWARDSW,
			data[j].ATTRIBUTES.PMAUTOSW,
			data[j].ATTRIBUTES.IPCLKPRI,
			data[j].ATTRIBUTES.PORTULOBSW,
			data[j].ATTRIBUTES.PORTDLOBSW,
			data[j].ATTRIBUTES.PORTULCACSW,
			data[j].ATTRIBUTES.PORTDLCACSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Gtransparagw data has been saved")
	}
	tx.Commit()
}


func insertGtrxgroup(eNodeBId string, baseName string, data []model.Gtrxgroup) {
	fmt.Println("[+] Processing Gtrxgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gtrxgroup` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GTRXGROUPID,
			data[j].ATTRIBUTES.GLOCELLID,
			data[j].ATTRIBUTES.SNDMD,
			data[j].ATTRIBUTES.RCVMD,
			data[j].ATTRIBUTES.WORKMODE,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Gtrxgroup data has been saved")
	}
	tx.Commit()
}


func insertGtrxgroupsectoreqm(eNodeBId string, baseName string, data []model.Gtrxgroupsectoreqm) {
	fmt.Println("[+] Processing Gtrxgroupsectoreqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `gtrxgroupsectoreqm` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GTRXGROUPID,
			data[j].ATTRIBUTES.SECTOREQMID,
			data[j].ATTRIBUTES.MAXPOWER,	

		)
		checkErr(err)
		//fmt.Println("[+] Gtrxgroupsectoreqm data has been saved")
	}
	tx.Commit()
}


func insertHighspdadaptionpara(eNodeBId string, baseName string, data []model.Highspdadaptionpara) {
	fmt.Println("[+] Processing Highspdadaptionpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `highspdadaptionpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UserLowSpeedJudgeNum,
			data[j].ATTRIBUTES.DopplerFactor,
			data[j].ATTRIBUTES.HighSpeedUserNumThd,
			data[j].ATTRIBUTES.HoHisJudgeHighThd,
			data[j].ATTRIBUTES.InterfAvoidCellNum,
			data[j].ATTRIBUTES.InterfAvoidMode,
			data[j].ATTRIBUTES.InterfBasedPowerOff,
			data[j].ATTRIBUTES.InterfBasedRbRatio,
			data[j].ATTRIBUTES.HighspdVectorUpdatePeriod,
			data[j].ATTRIBUTES.HighspdVectorUpdateSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Highspdadaptionpara data has been saved")
	}
	tx.Commit()
}


func insertHomeascomm(eNodeBId string, baseName string, data []model.Homeascomm) {
	fmt.Println("[+] Processing Homeascomm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `homeascomm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.GapPatternType,
			data[j].ATTRIBUTES.EutranFilterCoeffRsrp,
			data[j].ATTRIBUTES.EutranFilterCoeffRsrq,
			data[j].ATTRIBUTES.GeranFilterCoeff,
			data[j].ATTRIBUTES.SMeasureInd,
			data[j].ATTRIBUTES.SpeedDepParaInd,
			data[j].ATTRIBUTES.UtranFilterCoeffRscp,
			data[j].ATTRIBUTES.UtranFilterCoeffEcn0,
			data[j].ATTRIBUTES.OptHoPreFailPunishTimer,
			data[j].ATTRIBUTES.ResHoPreFailPunishTimer,
			data[j].ATTRIBUTES.NonResHoPreFailPunishTimes,
			data[j].ATTRIBUTES.NonResHoPreFailRetryTimes,
			data[j].ATTRIBUTES.DedicatedGapPatternType,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Homeascomm data has been saved")
	}
	tx.Commit()
}


func insertHtcdpa(eNodeBId string, baseName string, data []model.Htcdpa) {
	fmt.Println("[+] Processing Htcdpa data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `htcdpa` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.LTCP,
			data[j].ATTRIBUTES.HTCP,
			data[j].ATTRIBUTES.TLT,
			data[j].ATTRIBUTES.DBD,
			data[j].ATTRIBUTES.NTDI,
			data[j].ATTRIBUTES.NTDO,
			data[j].ATTRIBUTES.HTDO,	

		)
		checkErr(err)
		//fmt.Println("[+] Htcdpa data has been saved")
	}
	tx.Commit()
}


func insertIkecfg(eNodeBId string, baseName string, data []model.Ikecfg) {
	fmt.Println("[+] Processing Ikecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ikecfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IKELNM,
			data[j].ATTRIBUTES.IKEKLI,
			data[j].ATTRIBUTES.IKEKLT,
			data[j].ATTRIBUTES.NATKLI,
			data[j].ATTRIBUTES.DSCP,
			data[j].ATTRIBUTES.IPSECSWITCHBACK,
			data[j].ATTRIBUTES.IPSECSBWAITTIME,
			data[j].ATTRIBUTES.IPSECSBRANDTIME,
			data[j].ATTRIBUTES.IPSECSOWAITTIME,
			data[j].ATTRIBUTES.IPSECSORANDTIME,	

		)
		checkErr(err)
		//fmt.Println("[+] Ikecfg data has been saved")
	}
	tx.Commit()
}


func insertImagrp(eNodeBId string, baseName string, data []model.Imagrp) {
	fmt.Println("[+] Processing Imagrp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `imagrp` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.IMAGRPN,
			data[j].ATTRIBUTES.VER,
			data[j].ATTRIBUTES.CLKM,
			data[j].ATTRIBUTES.FRMLEN,
			data[j].ATTRIBUTES.MINLNK,
			data[j].ATTRIBUTES.DELAY,
			data[j].ATTRIBUTES.SCRAM,
			data[j].ATTRIBUTES.TS16,
			data[j].ATTRIBUTES.SBT,	

		)
		checkErr(err)
		//fmt.Println("[+] Imagrp data has been saved")
	}
	tx.Commit()
}


func insertImalnk(eNodeBId string, baseName string, data []model.Imalnk) {
	fmt.Println("[+] Processing Imalnk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `imalnk` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.IMALNKN,
			data[j].ATTRIBUTES.IMAGRPN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.IMAGRPSBT,	

		)
		checkErr(err)
		//fmt.Println("[+] Imalnk data has been saved")
	}
	tx.Commit()
}


func insertIngchktsk(eNodeBId string, baseName string, data []model.Ingchktsk) {
	fmt.Println("[+] Processing Ingchktsk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ingchktsk` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.FLINGCHKSW,
			data[j].ATTRIBUTES.FLINGCHKTM,
			data[j].ATTRIBUTES.FLINGCHKITEM,	

		)
		checkErr(err)
		//fmt.Println("[+] Ingchktsk data has been saved")
	}
	tx.Commit()
}


func insertInterfreqhogroup(eNodeBId string, baseName string, data []model.Interfreqhogroup) {
	fmt.Println("[+] Processing Interfreqhogroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interfreqhogroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterFreqHoGroupId,
			data[j].ATTRIBUTES.InterFreqHoA1A2Hyst,
			data[j].ATTRIBUTES.InterFreqHoA1A2TimeToTrig,
			data[j].ATTRIBUTES.InterFreqHoA1ThdRsrp,
			data[j].ATTRIBUTES.InterFreqHoA1ThdRsrq,
			data[j].ATTRIBUTES.InterFreqHoA2ThdRsrp,
			data[j].ATTRIBUTES.InterFreqHoA2ThdRsrq,
			data[j].ATTRIBUTES.InterFreqHoA4Hyst,
			data[j].ATTRIBUTES.InterFreqHoA4ThdRsrp,
			data[j].ATTRIBUTES.InterFreqHoA4ThdRsrq,
			data[j].ATTRIBUTES.InterFreqHoA4TimeToTrig,
			data[j].ATTRIBUTES.InterFreqLoadBasedHoA4ThdRsrp,
			data[j].ATTRIBUTES.InterFreqLoadBasedHoA4ThdRsrq,
			data[j].ATTRIBUTES.FreqPriInterFreqHoA1ThdRsrp,
			data[j].ATTRIBUTES.FreqPriInterFreqHoA1ThdRsrq,
			data[j].ATTRIBUTES.InterFreqHoA3Offset,
			data[j].ATTRIBUTES.A3InterFreqHoA1ThdRsrp,
			data[j].ATTRIBUTES.A3InterFreqHoA2ThdRsrp,
			data[j].ATTRIBUTES.FreqPriInterFreqHoA2ThdRsrp,
			data[j].ATTRIBUTES.FreqPriInterFreqHoA2ThdRsrq,
			data[j].ATTRIBUTES.InterFreqMlbA1A2ThdRsrp,
			data[j].ATTRIBUTES.InterFreqHoA5Thd1Rsrp,
			data[j].ATTRIBUTES.InterFreqHoA5Thd1Rsrq,
			data[j].ATTRIBUTES.SrvReqHoA4ThdRsrp,
			data[j].ATTRIBUTES.SrvReqHoA4ThdRsrq,
			data[j].ATTRIBUTES.MlbInterFreqHoA5Thd1Rsrp,
			data[j].ATTRIBUTES.MlbInterFreqHoA5Thd1Rsrq,
			data[j].ATTRIBUTES.UlBadQualHoA4Offset,
			data[j].ATTRIBUTES.A3InterFreqHoA1ThdRsrq,
			data[j].ATTRIBUTES.A3InterFreqHoA2ThdRsrq,
			data[j].ATTRIBUTES.UlHeavyTrafficMlbA4ThdRsrp,
			data[j].ATTRIBUTES.UlHeavyTrafficMlbA4ThdRsrq,
			data[j].ATTRIBUTES.InterFreqHoA1ThdRsrpCE,
			data[j].ATTRIBUTES.InterFreqHoA2ThdRsrpCE,	

		)
		checkErr(err)
		//fmt.Println("[+] Interfreqhogroup data has been saved")
	}
	tx.Commit()
}


func insertInterratcellshutdown(eNodeBId string, baseName string, data []model.Interratcellshutdown) {
	fmt.Println("[+] Processing Interratcellshutdown data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interratcellshutdown` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ForceShutdownSwitch,
			data[j].ATTRIBUTES.StartTime,
			data[j].ATTRIBUTES.StopTime,
			data[j].ATTRIBUTES.BearNumThd,
			data[j].ATTRIBUTES.DlPrbThd,
			data[j].ATTRIBUTES.UlPrbThd,
			data[j].ATTRIBUTES.ShutDownType,
			data[j].ATTRIBUTES.HighPriArpThd,
			data[j].ATTRIBUTES.UtranCellDlLoadThd,
			data[j].ATTRIBUTES.UtranCellDlLoadOffset,
			data[j].ATTRIBUTES.UtranCellUlLoadThd,
			data[j].ATTRIBUTES.UtranCellUlLoadOffset,
			data[j].ATTRIBUTES.GeranCellLoadThd,
			data[j].ATTRIBUTES.GeranCellLoadOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Interratcellshutdown data has been saved")
	}
	tx.Commit()
}


func insertInterrathocdma1xrttgroup(eNodeBId string, baseName string, data []model.Interrathocdma1xrttgroup) {
	fmt.Println("[+] Processing Interrathocdma1xrttgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathocdma1xrttgroup` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterRatHoCdma1xRttGroupId,
			data[j].ATTRIBUTES.InterRatHoCdmaB1Hyst,
			data[j].ATTRIBUTES.InterRatHoCdmaB1ThdPs,
			data[j].ATTRIBUTES.InterRatHoCdmaB1TimeToTrig,
			data[j].ATTRIBUTES.LdSvBasedHoCdmaB1ThdPs,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathocdma1xrttgroup data has been saved")
	}
	tx.Commit()
}


func insertInterrathocdmahrpdgroup(eNodeBId string, baseName string, data []model.Interrathocdmahrpdgroup) {
	fmt.Println("[+] Processing Interrathocdmahrpdgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathocdmahrpdgroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterRatHoCdmaHrpdGroupId,
			data[j].ATTRIBUTES.InterRatHoCdmaB1Hyst,
			data[j].ATTRIBUTES.InterRatHoCdmaB1ThdPs,
			data[j].ATTRIBUTES.InterRatHoCdmaB1TimeToTrig,
			data[j].ATTRIBUTES.LdSvBasedHoCdmaB1ThdPs,
			data[j].ATTRIBUTES.Cdma2000HrpdB2Thd1Rsrp,
			data[j].ATTRIBUTES.Cdma2000HrpdB2Thd1Rsrq,
			data[j].ATTRIBUTES.Cdma2000HrpdNonB2ThdRsrp,
			data[j].ATTRIBUTES.Cdma2000HrpdNonB2ThdRsrq,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathocdmahrpdgroup data has been saved")
	}
	tx.Commit()
}


func insertInterrathocomm(eNodeBId string, baseName string, data []model.Interrathocomm) {
	fmt.Println("[+] Processing Interrathocomm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathocomm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.InterRatHoMaxRprtCell,
			data[j].ATTRIBUTES.InterRatHoRprtAmount,
			data[j].ATTRIBUTES.InterRatHoGeranRprtInterval,
			data[j].ATTRIBUTES.InterRatHoUtranB1MeasQuan,
			data[j].ATTRIBUTES.InterRatHoUtranRprtInterval,
			data[j].ATTRIBUTES.InterRatHoCdma1xRttB1MeasQuan,
			data[j].ATTRIBUTES.InterRatCdma1xRttRprtInterval,
			data[j].ATTRIBUTES.InterRatHoCdmaHrpdB1MeasQuan,
			data[j].ATTRIBUTES.InterRatCdmaHrpdRprtInterval,
			data[j].ATTRIBUTES.InterRatHoA1A2TrigQuan,
			data[j].ATTRIBUTES.InterRatHoEventType,
			data[j].ATTRIBUTES.Cdma20001XrttMeasTimer,
			data[j].ATTRIBUTES.Cdma20001XrttJudgePnNum,
			data[j].ATTRIBUTES.Cdma2000HrpdFreqSelMode,
			data[j].ATTRIBUTES.Cdma20001XrttFreqSelMode,
			data[j].ATTRIBUTES.CdmaEcsfbPsConcurrentMode,
			data[j].ATTRIBUTES.Cdma20001XrttMeasMode,
			data[j].ATTRIBUTES.Cdma1XrttSectorIdSelMode,
			data[j].ATTRIBUTES.CellInfoMaxUtranCellNum,
			data[j].ATTRIBUTES.CellInfoMaxGeranCellNum,
			data[j].ATTRIBUTES.GeranCellNumForEmcRedirect,
			data[j].ATTRIBUTES.UtranCellNumForEmcRedirect,
			data[j].ATTRIBUTES.CdmaHrpdSectorIdSelMode,
			data[j].ATTRIBUTES.IRatBlindRedirPlmnCfgSimSw,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathocomm data has been saved")
	}
	tx.Commit()
}


func insertInterrathocommgroup(eNodeBId string, baseName string, data []model.Interrathocommgroup) {
	fmt.Println("[+] Processing Interrathocommgroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathocommgroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterRatHoCommGroupId,
			data[j].ATTRIBUTES.InterRatHoA1A2Hyst,
			data[j].ATTRIBUTES.InterRatHoA1A2TimeToTrig,
			data[j].ATTRIBUTES.InterRatHoA1ThdRsrp,
			data[j].ATTRIBUTES.InterRatHoA1ThdRsrq,
			data[j].ATTRIBUTES.InterRatHoA2ThdRsrp,
			data[j].ATTRIBUTES.InterRatHoA2ThdRsrq,
			data[j].ATTRIBUTES.GeranB2Thd1Rsrp,
			data[j].ATTRIBUTES.GeranB2Thd1Rsrq,
			data[j].ATTRIBUTES.UtranB2Thd1Rsrp,
			data[j].ATTRIBUTES.UtranB2Thd1Rsrq,
			data[j].ATTRIBUTES.DelIfMeasA2ThdRsrpOffset,
			data[j].ATTRIBUTES.DelIfMeasA2ThdRsrqOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathocommgroup data has been saved")
	}
	tx.Commit()
}


func insertInterrathogerangroup(eNodeBId string, baseName string, data []model.Interrathogerangroup) {
	fmt.Println("[+] Processing Interrathogerangroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathogerangroup` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterRatHoGeranGroupId,
			data[j].ATTRIBUTES.InterRatHoGeranB1Hyst,
			data[j].ATTRIBUTES.InterRatHoGeranB1Thd,
			data[j].ATTRIBUTES.InterRatHoGeranB1TimeToTrig,
			data[j].ATTRIBUTES.LdSvBasedHoGeranB1Thd,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathogerangroup data has been saved")
	}
	tx.Commit()
}


func insertInterrathoutrangroup(eNodeBId string, baseName string, data []model.Interrathoutrangroup) {
	fmt.Println("[+] Processing Interrathoutrangroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interrathoutrangroup` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.InterRatHoUtranGroupId,
			data[j].ATTRIBUTES.InterRatHoUtranB1ThdEcn0,
			data[j].ATTRIBUTES.InterRatHoUtranB1ThdRscp,
			data[j].ATTRIBUTES.InterRatHoUtranB1Hyst,
			data[j].ATTRIBUTES.InterRatHoUtranB1TimeToTrig,
			data[j].ATTRIBUTES.LdSvBasedHoUtranB1ThdEcn0,
			data[j].ATTRIBUTES.LdSvBasedHoUtranB1ThdRscp,
			data[j].ATTRIBUTES.HSInterRatHoUtranB1TimeTrig,	

		)
		checkErr(err)
		//fmt.Println("[+] Interrathoutrangroup data has been saved")
	}
	tx.Commit()
}


func insertInterratpolicycfggroup(eNodeBId string, baseName string, data []model.Interratpolicycfggroup) {
	fmt.Println("[+] Processing Interratpolicycfggroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `interratpolicycfggroup` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.InterRatPolicyCfgGroupId,
			data[j].ATTRIBUTES.UtranHoCfg,
			data[j].ATTRIBUTES.GeranGsmHoCfg,
			data[j].ATTRIBUTES.GeranGprsEdgeHoCfg,
			data[j].ATTRIBUTES.Cdma1xRttHoCfg,
			data[j].ATTRIBUTES.CdmaHrpdHoCfg,
			data[j].ATTRIBUTES.NoHoFlag,
			data[j].ATTRIBUTES.NoFastAnrFlag,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Interratpolicycfggroup data has been saved")
	}
	tx.Commit()
}


func insertIntrafreqhogroup(eNodeBId string, baseName string, data []model.Intrafreqhogroup) {
	fmt.Println("[+] Processing Intrafreqhogroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `intrafreqhogroup` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.IntraFreqHoGroupId,
			data[j].ATTRIBUTES.IntraFreqHoA3Hyst,
			data[j].ATTRIBUTES.IntraFreqHoA3Offset,
			data[j].ATTRIBUTES.IntraFreqHoA3TimeToTrig,
			data[j].ATTRIBUTES.HighSpeedA3TimeToTrig,
			data[j].ATTRIBUTES.IntraFreqHoA2ThldRsrpCE,	

		)
		checkErr(err)
		//fmt.Println("[+] Intrafreqhogroup data has been saved")
	}
	tx.Commit()
}


func insertIntrarathocomm(eNodeBId string, baseName string, data []model.Intrarathocomm) {
	fmt.Println("[+] Processing Intrarathocomm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `intrarathocomm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IntraRatHoMaxRprtCell,
			data[j].ATTRIBUTES.IntraRatHoRprtAmount,
			data[j].ATTRIBUTES.IntraFreqHoA3TrigQuan,
			data[j].ATTRIBUTES.IntraFreqHoA3RprtQuan,
			data[j].ATTRIBUTES.IntraFreqHoRprtInterval,
			data[j].ATTRIBUTES.InterFreqHoA4RprtQuan,
			data[j].ATTRIBUTES.InterFreqHoRprtInterval,
			data[j].ATTRIBUTES.InterFreqHoA1A2TrigQuan,
			data[j].ATTRIBUTES.FreqPriInterFreqHoA1TrigQuan,
			data[j].ATTRIBUTES.InterFreqHoA4TrigQuan,
			data[j].ATTRIBUTES.CovBasedIfHoWaitingTimer,
			data[j].ATTRIBUTES.A3InterFreqHoA1A2TrigQuan,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.FreqPriHoCandidateUeSelPer,	

		)
		checkErr(err)
		//fmt.Println("[+] Intrarathocomm data has been saved")
	}
	tx.Commit()
}


func insertIopscfg(eNodeBId string, baseName string, data []model.Iopscfg) {
	fmt.Println("[+] Processing Iopscfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `iopscfg` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IopsSwitch,
			data[j].ATTRIBUTES.EnterIopsThd,
			data[j].ATTRIBUTES.ExitIopsThd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Iopscfg data has been saved")
	}
	tx.Commit()
}


func insertIpclklnk(eNodeBId string, baseName string, data []model.Ipclklnk) {
	fmt.Println("[+] Processing Ipclklnk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ipclklnk` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LN,
			data[j].ATTRIBUTES.ICPT,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.CIP,
			data[j].ATTRIBUTES.SIP,
			data[j].ATTRIBUTES.DM,
			data[j].ATTRIBUTES.DELAYTYPE,
			data[j].ATTRIBUTES.CLASSIDENTIFY,
			data[j].ATTRIBUTES.CLASS0,
			data[j].ATTRIBUTES.CLASS1,
			data[j].ATTRIBUTES.CLASS2,
			data[j].ATTRIBUTES.CLASS3,
			data[j].ATTRIBUTES.CNM,
			data[j].ATTRIBUTES.CMPST,
			data[j].ATTRIBUTES.PRI,
			data[j].ATTRIBUTES.ANNFREQ,
			data[j].ATTRIBUTES.SYNCFREQ,
			data[j].ATTRIBUTES.MASTERPRIO,
			data[j].ATTRIBUTES.PROFILETYPE,
			data[j].ATTRIBUTES.NEGDURATION,
			data[j].ATTRIBUTES.IPMODE,
			data[j].ATTRIBUTES.PREHOPIP,
			data[j].ATTRIBUTES.PDELAYREQFREQ,
			data[j].ATTRIBUTES.DEVTYPE,
			data[j].ATTRIBUTES.DSTMLTMACTYPE,
			data[j].ATTRIBUTES.DELAYFREQ,
			data[j].ATTRIBUTES.IPSYNCMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Ipclklnk data has been saved")
	}
	tx.Commit()
}


func insertIpguard(eNodeBId string, baseName string, data []model.Ipguard) {
	fmt.Println("[+] Processing Ipguard data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ipguard` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ARPSPOOFCHKSW,
			data[j].ATTRIBUTES.ARPLRNSTRICTSW,
			data[j].ATTRIBUTES.INVALIDPKTCHKSW,
			data[j].ATTRIBUTES.IPSECREPLAYCHKSW,
			data[j].ATTRIBUTES.REDIRECTDOSCHKSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Ipguard data has been saved")
	}
	tx.Commit()
}


func insertIppath(eNodeBId string, baseName string, data []model.Ippath) {
	fmt.Println("[+] Processing Ippath data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ippath` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PATHID,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PATHTYPE,
			data[j].ATTRIBUTES.DSCP,
			data[j].ATTRIBUTES.LOCALIP,
			data[j].ATTRIBUTES.PEERIP,
			data[j].ATTRIBUTES.QT,
			data[j].ATTRIBUTES.IPMUXSWITCH,
			data[j].ATTRIBUTES.DESCRI,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.JNRSCGRP,
			data[j].ATTRIBUTES.STATICCHK,	

		)
		checkErr(err)
		//fmt.Println("[+] Ippath data has been saved")
	}
	tx.Commit()
}


func insertIppmsession(eNodeBId string, baseName string, data []model.Ippmsession) {
	fmt.Println("[+] Processing Ippmsession data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ippmsession` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IPPMSN,
			data[j].ATTRIBUTES.LOCALIP,
			data[j].ATTRIBUTES.PEERIP,
			data[j].ATTRIBUTES.IPPMDSCP,
			data[j].ATTRIBUTES.DIR,
			data[j].ATTRIBUTES.IPPMTYPE,
			data[j].ATTRIBUTES.BINDPATH,
			data[j].ATTRIBUTES.ASSOCIATEUPPATHPLR,
			data[j].ATTRIBUTES.PATHID,	

		)
		checkErr(err)
		//fmt.Println("[+] Ippmsession data has been saved")
	}
	tx.Commit()
}


func insertIprt(eNodeBId string, baseName string, data []model.Iprt) {
	fmt.Println("[+] Processing Iprt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `iprt` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RTIDX,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.RTTYPE,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.DSTIP,
			data[j].ATTRIBUTES.DSTMASK,
			data[j].ATTRIBUTES.NEXTHOP,
			data[j].ATTRIBUTES.MTUSWITCH,
			data[j].ATTRIBUTES.PREF,
			data[j].ATTRIBUTES.DESCRI,	

		)
		checkErr(err)
		//fmt.Println("[+] Iprt data has been saved")
	}
	tx.Commit()
}


func insertIratncellclassmgt(eNodeBId string, baseName string, data []model.Iratncellclassmgt) {
	fmt.Println("[+] Processing Iratncellclassmgt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `iratncellclassmgt` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RatType,
			data[j].ATTRIBUTES.StatPeriodForNcellClass,
			data[j].ATTRIBUTES.NCellMeasNumThd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Iratncellclassmgt data has been saved")
	}
	tx.Commit()
}


func insertIub(eNodeBId string, baseName string, data []model.Iub) {
	fmt.Println("[+] Processing Iub data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `iub` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IUBID,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.UPEPGROUPID,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.STATICCHKSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Iub data has been saved")
	}
	tx.Commit()
}


func insertIubcp(eNodeBId string, baseName string, data []model.Iubcp) {
	fmt.Println("[+] Processing Iubcp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `iubcp` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CPPT,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.CPPN,
			data[j].ATTRIBUTES.CPBEARID,
			data[j].ATTRIBUTES.BELONG,	

		)
		checkErr(err)
		//fmt.Println("[+] Iubcp data has been saved")
	}
	tx.Commit()
}


func insertLansw(eNodeBId string, baseName string, data []model.Lansw) {
	fmt.Println("[+] Processing Lansw data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lansw` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.BCPKTRATETHD,
			data[j].ATTRIBUTES.MACAGINGTIME,
			data[j].ATTRIBUTES.DYNAMICMACLRNSW,
			data[j].ATTRIBUTES.SRCMACATTACKCHKSW,
			data[j].ATTRIBUTES.SRCMACATTACKALMTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Lansw data has been saved")
	}
	tx.Commit()
}


func insertLanswitchport(eNodeBId string, baseName string, data []model.Lanswitchport) {
	fmt.Println("[+] Processing Lanswitchport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lanswitchport` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PORTIDX,
			data[j].ATTRIBUTES.PORTTYPE,
			data[j].ATTRIBUTES.USERLABLE,	

		)
		checkErr(err)
		//fmt.Println("[+] Lanswitchport data has been saved")
	}
	tx.Commit()
}


func insertLicensecontrolstrategy(eNodeBId string, baseName string, data []model.Licensecontrolstrategy) {
	fmt.Println("[+] Processing Licensecontrolstrategy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `licensecontrolstrategy` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SmoothUpgradeTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Licensecontrolstrategy data has been saved")
	}
	tx.Commit()
}


func insertLicratio(eNodeBId string, baseName string, data []model.Licratio) {
	fmt.Println("[+] Processing Licratio data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `licratio` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UpLicRatio,
			data[j].ATTRIBUTES.TrafficSharingType,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Licratio data has been saved")
	}
	tx.Commit()
}


func insertLineclk(eNodeBId string, baseName string, data []model.Lineclk) {
	fmt.Println("[+] Processing Lineclk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lineclk` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LN,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PRI,	

		)
		checkErr(err)
		//fmt.Println("[+] Lineclk data has been saved")
	}
	tx.Commit()
}


func insertLioptatomrule(eNodeBId string, baseName string, data []model.Lioptatomrule) {
	fmt.Println("[+] Processing Lioptatomrule data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lioptatomrule` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AtomRuleID,
			data[j].ATTRIBUTES.MeasureObjType,
			data[j].ATTRIBUTES.ConditionType,
			data[j].ATTRIBUTES.ThresholdforNumPara,
			data[j].ATTRIBUTES.MeasureObject,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Lioptatomrule data has been saved")
	}
	tx.Commit()
}


func insertLioptfeature(eNodeBId string, baseName string, data []model.Lioptfeature) {
	fmt.Println("[+] Processing Lioptfeature data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lioptfeature` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IOptFeatureID,
			data[j].ATTRIBUTES.IOptFeatureName,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Lioptfeature data has been saved")
	}
	tx.Commit()
}


func insertLioptfunction(eNodeBId string, baseName string, data []model.Lioptfunction) {
	fmt.Println("[+] Processing Lioptfunction data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lioptfunction` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IOptFunctionID,
			data[j].ATTRIBUTES.IOptFunctionName,
			data[j].ATTRIBUTES.IOptFeatureID,
			data[j].ATTRIBUTES.MeasureObjType,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Lioptfunction data has been saved")
	}
	tx.Commit()
}


func insertLioptrule(eNodeBId string, baseName string, data []model.Lioptrule) {
	fmt.Println("[+] Processing Lioptrule data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lioptrule` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RuleID,
			data[j].ATTRIBUTES.ActionType,
			data[j].ATTRIBUTES.AtomRuleRelationType,
			data[j].ATTRIBUTES.Period,
			data[j].ATTRIBUTES.Action,
			data[j].ATTRIBUTES.IOptFunctionID,
			data[j].ATTRIBUTES.PenaltyTime,
			data[j].ATTRIBUTES.AdaptiveRAT,
			data[j].ATTRIBUTES.ActiveStatus,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Lioptrule data has been saved")
	}
	tx.Commit()
}


func insertLioptrulemember(eNodeBId string, baseName string, data []model.Lioptrulemember) {
	fmt.Println("[+] Processing Lioptrulemember data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lioptrulemember` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RuleID,
			data[j].ATTRIBUTES.AtomRuleID,
			data[j].ATTRIBUTES.ActiveStatus,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Lioptrulemember data has been saved")
	}
	tx.Commit()
}


func insertLldpglobal(eNodeBId string, baseName string, data []model.Lldpglobal) {
	fmt.Println("[+] Processing Lldpglobal data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lldpglobal` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TXINTVAL,
			data[j].ATTRIBUTES.HOLDMULTI,
			data[j].ATTRIBUTES.REINITDELAY,
			data[j].ATTRIBUTES.DELAY,
			data[j].ATTRIBUTES.NOTIFYSW,
			data[j].ATTRIBUTES.NOTIFYINTERVAL,	

		)
		checkErr(err)
		//fmt.Println("[+] Lldpglobal data has been saved")
	}
	tx.Commit()
}


func insertLocalethport(eNodeBId string, baseName string, data []model.Localethport) {
	fmt.Println("[+] Processing Localethport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `localethport` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SWITCH,
			data[j].ATTRIBUTES.GRATUITOUSARPSW,
			data[j].ATTRIBUTES.IP6SW,	

		)
		checkErr(err)
		//fmt.Println("[+] Localethport data has been saved")
	}
	tx.Commit()
}


func insertLocalip(eNodeBId string, baseName string, data []model.Localip) {
	fmt.Println("[+] Processing Localip data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `localip` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IP,
			data[j].ATTRIBUTES.MASK,	

		)
		checkErr(err)
		//fmt.Println("[+] Localip data has been saved")
	}
	tx.Commit()
}


func insertLocalip6(eNodeBId string, baseName string, data []model.Localip6) {
	fmt.Println("[+] Processing Localip6 data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `localip6` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IP6,
			data[j].ATTRIBUTES.PFXLEN,	

		)
		checkErr(err)
		//fmt.Println("[+] Localip6 data has been saved")
	}
	tx.Commit()
}


func insertLocalwap(eNodeBId string, baseName string, data []model.Localwap) {
	fmt.Println("[+] Processing Localwap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `localwap` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SWITCH,
			data[j].ATTRIBUTES.SSID,
			data[j].ATTRIBUTES.SHOWSSID,
			data[j].ATTRIBUTES.AUTODISABLETIME,
			data[j].ATTRIBUTES.REGIONCODE,
			data[j].ATTRIBUTES.CHANNEL,	

		)
		checkErr(err)
		//fmt.Println("[+] Localwap data has been saved")
	}
	tx.Commit()
}


func insertLocation(eNodeBId string, baseName string, data []model.Location) {
	fmt.Println("[+] Processing Location data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `location` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ADDRESS,
			data[j].ATTRIBUTES.ALTITUDE,
			data[j].ATTRIBUTES.CITY,
			data[j].ATTRIBUTES.CONTACT,
			data[j].ATTRIBUTES.GCDF,
			data[j].ATTRIBUTES.LATITUDEDEGFORMAT,
			data[j].ATTRIBUTES.LOCATIONID,
			data[j].ATTRIBUTES.LOCATIONNAME,
			data[j].ATTRIBUTES.LONGITUDEDEGFORMAT,
			data[j].ATTRIBUTES.OFFICE,
			data[j].ATTRIBUTES.RANGE,
			data[j].ATTRIBUTES.REGION,
			data[j].ATTRIBUTES.TELEPHONE,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.PRECISE,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.LOCATIONTYPE,	

		)
		checkErr(err)
		//fmt.Println("[+] Location data has been saved")
	}
	tx.Commit()
}


func insertLwamgtcfg(eNodeBId string, baseName string, data []model.Lwamgtcfg) {
	fmt.Println("[+] Processing Lwamgtcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `lwamgtcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.LteWlanAggrW1ThdRssi,
			data[j].ATTRIBUTES.LteWlanAggrW2Thd1Rssi,
			data[j].ATTRIBUTES.LteWlanAggrW3ThdRssi,
			data[j].ATTRIBUTES.LwaW1TimeToTrigger,
			data[j].ATTRIBUTES.LwaW2TimeToTrigger,
			data[j].ATTRIBUTES.LwaW3TimeToTrigger,
			data[j].ATTRIBUTES.WlanMeasHyst,
			data[j].ATTRIBUTES.CellLwaAlgoSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Lwamgtcfg data has been saved")
	}
	tx.Commit()
}


func insertMainsalarmbind(eNodeBId string, baseName string, data []model.Mainsalarmbind) {
	fmt.Println("[+] Processing Mainsalarmbind data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mainsalarmbind` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ISDSWITCH,
			data[j].ATTRIBUTES.NMSACN,
			data[j].ATTRIBUTES.NMSASRN,
			data[j].ATTRIBUTES.NMSASN,
			data[j].ATTRIBUTES.NMSAPN,	

		)
		checkErr(err)
		//fmt.Println("[+] Mainsalarmbind data has been saved")
	}
	tx.Commit()
}


func insertManresalmrpt(eNodeBId string, baseName string, data []model.Manresalmrpt) {
	fmt.Println("[+] Processing Manresalmrpt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `manresalmrpt` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Manresalmrpt data has been saved")
	}
	tx.Commit()
}


func insertMbmspara(eNodeBId string, baseName string, data []model.Mbmspara) {
	fmt.Println("[+] Processing Mbmspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mbmspara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SyncPeriod,
			data[j].ATTRIBUTES.MbmsCtrlSwitch,
			data[j].ATTRIBUTES.M2DisconnProtectTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Mbmspara data has been saved")
	}
	tx.Commit()
}


func insertMmefeaturecfg(eNodeBId string, baseName string, data []model.Mmefeaturecfg) {
	fmt.Println("[+] Processing Mmefeaturecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mmefeaturecfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.S1InterfaceId,
			data[j].ATTRIBUTES.MdtEnable,
			data[j].ATTRIBUTES.CtrlMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Mmefeaturecfg data has been saved")
	}
	tx.Commit()
}


func insertMpt(eNodeBId string, baseName string, data []model.Mpt) {
	fmt.Println("[+] Processing Mpt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mpt` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.TYPE,
			data[j].ATTRIBUTES.SWITCH,
			data[j].ATTRIBUTES.BRDSPEC,
			data[j].ATTRIBUTES.MPTWORKMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Mpt data has been saved")
	}
	tx.Commit()
}


func insertMptresassignment(eNodeBId string, baseName string, data []model.Mptresassignment) {
	fmt.Println("[+] Processing Mptresassignment data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mptresassignment` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.MasterMptAssignmentMode,
			data[j].ATTRIBUTES.SlaveMptAssignmentMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Mptresassignment data has been saved")
	}
	tx.Commit()
}


func insertMro(eNodeBId string, baseName string, data []model.Mro) {
	fmt.Println("[+] Processing Mro data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `mro` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OptPeriod,
			data[j].ATTRIBUTES.NcellOptThd,
			data[j].ATTRIBUTES.StatNumThd,
			data[j].ATTRIBUTES.PingpongTimeThd,
			data[j].ATTRIBUTES.PingpongRatioThd,
			data[j].ATTRIBUTES.CoverAbnormalThd,
			data[j].ATTRIBUTES.ServingRsrpThd,
			data[j].ATTRIBUTES.NeighborRsrpThd,
			data[j].ATTRIBUTES.UePingPongNumThd,
			data[j].ATTRIBUTES.InterFreqMeasTooLateHoThd,
			data[j].ATTRIBUTES.InterFreqA2RollBackThd,
			data[j].ATTRIBUTES.MroOptMode,
			data[j].ATTRIBUTES.UnnecInterRatHoOptThd,
			data[j].ATTRIBUTES.UnnecInterRatHoRsrpThd,
			data[j].ATTRIBUTES.InterRatA2RsrpMinThd,
			data[j].ATTRIBUTES.InterRatStatNumThd,
			data[j].ATTRIBUTES.IntraRatTooEarlyHoRatioThd,
			data[j].ATTRIBUTES.IntraRatTooLateHoRatioThd,
			data[j].ATTRIBUTES.IntraRatAbnormalRatioThd,
			data[j].ATTRIBUTES.InterFreqA2RollBackPeriod,
			data[j].ATTRIBUTES.IntraRatHoTooEarlyTimeThd,
			data[j].ATTRIBUTES.InterRatAbnormalHoRatioThd,
			data[j].ATTRIBUTES.InterRatMeasTooLateHoThd,
			data[j].ATTRIBUTES.UnnecInterRatHoRatioThd,
			data[j].ATTRIBUTES.UnnecInterRatHoMeasTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Mro data has been saved")
	}
	tx.Commit()
}


func insertNbcelldlschcealgo(eNodeBId string, baseName string, data []model.Nbcelldlschcealgo) {
	fmt.Println("[+] Processing Nbcelldlschcealgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nbcelldlschcealgo` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.DlInitialTransRptCount,
			data[j].ATTRIBUTES.DlInitialMcs,
			data[j].ATTRIBUTES.UuMessageWaitingTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Nbcelldlschcealgo data has been saved")
	}
	tx.Commit()
}


func insertNbcellulschcealgo(eNodeBId string, baseName string, data []model.Nbcellulschcealgo) {
	fmt.Println("[+] Processing Nbcellulschcealgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nbcellulschcealgo` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.CoverageLevel,
			data[j].ATTRIBUTES.UlInitialMcs,
			data[j].ATTRIBUTES.UlInitialTransRptCount,
			data[j].ATTRIBUTES.AckNackTransRptCount,
			data[j].ATTRIBUTES.AckNackTransRptCountMsg4,
			data[j].ATTRIBUTES.NbLogicChSrProhibitTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Nbcellulschcealgo data has been saved")
	}
	tx.Commit()
}


func insertNcellclassmgt(eNodeBId string, baseName string, data []model.Ncellclassmgt) {
	fmt.Println("[+] Processing Ncellclassmgt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ncellclassmgt` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.StatPeriodForNcellClass,
			data[j].ATTRIBUTES.HoAttemptThd,
			data[j].ATTRIBUTES.CaSCellCfgThd,
			data[j].ATTRIBUTES.HoSuccThd,
			data[j].ATTRIBUTES.IntraRatNcellMgtMode,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Ncellclassmgt data has been saved")
	}
	tx.Commit()
}


func insertNcelldlrsrpmeaspara(eNodeBId string, baseName string, data []model.Ncelldlrsrpmeaspara) {
	fmt.Println("[+] Processing Ncelldlrsrpmeaspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ncelldlrsrpmeaspara` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlRsrpAutoNCellMeasSwitch,
			data[j].ATTRIBUTES.NCellDlRsrpMeasA3Offset,	

		)
		checkErr(err)
		//fmt.Println("[+] Ncelldlrsrpmeaspara data has been saved")
	}
	tx.Commit()
}


func insertNcellparacfg(eNodeBId string, baseName string, data []model.Ncellparacfg) {
	fmt.Println("[+] Processing Ncellparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ncellparacfg` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.RatType,
			data[j].ATTRIBUTES.NCellOdDisThd,
			data[j].ATTRIBUTES.HoStatThd,
			data[j].ATTRIBUTES.HoSuccThd,
			data[j].ATTRIBUTES.NcellNumForAnr,	

		)
		checkErr(err)
		//fmt.Println("[+] Ncellparacfg data has been saved")
	}
	tx.Commit()
}


func insertNcellsrsmeaspara(eNodeBId string, baseName string, data []model.Ncellsrsmeaspara) {
	fmt.Println("[+] Processing Ncellsrsmeaspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ncellsrsmeaspara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SrsAutoNCellMeasSwitch,
			data[j].ATTRIBUTES.NCellSrsMeasA3Offset,
			data[j].ATTRIBUTES.NCellMeasSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Ncellsrsmeaspara data has been saved")
	}
	tx.Commit()
}


func insertNe(eNodeBId string, baseName string, data []model.Ne) {
	fmt.Println("[+] Processing Ne data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ne` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LOCATION,
			data[j].ATTRIBUTES.SWVERSION,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.NERMVERSION,
			data[j].ATTRIBUTES.INTERFACEID,
			data[j].ATTRIBUTES.PRODUCTVERSION,
			data[j].ATTRIBUTES.LMTVERSION,
			data[j].ATTRIBUTES.DID,
			data[j].ATTRIBUTES.SITENAME,
			data[j].ATTRIBUTES.NENAME,
			data[j].ATTRIBUTES.HOTPATCHVERSION,
			data[j].ATTRIBUTES.CLOUDBBID,	

		)
		checkErr(err)
		//fmt.Println("[+] Ne data has been saved")
	}
	tx.Commit()
}


func insertNemnt(eNodeBId string, baseName string, data []model.Nemnt) {
	fmt.Println("[+] Processing Nemnt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nemnt` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.MNTMODE,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.ET,
			data[j].ATTRIBUTES.MMSETREMARK,	

		)
		checkErr(err)
		//fmt.Println("[+] Nemnt data has been saved")
	}
	tx.Commit()
}


func insertNode(eNodeBId string, baseName string, data []model.Node) {
	fmt.Println("[+] Processing Node data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `node` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PRODUCTTYPE,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.NERMVERSION,
			data[j].ATTRIBUTES.NODEID,
			data[j].ATTRIBUTES.NODENAME,
			data[j].ATTRIBUTES.WM,
			data[j].ATTRIBUTES.SWVERSION,
			data[j].ATTRIBUTES.HOTPATCHVERSION,
			data[j].ATTRIBUTES.PRODUCTVERSION,
			data[j].ATTRIBUTES.INTERFACEID,
			data[j].ATTRIBUTES.LMTVERSION,
			data[j].ATTRIBUTES.APPVERSION,
			data[j].ATTRIBUTES.APPHOTPATCHVERSION,
			data[j].ATTRIBUTES.LTEMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Node data has been saved")
	}
	tx.Commit()
}


func insertNodebalgpara(eNodeBId string, baseName string, data []model.Nodebalgpara) {
	fmt.Println("[+] Processing Nodebalgpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebalgpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.CERSVFOR2MSUSER,
			data[j].ATTRIBUTES.CCPICPHASESWCTRL,
			data[j].ATTRIBUTES.HSDPASCHPOOLSW,
			data[j].ATTRIBUTES.PWRCHKSW,
			data[j].ATTRIBUTES.CCPICPH2SIRTARMODSW,
			data[j].ATTRIBUTES.ADPRETRANSSW,
			data[j].ATTRIBUTES.FRTRJT,
			data[j].ATTRIBUTES.INTERBOARDICSW,
			data[j].ATTRIBUTES.SLEEPINGCELLDETECTSW,
			data[j].ATTRIBUTES.UTRPCEXTCNBAPCAP,
			data[j].ATTRIBUTES.RTWPSCALEPRECISIONOPT,
			data[j].ATTRIBUTES.CEDETECTPERIODLEVEL,
			data[j].ATTRIBUTES.ULRESWORKMODE,
			data[j].ATTRIBUTES.CEIMPROVEMENT2MSSW,
			data[j].ATTRIBUTES.INITSINGLEHARQSW,
			data[j].ATTRIBUTES.ULHITHPMACESW,
			data[j].ATTRIBUTES.ULINNERPCABNRISECTRLSW,
			data[j].ATTRIBUTES.KPIFAULTDETECTSW,
			data[j].ATTRIBUTES.KPIFAULTAUTORCVRMTHD,
			data[j].ATTRIBUTES.FIXEDCERSVFOR2MS,
			data[j].ATTRIBUTES.DLTXPWRMEASPERIOD,
			data[j].ATTRIBUTES.PWRPREDICTSWFORCAC,
			data[j].ATTRIBUTES.DLFPWINADJSW,
			data[j].ATTRIBUTES.HSPAMCINTERBRDSCHSW,
			data[j].ATTRIBUTES.R99FAULTTHD,
			data[j].ATTRIBUTES.UPAFAULTTHD,
			data[j].ATTRIBUTES.PROCUNITFAULTTHD,
			data[j].ATTRIBUTES.PROTECTIONTHD,
			data[j].ATTRIBUTES.PROCUNITFAULTCLEARTHD,
			data[j].ATTRIBUTES.ALMOCCURPRD,
			data[j].ATTRIBUTES.ALMCLEARPRD,
			data[j].ATTRIBUTES.PDMEASMSW,
			data[j].ATTRIBUTES.ULCCHOLPCMINSIRTARGET,
			data[j].ATTRIBUTES.HSUPASCHPOOLSW,
			data[j].ATTRIBUTES.ULCOVEXPCFGMOD,
			data[j].ATTRIBUTES.PWRCFGPROTECT,
			data[j].ATTRIBUTES.UPCEIMPROVEFORAMRWBSW,
			data[j].ATTRIBUTES.NBISADAPTIVESW,
			data[j].ATTRIBUTES.UX2SONSETUPSW,
			data[j].ATTRIBUTES.SLEEPINGCELLALARMSW,
			data[j].ATTRIBUTES.SLEEPINGCELLDETECTOPTION,
			data[j].ATTRIBUTES.RFFAULTCTRSW,
			data[j].ATTRIBUTES.CPCFDPCHDTXSW,
			data[j].ATTRIBUTES.IUBCONGESTKPIIMPROVESW,
			data[j].ATTRIBUTES.HSDPASERCELLCHOPTSW,
			data[j].ATTRIBUTES.CELLRTWPEXCESSALMSW,
			data[j].ATTRIBUTES.CELLRTWPEXCESSALMTHD,
			data[j].ATTRIBUTES.VAMPWRSAVSW,
			data[j].ATTRIBUTES.VAMPWRSAVSTARTTIME,
			data[j].ATTRIBUTES.VAMPWRSAVSTOPTIME,
			data[j].ATTRIBUTES.VAMPWRSAVTXPWRLOWTHD,
			data[j].ATTRIBUTES.VAMPWRSAVTXPWRHIGHTHD,
			data[j].ATTRIBUTES.DLAUTORESHUFFLESW,
			data[j].ATTRIBUTES.ADPETADJFLAG,
			data[j].ATTRIBUTES.AMRTRANSBWCOMP,
			data[j].ATTRIBUTES.SFSSWITCH,
			data[j].ATTRIBUTES.KPIFAULTDETECT,
			data[j].ATTRIBUTES.RRUINTELDORFEQBAND,
			data[j].ATTRIBUTES.OPERATORSHARECEMODE,
			data[j].ATTRIBUTES.OPERATORINDEXSELECTION,
			data[j].ATTRIBUTES.NBISRESELPERIOD,
			data[j].ATTRIBUTES.RTWPIMBALANCEALMTHD,
			data[j].ATTRIBUTES.NBISINIACTTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebalgpara data has been saved")
	}
	tx.Commit()
}


func insertNodebbbres(eNodeBId string, baseName string, data []model.Nodebbbres) {
	fmt.Println("[+] Processing Nodebbbres data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebbbres` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.NODEBBBRESID,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.CAPACITY,
			data[j].ATTRIBUTES.ADMSTATE,
			data[j].ATTRIBUTES.HCE,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebbbres data has been saved")
	}
	tx.Commit()
}


func insertNodebchrlevel(eNodeBId string, baseName string, data []model.Nodebchrlevel) {
	fmt.Println("[+] Processing Nodebchrlevel data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebchrlevel` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.COMMCHREVENTSW,
			data[j].ATTRIBUTES.USERCHREVENTSW,
			data[j].ATTRIBUTES.COMMCHRPRDEVENTSW,
			data[j].ATTRIBUTES.USERCHRPRDEVENTSW,
			data[j].ATTRIBUTES.HSDPAKQIDIAGTHPTHD,
			data[j].ATTRIBUTES.HSUPAKQIDIAGTHPTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebchrlevel data has been saved")
	}
	tx.Commit()
}


func insertNodebclspatimer(eNodeBId string, baseName string, data []model.Nodebclspatimer) {
	fmt.Println("[+] Processing Nodebclspatimer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebclspatimer` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LOWERLIMIT,
			data[j].ATTRIBUTES.UPPERLIMIT,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebclspatimer data has been saved")
	}
	tx.Commit()
}


func insertNodebfunction(eNodeBId string, baseName string, data []model.Nodebfunction) {
	fmt.Println("[+] Processing Nodebfunction data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebfunction` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.NODEBFUNCTIONNAME,
			data[j].ATTRIBUTES.APPLICATIONREF,
			data[j].ATTRIBUTES.NERMVERSION,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.NODEBID,
			data[j].ATTRIBUTES.PRODUCTVERSION,
			data[j].ATTRIBUTES.INTERFACEID,
			data[j].ATTRIBUTES.LMTVERSION,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebfunction data has been saved")
	}
	tx.Commit()
}


func insertNodeblicensealmthd(eNodeBId string, baseName string, data []model.Nodeblicensealmthd) {
	fmt.Println("[+] Processing Nodeblicensealmthd data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodeblicensealmthd` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.OTHD,
			data[j].ATTRIBUTES.OPRD,
			data[j].ATTRIBUTES.RTHD,
			data[j].ATTRIBUTES.RPRD,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodeblicensealmthd data has been saved")
	}
	tx.Commit()
}


func insertNodebmulticellgrp(eNodeBId string, baseName string, data []model.Nodebmulticellgrp) {
	fmt.Println("[+] Processing Nodebmulticellgrp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebmulticellgrp` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.MULTICELLGRPID,
			data[j].ATTRIBUTES.MULTICELLGRPTYPE,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.ULOCELLREF,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebmulticellgrp data has been saved")
	}
	tx.Commit()
}


func insertNodeboptdynadjpara(eNodeBId string, baseName string, data []model.Nodeboptdynadjpara) {
	fmt.Println("[+] Processing Nodeboptdynadjpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodeboptdynadjpara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DYNADJSW,
			data[j].ATTRIBUTES.DYNADJSTARTTIME,
			data[j].ATTRIBUTES.DYNADJENDTIME,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodeboptdynadjpara data has been saved")
	}
	tx.Commit()
}


func insertNodebpoweroutage(eNodeBId string, baseName string, data []model.Nodebpoweroutage) {
	fmt.Println("[+] Processing Nodebpoweroutage data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebpoweroutage` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ISDSW,
			data[j].ATTRIBUTES.BAKPWRSAVPOLICY,
			data[j].ATTRIBUTES.PHASE1PERIOD,
			data[j].ATTRIBUTES.PHASE2PERIOD,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebpoweroutage data has been saved")
	}
	tx.Commit()
}


func insertNodebresallocrule(eNodeBId string, baseName string, data []model.Nodebresallocrule) {
	fmt.Println("[+] Processing Nodebresallocrule data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebresallocrule` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RULE,
			data[j].ATTRIBUTES.SW,
			data[j].ATTRIBUTES.CSUSERNUM,
			data[j].ATTRIBUTES.PSUSERNUM,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.RESREAVESW,
			data[j].ATTRIBUTES.MULTICELLREBUILDSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebresallocrule data has been saved")
	}
	tx.Commit()
}


func insertNodebrsvdpara(eNodeBId string, baseName string, data []model.Nodebrsvdpara) {
	fmt.Println("[+] Processing Nodebrsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebrsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.NODEBRSVDPARA1,
			data[j].ATTRIBUTES.NODEBRSVDPARA2,
			data[j].ATTRIBUTES.NODEBRSVDPARA3,
			data[j].ATTRIBUTES.NODEBRSVDPARA4,
			data[j].ATTRIBUTES.NODEBRSVDPARA5,
			data[j].ATTRIBUTES.NODEBRSVDPARA6,
			data[j].ATTRIBUTES.NODEBRSVDPARA7,
			data[j].ATTRIBUTES.NODEBRSVDPARA8,
			data[j].ATTRIBUTES.NODEBRSVDPARA9,
			data[j].ATTRIBUTES.NODEBRSVDPARA10,
			data[j].ATTRIBUTES.NODEBRSVDPARA11,
			data[j].ATTRIBUTES.NODEBRSVDPARA12,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.NODEBRSVDPARA13,
			data[j].ATTRIBUTES.NODEBRSVDPARA14,
			data[j].ATTRIBUTES.NODEBRSVDPARA15,
			data[j].ATTRIBUTES.NODEBRSVDPARA16,
			data[j].ATTRIBUTES.NODEBRSVDPARA17,
			data[j].ATTRIBUTES.NODEBRSVDPARA18,
			data[j].ATTRIBUTES.NODEBRSVDPARA19,
			data[j].ATTRIBUTES.NODEBRSVDPARA20,
			data[j].ATTRIBUTES.NODEBRSVDPARA21,
			data[j].ATTRIBUTES.NODEBRSVDPARA22,
			data[j].ATTRIBUTES.NODEBRSVDPARA23,
			data[j].ATTRIBUTES.NODEBRSVDPARA24,
			data[j].ATTRIBUTES.NODEBRSVDPARA25,
			data[j].ATTRIBUTES.NODEBRSVDPARA26,
			data[j].ATTRIBUTES.NODEBRSVDPARA27,
			data[j].ATTRIBUTES.NODEBRSVDPARA28,
			data[j].ATTRIBUTES.NODEBRSVDPARA29,
			data[j].ATTRIBUTES.NODEBRSVDPARA30,
			data[j].ATTRIBUTES.NODEBRSVDPARA31,
			data[j].ATTRIBUTES.NODEBRSVDPARA32,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebrsvdpara data has been saved")
	}
	tx.Commit()
}


func insertNodebruleactionpara(eNodeBId string, baseName string, data []model.Nodebruleactionpara) {
	fmt.Println("[+] Processing Nodebruleactionpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebruleactionpara` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RULEID,
			data[j].ATTRIBUTES.MOCNAME,
			data[j].ATTRIBUTES.PARANAME,
			data[j].ATTRIBUTES.ADMINSTATE,
			data[j].ATTRIBUTES.PARAINTERVALVALUE,
			data[j].ATTRIBUTES.PARAENUMVALUE,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebruleactionpara data has been saved")
	}
	tx.Commit()
}


func insertNodebsmthpwr(eNodeBId string, baseName string, data []model.Nodebsmthpwr) {
	fmt.Println("[+] Processing Nodebsmthpwr data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebsmthpwr` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SMTHPWRSWITCH,
			data[j].ATTRIBUTES.STEP,
			data[j].ATTRIBUTES.PERIOD,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebsmthpwr data has been saved")
	}
	tx.Commit()
}


func insertNodebtrfoverloadthd(eNodeBId string, baseName string, data []model.Nodebtrfoverloadthd) {
	fmt.Println("[+] Processing Nodebtrfoverloadthd data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebtrfoverloadthd` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TRAFFICOVERLOADOTHD,
			data[j].ATTRIBUTES.TRAFFICOVERLOADOPRD,
			data[j].ATTRIBUTES.TRAFFICOVERLOADRTHD,
			data[j].ATTRIBUTES.TRAFFICOVERLOADRPRD,
			data[j].ATTRIBUTES.OBJID,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebtrfoverloadthd data has been saved")
	}
	tx.Commit()
}


func insertNodebueqosenhance(eNodeBId string, baseName string, data []model.Nodebueqosenhance) {
	fmt.Println("[+] Processing Nodebueqosenhance data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `nodebueqosenhance` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULGOLDUEBPS,
			data[j].ATTRIBUTES.ULSILVERUEBPS,
			data[j].ATTRIBUTES.ULCOPPERUEBPS,
			data[j].ATTRIBUTES.DLGOLDUEBPS,
			data[j].ATTRIBUTES.DLSILVERUEBPS,
			data[j].ATTRIBUTES.DLCOPPERUEBPS,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.LOWTRAFFENHUSERNUMTHD,
			data[j].ATTRIBUTES.TBSIZEFOR10MSLOWTRAFFENH,
			data[j].ATTRIBUTES.TBSIZEFOR2MSLOWTRAFFENH,	

		)
		checkErr(err)
		//fmt.Println("[+] Nodebueqosenhance data has been saved")
	}
	tx.Commit()
}


func insertNtpcp(eNodeBId string, baseName string, data []model.Ntpcp) {
	fmt.Println("[+] Processing Ntpcp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ntpcp` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.IP,
			data[j].ATTRIBUTES.SYNCCYCLE,
			data[j].ATTRIBUTES.PORT,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.AUTHMODE,
			data[j].ATTRIBUTES.KEY,
			data[j].ATTRIBUTES.KEYID,
			data[j].ATTRIBUTES.MASTERFLAG,	

		)
		checkErr(err)
		//fmt.Println("[+] Ntpcp data has been saved")
	}
	tx.Commit()
}


func insertOmch(eNodeBId string, baseName string, data []model.Omch) {
	fmt.Println("[+] Processing Omch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `omch` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.FLAG,
			data[j].ATTRIBUTES.IP,
			data[j].ATTRIBUTES.MASK,
			data[j].ATTRIBUTES.PEERIP,
			data[j].ATTRIBUTES.PEERMASK,
			data[j].ATTRIBUTES.BEAR,
			data[j].ATTRIBUTES.BRT,
			data[j].ATTRIBUTES.CHECKTYPE,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.RTIDX,
			data[j].ATTRIBUTES.BINDSECONDARYRT,	

		)
		checkErr(err)
		//fmt.Println("[+] Omch data has been saved")
	}
	tx.Commit()
}


func insertOp(eNodeBId string, baseName string, data []model.Op) {
	fmt.Println("[+] Processing Op data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `op` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SWOP,
			data[j].ATTRIBUTES.LOCKST,	

		)
		checkErr(err)
		//fmt.Println("[+] Op data has been saved")
	}
	tx.Commit()
}


func insertOutport(eNodeBId string, baseName string, data []model.Outport) {
	fmt.Println("[+] Processing Outport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `outport` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.NAME,
			data[j].ATTRIBUTES.SW,	

		)
		checkErr(err)
		//fmt.Println("[+] Outport data has been saved")
	}
	tx.Commit()
}


func insertParaautooptcfg(eNodeBId string, baseName string, data []model.Paraautooptcfg) {
	fmt.Println("[+] Processing Paraautooptcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `paraautooptcfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PUSCHRsrpHighThd4AutoOpt,
			data[j].ATTRIBUTES.PUCCHPcSINROffset4AutoOpt,
			data[j].ATTRIBUTES.P0NominalPUCCH4AutoOpt,
			data[j].ATTRIBUTES.HOTimesThd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Paraautooptcfg data has been saved")
	}
	tx.Commit()
}


func insertPccfreqcfg(eNodeBId string, baseName string, data []model.Pccfreqcfg) {
	fmt.Println("[+] Processing Pccfreqcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pccfreqcfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PccDlEarfcn,
			data[j].ATTRIBUTES.PreferredPccPriority,
			data[j].ATTRIBUTES.PccA4RsrpThd,
			data[j].ATTRIBUTES.PccA4RsrqThd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Pccfreqcfg data has been saved")
	}
	tx.Commit()
}


func insertPcchcfg(eNodeBId string, baseName string, data []model.Pcchcfg) {
	fmt.Println("[+] Processing Pcchcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pcchcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DefaultPagingCycle,
			data[j].ATTRIBUTES.Nb,
			data[j].ATTRIBUTES.PagingSentNum,
			data[j].ATTRIBUTES.MaxPagingRecordsNum,
			data[j].ATTRIBUTES.PagingStrategy,
			data[j].ATTRIBUTES.EnhPagingCR,
			data[j].ATTRIBUTES.EnhPagingSwitch,
			data[j].ATTRIBUTES.DefaultPagingCycleForNb,
			data[j].ATTRIBUTES.NbForNbIoT,
			data[j].ATTRIBUTES.MaxNumRepetitionForPaging,	

		)
		checkErr(err)
		//fmt.Println("[+] Pcchcfg data has been saved")
	}
	tx.Commit()
}


func insertPdcprohcpara(eNodeBId string, baseName string, data []model.Pdcprohcpara) {
	fmt.Println("[+] Processing Pdcprohcpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pdcprohcpara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RohcSwitch,
			data[j].ATTRIBUTES.HighestMode,
			data[j].ATTRIBUTES.Profiles,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Pdcprohcpara data has been saved")
	}
	tx.Commit()
}


func insertPdschcfg(eNodeBId string, baseName string, data []model.Pdschcfg) {
	fmt.Println("[+] Processing Pdschcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pdschcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ReferenceSignalPwr,
			data[j].ATTRIBUTES.Pb,
			data[j].ATTRIBUTES.ReferenceSignalPwrMargin,
			data[j].ATTRIBUTES.TxPowerOffsetAnt0,
			data[j].ATTRIBUTES.TxPowerOffsetAnt1,
			data[j].ATTRIBUTES.TxPowerOffsetAnt2,
			data[j].ATTRIBUTES.TxPowerOffsetAnt3,
			data[j].ATTRIBUTES.TxChnPowerCfgSw,	

		)
		checkErr(err)
		//fmt.Println("[+] Pdschcfg data has been saved")
	}
	tx.Commit()
}


func insertPeerclk(eNodeBId string, baseName string, data []model.Peerclk) {
	fmt.Println("[+] Processing Peerclk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `peerclk` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PS,	

		)
		checkErr(err)
		//fmt.Println("[+] Peerclk data has been saved")
	}
	tx.Commit()
}


func insertPeu(eNodeBId string, baseName string, data []model.Peu) {
	fmt.Println("[+] Processing Peu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `peu` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,	

		)
		checkErr(err)
		//fmt.Println("[+] Peu data has been saved")
	}
	tx.Commit()
}


func insertPhichcfg(eNodeBId string, baseName string, data []model.Phichcfg) {
	fmt.Println("[+] Processing Phichcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `phichcfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PhichDuration,
			data[j].ATTRIBUTES.PhichResource,	

		)
		checkErr(err)
		//fmt.Println("[+] Phichcfg data has been saved")
	}
	tx.Commit()
}


func insertPhyport(eNodeBId string, baseName string, data []model.Phyport) {
	fmt.Println("[+] Processing Phyport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `phyport` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.ADMINISTRATIVESTATE,	

		)
		checkErr(err)
		//fmt.Println("[+] Phyport data has been saved")
	}
	tx.Commit()
}


func insertPlrthreshold(eNodeBId string, baseName string, data []model.Plrthreshold) {
	fmt.Println("[+] Processing Plrthreshold data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `plrthreshold` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PLRAT,
			data[j].ATTRIBUTES.PLRDT,	

		)
		checkErr(err)
		//fmt.Println("[+] Plrthreshold data has been saved")
	}
	tx.Commit()
}


func insertPmtucfg(eNodeBId string, baseName string, data []model.Pmtucfg) {
	fmt.Println("[+] Processing Pmtucfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pmtucfg` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.UDPPORT,
			data[j].ATTRIBUTES.DETECTPERIOD,
			data[j].ATTRIBUTES.AGINGPERIOD,
			data[j].ATTRIBUTES.DSCP,	

		)
		checkErr(err)
		//fmt.Println("[+] Pmtucfg data has been saved")
	}
	tx.Commit()
}


func insertPortpolicy(eNodeBId string, baseName string, data []model.Portpolicy) {
	fmt.Println("[+] Processing Portpolicy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `portpolicy` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.USBACCESSSECPOLICY,	

		)
		checkErr(err)
		//fmt.Println("[+] Portpolicy data has been saved")
	}
	tx.Commit()
}


func insertPri2que(eNodeBId string, baseName string, data []model.Pri2que) {
	fmt.Println("[+] Processing Pri2que data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pri2que` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.PRI0,
			data[j].ATTRIBUTES.PRI1,
			data[j].ATTRIBUTES.PRI2,
			data[j].ATTRIBUTES.PRI3,
			data[j].ATTRIBUTES.PRI4,
			data[j].ATTRIBUTES.PRI5,
			data[j].ATTRIBUTES.PRI6,	

		)
		checkErr(err)
		//fmt.Println("[+] Pri2que data has been saved")
	}
	tx.Commit()
}


func insertPrivatecabandcomb(eNodeBId string, baseName string, data []model.Privatecabandcomb) {
	fmt.Println("[+] Processing Privatecabandcomb data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `privatecabandcomb` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PrivateCaCombId,
			data[j].ATTRIBUTES.MaxAggregatedBw,
			data[j].ATTRIBUTES.BwCombSetId,
			data[j].ATTRIBUTES.CombBand1Id,
			data[j].ATTRIBUTES.CombBand2Id,
			data[j].ATTRIBUTES.CombBand3Id,
			data[j].ATTRIBUTES.CombBand4Id,
			data[j].ATTRIBUTES.CombBand1Bw,
			data[j].ATTRIBUTES.CombBand2Bw,
			data[j].ATTRIBUTES.CombBand3Bw,
			data[j].ATTRIBUTES.CombBand4Bw,
			data[j].ATTRIBUTES.CombBand5Id,
			data[j].ATTRIBUTES.CombBand5Bw,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Privatecabandcomb data has been saved")
	}
	tx.Commit()
}


func insertPsuis(eNodeBId string, baseName string, data []model.Psuis) {
	fmt.Println("[+] Processing Psuis data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `psuis` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PSUISS,	

		)
		checkErr(err)
		//fmt.Println("[+] Psuis data has been saved")
	}
	tx.Commit()
}


func insertPucchcfg(eNodeBId string, baseName string, data []model.Pucchcfg) {
	fmt.Println("[+] Processing Pucchcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pucchcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DeltaShift,
			data[j].ATTRIBUTES.NaSriChNum,
			data[j].ATTRIBUTES.CqiRbNum,
			data[j].ATTRIBUTES.PucchExtendedRBNum,
			data[j].ATTRIBUTES.Format1ChAllocMode,
			data[j].ATTRIBUTES.SriPeriodAdaptive,
			data[j].ATTRIBUTES.Format2ChAllocMode,
			data[j].ATTRIBUTES.PucchAllocPolicy,
			data[j].ATTRIBUTES.Format3RbNum,
			data[j].ATTRIBUTES.Max2CCAckChNum,	

		)
		checkErr(err)
		//fmt.Println("[+] Pucchcfg data has been saved")
	}
	tx.Commit()
}


func insertPuschcfg(eNodeBId string, baseName string, data []model.Puschcfg) {
	fmt.Println("[+] Processing Puschcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `puschcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Nsb,
			data[j].ATTRIBUTES.HoppingMode,
			data[j].ATTRIBUTES.HoppingOffset,
			data[j].ATTRIBUTES.GroupHoppingEnabled,
			data[j].ATTRIBUTES.GroupAssignPUSCH,
			data[j].ATTRIBUTES.SeqHoppingEnabled,
			data[j].ATTRIBUTES.CyclicShift,
			data[j].ATTRIBUTES.Qam64Enabled,
			data[j].ATTRIBUTES.R12Qam64Enabled,	

		)
		checkErr(err)
		//fmt.Println("[+] Puschcfg data has been saved")
	}
	tx.Commit()
}


func insertPuschparam(eNodeBId string, baseName string, data []model.Puschparam) {
	fmt.Println("[+] Processing Puschparam data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `puschparam` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DeltaOffsetCqiIndex,
			data[j].ATTRIBUTES.DeltaOffsetRiIndex,
			data[j].ATTRIBUTES.DeltaOffsetAckIndex,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Puschparam data has been saved")
	}
	tx.Commit()
}


func insertPwdpolicy(eNodeBId string, baseName string, data []model.Pwdpolicy) {
	fmt.Println("[+] Processing Pwdpolicy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `pwdpolicy` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PWDMINLEN,
			data[j].ATTRIBUTES.COMPLICACY,
			data[j].ATTRIBUTES.PASSREPLMT,
			data[j].ATTRIBUTES.MAXPERIOD,
			data[j].ATTRIBUTES.MINPERIOD,
			data[j].ATTRIBUTES.MAXMISSTIMES,
			data[j].ATTRIBUTES.AUTOUNLOCKTIME,
			data[j].ATTRIBUTES.RESETINTERVAL,
			data[j].ATTRIBUTES.PWDEXPRT,
			data[j].ATTRIBUTES.FirstLoginMustModPWD,
			data[j].ATTRIBUTES.MAXREPEATCHARTIMES,
			data[j].ATTRIBUTES.DICTCHKPWD,
			data[j].ATTRIBUTES.MAXCCUN,	

		)
		checkErr(err)
		//fmt.Println("[+] Pwdpolicy data has been saved")
	}
	tx.Commit()
}


func insertQcipara(eNodeBId string, baseName string, data []model.Qcipara) {
	fmt.Println("[+] Processing Qcipara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `qcipara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.PreAllocationWeight,
			data[j].ATTRIBUTES.InterRatPolicyCfgGroupId,
			data[j].ATTRIBUTES.LogicalChannelPriority,
			data[j].ATTRIBUTES.RlcPdcpParaGroupId,
			data[j].ATTRIBUTES.UeInactiveTimerForQci,
			data[j].ATTRIBUTES.UlSynTimerForQci,
			data[j].ATTRIBUTES.UeInactivityTimerDynDrxQci,
			data[j].ATTRIBUTES.UeInactiveTimerPri,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.PrioritisedBitRate,
			data[j].ATTRIBUTES.FlowCtrlType,
			data[j].ATTRIBUTES.UlschPriorityFactor,
			data[j].ATTRIBUTES.DlschPriorityFactor,
			data[j].ATTRIBUTES.UlMinGbr,
			data[j].ATTRIBUTES.DlMinGbr,
			data[j].ATTRIBUTES.EmtcModeARlcParaGroupId,
			data[j].ATTRIBUTES.EmtcModeBRlcParaGroupId,
			data[j].ATTRIBUTES.NbRlcPdcpParaGroupId,
			data[j].ATTRIBUTES.CiotUeInactiveTimer,
			data[j].ATTRIBUTES.ServiceType,
			data[j].ATTRIBUTES.FreeUserFlag,
			data[j].ATTRIBUTES.LtePttDedicatedExtendedQci,
			data[j].ATTRIBUTES.ExtQciCounterIndex,	

		)
		checkErr(err)
		//fmt.Println("[+] Qcipara data has been saved")
	}
	tx.Commit()
}


func insertQoehocommoncfg(eNodeBId string, baseName string, data []model.Qoehocommoncfg) {
	fmt.Println("[+] Processing Qoehocommoncfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `qoehocommoncfg` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.QoEBasedHandoverStat,
			data[j].ATTRIBUTES.QoEBasedHandoverLast,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Qoehocommoncfg data has been saved")
	}
	tx.Commit()
}


func insertRachcfg(eNodeBId string, baseName string, data []model.Rachcfg) {
	fmt.Println("[+] Processing Rachcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rachcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.PwrRampingStep,
			data[j].ATTRIBUTES.PreambInitRcvTargetPwr,
			data[j].ATTRIBUTES.MessageSizeGroupA,
			data[j].ATTRIBUTES.PrachFreqOffset,
			data[j].ATTRIBUTES.PrachConfigIndexCfgInd,
			data[j].ATTRIBUTES.PreambleTransMax,
			data[j].ATTRIBUTES.ContentionResolutionTimer,
			data[j].ATTRIBUTES.MaxHarqMsg3Tx,
			data[j].ATTRIBUTES.RandomPreambleRatio,
			data[j].ATTRIBUTES.RaPreambleGrpARatio,
			data[j].ATTRIBUTES.PrachFreqOffsetStrategy,
			data[j].ATTRIBUTES.PrachStartTimeCfgInd,
			data[j].ATTRIBUTES.NbCyclicPrefixLength,
			data[j].ATTRIBUTES.NbRsrpFirstThreshold,
			data[j].ATTRIBUTES.NbRsrpSecondThreshold,
			data[j].ATTRIBUTES.UeRaInfoReportRatioThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Rachcfg data has been saved")
	}
	tx.Commit()
}


func insertRet(eNodeBId string, baseName string, data []model.Ret) {
	fmt.Println("[+] Processing Ret data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ret` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DEVICENO,
			data[j].ATTRIBUTES.DEVICENAME,
			data[j].ATTRIBUTES.CTRLCN,
			data[j].ATTRIBUTES.CTRLSRN,
			data[j].ATTRIBUTES.CTRLSN,
			data[j].ATTRIBUTES.VENDORCODE,
			data[j].ATTRIBUTES.SERIALNO,
			data[j].ATTRIBUTES.RETTYPE,
			data[j].ATTRIBUTES.POLARTYPE,
			data[j].ATTRIBUTES.SCENARIO,
			data[j].ATTRIBUTES.SUBUNITNUM,	

		)
		checkErr(err)
		//fmt.Println("[+] Ret data has been saved")
	}
	tx.Commit()
}


func insertRetdevicedata(eNodeBId string, baseName string, data []model.Retdevicedata) {
	fmt.Println("[+] Processing Retdevicedata data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `retdevicedata` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DEVICENO,
			data[j].ATTRIBUTES.SUBUNITNO,
			data[j].ATTRIBUTES.BEARING,
			data[j].ATTRIBUTES.MODELNO,
			data[j].ATTRIBUTES.BSID,
			data[j].ATTRIBUTES.BEAMWIDTH1,
			data[j].ATTRIBUTES.BEAMWIDTH2,
			data[j].ATTRIBUTES.BEAMWIDTH3,
			data[j].ATTRIBUTES.BEAMWIDTH4,
			data[j].ATTRIBUTES.GAIN1,
			data[j].ATTRIBUTES.GAIN2,
			data[j].ATTRIBUTES.GAIN3,
			data[j].ATTRIBUTES.GAIN4,
			data[j].ATTRIBUTES.DATE,
			data[j].ATTRIBUTES.TILT,
			data[j].ATTRIBUTES.INSTALLERID,
			data[j].ATTRIBUTES.BAND1,
			data[j].ATTRIBUTES.BAND2,
			data[j].ATTRIBUTES.BAND3,
			data[j].ATTRIBUTES.BAND4,
			data[j].ATTRIBUTES.SECTORID,
			data[j].ATTRIBUTES.SERIALNO,	

		)
		checkErr(err)
		//fmt.Println("[+] Retdevicedata data has been saved")
	}
	tx.Commit()
}


func insertRetport(eNodeBId string, baseName string, data []model.Retport) {
	fmt.Println("[+] Processing Retport data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `retport` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.PWRSWITCH,
			data[j].ATTRIBUTES.THRESHOLDTYPE,
			data[j].ATTRIBUTES.UOTHD,
			data[j].ATTRIBUTES.UCTHD,
			data[j].ATTRIBUTES.OOTHD,
			data[j].ATTRIBUTES.OCTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Retport data has been saved")
	}
	tx.Commit()
}


func insertRetsubunit(eNodeBId string, baseName string, data []model.Retsubunit) {
	fmt.Println("[+] Processing Retsubunit data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `retsubunit` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.DEVICENO,
			data[j].ATTRIBUTES.SUBUNITNO,
			data[j].ATTRIBUTES.CONNCN1,
			data[j].ATTRIBUTES.CONNSRN1,
			data[j].ATTRIBUTES.CONNSN1,
			data[j].ATTRIBUTES.CONNPN1,
			data[j].ATTRIBUTES.CONNCN2,
			data[j].ATTRIBUTES.CONNSRN2,
			data[j].ATTRIBUTES.CONNSN2,
			data[j].ATTRIBUTES.CONNPN2,
			data[j].ATTRIBUTES.TILT,
			data[j].ATTRIBUTES.AER,
			data[j].ATTRIBUTES.SUBNAME,	

		)
		checkErr(err)
		//fmt.Println("[+] Retsubunit data has been saved")
	}
	tx.Commit()
}


func insertRfu(eNodeBId string, baseName string, data []model.Rfu) {
	fmt.Println("[+] Processing Rfu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rfu` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.ADMSTATE,
			data[j].ATTRIBUTES.IFFREQ,
			data[j].ATTRIBUTES.ALMPROCSW,
			data[j].ATTRIBUTES.ALMPROCTHRHLD,
			data[j].ATTRIBUTES.ALMTHRHLD,
			data[j].ATTRIBUTES.PS,
			data[j].ATTRIBUTES.RCN,
			data[j].ATTRIBUTES.TP,
			data[j].ATTRIBUTES.RS,
			data[j].ATTRIBUTES.RXNUM,
			data[j].ATTRIBUTES.TXNUM,
			data[j].ATTRIBUTES.RFTXSIGNDETECTSW,
			data[j].ATTRIBUTES.RFTXSIGNDETECTPERIOD,
			data[j].ATTRIBUTES.RFTXSIGNDETECTTHLD,
			data[j].ATTRIBUTES.RT,
			data[j].ATTRIBUTES.IFOFFSET,
			data[j].ATTRIBUTES.RFDS,
			data[j].ATTRIBUTES.FMBWH,
			data[j].ATTRIBUTES.LCPSW,
			data[j].ATTRIBUTES.FLAG,
			data[j].ATTRIBUTES.RUSPEC,
			data[j].ATTRIBUTES.RFCONNCN2,
			data[j].ATTRIBUTES.RFCONNSN2,
			data[j].ATTRIBUTES.RFCONNSRN2,
			data[j].ATTRIBUTES.RFCONNTYPE,
			data[j].ATTRIBUTES.PAEFFSWITCH,
			data[j].ATTRIBUTES.SCR,
			data[j].ATTRIBUTES.RXFREQBANDMUTUALSW,
			data[j].ATTRIBUTES.MNTMODE,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.ET,
			data[j].ATTRIBUTES.MMSETREMARK,
			data[j].ATTRIBUTES.LEDSW,
			data[j].ATTRIBUTES.CUSTOMEDRFSPECSW,
			data[j].ATTRIBUTES.PSGID,
			data[j].ATTRIBUTES.WIFISW,	

		)
		checkErr(err)
		//fmt.Println("[+] Rfu data has been saved")
	}
	tx.Commit()
}


func insertRlcpdcpparagroup(eNodeBId string, baseName string, data []model.Rlcpdcpparagroup) {
	fmt.Println("[+] Processing Rlcpdcpparagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rlcpdcpparagroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RlcPdcpParaGroupId,
			data[j].ATTRIBUTES.DiscardTimer,
			data[j].ATTRIBUTES.RlcMode,
			data[j].ATTRIBUTES.PdcpSnSize,
			data[j].ATTRIBUTES.UlRlcSnSize,
			data[j].ATTRIBUTES.DlRlcSnSize,
			data[j].ATTRIBUTES.UeUmReorderingTimer,
			data[j].ATTRIBUTES.ENodeBUmReorderingTimer,
			data[j].ATTRIBUTES.NonsptUmUeAdaptSwitch,
			data[j].ATTRIBUTES.UlDlDiscardtimerSwitch,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.UeMaxRetxThreshold,
			data[j].ATTRIBUTES.ENodeBMaxRetxThreshold,
			data[j].ATTRIBUTES.UePollByte,
			data[j].ATTRIBUTES.ENodeBPollByte,
			data[j].ATTRIBUTES.UePollPdu,
			data[j].ATTRIBUTES.ENodeBPollPdu,
			data[j].ATTRIBUTES.UePollRetransmitTimer,
			data[j].ATTRIBUTES.ENodeBPollRetransmitTimer,
			data[j].ATTRIBUTES.UeStatusProhibitTimer,
			data[j].ATTRIBUTES.ENodeBStatusProhibitTimer,
			data[j].ATTRIBUTES.UeAmReorderingTimer,
			data[j].ATTRIBUTES.ENodeBAmReorderingTimer,
			data[j].ATTRIBUTES.PdcpStatusRptReq,
			data[j].ATTRIBUTES.RlcParaAdaptSwitch,
			data[j].ATTRIBUTES.ENodeBPollRtrTimerPreset,
			data[j].ATTRIBUTES.ENodeBStatProhTimerPreset,
			data[j].ATTRIBUTES.UePollRtrTimerPreset,
			data[j].ATTRIBUTES.UeStatProhTimerPreset,
			data[j].ATTRIBUTES.CaUeRlcParaAdptiveThd,
			data[j].ATTRIBUTES.CaUeReorderingTimer,
			data[j].ATTRIBUTES.CaUeStatProhTimer,
			data[j].ATTRIBUTES.AmPdcpSnSize,
			data[j].ATTRIBUTES.ENodeBReorderingTimerAdapt,
			data[j].ATTRIBUTES.UePdcpReorderingTimer,
			data[j].ATTRIBUTES.CatType,
			data[j].ATTRIBUTES.NbPdcpDiscardTimer,
			data[j].ATTRIBUTES.NbDlPdcpDiscardTimer,
			data[j].ATTRIBUTES.NbEnodebPollRetxTimer,
			data[j].ATTRIBUTES.NbUePollRetxTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Rlcpdcpparagroup data has been saved")
	}
	tx.Commit()
}


func insertRrcconnstatetimer(eNodeBId string, baseName string, data []model.Rrcconnstatetimer) {
	fmt.Println("[+] Processing Rrcconnstatetimer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rrcconnstatetimer` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.T302,
			data[j].ATTRIBUTES.T320ForLoadBalance,
			data[j].ATTRIBUTES.T304ForEutran,
			data[j].ATTRIBUTES.T304ForGeran,
			data[j].ATTRIBUTES.T320ForOther,
			data[j].ATTRIBUTES.UeInactiveTimer,
			data[j].ATTRIBUTES.UlSynTimer,
			data[j].ATTRIBUTES.FilterReptRrcConnReqTimer,
			data[j].ATTRIBUTES.UeInactivityTimerDynDrx,
			data[j].ATTRIBUTES.UlSynTimerDynDrx,
			data[j].ATTRIBUTES.UeInactiveTimerQci1,
			data[j].ATTRIBUTES.UeInactTimerDynDrxQci1,
			data[j].ATTRIBUTES.RrcConnRelTimer,
			data[j].ATTRIBUTES.DRXRrcConnRelTimerOffset,
			data[j].ATTRIBUTES.SRLTERrcConnRelTimerOffset,
			data[j].ATTRIBUTES.ReptRrcReestProtectTimer,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.NBUeInactiveTimer,
			data[j].ATTRIBUTES.ExtendedWaitTime,
			data[j].ATTRIBUTES.MTCUeInactiveTimer,
			data[j].ATTRIBUTES.PowerPrefIndicationTimer,
			data[j].ATTRIBUTES.OutOfSyncRelTimerOffset,	

		)
		checkErr(err)
		//fmt.Println("[+] Rrcconnstatetimer data has been saved")
	}
	tx.Commit()
}


func insertRru(eNodeBId string, baseName string, data []model.Rru) {
	fmt.Println("[+] Processing Rru data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rru` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.ADMSTATE,
			data[j].ATTRIBUTES.IFFREQ,
			data[j].ATTRIBUTES.ALMPROCSW,
			data[j].ATTRIBUTES.ALMPROCTHRHLD,
			data[j].ATTRIBUTES.ALMTHRHLD,
			data[j].ATTRIBUTES.PS,
			data[j].ATTRIBUTES.RCN,
			data[j].ATTRIBUTES.TP,
			data[j].ATTRIBUTES.RS,
			data[j].ATTRIBUTES.RXNUM,
			data[j].ATTRIBUTES.TXNUM,
			data[j].ATTRIBUTES.RFTXSIGNDETECTSW,
			data[j].ATTRIBUTES.RFTXSIGNDETECTPERIOD,
			data[j].ATTRIBUTES.RFTXSIGNDETECTTHLD,
			data[j].ATTRIBUTES.RN,
			data[j].ATTRIBUTES.RT,
			data[j].ATTRIBUTES.IFOFFSET,
			data[j].ATTRIBUTES.RFDS,
			data[j].ATTRIBUTES.FMBWH,
			data[j].ATTRIBUTES.LCPSW,
			data[j].ATTRIBUTES.FLAG,
			data[j].ATTRIBUTES.RUSPEC,
			data[j].ATTRIBUTES.RFCONNCN2,
			data[j].ATTRIBUTES.RFCONNSN2,
			data[j].ATTRIBUTES.RFCONNSRN2,
			data[j].ATTRIBUTES.RFCONNTYPE,
			data[j].ATTRIBUTES.PAEFFSWITCH,
			data[j].ATTRIBUTES.SCR,
			data[j].ATTRIBUTES.RXFREQBANDMUTUALSW,
			data[j].ATTRIBUTES.REMOTEFLAG,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.RFDCPWROFFALMDETECTSW,
			data[j].ATTRIBUTES.BATTOUTPUNDERVOLTTHLD,
			data[j].ATTRIBUTES.MNTMODE,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.ET,
			data[j].ATTRIBUTES.MMSETREMARK,
			data[j].ATTRIBUTES.LEDSW,
			data[j].ATTRIBUTES.CUSTOMEDRFSPECSW,
			data[j].ATTRIBUTES.PSGID,
			data[j].ATTRIBUTES.WIFISW,
			data[j].ATTRIBUTES.LOCATIONNAME,
			data[j].ATTRIBUTES.CIRCUITBREAKERVALUE,	

		)
		checkErr(err)
		//fmt.Println("[+] Rru data has been saved")
	}
	tx.Commit()
}


func insertRruchain(eNodeBId string, baseName string, data []model.Rruchain) {
	fmt.Println("[+] Processing Rruchain data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rruchain` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RCN,
			data[j].ATTRIBUTES.TT,
			data[j].ATTRIBUTES.BM,
			data[j].ATTRIBUTES.HCN,
			data[j].ATTRIBUTES.HSRN,
			data[j].ATTRIBUTES.HSN,
			data[j].ATTRIBUTES.HPN,
			data[j].ATTRIBUTES.BRKPOS1,
			data[j].ATTRIBUTES.BRKPOS2,
			data[j].ATTRIBUTES.AT,
			data[j].ATTRIBUTES.CR,
			data[j].ATTRIBUTES.LSN,
			data[j].ATTRIBUTES.PROTOCOL,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.CONNPORTNUM,	

		)
		checkErr(err)
		//fmt.Println("[+] Rruchain data has been saved")
	}
	tx.Commit()
}


func insertRrujointcalparacfg(eNodeBId string, baseName string, data []model.Rrujointcalparacfg) {
	fmt.Println("[+] Processing Rrujointcalparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rrujointcalparacfg` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TxChnCalSwitch,
			data[j].ATTRIBUTES.TxChnCalTime,
			data[j].ATTRIBUTES.TxChnCalPeriod,
			data[j].ATTRIBUTES.TxChnAntSpacing,
			data[j].ATTRIBUTES.AauPassivePortCalibPeriod,	

		)
		checkErr(err)
		//fmt.Println("[+] Rrujointcalparacfg data has been saved")
	}
	tx.Commit()
}


func insertRscgrp(eNodeBId string, baseName string, data []model.Rscgrp) {
	fmt.Println("[+] Processing Rscgrp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rscgrp` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.BEAR,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.RSCGRPID,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.RU,
			data[j].ATTRIBUTES.TXBW,
			data[j].ATTRIBUTES.RXBW,
			data[j].ATTRIBUTES.TXCBS,
			data[j].ATTRIBUTES.TXEBS,
			data[j].ATTRIBUTES.OID,
			data[j].ATTRIBUTES.WEIGHT,
			data[j].ATTRIBUTES.TXCIR,
			data[j].ATTRIBUTES.RXCIR,
			data[j].ATTRIBUTES.TXPIR,
			data[j].ATTRIBUTES.RXPIR,
			data[j].ATTRIBUTES.TXPBS,	

		)
		checkErr(err)
		//fmt.Println("[+] Rscgrp data has been saved")
	}
	tx.Commit()
}


func insertRscgrpalg(eNodeBId string, baseName string, data []model.Rscgrpalg) {
	fmt.Println("[+] Processing Rscgrpalg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rscgrpalg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.BEAR,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.RSCGRPID,
			data[j].ATTRIBUTES.TXSSW,
			data[j].ATTRIBUTES.TXBWASW,
			data[j].ATTRIBUTES.RXBWASW,
			data[j].ATTRIBUTES.PLRDTH,
			data[j].ATTRIBUTES.DDTH,
			data[j].ATTRIBUTES.TXBWAMIN,
			data[j].ATTRIBUTES.RXBWAMIN,
			data[j].ATTRIBUTES.TCSW,
			data[j].ATTRIBUTES.PQN,
			data[j].ATTRIBUTES.CTTH,
			data[j].ATTRIBUTES.CCTTH,
			data[j].ATTRIBUTES.TXRSVBW,
			data[j].ATTRIBUTES.RXRSVBW,
			data[j].ATTRIBUTES.DROPPKTNUM,	

		)
		checkErr(err)
		//fmt.Println("[+] Rscgrpalg data has been saved")
	}
	tx.Commit()
}


func insertRxbranch(eNodeBId string, baseName string, data []model.Rxbranch) {
	fmt.Println("[+] Processing Rxbranch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `rxbranch` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.RXNO,
			data[j].ATTRIBUTES.RXSW,
			data[j].ATTRIBUTES.ATTEN,
			data[j].ATTRIBUTES.RTWPINITADJ0,
			data[j].ATTRIBUTES.RTWPINITADJ1,
			data[j].ATTRIBUTES.RTWPINITADJ2,
			data[j].ATTRIBUTES.RTWPINITADJ3,
			data[j].ATTRIBUTES.RTWPINITADJ4,
			data[j].ATTRIBUTES.RTWPINITADJ5,
			data[j].ATTRIBUTES.RTWPINITADJ6,
			data[j].ATTRIBUTES.RTWPINITADJ7,	

		)
		checkErr(err)
		//fmt.Println("[+] Rxbranch data has been saved")
	}
	tx.Commit()
}


func insertS1(eNodeBId string, baseName string, data []model.S1) {
	fmt.Println("[+] Processing S1 data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `s1` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.S1Id,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.UpEpGroupId,
			data[j].ATTRIBUTES.MmeRelease,
			data[j].ATTRIBUTES.UserLabel,
			data[j].ATTRIBUTES.EpGroupCfgFlag,
			data[j].ATTRIBUTES.Priority,
			data[j].ATTRIBUTES.CnOpSharingGroupId,
			data[j].ATTRIBUTES.MdtEnable,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] S1 data has been saved")
	}
	tx.Commit()
}


func insertS1interface(eNodeBId string, baseName string, data []model.S1interface) {
	fmt.Println("[+] Processing S1interface data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `s1interface` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.S1InterfaceId,
			data[j].ATTRIBUTES.S1CpBearerId,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.MmeRelease,
			data[j].ATTRIBUTES.S1InterfaceIsBlock,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.AutoCfgFlag,
			data[j].ATTRIBUTES.Priority,
			data[j].ATTRIBUTES.CnOpSharingGroupId,
			data[j].ATTRIBUTES.ServedGummeis,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] S1interface data has been saved")
	}
	tx.Commit()
}


func insertS1reesttimer(eNodeBId string, baseName string, data []model.S1reesttimer) {
	fmt.Println("[+] Processing S1reesttimer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `s1reesttimer` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.S1ReestMinTimer,
			data[j].ATTRIBUTES.S1ReestMaxTimer,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] S1reesttimer data has been saved")
	}
	tx.Commit()
}


func insertSaallnk(eNodeBId string, baseName string, data []model.Saallnk) {
	fmt.Println("[+] Processing Saallnk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `saallnk` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SAALNO,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PT,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.JNRSCGRP,
			data[j].ATTRIBUTES.VPI,
			data[j].ATTRIBUTES.VCI,
			data[j].ATTRIBUTES.ST,
			data[j].ATTRIBUTES.PCR,
			data[j].ATTRIBUTES.CCTM,
			data[j].ATTRIBUTES.POLL,
			data[j].ATTRIBUTES.IDLE,
			data[j].ATTRIBUTES.NRTM,
			data[j].ATTRIBUTES.KATM,
			data[j].ATTRIBUTES.MAXCC,
			data[j].ATTRIBUTES.MAXPD,
			data[j].ATTRIBUTES.MAXLE,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.RU,	

		)
		checkErr(err)
		//fmt.Println("[+] Saallnk data has been saved")
	}
	tx.Commit()
}


func insertScappparacfg(eNodeBId string, baseName string, data []model.Scappparacfg) {
	fmt.Println("[+] Processing Scappparacfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `scappparacfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AppDnsId,
			data[j].ATTRIBUTES.MatchRule,
			data[j].ATTRIBUTES.AppName,
			data[j].ATTRIBUTES.AppIdentType,
			data[j].ATTRIBUTES.AppIpv4,
			data[j].ATTRIBUTES.AppCfgTargetInd,
			data[j].ATTRIBUTES.MainAppDnsFlag,
			data[j].ATTRIBUTES.AsParaGroupID,
			data[j].ATTRIBUTES.MainAppDnsId,
			data[j].ATTRIBUTES.ServiceEndTimeThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Scappparacfg data has been saved")
	}
	tx.Commit()
}


func insertSccfreqcfg(eNodeBId string, baseName string, data []model.Sccfreqcfg) {
	fmt.Println("[+] Processing Sccfreqcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sccfreqcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PccDlEarfcn,
			data[j].ATTRIBUTES.SccDlEarfcn,
			data[j].ATTRIBUTES.SccPriority,
			data[j].ATTRIBUTES.SccA2Offset,
			data[j].ATTRIBUTES.SccA4Offset,
			data[j].ATTRIBUTES.BlindScellAddThd,
			data[j].ATTRIBUTES.BlindScellDelThd,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.SpidGrpId,	

		)
		checkErr(err)
		//fmt.Println("[+] Sccfreqcfg data has been saved")
	}
	tx.Commit()
}


func insertScpolicy(eNodeBId string, baseName string, data []model.Scpolicy) {
	fmt.Println("[+] Processing Scpolicy data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `scpolicy` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ScAlgoSwitch,
			data[j].ATTRIBUTES.VideoInitBufTime,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Scpolicy data has been saved")
	}
	tx.Commit()
}


func insertScserviceqos(eNodeBId string, baseName string, data []model.Scserviceqos) {
	fmt.Println("[+] Processing Scserviceqos data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `scserviceqos` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ScQosId,
			data[j].ATTRIBUTES.AppDns,
			data[j].ATTRIBUTES.DlSgbr,
			data[j].ATTRIBUTES.AppIdentType,	

		)
		checkErr(err)
		//fmt.Println("[+] Scserviceqos data has been saved")
	}
	tx.Commit()
}


func insertSctphost(eNodeBId string, baseName string, data []model.Sctphost) {
	fmt.Println("[+] Processing Sctphost data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sctphost` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SCTPHOSTID,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.IPVERSION,
			data[j].ATTRIBUTES.SIGIP1V4,
			data[j].ATTRIBUTES.SIGIP1SECSWITCH,
			data[j].ATTRIBUTES.SIGIP2V4,
			data[j].ATTRIBUTES.SIGIP2SECSWITCH,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.SCTPTEMPLATEID,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.SIMPLEMODESWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Sctphost data has been saved")
	}
	tx.Commit()
}


func insertSctplnk(eNodeBId string, baseName string, data []model.Sctplnk) {
	fmt.Println("[+] Processing Sctplnk data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sctplnk` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SCTPNO,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MAXSTREAM,
			data[j].ATTRIBUTES.CTRLMODE,
			data[j].ATTRIBUTES.LOCIP,
			data[j].ATTRIBUTES.SECLOCIP,
			data[j].ATTRIBUTES.LOCPORT,
			data[j].ATTRIBUTES.PEERIP,
			data[j].ATTRIBUTES.SECPEERIP,
			data[j].ATTRIBUTES.PEERPORT,
			data[j].ATTRIBUTES.RTOMIN,
			data[j].ATTRIBUTES.RTOMAX,
			data[j].ATTRIBUTES.RTOINIT,
			data[j].ATTRIBUTES.RTOALPHA,
			data[j].ATTRIBUTES.RTOBETA,
			data[j].ATTRIBUTES.HBINTER,
			data[j].ATTRIBUTES.MAXASSOCRETR,
			data[j].ATTRIBUTES.MAXPATHRETR,
			data[j].ATTRIBUTES.CHKSUMTYPE,
			data[j].ATTRIBUTES.AUTOSWITCH,
			data[j].ATTRIBUTES.SWITCHBACKHBNUM,
			data[j].ATTRIBUTES.BLKFLAG,
			data[j].ATTRIBUTES.TSACK,
			data[j].ATTRIBUTES.DESCRI,
			data[j].ATTRIBUTES.AUTOCFGFLAG,
			data[j].ATTRIBUTES.VRFIDX,	

		)
		checkErr(err)
		//fmt.Println("[+] Sctplnk data has been saved")
	}
	tx.Commit()
}


func insertSctppeer(eNodeBId string, baseName string, data []model.Sctppeer) {
	fmt.Println("[+] Processing Sctppeer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sctppeer` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SCTPPEERID,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.IPVERSION,
			data[j].ATTRIBUTES.SIGIP1V4,
			data[j].ATTRIBUTES.SIGIP1SECSWITCH,
			data[j].ATTRIBUTES.SIGIP2V4,
			data[j].ATTRIBUTES.SIGIP2SECSWITCH,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.REMOTEID,
			data[j].ATTRIBUTES.CTRLMODE,
			data[j].ATTRIBUTES.AUTOCFGFLAG,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.BLKFLAG,
			data[j].ATTRIBUTES.SIMPLEMODESWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Sctppeer data has been saved")
	}
	tx.Commit()
}


func insertSctptemplate(eNodeBId string, baseName string, data []model.Sctptemplate) {
	fmt.Println("[+] Processing Sctptemplate data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sctptemplate` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SCTPTEMPLATEID,
			data[j].ATTRIBUTES.RTOMIN,
			data[j].ATTRIBUTES.RTOMAX,
			data[j].ATTRIBUTES.RTOINIT,
			data[j].ATTRIBUTES.RTOALPHA,
			data[j].ATTRIBUTES.RTOBETA,
			data[j].ATTRIBUTES.HBINTER,
			data[j].ATTRIBUTES.MAXASSOCRETR,
			data[j].ATTRIBUTES.MAXPATHRETR,
			data[j].ATTRIBUTES.SWITCHBACKFLAG,
			data[j].ATTRIBUTES.SWITCHBACHHBNUM,
			data[j].ATTRIBUTES.TSACK,
			data[j].ATTRIBUTES.CHKSUMTYPE,
			data[j].ATTRIBUTES.MAXSTREAM,	

		)
		checkErr(err)
		//fmt.Println("[+] Sctptemplate data has been saved")
	}
	tx.Commit()
}


func insertSector(eNodeBId string, baseName string, data []model.Sector) {
	fmt.Println("[+] Processing Sector data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sector` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SECTORID,
			data[j].ATTRIBUTES.SECNAME,
			data[j].ATTRIBUTES.LOCATIONNAME,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.ANTAZIMUTH,
			data[j].ATTRIBUTES.OLDSECTORID,
			data[j].ATTRIBUTES.SECTORIDFORCONVERSION,
			data[j].ATTRIBUTES.SECTORTYPEUMTS,
			data[j].ATTRIBUTES.RXANTNUM,
			data[j].ATTRIBUTES.DIVMODE,
			data[j].ATTRIBUTES.COVERTYPE,
			data[j].ATTRIBUTES.RFCONNMODE,
			data[j].ATTRIBUTES.SECTORMODELTE,
			data[j].ATTRIBUTES.ANTENNAMODE,
			data[j].ATTRIBUTES.SECTORCOMBIND,
			data[j].ATTRIBUTES.OMNIFLAG,
			data[j].ATTRIBUTES.ORIENTOFMAJORAXIS,
			data[j].ATTRIBUTES.CONFIDENCE,
			data[j].ATTRIBUTES.UNCERTSEMIMAJOR,
			data[j].ATTRIBUTES.UNCERTSEMIMINOR,
			data[j].ATTRIBUTES.UNCERTALTITUDE,
			data[j].ATTRIBUTES.SECTORANTENNA,	

		)
		checkErr(err)
		//fmt.Println("[+] Sector data has been saved")
	}
	tx.Commit()
}


func insertSectoreqm(eNodeBId string, baseName string, data []model.Sectoreqm) {
	fmt.Println("[+] Processing Sectoreqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sectoreqm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SECTOREQMID,
			data[j].ATTRIBUTES.SECTORID,
			data[j].ATTRIBUTES.SECTOREQMANTENNA,
			data[j].ATTRIBUTES.ANTCFGMODE,
			data[j].ATTRIBUTES.RRUCN,
			data[j].ATTRIBUTES.RRUSRN,
			data[j].ATTRIBUTES.RRUSN,
			data[j].ATTRIBUTES.BEAMSHAPE,
			data[j].ATTRIBUTES.BEAMLAYERSPLIT,
			data[j].ATTRIBUTES.BEAMAZIMUTHOFFSET,	

		)
		checkErr(err)
		//fmt.Println("[+] Sectoreqm data has been saved")
	}
	tx.Commit()
}


func insertServicediffsetting(eNodeBId string, baseName string, data []model.Servicediffsetting) {
	fmt.Println("[+] Processing Servicediffsetting data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `servicediffsetting` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.QueueWeight0,
			data[j].ATTRIBUTES.QueueWeight1,
			data[j].ATTRIBUTES.QueueWeight2,
			data[j].ATTRIBUTES.QueueWeight3,
			data[j].ATTRIBUTES.QueueWeight4,
			data[j].ATTRIBUTES.QueueWeight5,
			data[j].ATTRIBUTES.QueueWeight6,
			data[j].ATTRIBUTES.QueueWeight7,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Servicediffsetting data has been saved")
	}
	tx.Commit()
}


func insertServiceidentifypara(eNodeBId string, baseName string, data []model.Serviceidentifypara) {
	fmt.Println("[+] Processing Serviceidentifypara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `serviceidentifypara` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.HeartbeatPacketLengthThld,
			data[j].ATTRIBUTES.HeartbeatPacketNumberThld,
			data[j].ATTRIBUTES.MassFlowBigPacketRateThld,
			data[j].ATTRIBUTES.MassFlowDuration,
			data[j].ATTRIBUTES.MassFlowPacketNumberThld,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Serviceidentifypara data has been saved")
	}
	tx.Commit()
}


func insertServiceifdlearfcngrp(eNodeBId string, baseName string, data []model.Serviceifdlearfcngrp) {
	fmt.Println("[+] Processing Serviceifdlearfcngrp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `serviceifdlearfcngrp` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.ServiceIfHoCfgGroupId,
			data[j].ATTRIBUTES.DlEarfcnIndex,
			data[j].ATTRIBUTES.DlEarfcn,
			data[j].ATTRIBUTES.ServiceHoFreqPriority,	

		)
		checkErr(err)
		//fmt.Println("[+] Serviceifdlearfcngrp data has been saved")
	}
	tx.Commit()
}


func insertServiceifhocfggroup(eNodeBId string, baseName string, data []model.Serviceifhocfggroup) {
	fmt.Println("[+] Processing Serviceifhocfggroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `serviceifhocfggroup` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.ServiceIfHoCfgGroupId,
			data[j].ATTRIBUTES.InterFreqHoState,
			data[j].ATTRIBUTES.A4RptWaitingTimer,
			data[j].ATTRIBUTES.A4TimeToTriger,	

		)
		checkErr(err)
		//fmt.Println("[+] Serviceifhocfggroup data has been saved")
	}
	tx.Commit()
}


func insertServiceirhocfggroup(eNodeBId string, baseName string, data []model.Serviceirhocfggroup) {
	fmt.Println("[+] Processing Serviceirhocfggroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `serviceirhocfggroup` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.ServiceIrHoCfgGroupId,
			data[j].ATTRIBUTES.InterRatHoState,
			data[j].ATTRIBUTES.ServiceIrMeasMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Serviceirhocfggroup data has been saved")
	}
	tx.Commit()
}


func insertSfp(eNodeBId string, baseName string, data []model.Sfp) {
	fmt.Println("[+] Processing Sfp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `sfp` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MODULEID,
			data[j].ATTRIBUTES.PT,	

		)
		checkErr(err)
		//fmt.Println("[+] Sfp data has been saved")
	}
	tx.Commit()
}


func insertSimuload(eNodeBId string, baseName string, data []model.Simuload) {
	fmt.Println("[+] Processing Simuload data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `simuload` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SimuLoadCfgIndex,
			data[j].ATTRIBUTES.SimuLoadRbThd,
			data[j].ATTRIBUTES.SimuLoadPwrThd,
			data[j].ATTRIBUTES.SimuLoadReportPeriod,
			data[j].ATTRIBUTES.SimuLoadStatPeriod,
			data[j].ATTRIBUTES.FreqSelSwitch,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Simuload data has been saved")
	}
	tx.Commit()
}


func insertSingleipswitch(eNodeBId string, baseName string, data []model.Singleipswitch) {
	fmt.Println("[+] Processing Singleipswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `singleipswitch` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SINGLEIPSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Singleipswitch data has been saved")
	}
	tx.Commit()
}


func insertSrbcfg(eNodeBId string, baseName string, data []model.Srbcfg) {
	fmt.Println("[+] Processing Srbcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `srbcfg` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SrbRlcParaAdaptSwitch,
			data[j].ATTRIBUTES.SrbPollTimerAdjustStep,
			data[j].ATTRIBUTES.SrbPollTimerMaxAdjustValue,
			data[j].ATTRIBUTES.SrbPollTimerAdjUserNumThd,
			data[j].ATTRIBUTES.RrcConnRelMaxRetxThd,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.SrbPollTimerPreset,
			data[j].ATTRIBUTES.HOCmdRRCRlsMsgRetxAdpSw,	

		)
		checkErr(err)
		//fmt.Println("[+] Srbcfg data has been saved")
	}
	tx.Commit()
}


func insertSrbrlcpdcpcfg(eNodeBId string, baseName string, data []model.Srbrlcpdcpcfg) {
	fmt.Println("[+] Processing Srbrlcpdcpcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `srbrlcpdcpcfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SrbId,
			data[j].ATTRIBUTES.UeMaxRetxThreshold,
			data[j].ATTRIBUTES.ENodeBMaxRetxThreshold,
			data[j].ATTRIBUTES.UePollByte,
			data[j].ATTRIBUTES.ENodeBPollByte,
			data[j].ATTRIBUTES.UePollPdu,
			data[j].ATTRIBUTES.ENodeBPollPdu,
			data[j].ATTRIBUTES.UePollRetransmitTimer,
			data[j].ATTRIBUTES.ENodeBPollRetransmitTimer,
			data[j].ATTRIBUTES.UeStatusProhibitTimer,
			data[j].ATTRIBUTES.ENodeBStatusProhibitTimer,
			data[j].ATTRIBUTES.UeAmReorderingTimer,
			data[j].ATTRIBUTES.ENodeBAmReorderingTimer,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.NbENodeBPollRetransTimer,
			data[j].ATTRIBUTES.NbUePollRetransTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Srbrlcpdcpcfg data has been saved")
	}
	tx.Commit()
}


func insertSrciprt(eNodeBId string, baseName string, data []model.Srciprt) {
	fmt.Println("[+] Processing Srciprt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `srciprt` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SRCRTIDX,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SBT,
			data[j].ATTRIBUTES.SRCIP,
			data[j].ATTRIBUTES.RTTYPE,
			data[j].ATTRIBUTES.NEXTHOP,
			data[j].ATTRIBUTES.PREF,
			data[j].ATTRIBUTES.USERLABEL,	

		)
		checkErr(err)
		//fmt.Println("[+] Srciprt data has been saved")
	}
	tx.Commit()
}


func insertSrsadaptivecfg(eNodeBId string, baseName string, data []model.Srsadaptivecfg) {
	fmt.Println("[+] Processing Srsadaptivecfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `srsadaptivecfg` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SrsPeriodAdaptive,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Srsadaptivecfg data has been saved")
	}
	tx.Commit()
}


func insertSrscfg(eNodeBId string, baseName string, data []model.Srscfg) {
	fmt.Println("[+] Processing Srscfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `srscfg` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.SrsSubframeCfg,
			data[j].ATTRIBUTES.AnSrsSimuTrans,
			data[j].ATTRIBUTES.SrsCfgInd,
			data[j].ATTRIBUTES.TddSrsCfgMode,
			data[j].ATTRIBUTES.FddSrsCfgMode,
			data[j].ATTRIBUTES.SrsAlgoOptSwitch,
			data[j].ATTRIBUTES.SrsCfgPolicySwitch,
			data[j].ATTRIBUTES.SrsResOptSwitch,	

		)
		checkErr(err)
		//fmt.Println("[+] Srscfg data has been saved")
	}
	tx.Commit()
}


func insertSsl(eNodeBId string, baseName string, data []model.Ssl) {
	fmt.Println("[+] Processing Ssl data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ssl` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CONNTYPE,
			data[j].ATTRIBUTES.AUTHMODE,
			data[j].ATTRIBUTES.RENEGO,
			data[j].ATTRIBUTES.RENEGOINTERVAL,
			data[j].ATTRIBUTES.LOWESTCSLEVEL,
			data[j].ATTRIBUTES.VERSION,	

		)
		checkErr(err)
		//fmt.Println("[+] Ssl data has been saved")
	}
	tx.Commit()
}


func insertStandardqci(eNodeBId string, baseName string, data []model.Standardqci) {
	fmt.Println("[+] Processing Standardqci data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `standardqci` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.PreAllocationWeight,
			data[j].ATTRIBUTES.InterRatPolicyCfgGroupId,
			data[j].ATTRIBUTES.RlcPdcpParaGroupId,
			data[j].ATTRIBUTES.LogicalChannelPriority,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.PrioritisedBitRate,
			data[j].ATTRIBUTES.FlowCtrlType,
			data[j].ATTRIBUTES.UlschPriorityFactor,
			data[j].ATTRIBUTES.DlschPriorityFactor,
			data[j].ATTRIBUTES.UlMinGbr,
			data[j].ATTRIBUTES.DlMinGbr,	

		)
		checkErr(err)
		//fmt.Println("[+] Standardqci data has been saved")
	}
	tx.Commit()
}


func insertSubrack(eNodeBId string, baseName string, data []model.Subrack) {
	fmt.Println("[+] Processing Subrack data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `subrack` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.TYPE,
			data[j].ATTRIBUTES.DESC,	

		)
		checkErr(err)
		//fmt.Println("[+] Subrack data has been saved")
	}
	tx.Commit()
}


func insertSynceth(eNodeBId string, baseName string, data []model.Synceth) {
	fmt.Println("[+] Processing Synceth data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `synceth` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LN,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.QL,
			data[j].ATTRIBUTES.SSM,
			data[j].ATTRIBUTES.PRI,
			data[j].ATTRIBUTES.TYPE,	

		)
		checkErr(err)
		//fmt.Println("[+] Synceth data has been saved")
	}
	tx.Commit()
}


func insertTacalg(eNodeBId string, baseName string, data []model.Tacalg) {
	fmt.Println("[+] Processing Tacalg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tacalg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RSCGRPULCACSWITCH,
			data[j].ATTRIBUTES.RSCGRPDLCACSWITCH,
			data[j].ATTRIBUTES.TRMULHOCACTH,
			data[j].ATTRIBUTES.TRMDLHOCACTH,
			data[j].ATTRIBUTES.TRMULGOLDCACTH,
			data[j].ATTRIBUTES.TRMDLGOLDCACTH,
			data[j].ATTRIBUTES.TRMULSILVERCACTH,
			data[j].ATTRIBUTES.TRMDLSILVERCACTH,
			data[j].ATTRIBUTES.TRMULBRONZECACTH,
			data[j].ATTRIBUTES.TRMDLBRONZECACTH,
			data[j].ATTRIBUTES.TRMULGBRCACTH,
			data[j].ATTRIBUTES.TRMDLGBRCACTH,
			data[j].ATTRIBUTES.TRMULPRESW,
			data[j].ATTRIBUTES.TRMDLPRESW,
			data[j].ATTRIBUTES.PORTULOBSW,
			data[j].ATTRIBUTES.PORTDLOBSW,
			data[j].ATTRIBUTES.PORTULCACSW,
			data[j].ATTRIBUTES.PORTDLCACSW,
			data[j].ATTRIBUTES.EMCTACPSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Tacalg data has been saved")
	}
	tx.Commit()
}


func insertTasm(eNodeBId string, baseName string, data []model.Tasm) {
	fmt.Println("[+] Processing Tasm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tasm` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.SRCNO,
			data[j].ATTRIBUTES.CLKSRC,
			data[j].ATTRIBUTES.MODE,
			data[j].ATTRIBUTES.ALGORITHM,
			data[j].ATTRIBUTES.FREECOUNT,
			data[j].ATTRIBUTES.SEARCHCOUNT,
			data[j].ATTRIBUTES.HOLDCOUNT,
			data[j].ATTRIBUTES.LOCKCOUNT,
			data[j].ATTRIBUTES.SAMPLETIME,
			data[j].ATTRIBUTES.SYNMODE,
			data[j].ATTRIBUTES.PERIOD,
			data[j].ATTRIBUTES.TM,
			data[j].ATTRIBUTES.CLKSYNCMODE,
			data[j].ATTRIBUTES.QL,
			data[j].ATTRIBUTES.FREQUENCE,
			data[j].ATTRIBUTES.NETMODE,
			data[j].ATTRIBUTES.FLAG,
			data[j].ATTRIBUTES.DAY,
			data[j].ATTRIBUTES.GSM_FBOFFSET,
			data[j].ATTRIBUTES.SYSCLKSRC,
			data[j].ATTRIBUTES.CLOUDSRC,
			data[j].ATTRIBUTES.FRAMESYNCSW,
			data[j].ATTRIBUTES.GSMFRAMESYNCSW,
			data[j].ATTRIBUTES.STANDARD,
			data[j].ATTRIBUTES.ENSYSCLKCHKSW,
			data[j].ATTRIBUTES.ATRSW,
			data[j].ATTRIBUTES.FNSYNCSW,
			data[j].ATTRIBUTES.DATE,
			data[j].ATTRIBUTES.TIME,
			data[j].ATTRIBUTES.LEAPSECONDSCHGDATE,
			data[j].ATTRIBUTES.LEAPSECONDSCHGTIME,
			data[j].ATTRIBUTES.CRTGPSTOUTCLEAPSECONDS,
			data[j].ATTRIBUTES.NEXTGPSTOUTCLEAPSECONDS,
			data[j].ATTRIBUTES.LPFNSYNCSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Tasm data has been saved")
	}
	tx.Commit()
}


func insertTbdspinfo(eNodeBId string, baseName string, data []model.Tbdspinfo) {
	fmt.Println("[+] Processing Tbdspinfo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tbdspinfo` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LANNO,
			data[j].ATTRIBUTES.VERSION,	

		)
		checkErr(err)
		//fmt.Println("[+] Tbdspinfo data has been saved")
	}
	tx.Commit()
}


func insertTblangno(eNodeBId string, baseName string, data []model.Tblangno) {
	fmt.Println("[+] Processing Tblangno data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tblangno` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LANNO,	

		)
		checkErr(err)
		//fmt.Println("[+] Tblangno data has been saved")
	}
	tx.Commit()
}


func insertTcpackctrlalgo(eNodeBId string, baseName string, data []model.Tcpackctrlalgo) {
	fmt.Println("[+] Processing Tcpackctrlalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tcpackctrlalgo` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.AckCtrlSwitch,
			data[j].ATTRIBUTES.DlMaxThroughput,
			data[j].ATTRIBUTES.CtrlTimerLength,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Tcpackctrlalgo data has been saved")
	}
	tx.Commit()
}


func insertTcpacklimitalg(eNodeBId string, baseName string, data []model.Tcpacklimitalg) {
	fmt.Println("[+] Processing Tcpacklimitalg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tcpacklimitalg` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TCPACKLIMITSWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Tcpacklimitalg data has been saved")
	}
	tx.Commit()
}


func insertTcpmssctrl(eNodeBId string, baseName string, data []model.Tcpmssctrl) {
	fmt.Println("[+] Processing Tcpmssctrl data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tcpmssctrl` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TcpMssCtrlSwitch,
			data[j].ATTRIBUTES.TcpMssThd,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Tcpmssctrl data has been saved")
	}
	tx.Commit()
}


func insertTcu(eNodeBId string, baseName string, data []model.Tcu) {
	fmt.Println("[+] Processing Tcu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tcu` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.MCN,
			data[j].ATTRIBUTES.MSRN,
			data[j].ATTRIBUTES.MPN,
			data[j].ATTRIBUTES.ADDR,
			data[j].ATTRIBUTES.TLTHD,
			data[j].ATTRIBUTES.TUTHD,
			data[j].ATTRIBUTES.SBAF,
			data[j].ATTRIBUTES.TCMODE,	

		)
		checkErr(err)
		//fmt.Println("[+] Tcu data has been saved")
	}
	tx.Commit()
}


func insertTddframeoffset(eNodeBId string, baseName string, data []model.Tddframeoffset) {
	fmt.Println("[+] Processing Tddframeoffset data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tddframeoffset` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TimeOffset,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Tddframeoffset data has been saved")
	}
	tx.Commit()
}


func insertTddresmodeswitch(eNodeBId string, baseName string, data []model.Tddresmodeswitch) {
	fmt.Println("[+] Processing Tddresmodeswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tddresmodeswitch` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ClkResExcludeSwitch,
			data[j].ATTRIBUTES.BbResExclusiveSwitch,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Tddresmodeswitch data has been saved")
	}
	tx.Commit()
}


func insertTimealignmenttimer(eNodeBId string, baseName string, data []model.Timealignmenttimer) {
	fmt.Println("[+] Processing Timealignmenttimer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `timealignmenttimer` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TimeAlignmentTimer,
			data[j].ATTRIBUTES.TimingAdvCmdOptSwitch,
			data[j].ATTRIBUTES.TimingMeasMode,
			data[j].ATTRIBUTES.TACmdSendPeriod,
			data[j].ATTRIBUTES.TimingResOptSwitch,
			data[j].ATTRIBUTES.PucchTimeOffsetCompVal,
			data[j].ATTRIBUTES.HighSpeedTaCmdSendPeriod,
			data[j].ATTRIBUTES.TaEnhance,	

		)
		checkErr(err)
		//fmt.Println("[+] Timealignmenttimer data has been saved")
	}
	tx.Commit()
}


func insertTimesrc(eNodeBId string, baseName string, data []model.Timesrc) {
	fmt.Println("[+] Processing Timesrc data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `timesrc` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TIMESRC,
			data[j].ATTRIBUTES.AUTOSWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Timesrc data has been saved")
	}
	tx.Commit()
}


func insertTimethrd(eNodeBId string, baseName string, data []model.Timethrd) {
	fmt.Println("[+] Processing Timethrd data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `timethrd` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.THRD,
			data[j].ATTRIBUTES.SWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Timethrd data has been saved")
	}
	tx.Commit()
}


func insertTldralg(eNodeBId string, baseName string, data []model.Tldralg) {
	fmt.Println("[+] Processing Tldralg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tldralg` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TRMULLDRTRGTH,
			data[j].ATTRIBUTES.TRMDLLDRTRGTH,
			data[j].ATTRIBUTES.TRMULLDRCLRTH,
			data[j].ATTRIBUTES.TRMDLLDRCLRTH,
			data[j].ATTRIBUTES.TRMULMLDTRGTH,
			data[j].ATTRIBUTES.TRMDLMLDTRGTH,
			data[j].ATTRIBUTES.TRMULMLDCLRTH,
			data[j].ATTRIBUTES.TRMDLMLDCLRTH,	

		)
		checkErr(err)
		//fmt.Println("[+] Tldralg data has been saved")
	}
	tx.Commit()
}


func insertTlfrswitch(eNodeBId string, baseName string, data []model.Tlfrswitch) {
	fmt.Println("[+] Processing Tlfrswitch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tlfrswitch` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TLFRSWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Tlfrswitch data has been saved")
	}
	tx.Commit()
}


func insertTolcalg(eNodeBId string, baseName string, data []model.Tolcalg) {
	fmt.Println("[+] Processing Tolcalg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tolcalg` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TRMULOLCSWITCH,
			data[j].ATTRIBUTES.TRMDLOLCSWITCH,
			data[j].ATTRIBUTES.TRMULOLCTRIGTH,
			data[j].ATTRIBUTES.TRMULOLCRELTH,
			data[j].ATTRIBUTES.TRMOLCRELBEARERNUM,
			data[j].ATTRIBUTES.TRMDLOLCTRIGTH,
			data[j].ATTRIBUTES.TRMDLOLCRELTH,	

		)
		checkErr(err)
		//fmt.Println("[+] Tolcalg data has been saved")
	}
	tx.Commit()
}


func insertTpealgo(eNodeBId string, baseName string, data []model.Tpealgo) {
	fmt.Println("[+] Processing Tpealgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tpealgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.PortListNum,
			data[j].ATTRIBUTES.Port1,
			data[j].ATTRIBUTES.Port2,
			data[j].ATTRIBUTES.Port3,
			data[j].ATTRIBUTES.Port4,
			data[j].ATTRIBUTES.Port5,
			data[j].ATTRIBUTES.Port6,
			data[j].ATTRIBUTES.Port7,
			data[j].ATTRIBUTES.Port8,
			data[j].ATTRIBUTES.Port9,
			data[j].ATTRIBUTES.Port10,
			data[j].ATTRIBUTES.Port11,
			data[j].ATTRIBUTES.Port12,
			data[j].ATTRIBUTES.Port13,
			data[j].ATTRIBUTES.Port14,
			data[j].ATTRIBUTES.Port15,
			data[j].ATTRIBUTES.Port16,
			data[j].ATTRIBUTES.Port17,
			data[j].ATTRIBUTES.Port18,
			data[j].ATTRIBUTES.Port19,
			data[j].ATTRIBUTES.Port20,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Tpealgo data has been saved")
	}
	tx.Commit()
}


func insertTrp(eNodeBId string, baseName string, data []model.Trp) {
	fmt.Println("[+] Processing Trp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `trp` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.SUBTYPE,	

		)
		checkErr(err)
		//fmt.Println("[+] Trp data has been saved")
	}
	tx.Commit()
}


func insertTrustcert(eNodeBId string, baseName string, data []model.Trustcert) {
	fmt.Println("[+] Processing Trustcert data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `trustcert` VALUES(?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CERTNAME,	

		)
		checkErr(err)
		//fmt.Println("[+] Trustcert data has been saved")
	}
	tx.Commit()
}


func insertTwampresponder(eNodeBId string, baseName string, data []model.Twampresponder) {
	fmt.Println("[+] Processing Twampresponder data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `twampresponder` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.RESPONDERID,
			data[j].ATTRIBUTES.LOCALIP,
			data[j].ATTRIBUTES.VRFINDEX,
			data[j].ATTRIBUTES.DSCP,
			data[j].ATTRIBUTES.REFWAIT,
			data[j].ATTRIBUTES.SERVWAIT,
			data[j].ATTRIBUTES.LOCALPORT,
			data[j].ATTRIBUTES.TWAMPARCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Twampresponder data has been saved")
	}
	tx.Commit()
}


func insertTxbranch(eNodeBId string, baseName string, data []model.Txbranch) {
	fmt.Println("[+] Processing Txbranch data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `txbranch` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.TXNO,
			data[j].ATTRIBUTES.TXSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Txbranch data has been saved")
	}
	tx.Commit()
}


func insertTypdrbbsr(eNodeBId string, baseName string, data []model.Typdrbbsr) {
	fmt.Println("[+] Processing Typdrbbsr data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `typdrbbsr` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.TPerodicBSRTimer,
			data[j].ATTRIBUTES.RetxBsrTimer,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Typdrbbsr data has been saved")
	}
	tx.Commit()
}


func insertTz(eNodeBId string, baseName string, data []model.Tz) {
	fmt.Println("[+] Processing Tz data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `tz` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ZONET,
			data[j].ATTRIBUTES.DST,	

		)
		checkErr(err)
		//fmt.Println("[+] Tz data has been saved")
	}
	tx.Commit()
}


func insertUdpping(eNodeBId string, baseName string, data []model.Udpping) {
	fmt.Println("[+] Processing Udpping data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `udpping` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.TIMEOUTTHD,
			data[j].ATTRIBUTES.TIMEOUTCNT,
			data[j].ATTRIBUTES.DSCP,	

		)
		checkErr(err)
		//fmt.Println("[+] Udpping data has been saved")
	}
	tx.Commit()
}


func insertUdt(eNodeBId string, baseName string, data []model.Udt) {
	fmt.Println("[+] Processing Udt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `udt` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UDTNO,
			data[j].ATTRIBUTES.UDTPARAGRPID,	

		)
		checkErr(err)
		//fmt.Println("[+] Udt data has been saved")
	}
	tx.Commit()
}


func insertUdtparagrp(eNodeBId string, baseName string, data []model.Udtparagrp) {
	fmt.Println("[+] Processing Udtparagrp data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `udtparagrp` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UDTPARAGRPID,
			data[j].ATTRIBUTES.PRIRULE,
			data[j].ATTRIBUTES.PRI,
			data[j].ATTRIBUTES.ACTFACTOR,
			data[j].ATTRIBUTES.PRIMTRANRSCTYPE,
			data[j].ATTRIBUTES.PRIMPTLOADTH,
			data[j].ATTRIBUTES.PRIM2SECPTLOADRATH,	

		)
		checkErr(err)
		//fmt.Println("[+] Udtparagrp data has been saved")
	}
	tx.Commit()
}


func insertUecooperationpara(eNodeBId string, baseName string, data []model.Uecooperationpara) {
	fmt.Println("[+] Processing Uecooperationpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `uecooperationpara` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.MacCeMsgLCID,
			data[j].ATTRIBUTES.A2Offset,
			data[j].ATTRIBUTES.A3Offset,
			data[j].ATTRIBUTES.UEAwarePowerSavingSwitch,
			data[j].ATTRIBUTES.UEForbiddenMsgThd,
			data[j].ATTRIBUTES.SpecUserCooperationSwitch,
			data[j].ATTRIBUTES.DsdsCooperationRptSwitch,
			data[j].ATTRIBUTES.DsdsLcid,	

		)
		checkErr(err)
		//fmt.Println("[+] Uecooperationpara data has been saved")
	}
	tx.Commit()
}


func insertUeiu(eNodeBId string, baseName string, data []model.Ueiu) {
	fmt.Println("[+] Processing Ueiu data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ueiu` VALUES(?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,	

		)
		checkErr(err)
		//fmt.Println("[+] Ueiu data has been saved")
	}
	tx.Commit()
}


func insertUespecdrxparagroup(eNodeBId string, baseName string, data []model.Uespecdrxparagroup) {
	fmt.Println("[+] Processing Uespecdrxparagroup data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `uespecdrxparagroup` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DrxParaGroupIndex,
			data[j].ATTRIBUTES.DrxParaGroupStat,
			data[j].ATTRIBUTES.OnDurationTimer,
			data[j].ATTRIBUTES.DrxInactivityTimer,
			data[j].ATTRIBUTES.DrxReTxTimer,
			data[j].ATTRIBUTES.LongDrxCycle,
			data[j].ATTRIBUTES.SupportShortDrx,
			data[j].ATTRIBUTES.ShortDrxCycle,
			data[j].ATTRIBUTES.DrxShortCycleTimer,	

		)
		checkErr(err)
		//fmt.Println("[+] Uespecdrxparagroup data has been saved")
	}
	tx.Commit()
}


func insertUetimerconst(eNodeBId string, baseName string, data []model.Uetimerconst) {
	fmt.Println("[+] Processing Uetimerconst data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `uetimerconst` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.T300,
			data[j].ATTRIBUTES.T301,
			data[j].ATTRIBUTES.T310,
			data[j].ATTRIBUTES.T311,
			data[j].ATTRIBUTES.N311,
			data[j].ATTRIBUTES.N310,
			data[j].ATTRIBUTES.T325,
			data[j].ATTRIBUTES.N313,
			data[j].ATTRIBUTES.N314,
			data[j].ATTRIBUTES.T307,
			data[j].ATTRIBUTES.T313,
			data[j].ATTRIBUTES.T300ForNb,
			data[j].ATTRIBUTES.T310ForNb,
			data[j].ATTRIBUTES.T301ForNb,
			data[j].ATTRIBUTES.T311ForNb,
			data[j].ATTRIBUTES.T300CE,
			data[j].ATTRIBUTES.T301CE,
			data[j].ATTRIBUTES.T310CE,
			data[j].ATTRIBUTES.T311CE,	

		)
		checkErr(err)
		//fmt.Println("[+] Uetimerconst data has been saved")
	}
	tx.Commit()
}


func insertUlcsalgopara(eNodeBId string, baseName string, data []model.Ulcsalgopara) {
	fmt.Println("[+] Processing Ulcsalgopara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulcsalgopara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlCoPcInThd,
			data[j].ATTRIBUTES.UlCoPcRbkRsrpThd,
			data[j].ATTRIBUTES.UlCoPcUserNumThd,
			data[j].ATTRIBUTES.UlCoResAllocBandMode,
			data[j].ATTRIBUTES.UlCoResAllocRbLoadThld,
			data[j].ATTRIBUTES.UlCsA3Offset,
			data[j].ATTRIBUTES.UlCsSw,
			data[j].ATTRIBUTES.UlCraInThld,
			data[j].ATTRIBUTES.UlCraUserNumThld,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulcsalgopara data has been saved")
	}
	tx.Commit()
}


func insertUlinterfsuppresscfg(eNodeBId string, baseName string, data []model.Ulinterfsuppresscfg) {
	fmt.Println("[+] Processing Ulinterfsuppresscfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulinterfsuppresscfg` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UlInterfsuppressionThd,
			data[j].ATTRIBUTES.P0UePuschOffset,
			data[j].ATTRIBUTES.RASignalMCSOffset,
			data[j].ATTRIBUTES.P0UePucchOffset,
			data[j].ATTRIBUTES.DeltaMSG2Offset,
			data[j].ATTRIBUTES.StrongInfUserThdRatio,
			data[j].ATTRIBUTES.ULInfStatisticPeriod,
			data[j].ATTRIBUTES.ULInfStatisticPeriodNum,
			data[j].ATTRIBUTES.VolteMCSOffset,
			data[j].ATTRIBUTES.VoltePucchPcTarSinrOffset,
			data[j].ATTRIBUTES.VoltePuschPsdCtrlTarget,
			data[j].ATTRIBUTES.RemoteInfULEnhanceSw,
			data[j].ATTRIBUTES.PuschEnhDeltaSinrThd,
			data[j].ATTRIBUTES.RASignalPcSwitch,
			data[j].ATTRIBUTES.UlInterPwrDiffThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulinterfsuppresscfg data has been saved")
	}
	tx.Commit()
}


func insertUlocell(eNodeBId string, baseName string, data []model.Ulocell) {
	fmt.Println("[+] Processing Ulocell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.VIRTUALCPC,
			data[j].ATTRIBUTES.RADIUS,
			data[j].ATTRIBUTES.BOOST,
			data[j].ATTRIBUTES.HORAD,
			data[j].ATTRIBUTES.TIMELIMIT,
			data[j].ATTRIBUTES.ULCOVEXPANSION,
			data[j].ATTRIBUTES.ULBASEBANDEQMID,
			data[j].ATTRIBUTES.PRECISECHEST2MS,
			data[j].ATTRIBUTES.DLBASEBANDEQMID,
			data[j].ATTRIBUTES.INTERNODEBULCOMP,
			data[j].ATTRIBUTES.ULFREQIND,
			data[j].ATTRIBUTES.SPEEDBASEDDEMODSW,
			data[j].ATTRIBUTES.OBJID,
			data[j].ATTRIBUTES.NBIS,
			data[j].ATTRIBUTES.FREQUENCYBANDWIDTH,
			data[j].ATTRIBUTES.HSDPA4C,
			data[j].ATTRIBUTES.DL64QAM,
			data[j].ATTRIBUTES.BBUNITREF,
			data[j].ATTRIBUTES.MULTIRRUSTATICDESSW,
			data[j].ATTRIBUTES.AIR,
			data[j].ATTRIBUTES.TTW,
			data[j].ATTRIBUTES.LOCELLPRI,
			data[j].ATTRIBUTES.LOCELLTYPE,
			data[j].ATTRIBUTES.TURBOIC,
			data[j].ATTRIBUTES.ERACHHSDPCCH,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.ULFREQ,
			data[j].ATTRIBUTES.DLFREQ,
			data[j].ATTRIBUTES.MAXPWR,
			data[j].ATTRIBUTES.DEFPWRLVL,
			data[j].ATTRIBUTES.DLRESMODE,
			data[j].ATTRIBUTES.DI,
			data[j].ATTRIBUTES.HISPM,
			data[j].ATTRIBUTES.RMTCM,
			data[j].ATTRIBUTES.UL16QAM,
			data[j].ATTRIBUTES.FDEMODE,
			data[j].ATTRIBUTES.ULL2PLUS,
			data[j].ATTRIBUTES.ICMODE,
			data[j].ATTRIBUTES.ERACH,
			data[j].ATTRIBUTES.RSV,
			data[j].ATTRIBUTES.GUPOWERSHARE,
			data[j].ATTRIBUTES.HSPAUSERNUMEXT,
			data[j].ATTRIBUTES.IRCSW,
			data[j].ATTRIBUTES.TURBOICPHASE2,
			data[j].ATTRIBUTES.MFHSDPASW,
			data[j].ATTRIBUTES.ULCOMPSW,
			data[j].ATTRIBUTES.VAM,
			data[j].ATTRIBUTES.CELLSCALEIND,
			data[j].ATTRIBUTES.INTERNODEBHSDPCCHCOMP,
			data[j].ATTRIBUTES.ULSUPERNARROWBANDFILTER,
			data[j].ATTRIBUTES.DLADAPTIVEBPFILTER,
			data[j].ATTRIBUTES.GSMQLBASEDROTDYNACTRL,
			data[j].ATTRIBUTES.INTEL2TVAM,
			data[j].ATTRIBUTES.DLASYMFILTER,
			data[j].ATTRIBUTES.DLASYMFILTERNFCELLID,
			data[j].ATTRIBUTES.DLASYMLEFTBANDWIDTH,
			data[j].ATTRIBUTES.DLASYMRIGHTBANDWIDTH,
			data[j].ATTRIBUTES.NBISADAPTIVEFORBIDSW,
			data[j].ATTRIBUTES.PRECDWSRSW,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocell data has been saved")
	}
	tx.Commit()
}


func insertUlocellalgpara(eNodeBId string, baseName string, data []model.Ulocellalgpara) {
	fmt.Println("[+] Processing Ulocellalgpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellalgpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.RTWPSIRTGTADJSW,
			data[j].ATTRIBUTES.SIB7RTWPOPTSW,
			data[j].ATTRIBUTES.ANTIANTENNAIMBALANCESW,
			data[j].ATTRIBUTES.R99FLOWCTRLSW,
			data[j].ATTRIBUTES.CQISAMPLEPERIOD,
			data[j].ATTRIBUTES.CQIUSERNUM,
			data[j].ATTRIBUTES.HIGHSPEEDRAKESW,
			data[j].ATTRIBUTES.DPCHMAXTXPWRRESTRSW,
			data[j].ATTRIBUTES.DPCHMAXPWRRTRLOADSTAT,
			data[j].ATTRIBUTES.SIB7MAXRTWP,
			data[j].ATTRIBUTES.ULCCHOLPCSW,
			data[j].ATTRIBUTES.ULCCHOLPCLOADHIGHTHD,
			data[j].ATTRIBUTES.ULCCHOLPCLOADLOWTHD,
			data[j].ATTRIBUTES.RADIOQUALITYDPCHPCSW,
			data[j].ATTRIBUTES.RADIOQUALITYDPCHPCLOADSTAT,
			data[j].ATTRIBUTES.RADIOQUALITYDPCHPCPO,
			data[j].ATTRIBUTES.HTOHPWSHSW,
			data[j].ATTRIBUTES.SHARINGMARGIN,
			data[j].ATTRIBUTES.MAXSHARINGRATIO,
			data[j].ATTRIBUTES.SRBOVERHSDPAOPTSW,
			data[j].ATTRIBUTES.DLPWRLOADTHD,
			data[j].ATTRIBUTES.SINGLEHARQACTTHRD,
			data[j].ATTRIBUTES.R99TOHPWSHSW,
			data[j].ATTRIBUTES.HTOHRTPWSHSW,
			data[j].ATTRIBUTES.DPCHTPCPOADJSW,
			data[j].ATTRIBUTES.HSDPACELLTHPTHD,
			data[j].ATTRIBUTES.HSDPAPWRMGNCANCELSW,
			data[j].ATTRIBUTES.RADIOQUALITYDPCHPCSHOSW,
			data[j].ATTRIBUTES.ULCARLEVELSCHSW,
			data[j].ATTRIBUTES.RADIOQUALITYDPCHPCENHSW,
			data[j].ATTRIBUTES.RADIOQUALITYFDPCHPCPO,
			data[j].ATTRIBUTES.COVERIMPBASEDONADVDEM,
			data[j].ATTRIBUTES.SRBOVERHSDPAMMTOPTSW,
			data[j].ATTRIBUTES.TURBOICENHANCEDSW,
			data[j].ATTRIBUTES.AMRIMPBASEDONPLVASW,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellalgpara data has been saved")
	}
	tx.Commit()
}


func insertUlocellmacepara(eNodeBId string, baseName string, data []model.Ulocellmacepara) {
	fmt.Println("[+] Processing Ulocellmacepara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellmacepara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.SCHEDULEPARA,
			data[j].ATTRIBUTES.HSUPAOLSCHSW,
			data[j].ATTRIBUTES.MAXDELTAOFTARGETROT,
			data[j].ATTRIBUTES.EAGCHPCPARA,
			data[j].ATTRIBUTES.EAGCHPCMOD,
			data[j].ATTRIBUTES.EAGCHPWROFFSET,
			data[j].ATTRIBUTES.EAGCHPOWER,
			data[j].ATTRIBUTES.MAXAGCHPOWER,
			data[j].ATTRIBUTES.MINAGCHPOWER,
			data[j].ATTRIBUTES.SERGCHPCPARA,
			data[j].ATTRIBUTES.SERGCHPCMOD,
			data[j].ATTRIBUTES.SERGCHPWROFFSET,
			data[j].ATTRIBUTES.SERGCHPOWER,
			data[j].ATTRIBUTES.NSERGCHPCPARA,
			data[j].ATTRIBUTES.NSERGCHPCMOD,
			data[j].ATTRIBUTES.NSERGCHPWROFFSET,
			data[j].ATTRIBUTES.NSERGCHPOWER,
			data[j].ATTRIBUTES.SEHICHPCPARA,
			data[j].ATTRIBUTES.SEHICHPWROFFSET,
			data[j].ATTRIBUTES.SEHICHPOWER,
			data[j].ATTRIBUTES.SRLEHICHPWROFFSET,
			data[j].ATTRIBUTES.SINGLERLEHICHPOWER,
			data[j].ATTRIBUTES.NSEHICHPCPARA,
			data[j].ATTRIBUTES.NSEHICHPWROFFSET,
			data[j].ATTRIBUTES.NSEHICHPOWER,
			data[j].ATTRIBUTES.OUTERSYSINTERSCHSW,
			data[j].ATTRIBUTES.OWNCELLULLOADRATIO,
			data[j].ATTRIBUTES.RTWPMEAOPTIMSW,
			data[j].ATTRIBUTES.LOADTHRESH4MINULCOV,
			data[j].ATTRIBUTES.HSUPATDSCHSW,
			data[j].ATTRIBUTES.HSUPATDALIGNUENUM,
			data[j].ATTRIBUTES.HSUPATDSCHENSW,
			data[j].ATTRIBUTES.RTWPSINGLEHARQSCHSW,
			data[j].ATTRIBUTES.ULSECCELLACTOPTSW,
			data[j].ATTRIBUTES.ULSECCELLDEACTCOVTHR,
			data[j].ATTRIBUTES.HSUPATOTALTHRGUARANTEESW,
			data[j].ATTRIBUTES.HSUPATOTALTHRTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellmacepara data has been saved")
	}
	tx.Commit()
}


func insertUlocellmachspara(eNodeBId string, baseName string, data []model.Ulocellmachspara) {
	fmt.Println("[+] Processing Ulocellmachspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellmachspara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.RSCALLOCM,
			data[j].ATTRIBUTES.SM,
			data[j].ATTRIBUTES.PWRMGN,
			data[j].ATTRIBUTES.HSSCCHPWRCMINDCH,
			data[j].ATTRIBUTES.SCCHPWR,
			data[j].ATTRIBUTES.HSSCCHFERTRGTINDCH,
			data[j].ATTRIBUTES.RSCLMSW,
			data[j].ATTRIBUTES.DYNCODESW,
			data[j].ATTRIBUTES.CME16QAMSW,
			data[j].ATTRIBUTES.MXPWRPHUSR,
			data[j].ATTRIBUTES.CQIFA,
			data[j].ATTRIBUTES.MAXEFACHHARQRT,
			data[j].ATTRIBUTES.MAXNONCONVERHARQRT,
			data[j].ATTRIBUTES.HSSCCHPWRCMINEFACH,
			data[j].ATTRIBUTES.CQIADJALGOFCON,
			data[j].ATTRIBUTES.CQIADJALGOFNONCON,
			data[j].ATTRIBUTES.RBLERTARGET,
			data[j].ATTRIBUTES.IBLERTARGET,
			data[j].ATTRIBUTES.SECCELLACTDEASW,
			data[j].ATTRIBUTES.IICSW,
			data[j].ATTRIBUTES.DTXALGOSW,
			data[j].ATTRIBUTES.MIMOPRIMESW,
			data[j].ATTRIBUTES.LOCWEIGHT,
			data[j].ATTRIBUTES.EXTRAPOWER,
			data[j].ATTRIBUTES.MAXEFACHHSHARQRT,
			data[j].ATTRIBUTES.BURSTBLEROPTSW,
			data[j].ATTRIBUTES.DLBESCHWHENULURLBSW,
			data[j].ATTRIBUTES.EXTRAPOWERFORSRB,
			data[j].ATTRIBUTES.CODEOPTSW,
			data[j].ATTRIBUTES.PREALLOCBWOPTISW,
			data[j].ATTRIBUTES.SFDCMACEHSDSCTMR,
			data[j].ATTRIBUTES.DF3CMACEHSDSCTMR,
			data[j].ATTRIBUTES.HSDPAL2OPTSW,
			data[j].ATTRIBUTES.PWRMGNDELTAFORPT,
			data[j].ATTRIBUTES.VIDEOCQITHD,
			data[j].ATTRIBUTES.HSDPAFULLBUFDATATHD,
			data[j].ATTRIBUTES.HSDPAFULLBUFTHROUTHD,
			data[j].ATTRIBUTES.HSSCCHCQIPWRINDCHOPTSW,
			data[j].ATTRIBUTES.MAXEXTRAPWRFORHSSCCHINDCH,
			data[j].ATTRIBUTES.SERVICECQITHD,
			data[j].ATTRIBUTES.EFACHDTCHRSCALLOCMODE,
			data[j].ATTRIBUTES.CQIADJBYIBLERADJSW,
			data[j].ATTRIBUTES.CQIADJBYDYNBLERADJSW,
			data[j].ATTRIBUTES.ACTSECCELLTHPTHD,
			data[j].ATTRIBUTES.DEASECCELLTHPTHD,
			data[j].ATTRIBUTES.ACTSECCELLDATABUFTHD,
			data[j].ATTRIBUTES.DEASECCELLDATABUFTHD,
			data[j].ATTRIBUTES.INDEPTCARSCHSW,
			data[j].ATTRIBUTES.BUFFERSMSW,
			data[j].ATTRIBUTES.DYNCODEOPTSW,
			data[j].ATTRIBUTES.HARQRTSCHOPTSW,
			data[j].ATTRIBUTES.MCDELTACQITHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellmachspara data has been saved")
	}
	tx.Commit()
}


func insertUlocellnoaccesspara(eNodeBId string, baseName string, data []model.Ulocellnoaccesspara) {
	fmt.Println("[+] Processing Ulocellnoaccesspara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellnoaccesspara` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.NOUETIMER,
			data[j].ATTRIBUTES.NOFSTRLTIMER,
			data[j].ATTRIBUTES.AUTORCVRMTHD,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellnoaccesspara data has been saved")
	}
	tx.Commit()
}


func insertUlocellr99algpara(eNodeBId string, baseName string, data []model.Ulocellr99algpara) {
	fmt.Println("[+] Processing Ulocellr99algpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellr99algpara` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.PWRCNGCACSW,
			data[j].ATTRIBUTES.SHORTTH,
			data[j].ATTRIBUTES.LONGTH,
			data[j].ATTRIBUTES.SHORTTHFORRTRRC,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellr99algpara data has been saved")
	}
	tx.Commit()
}


func insertUlocellrsclmtpara(eNodeBId string, baseName string, data []model.Ulocellrsclmtpara) {
	fmt.Println("[+] Processing Ulocellrsclmtpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellrsclmtpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.CME8KRSCLMT,
			data[j].ATTRIBUTES.CME16KRSCLMT,
			data[j].ATTRIBUTES.CME32KRSCLMT,
			data[j].ATTRIBUTES.CME64KRSCLMT,
			data[j].ATTRIBUTES.CME128KRSCLMT,
			data[j].ATTRIBUTES.CME256KRSCLMT,
			data[j].ATTRIBUTES.CME384KRSCLMT,
			data[j].ATTRIBUTES.CME512KRSCLMT,
			data[j].ATTRIBUTES.CME768KRSCLMT,
			data[j].ATTRIBUTES.CME1024KRSCLMT,
			data[j].ATTRIBUTES.CME1536KRSCLMT,
			data[j].ATTRIBUTES.CME1800KRSCLMT,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellrsclmtpara data has been saved")
	}
	tx.Commit()
}


func insertUlocellrsvdpara(eNodeBId string, baseName string, data []model.Ulocellrsvdpara) {
	fmt.Println("[+] Processing Ulocellrsvdpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellrsvdpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA1,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA2,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA3,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA4,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA5,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA6,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA7,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA8,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA9,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA10,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA11,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA12,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA13,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA14,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA15,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA16,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA17,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA18,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA19,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA20,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA21,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA22,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA23,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA24,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA25,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA26,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA27,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA28,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA29,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA30,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA31,
			data[j].ATTRIBUTES.ULOCELLRSVDPARA32,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellrsvdpara data has been saved")
	}
	tx.Commit()
}


func insertUlocellsectoreqm(eNodeBId string, baseName string, data []model.Ulocellsectoreqm) {
	fmt.Println("[+] Processing Ulocellsectoreqm data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulocellsectoreqm` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULOCELLID,
			data[j].ATTRIBUTES.SECTOREQMID,
			data[j].ATTRIBUTES.MAXPWR,
			data[j].ATTRIBUTES.SECTOREQMPROPERTY,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulocellsectoreqm data has been saved")
	}
	tx.Commit()
}


func insertUlzerobufferzone(eNodeBId string, baseName string, data []model.Ulzerobufferzone) {
	fmt.Println("[+] Processing Ulzerobufferzone data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `ulzerobufferzone` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.ULSharedFreqStartRB,
			data[j].ATTRIBUTES.ULSharedFreqEndRB,
			data[j].ATTRIBUTES.ULZeroBufZonePRBThd,
			data[j].ATTRIBUTES.ULZeroBufZonePRBOffset,
			data[j].ATTRIBUTES.ULZeroBufZoneRsrpThd,
			data[j].ATTRIBUTES.ULZeroBufZoneRsrpOffset,
			data[j].ATTRIBUTES.ULZeroBufZoneB1RmvOffset,
			data[j].ATTRIBUTES.ULZeroBufZoneUtranRscpThd,
			data[j].ATTRIBUTES.ULZeroBufZoInterRatB1Timer,	

		)
		checkErr(err)
		//fmt.Println("[+] Ulzerobufferzone data has been saved")
	}
	tx.Commit()
}


func insertUpptsinterfcfg(eNodeBId string, baseName string, data []model.Upptsinterfcfg) {
	fmt.Println("[+] Processing Upptsinterfcfg data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `upptsinterfcfg` VALUES(?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TestPeriod,
			data[j].ATTRIBUTES.TestThreshold,
			data[j].ATTRIBUTES.TestHysteresis,	

		)
		checkErr(err)
		//fmt.Println("[+] Upptsinterfcfg data has been saved")
	}
	tx.Commit()
}


func insertUsb(eNodeBId string, baseName string, data []model.Usb) {
	fmt.Println("[+] Processing Usb data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `usb` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.CN,
			data[j].ATTRIBUTES.SRN,
			data[j].ATTRIBUTES.SN,
			data[j].ATTRIBUTES.PN,
			data[j].ATTRIBUTES.SWITCH,	

		)
		checkErr(err)
		//fmt.Println("[+] Usb data has been saved")
	}
	tx.Commit()
}


func insertUserplanehost(eNodeBId string, baseName string, data []model.Userplanehost) {
	fmt.Println("[+] Processing Userplanehost data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `userplanehost` VALUES(?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UPHOSTID,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.IPVERSION,
			data[j].ATTRIBUTES.LOCIPV4,
			data[j].ATTRIBUTES.IPSECSWITCH,
			data[j].ATTRIBUTES.USERLABEL,	

		)
		checkErr(err)
		//fmt.Println("[+] Userplanehost data has been saved")
	}
	tx.Commit()
}


func insertUserplanepeer(eNodeBId string, baseName string, data []model.Userplanepeer) {
	fmt.Println("[+] Processing Userplanepeer data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `userplanepeer` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.UPPEERID,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.IPVERSION,
			data[j].ATTRIBUTES.PEERIPV4,
			data[j].ATTRIBUTES.IPSECSWITCH,
			data[j].ATTRIBUTES.REMOTEID,
			data[j].ATTRIBUTES.CTRLMODE,
			data[j].ATTRIBUTES.AUTOCFGFLAG,
			data[j].ATTRIBUTES.USERLABEL,
			data[j].ATTRIBUTES.STATICCHK,	

		)
		checkErr(err)
		//fmt.Println("[+] Userplanepeer data has been saved")
	}
	tx.Commit()
}


func insertUserpriority(eNodeBId string, baseName string, data []model.Userpriority) {
	fmt.Println("[+] Processing Userpriority data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `userpriority` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ArpSchSwitch,
			data[j].ATTRIBUTES.Arp1Priority,
			data[j].ATTRIBUTES.Arp2Priority,
			data[j].ATTRIBUTES.Arp3Priority,
			data[j].ATTRIBUTES.Arp4Priority,
			data[j].ATTRIBUTES.Arp5Priority,
			data[j].ATTRIBUTES.Arp6Priority,
			data[j].ATTRIBUTES.Arp7Priority,
			data[j].ATTRIBUTES.Arp8Priority,
			data[j].ATTRIBUTES.Arp9Priority,
			data[j].ATTRIBUTES.Arp10Priority,
			data[j].ATTRIBUTES.Arp11Priority,
			data[j].ATTRIBUTES.Arp12Priority,
			data[j].ATTRIBUTES.Arp13Priority,
			data[j].ATTRIBUTES.Arp14Priority,
			data[j].ATTRIBUTES.Arp15Priority,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Userpriority data has been saved")
	}
	tx.Commit()
}


func insertUserqcipriority(eNodeBId string, baseName string, data []model.Userqcipriority) {
	fmt.Println("[+] Processing Userqcipriority data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `userqcipriority` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Qci,
			data[j].ATTRIBUTES.GoldUlSchPriorityFactor,
			data[j].ATTRIBUTES.GoldDlSchPriorityFactor,
			data[j].ATTRIBUTES.SilverUlSchPriorityFactor,
			data[j].ATTRIBUTES.SilverDlSchPriorityFactor,
			data[j].ATTRIBUTES.BronzeUlSchPriorityFactor,
			data[j].ATTRIBUTES.BronzeDlSchPriorityFactor,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Userqcipriority data has been saved")
	}
	tx.Commit()
}


func insertUservpfmpara(eNodeBId string, baseName string, data []model.Uservpfmpara) {
	fmt.Println("[+] Processing Uservpfmpara data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `uservpfmpara` VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.DlUserThruputThd0,
			data[j].ATTRIBUTES.DlUserThruputThd1,
			data[j].ATTRIBUTES.DlUserThruputThd2,
			data[j].ATTRIBUTES.DlUserThruputThd3,
			data[j].ATTRIBUTES.DlUserThruputThd4,
			data[j].ATTRIBUTES.UlAccessDelayGoodThd,
			data[j].ATTRIBUTES.UlAccessDelayBadThd,
			data[j].ATTRIBUTES.DlAccessDelayGoodThd,
			data[j].ATTRIBUTES.DlAccessDelayBadThd,	

		)
		checkErr(err)
		//fmt.Println("[+] Uservpfmpara data has been saved")
	}
	tx.Commit()
}


func insertUtranexternalcell(eNodeBId string, baseName string, data []model.Utranexternalcell) {
	fmt.Println("[+] Processing Utranexternalcell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `utranexternalcell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.RncId,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.UtranDlArfcn,
			data[j].ATTRIBUTES.UtranUlArfcnCfgInd,
			data[j].ATTRIBUTES.UtranFddTddType,
			data[j].ATTRIBUTES.RacCfgInd,
			data[j].ATTRIBUTES.Rac,
			data[j].ATTRIBUTES.PScrambCode,
			data[j].ATTRIBUTES.Lac,
			data[j].ATTRIBUTES.CellName,
			data[j].ATTRIBUTES.CsPsHOInd,
			data[j].ATTRIBUTES.UtranExternalCellSlaveBand,
			data[j].ATTRIBUTES.RoamingAreaHoInd,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.ObjId,
			data[j].ATTRIBUTES.UtranUlArfcn,	

		)
		checkErr(err)
		//fmt.Println("[+] Utranexternalcell data has been saved")
	}
	tx.Commit()
}


func insertUtranncell(eNodeBId string, baseName string, data []model.Utranncell) {
	fmt.Println("[+] Processing Utranncell data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `utranncell` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.RncId,
			data[j].ATTRIBUTES.CellId,
			data[j].ATTRIBUTES.NoHoFlag,
			data[j].ATTRIBUTES.NoRmvFlag,
			data[j].ATTRIBUTES.AnrFlag,
			data[j].ATTRIBUTES.BlindHoPriority,
			data[j].ATTRIBUTES.CellMeasPriority,
			data[j].ATTRIBUTES.OverlapInd,
			data[j].ATTRIBUTES.NCellMeasPriority,
			data[j].ATTRIBUTES.LocalCellName,
			data[j].ATTRIBUTES.NeighbourCellName,
			data[j].ATTRIBUTES.CtrlMode,	

		)
		checkErr(err)
		//fmt.Println("[+] Utranncell data has been saved")
	}
	tx.Commit()
}


func insertUtrannfreq(eNodeBId string, baseName string, data []model.Utrannfreq) {
	fmt.Println("[+] Processing Utrannfreq data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `utrannfreq` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UtranDlArfcn,
			data[j].ATTRIBUTES.UtranVersion,
			data[j].ATTRIBUTES.UtranFddTddType,
			data[j].ATTRIBUTES.UtranUlArfcnCfgInd,
			data[j].ATTRIBUTES.CellReselPriorityCfgInd,
			data[j].ATTRIBUTES.CellReselPriority,
			data[j].ATTRIBUTES.PmaxUtran,
			data[j].ATTRIBUTES.OffsetFreq,
			data[j].ATTRIBUTES.Qqualmin,
			data[j].ATTRIBUTES.QRxLevMin,
			data[j].ATTRIBUTES.ThreshXHigh,
			data[j].ATTRIBUTES.ThreshXLow,
			data[j].ATTRIBUTES.ThreshXHighQ,
			data[j].ATTRIBUTES.ThreshXLowQ,
			data[j].ATTRIBUTES.PsPriority,
			data[j].ATTRIBUTES.CsPriority,
			data[j].ATTRIBUTES.ConnFreqPriority,
			data[j].ATTRIBUTES.CsPsMixedPriority,
			data[j].ATTRIBUTES.ContinuCoverageIndication,
			data[j].ATTRIBUTES.MlbFreqPriority,
			data[j].ATTRIBUTES.MasterBandFlag,
			data[j].ATTRIBUTES.UtranRanSharingInd,
			data[j].ATTRIBUTES.AnrInd,
			data[j].ATTRIBUTES.SrvccPriority,
			data[j].ATTRIBUTES.ULSharedFreqInd,
			data[j].ATTRIBUTES.FreqPriorityForAnr,
			data[j].ATTRIBUTES.MlbTargetInd,
			data[j].ATTRIBUTES.UtranFreqHighSpeedFlag,	

		)
		checkErr(err)
		//fmt.Println("[+] Utrannfreq data has been saved")
	}
	tx.Commit()
}


func insertUtranranshare(eNodeBId string, baseName string, data []model.Utranranshare) {
	fmt.Println("[+] Processing Utranranshare data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `utranranshare` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.UtranDlArfcn,
			data[j].ATTRIBUTES.Mcc,
			data[j].ATTRIBUTES.Mnc,
			data[j].ATTRIBUTES.CellReselPriorityCfgInd,	

		)
		checkErr(err)
		//fmt.Println("[+] Utranranshare data has been saved")
	}
	tx.Commit()
}


func insertVlanclass(eNodeBId string, baseName string, data []model.Vlanclass) {
	fmt.Println("[+] Processing Vlanclass data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `vlanclass` VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.VLANGROUPNO,
			data[j].ATTRIBUTES.TRAFFIC,
			data[j].ATTRIBUTES.SRVPRIO,
			data[j].ATTRIBUTES.VLANID,
			data[j].ATTRIBUTES.VLANPRIO,	

		)
		checkErr(err)
		//fmt.Println("[+] Vlanclass data has been saved")
	}
	tx.Commit()
}


func insertVlanmap(eNodeBId string, baseName string, data []model.Vlanmap) {
	fmt.Println("[+] Processing Vlanmap data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `vlanmap` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.NEXTHOPIP,
			data[j].ATTRIBUTES.MASK,
			data[j].ATTRIBUTES.VLANMODE,
			data[j].ATTRIBUTES.VLANID,
			data[j].ATTRIBUTES.VLANPRIO,
			data[j].ATTRIBUTES.SETPRIO,
			data[j].ATTRIBUTES.VLANGROUPNO,	

		)
		checkErr(err)
		//fmt.Println("[+] Vlanmap data has been saved")
	}
	tx.Commit()
}


func insertVqmalgo(eNodeBId string, baseName string, data []model.Vqmalgo) {
	fmt.Println("[+] Processing Vqmalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `vqmalgo` VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.ULDelayJitter,
			data[j].ATTRIBUTES.VqiGoodThd,
			data[j].ATTRIBUTES.VqiPoorThd,
			data[j].ATTRIBUTES.VqiBadThd,
			data[j].ATTRIBUTES.AmrWbRadioLowQualRelThd,
			data[j].ATTRIBUTES.AmrNbRadioLowQualRelThd,
			data[j].ATTRIBUTES.AmrSilentPeriodNum,
			data[j].ATTRIBUTES.AmrLowQualRelPeriodNum,
			data[j].ATTRIBUTES.VqiExcellentThd,
			data[j].ATTRIBUTES.AmrNbSilentThd,
			data[j].ATTRIBUTES.AmrWbRcvLowQualRelThd,
			data[j].ATTRIBUTES.AmrNbRcvLowQualRelThd,
			data[j].ATTRIBUTES.AmrWbSilentThd,
			data[j].ATTRIBUTES.VQMAlgoPeriod,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] Vqmalgo data has been saved")
	}
	tx.Commit()
}


func insertVrf(eNodeBId string, baseName string, data []model.Vrf) {
	fmt.Println("[+] Processing Vrf data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `vrf` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.VRFIDX,
			data[j].ATTRIBUTES.USERLABEL,	

		)
		checkErr(err)
		//fmt.Println("[+] Vrf data has been saved")
	}
	tx.Commit()
}


func insertWeblmt(eNodeBId string, baseName string, data []model.Weblmt) {
	fmt.Println("[+] Processing Weblmt data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `weblmt` VALUES(?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.POLICY,
			data[j].ATTRIBUTES.SSLVER,	

		)
		checkErr(err)
		//fmt.Println("[+] Weblmt data has been saved")
	}
	tx.Commit()
}


func insertWtcpproxyalgo(eNodeBId string, baseName string, data []model.Wtcpproxyalgo) {
	fmt.Println("[+] Processing Wtcpproxyalgo data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `wtcpproxyalgo` VALUES(?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.LocalCellId,
			data[j].ATTRIBUTES.TcpAccelerationSwitch,
			data[j].ATTRIBUTES.TCPStatisticsSwitch,
			data[j].ATTRIBUTES.MaxRttStatisticsThd,
			data[j].ATTRIBUTES.MaxProxyPktNum,
			data[j].ATTRIBUTES.AckSplitCount,
			data[j].ATTRIBUTES.TcpAsSchWeightFactor,	

		)
		checkErr(err)
		//fmt.Println("[+] Wtcpproxyalgo data has been saved")
	}
	tx.Commit()
}


func insertX2(eNodeBId string, baseName string, data []model.X2) {
	fmt.Println("[+] Processing X2 data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `x2` VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.X2Id,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.CpEpGroupId,
			data[j].ATTRIBUTES.UpEpGroupId,
			data[j].ATTRIBUTES.TargetENodeBRelease,
			data[j].ATTRIBUTES.UserLabel,
			data[j].ATTRIBUTES.EpGroupCfgFlag,
			data[j].ATTRIBUTES.CnOpSharingGroupId,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] X2 data has been saved")
	}
	tx.Commit()
}


func insertX2interface(eNodeBId string, baseName string, data []model.X2interface) {
	fmt.Println("[+] Processing X2interface data")
	tx, err := db.Begin()
	checkErr(err)
	stmt, err = tx.Prepare("INSERT IGNORE INTO `x2interface` VALUES(?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	for j = 0; j < len(data); j++ {

		_, err = stmt.Exec(eNodeBId, baseName,
			data[j].ATTRIBUTES.X2InterfaceId,
			data[j].ATTRIBUTES.X2CpBearerId,
			data[j].ATTRIBUTES.CnOperatorId,
			data[j].ATTRIBUTES.TargetENodeBRelease,
			data[j].ATTRIBUTES.CtrlMode,
			data[j].ATTRIBUTES.AutoCfgFlag,
			data[j].ATTRIBUTES.CnOpSharingGroupId,
			data[j].ATTRIBUTES.ObjId,	

		)
		checkErr(err)
		//fmt.Println("[+] X2interface data has been saved")
	}
	tx.Commit()
}


