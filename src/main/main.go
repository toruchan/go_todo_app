package main

import (
    "controller"
	"net/http"
)

func main() {
    http.HandleFunc("/", controller.IndexGET)
	http.HandleFunc("/addTask", controller.AddTask)
    http.ListenAndServe(":8080", nil)
	
}
