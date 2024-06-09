FROM golang:1.22.2 as build

WORKDIR /usr/src/app

COPY . .

RUN go build -o ./main ./cmd/chat/main.go


FROM ubuntu

WORKDIR /usr/app
RUN mkdir "config"

COPY --from=build /usr/src/app/main ./main
COPY --from=build /usr/src/app/deploy/config.yaml ./config/base.yaml

CMD ["./main"]