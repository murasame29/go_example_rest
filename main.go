package main

import (
	"go_example_crud/handle"
	"log"
	"net/http"
)

func main() {
	//ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"pong"}`))
	})
	//pong
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ping"}`))
	})

	// ユーザ操作
	http.HandleFunc("/users/", handle.UsersHandle)

	log.Println("opening localhost:8080 server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe :%s", err)
	}
}
