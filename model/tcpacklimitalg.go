package model

import "encoding/xml"

type Tcpacklimitalg struct {
	XMLName xml.Name `xml:"TCPACKLIMITALG"`
	ATTRIBUTES TcpacklimitalgAttributes `xml:"attributes"`
}

type TcpacklimitalgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TCPACKLIMITSWITCH string `xml:"TCPACKLIMITSWITCH"`
}

