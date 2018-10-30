package model

import "encoding/xml"

type Cellhoparacfg struct {
	XMLName xml.Name `xml:"CellHoParaCfg"`
	ATTRIBUTES CellhoparacfgAttributes `xml:"attributes"`
}

type CellhoparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BlindHoA1A2ThdRsrp string `xml:"BlindHoA1A2ThdRsrp"`
	BlindHoA1A2ThdRsrq string `xml:"BlindHoA1A2ThdRsrq"`
	HighSpeedThreshold string `xml:"HighSpeedThreshold"`
	HoModeSwitch string `xml:"HoModeSwitch"`
	SrvccHoOptSwitch string `xml:"SrvccHoOptSwitch"`
	UlBadQualHoMcsThd string `xml:"UlBadQualHoMcsThd"`
	UlBadQualHoIblerThd string `xml:"UlBadQualHoIblerThd"`
	SpeedEvaluatedPeriod string `xml:"SpeedEvaluatedPeriod"`
	SpeedEvaluatedNum string `xml:"SpeedEvaluatedNum"`
	L2UCsfbMRProMode string `xml:"L2UCsfbMRProMode"`
	CsfbMRWaitingTimer string `xml:"CsfbMRWaitingTimer"`
	CellHoAlgoSwitch string `xml:"CellHoAlgoSwitch"`
	SpeedEvaluatedNumForFastUe string `xml:"SpeedEvaluatedNumForFastUe"`
	HoUseInactiveTimerSwitch string `xml:"HoUseInactiveTimerSwitch"`
	HSCellSleepUserNum string `xml:"HSCellSleepUserNum"`
	LowSpeedUsersSelProTimer string `xml:"LowSpeedUsersSelProTimer"`
	HSCellHoInProtectTimer string `xml:"HSCellHoInProtectTimer"`
	FlashSrvccSwitch string `xml:"FlashSrvccSwitch"`
	UlPoorCoverPathLossThd string `xml:"UlPoorCoverPathLossThd"`
	UlPoorCoverSinrThd string `xml:"UlPoorCoverSinrThd"`
	HoMrDelayTimerQci1 string `xml:"HoMrDelayTimerQci1"`
	VoLTEQualityHoAlgoSwitch string `xml:"VoLTEQualityHoAlgoSwitch"`
	Qci1PlrThldForVolteQualHo string `xml:"Qci1PlrThldForVolteQualHo"`
	VolteQualPktLossPeriod string `xml:"VolteQualPktLossPeriod"`
	CovBasedIntraRatMeasTime string `xml:"CovBasedIntraRatMeasTime"`
	HighSpeedUeJudgeMode string `xml:"HighSpeedUeJudgeMode"`
	RingingMsgCheckSw string `xml:"RingingMsgCheckSw"`
	SrvccMrDelayTimer string `xml:"SrvccMrDelayTimer"`
	VolteQualIfHoQci1PlrThld string `xml:"VolteQualIfHoQci1PlrThld"`
	VolteQualRecoveryQci1PlrThld string `xml:"VolteQualRecoveryQci1PlrThld"`
	InterRatUuHoFailRetryTimes string `xml:"InterRatUuHoFailRetryTimes"`
}

