package model

import "encoding/xml"

type Cpriport struct {
	XMLName xml.Name `xml:"CPRIPORT"`
	ATTRIBUTES CpriportAttributes `xml:"attributes"`
}

type CpriportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	OPTN string `xml:"OPTN"`
	ADMINISTRATIVESTATE string `xml:"ADMINISTRATIVESTATE"`
	PT string `xml:"PT"`
	SPN string `xml:"SPN"`
}

