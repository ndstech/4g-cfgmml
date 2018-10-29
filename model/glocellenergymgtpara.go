package model

import "encoding/xml"

type Glocellenergymgtpara struct {
	XMLName xml.Name `xml:"GLOCELLENERGYMGTPARA"`
	ATTRIBUTES GlocellenergymgtparaAttributes `xml:"attributes"`
}

type GlocellenergymgtparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GLOCELLID string `xml:"GLOCELLID"`
	MAINBCCHPWRDTEN string `xml:"MAINBCCHPWRDTEN"`
}

