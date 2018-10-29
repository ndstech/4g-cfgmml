package model

import "encoding/xml"

type Gtrxgroupsectoreqm struct {
	XMLName xml.Name `xml:"GTRXGROUPSECTOREQM"`
	ATTRIBUTES GtrxgroupsectoreqmAttributes `xml:"attributes"`
}

type GtrxgroupsectoreqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GTRXGROUPID string `xml:"GTRXGROUPID"`
	SECTOREQMID string `xml:"SECTOREQMID"`
	MAXPOWER string `xml:"MAXPOWER"`
}

