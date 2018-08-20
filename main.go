package main

import (
	"fmt"
	"net/http"
	"os"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	port := os.Getenv("PORT")
	http.Handle("/", String("Hello World."))
	http.ListenAndServe(":"+port, nil)
}
