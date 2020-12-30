# in order for this to work, install rust plugin for protoc

# ensure rust is installed, and cargo is in your path
# cargo install protobuf-codegen

# gen:
	# protoc --rust_out=../fighter-service/src --go_out=./dist/go  ./src/*

gen-rust:
	cargo build

gen-go:
	protoc --go_out=plugins=grpc:./../battle-service/proto/ *.proto
 