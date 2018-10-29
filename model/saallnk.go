package model

import "encoding/xml"

type Saallnk struct {
	XMLName xml.Name `xml:"SAALLNK"`
	ATTRIBUTES SaallnkAttributes `xml:"attributes"`
}

type SaallnkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SAALNO string `xml:"SAALNO"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	JNRSCGRP string `xml:"JNRSCGRP"`
	VPI string `xml:"VPI"`
	VCI string `xml:"VCI"`
	ST string `xml:"ST"`
	PCR string `xml:"PCR"`
	CCTM string `xml:"CCTM"`
	POLL string `xml:"POLL"`
	IDLE string `xml:"IDLE"`
	NRTM string `xml:"NRTM"`
	KATM string `xml:"KATM"`
	MAXCC string `xml:"MAXCC"`
	MAXPD string `xml:"MAXPD"`
	MAXLE string `xml:"MAXLE"`
	SBT string `xml:"SBT"`
	RU string `xml:"RU"`
}

