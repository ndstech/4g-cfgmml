package model

import "encoding/xml"

type Wtcpproxyalgo struct {
	XMLName xml.Name `xml:"WtcpProxyAlgo"`
	ATTRIBUTES WtcpproxyalgoAttributes `xml:"attributes"`
}

type WtcpproxyalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TcpAccelerationSwitch string `xml:"TcpAccelerationSwitch"`
	TCPStatisticsSwitch string `xml:"TCPStatisticsSwitch"`
	MaxRttStatisticsThd string `xml:"MaxRttStatisticsThd"`
	MaxProxyPktNum string `xml:"MaxProxyPktNum"`
	AckSplitCount string `xml:"AckSplitCount"`
	TcpAsSchWeightFactor string `xml:"TcpAsSchWeightFactor"`
}

