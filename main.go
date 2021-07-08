package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var LastVersion []byte

func main() {
	savedVersion := os.Getenv("LAST_VERSION")
	LastVersion = []byte(savedVersion)

	key := os.Getenv("WS_SECRET_KEY")

	store := newStore()
	go store.run()

	http.HandleFunc("/release", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != key {
			http.Error(w, "Not Authorized", http.StatusForbidden)
		} else {
			b, _ := ioutil.ReadAll(r.Body)

			LastVersion = b

			go func() {
				store.broadcast <- b
			}()

			fmt.Fprintf(w, "Ok")
		}
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(store, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
