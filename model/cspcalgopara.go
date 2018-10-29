package model

import "encoding/xml"

type Cspcalgopara struct {
	XMLName xml.Name `xml:"CSPCALGOPARA"`
	ATTRIBUTES CspcalgoparaAttributes `xml:"attributes"`
}

type CspcalgoparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CspcAlgoSwitch string `xml:"CspcAlgoSwitch"`
	CspcRsrpMeasMode string `xml:"CspcRsrpMeasMode"`
	CspcClusterMode string `xml:"CspcClusterMode"`
	CspcClusterPartPeriod string `xml:"CspcClusterPartPeriod"`
	CspcComputeSwitch string `xml:"CspcComputeSwitch"`
	CspcFullPowerSubframeRatio string `xml:"CspcFullPowerSubframeRatio"`
	CspcPowerConfigDelay string `xml:"CspcPowerConfigDelay"`
	CspcScheduleUeSpec string `xml:"CspcScheduleUeSpec"`
	CspcEnableDlPrbRatioThd string `xml:"CspcEnableDlPrbRatioThd"`
	TddCspcAlgoSwitch string `xml:"TddCspcAlgoSwitch"`
	CspcCapacityFactor string `xml:"CspcCapacityFactor"`
	ObjId string `xml:"objId"`
}

