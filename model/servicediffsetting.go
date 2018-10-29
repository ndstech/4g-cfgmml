package model

import "encoding/xml"

type Servicediffsetting struct {
	XMLName xml.Name `xml:"SERVICEDIFFSETTING"`
	ATTRIBUTES ServicediffsettingAttributes `xml:"attributes"`
}

type ServicediffsettingAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	QueueWeight0 string `xml:"QueueWeight0"`
	QueueWeight1 string `xml:"QueueWeight1"`
	QueueWeight2 string `xml:"QueueWeight2"`
	QueueWeight3 string `xml:"QueueWeight3"`
	QueueWeight4 string `xml:"QueueWeight4"`
	QueueWeight5 string `xml:"QueueWeight5"`
	QueueWeight6 string `xml:"QueueWeight6"`
	QueueWeight7 string `xml:"QueueWeight7"`
	ObjId string `xml:"objId"`
}

