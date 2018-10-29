package model

import "encoding/xml"

type Cnoperator struct {
	XMLName xml.Name `xml:"CNOPERATOR"`
	ATTRIBUTES CnoperatorAttributes `xml:"attributes"`
}

type CnoperatorAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	CnOperatorName string `xml:"CnOperatorName"`
	CnOperatorType string `xml:"CnOperatorType"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	ObjId string `xml:"objId"`
	PlmnMode string `xml:"PlmnMode"`
	OperatorFunSwitch string `xml:"OperatorFunSwitch"`
}

