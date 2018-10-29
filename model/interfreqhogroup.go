package model

import "encoding/xml"

type Interfreqhogroup struct {
	XMLName xml.Name `xml:"INTERFREQHOGROUP"`
	ATTRIBUTES InterfreqhogroupAttributes `xml:"attributes"`
}

type InterfreqhogroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterFreqHoGroupId string `xml:"InterFreqHoGroupId"`
	InterFreqHoA1A2Hyst string `xml:"InterFreqHoA1A2Hyst"`
	InterFreqHoA1A2TimeToTrig string `xml:"InterFreqHoA1A2TimeToTrig"`
	InterFreqHoA1ThdRsrp string `xml:"InterFreqHoA1ThdRsrp"`
	InterFreqHoA1ThdRsrq string `xml:"InterFreqHoA1ThdRsrq"`
	InterFreqHoA2ThdRsrp string `xml:"InterFreqHoA2ThdRsrp"`
	InterFreqHoA2ThdRsrq string `xml:"InterFreqHoA2ThdRsrq"`
	InterFreqHoA4Hyst string `xml:"InterFreqHoA4Hyst"`
	InterFreqHoA4ThdRsrp string `xml:"InterFreqHoA4ThdRsrp"`
	InterFreqHoA4ThdRsrq string `xml:"InterFreqHoA4ThdRsrq"`
	InterFreqHoA4TimeToTrig string `xml:"InterFreqHoA4TimeToTrig"`
	InterFreqLoadBasedHoA4ThdRsrp string `xml:"InterFreqLoadBasedHoA4ThdRsrp"`
	InterFreqLoadBasedHoA4ThdRsrq string `xml:"InterFreqLoadBasedHoA4ThdRsrq"`
	FreqPriInterFreqHoA1ThdRsrp string `xml:"FreqPriInterFreqHoA1ThdRsrp"`
	FreqPriInterFreqHoA1ThdRsrq string `xml:"FreqPriInterFreqHoA1ThdRsrq"`
	InterFreqHoA3Offset string `xml:"InterFreqHoA3Offset"`
	A3InterFreqHoA1ThdRsrp string `xml:"A3InterFreqHoA1ThdRsrp"`
	A3InterFreqHoA2ThdRsrp string `xml:"A3InterFreqHoA2ThdRsrp"`
	FreqPriInterFreqHoA2ThdRsrp string `xml:"FreqPriInterFreqHoA2ThdRsrp"`
	FreqPriInterFreqHoA2ThdRsrq string `xml:"FreqPriInterFreqHoA2ThdRsrq"`
	InterFreqMlbA1A2ThdRsrp string `xml:"InterFreqMlbA1A2ThdRsrp"`
	InterFreqHoA5Thd1Rsrp string `xml:"InterFreqHoA5Thd1Rsrp"`
	InterFreqHoA5Thd1Rsrq string `xml:"InterFreqHoA5Thd1Rsrq"`
	SrvReqHoA4ThdRsrp string `xml:"SrvReqHoA4ThdRsrp"`
	SrvReqHoA4ThdRsrq string `xml:"SrvReqHoA4ThdRsrq"`
	MlbInterFreqHoA5Thd1Rsrp string `xml:"MlbInterFreqHoA5Thd1Rsrp"`
	MlbInterFreqHoA5Thd1Rsrq string `xml:"MlbInterFreqHoA5Thd1Rsrq"`
	UlBadQualHoA4Offset string `xml:"UlBadQualHoA4Offset"`
	A3InterFreqHoA1ThdRsrq string `xml:"A3InterFreqHoA1ThdRsrq"`
	A3InterFreqHoA2ThdRsrq string `xml:"A3InterFreqHoA2ThdRsrq"`
	UlHeavyTrafficMlbA4ThdRsrp string `xml:"UlHeavyTrafficMlbA4ThdRsrp"`
	UlHeavyTrafficMlbA4ThdRsrq string `xml:"UlHeavyTrafficMlbA4ThdRsrq"`
	InterFreqHoA1ThdRsrpCE string `xml:"InterFreqHoA1ThdRsrpCE"`
	InterFreqHoA2ThdRsrpCE string `xml:"InterFreqHoA2ThdRsrpCE"`
}

