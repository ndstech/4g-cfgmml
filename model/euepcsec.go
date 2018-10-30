package model

import "encoding/xml"

type Euepcsec struct {
	XMLName xml.Name `xml:"EUEPCSec"`
	ATTRIBUTES EuepcsecAttributes `xml:"attributes"`
}

type EuepcsecAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CiphAlgo string `xml:"CiphAlgo"`
	IntAlgo string `xml:"IntAlgo"`
	ObjId string `xml:"objId"`
}

