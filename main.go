package main

import (
	"context"
	"net"

	test2pb "github.com/martin-lin-cw/test2/gen/proto/test2/v1"
	"google.golang.org/grpc"
)

type Test2Service struct {
	test2pb.UnsafeTest2ServiceServer
}

// Hello2 implements test2pb.Test2ServiceServer.
func (t *Test2Service) Hello2(ctx context.Context, req *test2pb.Hello2Request) (*test2pb.Hello2Response, error) {
	return &test2pb.Hello2Response{Result: "Hello2"}, nil
}

func NewTest2Service() test2pb.Test2ServiceServer {
	return &Test2Service{}
}

func main() {
	println("Hello, test2")

	server := grpc.NewServer()

	test2Service := NewTest2Service()

	test2pb.RegisterTest2ServiceServer(server, test2Service)

	ln, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	server.Serve(ln)
}
