package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		fmt.Fprintln(w, "# HELP go_app_status Az alkalmazas futasi allapota")
		fmt.Fprintln(w, "# TYPE go_app_status gauge")
		fmt.Fprintln(w, "go_app_status 1")
		fmt.Fprintln(w, "# HELP go_app_random_value Egy tetszoleges statikus ertek")
		fmt.Fprintln(w, "# TYPE go_app_random_value gauge")
		fmt.Fprintln(w, "go_app_random_value 42")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go App fut! A metrikakat itt talalod: /metrics")
	})

	fmt.Println("Alkalmazás indul a 8080-as porton...")
	http.ListenAndServe(":8080", nil)
}
