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

type AccountServer struct {
	Register      *http.Server
	Remove        *http.Server
	List          *http.Server
	Search        *http.Server
	AddContact    *http.Server
	RemoveContact *http.Server
}

func decodeHTTPRegisterRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeHTTPRemoveRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.RemoveRequest
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

func decodeHTTPListRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.ListRequest
	ok, _, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.Token = r.Header.Get("Authorization")
	return request, nil
}

func decodeHTTPSearchRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.SearchRequest
	ok, _, err := verify.Verify(r)
	if !ok {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeHTTPContactRequest(_ context.Context, r *net.Request) (interface{}, error) {
	var request endpoints.ContactRequest
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

func encodeHTTPRAccountResponse(_ context.Context, w net.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeAccountHandler(cfg config.Cfg) *AccountServer {
	srv := service.NewAccountService(&cfg)
	en := endpoints.MakeAccountEndpoint(srv)
	return &AccountServer{
		Register:      http.NewServer(en.Register, decodeHTTPRegisterRequest, encodeHTTPRAccountResponse),
		Remove:        http.NewServer(en.Remove, decodeHTTPRemoveRequest, encodeHTTPRAccountResponse),
		List:          http.NewServer(en.List, decodeHTTPListRequest, encodeHTTPRAccountResponse),
		Search:        http.NewServer(en.Search, decodeHTTPSearchRequest, encodeHTTPRAccountResponse),
		AddContact:    http.NewServer(en.AddContact, decodeHTTPContactRequest, encodeHTTPResponse),
		RemoveContact: http.NewServer(en.RemoveContact, decodeHTTPContactRequest, encodeHTTPResponse),
	}
}
