package model

import "encoding/xml"

type Userqcipriority struct {
	XMLName xml.Name `xml:"USERQCIPRIORITY"`
	ATTRIBUTES UserqcipriorityAttributes `xml:"attributes"`
}

type UserqcipriorityAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Qci string `xml:"Qci"`
	GoldUlSchPriorityFactor string `xml:"GoldUlSchPriorityFactor"`
	GoldDlSchPriorityFactor string `xml:"GoldDlSchPriorityFactor"`
	SilverUlSchPriorityFactor string `xml:"SilverUlSchPriorityFactor"`
	SilverDlSchPriorityFactor string `xml:"SilverDlSchPriorityFactor"`
	BronzeUlSchPriorityFactor string `xml:"BronzeUlSchPriorityFactor"`
	BronzeDlSchPriorityFactor string `xml:"BronzeDlSchPriorityFactor"`
	ObjId string `xml:"objId"`
}

