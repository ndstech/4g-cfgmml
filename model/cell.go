package model

import "encoding/xml"

type Cell struct {
	XMLName    xml.Name       `xml:"Cell"`
	ATTRIBUTES CellAttributes `xml:"attributes"`
}

type CellAttributes struct {
	XMLName                    xml.Name `xml:"attributes"`
	LocalCellId                string   `xml:"LocalCellId"`
	CellName                   string   `xml:"CellName"`
	CsgInd                     string   `xml:"CsgInd"`
	UlCyclicPrefix             string   `xml:"UlCyclicPrefix"`
	DlCyclicPrefix             string   `xml:"DlCyclicPrefix"`
	FreqBand                   string   `xml:"FreqBand"`
	UlEarfcnCfgInd             string   `xml:"UlEarfcnCfgInd"`
	DlEarfcn                   string   `xml:"DlEarfcn"`
	UlBandWidth                string   `xml:"UlBandWidth"`
	DlBandWidth                string   `xml:"DlBandWidth"`
	CellId                     string   `xml:"CellId"`
	PhyCellId                  string   `xml:"PhyCellId"`
	AdditionalSpectrumEmission string   `xml:"AdditionalSpectrumEmission"`
	CellActiveState            string   `xml:"CellActiveState"`
	CellAdminState             string   `xml:"CellAdminState"`
	FddTddInd                  string   `xml:"FddTddInd"`
	CellSpecificOffset         string   `xml:"CellSpecificOffset"`
	QoffsetFreq                string   `xml:"QoffsetFreq"`
	RootSequenceIdx            string   `xml:"RootSequenceIdx"`
	HighSpeedFlag              string   `xml:"HighSpeedFlag"`
	PreambleFmt                string   `xml:"PreambleFmt"`
	CellRadius                 string   `xml:"CellRadius"`
	CustomizedBandWidthCfgInd  string   `xml:"CustomizedBandWidthCfgInd"`
	EmergencyAreaIdCfgInd      string   `xml:"EmergencyAreaIdCfgInd"`
	UePowerMaxCfgInd           string   `xml:"UePowerMaxCfgInd"`
	MultiRruCellFlag           string   `xml:"MultiRruCellFlag"`
	CPRICompression            string   `xml:"CPRICompression"`
	AirCellFlag                string   `xml:"AirCellFlag"`
	CrsPortNum                 string   `xml:"CrsPortNum"`
	TxRxMode                   string   `xml:"TxRxMode"`
	UserLabel                  string   `xml:"UserLabel"`
	WorkMode                   string   `xml:"WorkMode"`
	EuCellStandbyMode          string   `xml:"EuCellStandbyMode"`
	SfnMasterCellLabel         string   `xml:"SfnMasterCellLabel"`
	CellSlaveBand              string   `xml:"CellSlaveBand"`
	CnOpSharingGroupId         string   `xml:"CnOpSharingGroupId"`
	IntraFreqRanSharingInd     string   `xml:"IntraFreqRanSharingInd"`
	IntraFreqAnrInd            string   `xml:"IntraFreqAnrInd"`
	CellScaleInd               string   `xml:"CellScaleInd"`
	FreqPriorityForAnr         string   `xml:"FreqPriorityForAnr"`
	CellRadiusStartLocation    string   `xml:"CellRadiusStartLocation"`
	SpecifiedCellFlag          string   `xml:"SpecifiedCellFlag"`
	MultiCellShareMode         string   `xml:"MultiCellShareMode"`
	ObjId                      string   `xml:"objId"`
	NbCellFlag                 string   `xml:"NbCellFlag"`
	SubframeAssignment         string   `xml:"SubframeAssignment"`
	SpecialSubframePatterns    string   `xml:"SpecialSubframePatterns"`
	CrsPortMap                 string   `xml:"CrsPortMap"`
	CsiRsPeriod                string   `xml:"CsiRsPeriod"`
	UlEarfcn                   string   `xml:"UlEarfcn"`
	UePowerMax                 string   `xml:"UePowerMax"`
}
