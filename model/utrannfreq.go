package model

import "encoding/xml"

type Utrannfreq struct {
	XMLName xml.Name `xml:"UTRANNFREQ"`
	ATTRIBUTES UtrannfreqAttributes `xml:"attributes"`
}

type UtrannfreqAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UtranDlArfcn string `xml:"UtranDlArfcn"`
	UtranVersion string `xml:"UtranVersion"`
	UtranFddTddType string `xml:"UtranFddTddType"`
	UtranUlArfcnCfgInd string `xml:"UtranUlArfcnCfgInd"`
	CellReselPriorityCfgInd string `xml:"CellReselPriorityCfgInd"`
	CellReselPriority string `xml:"CellReselPriority"`
	PmaxUtran string `xml:"PmaxUtran"`
	OffsetFreq string `xml:"OffsetFreq"`
	Qqualmin string `xml:"Qqualmin"`
	QRxLevMin string `xml:"QRxLevMin"`
	ThreshXHigh string `xml:"ThreshXHigh"`
	ThreshXLow string `xml:"ThreshXLow"`
	ThreshXHighQ string `xml:"ThreshXHighQ"`
	ThreshXLowQ string `xml:"ThreshXLowQ"`
	PsPriority string `xml:"PsPriority"`
	CsPriority string `xml:"CsPriority"`
	ConnFreqPriority string `xml:"ConnFreqPriority"`
	CsPsMixedPriority string `xml:"CsPsMixedPriority"`
	ContinuCoverageIndication string `xml:"ContinuCoverageIndication"`
	MlbFreqPriority string `xml:"MlbFreqPriority"`
	MasterBandFlag string `xml:"MasterBandFlag"`
	UtranRanSharingInd string `xml:"UtranRanSharingInd"`
	AnrInd string `xml:"AnrInd"`
	SrvccPriority string `xml:"SrvccPriority"`
	ULSharedFreqInd string `xml:"ULSharedFreqInd"`
	FreqPriorityForAnr string `xml:"FreqPriorityForAnr"`
	MlbTargetInd string `xml:"MlbTargetInd"`
	UtranFreqHighSpeedFlag string `xml:"UtranFreqHighSpeedFlag"`
}

