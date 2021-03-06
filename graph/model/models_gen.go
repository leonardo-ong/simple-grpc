// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Pokemon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  *Type  `json:"types"`
}

type Type struct {
	TypeID   string `json:"type_id"`
	TypeName string `json:"type_name"`
}
