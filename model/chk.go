package model

import "encoding/xml"

type Chk struct {
	XMLName xml.Name `xml:"CHK"`
	ATTRIBUTES ChkAttributes `xml:"attributes"`
}

type ChkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENABLEFLAG string `xml:"ENABLEFLAG"`
}

