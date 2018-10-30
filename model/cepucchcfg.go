package model

import "encoding/xml"

type Cepucchcfg struct {
	XMLName xml.Name `xml:"CePucchCfg"`
	ATTRIBUTES CepucchcfgAttributes `xml:"attributes"`
}

type CepucchcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	PucchRepNum string `xml:"PucchRepNum"`
}

