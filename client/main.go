package main

import (
	"bufio"
	"context"
	pb "eventpush/protos/event"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	userId  = flag.Int("user", 1, "user Id")
	name    = flag.String("name", "", "user name")
	channel = flag.String("channel", "", "channel")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEventPushClient(conn)

	// Contact the server and print out its response.
	ctx := context.Background()

	go join(ctx, c)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go send(ctx, c, scanner.Text())
	}
}

func join(ctx context.Context, c pb.EventPushClient) {
	ch := *channel
	if ch == "" {
		ch = fmt.Sprintf("CH-%v", *userId)
	}

	req := &pb.JoinReq{
		User: &pb.User{
			Name: *name,
			Id:   int32(*userId),
		},
		Channel: ch,
	}

	stream, err := c.Join(ctx, req)

	if err != nil {
		log.Fatalf("could not StreamListTopK: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("end")
			break
		}
		if err != nil {
			log.Fatalf("%v.join(_) = _, %v", c, err)
		}
		if res.From.Id != req.User.Id {
			log.Printf("[%v:%v]: %v", res.From.Id, res.From.Name, res.Message)
		}

	}
}

func send(ctx context.Context, c pb.EventPushClient, msg string) {
	ch := *channel
	if ch == "" {
		ch = fmt.Sprintf("CH-%v", *userId)
	}

	req := &pb.SendReq{
		To:      int32(*userId),
		Message: msg,
		Channel: ch,
	}

	_, err := c.SendMsg(ctx, req)
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}

}
