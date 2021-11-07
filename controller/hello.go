package controller

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}
