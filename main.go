package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/* Solo usamos esta funcion para crear de forma segura una llave para el archivo .env
func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
*/

func main() {
	config.Carregar()
	fmt.Printf("Corriendo en conexion de base de datos %s\n", config.StringConexaoBanco)
	fmt.Printf("Corriendo API en el puerto %d...\n", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
