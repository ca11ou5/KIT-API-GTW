proto:
	protoc ./internal/auth/pb/auth.proto --go_out=plugins=grpc:.

buildimage:
	docker build -t api-gtw .

compose:
	docker compose -p kit up