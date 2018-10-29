package model

import "encoding/xml"

type Sctptemplate struct {
	XMLName xml.Name `xml:"SCTPTEMPLATE"`
	ATTRIBUTES SctptemplateAttributes `xml:"attributes"`
}

type SctptemplateAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SCTPTEMPLATEID string `xml:"SCTPTEMPLATEID"`
	RTOMIN string `xml:"RTOMIN"`
	RTOMAX string `xml:"RTOMAX"`
	RTOINIT string `xml:"RTOINIT"`
	RTOALPHA string `xml:"RTOALPHA"`
	RTOBETA string `xml:"RTOBETA"`
	HBINTER string `xml:"HBINTER"`
	MAXASSOCRETR string `xml:"MAXASSOCRETR"`
	MAXPATHRETR string `xml:"MAXPATHRETR"`
	SWITCHBACKFLAG string `xml:"SWITCHBACKFLAG"`
	SWITCHBACHHBNUM string `xml:"SWITCHBACHHBNUM"`
	TSACK string `xml:"TSACK"`
	CHKSUMTYPE string `xml:"CHKSUMTYPE"`
	MAXSTREAM string `xml:"MAXSTREAM"`
}

