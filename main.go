package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("WS_SECRET_KEY")
	origin := os.Getenv("WS_ORIGIN")

	store := newStore()
	go store.run()

	http.HandleFunc("/release", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != key {
			http.Error(w, "Not Authorized", http.StatusForbidden)
		} else {
			b, _ := ioutil.ReadAll(r.Body)

			go func() {
				store.broadcast <- b
			}()

			fmt.Fprintf(w, "Ok")
		}
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(store, w, r, origin)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
