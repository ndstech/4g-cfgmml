package model

import "encoding/xml"

type Cellpdcchalgo struct {
	XMLName xml.Name `xml:"CellPdcchAlgo"`
	ATTRIBUTES CellpdcchalgoAttributes `xml:"attributes"`
}

type CellpdcchalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ComSigCongregLv string `xml:"ComSigCongregLv"`
	CceRatioAdjSwitch string `xml:"CceRatioAdjSwitch"`
	SfnPdcchDcsThd string `xml:"SfnPdcchDcsThd"`
	InitPdcchSymNum string `xml:"InitPdcchSymNum"`
	VirtualLoadPro string `xml:"VirtualLoadPro"`
	PdcchInitialCceAdjustValue string `xml:"PdcchInitialCceAdjustValue"`
	PdcchSymNumSwitch string `xml:"PdcchSymNumSwitch"`
	CceUseRatio string `xml:"CceUseRatio"`
	PdcchAggLvlCLAdjustSwitch string `xml:"PdcchAggLvlCLAdjustSwitch"`
	DPDVirtualLoadSwitch string `xml:"DPDVirtualLoadSwitch"`
	DPDVirtualLoadType string `xml:"DPDVirtualLoadType"`
	AggLvlSelStrageForDualCW string `xml:"AggLvlSelStrageForDualCW"`
	PdcchCapacityImproveSwitch string `xml:"PdcchCapacityImproveSwitch"`
	PdcchMaxCodeRate string `xml:"PdcchMaxCodeRate"`
	ULDLPdcchSymNum string `xml:"ULDLPdcchSymNum"`
	PDCCHAggLvlAdaptStrage string `xml:"PDCCHAggLvlAdaptStrage"`
	HysForCfiBasedPreSch string `xml:"HysForCfiBasedPreSch"`
	SfnPdcchSdmaThd string `xml:"SfnPdcchSdmaThd"`
	UlPdcchAllocImproveSwitch string `xml:"UlPdcchAllocImproveSwitch"`
	CceMaxInitialRatio string `xml:"CceMaxInitialRatio"`
	PdcchPowerEnhancedSwitch string `xml:"PdcchPowerEnhancedSwitch"`
	PdcchBlerTarget string `xml:"PdcchBlerTarget"`
	HLNetAccSigAggLvlSelEnhSw string `xml:"HLNetAccSigAggLvlSelEnhSw"`
	EpdcchAlgoSwitch string `xml:"EpdcchAlgoSwitch"`
	CceThdtoEnableEpdcch string `xml:"CceThdtoEnableEpdcch"`
	RbThdtoEnableEpdcch string `xml:"RbThdtoEnableEpdcch"`
	CceThdtoDisableEpdcch string `xml:"CceThdtoDisableEpdcch"`
	EcceThdtoDisableEpdcch string `xml:"EcceThdtoDisableEpdcch"`
	RbThdtoDisableEpdcch string `xml:"RbThdtoDisableEpdcch"`
	HLNetAccSigAttempt string `xml:"HLNetAccSigAttempt"`
	EpdcchSfCfg string `xml:"EpdcchSfCfg"`
	UlDlEcceInitialRatioAdjSw string `xml:"UlDlEcceInitialRatioAdjSw"`
	CapacityBfOpt string `xml:"CapacityBfOpt"`
	CCEThdEnableFlowCtrl string `xml:"CCEThdEnableFlowCtrl"`
	CCEThdDisableFlowCtrl string `xml:"CCEThdDisableFlowCtrl"`
	PDCCHPwrUpperLimitOptSw string `xml:"PDCCHPwrUpperLimitOptSw"`
	VoltePdcchSinrOffset string `xml:"VoltePdcchSinrOffset"`
	PdcchSparePowerAllocStrage string `xml:"PdcchSparePowerAllocStrage"`
	PdcchDlAggLvlSlcEhnSwitch string `xml:"PdcchDlAggLvlSlcEhnSwitch"`
	PdcchOutLoopAdjBaseStep string `xml:"PdcchOutLoopAdjBaseStep"`
	PdcchOutLoopAdjLowerLimit string `xml:"PdcchOutLoopAdjLowerLimit"`
	PdcchAdjAlgoSwitch string `xml:"PdcchAdjAlgoSwitch"`
	SplitBeamPdcchSdmaThd string `xml:"SplitBeamPdcchSdmaThd"`
	NackDtxRatioThd string `xml:"NackDtxRatioThd"`
	SplitBeamPdcchSdmaSw string `xml:"SplitBeamPdcchSdmaSw"`
	PdcchPwrCtrlUserNumThd string `xml:"PdcchPwrCtrlUserNumThd"`
	PdcchBfGainOffset string `xml:"PdcchBfGainOffset"`
}

