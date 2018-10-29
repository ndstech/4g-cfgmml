package model

import "encoding/xml"

type Cellricalgo struct {
	XMLName xml.Name `xml:"CELLRICALGO"`
	ATTRIBUTES CellricalgoAttributes `xml:"attributes"`
}

type CellricalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	RicAlgoSwitch string `xml:"RicAlgoSwitch"`
	MuteUpPTSSymNum string `xml:"MuteUpPTSSymNum"`
	MuteULSym string `xml:"MuteULSym"`
	UlInterferenceThd string `xml:"UlInterferenceThd"`
	DuctingRemoteInfThd string `xml:"DuctingRemoteInfThd"`
	InfAvoidDetPeriodNum string `xml:"InfAvoidDetPeriodNum"`
	InfAvoidRecDetPeriodNum string `xml:"InfAvoidRecDetPeriodNum"`
	RemoteInfAdpAvoidSwitch string `xml:"RemoteInfAdpAvoidSwitch"`
	UlInterferenceSymbMaxDiff string `xml:"UlInterferenceSymbMaxDiff"`
	DuctDLSubfrmShutoffSwitch string `xml:"DuctDLSubfrmShutoffSwitch"`
	UlInterferenceDuration string `xml:"UlInterferenceDuration"`
	DuctDLSubfrmShutoffInfThld string `xml:"DuctDLSubfrmShutoffInfThld"`
	DuctDLSubfrmShutoffOptSw string `xml:"DuctDLSubfrmShutoffOptSw"`
	DuctInfDetPeriodCount string `xml:"DuctInfDetPeriodCount"`
	DuctInfPwrDiffInThld string `xml:"DuctInfPwrDiffInThld"`
	DuctInfPwrDiffOutThld string `xml:"DuctInfPwrDiffOutThld"`
	DuctInfRateThld string `xml:"DuctInfRateThld"`
	DuctSubfrmShutoffDepth string `xml:"DuctSubfrmShutoffDepth"`
}

