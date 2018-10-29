package model

import "encoding/xml"

type Rscgrpalg struct {
	XMLName xml.Name `xml:"RSCGRPALG"`
	ATTRIBUTES RscgrpalgAttributes `xml:"attributes"`
}

type RscgrpalgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	BEAR string `xml:"BEAR"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	RSCGRPID string `xml:"RSCGRPID"`
	TXSSW string `xml:"TXSSW"`
	TXBWASW string `xml:"TXBWASW"`
	RXBWASW string `xml:"RXBWASW"`
	PLRDTH string `xml:"PLRDTH"`
	DDTH string `xml:"DDTH"`
	TXBWAMIN string `xml:"TXBWAMIN"`
	RXBWAMIN string `xml:"RXBWAMIN"`
	TCSW string `xml:"TCSW"`
	PQN string `xml:"PQN"`
	CTTH string `xml:"CTTH"`
	CCTTH string `xml:"CCTTH"`
	TXRSVBW string `xml:"TXRSVBW"`
	RXRSVBW string `xml:"RXRSVBW"`
	DROPPKTNUM string `xml:"DROPPKTNUM"`
}

