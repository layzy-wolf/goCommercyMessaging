package transport

import (
	"ApiGateway/config"
	endpoints "ApiGateway/internal/endpoint"
	"ApiGateway/internal/service"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/transport/http"
	net "net/http"
)

type GroupServer struct {
	Register        *http.Server
	GetCode         *http.Server
	AddToGroup      *http.Server
	RemoveFromGroup *http.Server
	GetMembers      *http.Server
}

func decodeHTTPGroupRegRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.GroupRegisterRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func decodeHTTPGetCodeRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.GetCodeRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func decodeHTTPAddToGroupRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.AddToGroupRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func decodeHTTPRemoveGrRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.RemoveFromGroupRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func decodeHTTPGetMembersRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.GetMembersRequest
	ok, usr, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.User = usr
	return request, nil
}

func encodeHTTPResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewGroupHandler(cfg config.Cfg) *GroupServer {
	srv := service.NewGroupChatService(cfg)
	en := endpoints.MakeGroupEndpoints(srv)
	handler := &GroupServer{
		Register:        http.NewServer(en.Register, decodeHTTPGroupRegRequest, encodeHTTPResponse),
		GetCode:         http.NewServer(en.GetCode, decodeHTTPGetCodeRequest, encodeHTTPResponse),
		AddToGroup:      http.NewServer(en.AddToGroup, decodeHTTPAddToGroupRequest, encodeHTTPResponse),
		RemoveFromGroup: http.NewServer(en.RemoveFromGroup, decodeHTTPRemoveGrRequest, encodeHTTPResponse),
		GetMembers:      http.NewServer(en.GetMembers, decodeHTTPGetMembersRequest, encodeHTTPResponse),
	}
	return handler
}
