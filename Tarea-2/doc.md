# Servicio suma

Este servicio devuelve el cambio de dolar actual
### URI

   	<http://10.1.1.33:3002/tipoCambio/tipoCambioDia>

### Metodo:
  

 	 `GET`
    
### Parametros de la URI

#### Requerido:
 
  			Ninguno
  	
#### Parámetros de datos:

  			Ninguno
  
### Respuesta exitosa:
  
	 Entrega el tipo de cambio del dolar y la fecha en formato JSON

  *  **Codigo:** 200 OK
  *  **Contenido:** 
    ```json
    {
        "Fecha": "13/08/2018",
        "Referencia": "7.48156"
    }
    ```
 
### Respuesta de error:

  	No pudo obtener el valor del dolar actual.

  * **Codigo:** 404 NOT FOUND <br />
  * **Contenido:** 
    ```json
    { "error" : "No se pudo obtener el valor actual del dolar..." }
    ```

 ### Llamada simple:

  ```javascript
    $.ajax({
      url: "http://10.1.1.33:3002/tipoCambio/tipoCambioDia",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
#### Nota:

	 Para realizar un test puede utilizar la herramienta [**Postman**](https://www.getpostman.com/).