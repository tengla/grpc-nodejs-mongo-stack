package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/tengla/grpc-nodejs-mongo-stack/client/people"
	"google.golang.org/grpc"
)

var (
	serverAddr   = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	insertPerson = flag.String("insert", "", "Insert a person")
	deletePerson = flag.String("delete", "", "Delete a person by id")
)

type InsertPerson struct {
	Name string
	Age  int64
}

func main() {

	flag.Parse()

	var pax InsertPerson
	json.Unmarshal([]byte(*insertPerson), &pax)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	log.Printf("Attempting connect to %v", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := people.NewPeopleServiceClient(conn)

	log.Println("")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if len(*deletePerson) > 0 {

		log.Printf("Delete %s", *deletePerson)

		res, err := client.Remove(ctx, &people.PersonId{
			Id: *deletePerson,
		})

		if err != nil {
			log.Fatalf("Failed to delete, %v", err)
		}
		log.Printf("%v", res)
	} else if len(pax.Name) > 0 {

		log.Printf("Insert %s %d", pax.Name, pax.Age)
		var p = &people.Person{
			Name: pax.Name,
			Age:  int32(pax.Age),
		}
		res, err := client.Insert(ctx, p)
		if err != nil {
			log.Fatalf("Failed to insert, %v", err)
		}
		log.Printf("%v", res)
	} else {

		people, err := client.GetAll(ctx, &people.Empty{})
		if err != nil {
			log.Fatalf("Failed to get all, %v", err)
		}
		for _, person := range people.People {
			log.Printf("Person { Id = '%s', Name = '%s', Age = %d }",
				person.Id, person.Name, person.Age)
		}
	}
}
