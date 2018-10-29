package model

import "encoding/xml"

type Celldlpcpdsch struct {
	XMLName xml.Name `xml:"CELLDLPCPDSCH"`
	ATTRIBUTES CelldlpcpdschAttributes `xml:"attributes"`
}

type CelldlpcpdschAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CcuPa string `xml:"CcuPa"`
	CeuPa string `xml:"CeuPa"`
	CellEdgLoadEvalPrd string `xml:"CellEdgLoadEvalPrd"`
	NeighCellEdgLoadThd string `xml:"NeighCellEdgLoadThd"`
	ExceedNCellEdgLoadThdRatio string `xml:"ExceedNCellEdgLoadThdRatio"`
	BFUserPwrBoostPrd string `xml:"BFUserPwrBoostPrd"`
	BFUserAdptPwrA3Offset string `xml:"BFUserAdptPwrA3Offset"`
	BFUserAdptPwrA6Offset string `xml:"BFUserAdptPwrA6Offset"`
}

