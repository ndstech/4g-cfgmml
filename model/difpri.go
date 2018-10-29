package model

import "encoding/xml"

type Difpri struct {
	XMLName xml.Name `xml:"DIFPRI"`
	ATTRIBUTES DifpriAttributes `xml:"attributes"`
}

type DifpriAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PRIRULE string `xml:"PRIRULE"`
	SIGPRI string `xml:"SIGPRI"`
	OMHIGHPRI string `xml:"OMHIGHPRI"`
	OMLOWPRI string `xml:"OMLOWPRI"`
	IPCLKPRI string `xml:"IPCLKPRI"`
}

