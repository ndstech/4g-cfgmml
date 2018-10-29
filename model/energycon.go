package model

import "encoding/xml"

type Energycon struct {
	XMLName xml.Name `xml:"ENERGYCON"`
	ATTRIBUTES EnergyconAttributes `xml:"attributes"`
}

type EnergyconAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	MP string `xml:"MP"`
}

