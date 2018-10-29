package model

import "encoding/xml"

type Rxbranch struct {
	XMLName xml.Name `xml:"RXBRANCH"`
	ATTRIBUTES RxbranchAttributes `xml:"attributes"`
}

type RxbranchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	RXNO string `xml:"RXNO"`
	RXSW string `xml:"RXSW"`
	ATTEN string `xml:"ATTEN"`
	RTWPINITADJ0 string `xml:"RTWPINITADJ0"`
	RTWPINITADJ1 string `xml:"RTWPINITADJ1"`
	RTWPINITADJ2 string `xml:"RTWPINITADJ2"`
	RTWPINITADJ3 string `xml:"RTWPINITADJ3"`
	RTWPINITADJ4 string `xml:"RTWPINITADJ4"`
	RTWPINITADJ5 string `xml:"RTWPINITADJ5"`
	RTWPINITADJ6 string `xml:"RTWPINITADJ6"`
	RTWPINITADJ7 string `xml:"RTWPINITADJ7"`
}

