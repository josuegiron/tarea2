package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
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

	myMsg, err := CallSoapXML("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx", request)
	if err != nil {
		stringErr := map[string]string{"error": "No se pudo obtener el valor actual del dolar..."}
		json.NewEncoder(w).Encode(stringErr)
	} else {
		json.NewEncoder(w).Encode(myMsg)
	}

	//	4.2) RESPONDER EN JSON

}

func CallSoapXML(url string, body []byte) (tcr structures.TipoCambioDiaResponse, err error) {

	var myMessage structures.Envelope

	//	2) CONSUMIENDO EL SERVICIO

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return tcr, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.banguat.gob.gt/variables/ws/TipoCambioDia")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return tcr, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tcr, err
	}

	//	3) TRANSFORMAR EL XML A STRING Y A ESTRUCTURA

	err2 := xml.Unmarshal([]byte(respBody), &myMessage)
	if err2 != nil {
		return tcr, err2
	}

	return myMessage.Body.TipoCambioDiaResponse, err
}
