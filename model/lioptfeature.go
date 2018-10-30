package model

import "encoding/xml"

type Lioptfeature struct {
	XMLName xml.Name `xml:"LIOptFeature"`
	ATTRIBUTES LioptfeatureAttributes `xml:"attributes"`
}

type LioptfeatureAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IOptFeatureID string `xml:"IOptFeatureID"`
	IOptFeatureName string `xml:"IOptFeatureName"`
	ObjId string `xml:"objId"`
}

