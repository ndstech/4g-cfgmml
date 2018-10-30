package model

import "encoding/xml"

type Cellmbmscfg struct {
	XMLName xml.Name `xml:"CellMBMSCfg"`
	ATTRIBUTES CellmbmscfgAttributes `xml:"attributes"`
}

type CellmbmscfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MBSFNSyncAreaId string `xml:"MBSFNSyncAreaId"`
	MBMSSwitch string `xml:"MBMSSwitch"`
	ServiceAreaIdList string `xml:"ServiceAreaIdList"`
	MBMSServiceSwitch string `xml:"MBMSServiceSwitch"`
	NCellMbmsCfgSwitch string `xml:"NCellMbmsCfgSwitch"`
}

