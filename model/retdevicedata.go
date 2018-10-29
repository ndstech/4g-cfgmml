package model

import "encoding/xml"

type Retdevicedata struct {
	XMLName xml.Name `xml:"RETDEVICEDATA"`
	ATTRIBUTES RetdevicedataAttributes `xml:"attributes"`
}

type RetdevicedataAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DEVICENO string `xml:"DEVICENO"`
	SUBUNITNO string `xml:"SUBUNITNO"`
	BEARING string `xml:"BEARING"`
	MODELNO string `xml:"MODELNO"`
	BSID string `xml:"BSID"`
	BEAMWIDTH1 string `xml:"BEAMWIDTH1"`
	BEAMWIDTH2 string `xml:"BEAMWIDTH2"`
	BEAMWIDTH3 string `xml:"BEAMWIDTH3"`
	BEAMWIDTH4 string `xml:"BEAMWIDTH4"`
	GAIN1 string `xml:"GAIN1"`
	GAIN2 string `xml:"GAIN2"`
	GAIN3 string `xml:"GAIN3"`
	GAIN4 string `xml:"GAIN4"`
	DATE string `xml:"DATE"`
	TILT string `xml:"TILT"`
	INSTALLERID string `xml:"INSTALLERID"`
	BAND1 string `xml:"BAND1"`
	BAND2 string `xml:"BAND2"`
	BAND3 string `xml:"BAND3"`
	BAND4 string `xml:"BAND4"`
	SECTORID string `xml:"SECTORID"`
	SERIALNO string `xml:"SERIALNO"`
}

