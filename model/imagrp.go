package model

import "encoding/xml"

type Imagrp struct {
	XMLName xml.Name `xml:"IMAGRP"`
	ATTRIBUTES ImagrpAttributes `xml:"attributes"`
}

type ImagrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	IMAGRPN string `xml:"IMAGRPN"`
	VER string `xml:"VER"`
	CLKM string `xml:"CLKM"`
	FRMLEN string `xml:"FRMLEN"`
	MINLNK string `xml:"MINLNK"`
	DELAY string `xml:"DELAY"`
	SCRAM string `xml:"SCRAM"`
	TS16 string `xml:"TS16"`
	SBT string `xml:"SBT"`
}

