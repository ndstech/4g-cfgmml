package model

import "encoding/xml"

type Gerannfreqgroup struct {
	XMLName xml.Name `xml:"GeranNfreqGroup"`
	ATTRIBUTES GerannfreqgroupAttributes `xml:"attributes"`
}

type GerannfreqgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BcchGroupId string `xml:"BcchGroupId"`
	GeranVersion string `xml:"GeranVersion"`
	StartingArfcn string `xml:"StartingArfcn"`
	BandIndicator string `xml:"BandIndicator"`
	CellReselPriorityCfgInd string `xml:"CellReselPriorityCfgInd"`
	CellReselPriority string `xml:"CellReselPriority"`
	PmaxGeranCfgInd string `xml:"PmaxGeranCfgInd"`
	QRxLevMin string `xml:"QRxLevMin"`
	ThreshXHigh string `xml:"ThreshXHigh"`
	ThreshXLow string `xml:"ThreshXLow"`
	OffsetFreq string `xml:"OffsetFreq"`
	NccPermitted string `xml:"NccPermitted"`
	ConnFreqPriority string `xml:"ConnFreqPriority"`
	ContinuCoverageIndication string `xml:"ContinuCoverageIndication"`
	GeranRanSharingInd string `xml:"GeranRanSharingInd"`
	AnrInd string `xml:"AnrInd"`
	FreqPriorityForAnr string `xml:"FreqPriorityForAnr"`
	PmaxGeran string `xml:"PmaxGeran"`
}

