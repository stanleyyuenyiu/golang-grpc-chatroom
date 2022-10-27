package main

import (
	"bufio"
	"context"
	pb "eventpush/protos/event"
	service "eventpush/service/eventpush"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	service.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	go adminClient()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func adminClient() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEventPushClient(conn)

	ctx := context.Background()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := &pb.BoardCastReq{
			Message: scanner.Text(),
		}
		go c.BoardCast(ctx, req)
	}
}
