package model

import "encoding/xml"

type Globalprocswitch struct {
	XMLName xml.Name `xml:"GLOBALPROCSWITCH"`
	ATTRIBUTES GlobalprocswitchAttributes `xml:"attributes"`
}

type GlobalprocswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	X2SonSetupSwitch string `xml:"X2SonSetupSwitch"`
	X2SonLinkSetupType string `xml:"X2SonLinkSetupType"`
	SriAdaptiveSwitch string `xml:"SriAdaptiveSwitch"`
	LcgProfile string `xml:"LcgProfile"`
	RncPoolHoRimSwitch string `xml:"RncPoolHoRimSwitch"`
	UuMsgSimulSendSwitch string `xml:"UuMsgSimulSendSwitch"`
	X2BasedUptENodeBCfgSwitch string `xml:"X2BasedUptENodeBCfgSwitch"`
	TargetOpBasedX2Switch string `xml:"TargetOpBasedX2Switch"`
	X2SonTNLSelectMode string `xml:"X2SonTNLSelectMode"`
	UtranLoadTransChan string `xml:"UtranLoadTransChan"`
	S1HoInDataFwdSwitch string `xml:"S1HoInDataFwdSwitch"`
	RrcReestProtectThd string `xml:"RrcReestProtectThd"`
	X2SonDeleteTimer string `xml:"X2SonDeleteTimer"`
	RimCodingPolicy string `xml:"RimCodingPolicy"`
	L2GUHoWithLCapSwitch string `xml:"L2GUHoWithLCapSwitch"`
	DiffOpWithSameMmecSwitch string `xml:"DiffOpWithSameMmecSwitch"`
	PeerReqBasedX2DelSwitch string `xml:"PeerReqBasedX2DelSwitch"`
	UlPdcpSduRcvStatSendSwitch string `xml:"UlPdcpSduRcvStatSendSwitch"`
	ProtocolMsgOptSwitch string `xml:"ProtocolMsgOptSwitch"`
	X2BasedDelNcellCfgSwitch string `xml:"X2BasedDelNcellCfgSwitch"`
	X2ServedCellType string `xml:"X2ServedCellType"`
	EnbTrigMmeLoadRebalSwitch string `xml:"EnbTrigMmeLoadRebalSwitch"`
	UeRelReSynTimes string `xml:"UeRelReSynTimes"`
	UeRelChkLostSwitch string `xml:"UeRelChkLostSwitch"`
	UeLinkAbnormalDetectSwitch string `xml:"UeLinkAbnormalDetectSwitch"`
	S1DefaultPagingDrxSelect string `xml:"S1DefaultPagingDrxSelect"`
	EnhancedPhrRptCtrlSw string `xml:"EnhancedPhrRptCtrlSw"`
	EutranLoadTransSwitch string `xml:"EutranLoadTransSwitch"`
	ProtocolSupportSwitch string `xml:"ProtocolSupportSwitch"`
	CellTrafficTraceMsgSwitch string `xml:"CellTrafficTraceMsgSwitch"`
	PreferSigExtend string `xml:"PreferSigExtend"`
	MmeSelectProcSwitch string `xml:"MmeSelectProcSwitch"`
	S1OfflineCommitSwitch string `xml:"S1OfflineCommitSwitch"`
	ProtocolCompatibilitySw string `xml:"ProtocolCompatibilitySw"`
	X2SctpEstType string `xml:"X2SctpEstType"`
	X2SonDeleteSwitch string `xml:"X2SonDeleteSwitch"`
	X2SonDeleteTimerforX2Fault string `xml:"X2SonDeleteTimerforX2Fault"`
	X2SonDeleteTimerforX2Usage string `xml:"X2SonDeleteTimerforX2Usage"`
	X2SonDeleteHoInNumThd string `xml:"X2SonDeleteHoInNumThd"`
	X2SonDeleteHoOutNumThd string `xml:"X2SonDeleteHoOutNumThd"`
	VoipWithGapMode string `xml:"VoipWithGapMode"`
	IntraEnodebHoStaticSw string `xml:"IntraEnodebHoStaticSw"`
	MaxSyncUserNumPerBbi string `xml:"MaxSyncUserNumPerBbi"`
	X2SetupFailureTimetoWait string `xml:"X2SetupFailureTimetoWait"`
	X2SetupFailureNumThd string `xml:"X2SetupFailureNumThd"`
	X2SetupFailureNumTimer string `xml:"X2SetupFailureNumTimer"`
	EX2AutoDeleteSwitch string `xml:"eX2AutoDeleteSwitch"`
	EX2AutoDeleteTimerforFault string `xml:"eX2AutoDeleteTimerforFault"`
	X2SonDeleteMode string `xml:"X2SonDeleteMode"`
	X2SonSetupNumThd string `xml:"X2SonSetupNumThd"`
	X2SonSetupTimer string `xml:"X2SonSetupTimer"`
	SecKeyRecfgSwitch string `xml:"SecKeyRecfgSwitch"`
	X2BasedUptENodeBPolicy string `xml:"X2BasedUptENodeBPolicy"`
	MMEDomNameMode string `xml:"MMEDomNameMode"`
	RrcConnPunishThd string `xml:"RrcConnPunishThd"`
	X2BasedUptNcellCfgSwitch string `xml:"X2BasedUptNcellCfgSwitch"`
	HoProcCtrlSwitch string `xml:"HoProcCtrlSwitch"`
	RrcReestOptSwitch string `xml:"RrcReestOptSwitch"`
	S1MMESonSwitch string `xml:"S1MMESonSwitch"`
	PrivateMdtUeSelSwitch string `xml:"PrivateMdtUeSelSwitch"`
	QciUpdParaCheckSwitch string `xml:"QciUpdParaCheckSwitch"`
	UeCompatSwitch string `xml:"UeCompatSwitch"`
	S1MmePrivFeatureInd string `xml:"S1MmePrivFeatureInd"`
	SctpAbortSmoothSwitch string `xml:"SctpAbortSmoothSwitch"`
	MaxUserNumPerCell string `xml:"MaxUserNumPerCell"`
	CsfbFlowOptSwitch string `xml:"CsfbFlowOptSwitch"`
	X2InitFailDelSwitch string `xml:"X2InitFailDelSwitch"`
	X2DynBlacklistAgingTimer string `xml:"X2DynBlacklistAgingTimer"`
	EX2AutoSetupSwitch string `xml:"eX2AutoSetupSwitch"`
	EX2DynBlackListSwitch string `xml:"eX2DynBlackListSwitch"`
	EX2DynBlackListAgingTimer string `xml:"eX2DynBlackListAgingTimer"`
	PRSMutingCtrlSwitch string `xml:"PRSMutingCtrlSwitch"`
	RrcConnReqStatSwitch string `xml:"RrcConnReqStatSwitch"`
	QciParaEffectFlag string `xml:"QciParaEffectFlag"`
	RimFirstMultiReportSwitch string `xml:"RimFirstMultiReportSwitch"`
	UuMsgTimeOutRelCause string `xml:"UuMsgTimeOutRelCause"`
	SigPoolOptPolicy string `xml:"SigPoolOptPolicy"`
	EnhancedRRCReestProtectThd string `xml:"EnhancedRRCReestProtectThd"`
	ObjId string `xml:"objId"`
	S1FaultSelfRecoverySwitch string `xml:"S1FaultSelfRecoverySwitch"`
	ItfTypeForNonIdealModeServ string `xml:"ItfTypeForNonIdealModeServ"`
	S1DefaultPagingDrxForNb string `xml:"S1DefaultPagingDrxForNb"`
	X2SonSecMode string `xml:"X2SonSecMode"`
	ProcTypeForNonIdealServData string `xml:"ProcTypeForNonIdealServData"`
	Ex2DeleteTimerForEx2Usage string `xml:"Ex2DeleteTimerForEx2Usage"`
	SinglePlmnHostSplitSwitch string `xml:"SinglePlmnHostSplitSwitch"`
	InitMsgProtectThld string `xml:"InitMsgProtectThld"`
}

