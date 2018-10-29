package model

import "encoding/xml"

type Celldrxpara struct {
	XMLName xml.Name `xml:"CELLDRXPARA"`
	ATTRIBUTES CelldrxparaAttributes `xml:"attributes"`
}

type CelldrxparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	FddEnterDrxThd string `xml:"FddEnterDrxThd"`
	FddExitDrxThd string `xml:"FddExitDrxThd"`
	TddEnterDrxThdUl string `xml:"TddEnterDrxThdUl"`
	TddExitDrxThdUl string `xml:"TddExitDrxThdUl"`
	TddEnterDrxThdDl string `xml:"TddEnterDrxThdDl"`
	TddExitDrxThdDl string `xml:"TddExitDrxThdDl"`
	DataAmountStatTimer string `xml:"DataAmountStatTimer"`
	TddPowerSaveMeasSwitch string `xml:"TddPowerSaveMeasSwitch"`
	TddPowerSaveStatTimer string `xml:"TddPowerSaveStatTimer"`
	TddPowerSavingEnterDrxThd string `xml:"TddPowerSavingEnterDrxThd"`
	TddPowerSavingExitDrxThd string `xml:"TddPowerSavingExitDrxThd"`
	LongDrxCycleUnsync string `xml:"LongDrxCycleUnsync"`
	CqiMask string `xml:"CqiMask"`
	OndurationTimerUnsync string `xml:"OndurationTimerUnsync"`
	DrxInactivityTimerUnsync string `xml:"DrxInactivityTimerUnsync"`
	DrxPolicyMode string `xml:"DrxPolicyMode"`
	DrxStartOffsetOptSwitch string `xml:"DrxStartOffsetOptSwitch"`
	DrxRcvDtxProSwitch string `xml:"DrxRcvDtxProSwitch"`
	DrxForMeasSwitch string `xml:"DrxForMeasSwitch"`
	LongDrxCycleForMeas string `xml:"LongDrxCycleForMeas"`
	OnDurTimerForMeas string `xml:"OnDurTimerForMeas"`
	DrxInactTimerForMeas string `xml:"DrxInactTimerForMeas"`
	DrxReTxTimerForMeas string `xml:"DrxReTxTimerForMeas"`
	ShortDrxSwForMeas string `xml:"ShortDrxSwForMeas"`
	ShortDrxCycleForMeas string `xml:"ShortDrxCycleForMeas"`
	ShortCycleTimerForMeas string `xml:"ShortCycleTimerForMeas"`
	DrxSrDetectOptSwitch string `xml:"DrxSrDetectOptSwitch"`
	DrxStartoffsetAdjustSW string `xml:"DrxStartoffsetAdjustSW"`
	MeasDrxSpecSchSwitch string `xml:"MeasDrxSpecSchSwitch"`
	CovGsmMeasDrxCfgSwitch string `xml:"CovGsmMeasDrxCfgSwitch"`
	OnDurationUnextendSwitch string `xml:"OnDurationUnextendSwitch"`
	DrxStateDuringUlHarqRetx string `xml:"DrxStateDuringUlHarqRetx"`
	DrxAlgSwitch string `xml:"DrxAlgSwitch"`
	ShortDrxCycleSwitch string `xml:"ShortDrxCycleSwitch"`
	VolteGapDrxExclusiveSwitch string `xml:"VolteGapDrxExclusiveSwitch"`
	DrxStopSrPendingSw string `xml:"DrxStopSrPendingSw"`
	GapDrxExclusiveSwitch string `xml:"GapDrxExclusiveSwitch"`
	NbDrxInactivityTimer string `xml:"NbDrxInactivityTimer"`
	NbDrxReTxTimer string `xml:"NbDrxReTxTimer"`
	NbDrxUlReTxTimer string `xml:"NbDrxUlReTxTimer"`
	NbLongDrxCycle string `xml:"NbLongDrxCycle"`
	NbOnDurationTimer string `xml:"NbOnDurationTimer"`
	DrxOdPreSchSwitch string `xml:"DrxOdPreSchSwitch"`
	DrxUeSrsOptSwitch string `xml:"DrxUeSrsOptSwitch"`
	SinrThldForVolteDrxCtrl string `xml:"SinrThldForVolteDrxCtrl"`
	VoltePlrThldForExitingDrx string `xml:"VoltePlrThldForExitingDrx"`
}

