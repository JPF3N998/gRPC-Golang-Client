package main

import (
	"context"
	"log"
	"time"

	pb "local/pokemon"

	"google.golang.org/grpc"
)

func printPokemonInfo(client pb.SearchPokedexClient, name string) {

	request := pb.SearchRequest{
		Name: name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pkmn, err := client.GetPokemon(ctx, &request)
	if err != nil {
		log.Fatalf("%v.GetPokemon(_) = _, %v: ", client, err)
	}
	log.Println(pkmn)
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
	client := pb.NewSearchPokedexClient(conn)

	printPokemonInfo(client, "bulbasaur")
}
