package main

import (
	"fmt"
	"log"
	tutorial "protobuf-dojo/pkg/pb"
	"protobuf-dojo/util"

	"google.golang.org/protobuf/proto"
)

func main() {
	util.Util()

	// Create a Person object
	p := &tutorial.Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	// Serialize object to bytes
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	newP := &tutorial.Person{}
	if err := proto.Unmarshal(data, newP); err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Printf("Deserialized data: %+v\n", newP)
}
