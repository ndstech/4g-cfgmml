package model

import "encoding/xml"

type Cellulschalgo struct {
	XMLName xml.Name `xml:"CellUlschAlgo"`
	ATTRIBUTES CellulschalgoAttributes `xml:"attributes"`
}

type CellulschalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlschStrategy string `xml:"UlschStrategy"`
	AdaptHarqSwitch string `xml:"AdaptHarqSwitch"`
	SinrAdjustTargetIbler string `xml:"SinrAdjustTargetIbler"`
	PreAllocationBandwidthRatio string `xml:"PreAllocationBandwidthRatio"`
	PreAllocationMinPeriod string `xml:"PreAllocationMinPeriod"`
	PreAllocationSize string `xml:"PreAllocationSize"`
	UlHoppingType string `xml:"UlHoppingType"`
	FreeUserUlRbUsedRatio string `xml:"FreeUserUlRbUsedRatio"`
	UlSrSchDateLen string `xml:"UlSrSchDateLen"`
	SpsRelThd string `xml:"SpsRelThd"`
	SmartPreAllocationDuration string `xml:"SmartPreAllocationDuration"`
	UlEpfCapacityFactor string `xml:"UlEpfCapacityFactor"`
	UlRbAllocationStrategy string `xml:"UlRbAllocationStrategy"`
	DopMeasLevel string `xml:"DopMeasLevel"`
	UlHarqMaxTxNum string `xml:"UlHarqMaxTxNum"`
	SmartPreAllocDuraForSparse string `xml:"SmartPreAllocDuraForSparse"`
	UlSpsInterval string `xml:"UlSpsInterval"`
	SriFalseDetThdSwitch string `xml:"SriFalseDetThdSwitch"`
	NoUlSchTtiNumAffterGap string `xml:"NoUlSchTtiNumAffterGap"`
	UlDelaySchStrategy string `xml:"UlDelaySchStrategy"`
	UlSchAbnUeThd string `xml:"UlSchAbnUeThd"`
	SpecificPktSizeThd string `xml:"SpecificPktSizeThd"`
	SrMaskSwitch string `xml:"SrMaskSwitch"`
	PuschDtxSchStrategy string `xml:"PuschDtxSchStrategy"`
	UlVoipRlcMaxSegNum string `xml:"UlVoipRlcMaxSegNum"`
	UlEnhencedVoipSchSw string `xml:"UlEnhencedVoipSchSw"`
	UlInBasedFssSinrThld string `xml:"UlInBasedFssSinrThld"`
	UlCamcDlRsrpOffset string `xml:"UlCamcDlRsrpOffset"`
	StatisticNumThdForTtibTrig string `xml:"StatisticNumThdForTtibTrig"`
	StatisticNumThdForTtibExit string `xml:"StatisticNumThdForTtibExit"`
	HystToExitTtiBundling string `xml:"HystToExitTtiBundling"`
	TtiBundlingRlcMaxSegNum string `xml:"TtiBundlingRlcMaxSegNum"`
	TtiBundlingHarqMaxTxNum string `xml:"TtiBundlingHarqMaxTxNum"`
	TtiBundlingTriggerStrategy string `xml:"TtiBundlingTriggerStrategy"`
	DopAlgoSwitch string `xml:"DopAlgoSwitch"`
	EnhancedVmimoSwitch string `xml:"EnhancedVmimoSwitch"`
	UeNumThdInPdcchPuschBal string `xml:"UeNumThdInPdcchPuschBal"`
	DataThdInPdcchPuschBal string `xml:"DataThdInPdcchPuschBal"`
	HeadOverheadForUlSch string `xml:"HeadOverheadForUlSch"`
	PreAllocMinPeriodForSparse string `xml:"PreAllocMinPeriodForSparse"`
	PreallocationSizeForSparse string `xml:"PreallocationSizeForSparse"`
	UlInterfRandomMode string `xml:"UlInterfRandomMode"`
	SinrAdjTargetIblerforVoLTE string `xml:"SinrAdjTargetIblerforVoLTE"`
	SfnUlLoadPeriod string `xml:"SfnUlLoadPeriod"`
	MaxLayerHOVMIMO string `xml:"MaxLayerHOVMIMO"`
	UlCompenSchPeriodinSpurt string `xml:"UlCompenSchPeriodinSpurt"`
	UlCompenSchPeriodinSilence string `xml:"UlCompenSchPeriodinSilence"`
	UlTargetIBlerAdaptType string `xml:"UlTargetIBlerAdaptType"`
	AperiodicCsiUlTxMode string `xml:"AperiodicCsiUlTxMode"`
	UlVoipRsvRbStart string `xml:"UlVoipRsvRbStart"`
	UlVoipRsvRbNum string `xml:"UlVoipRsvRbNum"`
	UlIBlerAdaptBigTrafficSw string `xml:"UlIBlerAdaptBigTrafficSw"`
	VmimoOptAlgoSwitch string `xml:"VmimoOptAlgoSwitch"`
	UlCamcInterfTh string `xml:"UlCamcInterfTh"`
	UlIcsA3Offset string `xml:"UlIcsA3Offset"`
	UlMcsLowThreshold string `xml:"UlMcsLowThreshold"`
	UserSpeedUlSchPriWeight string `xml:"UserSpeedUlSchPriWeight"`
	SinrAdjTargetIblerforPtt string `xml:"SinrAdjTargetIblerforPtt"`
	TarRruSelRsrpOffsetThd string `xml:"TarRruSelRsrpOffsetThd"`
	MaxUlSchRbNum string `xml:"MaxUlSchRbNum"`
	UlExtVolteSchSw string `xml:"UlExtVolteSchSw"`
	UlVolteDeltaSinrForNack string `xml:"UlVolteDeltaSinrForNack"`
	InitUlSinrAdjust string `xml:"InitUlSinrAdjust"`
	UlSinrAdjustStep string `xml:"UlSinrAdjustStep"`
	UlSinrFilterCoef string `xml:"UlSinrFilterCoef"`
	SfnUlPairRsrpThd string `xml:"SfnUlPairRsrpThd"`
	UlSrSchDataVolAdptOptUpThd string `xml:"UlSrSchDataVolAdptOptUpThd"`
	TtiBundlingRetxStrategy string `xml:"TtiBundlingRetxStrategy"`
	UlVoLTERetransSchStrategy string `xml:"UlVoLTERetransSchStrategy"`
	NbUlHarqMaxTxCount string `xml:"NbUlHarqMaxTxCount"`
	PrealloSystemBwCoeff string `xml:"PrealloSystemBwCoeff"`
	SmartPrealloDurationCoeff string `xml:"SmartPrealloDurationCoeff"`
	HighSpeedSdmaIsolationThld string `xml:"HighSpeedSdmaIsolationThld"`
	VmimoPairingStrategy string `xml:"VmimoPairingStrategy"`
	VMIMOEgdeResRatio string `xml:"VMIMOEgdeResRatio"`
	MaxLayerMMVMIMO string `xml:"MaxLayerMMVMIMO"`
	UlSrSchDateLenForVoLTE string `xml:"UlSrSchDateLenForVoLTE"`
	AmrcDecreasingPeriod string `xml:"AmrcDecreasingPeriod"`
	SinrThdForVoLTERateCtrl string `xml:"SinrThdForVoLTERateCtrl"`
	RateCtrlCmrProcessStrategy string `xml:"RateCtrlCmrProcessStrategy"`
	SinrAdjTargetIblerforQCI2 string `xml:"SinrAdjTargetIblerforQCI2"`
}

