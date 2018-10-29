package model

import "encoding/xml"

type Enodebusparacfg struct {
	XMLName xml.Name `xml:"ENODEBUSPARACFG"`
	ATTRIBUTES EnodebusparacfgAttributes `xml:"attributes"`
}

type EnodebusparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UsStrategy string `xml:"UsStrategy"`
	UsSPIDConfig string `xml:"UsSPIDConfig"`
	UsUeInactiveTimer string `xml:"UsUeInactiveTimer"`
	IPDetectInterval string `xml:"IPDetectInterval"`
	UsUeDlPriorityFactor string `xml:"UsUeDlPriorityFactor"`
	UsLPktUlschPriFactor string `xml:"UsLPktUlschPriFactor"`
	UsSPktUlschPriFactor string `xml:"UsSPktUlschPriFactor"`
	UsUESTMSIKeepAliveTimer string `xml:"UsUESTMSIKeepAliveTimer"`
	BigPackageIdentify string `xml:"BigPackageIdentify"`
	ObjId string `xml:"objId"`
	UlDlUsUserDetectPeriod string `xml:"UlDlUsUserDetectPeriod"`
}

