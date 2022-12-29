package paging

import "github.com/graphql-go/graphql"

var PageCursor = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageCursor",
	Fields: graphql.Fields{
		"cursor": &graphql.Field{
			Type: Cursor,
		},
		"pageNumber": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})
var PageCursors = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageCursors",
	Fields: graphql.Fields{
		"first": &graphql.Field{
			Type: PageCursor,
		},
		"around": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PageCursor))),
		},
		"last": &graphql.Field{
			Type: PageCursor,
		},
	},
})
var PageInfo = graphql.NewObject(graphql.ObjectConfig{
	Description: "https://relay.dev/graphql/connections.htm#sec-Connection-Types.Fields.PageInfo",
	Name:        "PageInfo",
	Fields: graphql.Fields{
		"hasNextPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"hasPreviousPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"startCursor": &graphql.Field{
			Type: Cursor,
		},
		"endCursor": &graphql.Field{
			Type: Cursor,
		},
	},
})
