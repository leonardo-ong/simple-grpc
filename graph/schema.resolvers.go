package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"grpc_pokedex/graph/generated"
	"grpc_pokedex/graph/model"
	"log"

	grpcModel "grpc_pokedex/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func servicePokedex() grpcModel.PokedexClient {

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to :8000", err)
	}

	return grpcModel.NewPokedexClient(conn)

}

func (r *queryResolver) Pokemons(ctx context.Context) ([]*model.Pokemon, error) {
	pokedex := servicePokedex()

	var pokemons []*model.Pokemon

	pokemonList, err := pokedex.GetPokemonList(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, val := range pokemonList.List {
		var pokemon model.Pokemon

		pokemon.ID = string(val.Id)
		pokemon.Name = val.Name
		pokemon.Height = int(val.Height)
		pokemon.Weight = int(val.Weight)
		pokemon.Types = &model.Type{TypeName: val.Types}

		pokemons = append(pokemons, &pokemon)
	}

	return pokemons, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
