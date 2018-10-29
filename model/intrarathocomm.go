package model

import "encoding/xml"

type Intrarathocomm struct {
	XMLName xml.Name `xml:"INTRARATHOCOMM"`
	ATTRIBUTES IntrarathocommAttributes `xml:"attributes"`
}

type IntrarathocommAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IntraRatHoMaxRprtCell string `xml:"IntraRatHoMaxRprtCell"`
	IntraRatHoRprtAmount string `xml:"IntraRatHoRprtAmount"`
	IntraFreqHoA3TrigQuan string `xml:"IntraFreqHoA3TrigQuan"`
	IntraFreqHoA3RprtQuan string `xml:"IntraFreqHoA3RprtQuan"`
	IntraFreqHoRprtInterval string `xml:"IntraFreqHoRprtInterval"`
	InterFreqHoA4RprtQuan string `xml:"InterFreqHoA4RprtQuan"`
	InterFreqHoRprtInterval string `xml:"InterFreqHoRprtInterval"`
	InterFreqHoA1A2TrigQuan string `xml:"InterFreqHoA1A2TrigQuan"`
	FreqPriInterFreqHoA1TrigQuan string `xml:"FreqPriInterFreqHoA1TrigQuan"`
	InterFreqHoA4TrigQuan string `xml:"InterFreqHoA4TrigQuan"`
	CovBasedIfHoWaitingTimer string `xml:"CovBasedIfHoWaitingTimer"`
	A3InterFreqHoA1A2TrigQuan string `xml:"A3InterFreqHoA1A2TrigQuan"`
	ObjId string `xml:"objId"`
	FreqPriHoCandidateUeSelPer string `xml:"FreqPriHoCandidateUeSelPer"`
}

