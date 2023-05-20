package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"pong"}`))
	})
	//pong
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ping"}`))
	})
	// ユーザ操作
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			fmt.Printf("/user :%s & %d \n", r.Method, http.StatusOK)
		case "POST":
			fmt.Printf("/user :%s & %d \n", r.Method, http.StatusOK)
		case "PUT":
			fmt.Printf("/user :%s & %d \n", r.Method, http.StatusOK)
		case "DELETE":
			fmt.Printf("/user :%s & %d \n", r.Method, http.StatusOK)
		default:
			fmt.Printf("/user :%s & %d \n", r.Method, http.StatusBadRequest)
		}

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe :%s", err)
	}
}
