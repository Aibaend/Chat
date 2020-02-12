	package main

import (
	"log"
	"net/http"
	"os"
	"github.com/matryer/goblueprints/chapter1/trace"
)



// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	r := newRoom("first")
	r.tracer = trace.New(os.Stdout)

	// Configure websocket route
	http.Handle("/ws", r)
	//http.HandleFunc("create/room", CreateRoom)
	go r.run()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


