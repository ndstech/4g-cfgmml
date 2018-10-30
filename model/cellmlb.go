package model

import "encoding/xml"

type Cellmlb struct {
	XMLName xml.Name `xml:"CellMLB"`
	ATTRIBUTES CellmlbAttributes `xml:"attributes"`
}

type CellmlbAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterFreqMlbThd string `xml:"InterFreqMlbThd"`
	InterRatMlbThd string `xml:"InterRatMlbThd"`
	LoadOffset string `xml:"LoadOffset"`
	LoadDiffThd string `xml:"LoadDiffThd"`
	InterRatMlbUeNumThd string `xml:"InterRatMlbUeNumThd"`
	InitValidPeriod string `xml:"InitValidPeriod"`
	LoadTransferFactor string `xml:"LoadTransferFactor"`
	MlbTriggerMode string `xml:"MlbTriggerMode"`
	InterFreqMlbUeNumThd string `xml:"InterFreqMlbUeNumThd"`
	MlbUeNumOffset string `xml:"MlbUeNumOffset"`
	MlbMaxUeNum string `xml:"MlbMaxUeNum"`
	MlbUeSelectPrbThd string `xml:"MlbUeSelectPrbThd"`
	DlDataMlbMode string `xml:"DlDataMlbMode"`
	InterFreqMLBRanShareMode string `xml:"InterFreqMLBRanShareMode"`
	UeNumDiffThd string `xml:"UeNumDiffThd"`
	HotSpotUeMode string `xml:"HotSpotUeMode"`
	InterFreqIdleMlbMode string `xml:"InterFreqIdleMlbMode"`
	MlbMinUeNumThd string `xml:"MlbMinUeNumThd"`
	MlbMinUeNumOffset string `xml:"MlbMinUeNumOffset"`
	InterFreqUeTrsfType string `xml:"InterFreqUeTrsfType"`
	InterFreqIdleMlbUeNumThd string `xml:"InterFreqIdleMlbUeNumThd"`
	InterRatIdleMlbUeNumThd string `xml:"InterRatIdleMlbUeNumThd"`
	InterFreqLoadEvalPrd string `xml:"InterFreqLoadEvalPrd"`
	InterRatLoadEvalPrd string `xml:"InterRatLoadEvalPrd"`
	FreqSelectStrategy string `xml:"FreqSelectStrategy"`
	LoadBalanceNCellScope string `xml:"LoadBalanceNCellScope"`
	InterRatMlbUeNumOffset string `xml:"InterRatMlbUeNumOffset"`
	IdleUeSelFreqScope string `xml:"IdleUeSelFreqScope"`
	InterRatMlbUeSelStrategy string `xml:"InterRatMlbUeSelStrategy"`
	InterRatMlbUeSelPrbThd string `xml:"InterRatMlbUeSelPrbThd"`
	PrbLoadCalcMethod string `xml:"PrbLoadCalcMethod"`
	MlbUeSelectPunishTimer string `xml:"MlbUeSelectPunishTimer"`
	MlbHoCellSelectStrategy string `xml:"MlbHoCellSelectStrategy"`
	InterRatMlbTriggerMode string `xml:"InterRatMlbTriggerMode"`
	InterRatMlbUeNumModeThd string `xml:"InterRatMlbUeNumModeThd"`
	PunishJudgePrdNum string `xml:"PunishJudgePrdNum"`
	FreqPunishPrdNum string `xml:"FreqPunishPrdNum"`
	CellPunishPrdNum string `xml:"CellPunishPrdNum"`
	MultiRruMode string `xml:"MultiRruMode"`
	InterFreqOffloadOffset string `xml:"InterFreqOffloadOffset"`
	InterFrqUeNumOffloadOffset string `xml:"InterFrqUeNumOffloadOffset"`
	CellCapacityScaleFactor string `xml:"CellCapacityScaleFactor"`
	InterRatMlbMaxUeNum string `xml:"InterRatMlbMaxUeNum"`
	InterRatMlbHoFailPunish string `xml:"InterRatMlbHoFailPunish"`
	EutranLoadJudgeStrategy string `xml:"EutranLoadJudgeStrategy"`
	MlbTrigJudgePeriod string `xml:"MlbTrigJudgePeriod"`
	InterFreqMlbStrategy string `xml:"InterFreqMlbStrategy"`
	MaxSpectralEfficiencyValue string `xml:"MaxSpectralEfficiencyValue"`
	MinSpectralEfficiencyValue string `xml:"MinSpectralEfficiencyValue"`
	SpectralEffAdjustMaxStep string `xml:"SpectralEffAdjustMaxStep"`
	UeNumDiffOffsetTransCa string `xml:"UeNumDiffOffsetTransCa"`
	MlbIdleUeNumAdjFactor string `xml:"MlbIdleUeNumAdjFactor"`
	IdleMlbUeNumDiffThd string `xml:"IdleMlbUeNumDiffThd"`
	L2USmartOffloadThd string `xml:"L2USmartOffloadThd"`
	L2USmartOffloadOffset string `xml:"L2USmartOffloadOffset"`
	InterFreqMlbUlThd string `xml:"InterFreqMlbUlThd"`
	UeDlPrbLowThdOffset string `xml:"UeDlPrbLowThdOffset"`
	UeUlPrbHighThdOffset string `xml:"UeUlPrbHighThdOffset"`
	UeUlPrbLowThdOffset string `xml:"UeUlPrbLowThdOffset"`
	VideoLoadHighThd string `xml:"VideoLoadHighThd"`
	VideoLoadLowThd string `xml:"VideoLoadLowThd"`
	VideoDlPrbThd string `xml:"VideoDlPrbThd"`
	InterFIdleUeNumOffloadOfs string `xml:"InterFIdleUeNumOffloadOfs"`
	UlExperienceDiffThd string `xml:"UlExperienceDiffThd"`
	UlExperienceEvalPrd string `xml:"UlExperienceEvalPrd"`
	UlExperienceMaxUeNum string `xml:"UlExperienceMaxUeNum"`
	UlExperienceOffloadThd string `xml:"UlExperienceOffloadThd"`
	UlExperienceOffset string `xml:"UlExperienceOffset"`
}

