# twirp demo

<details>

<summary>Generate code</summary>

```sh
protoc --twirp_out=. --go_out=. proto/rpc.proto
```

</details>

<details>

<summary>Run service</summary>

```go
go run cmd/demosvc/main.go
```

</details>

<details>

<summary>Make a request (/w cURL)</summary>

```sh
curl \
  --header 'Content-Type: application/json' \
  --data '{"foo": "bar"}' \
  --request "POST" \
  http://localhost:3000/twirp/Foo/Ping
```

</details>
