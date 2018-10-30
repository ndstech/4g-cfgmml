package model

import "encoding/xml"

type Enbcellqcirsvdpara struct {
	XMLName xml.Name `xml:"eNBCellQciRsvdPara"`
	ATTRIBUTES EnbcellqcirsvdparaAttributes `xml:"attributes"`
}

type EnbcellqcirsvdparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Qci string `xml:"Qci"`
	RsvdSwPara0 string `xml:"RsvdSwPara0"`
	RsvdSwPara1 string `xml:"RsvdSwPara1"`
	RsvdSwPara2 string `xml:"RsvdSwPara2"`
	RsvdPara2 string `xml:"RsvdPara2"`
	RsvdPara3 string `xml:"RsvdPara3"`
	RsvdPara4 string `xml:"RsvdPara4"`
	RsvdPara5 string `xml:"RsvdPara5"`
	RsvdPara6 string `xml:"RsvdPara6"`
	RsvdPara7 string `xml:"RsvdPara7"`
	RsvdPara8 string `xml:"RsvdPara8"`
	RsvdPara9 string `xml:"RsvdPara9"`
	RsvdPara10 string `xml:"RsvdPara10"`
	RsvdPara11 string `xml:"RsvdPara11"`
	RsvdPara12 string `xml:"RsvdPara12"`
	RsvdPara13 string `xml:"RsvdPara13"`
	RsvdPara14 string `xml:"RsvdPara14"`
	RsvdPara15 string `xml:"RsvdPara15"`
	RsvdPara16 string `xml:"RsvdPara16"`
	RsvdPara17 string `xml:"RsvdPara17"`
	RsvdPara18 string `xml:"RsvdPara18"`
	RsvdPara19 string `xml:"RsvdPara19"`
	RsvdPara20 string `xml:"RsvdPara20"`
	RsvdPara21 string `xml:"RsvdPara21"`
}

