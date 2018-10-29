package model

import "encoding/xml"

type Celldynacbaralgopara struct {
	XMLName xml.Name `xml:"CELLDYNACBARALGOPARA"`
	ATTRIBUTES CelldynacbaralgoparaAttributes `xml:"attributes"`
}

type CelldynacbaralgoparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DynAcBarStatPeriod string `xml:"DynAcBarStatPeriod"`
	DynAcBarTriggerThd string `xml:"DynAcBarTriggerThd"`
	DynAcBarCancelThd string `xml:"DynAcBarCancelThd"`
	SsacTriggerCondSatiPeriods string `xml:"SsacTriggerCondSatiPeriods"`
	SsacCancelCondSatiPeriods string `xml:"SsacCancelCondSatiPeriods"`
	DisasterReferenceInd string `xml:"DisasterReferenceInd"`
	DisasterDuration string `xml:"DisasterDuration"`
	MoTriggerCondSatiPeriods string `xml:"MoTriggerCondSatiPeriods"`
	MoCancelCondSatiPeriods string `xml:"MoCancelCondSatiPeriods"`
	MoFactorAdjStep string `xml:"MoFactorAdjStep"`
	MoFactorRetreatStep string `xml:"MoFactorRetreatStep"`
	SsacFactorAdjStep string `xml:"SsacFactorAdjStep"`
	SsacFactorRetreatStep string `xml:"SsacFactorRetreatStep"`
}

