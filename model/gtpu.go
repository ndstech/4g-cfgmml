package model

import "encoding/xml"

type Gtpu struct {
	XMLName xml.Name `xml:"GTPU"`
	ATTRIBUTES GtpuAttributes `xml:"attributes"`
}

type GtpuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TIMEOUTTH string `xml:"TIMEOUTTH"`
	TIMEOUTCNT string `xml:"TIMEOUTCNT"`
	DSCP string `xml:"DSCP"`
	STATICCHK string `xml:"STATICCHK"`
	ONLYUPIPCHK string `xml:"ONLYUPIPCHK"`
	ULGTPUSNPADSW string `xml:"ULGTPUSNPADSW"`
	DLGTPUSNCHKSW string `xml:"DLGTPUSNCHKSW"`
}

