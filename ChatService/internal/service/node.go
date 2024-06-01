package service

import (
	pb "ChatService/api/grpc/chat.v1"
	"errors"
	"fmt"
)

type List interface {
	Insert(string, *pb.Chat_ForwardMessageServer)
	Search(string) (Node, error)
	Delete(string) error
}

type Node struct {
	Data  string
	Next  *Node
	Chain *pb.Chat_ForwardMessageServer
}

// LinkedList implementation /*
type LinkedList struct {
	Head   *Node
	Length int
}

/*
Insert data to end of linked list
*/
func (l *LinkedList) Insert(data string, stream *pb.Chat_ForwardMessageServer) {
	temp1 := new(Node)
	temp1.Data = data
	temp1.Chain = stream

	if l.Head == nil {
		l.Head = temp1
	} else {
		temp2 := l.Head
		for temp2.Next != nil {
			temp2 = temp2.Next
		}
		temp2.Next = temp1
	}
	l.Length += 1
}

/*
Search data in linked list
*/
func (l *LinkedList) Search(data string) (Node, error) {
	temp := l.Head

	for temp != nil {
		if temp.Data == data {
			return *temp, nil
		}
		temp = temp.Next
	}
	return Node{}, fmt.Errorf("non element")
}

/*
Delete elements from linked list by them data
*/
func (l *LinkedList) Delete(data string) error {
	temp1 := l.Head
	temp2 := temp1.Next

	for temp2 != nil {
		if temp2.Data == data {
			temp1.Next = temp2.Next
			temp2 = nil
			l.Length -= 1
			return nil
		} else {
			temp1 = temp1.Next
			temp2 = temp2.Next
		}
	}
	return errors.New("field does not exist")
}
