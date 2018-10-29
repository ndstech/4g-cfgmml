package model

import "encoding/xml"

type Ulocellnoaccesspara struct {
	XMLName xml.Name `xml:"ULOCELLNOACCESSPARA"`
	ATTRIBUTES UlocellnoaccessparaAttributes `xml:"attributes"`
}

type UlocellnoaccessparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULOCELLID string `xml:"ULOCELLID"`
	NOUETIMER string `xml:"NOUETIMER"`
	NOFSTRLTIMER string `xml:"NOFSTRLTIMER"`
	AUTORCVRMTHD string `xml:"AUTORCVRMTHD"`
}

