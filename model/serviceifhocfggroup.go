package model

import "encoding/xml"

type Serviceifhocfggroup struct {
	XMLName xml.Name `xml:"SERVICEIFHOCFGGROUP"`
	ATTRIBUTES ServiceifhocfggroupAttributes `xml:"attributes"`
}

type ServiceifhocfggroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	ServiceIfHoCfgGroupId string `xml:"ServiceIfHoCfgGroupId"`
	InterFreqHoState string `xml:"InterFreqHoState"`
	A4RptWaitingTimer string `xml:"A4RptWaitingTimer"`
	A4TimeToTriger string `xml:"A4TimeToTriger"`
}

