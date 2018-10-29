package model

import "encoding/xml"

type Cellservicediffcfg struct {
	XMLName xml.Name `xml:"CELLSERVICEDIFFCFG"`
	ATTRIBUTES CellservicediffcfgAttributes `xml:"attributes"`
}

type CellservicediffcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ServiceDiffSwitch string `xml:"ServiceDiffSwitch"`
	CellConStatePrbThd string `xml:"CellConStatePrbThd"`
	CellOverloadStatePrbThd string `xml:"CellOverloadStatePrbThd"`
	CellConStateUeNumThd string `xml:"CellConStateUeNumThd"`
	CellOverloadStateUeNumThd string `xml:"CellOverloadStateUeNumThd"`
	QueueServiceWeight0 string `xml:"QueueServiceWeight0"`
	QueueServiceWeight1 string `xml:"QueueServiceWeight1"`
	QueueServiceWeight2 string `xml:"QueueServiceWeight2"`
	QueueServiceWeight3 string `xml:"QueueServiceWeight3"`
	QueueServiceWeight4 string `xml:"QueueServiceWeight4"`
	QueueServiceWeight5 string `xml:"QueueServiceWeight5"`
	QueueServiceWeight6 string `xml:"QueueServiceWeight6"`
}

