package main

import (
	"fmt"
	"github.com/moesif/moesifmiddleware-go"
	"log"
	"net/http"
	"os"
)

var moesifOptions = map[string]interface{} {
	"Application_Id": os.Getenv("MOESIF_APPLICATION_ID"),
	"Log_Body": true,
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := "_1"
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, name)
}

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.Handle("/", moesifmiddleware.MoesifMiddleware(http.HandlerFunc(handle), moesifOptions))
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