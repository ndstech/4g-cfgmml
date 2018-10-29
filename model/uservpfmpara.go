package model

import "encoding/xml"

type Uservpfmpara struct {
	XMLName xml.Name `xml:"USERVPFMPARA"`
	ATTRIBUTES UservpfmparaAttributes `xml:"attributes"`
}

type UservpfmparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlUserThruputThd0 string `xml:"DlUserThruputThd0"`
	DlUserThruputThd1 string `xml:"DlUserThruputThd1"`
	DlUserThruputThd2 string `xml:"DlUserThruputThd2"`
	DlUserThruputThd3 string `xml:"DlUserThruputThd3"`
	DlUserThruputThd4 string `xml:"DlUserThruputThd4"`
	UlAccessDelayGoodThd string `xml:"UlAccessDelayGoodThd"`
	UlAccessDelayBadThd string `xml:"UlAccessDelayBadThd"`
	DlAccessDelayGoodThd string `xml:"DlAccessDelayGoodThd"`
	DlAccessDelayBadThd string `xml:"DlAccessDelayBadThd"`
}

