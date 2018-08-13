package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./structures"
	"github.com/gorilla/mux"
)

var request = []byte(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <TipoCambioDia xmlns="http://www.banguat.gob.gt/variables/ws/" />
  </soap:Body>
</soap:Envelope>
`)

func main() {

	//	4.1) API REST

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("GET")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

	myMsg := CallSoapXML("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx", request)

	//	4.2) RESPONDER EN JSON

	json.NewEncoder(w).Encode(myMsg)
}

func CallSoapXML(url string, body []byte) structures.TipoCambioDiaResponse {

	//	2) CONSUMIENDO EL SERVICIO

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("error 1: %v", err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.banguat.gob.gt/variables/ws/TipoCambioDia")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error 2: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	var myMessage structures.Envelope

	//	3) TRANSFORMAR EL XML A STRING Y A ESTRUCTURA

	err2 := xml.Unmarshal([]byte(respBody), &myMessage)
	if err2 != nil {
		fmt.Printf("error 3: %v", err2)
	}

	return myMessage.Body.TipoCambioDiaResponse
}
