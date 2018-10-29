package model

import "encoding/xml"

type Application struct {
	XMLName xml.Name `xml:"APPLICATION"`
	ATTRIBUTES ApplicationAttributes `xml:"attributes"`
}

type ApplicationAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AID string `xml:"AID"`
	AT string `xml:"AT"`
	AN string `xml:"AN"`
	SWVERSION string `xml:"SWVERSION"`
	HOTPATCHVERSION string `xml:"HOTPATCHVERSION"`
	AV string `xml:"AV"`
	APPHOTPATCHVERSION string `xml:"APPHOTPATCHVERSION"`
	APPMNTMODE string `xml:"APPMNTMODE"`
	APPST string `xml:"APPST"`
	APPET string `xml:"APPET"`
	APPMMSETREMARK string `xml:"APPMMSETREMARK"`
	EXTAPPTYPE string `xml:"EXTAPPTYPE"`
}

