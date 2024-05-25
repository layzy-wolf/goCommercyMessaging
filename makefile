run:
	go run ./AuthService/cmd/main.go --config=./AuthService/config/base.yaml
	go run ./GroupService/cmd/main.go --config=./GroupService/config/base.yaml
	go run ./ChatService/cmd/store/main.go --config=./ChatService/config/base.yaml
	go run ./ChatService/cmd/chat/main.go --config=./ChatService/config/base.yaml
	go run ./APIGateway/cmd/main.go --config=./APIGateway/config/base.yaml