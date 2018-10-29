package model

import "encoding/xml"

type Glocellalgpara struct {
	XMLName xml.Name `xml:"GLOCELLALGPARA"`
	ATTRIBUTES GlocellalgparaAttributes `xml:"attributes"`
}

type GlocellalgparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GLOCELLID string `xml:"GLOCELLID"`
	GUDEGRATEPWRCTRL string `xml:"GUDEGRATEPWRCTRL"`
	GMSKDELAY string `xml:"GMSKDELAY"`
	DIVERT8PSKDELAY string `xml:"DIVERT8PSKDELAY"`
	DIVERT16QAMDELAY string `xml:"DIVERT16QAMDELAY"`
	DIVERT32QAMDELAY string `xml:"DIVERT32QAMDELAY"`
	GMSKDELAYDYNADJSW string `xml:"GMSKDELAYDYNADJSW"`
}

