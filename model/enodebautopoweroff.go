package model

import "encoding/xml"

type Enodebautopoweroff struct {
	XMLName xml.Name `xml:"eNodeBAutoPowerOff"`
	ATTRIBUTES EnodebautopoweroffAttributes `xml:"attributes"`
}

type EnodebautopoweroffAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AutoPowerOffSwitch string `xml:"AutoPowerOffSwitch"`
	PowerOffTime string `xml:"PowerOffTime"`
	PowerOnTime string `xml:"PowerOnTime"`
	ObjId string `xml:"objId"`
}

