package model

import "encoding/xml"

type Eucoschdlcompcfg struct {
	XMLName xml.Name `xml:"EUCOSCHDLCOMPCFG"`
	ATTRIBUTES EucoschdlcompcfgAttributes `xml:"attributes"`
}

type EucoschdlcompcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CordInfoEffDelay string `xml:"CordInfoEffDelay"`
	InterEnbDlCompSwitch string `xml:"InterEnbDlCompSwitch"`
	EuCoSchUeSpec string `xml:"EuCoSchUeSpec"`
	CordInfoEffDelayMode string `xml:"CordInfoEffDelayMode"`
	ObjId string `xml:"objId"`
}

