package model

import "encoding/xml"

type Retsubunit struct {
	XMLName xml.Name `xml:"RETSUBUNIT"`
	ATTRIBUTES RetsubunitAttributes `xml:"attributes"`
}

type RetsubunitAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DEVICENO string `xml:"DEVICENO"`
	SUBUNITNO string `xml:"SUBUNITNO"`
	CONNCN1 string `xml:"CONNCN1"`
	CONNSRN1 string `xml:"CONNSRN1"`
	CONNSN1 string `xml:"CONNSN1"`
	CONNPN1 string `xml:"CONNPN1"`
	CONNCN2 string `xml:"CONNCN2"`
	CONNSRN2 string `xml:"CONNSRN2"`
	CONNSN2 string `xml:"CONNSN2"`
	CONNPN2 string `xml:"CONNPN2"`
	TILT string `xml:"TILT"`
	AER string `xml:"AER"`
	SUBNAME string `xml:"SUBNAME"`
}

