package foo

import (
	"context"

	"github.com/lexor/twirp-demo/internal/rpc"
)

type FooService struct{}

func NewFooService() *FooService {
	return &FooService{}
}

func (c *FooService) Ping(ctx context.Context, req *rpc.PingReq) (*rpc.PongRes, error) {
	return &rpc.PongRes{
		Bar: req.GetFoo(),
	}, nil
}
