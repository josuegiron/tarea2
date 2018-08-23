package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tipoCambio/tipoCambioDia", TipoCambioDia).Methods("GET")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func TipoCambioDia(w http.ResponseWriter, r *http.Request) {

	//	OBTIENE EL CACHE
	myFecha, myReferencia, err := GetCache()
	if err != nil{
		stringErr := map[string]string{"error": "No se pudo obtener el valor actual del dolar..."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(stringErr)
	}else{
		tipoCambioDia := map[string]string{"Fecha" : myFecha, "Referencia": myReferencia}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tipoCambioDia)
	}
}
//	FUNCIONES HACIA REDIS:

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

//	docker build -t josuegiron/api-suma-go .
//	docker run -p 3001:3001 josuegiron/api-suma-go
//	docker tag josuegiron/api-suma-go josuegiron/api-suma-go:version1
//  docker push josuegiron/api-suma-go:version1
//	docker stack deploy -c docker-compose.yml api-suma-go-balanceada