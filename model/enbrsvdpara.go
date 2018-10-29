package model

import "encoding/xml"

type Enbrsvdpara struct {
	XMLName xml.Name `xml:"ENBRSVDPARA"`
	ATTRIBUTES EnbrsvdparaAttributes `xml:"attributes"`
}

type EnbrsvdparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RsvdSwPara0 string `xml:"RsvdSwPara0"`
	RsvdSwPara1 string `xml:"RsvdSwPara1"`
	RsvdSwPara2 string `xml:"RsvdSwPara2"`
	RsvdSwPara3 string `xml:"RsvdSwPara3"`
	RsvdPara3 string `xml:"RsvdPara3"`
	RsvdPara4 string `xml:"RsvdPara4"`
	RsvdPara5 string `xml:"RsvdPara5"`
	RsvdPara6 string `xml:"RsvdPara6"`
	RsvdPara10 string `xml:"RsvdPara10"`
	RsvdPara11 string `xml:"RsvdPara11"`
	RsvdPara12 string `xml:"RsvdPara12"`
	RsvdPara13 string `xml:"RsvdPara13"`
	RsvdPara20 string `xml:"RsvdPara20"`
	RsvdPara21 string `xml:"RsvdPara21"`
	RsvdPara22 string `xml:"RsvdPara22"`
	RsvdPara23 string `xml:"RsvdPara23"`
	RsvdPara24 string `xml:"RsvdPara24"`
	RsvdPara25 string `xml:"RsvdPara25"`
	RsvdPara26 string `xml:"RsvdPara26"`
	RsvdPara27 string `xml:"RsvdPara27"`
	RsvdPara28 string `xml:"RsvdPara28"`
	RsvdPara29 string `xml:"RsvdPara29"`
	RsvdPara31 string `xml:"RsvdPara31"`
	RsvdPara32 string `xml:"RsvdPara32"`
	RsvdPara33 string `xml:"RsvdPara33"`
	RsvdPara34 string `xml:"RsvdPara34"`
	RsvdPara35 string `xml:"RsvdPara35"`
	RsvdPara36 string `xml:"RsvdPara36"`
	RsvdPara37 string `xml:"RsvdPara37"`
	RsvdPara38 string `xml:"RsvdPara38"`
	RsvdPara39 string `xml:"RsvdPara39"`
	RsvdPara40 string `xml:"RsvdPara40"`
	RsvdPara41 string `xml:"RsvdPara41"`
	RsvdPara42 string `xml:"RsvdPara42"`
	RsvdPara43 string `xml:"RsvdPara43"`
	RsvdPara44 string `xml:"RsvdPara44"`
	RsvdPara45 string `xml:"RsvdPara45"`
	RsvdPara46 string `xml:"RsvdPara46"`
	RsvdPara47 string `xml:"RsvdPara47"`
	RsvdPara48 string `xml:"RsvdPara48"`
	RsvdPara49 string `xml:"RsvdPara49"`
	RsvdPara50 string `xml:"RsvdPara50"`
	RsvdPara51 string `xml:"RsvdPara51"`
	RsvdPara52 string `xml:"RsvdPara52"`
	RsvdPara53 string `xml:"RsvdPara53"`
	ObjId string `xml:"objId"`
}

