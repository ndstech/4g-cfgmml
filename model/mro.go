package model

import "encoding/xml"

type Mro struct {
	XMLName xml.Name `xml:"MRO"`
	ATTRIBUTES MroAttributes `xml:"attributes"`
}

type MroAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OptPeriod string `xml:"OptPeriod"`
	NcellOptThd string `xml:"NcellOptThd"`
	StatNumThd string `xml:"StatNumThd"`
	PingpongTimeThd string `xml:"PingpongTimeThd"`
	PingpongRatioThd string `xml:"PingpongRatioThd"`
	CoverAbnormalThd string `xml:"CoverAbnormalThd"`
	ServingRsrpThd string `xml:"ServingRsrpThd"`
	NeighborRsrpThd string `xml:"NeighborRsrpThd"`
	UePingPongNumThd string `xml:"UePingPongNumThd"`
	InterFreqMeasTooLateHoThd string `xml:"InterFreqMeasTooLateHoThd"`
	InterFreqA2RollBackThd string `xml:"InterFreqA2RollBackThd"`
	MroOptMode string `xml:"MroOptMode"`
	UnnecInterRatHoOptThd string `xml:"UnnecInterRatHoOptThd"`
	UnnecInterRatHoRsrpThd string `xml:"UnnecInterRatHoRsrpThd"`
	InterRatA2RsrpMinThd string `xml:"InterRatA2RsrpMinThd"`
	InterRatStatNumThd string `xml:"InterRatStatNumThd"`
	IntraRatTooEarlyHoRatioThd string `xml:"IntraRatTooEarlyHoRatioThd"`
	IntraRatTooLateHoRatioThd string `xml:"IntraRatTooLateHoRatioThd"`
	IntraRatAbnormalRatioThd string `xml:"IntraRatAbnormalRatioThd"`
	InterFreqA2RollBackPeriod string `xml:"InterFreqA2RollBackPeriod"`
	IntraRatHoTooEarlyTimeThd string `xml:"IntraRatHoTooEarlyTimeThd"`
	InterRatAbnormalHoRatioThd string `xml:"InterRatAbnormalHoRatioThd"`
	InterRatMeasTooLateHoThd string `xml:"InterRatMeasTooLateHoThd"`
	UnnecInterRatHoRatioThd string `xml:"UnnecInterRatHoRatioThd"`
	UnnecInterRatHoMeasTime string `xml:"UnnecInterRatHoMeasTime"`
	ObjId string `xml:"objId"`
}

