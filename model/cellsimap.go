package model

import "encoding/xml"

type Cellsimap struct {
	XMLName xml.Name `xml:"CELLSIMAP"`
	ATTRIBUTES CellsimapAttributes `xml:"attributes"`
}

type CellsimapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SiMapSwitch string `xml:"SiMapSwitch"`
	Sib2Period string `xml:"Sib2Period"`
	Sib3Period string `xml:"Sib3Period"`
	Sib4Period string `xml:"Sib4Period"`
	Sib5Period string `xml:"Sib5Period"`
	Sib6Period string `xml:"Sib6Period"`
	Sib7Period string `xml:"Sib7Period"`
	Sib8Period string `xml:"Sib8Period"`
	Sib10Period string `xml:"Sib10Period"`
	Sib11Period string `xml:"Sib11Period"`
	EtwsPnDuration string `xml:"EtwsPnDuration"`
	EtwsSnOverlapPolicy string `xml:"EtwsSnOverlapPolicy"`
	EtwsPnOverlapPolicy string `xml:"EtwsPnOverlapPolicy"`
	Sib12Period string `xml:"Sib12Period"`
	SiTransEcr string `xml:"SiTransEcr"`
	Sib13Period string `xml:"Sib13Period"`
	Sib15Period string `xml:"Sib15Period"`
	Sib16Period string `xml:"Sib16Period"`
	SibTransCtrlSwitch string `xml:"SibTransCtrlSwitch"`
	SiSwitch string `xml:"SiSwitch"`
	SiSchResRatio string `xml:"SiSchResRatio"`
	SibUpdOptSwitch string `xml:"SibUpdOptSwitch"`
	Sib14Period string `xml:"Sib14Period"`
	NbSib1RepetitionNum string `xml:"NbSib1RepetitionNum"`
	NbSib2Period string `xml:"NbSib2Period"`
	NbSib3Period string `xml:"NbSib3Period"`
	NbSib4Period string `xml:"NbSib4Period"`
	NbSib5Period string `xml:"NbSib5Period"`
	NbSib14Period string `xml:"NbSib14Period"`
	NbSib16Period string `xml:"NbSib16Period"`
}

