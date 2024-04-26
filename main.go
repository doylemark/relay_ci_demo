package main

import (
	"net/http"
	"time"
)

func hello_world(w http.ResponseWriter, req *http.Request) {
    current_time := time.Now().Format(time.Stamp)
    quote := "Hello World!"
    w.Write([]byte(current_time + ": " + quote + "\n"))
}

func main() {
	http.HandleFunc("/", hello_world)
	http.ListenAndServe(":8080", nil)
}
