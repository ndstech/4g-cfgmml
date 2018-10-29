package model

import "encoding/xml"

type Glocellrsvdpara struct {
	XMLName xml.Name `xml:"GLOCELLRSVDPARA"`
	ATTRIBUTES GlocellrsvdparaAttributes `xml:"attributes"`
}

type GlocellrsvdparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GLOCELLID string `xml:"GLOCELLID"`
	GLOCELLRSVDPARA1 string `xml:"GLOCELLRSVDPARA1"`
	GLOCELLRSVDPARA2 string `xml:"GLOCELLRSVDPARA2"`
	GLOCELLRSVDPARA3 string `xml:"GLOCELLRSVDPARA3"`
	GLOCELLRSVDPARA4 string `xml:"GLOCELLRSVDPARA4"`
	GLOCELLRSVDPARA5 string `xml:"GLOCELLRSVDPARA5"`
	GLOCELLRSVDPARA6 string `xml:"GLOCELLRSVDPARA6"`
}

