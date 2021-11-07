package controller

import (
	"dev-1/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WriteToFile(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Error 1")
		return
	}
	done, b, err := make(chan bool), make(chan []byte), make(chan error)
	go utils.WriteToFile(done, b, err)
	bb, e := ioutil.ReadAll(req.Body)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Error 2")
		return
	}
	b <- bb
	if <-err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Error 3")
		return
	}
	done <- true
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Success")
	return
}
