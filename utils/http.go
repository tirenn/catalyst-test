package utils

import (
	"net/http"
)

func OK(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Created(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
