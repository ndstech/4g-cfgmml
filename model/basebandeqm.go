package model

import "encoding/xml"

type Basebandeqm struct {
	XMLName xml.Name `xml:"BASEBANDEQM"`
	ATTRIBUTES BasebandeqmAttributes `xml:"attributes"`
}

type BasebandeqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BASEBANDEQMID string `xml:"BASEBANDEQMID"`
	BASEBANDEQMTYPE string `xml:"BASEBANDEQMTYPE"`
	UMTSDEMMODE string `xml:"UMTSDEMMODE"`
	BASEBANDEQMBOARD string `xml:"BASEBANDEQMBOARD"`
}

