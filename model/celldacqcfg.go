package model

import "encoding/xml"

type Celldacqcfg struct {
	XMLName xml.Name `xml:"CellDacqCfg"`
	ATTRIBUTES CelldacqcfgAttributes `xml:"attributes"`
}

type CelldacqcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlAmbrLimitFirstRank string `xml:"DlAmbrLimitFirstRank"`
	DlAmbrLimitSecRank string `xml:"DlAmbrLimitSecRank"`
	UlAmbrLimitFirstRank string `xml:"UlAmbrLimitFirstRank"`
	UlAmbrLimitSecRank string `xml:"UlAmbrLimitSecRank"`
	DlPrbUsageFirstRank string `xml:"DlPrbUsageFirstRank"`
	DlPrbUsageSecRank string `xml:"DlPrbUsageSecRank"`
	UlPrbUsageFirstRank string `xml:"UlPrbUsageFirstRank"`
	UlPrbUsageSecRank string `xml:"UlPrbUsageSecRank"`
	DacqExecutionDuration string `xml:"DacqExecutionDuration"`
	DacqCongMonDur string `xml:"DacqCongMonDur"`
	AmbrLimitSchUserNum string `xml:"AmbrLimitSchUserNum"`
	AmbrProtectSchUserNum string `xml:"AmbrProtectSchUserNum"`
	DlAmbrLimitThirdRank string `xml:"DlAmbrLimitThirdRank"`
	DlPrbUsageThirdRank string `xml:"DlPrbUsageThirdRank"`
	UlAmbrLimitThirdRank string `xml:"UlAmbrLimitThirdRank"`
	UlPrbUsageThirdRank string `xml:"UlPrbUsageThirdRank"`
}

