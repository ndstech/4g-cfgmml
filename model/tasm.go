package model

import "encoding/xml"

type Tasm struct {
	XMLName xml.Name `xml:"TASM"`
	ATTRIBUTES TasmAttributes `xml:"attributes"`
}

type TasmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SRCNO string `xml:"SRCNO"`
	CLKSRC string `xml:"CLKSRC"`
	MODE string `xml:"MODE"`
	ALGORITHM string `xml:"ALGORITHM"`
	FREECOUNT string `xml:"FREECOUNT"`
	SEARCHCOUNT string `xml:"SEARCHCOUNT"`
	HOLDCOUNT string `xml:"HOLDCOUNT"`
	LOCKCOUNT string `xml:"LOCKCOUNT"`
	SAMPLETIME string `xml:"SAMPLETIME"`
	SYNMODE string `xml:"SYNMODE"`
	PERIOD string `xml:"PERIOD"`
	TM string `xml:"TM"`
	CLKSYNCMODE string `xml:"CLKSYNCMODE"`
	QL string `xml:"QL"`
	FREQUENCE string `xml:"FREQUENCE"`
	NETMODE string `xml:"NETMODE"`
	FLAG string `xml:"FLAG"`
	DAY string `xml:"DAY"`
	GSM_FBOFFSET string `xml:"GSM_FBOFFSET"`
	SYSCLKSRC string `xml:"SYSCLKSRC"`
	CLOUDSRC string `xml:"CLOUDSRC"`
	FRAMESYNCSW string `xml:"FRAMESYNCSW"`
	GSMFRAMESYNCSW string `xml:"GSMFRAMESYNCSW"`
	STANDARD string `xml:"STANDARD"`
	ENSYSCLKCHKSW string `xml:"ENSYSCLKCHKSW"`
	ATRSW string `xml:"ATRSW"`
	FNSYNCSW string `xml:"FNSYNCSW"`
	DATE string `xml:"DATE"`
	TIME string `xml:"TIME"`
	LEAPSECONDSCHGDATE string `xml:"LEAPSECONDSCHGDATE"`
	LEAPSECONDSCHGTIME string `xml:"LEAPSECONDSCHGTIME"`
	CRTGPSTOUTCLEAPSECONDS string `xml:"CRTGPSTOUTCLEAPSECONDS"`
	NEXTGPSTOUTCLEAPSECONDS string `xml:"NEXTGPSTOUTCLEAPSECONDS"`
	LPFNSYNCSW string `xml:"LPFNSYNCSW"`
}

