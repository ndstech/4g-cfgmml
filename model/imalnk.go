package model

import "encoding/xml"

type Imalnk struct {
	XMLName xml.Name `xml:"IMALNK"`
	ATTRIBUTES ImalnkAttributes `xml:"attributes"`
}

type ImalnkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	IMALNKN string `xml:"IMALNKN"`
	IMAGRPN string `xml:"IMAGRPN"`
	SBT string `xml:"SBT"`
	IMAGRPSBT string `xml:"IMAGRPSBT"`
}

