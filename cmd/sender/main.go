package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"

	pb "github.com/chuhlomin/protobuf-telegram-post/pkg/demo"
)

func main() {
	p := pb.Person{
		Name:          "John Doe",
		Id:            1234,
		HasPonycopter: true,
	}
	message, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalf("Failed to marshal %v", err)
	}

	if err = ioutil.WriteFile("person.pb", message, 0644); err != nil {
		log.Fatalf("Failed to write file `person.pb`: %v", err)
	}
	log.Printf("File `person.pb` updated with content: %s", string(message))
}
