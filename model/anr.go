package model

import "encoding/xml"

type Anr struct {
	XMLName    xml.Name      `xml:"ANR"`
	ATTRIBUTES AnrAttributes `xml:"attributes"`
}

type AnrAttributes struct {
	XMLName                  xml.Name `xml:"attributes"`
	DelCellThd               string   `xml:"DelCellThd"`
	NcellHoStatNum           string   `xml:"NcellHoStatNum"`
	StatisticPeriod          string   `xml:"StatisticPeriod"`
	FastAnrRprtAmount        string   `xml:"FastAnrRprtAmount"`
	FastAnrRprtInterval      string   `xml:"FastAnrRprtInterval"`
	FastAnrCheckPeriod       string   `xml:"FastAnrCheckPeriod"`
	FastAnrRsrpThd           string   `xml:"FastAnrRsrpThd"`
	FastAnrIntraRatMeasUeNum string   `xml:"FastAnrIntraRatMeasUeNum"`
	FastAnrInterRatMeasUeNum string   `xml:"FastAnrInterRatMeasUeNum"`
	FastAnrIntraRatUeNumThd  string   `xml:"FastAnrIntraRatUeNumThd"`
	FastAnrInterRatUeNumThd  string   `xml:"FastAnrInterRatUeNumThd"`
	OptMode                  string   `xml:"OptMode"`
	StatisticPeriodForNRTDel string   `xml:"StatisticPeriodForNRTDel"`
	StatisticNumForNRTDel    string   `xml:"StatisticNumForNRTDel"`
	ActivePciConflictSwitch  string   `xml:"ActivePciConflictSwitch"`
	StartTime                string   `xml:"StartTime"`
	StopTime                 string   `xml:"StopTime"`
	FastAnrRscpThd           string   `xml:"FastAnrRscpThd"`
	FastAnrRssiThd           string   `xml:"FastAnrRssiThd"`
	FastAnrCdma1xrttPilotThd string   `xml:"FastAnrCdma1xrttPilotThd"`
	FastAnrCdmahrpdPilotThd  string   `xml:"FastAnrCdmahrpdPilotThd"`
	NoHoSetThd               string   `xml:"NoHoSetThd"`
	NcellHoForNRTDelThd      string   `xml:"NcellHoForNRTDelThd"`
	FastAnrMode              string   `xml:"FastAnrMode"`
	EventAnrMode             string   `xml:"EventAnrMode"`
	CaUeChoseMode            string   `xml:"CaUeChoseMode"`
	AnrControlledHoStrategy  string   `xml:"AnrControlledHoStrategy"`
	SmartPreallocationMode   string   `xml:"SmartPreallocationMode"`
	NoHoSetMode              string   `xml:"NoHoSetMode"`
	UtranEventAnrMode        string   `xml:"UtranEventAnrMode"`
	GeranEventAnrMode        string   `xml:"GeranEventAnrMode"`
	EventAnrWithVoipMode     string   `xml:"EventAnrWithVoipMode"`
	UtranEventAnrCgiTimer    string   `xml:"UtranEventAnrCgiTimer"`
	GeranEventAnrCgiTimer    string   `xml:"GeranEventAnrCgiTimer"`
	NrtDelMode               string   `xml:"NrtDelMode"`
	OptModeStrategy          string   `xml:"OptModeStrategy"`
	UtranNcellHoForNRTDelThd string   `xml:"UtranNcellHoForNRTDelThd"`
	GeranNcellHoForNRTDelThd string   `xml:"GeranNcellHoForNRTDelThd"`
	NcellDelPunishPeriod     string   `xml:"NcellDelPunishPeriod"`
	EutranNcellDelPunNum     string   `xml:"EutranNcellDelPunNum"`
	UtranNcellDelPunNum      string   `xml:"UtranNcellDelPunNum"`
	NcellCaThdForNRTDel      string   `xml:"NcellCaThdForNRTDel"`
	HoSucRateForCgiRead      string   `xml:"HoSucRateForCgiRead"`
	StaPeriodForIRatNRTDel   string   `xml:"StaPeriodForIRatNRTDel"`
	StaNumForIRatNRTDel      string   `xml:"StaNumForIRatNRTDel"`
	PeriodForNCellRanking    string   `xml:"PeriodForNCellRanking"`
	StatPeriodCoeff          string   `xml:"StatPeriodCoeff"`
	ObjId                    string   `xml:"objId"`
}
