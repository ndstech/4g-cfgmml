package model

import "encoding/xml"

type Tpealgo struct {
	XMLName xml.Name `xml:"TpeAlgo"`
	ATTRIBUTES TpealgoAttributes `xml:"attributes"`
}

type TpealgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PortListNum string `xml:"PortListNum"`
	Port1 string `xml:"Port1"`
	Port2 string `xml:"Port2"`
	Port3 string `xml:"Port3"`
	Port4 string `xml:"Port4"`
	Port5 string `xml:"Port5"`
	Port6 string `xml:"Port6"`
	Port7 string `xml:"Port7"`
	Port8 string `xml:"Port8"`
	Port9 string `xml:"Port9"`
	Port10 string `xml:"Port10"`
	Port11 string `xml:"Port11"`
	Port12 string `xml:"Port12"`
	Port13 string `xml:"Port13"`
	Port14 string `xml:"Port14"`
	Port15 string `xml:"Port15"`
	Port16 string `xml:"Port16"`
	Port17 string `xml:"Port17"`
	Port18 string `xml:"Port18"`
	Port19 string `xml:"Port19"`
	Port20 string `xml:"Port20"`
	ObjId string `xml:"objId"`
}

