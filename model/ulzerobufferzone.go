package model

import "encoding/xml"

type Ulzerobufferzone struct {
	XMLName xml.Name `xml:"ULZeroBufferZone"`
	ATTRIBUTES UlzerobufferzoneAttributes `xml:"attributes"`
}

type UlzerobufferzoneAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ULSharedFreqStartRB string `xml:"ULSharedFreqStartRB"`
	ULSharedFreqEndRB string `xml:"ULSharedFreqEndRB"`
	ULZeroBufZonePRBThd string `xml:"ULZeroBufZonePRBThd"`
	ULZeroBufZonePRBOffset string `xml:"ULZeroBufZonePRBOffset"`
	ULZeroBufZoneRsrpThd string `xml:"ULZeroBufZoneRsrpThd"`
	ULZeroBufZoneRsrpOffset string `xml:"ULZeroBufZoneRsrpOffset"`
	ULZeroBufZoneB1RmvOffset string `xml:"ULZeroBufZoneB1RmvOffset"`
	ULZeroBufZoneUtranRscpThd string `xml:"ULZeroBufZoneUtranRscpThd"`
	ULZeroBufZoInterRatB1Timer string `xml:"ULZeroBufZoInterRatB1Timer"`
}

