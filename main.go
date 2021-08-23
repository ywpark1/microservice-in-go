package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(req.Body)

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest) // status code send
			// rw.Write([]byte("Oops"))
			return
		}

		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d) // response
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)

}
