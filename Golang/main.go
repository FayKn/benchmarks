package main

import "net/http"

func main() {
	http.HandleFunc("/api/steamUsers", steamUsers)

	http.ListenAndServe(":8000", nil)
}
