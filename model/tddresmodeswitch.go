package model

import "encoding/xml"

type Tddresmodeswitch struct {
	XMLName xml.Name `xml:"TddResModeSwitch"`
	ATTRIBUTES TddresmodeswitchAttributes `xml:"attributes"`
}

type TddresmodeswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ClkResExcludeSwitch string `xml:"ClkResExcludeSwitch"`
	BbResExclusiveSwitch string `xml:"BbResExclusiveSwitch"`
	ObjId string `xml:"objId"`
}

