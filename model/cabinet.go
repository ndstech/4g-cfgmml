package model

import "encoding/xml"

type Cabinet struct {
	XMLName xml.Name `xml:"CABINET"`
	ATTRIBUTES CabinetAttributes `xml:"attributes"`
}

type CabinetAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	TYPE string `xml:"TYPE"`
	DESC string `xml:"DESC"`
	LOCATIONNAME string `xml:"LOCATIONNAME"`
}

