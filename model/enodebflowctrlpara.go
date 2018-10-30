package model

import "encoding/xml"

type Enodebflowctrlpara struct {
	XMLName xml.Name `xml:"eNodeBFlowCtrlPara"`
	ATTRIBUTES EnodebflowctrlparaAttributes `xml:"attributes"`
}

type EnodebflowctrlparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AdaptUnsyncTimerLen string `xml:"AdaptUnsyncTimerLen"`
	AdaptUnsyncUserNumThd string `xml:"AdaptUnsyncUserNumThd"`
	DynAcBarPolicyMode string `xml:"DynAcBarPolicyMode"`
	CpuLoadThd string `xml:"CpuLoadThd"`
	ObjId string `xml:"objId"`
	McpttMsgCntSwitch string `xml:"McpttMsgCntSwitch"`
	DynAcBarringTrigAllCellSw string `xml:"DynAcBarringTrigAllCellSw"`
}

