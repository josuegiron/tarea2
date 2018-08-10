package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"./structures"
	"github.com/gorilla/mux"
)

var request = []byte(`
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <TipoCambioDia xmlns="http://www.banguat.gob.gt/variables/ws/" />
  </soap:Body>
</soap:Envelope>
`)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("POST")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

	CallSoap("http://www.banguat.gob.gt/variables/ws/TipoCambio.asmx?op=TipoCambioDia", request)

}

func CallSoap(url string, body []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error 1: %v", err)
	}
	req.Header.Set("SOAPAction", "http://www.banguat.gob.gt/variables/ws/TipoCambio.asmx?op=TipoCambioDia")
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("charset", "utf-8")

	var httpClient = &http.Client{Timeout: time.Second * 5}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("error 2: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	var myMessage structures.Envelope

	err2 := xml.Unmarshal([]byte(respBody), &myMessage)
	if err2 != nil {
		fmt.Printf("error 3: %v", err2)

	}

	var myMsg = myMessage.Body.TipoCambioDiaResponse
	soapTotalItems := myMsg.TipoCambioDiaResult.TotalItems
	//soapFecha := myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha
	//soapReferencia := myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia

	fmt.Println("TotalItems: ", soapTotalItems)
	//fmt.Println("Fecha: ", soapFecha)
	//fmt.Println("Referencia: ", soapReferencia)

}

func ResponseSoap() {

}

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada
