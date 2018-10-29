package model

import "encoding/xml"

type Cerachcfg struct {
	XMLName xml.Name `xml:"CERACHCFG"`
	ATTRIBUTES CerachcfgAttributes `xml:"attributes"`
}

type CerachcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	ContentionResolutionTimer string `xml:"ContentionResolutionTimer"`
	MaxNumPrbAttempt string `xml:"MaxNumPrbAttempt"`
	PreambleRatio string `xml:"PreambleRatio"`
	PreambleRepetitionNum string `xml:"PreambleRepetitionNum"`
	RandomPreambleRatio string `xml:"RandomPreambleRatio"`
}

