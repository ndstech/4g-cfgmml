package model

import "encoding/xml"

type Rachcfg struct {
	XMLName xml.Name `xml:"RACHCFG"`
	ATTRIBUTES RachcfgAttributes `xml:"attributes"`
}

type RachcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PwrRampingStep string `xml:"PwrRampingStep"`
	PreambInitRcvTargetPwr string `xml:"PreambInitRcvTargetPwr"`
	MessageSizeGroupA string `xml:"MessageSizeGroupA"`
	PrachFreqOffset string `xml:"PrachFreqOffset"`
	PrachConfigIndexCfgInd string `xml:"PrachConfigIndexCfgInd"`
	PreambleTransMax string `xml:"PreambleTransMax"`
	ContentionResolutionTimer string `xml:"ContentionResolutionTimer"`
	MaxHarqMsg3Tx string `xml:"MaxHarqMsg3Tx"`
	RandomPreambleRatio string `xml:"RandomPreambleRatio"`
	RaPreambleGrpARatio string `xml:"RaPreambleGrpARatio"`
	PrachFreqOffsetStrategy string `xml:"PrachFreqOffsetStrategy"`
	PrachStartTimeCfgInd string `xml:"PrachStartTimeCfgInd"`
	NbCyclicPrefixLength string `xml:"NbCyclicPrefixLength"`
	NbRsrpFirstThreshold string `xml:"NbRsrpFirstThreshold"`
	NbRsrpSecondThreshold string `xml:"NbRsrpSecondThreshold"`
	UeRaInfoReportRatioThd string `xml:"UeRaInfoReportRatioThd"`
}

