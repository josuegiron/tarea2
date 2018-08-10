package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"./structures"
	"github.com/gorilla/mux"
	"github.com/tiaguinho/gosoap"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("POST")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

	var miTipoCambioDiaResponse structures.TipoCambioDiaResponse

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

	soap.Unmarshal(&miTipoCambioDiaResponse)

	fmt.Println("TotalItems: ", miTipoCambioDiaResponse.TipoCambioDiaResult.TotalItems)
	fmt.Println("Fecha: ", miTipoCambioDiaResponse.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha)
	fmt.Println("Referencia: ", miTipoCambioDiaResponse.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia)
}

var request = []byte(`
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <TipoCambioDia xmlns="http://www.banguat.gob.gt/variables/ws/" />
  </soap:Body>
</soap:Envelope>
`)

func CallSoap(url string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("SOAPAction", fmt.Sprintf("%s/%s", c.URL, c.Method))

}

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada
