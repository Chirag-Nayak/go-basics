package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleRootGet)

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good Bye World ! !! ")
		fmt.Fprintf(w, "Good Bye  World ! ! !")
	})

	log.Println("Server started on :9090 ! ! !")
	http.ListenAndServe(":9090", nil)
}

func handleRootGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World ! ! !")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Ooops !"))
		http.Error(w, "Oooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data received is: %s\n,", data)
	fmt.Fprintf(w, "Hello World ! ! !")
}
