package model

import "encoding/xml"

type Eucoschulicscfg struct {
	XMLName xml.Name `xml:"EuCoSchULICSCfg"`
	ATTRIBUTES EucoschulicscfgAttributes `xml:"attributes"`
}

type EucoschulicscfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UlIcsAlgoSwitch string `xml:"UlIcsAlgoSwitch"`
	ObjId string `xml:"objId"`
}

