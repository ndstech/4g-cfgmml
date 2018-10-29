package model

import "encoding/xml"

type Tolcalg struct {
	XMLName xml.Name `xml:"TOLCALG"`
	ATTRIBUTES TolcalgAttributes `xml:"attributes"`
}

type TolcalgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TRMULOLCSWITCH string `xml:"TRMULOLCSWITCH"`
	TRMDLOLCSWITCH string `xml:"TRMDLOLCSWITCH"`
	TRMULOLCTRIGTH string `xml:"TRMULOLCTRIGTH"`
	TRMULOLCRELTH string `xml:"TRMULOLCRELTH"`
	TRMOLCRELBEARERNUM string `xml:"TRMOLCRELBEARERNUM"`
	TRMDLOLCTRIGTH string `xml:"TRMDLOLCTRIGTH"`
	TRMDLOLCRELTH string `xml:"TRMDLOLCRELTH"`
}

