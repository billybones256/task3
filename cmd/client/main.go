package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"task3/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	days, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	connection, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewSoupClient(connection)
	res, err := client.BuildMenu(context.Background(), &api.SoupRequest{Days: int32(days)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetAnswer())
}
