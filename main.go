package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	name := "_1"
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, name)
}

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	StartServer(port, handle)
}