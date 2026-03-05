build_pager:
	protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      ./proto/v1/pager.proto


config ?=./config/config.yaml
start_server:
	go run ./internal/controller/cli/server/v1/server.go --config=$(config)

port ?=50051 # Take value from ./config/config.yaml
user_list:
	go run ./internal/controller/cli/client/v1/user_list/user_list.go --port=$(port)

register_user:
	go run ./internal/controller/cli/client/v1/register/register_user.go --port=$(port) --name="$(name)" --email=$(email) --password=$(password)

password ?=123456
login_user:
	go run ./internal/controller/cli/client/v1/login/login_user.go --port=$(port) --email=$(email) --password=$(password)

start_app:
	go run ./cmd/app/main.go --config=$(config)
