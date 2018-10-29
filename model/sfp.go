package model

import "encoding/xml"

type Sfp struct {
	XMLName xml.Name `xml:"SFP"`
	ATTRIBUTES SfpAttributes `xml:"attributes"`
}

type SfpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	MODULEID string `xml:"MODULEID"`
	PT string `xml:"PT"`
}

