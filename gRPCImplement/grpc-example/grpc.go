package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// === Proto-like Definitions ===
type HelloRequest struct {
	Name string
}
type HelloReply struct {
	Message string
}

// === gRPC Interface Definitions ===
// Normally generated from .proto file
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}
type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello, " + req.Name}, nil
}

// === gRPC Registration (Mock) ===
func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "greet.Greeter",
		HandlerType: (*GreeterServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(HelloRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					return srv.(GreeterServer).SayHello(ctx, in)
				},
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "mock",
	}, srv)
}

// === Main ===
func main() {
	// Start server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		s := grpc.NewServer()
		RegisterGreeterServer(s, &greeterServer{})
		reflection.Register(s)

		log.Println("Server started on port :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Let the server start
	time.Sleep(time.Second)

	// Create client connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Call SayHello
	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Println("Client received:", resp.Message)
}

// === Client Interface ===
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/greet.Greeter/SayHello", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
