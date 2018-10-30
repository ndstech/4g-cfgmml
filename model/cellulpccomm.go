package model

import "encoding/xml"

type Cellulpccomm struct {
	XMLName xml.Name `xml:"CellUlpcComm"`
	ATTRIBUTES CellulpccommAttributes `xml:"attributes"`
}

type CellulpccommAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PassLossCoeff string `xml:"PassLossCoeff"`
	P0NominalPUCCH string `xml:"P0NominalPUCCH"`
	P0NominalPUSCH string `xml:"P0NominalPUSCH"`
	DeltaFPUCCHFormat1 string `xml:"DeltaFPUCCHFormat1"`
	DeltaFPUCCHFormat1b string `xml:"DeltaFPUCCHFormat1b"`
	DeltaFPUCCHFormat2 string `xml:"DeltaFPUCCHFormat2"`
	DeltaFPUCCHFormat2a string `xml:"DeltaFPUCCHFormat2a"`
	DeltaFPUCCHFormat2b string `xml:"DeltaFPUCCHFormat2b"`
	DeltaPreambleMsg3 string `xml:"DeltaPreambleMsg3"`
	DeltaMsg2 string `xml:"DeltaMsg2"`
	DeltaFPUCCHFormat1bcs string `xml:"DeltaFPUCCHFormat1bcs"`
	DeltaFPUCCHFormat3 string `xml:"DeltaFPUCCHFormat3"`
}

