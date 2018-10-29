package model

import "encoding/xml"

type Ftpcltport struct {
	XMLName xml.Name `xml:"FTPCLTPORT"`
	ATTRIBUTES FtpcltportAttributes `xml:"attributes"`
}

type FtpcltportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SERVERIP string `xml:"SERVERIP"`
	PORT string `xml:"PORT"`
	STARTDATAPORT string `xml:"STARTDATAPORT"`
	ENDDATAPORT string `xml:"ENDDATAPORT"`
}

