package model

import "encoding/xml"

type Rfu struct {
	XMLName xml.Name `xml:"RFU"`
	ATTRIBUTES RfuAttributes `xml:"attributes"`
}

type RfuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	ADMSTATE string `xml:"ADMSTATE"`
	IFFREQ string `xml:"IFFREQ"`
	ALMPROCSW string `xml:"ALMPROCSW"`
	ALMPROCTHRHLD string `xml:"ALMPROCTHRHLD"`
	ALMTHRHLD string `xml:"ALMTHRHLD"`
	PS string `xml:"PS"`
	RCN string `xml:"RCN"`
	TP string `xml:"TP"`
	RS string `xml:"RS"`
	RXNUM string `xml:"RXNUM"`
	TXNUM string `xml:"TXNUM"`
	RFTXSIGNDETECTSW string `xml:"RFTXSIGNDETECTSW"`
	RFTXSIGNDETECTPERIOD string `xml:"RFTXSIGNDETECTPERIOD"`
	RFTXSIGNDETECTTHLD string `xml:"RFTXSIGNDETECTTHLD"`
	RT string `xml:"RT"`
	IFOFFSET string `xml:"IFOFFSET"`
	RFDS string `xml:"RFDS"`
	FMBWH string `xml:"FMBWH"`
	LCPSW string `xml:"LCPSW"`
	FLAG string `xml:"FLAG"`
	RUSPEC string `xml:"RUSPEC"`
	RFCONNCN2 string `xml:"RFCONNCN2"`
	RFCONNSN2 string `xml:"RFCONNSN2"`
	RFCONNSRN2 string `xml:"RFCONNSRN2"`
	RFCONNTYPE string `xml:"RFCONNTYPE"`
	PAEFFSWITCH string `xml:"PAEFFSWITCH"`
	SCR string `xml:"SCR"`
	RXFREQBANDMUTUALSW string `xml:"RXFREQBANDMUTUALSW"`
	MNTMODE string `xml:"MNTMODE"`
	ST string `xml:"ST"`
	ET string `xml:"ET"`
	MMSETREMARK string `xml:"MMSETREMARK"`
	LEDSW string `xml:"LEDSW"`
	CUSTOMEDRFSPECSW string `xml:"CUSTOMEDRFSPECSW"`
	PSGID string `xml:"PSGID"`
	WIFISW string `xml:"WIFISW"`
}

