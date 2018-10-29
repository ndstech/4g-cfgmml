package model

import "encoding/xml"

type Qoehocommoncfg struct {
	XMLName xml.Name `xml:"QOEHOCOMMONCFG"`
	ATTRIBUTES QoehocommoncfgAttributes `xml:"attributes"`
}

type QoehocommoncfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	QoEBasedHandoverStat string `xml:"QoEBasedHandoverStat"`
	QoEBasedHandoverLast string `xml:"QoEBasedHandoverLast"`
	ObjId string `xml:"objId"`
}

