package model

import "encoding/xml"

type Homeascomm struct {
	XMLName xml.Name `xml:"HoMeasComm"`
	ATTRIBUTES HomeascommAttributes `xml:"attributes"`
}

type HomeascommAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GapPatternType string `xml:"GapPatternType"`
	EutranFilterCoeffRsrp string `xml:"EutranFilterCoeffRsrp"`
	EutranFilterCoeffRsrq string `xml:"EutranFilterCoeffRsrq"`
	GeranFilterCoeff string `xml:"GeranFilterCoeff"`
	SMeasureInd string `xml:"SMeasureInd"`
	SpeedDepParaInd string `xml:"SpeedDepParaInd"`
	UtranFilterCoeffRscp string `xml:"UtranFilterCoeffRscp"`
	UtranFilterCoeffEcn0 string `xml:"UtranFilterCoeffEcn0"`
	OptHoPreFailPunishTimer string `xml:"OptHoPreFailPunishTimer"`
	ResHoPreFailPunishTimer string `xml:"ResHoPreFailPunishTimer"`
	NonResHoPreFailPunishTimes string `xml:"NonResHoPreFailPunishTimes"`
	NonResHoPreFailRetryTimes string `xml:"NonResHoPreFailRetryTimes"`
	DedicatedGapPatternType string `xml:"DedicatedGapPatternType"`
	ObjId string `xml:"objId"`
}

