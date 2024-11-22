package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jebus24/mus/api/router"
	"github.com/jebus24/mus/config"
)

func init() {
	config.Load()
}

func Run() {
	fmt.Printf("Listening...", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
