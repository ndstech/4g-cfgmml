package model

import "encoding/xml"

type Tbdspinfo struct {
	XMLName xml.Name `xml:"TBDSPINFO"`
	ATTRIBUTES TbdspinfoAttributes `xml:"attributes"`
}

type TbdspinfoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LANNO string `xml:"LANNO"`
	VERSION string `xml:"VERSION"`
}

