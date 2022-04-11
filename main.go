package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Running on port 5002")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5002", r))
}
