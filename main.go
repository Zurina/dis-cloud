package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/host_id", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		fmt.Fprint(w, id.String())
	})

	http.ListenAndServe(":80", r)
}
