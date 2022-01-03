package main

import (
	"math/rand"
	"net/http"
	"os"
	"random_fact_generator_api/routes"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().UnixNano()))
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/api", routes.Index)
	http.HandleFunc("/api/docs", routes.Docs)
	http.HandleFunc("/api/fact", routes.API)
	http.ListenAndServe(os.Getenv("ip_address"), nil)
}
