package handle

import (
	"fmt"
	"net/http"
)

func UserHandle(w http.ResponseWriter, r *http.Request) {
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

}
