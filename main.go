package main

import (
	"api-rede-social/src/config"
	"api-rede-social/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	r := router.Create()
	APIPort := config.APIPort
	fmt.Printf("API listening on port %v", APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", APIPort), r))
}
