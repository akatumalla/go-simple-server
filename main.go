package main

import (

	"log"

	"net/http"
	"os"


	"github.com/joho/godotenv"

)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Simple Go server deployed using Terraform, Ansible and CircleCI!</h1>"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+"80", mux)
}
