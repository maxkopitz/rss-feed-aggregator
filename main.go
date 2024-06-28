package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
}

func main() {

	filepathRoot := "./app"
	flag.Parse()
	apiCfg := &apiConfig{}
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", apiCfg.handlerReadiness)
	mux.HandleFunc("GET /api/err", apiCfg.handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
