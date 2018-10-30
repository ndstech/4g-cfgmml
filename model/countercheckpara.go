package model

import "encoding/xml"

type Countercheckpara struct {
	XMLName xml.Name `xml:"CounterCheckPara"`
	ATTRIBUTES CountercheckparaAttributes `xml:"attributes"`
}

type CountercheckparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CounterCheckTimer string `xml:"CounterCheckTimer"`
	CounterCheckCountNum string `xml:"CounterCheckCountNum"`
	CounterCheckUserRelSwitch string `xml:"CounterCheckUserRelSwitch"`
	ObjId string `xml:"objId"`
}

