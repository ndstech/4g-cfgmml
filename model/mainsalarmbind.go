package model

import "encoding/xml"

type Mainsalarmbind struct {
	XMLName xml.Name `xml:"MAINSALARMBIND"`
	ATTRIBUTES MainsalarmbindAttributes `xml:"attributes"`
}

type MainsalarmbindAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ISDSWITCH string `xml:"ISDSWITCH"`
	NMSACN string `xml:"NMSACN"`
	NMSASRN string `xml:"NMSASRN"`
	NMSASN string `xml:"NMSASN"`
	NMSAPN string `xml:"NMSAPN"`
}

