package model

import "encoding/xml"

type Tacalg struct {
	XMLName xml.Name `xml:"TACALG"`
	ATTRIBUTES TacalgAttributes `xml:"attributes"`
}

type TacalgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RSCGRPULCACSWITCH string `xml:"RSCGRPULCACSWITCH"`
	RSCGRPDLCACSWITCH string `xml:"RSCGRPDLCACSWITCH"`
	TRMULHOCACTH string `xml:"TRMULHOCACTH"`
	TRMDLHOCACTH string `xml:"TRMDLHOCACTH"`
	TRMULGOLDCACTH string `xml:"TRMULGOLDCACTH"`
	TRMDLGOLDCACTH string `xml:"TRMDLGOLDCACTH"`
	TRMULSILVERCACTH string `xml:"TRMULSILVERCACTH"`
	TRMDLSILVERCACTH string `xml:"TRMDLSILVERCACTH"`
	TRMULBRONZECACTH string `xml:"TRMULBRONZECACTH"`
	TRMDLBRONZECACTH string `xml:"TRMDLBRONZECACTH"`
	TRMULGBRCACTH string `xml:"TRMULGBRCACTH"`
	TRMDLGBRCACTH string `xml:"TRMDLGBRCACTH"`
	TRMULPRESW string `xml:"TRMULPRESW"`
	TRMDLPRESW string `xml:"TRMDLPRESW"`
	PORTULOBSW string `xml:"PORTULOBSW"`
	PORTDLOBSW string `xml:"PORTDLOBSW"`
	PORTULCACSW string `xml:"PORTULCACSW"`
	PORTDLCACSW string `xml:"PORTDLCACSW"`
	EMCTACPSW string `xml:"EMCTACPSW"`
}

