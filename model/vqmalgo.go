package model

import "encoding/xml"

type Vqmalgo struct {
	XMLName xml.Name `xml:"VQMALGO"`
	ATTRIBUTES VqmalgoAttributes `xml:"attributes"`
}

type VqmalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULDelayJitter string `xml:"ULDelayJitter"`
	VqiGoodThd string `xml:"VqiGoodThd"`
	VqiPoorThd string `xml:"VqiPoorThd"`
	VqiBadThd string `xml:"VqiBadThd"`
	AmrWbRadioLowQualRelThd string `xml:"AmrWbRadioLowQualRelThd"`
	AmrNbRadioLowQualRelThd string `xml:"AmrNbRadioLowQualRelThd"`
	AmrSilentPeriodNum string `xml:"AmrSilentPeriodNum"`
	AmrLowQualRelPeriodNum string `xml:"AmrLowQualRelPeriodNum"`
	VqiExcellentThd string `xml:"VqiExcellentThd"`
	AmrNbSilentThd string `xml:"AmrNbSilentThd"`
	AmrWbRcvLowQualRelThd string `xml:"AmrWbRcvLowQualRelThd"`
	AmrNbRcvLowQualRelThd string `xml:"AmrNbRcvLowQualRelThd"`
	AmrWbSilentThd string `xml:"AmrWbSilentThd"`
	VQMAlgoPeriod string `xml:"VQMAlgoPeriod"`
	ObjId string `xml:"objId"`
}

