package model

import "encoding/xml"

type Cellulunifiedolc struct {
	XMLName xml.Name `xml:"CellULUnifiedOLC"`
	ATTRIBUTES CellulunifiedolcAttributes `xml:"attributes"`
}

type CellulunifiedolcAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	RrcRejectRateHighThd string `xml:"RrcRejectRateHighThd"`
	RrcRejectRateLowThd string `xml:"RrcRejectRateLowThd"`
	RrcReqNumHighThd string `xml:"RrcReqNumHighThd"`
	RrcReqNumLowThd string `xml:"RrcReqNumLowThd"`
	UeNumHighThd string `xml:"UeNumHighThd"`
	UeNumLowThd string `xml:"UeNumLowThd"`
}

