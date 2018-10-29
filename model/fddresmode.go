package model

import "encoding/xml"

type Fddresmode struct {
	XMLName xml.Name `xml:"FDDRESMODE"`
	ATTRIBUTES FddresmodeAttributes `xml:"attributes"`
}

type FddresmodeAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BbCapabilityMode string `xml:"BbCapabilityMode"`
	SfnCapabilityMode string `xml:"SfnCapabilityMode"`
	BbResExclusiveSwitch string `xml:"BbResExclusiveSwitch"`
	ObjId string `xml:"objId"`
}

