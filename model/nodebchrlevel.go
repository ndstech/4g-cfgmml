package model

import "encoding/xml"

type Nodebchrlevel struct {
	XMLName xml.Name `xml:"NODEBCHRLEVEL"`
	ATTRIBUTES NodebchrlevelAttributes `xml:"attributes"`
}

type NodebchrlevelAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	COMMCHREVENTSW string `xml:"COMMCHREVENTSW"`
	USERCHREVENTSW string `xml:"USERCHREVENTSW"`
	COMMCHRPRDEVENTSW string `xml:"COMMCHRPRDEVENTSW"`
	USERCHRPRDEVENTSW string `xml:"USERCHRPRDEVENTSW"`
	HSDPAKQIDIAGTHPTHD string `xml:"HSDPAKQIDIAGTHPTHD"`
	HSUPAKQIDIAGTHPTHD string `xml:"HSUPAKQIDIAGTHPTHD"`
}

