package structures

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Header struct {
	Header interface{} `xml:",any"`
}

type Body struct {
	Content interface{} `xml:",any"`
	// TipoCambioDiaResponse TipoCambioDiaResponse `xml:"http://www.banguat.gob.gt/variables/ws/ TipoCambioDiaResponse"`
}

type TipoCambioDiaResponse struct {
	TipoCambioDiaResult TipoCambioDiaResult `xml:"TipoCambioDiaResult"`
}

type TipoCambioDiaResult struct {
	CambioDolar CambioDolar `xml:"CambioDolar"`
	TotalItems  int         `xml:"TotalItems"`
}

type TipoCambioDiaRequest struct {
	XMLName xml.Name `xml:"TipoCambioDiaRequest"`
}

type CambioDolar struct {
	VarDolar []VarDolar `xml:"VarDolar"`
}

type VarDolar struct {
	Fecha      string `xml:"fecha"`
	Referencia string `xml:"referencia"`
}
