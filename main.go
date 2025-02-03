package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Define the structure for a valid JSON response
type Resp struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// Define the structure for error response
type ErrorResp struct {
	Number string `json:"alphabet"`
	Error  bool   `json:"error"`
}

// Define the entrypoint
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number", numberHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
