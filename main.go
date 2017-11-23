package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	hostPort := net.JoinHostPort("0.0.0.0", os.Getenv("PORT"))
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/healthy", healthCheck)

	log.Println("Starting helloweb app")
	log.Printf("Listening on %s\n", hostPort)

	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
