package model

import "encoding/xml"

type Licratio struct {
	XMLName xml.Name `xml:"LicRatio"`
	ATTRIBUTES LicratioAttributes `xml:"attributes"`
}

type LicratioAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UpLicRatio string `xml:"UpLicRatio"`
	TrafficSharingType string `xml:"TrafficSharingType"`
	ObjId string `xml:"objId"`
}

