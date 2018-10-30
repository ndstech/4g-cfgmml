package model

import "encoding/xml"

type Eucellsectoreqm struct {
	XMLName xml.Name `xml:"eUCellSectorEqm"`
	ATTRIBUTES EucellsectoreqmAttributes `xml:"attributes"`
}

type EucellsectoreqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SectorEqmId string `xml:"SectorEqmId"`
	ReferenceSignalPwr string `xml:"ReferenceSignalPwr"`
	BaseBandEqmId string `xml:"BaseBandEqmId"`
	ReferenceSignalPwrMargin string `xml:"ReferenceSignalPwrMargin"`
	SectorCpriCompression string `xml:"SectorCpriCompression"`
	AutoCfgFlag string `xml:"AutoCfgFlag"`
	WeightNO string `xml:"WeightNO"`
	VisualCellId string `xml:"VisualCellId"`
	PrbIdList string `xml:"PrbIdList"`
	SectorEqmCombineGrpId string `xml:"SectorEqmCombineGrpId"`
	CellBeamMode string `xml:"CellBeamMode"`
	BeamId string `xml:"BeamId"`
}

