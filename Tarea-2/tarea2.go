package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"./structures"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	Execute()
}

func Execute() {
	Here:
	RequestServer()
	time.Sleep(55 * time.Second)
	goto Here
}

func RequestServer()  {

	myMsg, err :=  CallSoapXML("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx")
	log.Println("Solicita al SOAP WS")
	if err != nil {
		log.Println("Error al solicitar al Soap Web Service...")
	} else {

		var myFecha = myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha
		var myReferencia = myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia
		SetCache(myFecha, myReferencia)

	}
}


func CallSoapXML(url string) (tcr structures.TipoCambioDiaResponse, err error) {

	myRequest := structures.Envelope{
		Body: structures.Body{
			Content: structures.TipoCambioDiaRequest{},
		},
	}

	rawRequest, _ := xml.Marshal(myRequest)
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

	//	TRANSFORMAR EL XML A STRING Y A ESTRUCTURA
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

//	FUNCIONES HACIA REDIS:
func SetCache(fecha string, referencia string){

	conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
	defer conn.Close()
	
	err = conn.Cmd("HMSET", "tipoCambioDia", "fecha", fecha, "referencia", referencia).Err
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Cmd("EXPIRE", "tipoCambioDia", "60").Err
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Actualiza el cache en redis...")
}
	

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada