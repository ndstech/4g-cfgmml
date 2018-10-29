package model

import "encoding/xml"

type Glocellrlalmpara struct {
	XMLName xml.Name `xml:"GLOCELLRLALMPARA"`
	ATTRIBUTES GlocellrlalmparaAttributes `xml:"attributes"`
}

type GlocellrlalmparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GLOCELLID string `xml:"GLOCELLID"`
	ARELWTHD string `xml:"ARELWTHD"`
	AREDISTHD string `xml:"AREDISTHD"`
	NOTRASP string `xml:"NOTRASP"`
	WLASTIME string `xml:"WLASTIME"`
	WLAETIME string `xml:"WLAETIME"`
	BWTHD string `xml:"BWTHD"`
	GUNFAULTDECTSW string `xml:"GUNFAULTDECTSW"`
	GUNCHDROPDECTTHRD string `xml:"GUNCHDROPDECTTHRD"`
	GUNCHDROPCLRTHRD string `xml:"GUNCHDROPCLRTHRD"`
	GUNBADQUALDETCTTHRD string `xml:"GUNBADQUALDETCTTHRD"`
	GUNBADQUALCLRTHRD string `xml:"GUNBADQUALCLRTHRD"`
	RSSIUNBALANCEDTHRD string `xml:"RSSIUNBALANCEDTHRD"`
}

