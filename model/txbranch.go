package model

import "encoding/xml"

type Txbranch struct {
	XMLName xml.Name `xml:"TXBRANCH"`
	ATTRIBUTES TxbranchAttributes `xml:"attributes"`
}

type TxbranchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	TXNO string `xml:"TXNO"`
	TXSW string `xml:"TXSW"`
}

