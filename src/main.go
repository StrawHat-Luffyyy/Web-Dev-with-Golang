package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)

}
