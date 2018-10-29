package model

import "encoding/xml"

type Subrack struct {
	XMLName xml.Name `xml:"SUBRACK"`
	ATTRIBUTES SubrackAttributes `xml:"attributes"`
}

type SubrackAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	TYPE string `xml:"TYPE"`
	DESC string `xml:"DESC"`
}

