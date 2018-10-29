package model

import "encoding/xml"

type Enodebresmodealgo struct {
	XMLName xml.Name `xml:"ENODEBRESMODEALGO"`
	ATTRIBUTES EnodebresmodealgoAttributes `xml:"attributes"`
}

type EnodebresmodealgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BbpResAutoConfigSw string `xml:"BbpResAutoConfigSw"`
	TddBbpResDeployAlgo string `xml:"TddBbpResDeployAlgo"`
	FddBbpResDeployAlgo string `xml:"FddBbpResDeployAlgo"`
	BbpResAutoRecfgTimer string `xml:"BbpResAutoRecfgTimer"`
	ENBCellNumCapabilityMode string `xml:"eNBCellNumCapabilityMode"`
	ObjId string `xml:"objId"`
	BbpCpriSharingSwitch string `xml:"BbpCpriSharingSwitch"`
	EnbCellDstDeploySw string `xml:"EnbCellDstDeploySw"`
}

