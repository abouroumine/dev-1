package controller

import (
	"dev-1/utils"
	"fmt"
	"net/http"
)

func ReadFile(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Error 1")
		return
	}
	b, err := make(chan []byte), make(chan error)
	go utils.ReadFile(b, err)
	if <-err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Error 2")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, string(<-b))
	return
}
