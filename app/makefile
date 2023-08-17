gen-all:
	# protoc -I proto/todos.proto --go-grpc_out=./http
	# protoc -I proto/todos.proto --go-grpc_out=./todos
	protoc -I proto --go_out=./http proto/todos.proto
	protoc -I proto --go_out=./todos proto/todos.proto

http-up:
	cd http && go run cmd/main.go
todos-up:
	cd todos && go run cmd/rabbitmq/main.go

all-up: http-up todos-up

.PHONY: all-up gen-all http-up todos-up
