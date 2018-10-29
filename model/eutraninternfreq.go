package model

import "encoding/xml"

type Eutraninternfreq struct {
	XMLName xml.Name `xml:"EUTRANINTERNFREQ"`
	ATTRIBUTES EutraninternfreqAttributes `xml:"attributes"`
}

type EutraninternfreqAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlEarfcn string `xml:"DlEarfcn"`
	UlEarfcnCfgInd string `xml:"UlEarfcnCfgInd"`
	CellReselPriorityCfgInd string `xml:"CellReselPriorityCfgInd"`
	CellReselPriority string `xml:"CellReselPriority"`
	EutranReselTime string `xml:"EutranReselTime"`
	SpeedDependSPCfgInd string `xml:"SpeedDependSPCfgInd"`
	MeasBandWidth string `xml:"MeasBandWidth"`
	QoffsetFreq string `xml:"QoffsetFreq"`
	ThreshXhigh string `xml:"ThreshXhigh"`
	ThreshXlow string `xml:"ThreshXlow"`
	QRxLevMin string `xml:"QRxLevMin"`
	PmaxCfgInd string `xml:"PmaxCfgInd"`
	NeighCellConfig string `xml:"NeighCellConfig"`
	PresenceAntennaPort1 string `xml:"PresenceAntennaPort1"`
	InterFreqHoEventType string `xml:"InterFreqHoEventType"`
	ThreshXhighQ string `xml:"ThreshXhighQ"`
	ThreshXlowQ string `xml:"ThreshXlowQ"`
	QqualMinCfgInd string `xml:"QqualMinCfgInd"`
	ConnFreqPriority string `xml:"ConnFreqPriority"`
	MlbTargetInd string `xml:"MlbTargetInd"`
	FreqPriBasedHoMeasFlag string `xml:"FreqPriBasedHoMeasFlag"`
	IdleMlbUEReleaseRatio string `xml:"IdleMlbUEReleaseRatio"`
	MlbFreqPriority string `xml:"MlbFreqPriority"`
	QoffsetFreqConn string `xml:"QoffsetFreqConn"`
	MeasFreqPriority string `xml:"MeasFreqPriority"`
	IfHoThdRsrpOffset string `xml:"IfHoThdRsrpOffset"`
	IfMlbThdRsrpOffset string `xml:"IfMlbThdRsrpOffset"`
	MasterBandFlag string `xml:"MasterBandFlag"`
	InterFreqRanSharingInd string `xml:"InterFreqRanSharingInd"`
	InterFreqHighSpeedFlag string `xml:"InterFreqHighSpeedFlag"`
	AnrInd string `xml:"AnrInd"`
	VoipPriority string `xml:"VoipPriority"`
	PsPriority string `xml:"PsPriority"`
	VolteHoTargetInd string `xml:"VolteHoTargetInd"`
	FreqPriorityForAnr string `xml:"FreqPriorityForAnr"`
	BackoffTargetInd string `xml:"BackoffTargetInd"`
	MlbInterFreqHoEventType string `xml:"MlbInterFreqHoEventType"`
	MobilityTargetInd string `xml:"MobilityTargetInd"`
	MlbInterFreqEffiRatio string `xml:"MlbInterFreqEffiRatio"`
	SnrBasedUeSelectionMode string `xml:"SnrBasedUeSelectionMode"`
	UlTrafficMlbTargetInd string `xml:"UlTrafficMlbTargetInd"`
	UlTrafficMlbPriority string `xml:"UlTrafficMlbPriority"`
	MlbInterFreqHoA3Offset string `xml:"MlbInterFreqHoA3Offset"`
	IfSrvHoThdRsrpOffset string `xml:"IfSrvHoThdRsrpOffset"`
	IfSrvHoThdRsrqOffset string `xml:"IfSrvHoThdRsrqOffset"`
	MlbFreqUlPriority string `xml:"MlbFreqUlPriority"`
	InterFreqMlbDlPrbOffset string `xml:"InterFreqMlbDlPrbOffset"`
	InterFreqMlbUlPrbOffset string `xml:"InterFreqMlbUlPrbOffset"`
	NcellNumForAnr string `xml:"NcellNumForAnr"`
	MeasPerformanceDemand string `xml:"MeasPerformanceDemand"`
	CtrlMode string `xml:"CtrlMode"`
	DlFreqOffset string `xml:"DlFreqOffset"`
	IfBackoffThdRsrpOffset string `xml:"IfBackoffThdRsrpOffset"`
	IfBackoffThdRsrqOffset string `xml:"IfBackoffThdRsrqOffset"`
	VoLTEQualityIfHoTargetInd string `xml:"VoLTEQualityIfHoTargetInd"`
	IdleMlbeMtcUEReleaseRatio string `xml:"IdleMlbeMtcUEReleaseRatio"`
	InterFreqCioAdjLimitCfgInd string `xml:"InterFreqCioAdjLimitCfgInd"`
	InterFreq4TInd string `xml:"InterFreq4TInd"`
}

