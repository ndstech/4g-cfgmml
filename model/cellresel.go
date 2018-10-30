package model

import "encoding/xml"

type Cellresel struct {
	XMLName xml.Name `xml:"CellResel"`
	ATTRIBUTES CellreselAttributes `xml:"attributes"`
}

type CellreselAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Qhyst string `xml:"Qhyst"`
	SpeedDepReselCfgInd string `xml:"SpeedDepReselCfgInd"`
	SNonIntraSearchCfgInd string `xml:"SNonIntraSearchCfgInd"`
	SNonIntraSearch string `xml:"SNonIntraSearch"`
	ThrshServLow string `xml:"ThrshServLow"`
	CellReselPriority string `xml:"CellReselPriority"`
	QRxLevMin string `xml:"QRxLevMin"`
	PMaxCfgInd string `xml:"PMaxCfgInd"`
	SIntraSearchCfgInd string `xml:"SIntraSearchCfgInd"`
	SIntraSearch string `xml:"SIntraSearch"`
	MeasBandWidthCfgInd string `xml:"MeasBandWidthCfgInd"`
	TReselEutran string `xml:"TReselEutran"`
	SpeedStateSfCfgInd string `xml:"SpeedStateSfCfgInd"`
	TReselEutranSfMedium string `xml:"TReselEutranSfMedium"`
	TReselEutranSfHigh string `xml:"TReselEutranSfHigh"`
	NeighCellConfig string `xml:"NeighCellConfig"`
	PresenceAntennaPort1 string `xml:"PresenceAntennaPort1"`
	SIntraSearchQ string `xml:"SIntraSearchQ"`
	SNonIntraSearchQ string `xml:"SNonIntraSearchQ"`
	ThrshServLowQCfgInd string `xml:"ThrshServLowQCfgInd"`
	QQualMinCfgInd string `xml:"QQualMinCfgInd"`
	QQualMin string `xml:"QQualMin"`
	TReselForNb string `xml:"TReselForNb"`
	TReselInterFreqForNb string `xml:"TReselInterFreqForNb"`
	TReselEutranCE string `xml:"TReselEutranCE"`
}

