package model

import "encoding/xml"

type Ulcsalgopara struct {
	XMLName xml.Name `xml:"ULCSALGOPARA"`
	ATTRIBUTES UlcsalgoparaAttributes `xml:"attributes"`
}

type UlcsalgoparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlCoPcInThd string `xml:"UlCoPcInThd"`
	UlCoPcRbkRsrpThd string `xml:"UlCoPcRbkRsrpThd"`
	UlCoPcUserNumThd string `xml:"UlCoPcUserNumThd"`
	UlCoResAllocBandMode string `xml:"UlCoResAllocBandMode"`
	UlCoResAllocRbLoadThld string `xml:"UlCoResAllocRbLoadThld"`
	UlCsA3Offset string `xml:"UlCsA3Offset"`
	UlCsSw string `xml:"UlCsSw"`
	UlCraInThld string `xml:"UlCraInThld"`
	UlCraUserNumThld string `xml:"UlCraUserNumThld"`
}

