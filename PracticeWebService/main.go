package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		bodyData, _ := ioutil.ReadAll(r.Body)
		w.Write(bodyData)
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
