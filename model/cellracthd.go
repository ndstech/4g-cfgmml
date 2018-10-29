package model

import "encoding/xml"

type Cellracthd struct {
	XMLName xml.Name `xml:"CELLRACTHD"`
	ATTRIBUTES CellracthdAttributes `xml:"attributes"`
}

type CellracthdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	GoldServiceArpThd string `xml:"GoldServiceArpThd"`
	SilverServiceArpThd string `xml:"SilverServiceArpThd"`
	Qci1HoThd string `xml:"Qci1HoThd"`
	Qci2HoThd string `xml:"Qci2HoThd"`
	Qci3HoThd string `xml:"Qci3HoThd"`
	Qci4HoThd string `xml:"Qci4HoThd"`
	NewGoldServiceOffset string `xml:"NewGoldServiceOffset"`
	NewSilverServiceOffset string `xml:"NewSilverServiceOffset"`
	NewCopperServiceOffset string `xml:"NewCopperServiceOffset"`
	Qci1CongThd string `xml:"Qci1CongThd"`
	Qci2CongThd string `xml:"Qci2CongThd"`
	Qci3CongThd string `xml:"Qci3CongThd"`
	Qci4CongThd string `xml:"Qci4CongThd"`
	CongRelOffset string `xml:"CongRelOffset"`
	UlRbHighThd string `xml:"UlRbHighThd"`
	UlRbLowThd string `xml:"UlRbLowThd"`
	AcReservedUserNumber string `xml:"AcReservedUserNumber"`
	VolteReservedNumber string `xml:"VolteReservedNumber"`
	VoltePrefAdmissionTimer string `xml:"VoltePrefAdmissionTimer"`
	CellCapacityMode string `xml:"CellCapacityMode"`
	LoadHoAdmitOffset string `xml:"LoadHoAdmitOffset"`
	VoipOverAdmitOffset string `xml:"VoipOverAdmitOffset"`
	AcUserNumber string `xml:"AcUserNumber"`
	MoSigArpOverride string `xml:"MoSigArpOverride"`
	UlSatisThdforVolteLoadAmrc string `xml:"UlSatisThdforVolteLoadAmrc"`
	CceUsageThd string `xml:"CceUsageThd"`
	PreResNeedTuningFactor string `xml:"PreResNeedTuningFactor"`
	CceAlFailHighThd string `xml:"CceAlFailHighThd"`
	CqiFarThd string `xml:"CqiFarThd"`
	DlExperienceThd string `xml:"DlExperienceThd"`
	RbCongHighThd string `xml:"RbCongHighThd"`
	UlExperienceThd string `xml:"UlExperienceThd"`
	VolteArpOverride string `xml:"VolteArpOverride"`
	CceThdforVolteLoadAmrc string `xml:"CceThdforVolteLoadAmrc"`
	UlRbThdforVolteLoadAmrc string `xml:"UlRbThdforVolteLoadAmrc"`
	CongBearRelPeriod string `xml:"CongBearRelPeriod"`
	PreemptableBearerNum string `xml:"PreemptableBearerNum"`
	Qci65HoThd string `xml:"Qci65HoThd"`
}

