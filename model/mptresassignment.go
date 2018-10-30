package model

import "encoding/xml"

type Mptresassignment struct {
	XMLName xml.Name `xml:"MptResAssignment"`
	ATTRIBUTES MptresassignmentAttributes `xml:"attributes"`
}

type MptresassignmentAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	MasterMptAssignmentMode string `xml:"MasterMptAssignmentMode"`
	SlaveMptAssignmentMode string `xml:"SlaveMptAssignmentMode"`
	ObjId string `xml:"objId"`
}

