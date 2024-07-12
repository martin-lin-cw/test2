package main

import (
	"context"
	"fmt"
	"net"
	"time"

	test1pb "github.com/martin-lin-cw/test2/gen/proto/test1/v1"
	test2pb "github.com/martin-lin-cw/test2/gen/proto/test2/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Test2Service struct {
	test2pb.UnsafeTest2ServiceServer
}

// Hello2 implements test2pb.Test2ServiceServer.
func (t *Test2Service) Hello2(ctx context.Context, req *test2pb.Hello2Request) (*test2pb.Hello2Response, error) {
	fmt.Printf("called at %v\n", time.Now())
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

	go server.Serve(ln)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Errorf("new client: %w", err))
	}

	client := test1pb.NewTest1ServiceClient(conn)
	fmt.Println("client created")

	for i := 0; i < 10; i++ {
		resp, err := client.Hello1(context.Background(), &test1pb.Hello1Request{})
		if err != nil {
			fmt.Println(fmt.Errorf("call test2 Hello2: %w", err))
		}

		fmt.Println(resp.GetResult())
	}

	time.Sleep(time.Second * 10)
	conn.Close()
}
