package main

import (
	_ "expvar"
	"net/http"
	"os"

	"github.com/lindep/golang_exp/micro_service/hello_world/transport"

	"github.com/go-kit/kit/log"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHTTPSServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
