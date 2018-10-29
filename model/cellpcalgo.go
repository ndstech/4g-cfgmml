package model

import "encoding/xml"

type Cellpcalgo struct {
	XMLName xml.Name `xml:"CELLPCALGO"`
	ATTRIBUTES CellpcalgoAttributes `xml:"attributes"`
}

type CellpcalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SrsPcStrategy string `xml:"SrsPcStrategy"`
	PucchCloseLoopPcType string `xml:"PucchCloseLoopPcType"`
	PucchPcPeriod string `xml:"PucchPcPeriod"`
	PucchPcTargetSinrOffset string `xml:"PucchPcTargetSinrOffset"`
	PuschRsrpHighThd string `xml:"PuschRsrpHighThd"`
	PuschIoTCtrlA3Offset string `xml:"PuschIoTCtrlA3Offset"`
	IoTNearPointOptSwitch string `xml:"IoTNearPointOptSwitch"`
	IoTNearPointPLThreshold string `xml:"IoTNearPointPLThreshold"`
	SrsPcSinrTarget string `xml:"SrsPcSinrTarget"`
	SrsPcRsrpTarget string `xml:"SrsPcRsrpTarget"`
	PUSCHPsdCtrlTarget string `xml:"PUSCHPsdCtrlTarget"`
	CloseLoopOptPUSCHType string `xml:"CloseLoopOptPUSCHType"`
	PUSCHOptIBLERThreshold string `xml:"PUSCHOptIBLERThreshold"`
	PUSCHPsdCtrlTargetForUs string `xml:"PUSCHPsdCtrlTargetForUs"`
	IoTCtrlINCorrectSwitch string `xml:"IoTCtrlINCorrectSwitch"`
	IoTCtrlEUPLThreshold string `xml:"IoTCtrlEUPLThreshold"`
	IoTCtrlNIThreshold string `xml:"IoTCtrlNIThreshold"`
	NearPointUePuschType string `xml:"NearPointUePuschType"`
	VoLteSinrHighThdOffset string `xml:"VoLteSinrHighThdOffset"`
	VoltePUSCHPowerOffset string `xml:"VoltePUSCHPowerOffset"`
	PuschRsrpHighThdOffsetVoIP string `xml:"PuschRsrpHighThdOffsetVoIP"`
	DMSrsPcSinrOffset string `xml:"DMSrsPcSinrOffset"`
}

