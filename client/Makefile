
srcdir = ../protos
dstdir = .

protoc:
	protoc -I=$(srcdir) --go_out=$(dstdir) $(srcdir)/people.proto

build:
	go build -o dist/client main.go

clean:
	rm -rf dist