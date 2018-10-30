package model

import "encoding/xml"

type Cellprbvalmlb struct {
	XMLName xml.Name `xml:"CellPrbValMlb"`
	ATTRIBUTES CellprbvalmlbAttributes `xml:"attributes"`
}

type CellprbvalmlbAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PrbValMlbTrigThd string `xml:"PrbValMlbTrigThd"`
	PrbValMlbAdmitThd string `xml:"PrbValMlbAdmitThd"`
	MlbUeSelectPrbValThd string `xml:"MlbUeSelectPrbValThd"`
	PrbValFilterFactor string `xml:"PrbValFilterFactor"`
}

