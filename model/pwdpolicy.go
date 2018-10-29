package model

import "encoding/xml"

type Pwdpolicy struct {
	XMLName xml.Name `xml:"PWDPOLICY"`
	ATTRIBUTES PwdpolicyAttributes `xml:"attributes"`
}

type PwdpolicyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PWDMINLEN string `xml:"PWDMINLEN"`
	COMPLICACY string `xml:"COMPLICACY"`
	PASSREPLMT string `xml:"PASSREPLMT"`
	MAXPERIOD string `xml:"MAXPERIOD"`
	MINPERIOD string `xml:"MINPERIOD"`
	MAXMISSTIMES string `xml:"MAXMISSTIMES"`
	AUTOUNLOCKTIME string `xml:"AUTOUNLOCKTIME"`
	RESETINTERVAL string `xml:"RESETINTERVAL"`
	PWDEXPRT string `xml:"PWDEXPRT"`
	FirstLoginMustModPWD string `xml:"FirstLoginMustModPWD"`
	MAXREPEATCHARTIMES string `xml:"MAXREPEATCHARTIMES"`
	DICTCHKPWD string `xml:"DICTCHKPWD"`
	MAXCCUN string `xml:"MAXCCUN"`
}

