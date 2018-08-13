**Servicio suma**
---
Este servicio realiza la operacion matemática suma.

* **URI**

   	<http://10.1.1.33:3000/operacion/suma>

* **Metodo:**
  

 	 `GET`
    
*  **Parametros de la URI**

	**Requerido:**
 
	* `sumandoA = [float]`
	* `sumandoB = [float]`

 	**Opcional:**
 	
 			Ninguno
  	
* **Parámetros de datos**

  			Ninguno
  
* **Respuesta exitosa:**
  
	 La respuesta devuelve el total de la suma en formato JSON

  * **Codigo:** 200 OK<br />
    **Contenido:** 
    ```json
    { "total" : 5.1 }
    ```
 
* **Respuesta de error:**

  	Si no se ingresa ningún parámetro o un numero `float`

  * **Codigo:** 400 BAD REQUEST <br />
    **Contenido:** 
    ```json
    { "error" : "No se ingreso un numero como sumando…" }
    ```

 * **Llamada simple:**

  ```javascript
    $.ajax({
      url: "http://10.1.1.33:3000/operacion/suma?sumandoA=2.0&sumandoB=3.1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
* **Nota:**

	 Para realizar un test puede utilizar la herramienta [**Postman**](https://www.getpostman.com/).