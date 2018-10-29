package model

import "encoding/xml"

type Scappparacfg struct {
	XMLName xml.Name `xml:"SCAPPPARACFG"`
	ATTRIBUTES ScappparacfgAttributes `xml:"attributes"`
}

type ScappparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AppDnsId string `xml:"AppDnsId"`
	MatchRule string `xml:"MatchRule"`
	AppName string `xml:"AppName"`
	AppIdentType string `xml:"AppIdentType"`
	AppIpv4 string `xml:"AppIpv4"`
	AppCfgTargetInd string `xml:"AppCfgTargetInd"`
	MainAppDnsFlag string `xml:"MainAppDnsFlag"`
	AsParaGroupID string `xml:"AsParaGroupID"`
	MainAppDnsId string `xml:"MainAppDnsId"`
	ServiceEndTimeThd string `xml:"ServiceEndTimeThd"`
}

