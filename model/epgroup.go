package model

import "encoding/xml"

type Epgroup struct {
	XMLName xml.Name `xml:"EPGROUP"`
	ATTRIBUTES EpgroupAttributes `xml:"attributes"`
}

type EpgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	EPGROUPID string `xml:"EPGROUPID"`
	VRFIDX string `xml:"VRFIDX"`
	SCTPHOSTREFS string `xml:"SCTPHOSTREFS"`
	SCTPPEERREFS string `xml:"SCTPPEERREFS"`
	USERPLANEHOSTREFS string `xml:"USERPLANEHOSTREFS"`
	USERPLANEPEERREFS string `xml:"USERPLANEPEERREFS"`
	PACKETFILTERSWITCH string `xml:"PACKETFILTERSWITCH"`
	USERLABEL string `xml:"USERLABEL"`
	TYPEFLAG string `xml:"TYPEFLAG"`
	LNKPFMSW string `xml:"LNKPFMSW"`
	CTRLMODE string `xml:"CTRLMODE"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
}

