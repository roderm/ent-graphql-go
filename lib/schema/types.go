package schema

import (
	"context"
	"strconv"

	"github.com/graphql-go/graphql"
)

var Types map[string]TypeResolver = make(map[string]TypeResolver)

type TypeResolver struct {
	Type    graphql.Type
	Resolve func(context.Context, int) (interface{}, error)
}

func GetID(input map[string]interface{}) int {
	idString, _ := input["id"].(string)
	id, _ := strconv.Atoi(idString)
	return id
}

var NodeType = graphql.NewInterface(graphql.InterfaceConfig{
	Name: "Node",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
})
