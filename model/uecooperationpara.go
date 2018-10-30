package model

import "encoding/xml"

type Uecooperationpara struct {
	XMLName xml.Name `xml:"UeCooperationPara"`
	ATTRIBUTES UecooperationparaAttributes `xml:"attributes"`
}

type UecooperationparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MacCeMsgLCID string `xml:"MacCeMsgLCID"`
	A2Offset string `xml:"A2Offset"`
	A3Offset string `xml:"A3Offset"`
	UEAwarePowerSavingSwitch string `xml:"UEAwarePowerSavingSwitch"`
	UEForbiddenMsgThd string `xml:"UEForbiddenMsgThd"`
	SpecUserCooperationSwitch string `xml:"SpecUserCooperationSwitch"`
	DsdsCooperationRptSwitch string `xml:"DsdsCooperationRptSwitch"`
	DsdsLcid string `xml:"DsdsLcid"`
}

