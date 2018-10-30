package model

import "encoding/xml"

type Celluemeascontrolcfg struct {
	XMLName xml.Name `xml:"CellUeMeasControlCfg"`
	ATTRIBUTES CelluemeascontrolcfgAttributes `xml:"attributes"`
}

type CelluemeascontrolcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MaxNonIntraMeasObjNum string `xml:"MaxNonIntraMeasObjNum"`
	MaxEutranFddMeasFreqNum string `xml:"MaxEutranFddMeasFreqNum"`
	MaxEutranTddMeasFreqNum string `xml:"MaxEutranTddMeasFreqNum"`
	MaxUtranFddMeasFreqNum string `xml:"MaxUtranFddMeasFreqNum"`
	MaxUtranTddMeasFreqNum string `xml:"MaxUtranTddMeasFreqNum"`
	MaxGeranMeasFreqNum string `xml:"MaxGeranMeasFreqNum"`
	MaxUeTddMeasFreqNum string `xml:"MaxUeTddMeasFreqNum"`
}

