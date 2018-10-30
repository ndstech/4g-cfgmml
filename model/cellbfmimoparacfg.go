package model

import "encoding/xml"

type Cellbfmimoparacfg struct {
	XMLName xml.Name `xml:"CellBfMimoParaCfg"`
	ATTRIBUTES CellbfmimoparacfgAttributes `xml:"attributes"`
}

type CellbfmimoparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BfMimoAdaptiveSwitch string `xml:"BfMimoAdaptiveSwitch"`
	BfMimoAdapWithoutTm2 string `xml:"BfMimoAdapWithoutTm2"`
	BfSingleToDualThdOffset string `xml:"BfSingleToDualThdOffset"`
	TM3Rank2ToDualBfThdOffset string `xml:"TM3Rank2ToDualBfThdOffset"`
	AsUeDualBfToTM3Rank2Offset string `xml:"AsUeDualBfToTM3Rank2Offset"`
	AsUeTM3Rank2ToDualBfOffset string `xml:"AsUeTM3Rank2ToDualBfOffset"`
	AsUeBfSingleToDualOffset string `xml:"AsUeBfSingleToDualOffset"`
	DualBfToTM3Rank2Offset string `xml:"DualBfToTM3Rank2Offset"`
	OffsetOfInStat string `xml:"OffsetOfInStat"`
	OffsetOfOutStat string `xml:"OffsetOfOutStat"`
	AntBasedBfMimoAlgoSelect string `xml:"AntBasedBfMimoAlgoSelect"`
	Tm3DirectToDualBfSwitch string `xml:"Tm3DirectToDualBfSwitch"`
	SccBfMimoAdaptiveSwitch string `xml:"SccBfMimoAdaptiveSwitch"`
	HighLowSpeedUeThdOffset string `xml:"HighLowSpeedUeThdOffset"`
	TmAccelerationSwitch string `xml:"TmAccelerationSwitch"`
	BfMimoAdapWithTm4Switch string `xml:"BfMimoAdapWithTm4Switch"`
	Tm3AndTm9ThdOffset string `xml:"Tm3AndTm9ThdOffset"`
	BfTo2LayerMubfThdOffset string `xml:"BfTo2LayerMubfThdOffset"`
	BfTo4LayerMubfThdOffset string `xml:"BfTo4LayerMubfThdOffset"`
	Tm3Rank2ToTm9Rank4Offset string `xml:"Tm3Rank2ToTm9Rank4Offset"`
	Tm9Rank4ToTm3Rank2Offset string `xml:"Tm9Rank4ToTm3Rank2Offset"`
	VolteMimoAdaptiveSwitch string `xml:"VolteMimoAdaptiveSwitch"`
	HoBfThdAdjSwitch string `xml:"HoBfThdAdjSwitch"`
	CaSccAddThldOffset string `xml:"CaSccAddThldOffset"`
	CaSccDelThldOffset string `xml:"CaSccDelThldOffset"`
	BfMimoAlgoOptSwitch string `xml:"BfMimoAlgoOptSwitch"`
	InitialBfMimoType string `xml:"InitialBfMimoType"`
	MultiLayerPairIsoThd string `xml:"MultiLayerPairIsoThd"`
}

