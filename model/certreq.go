package model

import "encoding/xml"

type Certreq struct {
	XMLName xml.Name `xml:"CERTREQ"`
	ATTRIBUTES CertreqAttributes `xml:"attributes"`
}

type CertreqAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	COMMNAME string `xml:"COMMNAME"`
	USERADDINFO string `xml:"USERADDINFO"`
	COUNTRY string `xml:"COUNTRY"`
	ORG string `xml:"ORG"`
	ORGUNIT string `xml:"ORGUNIT"`
	STATEPROVINCENAME string `xml:"STATEPROVINCENAME"`
	LOCALITY string `xml:"LOCALITY"`
	KEYUSAGE string `xml:"KEYUSAGE"`
	SIGNALG string `xml:"SIGNALG"`
	KEYSIZE string `xml:"KEYSIZE"`
	LOCALNAME string `xml:"LOCALNAME"`
	LOCALIP string `xml:"LOCALIP"`
}

