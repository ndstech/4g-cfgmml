package model

import "encoding/xml"

type Cellusparacfg struct {
	XMLName xml.Name `xml:"CELLUSPARACFG"`
	ATTRIBUTES CellusparacfgAttributes `xml:"attributes"`
}

type CellusparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UsAlgoSwitch string `xml:"UsAlgoSwitch"`
	UsVoIPPreAllocMinPeriod string `xml:"UsVoIPPreAllocMinPeriod"`
	UsDataPreAllocMinPeriod string `xml:"UsDataPreAllocMinPeriod"`
	UsVoIPPreAllocSize string `xml:"UsVoIPPreAllocSize"`
	UsDataPreAllocSize string `xml:"UsDataPreAllocSize"`
	UsVoIPSmartPreallocDura string `xml:"UsVoIPSmartPreallocDura"`
	UsDataSmartPreallocDura string `xml:"UsDataSmartPreallocDura"`
	UlUsVoIPIblerTarget string `xml:"UlUsVoIPIblerTarget"`
	NsNumThdForUsHo string `xml:"NsNumThdForUsHo"`
	UsPucchSinrTargetOffset string `xml:"UsPucchSinrTargetOffset"`
	UsPucchRsrpHighThdOffset string `xml:"UsPucchRsrpHighThdOffset"`
	UsDataRatFreqPriGroupId string `xml:"UsDataRatFreqPriGroupId"`
	UsGuaranteeDurTimer string `xml:"UsGuaranteeDurTimer"`
	UsGapMeasurementPeriod string `xml:"UsGapMeasurementPeriod"`
	UsHoA2ThdRsrp string `xml:"UsHoA2ThdRsrp"`
	UsHoA2ThdRsrq string `xml:"UsHoA2ThdRsrq"`
	UsVoIPRatFreqPriGroupId string `xml:"UsVoIPRatFreqPriGroupId"`
	NsRatFreqPriGroupId string `xml:"NsRatFreqPriGroupId"`
	UsPaPcOff string `xml:"UsPaPcOff"`
	UsPuschSinrHighThdOffset string `xml:"UsPuschSinrHighThdOffset"`
	NsRbRestrictRatio string `xml:"NsRbRestrictRatio"`
	UsDetectTimer string `xml:"UsDetectTimer"`
	UsDlMinGbr string `xml:"UsDlMinGbr"`
	UsUlMinGbr string `xml:"UsUlMinGbr"`
	UsNbInterfCtrl string `xml:"UsNbInterfCtrl"`
	UlUsVoipRsvStartRb string `xml:"UlUsVoipRsvStartRb"`
	UlUsVoipRsvEndRb string `xml:"UlUsVoipRsvEndRb"`
	SfCtrlPrbThd string `xml:"SfCtrlPrbThd"`
	UsVoIPKeepOnLen string `xml:"UsVoIPKeepOnLen"`
	NsIdleRatFreqPriGroupId string `xml:"NsIdleRatFreqPriGroupId"`
	UsDlPriorityRatio string `xml:"UsDlPriorityRatio"`
	UsUlPriorityRatio string `xml:"UsUlPriorityRatio"`
	UsA3HoA2ThdOffset string `xml:"UsA3HoA2ThdOffset"`
	UsA4A5HoA2ThdOffset string `xml:"UsA4A5HoA2ThdOffset"`
	UlUsDataRatFreqPriGroupId string `xml:"UlUsDataRatFreqPriGroupId"`
	UsNbUlSinrHighThdOffset string `xml:"UsNbUlSinrHighThdOffset"`
	UsNbUlRbRestrictRat string `xml:"UsNbUlRbRestrictRat"`
	UsA3HoA1ThdOffset string `xml:"UsA3HoA1ThdOffset"`
	UsA4A5HoA1ThdOffset string `xml:"UsA4A5HoA1ThdOffset"`
	UsDataPdcchSinrOffset string `xml:"UsDataPdcchSinrOffset"`
	UsVoipPdcchSinrOffset string `xml:"UsVoipPdcchSinrOffset"`
	UsVoipInitDlIblerTarget string `xml:"UsVoipInitDlIblerTarget"`
	UsVoipCompenSchPeriod string `xml:"UsVoipCompenSchPeriod"`
	CrsShutDownRsrpOffset string `xml:"CrsShutDownRsrpOffset"`
	CrsShutDownStrategy string `xml:"CrsShutDownStrategy"`
	VolteSilDelayHoRsrpThld string `xml:"VolteSilDelayHoRsrpThld"`
	UsCellSchOptSwitch string `xml:"UsCellSchOptSwitch"`
	MMPowerRatio string `xml:"MMPowerRatio"`
	NsDlFirstRetxRbRatio string `xml:"NsDlFirstRetxRbRatio"`
	NsMaxNumAllowSchPerTTI string `xml:"NsMaxNumAllowSchPerTTI"`
	NsNotAllowSchSubframe string `xml:"NsNotAllowSchSubframe"`
	SchGuaranteeUsNum string `xml:"SchGuaranteeUsNum"`
	UsWithGapMode string `xml:"UsWithGapMode"`
	LargePacketDlDataThld string `xml:"LargePacketDlDataThld"`
	UsNotSchRbgRsrpOffset string `xml:"UsNotSchRbgRsrpOffset"`
	UsNotSchRbgRsrpThld string `xml:"UsNotSchRbgRsrpThld"`
	UsAlgoExSwitch string `xml:"UsAlgoExSwitch"`
	UsSrsGuaranteeSwitch string `xml:"UsSrsGuaranteeSwitch"`
	UeMobilIdentifyStrategy string `xml:"UeMobilIdentifyStrategy"`
	UsSmoothHoCompThldOfs string `xml:"UsSmoothHoCompThldOfs"`
}

