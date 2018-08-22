package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"./structures"
	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("GET")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//	OBTIENE EL CACHE
	myFecha, myReferencia, err := GetCache()
	if err != nil {
		log.Println(err)
		//	LLAMA AL SERVIDOR
		myMsg, err := CallSoapXML("https://www.banguat.gob.gt/variables/ws/TipoCambio.asmx")
		if err != nil {
			stringErr := map[string]string{"error": "No se pudo obtener el valor actual del dolar..."}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(stringErr)
		} else {
			w.WriteHeader(http.StatusOK)
			//	ALMACENA EL CACHE:
			
			json.NewEncoder(w).Encode(myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0])
			SetCache(myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Fecha, myMsg.TipoCambioDiaResult.CambioDolar.VarDolar[0].Referencia)
			log.Println("Solicita al Servidor")
		}
		
	} else {	
		//	IMPLEMENTACION DEL CACHE
		w.WriteHeader(http.StatusOK)
		tipoCambioDia := map[string]string{"Fecha" : myFecha, "Referencia": myReferencia}
		json.NewEncoder(w).Encode(tipoCambioDia)
		//log.Println("Solicita al Cache")
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
}

func GetCache() ( string,  string, error ){
	conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
	defer conn.Close()
	
	fecha, err := conn.Cmd("HGET", "tipoCambioDia", "fecha").Str()
    if err != nil {
    	return "", "", err
	}
	referencia, err := conn.Cmd("HGET", "tipoCambioDia", "referencia").Str()
    if err != nil {
        return "", "", err
	}

	return fecha, referencia, nil
}
	
func GetExpireTime()(int64, error){
	conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
	defer conn.Close()
	
	time, err := conn.Cmd("TTL", "tipoCambioDia").Str()
    if err != nil {
    	return 0, err
	}else{
		return strconv.ParseInt(time, 10, 32)
	}
}

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada