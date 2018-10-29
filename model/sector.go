package model

import "encoding/xml"

type Sector struct {
	XMLName xml.Name `xml:"SECTOR"`
	ATTRIBUTES SectorAttributes `xml:"attributes"`
}

type SectorAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SECTORID string `xml:"SECTORID"`
	SECNAME string `xml:"SECNAME"`
	LOCATIONNAME string `xml:"LOCATIONNAME"`
	USERLABEL string `xml:"USERLABEL"`
	ANTAZIMUTH string `xml:"ANTAZIMUTH"`
	OLDSECTORID string `xml:"OLDSECTORID"`
	SECTORIDFORCONVERSION string `xml:"SECTORIDFORCONVERSION"`
	SECTORTYPEUMTS string `xml:"SECTORTYPEUMTS"`
	RXANTNUM string `xml:"RXANTNUM"`
	DIVMODE string `xml:"DIVMODE"`
	COVERTYPE string `xml:"COVERTYPE"`
	RFCONNMODE string `xml:"RFCONNMODE"`
	SECTORMODELTE string `xml:"SECTORMODELTE"`
	ANTENNAMODE string `xml:"ANTENNAMODE"`
	SECTORCOMBIND string `xml:"SECTORCOMBIND"`
	OMNIFLAG string `xml:"OMNIFLAG"`
	ORIENTOFMAJORAXIS string `xml:"ORIENTOFMAJORAXIS"`
	CONFIDENCE string `xml:"CONFIDENCE"`
	UNCERTSEMIMAJOR string `xml:"UNCERTSEMIMAJOR"`
	UNCERTSEMIMINOR string `xml:"UNCERTSEMIMINOR"`
	UNCERTALTITUDE string `xml:"UNCERTALTITUDE"`
	SECTORANTENNA string `xml:"SECTORANTENNA"`
}

