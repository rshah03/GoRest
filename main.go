package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
}

type cartItemRequestBody struct {
	Item string
}

func get(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`
	{
		"message": "GET called"	
	}`))
}
func post(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`
	{
		"message": "POST called"	
	}`))
}
func put(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write([]byte(`
	{
		"message": "PUT called"	
	}`))
}
func delete(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`
	{
		"message": "DELETE called"	
	}`))
}
func notFound(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte(`
	{
		"message": "Not found"	
	}`))
}

func main() {
	fmt.Println("Go Rest!")
	req := mux.NewRouter()
	api := req.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":8080", req))
}
