package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/loupax/extreme-microservices/app/advancewars/domain"
)

type battleRequest struct {
	Terrain domain.Terrain `json:"terrain"`
	Units   domain.Units   `json:"units"`
}

func main() {
	port := flag.String("port", "8081", "the port the guest authentication endpoint shoult listen to")

	flag.Parse()
	logger := log.New(
		os.Stdout,
		"[battle initialization]",
		log.Ldate|log.LUTC,
	)
	if err := run(*port, logger); err != nil {
		panic(err)
	}
}

func run(port string, logger *log.Logger) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		g := battleRequest{}
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			logger.Printf("failed decoding request %s", err)
		}
		logger.Println(r.Header.Get("Authenticated-User"))
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(fmt.Sprintf(`{"map_id":"%s"}`, domain.NewMapID().String())))
	})

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
