package model

import "encoding/xml"

type Ulocellalgpara struct {
	XMLName xml.Name `xml:"ULOCELLALGPARA"`
	ATTRIBUTES UlocellalgparaAttributes `xml:"attributes"`
}

type UlocellalgparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULOCELLID string `xml:"ULOCELLID"`
	RTWPSIRTGTADJSW string `xml:"RTWPSIRTGTADJSW"`
	SIB7RTWPOPTSW string `xml:"SIB7RTWPOPTSW"`
	ANTIANTENNAIMBALANCESW string `xml:"ANTIANTENNAIMBALANCESW"`
	R99FLOWCTRLSW string `xml:"R99FLOWCTRLSW"`
	CQISAMPLEPERIOD string `xml:"CQISAMPLEPERIOD"`
	CQIUSERNUM string `xml:"CQIUSERNUM"`
	HIGHSPEEDRAKESW string `xml:"HIGHSPEEDRAKESW"`
	DPCHMAXTXPWRRESTRSW string `xml:"DPCHMAXTXPWRRESTRSW"`
	DPCHMAXPWRRTRLOADSTAT string `xml:"DPCHMAXPWRRTRLOADSTAT"`
	SIB7MAXRTWP string `xml:"SIB7MAXRTWP"`
	ULCCHOLPCSW string `xml:"ULCCHOLPCSW"`
	ULCCHOLPCLOADHIGHTHD string `xml:"ULCCHOLPCLOADHIGHTHD"`
	ULCCHOLPCLOADLOWTHD string `xml:"ULCCHOLPCLOADLOWTHD"`
	RADIOQUALITYDPCHPCSW string `xml:"RADIOQUALITYDPCHPCSW"`
	RADIOQUALITYDPCHPCLOADSTAT string `xml:"RADIOQUALITYDPCHPCLOADSTAT"`
	RADIOQUALITYDPCHPCPO string `xml:"RADIOQUALITYDPCHPCPO"`
	HTOHPWSHSW string `xml:"HTOHPWSHSW"`
	SHARINGMARGIN string `xml:"SHARINGMARGIN"`
	MAXSHARINGRATIO string `xml:"MAXSHARINGRATIO"`
	SRBOVERHSDPAOPTSW string `xml:"SRBOVERHSDPAOPTSW"`
	DLPWRLOADTHD string `xml:"DLPWRLOADTHD"`
	SINGLEHARQACTTHRD string `xml:"SINGLEHARQACTTHRD"`
	R99TOHPWSHSW string `xml:"R99TOHPWSHSW"`
	HTOHRTPWSHSW string `xml:"HTOHRTPWSHSW"`
	DPCHTPCPOADJSW string `xml:"DPCHTPCPOADJSW"`
	HSDPACELLTHPTHD string `xml:"HSDPACELLTHPTHD"`
	HSDPAPWRMGNCANCELSW string `xml:"HSDPAPWRMGNCANCELSW"`
	RADIOQUALITYDPCHPCSHOSW string `xml:"RADIOQUALITYDPCHPCSHOSW"`
	ULCARLEVELSCHSW string `xml:"ULCARLEVELSCHSW"`
	RADIOQUALITYDPCHPCENHSW string `xml:"RADIOQUALITYDPCHPCENHSW"`
	RADIOQUALITYFDPCHPCPO string `xml:"RADIOQUALITYFDPCHPCPO"`
	COVERIMPBASEDONADVDEM string `xml:"COVERIMPBASEDONADVDEM"`
	SRBOVERHSDPAMMTOPTSW string `xml:"SRBOVERHSDPAMMTOPTSW"`
	TURBOICENHANCEDSW string `xml:"TURBOICENHANCEDSW"`
	AMRIMPBASEDONPLVASW string `xml:"AMRIMPBASEDONPLVASW"`
}

