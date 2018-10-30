package model

import "encoding/xml"

type Pcchcfg struct {
	XMLName xml.Name `xml:"PCCHCfg"`
	ATTRIBUTES PcchcfgAttributes `xml:"attributes"`
}

type PcchcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DefaultPagingCycle string `xml:"DefaultPagingCycle"`
	Nb string `xml:"Nb"`
	PagingSentNum string `xml:"PagingSentNum"`
	MaxPagingRecordsNum string `xml:"MaxPagingRecordsNum"`
	PagingStrategy string `xml:"PagingStrategy"`
	EnhPagingCR string `xml:"EnhPagingCR"`
	EnhPagingSwitch string `xml:"EnhPagingSwitch"`
	DefaultPagingCycleForNb string `xml:"DefaultPagingCycleForNb"`
	NbForNbIoT string `xml:"NbForNbIoT"`
	MaxNumRepetitionForPaging string `xml:"MaxNumRepetitionForPaging"`
}

