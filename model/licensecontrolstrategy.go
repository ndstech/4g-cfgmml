package model

import "encoding/xml"

type Licensecontrolstrategy struct {
	XMLName xml.Name `xml:"LicenseControlStrategy"`
	ATTRIBUTES LicensecontrolstrategyAttributes `xml:"attributes"`
}

type LicensecontrolstrategyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SmoothUpgradeTime string `xml:"SmoothUpgradeTime"`
	ObjId string `xml:"objId"`
}

