package model

import "encoding/xml"

type Ulocellr99algpara struct {
	XMLName xml.Name `xml:"ULOCELLR99ALGPARA"`
	ATTRIBUTES Ulocellr99algparaAttributes `xml:"attributes"`
}

type Ulocellr99algparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULOCELLID string `xml:"ULOCELLID"`
	PWRCNGCACSW string `xml:"PWRCNGCACSW"`
	SHORTTH string `xml:"SHORTTH"`
	LONGTH string `xml:"LONGTH"`
	SHORTTHFORRTRRC string `xml:"SHORTTHFORRTRRC"`
}

