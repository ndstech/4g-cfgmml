package model

import "encoding/xml"

type Clkdetect struct {
	XMLName xml.Name `xml:"ClkDetect"`
	ATTRIBUTES ClkdetectAttributes `xml:"attributes"`
}

type ClkdetectAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ClkAsyncDetectSwitch string `xml:"ClkAsyncDetectSwitch"`
	ClkAsyncInterfRptThld string `xml:"ClkAsyncInterfRptThld"`
	ClkAsyncPrbUsageThld string `xml:"ClkAsyncPrbUsageThld"`
	ClkAsyncSilDetInfDifThld string `xml:"ClkAsyncSilDetInfDifThld"`
	ClkAsyncSilDetInfThld string `xml:"ClkAsyncSilDetInfThld"`
	ObjId string `xml:"objId"`
}

