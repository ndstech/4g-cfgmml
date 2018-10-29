package model

import "encoding/xml"

type Rru struct {
	XMLName xml.Name `xml:"RRU"`
	ATTRIBUTES RruAttributes `xml:"attributes"`
}

type RruAttributes struct {
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
	RN string `xml:"RN"`
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
	REMOTEFLAG string `xml:"REMOTEFLAG"`
	USERLABEL string `xml:"USERLABEL"`
	RFDCPWROFFALMDETECTSW string `xml:"RFDCPWROFFALMDETECTSW"`
	BATTOUTPUNDERVOLTTHLD string `xml:"BATTOUTPUNDERVOLTTHLD"`
	MNTMODE string `xml:"MNTMODE"`
	ST string `xml:"ST"`
	ET string `xml:"ET"`
	MMSETREMARK string `xml:"MMSETREMARK"`
	LEDSW string `xml:"LEDSW"`
	CUSTOMEDRFSPECSW string `xml:"CUSTOMEDRFSPECSW"`
	PSGID string `xml:"PSGID"`
	WIFISW string `xml:"WIFISW"`
	LOCATIONNAME string `xml:"LOCATIONNAME"`
	CIRCUITBREAKERVALUE string `xml:"CIRCUITBREAKERVALUE"`
}

