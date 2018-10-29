package model

import "encoding/xml"

type Rscgrp struct {
	XMLName xml.Name `xml:"RSCGRP"`
	ATTRIBUTES RscgrpAttributes `xml:"attributes"`
}

type RscgrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	BEAR string `xml:"BEAR"`
	SBT string `xml:"SBT"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	RSCGRPID string `xml:"RSCGRPID"`
	USERLABEL string `xml:"USERLABEL"`
	RU string `xml:"RU"`
	TXBW string `xml:"TXBW"`
	RXBW string `xml:"RXBW"`
	TXCBS string `xml:"TXCBS"`
	TXEBS string `xml:"TXEBS"`
	OID string `xml:"OID"`
	WEIGHT string `xml:"WEIGHT"`
	TXCIR string `xml:"TXCIR"`
	RXCIR string `xml:"RXCIR"`
	TXPIR string `xml:"TXPIR"`
	RXPIR string `xml:"RXPIR"`
	TXPBS string `xml:"TXPBS"`
}

