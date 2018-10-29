package model

import "encoding/xml"

type Brdresassignment struct {
	XMLName xml.Name `xml:"BRDRESASSIGNMENT"`
	ATTRIBUTES BrdresassignmentAttributes `xml:"attributes"`
}

type BrdresassignmentAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	BRDASSIGNMENT string `xml:"BRDASSIGNMENT"`
}

