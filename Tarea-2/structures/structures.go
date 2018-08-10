package structures

import "encoding/xml"

type Envelope struct {
	XMLName struct{} `xml:"Envelope"`
	Header  Header
	Body    Body
}

type Header struct {
	XMLName  struct{} `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

type Body struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

type TipoCambioDiaRequest struct {
	XMLName xml.Name `xml:"http://www.banguat.gob.gt/variables/ws/ TipoCambioDia"`
}

type TipoCambioDiaResponse struct {
	XMLName             xml.Name            `xml:"http://www.banguat.gob.gt/variables/ws/ TipoCambioDiaResponse"`
	TipoCambioDiaResult TipoCambioDiaResult `xml:"TipoCambioDiaResult"`
}

type TipoCambioDiaResult struct {
	CambioDolar CambioDolar `xml:"CambioDolar"`
	TotalItems  int         `xml:"TotalItems"`
}

type CambioDolar struct {
	VarDolar []VarDolar `xml:"VarDolar"`
}

type VarDolar struct {
	Fecha      string  `xml:"fecha"`
	Referencia float64 `xml:"referencia"`
}
