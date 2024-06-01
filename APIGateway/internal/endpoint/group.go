package endpoint

import (
	"ApiGateway/internal/service"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"log"
)

type GroupEndpoints struct {
	Register        endpoint.Endpoint
	GetCode         endpoint.Endpoint
	AddToGroup      endpoint.Endpoint
	RemoveFromGroup endpoint.Endpoint
	GetMembers      endpoint.Endpoint
}

type GroupRegisterRequest struct {
	User string
	Name string `json:"name"`
}

type GetCodeRequest struct {
	Name string `json:"name"`
	User string
}

type AddToGroupRequest struct {
	Token string `json:"token"`
	User  string
}

type RemoveFromGroupRequest struct {
	User   string
	Group  string `json:"group"`
	Remove string `json:"remove"`
}

type GetMembersRequest struct {
	Name string `json:"name"`
	User string
}

type BoolResponse struct {
	Success bool `json:"success"`
}

type GetCodeResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type GetMembersResponse struct {
	Members []string `json:"members"`
}

func MakeGroupEndpoints(srv service.GroupService) GroupEndpoints {
	return GroupEndpoints{
		Register:        MakeGroupRegisterEndpoint(srv),
		GetCode:         MakeGetCodeEndpoint(srv),
		AddToGroup:      MakeAddToGroupEndpoint(srv),
		RemoveFromGroup: MakeRemoveFromGroupEndpoint(srv),
		GetMembers:      MakeGetMembersEndpoint(srv),
	}
}

func MakeGroupRegisterEndpoint(srv service.GroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GroupRegisterRequest)
		success := srv.Register(req.User, req.Name)
		return BoolResponse{Success: success}, nil
	}
}

func MakeGetCodeEndpoint(srv service.GroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCodeRequest)
		token, err := srv.GetCode(req.Name, req.User)
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return GetCodeResponse{
			Code:  token,
			Error: fmt.Sprintf("%v", err),
		}, nil
	}
}

func MakeAddToGroupEndpoint(srv service.GroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddToGroupRequest)
		success := srv.AddToGroup(req.Token, req.User)
		return BoolResponse{Success: success}, nil
	}
}

func MakeRemoveFromGroupEndpoint(srv service.GroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RemoveFromGroupRequest)
		success := srv.RemoveFromGroup(req.User, req.Group, req.Remove)
		return BoolResponse{Success: success}, nil
	}
}

func MakeGetMembersEndpoint(srv service.GroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetMembersRequest)
		members := srv.GetMembers(req.Name, req.User)
		return GetMembersResponse{Members: members}, nil
	}
}
