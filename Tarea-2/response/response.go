package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

var body = []byte(`
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <soap:Body>
        <TipoCambioDiaResponse xmlns="http://www.banguat.gob.gt/variables/ws/">
            <TipoCambioDiaResult>
                <CambioDolar>
                    <VarDolar>
                        <fecha>10/08/2018</fecha>
                        <referencia>7.48332</referencia>
                    </VarDolar>
                </CambioDolar>
                <TotalItems>1</TotalItems>
            </TipoCambioDiaResult>
        </TipoCambioDiaResponse>
    </soap:Body>
</soap:Envelope>`)

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

func main() {
	dec := xml.NewDecoder(bytes.NewReader(body))
	var miEnvelope Envelope

	v := Envelope{}
	err := dec.Decode(&v)
	if err != nil {
		panic(err)
	}

	err2 := xml.Unmarshal([]byte(body), &miEnvelope)
	if err2 != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var myMsg = v.Body.TipoCambioDiaResponse

	fmt.Println("TotalItems: ", myMsg.TipoCambioDiaResult.TotalItems)
	fmt.Println("Fecha: ", myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha)
	fmt.Println("Referencia: ", myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia)

	var myMsg2 = miEnvelope.Body.TipoCambioDiaResponse
	soapTotalItems := myMsg2.TipoCambioDiaResult.TotalItems
	soapFecha := myMsg2.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha
	soapReferencia := myMsg2.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia

	fmt.Println("TotalItems: ", soapTotalItems)
	fmt.Println("Fecha: ", soapFecha)
	fmt.Println("Referencia: ", soapReferencia)
}
