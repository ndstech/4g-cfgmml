package model

import "encoding/xml"

type Eucommcellsectoreqm struct {
	XMLName xml.Name `xml:"eUCommCellSectorEqm"`
	ATTRIBUTES EucommcellsectoreqmAttributes `xml:"attributes"`
}

type EucommcellsectoreqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ENodeBId string `xml:"eNodeBId"`
	SectorEqmId string `xml:"SectorEqmId"`
}

