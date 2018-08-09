package main

import (
	"fmt"

	"github.com/tiaguinho/gosoap"
)

func main() {
	var r TipoCambioDiaResponse

	soap, err := gosoap.SoapClient("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx?WSDL")
	if err != nil {
		fmt.Errorf("Error no definido: %s", err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	err = soap.Call("TipoCambioDia", params)
	if err != nil {
		fmt.Errorf("Error en la llamada SOAP: %s", err)
	}

	soap.Unmarshal(&r)

	fmt.Println("TotalItems: ", r.TipoCambioDiaResult.TotalItems)
	fmt.Println("Fecha: ", r.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha)
	fmt.Println("Referencia: ", r.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia)
}
