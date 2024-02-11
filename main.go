package main

import (
	"curso/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.Router()
	http.ListenAndServe(":8000", nil)
}
