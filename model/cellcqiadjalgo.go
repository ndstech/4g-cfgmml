package model

import "encoding/xml"

type Cellcqiadjalgo struct {
	XMLName xml.Name `xml:"CELLCQIADJALGO"`
	ATTRIBUTES CellcqiadjalgoAttributes `xml:"attributes"`
}

type CellcqiadjalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InitDlIblerTarget string `xml:"InitDlIblerTarget"`
	InitDeltaCqi string `xml:"InitDeltaCqi"`
	CqiAdjStep string `xml:"CqiAdjStep"`
	CqiFilterCoefForMcs string `xml:"CqiFilterCoefForMcs"`
	VolteNackDeltaCqi string `xml:"VolteNackDeltaCqi"`
	DlVolteCqiAdjOptSw string `xml:"DlVolteCqiAdjOptSw"`
	CqiOptSwitch string `xml:"CqiOptSwitch"`
	InitDlIblerTargetforQCI2 string `xml:"InitDlIblerTargetforQCI2"`
	InitDlIblerTargetforVoLTE string `xml:"InitDlIblerTargetforVoLTE"`
	CellDeltaCqiSampSelThd string `xml:"CellDeltaCqiSampSelThd"`
}

