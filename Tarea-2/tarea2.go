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

func main() {

	//	4.1) API REST

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("GET")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

	//	1) SERVICIO

	myMsg, err := CallSoapXML("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx")

	//	4.2) RESPONDER EN JSON

	if err != nil {
		stringErr := map[string]string{"error": "No se pudo obtener el valor actual del dolar..."}
		json.NewEncoder(w).Encode(stringErr)
	} else {
		json.NewEncoder(w).Encode(myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0])
	}
}

func CallSoapXML(url string) (tcr structures.TipoCambioDiaResponse, err error) {

	myRequest := structures.Envelope{
		Body: structures.Body{
			Content: structures.TipoCambioDiaRequest{},
		},
	}

	rawRequest, _ := xml.Marshal(myRequest)

	//	2) CONSUMIENDO EL SERVICIO

	req, err := http.NewRequest("POST", url, bytes.NewReader(rawRequest))
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

	var xmlResponse structures.TipoCambioDiaResponse
	myMessage := structures.Envelope{
		Body: structures.Body{
			Content: &xmlResponse,
		},
	}

	err2 := xml.Unmarshal([]byte(respBody), &myMessage)
	if err2 != nil {
		return tcr, err2
	}

	return xmlResponse, err
}
