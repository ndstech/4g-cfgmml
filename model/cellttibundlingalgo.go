package model

import "encoding/xml"

type Cellttibundlingalgo struct {
	XMLName xml.Name `xml:"CELLTTIBUNDLINGALGO"`
	ATTRIBUTES CellttibundlingalgoAttributes `xml:"attributes"`
}

type CellttibundlingalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SinrThdToTrigTtib string `xml:"SinrThdToTrigTtib"`
	SinrThdToTrigVideoTtib string `xml:"SinrThdToTrigVideoTtib"`
	TtiBundlingAlgoSw string `xml:"TtiBundlingAlgoSw"`
	R12TtiBHarqMaxTxNum string `xml:"R12TtiBHarqMaxTxNum"`
	R12TtiBundlingSwitch string `xml:"R12TtiBundlingSwitch"`
	SinrThdToTrigR12TtiB string `xml:"SinrThdToTrigR12TtiB"`
}

