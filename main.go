package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	proto "github.com/JPF3N998/gRPC-Golang-Server/proto"

	"google.golang.org/grpc"
)

func printPokemonInfo(client proto.PokedexClient, name string) {

	request := proto.SearchRequest{
		Name: name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pkmn, err := client.GetPokemon(ctx, &request)
	if err != nil {
		log.Fatalf("%v.GetPokemon(_) = _, %v: ", client, err)
	}
	prettyprinted, _ := json.MarshalIndent(pkmn, "", " ")
	log.Println(string(prettyprinted))
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:6809", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewPokedexClient(conn)

	printPokemonInfo(client, "bulbasaur")
}
