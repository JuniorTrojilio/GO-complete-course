package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Printf("Server started in http://localhost:%d\n", config.PORT)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(config.PORT), r))
}
