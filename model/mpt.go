package model

import "encoding/xml"

type Mpt struct {
	XMLName xml.Name `xml:"MPT"`
	ATTRIBUTES MptAttributes `xml:"attributes"`
}

type MptAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	TYPE string `xml:"TYPE"`
	SWITCH string `xml:"SWITCH"`
	BRDSPEC string `xml:"BRDSPEC"`
	MPTWORKMODE string `xml:"MPTWORKMODE"`
}

