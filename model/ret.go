package model

import "encoding/xml"

type Ret struct {
	XMLName xml.Name `xml:"RET"`
	ATTRIBUTES RetAttributes `xml:"attributes"`
}

type RetAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DEVICENO string `xml:"DEVICENO"`
	DEVICENAME string `xml:"DEVICENAME"`
	CTRLCN string `xml:"CTRLCN"`
	CTRLSRN string `xml:"CTRLSRN"`
	CTRLSN string `xml:"CTRLSN"`
	VENDORCODE string `xml:"VENDORCODE"`
	SERIALNO string `xml:"SERIALNO"`
	RETTYPE string `xml:"RETTYPE"`
	POLARTYPE string `xml:"POLARTYPE"`
	SCENARIO string `xml:"SCENARIO"`
	SUBUNITNUM string `xml:"SUBUNITNUM"`
}

