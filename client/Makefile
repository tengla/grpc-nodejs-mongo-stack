
srcdir = ../protos
dstdir = ./people

protoc:
	protoc \
		-I=$(srcdir) \
		--go_out=$(dstdir) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(dstdir) \
 		--go-grpc_opt=paths=source_relative \
		$(srcdir)/people.proto

build:
	go build -o dist/client main.go

clean:
	rm -rf dist