package model

import "encoding/xml"

type Plrthreshold struct {
	XMLName xml.Name `xml:"PLRTHRESHOLD"`
	ATTRIBUTES PlrthresholdAttributes `xml:"attributes"`
}

type PlrthresholdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PLRAT string `xml:"PLRAT"`
	PLRDT string `xml:"PLRDT"`
}

