package model

import "encoding/xml"

type Camgtcfg struct {
	XMLName xml.Name `xml:"CAMGTCFG"`
	ATTRIBUTES CamgtcfgAttributes `xml:"attributes"`
}

type CamgtcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CarrAggrA2ThdRsrp string `xml:"CarrAggrA2ThdRsrp"`
	CarrAggrA4ThdRsrp string `xml:"CarrAggrA4ThdRsrp"`
	CarrierMgtSwitch string `xml:"CarrierMgtSwitch"`
	ActiveBufferLenThd string `xml:"ActiveBufferLenThd"`
	DeactiveBufferLenThd string `xml:"DeactiveBufferLenThd"`
	ActiveBufferDelayThd string `xml:"ActiveBufferDelayThd"`
	DeactiveThroughputThd string `xml:"DeactiveThroughputThd"`
	CarrAggrA6Offset string `xml:"CarrAggrA6Offset"`
	SCellAgingTime string `xml:"SCellAgingTime"`
	CaA6ReportAmount string `xml:"CaA6ReportAmount"`
	CaA6ReportInterval string `xml:"CaA6ReportInterval"`
	SccDeactCqiThd string `xml:"SccDeactCqiThd"`
	SccCfgInterval string `xml:"SccCfgInterval"`
	CellCaAlgoSwitch string `xml:"CellCaAlgoSwitch"`
	UlCaActiveTimeToTrigger string `xml:"UlCaActiveTimeToTrigger"`
	CaAmbrThd string `xml:"CaAmbrThd"`
	MeasCycleSCell string `xml:"MeasCycleSCell"`
	CellMaxPccNumber string `xml:"CellMaxPccNumber"`
	PccUserNumberOffloadThd string `xml:"PccUserNumberOffloadThd"`
	EnhancedPccAnchorA1ThdRsrp string `xml:"EnhancedPccAnchorA1ThdRsrp"`
	AddedMeasCfg string `xml:"AddedMeasCfg"`
	BlindScellSampleNum string `xml:"BlindScellSampleNum"`
	OptMode string `xml:"OptMode"`
	SleepPeriod string `xml:"SleepPeriod"`
	StatPeriod string `xml:"StatPeriod"`
	RelaxedBHSccDeactCqiThd string `xml:"RelaxedBHSccDeactCqiThd"`
	RelaxedBackhaulCaMaxCcNum string `xml:"RelaxedBackhaulCaMaxCcNum"`
	DisCloudBBCaMaxCcNum string `xml:"DisCloudBBCaMaxCcNum"`
	CaA2TimeToTrigger string `xml:"CaA2TimeToTrigger"`
	CaA6TimeToTrigger string `xml:"CaA6TimeToTrigger"`
	CaA1TimeToTrigger string `xml:"CaA1TimeToTrigger"`
	CaA4TimeToTrigger string `xml:"CaA4TimeToTrigger"`
	SccQuietTime string `xml:"SccQuietTime"`
	FTRelaxedBHCaDLMaxCcNum string `xml:"FTRelaxedBHCaDLMaxCcNum"`
	FddTddCaUlMaxCcNum string `xml:"FddTddCaUlMaxCcNum"`
	FddTddCaDlMaxCcNum string `xml:"FddTddCaDlMaxCcNum"`
	LaaCarrAggrA1ThdRsrp string `xml:"LaaCarrAggrA1ThdRsrp"`
	LaaCarrAggrA2ThdRsrp string `xml:"LaaCarrAggrA2ThdRsrp"`
	FTCA2CCAnchorPolicy string `xml:"FTCA2CCAnchorPolicy"`
	FTCAMultiCCAnchorPolicy string `xml:"FTCAMultiCCAnchorPolicy"`
	UlHeavyTrafficMlbTFCAOptSw string `xml:"UlHeavyTrafficMlbTFCAOptSw"`
	SccReactivationTime string `xml:"SccReactivationTime"`
	FTRelaxedBHCaUlMaxCcNum string `xml:"FTRelaxedBHCaUlMaxCcNum"`
	RelaxedBHCaUlMaxCcNum string `xml:"RelaxedBHCaUlMaxCcNum"`
	CaTrafficDirectionPref string `xml:"CaTrafficDirectionPref"`
	CaMimoPriorityStrategySw string `xml:"CaMimoPriorityStrategySw"`
	EnhancedSccSelA1ThldRsrp string `xml:"EnhancedSccSelA1ThldRsrp"`
	FastScellSelAftScellRmvSw string `xml:"FastScellSelAftScellRmvSw"`
	RsrpOffset string `xml:"RsrpOffset"`
}

