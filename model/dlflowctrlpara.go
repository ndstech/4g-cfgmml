package model

import "encoding/xml"

type Dlflowctrlpara struct {
	XMLName xml.Name `xml:"DLFLOWCTRLPARA"`
	ATTRIBUTES DlflowctrlparaAttributes `xml:"attributes"`
}

type DlflowctrlparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	BEAR string `xml:"BEAR"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	SWITCH string `xml:"SWITCH"`
	TD string `xml:"TD"`
	DR string `xml:"DR"`
	ITM string `xml:"ITM"`
	BWPRTSWITCH string `xml:"BWPRTSWITCH"`
	FAIRSWITCH string `xml:"FAIRSWITCH"`
	FAIRRATIO string `xml:"FAIRRATIO"`
	DLIUBMINBW string `xml:"DLIUBMINBW"`
	DLFLOWCTRLENHSW string `xml:"DLFLOWCTRLENHSW"`
	DLIUBBWMINPRORATIO string `xml:"DLIUBBWMINPRORATIO"`
}

