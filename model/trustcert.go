package model

import "encoding/xml"

type Trustcert struct {
	XMLName xml.Name `xml:"TRUSTCERT"`
	ATTRIBUTES TrustcertAttributes `xml:"attributes"`
}

type TrustcertAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CERTNAME string `xml:"CERTNAME"`
}

