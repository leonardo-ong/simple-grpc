package main

import (
	"context"
	"log"
	"net"

	"grpc_pokedex/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var pokemonList *model.PokemonList

func init() {
	pokemonList = new(model.PokemonList)
	pokemonList.List = make([]*model.Pokemon, 0)

	bulbasaur := model.Pokemon{
		Id:     "001",
		Name:   "Bulbasaur",
		Height: "7",
		Weight: "69",
		Types:  "Grass, Poison",
	}

	charizard := model.Pokemon{
		Id:     "002",
		Name:   "Charizard",
		Height: "17",
		Weight: "905",
		Types:  "Fire, Flying",
	}

	pokemonList.List = append(pokemonList.List, &bulbasaur)
	pokemonList.List = append(pokemonList.List, &charizard)
}

type PokedexServer struct{}

func (PokedexServer) TotalPokemon(ctx context.Context, void *empty.Empty) (*model.PokemonCount, error) {
	count := model.PokemonCount{
		Count: 1118,
	}
	return &count, nil
}

func (PokedexServer) GetPokemon(ctx context.Context, param *model.PokemonId) (*model.Pokemon, error) {
	pokemonId := param.Id

	pokemon := model.Pokemon{
		Id:     pokemonId,
		Name:   "Bulbasaur",
		Height: "7",
		Weight: "69",
		Types:  "Grass, Poison",
	}

	return &pokemon, nil
}

func (PokedexServer) GetPokemonList(ctx context.Context, void *empty.Empty) (*model.PokemonList, error) {
	return pokemonList, nil
}

func main() {
	srv := grpc.NewServer()
	var pokedexSrv PokedexServer
	model.RegisterPokedexServer(srv, pokedexSrv)

	log.Println("Starting RPC server at :8000")

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("could not listen to :8000: %v", err)
	}

	log.Fatal(srv.Serve(l))
}
