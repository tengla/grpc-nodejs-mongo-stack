package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/tengla/grpc-nodejs-mongo-stack/client/people"
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
	defer conn.Close()

	client := people.NewPeopleServiceClient(conn)
	log.Printf("client: %v", client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	people, err := client.GetAll(ctx, &people.Empty{})
	if err != nil {
		log.Fatalf("Failed to get all, %v", err)
	}
	log.Printf("people: %v", people)
}
