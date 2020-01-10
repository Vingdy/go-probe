package main

import (
	"log"
	"net/http"
	"probe/curl"
)

var (
	Port = "11111"
)

func main() {
	http.HandleFunc("/", curl.GetIps)
	//http.HandleFunc("/", dingding.SendPost)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
