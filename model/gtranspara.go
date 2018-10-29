package model

import "encoding/xml"

type Gtranspara struct {
	XMLName xml.Name `xml:"GTRANSPARA"`
	ATTRIBUTES GtransparaAttributes `xml:"attributes"`
}

type GtransparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LPSCHSW string `xml:"LPSCHSW"`
	RATECFGTYPE string `xml:"RATECFGTYPE"`
	SBTIME string `xml:"SBTIME"`
	ARPAGINGTIME string `xml:"ARPAGINGTIME"`
	MODE string `xml:"MODE"`
	OAMTYPE string `xml:"OAMTYPE"`
	BYPASSSWITCH string `xml:"BYPASSSWITCH"`
	STATUS string `xml:"STATUS"`
	OMCHSWITCHBACK string `xml:"OMCHSWITCHBACK"`
	CPLBSW string `xml:"CPLBSW"`
	CPLISTENINGMODE string `xml:"CPLISTENINGMODE"`
	SCTPRTRPTSW string `xml:"SCTPRTRPTSW"`
	IPERRFRMSW string `xml:"IPERRFRMSW"`
	VLANID string `xml:"VLANID"`
	FORWARDMODE string `xml:"FORWARDMODE"`
	SAMEPORTFORWARDSW string `xml:"SAMEPORTFORWARDSW"`
	ONLYOMIP string `xml:"ONLYOMIP"`
	DNSPRD string `xml:"DNSPRD"`
	EPOBJAUTOCFGSW string `xml:"EPOBJAUTOCFGSW"`
	EPCFGMODE string `xml:"EPCFGMODE"`
	PMAUTOSW string `xml:"PMAUTOSW"`
	DIRECTIPSECPRIMATCHSW string `xml:"DIRECTIPSECPRIMATCHSW"`
	MODELSELECT string `xml:"MODELSELECT"`
}

