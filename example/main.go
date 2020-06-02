package main

import (
	"log"
	"net/http"

	rip "github.com/vikram1565/request-ip"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Your client ip is ", rip.GetClientIP(r))
	})
	log.Println("server started on: http://127.0.1.1:8000")
	serverError := http.ListenAndServe(":8000", nil)
	if serverError != nil {
		log.Println("Failed to start server: ", serverError)
	}
}
