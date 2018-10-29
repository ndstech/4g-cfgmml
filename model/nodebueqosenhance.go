package model

import "encoding/xml"

type Nodebueqosenhance struct {
	XMLName xml.Name `xml:"NODEBUEQOSENHANCE"`
	ATTRIBUTES NodebueqosenhanceAttributes `xml:"attributes"`
}

type NodebueqosenhanceAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULGOLDUEBPS string `xml:"ULGOLDUEBPS"`
	ULSILVERUEBPS string `xml:"ULSILVERUEBPS"`
	ULCOPPERUEBPS string `xml:"ULCOPPERUEBPS"`
	DLGOLDUEBPS string `xml:"DLGOLDUEBPS"`
	DLSILVERUEBPS string `xml:"DLSILVERUEBPS"`
	DLCOPPERUEBPS string `xml:"DLCOPPERUEBPS"`
	OBJID string `xml:"OBJID"`
	LOWTRAFFENHUSERNUMTHD string `xml:"LOWTRAFFENHUSERNUMTHD"`
	TBSIZEFOR10MSLOWTRAFFENH string `xml:"TBSIZEFOR10MSLOWTRAFFENH"`
	TBSIZEFOR2MSLOWTRAFFENH string `xml:"TBSIZEFOR2MSLOWTRAFFENH"`
}

