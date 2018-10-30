package model

import "encoding/xml"

type Uetimerconst struct {
	XMLName xml.Name `xml:"UeTimerConst"`
	ATTRIBUTES UetimerconstAttributes `xml:"attributes"`
}

type UetimerconstAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	T300 string `xml:"T300"`
	T301 string `xml:"T301"`
	T310 string `xml:"T310"`
	T311 string `xml:"T311"`
	N311 string `xml:"N311"`
	N310 string `xml:"N310"`
	T325 string `xml:"T325"`
	N313 string `xml:"N313"`
	N314 string `xml:"N314"`
	T307 string `xml:"T307"`
	T313 string `xml:"T313"`
	T300ForNb string `xml:"T300ForNb"`
	T310ForNb string `xml:"T310ForNb"`
	T301ForNb string `xml:"T301ForNb"`
	T311ForNb string `xml:"T311ForNb"`
	T300CE string `xml:"T300CE"`
	T301CE string `xml:"T301CE"`
	T310CE string `xml:"T310CE"`
	T311CE string `xml:"T311CE"`
}

