package model

import "encoding/xml"

type Ulocellsectoreqm struct {
	XMLName xml.Name `xml:"ULOCELLSECTOREQM"`
	ATTRIBUTES UlocellsectoreqmAttributes `xml:"attributes"`
}

type UlocellsectoreqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULOCELLID string `xml:"ULOCELLID"`
	SECTOREQMID string `xml:"SECTOREQMID"`
	MAXPWR string `xml:"MAXPWR"`
	SECTOREQMPROPERTY string `xml:"SECTOREQMPROPERTY"`
}

