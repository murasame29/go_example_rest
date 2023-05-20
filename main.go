package main

import (
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
	http.HandleFunc("/user")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe :%s", err)
	}
}
