package model

import "encoding/xml"

type Appcert struct {
	XMLName    xml.Name          `xml:"APPCERT"`
	ATTRIBUTES AppcertAttributes `xml:"attributes"`
}

type AppcertAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	APPTYPE string   `xml:"APPTYPE"`
	APPCERT string   `xml:"APPCERT"`
}
