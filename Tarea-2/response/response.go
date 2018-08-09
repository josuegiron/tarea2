package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

var body = []byte(`
        <TipoCambioDiaResponse xmlns="http://www.banguat.gob.gt/variables/ws/">
            <TipoCambioDiaResult>
                <CambioDolar>
                    <VarDolar>
                        <fecha>09/08/2018</fecha>
                        <referencia>7.48541</referencia>
					</VarDolar>
					<VarDolar>
                        <fecha>09/08/2010</fecha>
                        <referencia>7.48541</referencia>
					</VarDolar>
                </CambioDolar>
                <TotalItems>1</TotalItems>
            </TipoCambioDiaResult>
        </TipoCambioDiaResponse>`)

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
	Fecha      string `xml:"fecha"`
	Referencia string `xml:"referencia"`
}

func main() {
	dec := xml.NewDecoder(bytes.NewReader(body))
	var miTipoCambioDiaResponse TipoCambioDiaResponse
	v := TipoCambioDiaResponse{}
	err := dec.Decode(&miTipoCambioDiaResponse)
	if err != nil {
		panic(err)
	}

	err2 := xml.Unmarshal([]byte(body), &v)
	if err2 != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("TotalItems: ", v.TipoCambioDiaResult.TotalItems)
	fmt.Println("Fecha: ", v.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha)
	fmt.Println("Fecha: ", v.TipoCambioDiaResult.CambioDolar.VarDolar[1].Fecha)

	soapTotalItems := miTipoCambioDiaResponse.TipoCambioDiaResult.TotalItems
	soapFecha := miTipoCambioDiaResponse.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha
	soapReferencia := miTipoCambioDiaResponse.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia

	fmt.Println("TotalItems: ", soapTotalItems)
	fmt.Println("Fecha: ", soapFecha)
	fmt.Println("Referencia: ", soapReferencia)
}
