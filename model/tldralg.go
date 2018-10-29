package model

import "encoding/xml"

type Tldralg struct {
	XMLName xml.Name `xml:"TLDRALG"`
	ATTRIBUTES TldralgAttributes `xml:"attributes"`
}

type TldralgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TRMULLDRTRGTH string `xml:"TRMULLDRTRGTH"`
	TRMDLLDRTRGTH string `xml:"TRMDLLDRTRGTH"`
	TRMULLDRCLRTH string `xml:"TRMULLDRCLRTH"`
	TRMDLLDRCLRTH string `xml:"TRMDLLDRCLRTH"`
	TRMULMLDTRGTH string `xml:"TRMULMLDTRGTH"`
	TRMDLMLDTRGTH string `xml:"TRMDLMLDTRGTH"`
	TRMULMLDCLRTH string `xml:"TRMULMLDCLRTH"`
	TRMDLMLDCLRTH string `xml:"TRMDLMLDCLRTH"`
}

