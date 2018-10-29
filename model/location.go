package model

import "encoding/xml"

type Location struct {
	XMLName xml.Name `xml:"LOCATION"`
	ATTRIBUTES LocationAttributes `xml:"attributes"`
}

type LocationAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ADDRESS string `xml:"ADDRESS"`
	ALTITUDE string `xml:"ALTITUDE"`
	CITY string `xml:"CITY"`
	CONTACT string `xml:"CONTACT"`
	GCDF string `xml:"GCDF"`
	LATITUDEDEGFORMAT string `xml:"LATITUDEDEGFORMAT"`
	LOCATIONID string `xml:"LOCATIONID"`
	LOCATIONNAME string `xml:"LOCATIONNAME"`
	LONGITUDEDEGFORMAT string `xml:"LONGITUDEDEGFORMAT"`
	OFFICE string `xml:"OFFICE"`
	RANGE string `xml:"RANGE"`
	REGION string `xml:"REGION"`
	TELEPHONE string `xml:"TELEPHONE"`
	USERLABEL string `xml:"USERLABEL"`
	PRECISE string `xml:"PRECISE"`
	MODE string `xml:"MODE"`
	LOCATIONTYPE string `xml:"LOCATIONTYPE"`
}

