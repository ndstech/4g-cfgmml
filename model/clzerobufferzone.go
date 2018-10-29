package model

import "encoding/xml"

type Clzerobufferzone struct {
	XMLName xml.Name `xml:"CLZEROBUFFERZONE"`
	ATTRIBUTES ClzerobufferzoneAttributes `xml:"attributes"`
}

type ClzerobufferzoneAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ClZeroBufzoneDlSharedInd string `xml:"ClZeroBufzoneDlSharedInd"`
	ClZeroBufzoneUlSharedInd string `xml:"ClZeroBufzoneUlSharedInd"`
	ClSharedFreqStartRb1 string `xml:"ClSharedFreqStartRb1"`
	ClSharedFreqEndRb1 string `xml:"ClSharedFreqEndRb1"`
	ClSharedFreqStartRb2 string `xml:"ClSharedFreqStartRb2"`
	ClSharedFreqEndRb2 string `xml:"ClSharedFreqEndRb2"`
	ClZeroBufferZoneUlPrbThd string `xml:"ClZeroBufferZoneUlPrbThd"`
	ClZeroBufferZoneUlPrbOst string `xml:"ClZeroBufferZoneUlPrbOst"`
	UlNearPtUserRsrpThd string `xml:"UlNearPtUserRsrpThd"`
	DlNearPtUserRsrpThd string `xml:"DlNearPtUserRsrpThd"`
}

