package model

import "encoding/xml"

type Lansw struct {
	XMLName xml.Name `xml:"LANSW"`
	ATTRIBUTES LanswAttributes `xml:"attributes"`
}

type LanswAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BCPKTRATETHD string `xml:"BCPKTRATETHD"`
	MACAGINGTIME string `xml:"MACAGINGTIME"`
	DYNAMICMACLRNSW string `xml:"DYNAMICMACLRNSW"`
	SRCMACATTACKCHKSW string `xml:"SRCMACATTACKCHKSW"`
	SRCMACATTACKALMTHD string `xml:"SRCMACATTACKALMTHD"`
}

