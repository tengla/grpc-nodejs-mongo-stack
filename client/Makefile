
srcdir = ../server/protos
dstdir = ./people

build:
  go build -o dist/client main.go

protoc:
  protoc \
    -I=$(srcdir) \
    --go_out=$(dstdir) \
    --go_opt=paths=source_relative \
    --go-grpc_out=$(dstdir) \
    --go-grpc_opt=paths=source_relative \
    $(srcdir)/people.proto

run:
  go run main.go -server_addr localhost:30043

clean:
  rm -rf dist