package model

import "encoding/xml"

type Certcfg struct {
	XMLName xml.Name `xml:"CERTCFG"`
	ATTRIBUTES CertcfgAttributes `xml:"attributes"`
}

type CertcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IKECHECKSW string `xml:"IKECHECKSW"`
}

