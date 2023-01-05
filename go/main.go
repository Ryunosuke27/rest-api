package main

import (
	"fmt"
	"log"
	"net/http"
	"go/config"
	"go/utils"
)

func main() {
	config.LoadConfig()
	utils.LoggingSettings(config.Config.LogFile)
	http.HandleFunc("/hello", fooHandlerFunc)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
