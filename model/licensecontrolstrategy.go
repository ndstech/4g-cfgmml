package model

import "encoding/xml"

type Licensecontrolstrategy struct {
	XMLName xml.Name `xml:"LICENSECONTROLSTRATEGY"`
	ATTRIBUTES LicensecontrolstrategyAttributes `xml:"attributes"`
}

type LicensecontrolstrategyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SmoothUpgradeTime string `xml:"SmoothUpgradeTime"`
	ObjId string `xml:"objId"`
}

