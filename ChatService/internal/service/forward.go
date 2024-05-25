package service

import (
	pb "ChatService/api/grpc/chat.v1"
	chat "ChatService/api/grpc/chatStore.v1"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
)

type message struct {
	from      string
	to        string
	message   string
	timestamp string
}

func forward(ctx context.Context) {
	for {
		msg, _ := waitMsg()
		m := unParse(msg)

		if m.to[0:1] == "#" {
			var n []Node

			usrs, err := client.GetUsersFromGroup(ctx, &chat.Group{Name: m.to})

			if err != nil {
				log.Printf("E: %v", err)
			}

			for _, us := range usrs.Name {
				if m.from != us {
					v, err := list.Search(us)
					if err == nil {
						n = append(n, v)
					}
				}
			}

			for _, node := range n {
				c := *node.Chain

				if err = c.Send(&msg); err != nil {
					log.Printf("E: %v", err)
				}
			}
		} else {
			n, err := list.Search(m.to)

			if err != nil {
				log.Printf("E: %v", err)
			} else {
				c := *n.Chain

				if err = c.Send(&msg); err != nil {
					log.Printf("E: %v", err)
				}
			}
		}

		if m.timestamp == "" {
			if _, err := client.AddMessage(ctx, &chat.ChatMessage{
				From:      m.from,
				To:        m.to,
				Message:   m.message,
				Timestamp: m.timestamp,
			}); err != nil {
				log.Printf("E: %v", err)
			}
		} else {
			if _, err := client.EditMessage(ctx, &chat.ChatMessage{
				From:      m.from,
				Message:   m.message,
				To:        m.to,
				Timestamp: m.timestamp,
			}); err != nil {
				log.Printf("E: %v", err)
			}
		}

		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

func waitMsg() (pb.Message, bool) {
	select {
	case <-done:
		return msg, true
	}
}

func unParse(msg pb.Message) message {
	pH, _ := jwt.Parse(msg.MessageHead, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})
	pB, _ := jwt.Parse(msg.MessageBody, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &Secret, nil
	})

	return message{
		from:      fmt.Sprint(pH.Claims.(jwt.MapClaims)["from"]),
		to:        fmt.Sprint(pH.Claims.(jwt.MapClaims)["to"]),
		message:   fmt.Sprint(pB.Claims.(jwt.MapClaims)["message"]),
		timestamp: fmt.Sprint(pB.Claims.(jwt.MapClaims)["timestamp"]),
	}
}
