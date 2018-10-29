package model

import "encoding/xml"

type Nodebmulticellgrp struct {
	XMLName xml.Name `xml:"NODEBMULTICELLGRP"`
	ATTRIBUTES NodebmulticellgrpAttributes `xml:"attributes"`
}

type NodebmulticellgrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	MULTICELLGRPID string `xml:"MULTICELLGRPID"`
	MULTICELLGRPTYPE string `xml:"MULTICELLGRPTYPE"`
	OBJID string `xml:"OBJID"`
	ULOCELLREF string `xml:"ULOCELLREF"`
}

