package eventpush_service

import (
	"context"
	"errors"
	pb "eventpush/protos/event"
	"log"
	"sync"

	"google.golang.org/grpc"
)

var (
	mu sync.Mutex
)

type EventPushService struct {
	pb.UnimplementedEventPushServer
	channel map[string][]chan pb.EventStream
	users   map[int32]*pb.User
}

func (s *EventPushService) Join(req *pb.JoinReq, stream pb.EventPush_JoinServer) error {
	user := req.GetUser()
	ch := string(req.GetChannel())
	log.Printf("[Join] id: %v, name %v join channel %v", user.Id, user.Name, ch)

	mu.Lock()
	msgCh := make(chan pb.EventStream)
	s.channel[ch] = append(s.channel[ch], msgCh)
	s.users[user.Id] = user
	mu.Unlock()

	tmp := make(chan struct{})
	go func() {
		defer close(tmp)

		for {
			eventStream := <-msgCh
			if err := stream.Send(&pb.EventStream{
				Message: eventStream.Message,
				From:    eventStream.From,
			}); err != nil {
				break
			}
		}
	}()
	<-tmp

	return nil
}

func (s *EventPushService) SendMsg(ctx context.Context, req *pb.SendReq) (*pb.SendReqRes, error) {
	msg := string(req.GetMessage())
	userId := int32(req.GetTo())
	ch := string(req.GetChannel())

	mu.Lock()
	user, exist := s.users[userId]
	streams := s.channel[ch]
	mu.Unlock()

	if !exist {
		return &pb.SendReqRes{Response: false}, errors.New("undefined user")
	}

	log.Printf("[Send] id: %v, name %v send message via channel %v: %v", user.Id, user.Name, ch, msg)

	m := pb.EventStream{
		Message: msg,
		From:    user,
	}

	for _, msgChan := range streams {
		msgChan <- m
	}

	return &pb.SendReqRes{Response: true}, nil
}

func (s *EventPushService) BoardCast(ctx context.Context, req *pb.BoardCastReq) (*pb.SendReqRes, error) {
	msg := string(req.GetMessage())
	ch := string(req.GetChannel())

	m := pb.EventStream{
		Message: msg,
		From: &pb.User{
			Id:   0,
			Name: "Admin",
		},
	}
	log.Printf("[BoardCast] admin boardcast message via channel %v: %v", ch, msg)
	if ch == "" {
		for channel := range s.channel {
			for _, msgChan := range s.channel[channel] {
				msgChan <- m
			}
		}
	} else {
		streams := s.channel[ch]
		for _, msgChan := range streams {
			msgChan <- m
		}
	}

	return &pb.SendReqRes{Response: true}, nil
}

func Register(s grpc.ServiceRegistrar) {
	service := EventPushService{}
	service.channel = make(map[string][]chan pb.EventStream)
	service.users = make(map[int32]*pb.User)
	pb.RegisterEventPushServer(s, &service)
}
