package model

import "encoding/xml"

type Tblangno struct {
	XMLName xml.Name `xml:"TBLANGNO"`
	ATTRIBUTES TblangnoAttributes `xml:"attributes"`
}

type TblangnoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LANNO string `xml:"LANNO"`
}

