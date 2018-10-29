package model

import "encoding/xml"

type Nodebruleactionpara struct {
	XMLName xml.Name `xml:"NODEBRULEACTIONPARA"`
	ATTRIBUTES NodebruleactionparaAttributes `xml:"attributes"`
}

type NodebruleactionparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RULEID string `xml:"RULEID"`
	MOCNAME string `xml:"MOCNAME"`
	PARANAME string `xml:"PARANAME"`
	ADMINSTATE string `xml:"ADMINSTATE"`
	PARAINTERVALVALUE string `xml:"PARAINTERVALVALUE"`
	PARAENUMVALUE string `xml:"PARAENUMVALUE"`
	USERLABEL string `xml:"USERLABEL"`
	OBJID string `xml:"OBJID"`
}

