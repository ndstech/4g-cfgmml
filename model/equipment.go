package model

import "encoding/xml"

type Equipment struct {
	XMLName xml.Name `xml:"EQUIPMENT"`
	ATTRIBUTES EquipmentAttributes `xml:"attributes"`
}

type EquipmentAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BATIMS string `xml:"BATIMS"`
	PAE string `xml:"PAE"`
	MNTMODE string `xml:"MNTMODE"`
	ST string `xml:"ST"`
	ET string `xml:"ET"`
	MMSETREMARK string `xml:"MMSETREMARK"`
	DVAS string `xml:"DVAS"`
	DVAH string `xml:"DVAH"`
	DVAL string `xml:"DVAL"`
	ODIID string `xml:"ODIID"`
	DCALMSW string `xml:"DCALMSW"`
	EQUIPMENTTY string `xml:"EQUIPMENTTY"`
	PSUFP string `xml:"PSUFP"`
	PROTOCOL string `xml:"PROTOCOL"`
	SMARTTRX string `xml:"SMARTTRX"`
	ESN string `xml:"ESN"`
	SDBBLSD string `xml:"SDBBLSD"`
	APST string `xml:"APST"`
	NPST string `xml:"NPST"`
	SDRCONNSW string `xml:"SDRCONNSW"`
	PWRFAILOMCONNENHSW string `xml:"PWRFAILOMCONNENHSW"`
	OMCONNENHSWCTLTIME string `xml:"OMCONNENHSWCTLTIME"`
}

