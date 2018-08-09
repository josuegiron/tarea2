package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiaguinho/gosoap"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("POST")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

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

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada
