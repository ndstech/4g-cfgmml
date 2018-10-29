package model

import "encoding/xml"

type Scserviceqos struct {
	XMLName xml.Name `xml:"SCSERVICEQOS"`
	ATTRIBUTES ScserviceqosAttributes `xml:"attributes"`
}

type ScserviceqosAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ScQosId string `xml:"ScQosId"`
	AppDns string `xml:"AppDns"`
	DlSgbr string `xml:"DlSgbr"`
	AppIdentType string `xml:"AppIdentType"`
}

