.PHONY: proto up

proto:
	protoc --go_out=. --go-grpc_out=. proto/user.proto

up: proto
	docker compose up -d --build

down:
	docker compose down