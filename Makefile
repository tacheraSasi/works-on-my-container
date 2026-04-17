.PHONY: proto up

proto:
	protoc --go_out=. proto/user.proto

up: proto
	docker compose up -d --build

down:
	docker compose down