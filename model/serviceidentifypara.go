package model

import "encoding/xml"

type Serviceidentifypara struct {
	XMLName xml.Name `xml:"ServiceIdentifyPara"`
	ATTRIBUTES ServiceidentifyparaAttributes `xml:"attributes"`
}

type ServiceidentifyparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	HeartbeatPacketLengthThld string `xml:"HeartbeatPacketLengthThld"`
	HeartbeatPacketNumberThld string `xml:"HeartbeatPacketNumberThld"`
	MassFlowBigPacketRateThld string `xml:"MassFlowBigPacketRateThld"`
	MassFlowDuration string `xml:"MassFlowDuration"`
	MassFlowPacketNumberThld string `xml:"MassFlowPacketNumberThld"`
	ObjId string `xml:"objId"`
}

