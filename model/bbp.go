package model

import "encoding/xml"

type Bbp struct {
	XMLName xml.Name `xml:"BBP"`
	ATTRIBUTES BbpAttributes `xml:"attributes"`
}

type BbpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	TYPE string `xml:"TYPE"`
	ADMSTATE string `xml:"ADMSTATE"`
	TIME string `xml:"TIME"`
	BLKTP string `xml:"BLKTP"`
	HCE string `xml:"HCE"`
	CPRIEX string `xml:"CPRIEX"`
	BRDSPEC string `xml:"BRDSPEC"`
	CCNE string `xml:"CCNE"`
	BBWS string `xml:"BBWS"`
	SRT string `xml:"SRT"`
	CPRIEXMODE string `xml:"CPRIEXMODE"`
	WM string `xml:"WM"`
	CPRIITFTYPE string `xml:"CPRIITFTYPE"`
}

