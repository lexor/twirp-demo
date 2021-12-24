package main

import (
	"net/http"

	"github.com/lexor/twirp-demo/internal/foo"
	"github.com/lexor/twirp-demo/internal/rpc"
)

func main() {
	svc := foo.NewFooService()
	twirpServer := rpc.NewFooServer(svc)

	http.ListenAndServe(":3000", twirpServer)
}
