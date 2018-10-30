package model

import "encoding/xml"

type Cellacbar struct {
	XMLName xml.Name `xml:"CellAcBar"`
	ATTRIBUTES CellacbarAttributes `xml:"attributes"`
}

type CellacbarAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	AcBarringInfoCfgInd string `xml:"AcBarringInfoCfgInd"`
	VoLTEPreferCfgInd string `xml:"VoLTEPreferCfgInd"`
	AcBarringForEmergency string `xml:"AcBarringForEmergency"`
	AcBarringForMoDataCfgInd string `xml:"AcBarringForMoDataCfgInd"`
	AcBarringFactorForCall string `xml:"AcBarringFactorForCall"`
	AcBarTimeForCall string `xml:"AcBarTimeForCall"`
	AC11BarforCall string `xml:"AC11BarforCall"`
	AC12BarforCall string `xml:"AC12BarforCall"`
	AC13BarforCall string `xml:"AC13BarforCall"`
	AC14BarforCall string `xml:"AC14BarforCall"`
	AC15BarforCall string `xml:"AC15BarforCall"`
	AcBarringForMoSigCfgInd string `xml:"AcBarringForMoSigCfgInd"`
	AcBarringFactorForSig string `xml:"AcBarringFactorForSig"`
	AcBarTimeForSig string `xml:"AcBarTimeForSig"`
	AC11BarForSig string `xml:"AC11BarForSig"`
	AC12BarForSig string `xml:"AC12BarForSig"`
	AC13BarForSig string `xml:"AC13BarForSig"`
	AC14BarForSig string `xml:"AC14BarForSig"`
	AC15BarForSig string `xml:"AC15BarForSig"`
	AcBarForMVoiceCfgInd string `xml:"AcBarForMVoiceCfgInd"`
	AcBarForMVideoCfgInd string `xml:"AcBarForMVideoCfgInd"`
	AcBarForCsfbCfgInd string `xml:"AcBarForCsfbCfgInd"`
}

