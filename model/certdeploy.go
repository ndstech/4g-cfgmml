package model

import "encoding/xml"

type Certdeploy struct {
	XMLName xml.Name `xml:"CERTDEPLOY"`
	ATTRIBUTES CertdeployAttributes `xml:"attributes"`
}

type CertdeployAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DEPLOYTYPE string `xml:"DEPLOYTYPE"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
}

