package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lexor/twirp-demo/internal/rpc"
)

func main() {
	client := rpc.NewFooProtobufClient("http://localhost:3000", http.DefaultClient)

	resp, err := client.Ping(context.Background(), &rpc.PingReq{
		Foo: "foo",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp.GetBar())
}
