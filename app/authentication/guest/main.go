package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/loupax/extreme-microservices/app/authentication/domain"
)

type guestClaims struct {
	*jwt.StandardClaims
}

func main() {
	port := flag.String("port", "8080", "the port the guest authentication endpoint shoult listen to")

	flag.Parse()
	logger := log.New(
		os.Stdout,
		"[guest authentication token generation]",
		log.Ldate|log.LUTC,
	)
	if err := run(*port, logger); err != nil {
		panic(err)
	}

}

func run(port string, logger *log.Logger) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		gc := guestClaims{
			StandardClaims: &jwt.StandardClaims{
				Subject: domain.NewUserID().String(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodNone, gc)
		str, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
		if err != nil {
			logger.Printf("failed signing token %s", err)
		}
		rw.Write([]byte(str))
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
