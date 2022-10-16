# Echo

A simple [echo server](https://en.wikipedia.org/wiki/Echo_Protocol) implemented using the [gRPC framework](https://grpc.io/).

## Running locally

Start the server
```
go run echo_server/main.go
```

Run the client
```
go run echo_client/main.go
```

Or build the client and run commands
```
go build -o echo-cli echo_client/main.go
./echo-cli
./echo-cli -input foo
```

## Other commands

Generate code from protocol buffer:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    echo/echo.proto
```

Query with [grpcurl](https://github.com/fullstorydev/grpcurl)
```
cd echo
grpcurl -plaintext -proto ./echo.proto 127.0.0.1:50051 list
grpcurl -plaintext -proto ./echo.proto 127.0.0.1:50051 EchoService.Echo
grpcurl -d '{"input": "bar"}' -plaintext -proto ./echo.proto 127.0.0.1:50051 EchoService.Echo
```

Check with [grpcui](https://github.com/fullstorydev/grpcui)
```
cd echo
grpcui -plaintext -proto echo.proto 127.0.0.1:50051
```

