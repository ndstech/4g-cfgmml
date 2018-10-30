package model

import "encoding/xml"

type Srbcfg struct {
	XMLName xml.Name `xml:"SrbCfg"`
	ATTRIBUTES SrbcfgAttributes `xml:"attributes"`
}

type SrbcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SrbRlcParaAdaptSwitch string `xml:"SrbRlcParaAdaptSwitch"`
	SrbPollTimerAdjustStep string `xml:"SrbPollTimerAdjustStep"`
	SrbPollTimerMaxAdjustValue string `xml:"SrbPollTimerMaxAdjustValue"`
	SrbPollTimerAdjUserNumThd string `xml:"SrbPollTimerAdjUserNumThd"`
	RrcConnRelMaxRetxThd string `xml:"RrcConnRelMaxRetxThd"`
	ObjId string `xml:"objId"`
	SrbPollTimerPreset string `xml:"SrbPollTimerPreset"`
	HOCmdRRCRlsMsgRetxAdpSw string `xml:"HOCmdRRCRlsMsgRetxAdpSw"`
}

