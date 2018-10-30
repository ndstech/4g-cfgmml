package model

import "encoding/xml"

type Cellbf struct {
	XMLName xml.Name `xml:"CellBf"`
	ATTRIBUTES CellbfAttributes `xml:"attributes"`
}

type CellbfAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MaxBfRankPara string `xml:"MaxBfRankPara"`
	DualLayerBFAlgType string `xml:"DualLayerBFAlgType"`
	HighOrderMubfMaxLayer string `xml:"HighOrderMubfMaxLayer"`
	HighOrderMubfPairRule string `xml:"HighOrderMubfPairRule"`
	AntSelUEMubfPairMode string `xml:"AntSelUEMubfPairMode"`
	NonAntSelUEMubfPairMode string `xml:"NonAntSelUEMubfPairMode"`
	MovingUeMuBfScheme string `xml:"MovingUeMuBfScheme"`
	QualUEPortAvoidMode string `xml:"QualUEPortAvoidMode"`
	WaitPairingLayerThd string `xml:"WaitPairingLayerThd"`
	MassiveMIMOMubfPairRule string `xml:"MassiveMIMOMubfPairRule"`
	MultiLayerThdSwitchToTM7 string `xml:"MultiLayerThdSwitchToTM7"`
}

