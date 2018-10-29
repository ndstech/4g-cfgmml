package model

import "encoding/xml"

type Certchktsk struct {
	XMLName xml.Name `xml:"CERTCHKTSK"`
	ATTRIBUTES CertchktskAttributes `xml:"attributes"`
}

type CertchktskAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ISENABLE string `xml:"ISENABLE"`
	PERIOD string `xml:"PERIOD"`
	ALMRNG string `xml:"ALMRNG"`
	UPDATEMETHOD string `xml:"UPDATEMETHOD"`
}

