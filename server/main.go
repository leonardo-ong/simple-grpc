package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"grpc_pokedex/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var pokemonList *model.PokemonList

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:password@/go_db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// func init() {
// 	pokemonList = new(model.PokemonList)
// 	pokemonList.List = make([]*model.Pokemon, 0)

// 	bulbasaur := model.Pokemon{
// 		Id:     "001",
// 		Name:   "Bulbasaur",
// 		Height: "7",
// 		Weight: "69",
// 		Types:  "Grass, Poison",
// 	}

// 	charizard := model.Pokemon{
// 		Id:     "002",
// 		Name:   "Charizard",
// 		Height: "17",
// 		Weight: "905",
// 		Types:  "Fire, Flying",
// 	}

// 	pokemonList.List = append(pokemonList.List, &bulbasaur)
// 	pokemonList.List = append(pokemonList.List, &charizard)
// }

type PokedexServer struct{}

func (PokedexServer) TotalPokemon(ctx context.Context, void *empty.Empty) (*model.PokemonCount, error) {
	db := dbConn()
	defer db.Close()
	res, err := db.Query("SELECT COUNT(id) AS totalPokemon FROM pokedex")
	if err != nil {
		panic(err.Error())
	}
	var total int
	for res.Next() {
		var totalPokemon int
		err = res.Scan(&totalPokemon)
		if err != nil {
			panic(err.Error())
		}

		total = totalPokemon
	}

	count := model.PokemonCount{
		Count: int32(total),
	}
	return &count, nil
}

func (PokedexServer) GetPokemon(ctx context.Context, param *model.PokemonId) (*model.Pokemon, error) {
	db := dbConn()
	defer db.Close()
	pokemonId := param.Id

	res, err := db.Query("SELECT p.id, p.name, p.height, p.weight, t.type_name FROM pokedex as p INNER JOIN type as t ON p.types = t.type_id WHERE id=?", pokemonId)
	if err != nil {
		panic(err.Error())
	}

	var pokemon model.Pokemon

	for res.Next() {

		var id, weight, height int
		var name, types string

		err := res.Scan(&id, &name, &weight, &height, &types)

		if err != nil {
			panic(err.Error())
		}

		pokemon.Id = int32(id)
		pokemon.Name = name
		pokemon.Height = int32(height)
		pokemon.Weight = int32(weight)
		pokemon.Types = types
	}

	return &pokemon, nil
}

func (PokedexServer) GetPokemonList(ctx context.Context, void *empty.Empty) (*model.PokemonList, error) {
	db := dbConn()
	defer db.Close()

	res, err := db.Query("SELECT p.id, p.name, p.height, p.weight, t.type_name FROM pokedex as p INNER JOIN type as t ON p.types = t.type_id")
	if err != nil {
		panic(err.Error())
	}

	pokemonList = new(model.PokemonList)
	pokemonList.List = make([]*model.Pokemon, 0)

	for res.Next() {
		var pokemon model.Pokemon
		var id, weight, height int
		var name, types string

		err := res.Scan(&id, &name, &weight, &height, &types)

		if err != nil {
			panic(err.Error())
		}

		pokemon.Id = int32(id)
		pokemon.Name = name
		pokemon.Height = int32(height)
		pokemon.Weight = int32(weight)
		pokemon.Types = types
		pokemonList.List = append(pokemonList.List, &pokemon)
	}

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
