package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	for i := 1; i <= 100000000; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
	}
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", home)
	server.HandleFunc("/home", home)
	err := http.ListenAndServe(":1234", server)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("服务启动成功")
	}
}
