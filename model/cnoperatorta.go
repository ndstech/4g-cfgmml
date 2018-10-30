package model

import "encoding/xml"

type Cnoperatorta struct {
	XMLName xml.Name `xml:"CnOperatorTa"`
	ATTRIBUTES CnoperatortaAttributes `xml:"attributes"`
}

type CnoperatortaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TrackingAreaId string `xml:"TrackingAreaId"`
	CnOperatorId string `xml:"CnOperatorId"`
	Tac string `xml:"Tac"`
	ObjId string `xml:"objId"`
	NbIotTaFlag string `xml:"NbIotTaFlag"`
}

