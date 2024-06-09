package endpoint

import (
	"ApiGateway/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AccountEndpoints struct {
	Register      endpoint.Endpoint
	Remove        endpoint.Endpoint
	List          endpoint.Endpoint
	Search        endpoint.Endpoint
	AddContact    endpoint.Endpoint
	RemoveContact endpoint.Endpoint
}

type RegisterRequest struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type RemoveRequest struct {
	User string `json:"user"`
}

type RemoveResponse struct {
	Success bool `json:"success"`
}

type ListRequest struct {
	Token string `json:"token"`
}

type ContactRequest struct {
	User string `json:"user"`
	Chat string `json:"chat"`
}

type ListResponse struct {
	Users []string `json:"users"`
}

type SearchRequest struct {
	Condition string `json:"condition"`
}

type SearchResponse struct {
	Users []string `json:"users"`
}

type ContactResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

func MakeAccountEndpoint(srv service.AccountService) AccountEndpoints {
	return AccountEndpoints{
		Register:      MakeRegisterEndpoint(srv),
		Remove:        MakeRemoveEndpoint(srv),
		List:          MakeListEndpoint(srv),
		Search:        MakeSearchEndpoint(srv),
		AddContact:    MakeAddContactEndpoint(srv),
		RemoveContact: MakeRemoveContactEndpoint(srv),
	}
}

func MakeRegisterEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterRequest)
		success, err := srv.Register(req.Token)
		if err != nil {
			return RegisterResponse{
				Success: success,
				Error:   err.Error(),
			}, err
		}
		return RegisterResponse{
			Success: success,
			Error:   "",
		}, nil
	}
}

func MakeRemoveEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RemoveRequest)
		success := srv.Remove(req.User)
		return RemoveResponse{Success: success}, nil
	}
}

func MakeListEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ListRequest)
		users, err := srv.List(req.Token)
		return ListResponse{Users: users}, err
	}
}

func MakeSearchEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SearchRequest)
		users, err := srv.Search(req.Condition)
		return SearchResponse{Users: users}, err
	}
}

func MakeAddContactEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ContactRequest)
		success, err := srv.AddContact(req.User, req.Chat)
		return ContactResponse{Success: success}, err
	}
}
func MakeRemoveContactEndpoint(srv service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ContactRequest)
		success, err := srv.RemoveContact(req.User, req.Chat)
		return ContactResponse{Success: success}, err
	}
}
