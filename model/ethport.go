package model

import "encoding/xml"

type Ethport struct {
	XMLName xml.Name `xml:"ETHPORT"`
	ATTRIBUTES EthportAttributes `xml:"attributes"`
}

type EthportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	PN string `xml:"PN"`
	PA string `xml:"PA"`
	MTU string `xml:"MTU"`
	SPEED string `xml:"SPEED"`
	DUPLEX string `xml:"DUPLEX"`
	ARPPROXY string `xml:"ARPPROXY"`
	FC string `xml:"FC"`
	FERAT string `xml:"FERAT"`
	FERDT string `xml:"FERDT"`
	RXBCPKTALMOCRTHD string `xml:"RXBCPKTALMOCRTHD"`
	RXBCPKTALMCLRTHD string `xml:"RXBCPKTALMCLRTHD"`
	FIBERSPEEDMATCH string `xml:"FIBERSPEEDMATCH"`
	USERLABEL string `xml:"USERLABEL"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
}

