package model

import "encoding/xml"

type Cellvms struct {
	XMLName xml.Name `xml:"CellVMS"`
	ATTRIBUTES CellvmsAttributes `xml:"attributes"`
}

type CellvmsAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	VmsHoUeNumTh string `xml:"VmsHoUeNumTh"`
	VmsPrbDiffTh string `xml:"VmsPrbDiffTh"`
	VmsPrbLoadTh string `xml:"VmsPrbLoadTh"`
	VmsA3Offset string `xml:"VmsA3Offset"`
}

