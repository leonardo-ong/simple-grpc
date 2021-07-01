package main

import (
	"context"
	"fmt"
	"log"

	"grpc_pokedex/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func servicePokedex() model.PokedexClient {

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to :8000", err)
	}

	return model.NewPokedexClient(conn)

}

func main() {
	pokedex := servicePokedex()

	fmt.Println("\n", "=========> start client")

	totalPokemon, err := pokedex.TotalPokemon(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Total Pokemon : %v", totalPokemon.Count)

	pokemonId := model.PokemonId{
		Id: 1,
	}

	pokemon, err := pokedex.GetPokemon(context.Background(), &pokemonId)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Pokemon : \n Id : %v \n Name : %s \n Height : %v \n Weight : %v \n Types : %s", pokemon.Id, pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.Types)

	pokemonList, err := pokedex.GetPokemonList(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(pokemonList)

}
