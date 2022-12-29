package paging

import "github.com/graphql-go/graphql"

var Cursor = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Cursor",
	Serialize: func(value interface{}) interface{} {
		return value
	},
})

var Paging = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Paging",
	Fields: graphql.InputObjectConfigFieldMap{
		"after": &graphql.InputObjectFieldConfig{
			Type: Cursor,
		},
		"first": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"before": &graphql.InputObjectFieldConfig{
			Type: Cursor,
		},
		"last": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"offset": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
