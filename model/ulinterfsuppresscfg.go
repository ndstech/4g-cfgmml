package model

import "encoding/xml"

type Ulinterfsuppresscfg struct {
	XMLName xml.Name `xml:"UlInterfSuppressCfg"`
	ATTRIBUTES UlinterfsuppresscfgAttributes `xml:"attributes"`
}

type UlinterfsuppresscfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlInterfsuppressionThd string `xml:"UlInterfsuppressionThd"`
	P0UePuschOffset string `xml:"P0UePuschOffset"`
	RASignalMCSOffset string `xml:"RASignalMCSOffset"`
	P0UePucchOffset string `xml:"P0UePucchOffset"`
	DeltaMSG2Offset string `xml:"DeltaMSG2Offset"`
	StrongInfUserThdRatio string `xml:"StrongInfUserThdRatio"`
	ULInfStatisticPeriod string `xml:"ULInfStatisticPeriod"`
	ULInfStatisticPeriodNum string `xml:"ULInfStatisticPeriodNum"`
	VolteMCSOffset string `xml:"VolteMCSOffset"`
	VoltePucchPcTarSinrOffset string `xml:"VoltePucchPcTarSinrOffset"`
	VoltePuschPsdCtrlTarget string `xml:"VoltePuschPsdCtrlTarget"`
	RemoteInfULEnhanceSw string `xml:"RemoteInfULEnhanceSw"`
	PuschEnhDeltaSinrThd string `xml:"PuschEnhDeltaSinrThd"`
	RASignalPcSwitch string `xml:"RASignalPcSwitch"`
	UlInterPwrDiffThd string `xml:"UlInterPwrDiffThd"`
}

