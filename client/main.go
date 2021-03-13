package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {

	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	log.Printf("Attempting %v", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	people.GetPeople()
	log.Printf("conn: %v", conn)
	defer conn.Close()
}
