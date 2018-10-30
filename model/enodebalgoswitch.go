package model

import "encoding/xml"

type Enodebalgoswitch struct {
	XMLName xml.Name `xml:"ENodeBAlgoSwitch"`
	ATTRIBUTES EnodebalgoswitchAttributes `xml:"attributes"`
}

type EnodebalgoswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	HoAlgoSwitch string `xml:"HoAlgoSwitch"`
	HoModeSwitch string `xml:"HoModeSwitch"`
	DlIcicSwitch string `xml:"DlIcicSwitch"`
	AnrSwitch string `xml:"AnrSwitch"`
	RedirectSwitch string `xml:"RedirectSwitch"`
	MroSwitch string `xml:"MroSwitch"`
	MacAssemblyOptSwitch string `xml:"MacAssemblyOptSwitch"`
	TpeSwitch string `xml:"TpeSwitch"`
	SpidSelectPlmnAlgoSwitch string `xml:"SpidSelectPlmnAlgoSwitch"`
	UlIcicFreqSwitch string `xml:"UlIcicFreqSwitch"`
	LcsSwitch string `xml:"LcsSwitch"`
	TrmSwitch string `xml:"TrmSwitch"`
	PciConflictAlmSwitch string `xml:"PciConflictAlmSwitch"`
	PowerSaveSwitch string `xml:"PowerSaveSwitch"`
	RimSwitch string `xml:"RimSwitch"`
	RanSharingAnrSwitch string `xml:"RanSharingAnrSwitch"`
	FreqLayerSwtich string `xml:"FreqLayerSwtich"`
	CmasSwitch string `xml:"CmasSwitch"`
	VQMAlgoSwitch string `xml:"VQMAlgoSwitch"`
	UeNumPreemptSwitch string `xml:"UeNumPreemptSwitch"`
	MultiOpCtrlSwitch string `xml:"MultiOpCtrlSwitch"`
	OverBBUsSwitch string `xml:"OverBBUsSwitch"`
	OperatorSpecificAlgoSwitch string `xml:"OperatorSpecificAlgoSwitch"`
	HoSignalingOptSwitch string `xml:"HoSignalingOptSwitch"`
	CompatibilityCtrlSwitch string `xml:"CompatibilityCtrlSwitch"`
	BlindNcellOptSwitch string `xml:"BlindNcellOptSwitch"`
	RimOnEcoSwitch string `xml:"RimOnEcoSwitch"`
	EutranVoipSupportSwitch string `xml:"EutranVoipSupportSwitch"`
	CaAlgoSwitch string `xml:"CaAlgoSwitch"`
	MlbAlgoSwitch string `xml:"MlbAlgoSwitch"`
	HoCommOptSwitch string `xml:"HoCommOptSwitch"`
	HighLoadNetOptSwitch string `xml:"HighLoadNetOptSwitch"`
	SchOptSwitch string `xml:"SchOptSwitch"`
	PrachTimeStagSwitch string `xml:"PrachTimeStagSwitch"`
	HighSpeedRootSeqCSSwitch string `xml:"HighSpeedRootSeqCSSwitch"`
	DropPktsStatSwitch string `xml:"DropPktsStatSwitch"`
	NCellRankingSwitch string `xml:"NCellRankingSwitch"`
	EUCompactRANSwitch string `xml:"EUCompactRANSwitch"`
	BbpCollaborateSwitch string `xml:"BbpCollaborateSwitch"`
	RootSeqConflictDetSwitch string `xml:"RootSeqConflictDetSwitch"`
	IOptAlgoSwitch string `xml:"IOptAlgoSwitch"`
	ServiceHoMultiTargetFreqSw string `xml:"ServiceHoMultiTargetFreqSw"`
	SFSSwitch string `xml:"SFSSwitch"`
	PciConflictDetectSwitch string `xml:"PciConflictDetectSwitch"`
	CaLbAlgoSwitch string `xml:"CaLbAlgoSwitch"`
	UlSchOptSwitch string `xml:"UlSchOptSwitch"`
	CompactRANMultiAPN string `xml:"CompactRANMultiAPN"`
	ClkJumpCellReStartSwitch string `xml:"ClkJumpCellReStartSwitch"`
	E2EVQIAlgoSwitch string `xml:"E2EVQIAlgoSwitch"`
	TlFreqFrameOffsetSwitch string `xml:"TlFreqFrameOffsetSwitch"`
	FastEnhancePccAnchorSwitch string `xml:"FastEnhancePccAnchorSwitch"`
	HoWithSccCfgAddBlindSwitch string `xml:"HoWithSccCfgAddBlindSwitch"`
	ObjId string `xml:"objId"`
	AntCalibrationTimeSwitch string `xml:"AntCalibrationTimeSwitch"`
	PimSwitch string `xml:"PimSwitch"`
	FltrRptRrcConReqExtdSwitch string `xml:"FltrRptRrcConReqExtdSwitch"`
	SSRDAlgoSwitch string `xml:"SSRDAlgoSwitch"`
	IoTClkJumpCellResetSw string `xml:"IoTClkJumpCellResetSw"`
	CaAlgoExtSwitch string `xml:"CaAlgoExtSwitch"`
	UlResManageOptSw string `xml:"UlResManageOptSw"`
	LTEPreemptNbSwitch string `xml:"LTEPreemptNbSwitch"`
}

