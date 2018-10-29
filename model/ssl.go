package model

import "encoding/xml"

type Ssl struct {
	XMLName    xml.Name      `xml:"SSL"`
	ATTRIBUTES SslAttributes `xml:"attributes"`
}

type SslAttributes struct {
	XMLName        xml.Name `xml:"attributes"`
	CONNTYPE       string   `xml:"CONNTYPE"`
	AUTHMODE       string   `xml:"AUTHMODE"`
	RENEGO         string   `xml:"RENEGO"`
	RENEGOINTERVAL string   `xml:"RENEGOINTERVAL"`
	LOWESTCSLEVEL  string   `xml:"LOWESTCSLEVEL"`
	VERSION        string   `xml:"VERSION"`
}
