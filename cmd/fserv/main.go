package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var portInt = flag.Int("port", 8000, "port number")
var port = strconv.Itoa(*portInt)
var dir = flag.String("d", "./", "server root")

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	log.Println("Starting server on:" + port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)

	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}
