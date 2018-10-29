package model

import "encoding/xml"

type Ingchktsk struct {
	XMLName xml.Name `xml:"INGCHKTSK"`
	ATTRIBUTES IngchktskAttributes `xml:"attributes"`
}

type IngchktskAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	FLINGCHKSW string `xml:"FLINGCHKSW"`
	FLINGCHKTM string `xml:"FLINGCHKTM"`
	FLINGCHKITEM string `xml:"FLINGCHKITEM"`
}

