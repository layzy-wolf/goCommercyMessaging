package service

import (
	pb "ChatService/api/grpc/chat.v1"
	"context"
	"google.golang.org/grpc/metadata"
	"log"
)

func listen(ctx context.Context, stream pb.Chat_ForwardMessageServer) {
	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("%v", err)
			md, _ := metadata.FromIncomingContext(stream.Context())
			u := getUserFromMD(md)
			if err := forget(u); err != nil {
				log.Printf("E: %v", err)
			}
			return
		}

		updateMsg(in)

		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

func updateMsg(m *pb.Message) {
	msg = *m
	done <- true
	close(done)
	done = make(chan bool)
}
