package helper

import (
	"encoding/json"
	"net/http"
)

func ResponJson(w http.ResponseWriter, code int, payload interface{}) {
	r, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(r)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponJson(w, code, map[string]string{"message": message})
}