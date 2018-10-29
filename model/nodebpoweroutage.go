package model

import "encoding/xml"

type Nodebpoweroutage struct {
	XMLName xml.Name `xml:"NODEBPOWEROUTAGE"`
	ATTRIBUTES NodebpoweroutageAttributes `xml:"attributes"`
}

type NodebpoweroutageAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ISDSW string `xml:"ISDSW"`
	BAKPWRSAVPOLICY string `xml:"BAKPWRSAVPOLICY"`
	PHASE1PERIOD string `xml:"PHASE1PERIOD"`
	PHASE2PERIOD string `xml:"PHASE2PERIOD"`
	OBJID string `xml:"OBJID"`
}

