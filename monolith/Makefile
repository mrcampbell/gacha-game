build:
	protoc --go_out=plugins=grpc:./internal/protos/ -I ../protos/*.proto
	swag init -g internal/http/swagger.go
	go build -o server main.go

run: build
	./server

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' -R './docs/*' make run

	protoc -I ../protos/ --go_out=plugins=grpc:./ fighter.proto