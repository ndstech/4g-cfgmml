package model

import "encoding/xml"

type Serviceirhocfggroup struct {
	XMLName xml.Name `xml:"SERVICEIRHOCFGGROUP"`
	ATTRIBUTES ServiceirhocfggroupAttributes `xml:"attributes"`
}

type ServiceirhocfggroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	ServiceIrHoCfgGroupId string `xml:"ServiceIrHoCfgGroupId"`
	InterRatHoState string `xml:"InterRatHoState"`
	ServiceIrMeasMode string `xml:"ServiceIrMeasMode"`
}

