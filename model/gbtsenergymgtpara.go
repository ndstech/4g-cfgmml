package model

import "encoding/xml"

type Gbtsenergymgtpara struct {
	XMLName xml.Name `xml:"GBTSENERGYMGTPARA"`
	ATTRIBUTES GbtsenergymgtparaAttributes `xml:"attributes"`
}

type GbtsenergymgtparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	BAKPWRSAVMETHOD string `xml:"BAKPWRSAVMETHOD"`
	PAADJVOL string `xml:"PAADJVOL"`
}

