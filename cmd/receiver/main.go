package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"

	pb "github.com/chuhlomin/protobuf-telegram-post/pkg/demo"
)

func main() {
	message, err := ioutil.ReadFile("person.pb")
	if err != nil {
		log.Fatalf("Failed to read file `person.pb`: %v", err)
	}

	p := pb.Person{}
	if err = proto.Unmarshal(message, &p); err != nil {
		log.Fatalf("Failed to unmarshal message: %v", err)
	}

	log.Printf("File content: %s", p.String())
	// name:"John Doe"  id:1234  has_ponycopter:true
}
