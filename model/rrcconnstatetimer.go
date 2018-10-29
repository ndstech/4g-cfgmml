package model

import "encoding/xml"

type Rrcconnstatetimer struct {
	XMLName xml.Name `xml:"RRCCONNSTATETIMER"`
	ATTRIBUTES RrcconnstatetimerAttributes `xml:"attributes"`
}

type RrcconnstatetimerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	T302 string `xml:"T302"`
	T320ForLoadBalance string `xml:"T320ForLoadBalance"`
	T304ForEutran string `xml:"T304ForEutran"`
	T304ForGeran string `xml:"T304ForGeran"`
	T320ForOther string `xml:"T320ForOther"`
	UeInactiveTimer string `xml:"UeInactiveTimer"`
	UlSynTimer string `xml:"UlSynTimer"`
	FilterReptRrcConnReqTimer string `xml:"FilterReptRrcConnReqTimer"`
	UeInactivityTimerDynDrx string `xml:"UeInactivityTimerDynDrx"`
	UlSynTimerDynDrx string `xml:"UlSynTimerDynDrx"`
	UeInactiveTimerQci1 string `xml:"UeInactiveTimerQci1"`
	UeInactTimerDynDrxQci1 string `xml:"UeInactTimerDynDrxQci1"`
	RrcConnRelTimer string `xml:"RrcConnRelTimer"`
	DRXRrcConnRelTimerOffset string `xml:"DRXRrcConnRelTimerOffset"`
	SRLTERrcConnRelTimerOffset string `xml:"SRLTERrcConnRelTimerOffset"`
	ReptRrcReestProtectTimer string `xml:"ReptRrcReestProtectTimer"`
	ObjId string `xml:"objId"`
	NBUeInactiveTimer string `xml:"NBUeInactiveTimer"`
	ExtendedWaitTime string `xml:"ExtendedWaitTime"`
	MTCUeInactiveTimer string `xml:"MTCUeInactiveTimer"`
	PowerPrefIndicationTimer string `xml:"PowerPrefIndicationTimer"`
	OutOfSyncRelTimerOffset string `xml:"OutOfSyncRelTimerOffset"`
}

