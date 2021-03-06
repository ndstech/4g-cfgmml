package model

import "encoding/xml"

type Cellalgoswitch struct {
	XMLName xml.Name `xml:"CellAlgoSwitch"`
	ATTRIBUTES CellalgoswitchAttributes `xml:"attributes"`
}

type CellalgoswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	VolteRedirectSwitch string `xml:"VolteRedirectSwitch"`
	RachAlgoSwitch string `xml:"RachAlgoSwitch"`
	SrsAlgoSwitch string `xml:"SrsAlgoSwitch"`
	PucchAlgoSwitch string `xml:"PucchAlgoSwitch"`
	AqmAlgoSwitch string `xml:"AqmAlgoSwitch"`
	CqiAdjAlgoSwitch string `xml:"CqiAdjAlgoSwitch"`
	DynAdjVoltSwitch string `xml:"DynAdjVoltSwitch"`
	RacAlgoSwitch string `xml:"RacAlgoSwitch"`
	MlbAlgoSwitch string `xml:"MlbAlgoSwitch"`
	DlPcAlgoSwitch string `xml:"DlPcAlgoSwitch"`
	UlPcAlgoSwitch string `xml:"UlPcAlgoSwitch"`
	BfAlgoSwitch string `xml:"BfAlgoSwitch"`
	DlSchSwitch string `xml:"DlSchSwitch"`
	UlSchSwitch string `xml:"UlSchSwitch"`
	RanShareModeSwitch string `xml:"RanShareModeSwitch"`
	FreqPriorityHoSwitch string `xml:"FreqPriorityHoSwitch"`
	MuBfAlgoSwitch string `xml:"MuBfAlgoSwitch"`
	DistBasedHoSwitch string `xml:"DistBasedHoSwitch"`
	AcBarAlgoSwitch string `xml:"AcBarAlgoSwitch"`
	SfnUlSchSwitch string `xml:"SfnUlSchSwitch"`
	SfnDlSchSwitch string `xml:"SfnDlSchSwitch"`
	IrcSwitch string `xml:"IrcSwitch"`
	DynDrxSwitch string `xml:"DynDrxSwitch"`
	HighMobiTrigIdleModeSwitch string `xml:"HighMobiTrigIdleModeSwitch"`
	AvoidInterfSwitch string `xml:"AvoidInterfSwitch"`
	GLPwrShare string `xml:"GLPwrShare"`
	EicicSwitch string `xml:"EicicSwitch"`
	PUSCHMaxRBPUCCHAdjSwitch string `xml:"PUSCHMaxRBPUCCHAdjSwitch"`
	DlCompSwitch string `xml:"DlCompSwitch"`
	PsicSwitch string `xml:"PsicSwitch"`
	MlbHoMode string `xml:"MlbHoMode"`
	UplinkCompSwitch string `xml:"UplinkCompSwitch"`
	AntCalibrationAlgoSwitch string `xml:"AntCalibrationAlgoSwitch"`
	DynSpectrumShareSwitch string `xml:"DynSpectrumShareSwitch"`
	SfnLoadBasedAdptSwitch string `xml:"SfnLoadBasedAdptSwitch"`
	PuschIrcAlgoSwitch string `xml:"PuschIrcAlgoSwitch"`
	ReselecPriAdaptSwitch string `xml:"ReselecPriAdaptSwitch"`
	AnrFunctionSwitch string `xml:"AnrFunctionSwitch"`
	SfnAlgoSwitch string `xml:"SfnAlgoSwitch"`
	PrachIntrfRejSwitch string `xml:"PrachIntrfRejSwitch"`
	EnhMIMOSwitch string `xml:"EnhMIMOSwitch"`
	InterfRandSwitch string `xml:"InterfRandSwitch"`
	RepeaterSwitch string `xml:"RepeaterSwitch"`
	MultiFreqPriControlSwitch string `xml:"MultiFreqPriControlSwitch"`
	HarqAlgoSwitch string `xml:"HarqAlgoSwitch"`
	CovBasedInterFreqHoMode string `xml:"CovBasedInterFreqHoMode"`
	LteUtcBroadcastSwitch string `xml:"LteUtcBroadcastSwitch"`
	CellSchStrategySwitch string `xml:"CellSchStrategySwitch"`
	SsrdAlgoSwitch string `xml:"SsrdAlgoSwitch"`
	SfnUplinkCompSwitch string `xml:"SfnUplinkCompSwitch"`
	LowSpeedInterFreqHoSwitch string `xml:"LowSpeedInterFreqHoSwitch"`
	RelaySwitch string `xml:"RelaySwitch"`
	InterFreqDirectHoSwitch string `xml:"InterFreqDirectHoSwitch"`
	PwrDeratSwitch string `xml:"PwrDeratSwitch"`
	DetectionAlgoSwitch string `xml:"DetectionAlgoSwitch"`
	PucchIRCEnhance string `xml:"PucchIRCEnhance"`
	AcBarAlgoforDynSwitch string `xml:"AcBarAlgoforDynSwitch"`
	CreSwitch string `xml:"CreSwitch"`
	BackoffAlgoSwitch string `xml:"BackoffAlgoSwitch"`
	HoAllowedSwitch string `xml:"HoAllowedSwitch"`
	NCellClassMgtSw string `xml:"NCellClassMgtSw"`
	SpePCIBasedPolicySw string `xml:"SpePCIBasedPolicySw"`
	CellPIMInterMitigSwitch string `xml:"CellPIMInterMitigSwitch"`
	PrachJointReceptionSwitch string `xml:"PrachJointReceptionSwitch"`
	FeicicSwitch string `xml:"FeicicSwitch"`
	CamcSwitch string `xml:"CamcSwitch"`
	RruUeMapSwitch string `xml:"RruUeMapSwitch"`
	HighSpeedSchOptSwitch string `xml:"HighSpeedSchOptSwitch"`
	ServiceDiffSwitch string `xml:"ServiceDiffSwitch"`
	LtePttQoSSwitch string `xml:"LtePttQoSSwitch"`
	SrsPucchEnhancedSwitch string `xml:"SrsPucchEnhancedSwitch"`
	UEInactiveTimerQCI1Switch string `xml:"UEInactiveTimerQCI1Switch"`
	UlJRAntNumCombSw string `xml:"UlJRAntNumCombSw"`
	VamPhaseShiftAlgoSwitch string `xml:"VamPhaseShiftAlgoSwitch"`
	AnrAlgoSwitch string `xml:"AnrAlgoSwitch"`
	Dl256QamAlgoSwitch string `xml:"Dl256QamAlgoSwitch"`
	MlbAutoGroupSwitch string `xml:"MlbAutoGroupSwitch"`
	CaAutoGroupSwitch string `xml:"CaAutoGroupSwitch"`
	AttachCellSelfCfgSwitch string `xml:"AttachCellSelfCfgSwitch"`
	CellDlCoverEnhanceSwitch string `xml:"CellDlCoverEnhanceSwitch"`
	UlSchExtSwitch string `xml:"UlSchExtSwitch"`
	UdcAlgoSwitch string `xml:"UdcAlgoSwitch"`
	VoLTESwitch string `xml:"VoLTESwitch"`
	VoipFailDecSelfRecSwitch string `xml:"VoipFailDecSelfRecSwitch"`
	DeprioritisationDeliverInd string `xml:"DeprioritisationDeliverInd"`
	CmasSwitch string `xml:"CmasSwitch"`
	MFBIAlgoSwitch string `xml:"MFBIAlgoSwitch"`
	PciConflictReportSwitch string `xml:"PciConflictReportSwitch"`
	MroSwitch string `xml:"MroSwitch"`
	OpResourceGroupShareSwitch string `xml:"OpResourceGroupShareSwitch"`
	CellRecoverySwitch string `xml:"CellRecoverySwitch"`
	EnhancedMlbAlgoSwitch string `xml:"EnhancedMlbAlgoSwitch"`
	TrafficMlbSwitch string `xml:"TrafficMlbSwitch"`
	MtcSwitch string `xml:"MtcSwitch"`
	UlIcSwitch string `xml:"UlIcSwitch"`
	FcsMode string `xml:"FcsMode"`
	ScVideoOptSwitch string `xml:"ScVideoOptSwitch"`
	SpidSpecificHoSwitch string `xml:"SpidSpecificHoSwitch"`
	DMIMOAlgoSwitch string `xml:"DMIMOAlgoSwitch"`
	UlAmrcMode string `xml:"UlAmrcMode"`
	AmrcAlgoSwitch string `xml:"AmrcAlgoSwitch"`
	CrsIcSwitch string `xml:"CrsIcSwitch"`
	MeasOptAlgoSwitch string `xml:"MeasOptAlgoSwitch"`
	FreqLayerSwitch string `xml:"FreqLayerSwitch"`
	EmimoSwitch string `xml:"EmimoSwitch"`
	RohcSwitch string `xml:"RohcSwitch"`
	McpttSwitch string `xml:"McpttSwitch"`
	UlIcsRbRatio string `xml:"UlIcsRbRatio"`
	UlIcsVoLTEPLTh string `xml:"UlIcsVoLTEPLTh"`
	DacqSwitch string `xml:"DacqSwitch"`
	RrcReestDataFwdSwitch string `xml:"RrcReestDataFwdSwitch"`
	MTCCongControlSwitch string `xml:"MTCCongControlSwitch"`
	MTCPowerSavSwitch string `xml:"MTCPowerSavSwitch"`
	TcpCtrlSwitch string `xml:"TcpCtrlSwitch"`
	TurboReceiverSwitch string `xml:"TurboReceiverSwitch"`
	NaicsSwitch string `xml:"NaicsSwitch"`
	UplinkIcSwitch string `xml:"UplinkIcSwitch"`
	DacqEnhancementSwitch string `xml:"DacqEnhancementSwitch"`
	AsAlgoSwitch string `xml:"AsAlgoSwitch"`
	EnhChnCalSwitch string `xml:"EnhChnCalSwitch"`
	NbCellAlgoSwitch string `xml:"NbCellAlgoSwitch"`
	MultiCnConnFreqPriSw string `xml:"MultiCnConnFreqPriSw"`
	NoSrsSccBfAlgoSwitch string `xml:"NoSrsSccBfAlgoSwitch"`
	SpecUserAlgoSwitch string `xml:"SpecUserAlgoSwitch"`
	MimoUePaAdaptSw string `xml:"MimoUePaAdaptSw"`
	UlSuMimoAlgoSwitch string `xml:"UlSuMimoAlgoSwitch"`
	HighSpeedInterRatHoSwitch string `xml:"HighSpeedInterRatHoSwitch"`
	EmcSpsSchSwitch string `xml:"EmcSpsSchSwitch"`
	LbsSwitch string `xml:"LbsSwitch"`
	DtxDetectionAlgoSwitch string `xml:"DtxDetectionAlgoSwitch"`
	DlSchExtSwitch string `xml:"DlSchExtSwitch"`
	MPMUDetectSwitch string `xml:"MPMUDetectSwitch"`
	VmsSwitch string `xml:"VmsSwitch"`
	SmallBandOptSwitch string `xml:"SmallBandOptSwitch"`
	IdleModeEdrxSwitch string `xml:"IdleModeEdrxSwitch"`
}

