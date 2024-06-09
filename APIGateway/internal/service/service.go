package service

import "net/http"

var Secret = []byte("my-Secret")

type AuthService interface {
	Login(token string) (string, error)
	Verify(token string) bool
}

type ChatService interface {
	ForwardMessage(w http.ResponseWriter, r *http.Request)
	GetChats(user, time string) ([]map[string]string, error)
	GetMessages(from, to string, limit int64) ([]map[string]string, error)
}

type GroupService interface {
	Register(user, name string) bool
	GetCode(name, user string) (string, error)
	AddToGroup(token, user string) bool
	RemoveFromGroup(user, group, remove string) bool
	GetMembers(name, user string) []string
}

type AccountService interface {
	Register(token string) (bool, error)
	Remove(token string) bool
	List(token string) ([]string, error)
	Search(condition string) ([]string, error)
	AddContact(user, chat string) (bool, error)
	RemoveContact(user, chat string) (bool, error)
}
