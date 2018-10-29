package model

import "encoding/xml"

type Interrathocomm struct {
	XMLName xml.Name `xml:"INTERRATHOCOMM"`
	ATTRIBUTES InterrathocommAttributes `xml:"attributes"`
}

type InterrathocommAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	InterRatHoMaxRprtCell string `xml:"InterRatHoMaxRprtCell"`
	InterRatHoRprtAmount string `xml:"InterRatHoRprtAmount"`
	InterRatHoGeranRprtInterval string `xml:"InterRatHoGeranRprtInterval"`
	InterRatHoUtranB1MeasQuan string `xml:"InterRatHoUtranB1MeasQuan"`
	InterRatHoUtranRprtInterval string `xml:"InterRatHoUtranRprtInterval"`
	InterRatHoCdma1xRttB1MeasQuan string `xml:"InterRatHoCdma1xRttB1MeasQuan"`
	InterRatCdma1xRttRprtInterval string `xml:"InterRatCdma1xRttRprtInterval"`
	InterRatHoCdmaHrpdB1MeasQuan string `xml:"InterRatHoCdmaHrpdB1MeasQuan"`
	InterRatCdmaHrpdRprtInterval string `xml:"InterRatCdmaHrpdRprtInterval"`
	InterRatHoA1A2TrigQuan string `xml:"InterRatHoA1A2TrigQuan"`
	InterRatHoEventType string `xml:"InterRatHoEventType"`
	Cdma20001XrttMeasTimer string `xml:"Cdma20001XrttMeasTimer"`
	Cdma20001XrttJudgePnNum string `xml:"Cdma20001XrttJudgePnNum"`
	Cdma2000HrpdFreqSelMode string `xml:"Cdma2000HrpdFreqSelMode"`
	Cdma20001XrttFreqSelMode string `xml:"Cdma20001XrttFreqSelMode"`
	CdmaEcsfbPsConcurrentMode string `xml:"CdmaEcsfbPsConcurrentMode"`
	Cdma20001XrttMeasMode string `xml:"Cdma20001XrttMeasMode"`
	Cdma1XrttSectorIdSelMode string `xml:"Cdma1XrttSectorIdSelMode"`
	CellInfoMaxUtranCellNum string `xml:"CellInfoMaxUtranCellNum"`
	CellInfoMaxGeranCellNum string `xml:"CellInfoMaxGeranCellNum"`
	GeranCellNumForEmcRedirect string `xml:"GeranCellNumForEmcRedirect"`
	UtranCellNumForEmcRedirect string `xml:"UtranCellNumForEmcRedirect"`
	CdmaHrpdSectorIdSelMode string `xml:"CdmaHrpdSectorIdSelMode"`
	IRatBlindRedirPlmnCfgSimSw string `xml:"IRatBlindRedirPlmnCfgSimSw"`
	ObjId string `xml:"objId"`
}

