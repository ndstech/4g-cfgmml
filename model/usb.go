package model

import "encoding/xml"

type Usb struct {
	XMLName xml.Name `xml:"USB"`
	ATTRIBUTES UsbAttributes `xml:"attributes"`
}

type UsbAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	SWITCH string `xml:"SWITCH"`
}

