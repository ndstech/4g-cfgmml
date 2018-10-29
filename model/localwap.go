package model

import "encoding/xml"

type Localwap struct {
	XMLName xml.Name `xml:"LOCALWAP"`
	ATTRIBUTES LocalwapAttributes `xml:"attributes"`
}

type LocalwapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SWITCH string `xml:"SWITCH"`
	SSID string `xml:"SSID"`
	SHOWSSID string `xml:"SHOWSSID"`
	AUTODISABLETIME string `xml:"AUTODISABLETIME"`
	REGIONCODE string `xml:"REGIONCODE"`
	CHANNEL string `xml:"CHANNEL"`
}

