package model

import "encoding/xml"

type Udpping struct {
	XMLName xml.Name `xml:"UDPPING"`
	ATTRIBUTES UdppingAttributes `xml:"attributes"`
}

type UdppingAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TIMEOUTTHD string `xml:"TIMEOUTTHD"`
	TIMEOUTCNT string `xml:"TIMEOUTCNT"`
	DSCP string `xml:"DSCP"`
}

