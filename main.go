package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nantaphop/go-grpc-example/location"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Can't connect to server")
	}
	defer conn.Close()
	client := pb.NewLocationClient(conn)
	location := pb.LatLng{}
	location.Lat = 10.232332
	location.Lng = 10.232332
	msg, err := client.Print(context.Background(), &location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
