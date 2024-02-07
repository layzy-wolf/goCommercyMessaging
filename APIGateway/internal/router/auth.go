package router

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

var addr = "localhost:40444"

func Register() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewAuthClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")

		params := httprouter.ParamsFromContext(r.Context())

		res, err := c.Register(ctx, &pb.RegisterRequest{
			Email:  params.ByName("email"),
			Login:  params.ByName("login"),
			Passwd: params.ByName("password"),
		})

		if err != nil {
			log.Fatalf("could`t register: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("{Could`t register: bad request}")
			return
		}

		//TODO
		return
	}
}
