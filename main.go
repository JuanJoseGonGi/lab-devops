package main

import "net/http"

func main() {
	handler := http.FileServer(http.Dir("./static"))

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
