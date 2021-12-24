#!/bin/sh
protoc --twirp_out=. --go_out=. proto/rpc.proto