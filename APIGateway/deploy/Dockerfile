FROM golang:1.22.2 as build

WORKDIR /usr/src/app

COPY . .

WORKDIR /usr/src/app/APIGateway

RUN go build -o ./main ./cmd/main.go


FROM ubuntu

WORKDIR /usr/app
RUN mkdir "config"

COPY --from=build /usr/src/app/APIGateway/main ./main
COPY --from=build /usr/src/app/APIGateway/deploy/config.yaml ./config/base.yaml

CMD ["./main"]