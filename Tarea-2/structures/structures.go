package structures

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Header struct {
	Header struct{}
}

type Body struct {
	TipoCambioDiaResponse TipoCambioDiaResponse `xml:"http://www.banguat.gob.gt/variables/ws/ TipoCambioDiaResponse"`
}

type TipoCambioDiaResponse struct {
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
	Fecha      string `xml:"fecha"`
	Referencia string `xml:"referencia"`
}
