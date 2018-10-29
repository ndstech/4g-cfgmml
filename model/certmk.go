package model

import "encoding/xml"

type Certmk struct {
	XMLName xml.Name `xml:"CERTMK"`
	ATTRIBUTES CertmkAttributes `xml:"attributes"`
}

type CertmkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	APPCERT string `xml:"APPCERT"`
	CASW string `xml:"CASW"`
}

