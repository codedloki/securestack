package main

import (
	"log"
	"net/http"
)


func main(){
	http.Handle("/",http.FileServer(http.Dir("./ui/")))
	log.Println("Server running on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
