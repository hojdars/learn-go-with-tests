package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(buffer io.Writer, name string) {
	fmt.Fprintf(buffer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world!")
}

func main() {
	Greet(os.Stdout, "Hroch")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
