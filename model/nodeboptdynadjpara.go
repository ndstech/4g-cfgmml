package model

import "encoding/xml"

type Nodeboptdynadjpara struct {
	XMLName xml.Name `xml:"NODEBOPTDYNADJPARA"`
	ATTRIBUTES NodeboptdynadjparaAttributes `xml:"attributes"`
}

type NodeboptdynadjparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DYNADJSW string `xml:"DYNADJSW"`
	DYNADJSTARTTIME string `xml:"DYNADJSTARTTIME"`
	DYNADJENDTIME string `xml:"DYNADJENDTIME"`
	OBJID string `xml:"OBJID"`
}

