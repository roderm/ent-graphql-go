package paging

import "github.com/graphql-go/graphql"

var connTypes = make(map[string]*graphql.Object)

func NewConnection(gqltype graphql.Type) *graphql.Object {
	if _, ok := connTypes[gqltype.Name()+"Connection"]; !ok {
		edge := graphql.NewObject(graphql.ObjectConfig{
			Name: gqltype.Name() + "Edge",
			Fields: graphql.Fields{
				"cursor": &graphql.Field{
					Type: Cursor,
				},
				"node": &graphql.Field{
					Type: gqltype,
				},
			},
		})
		connTypes[gqltype.Name()+"Connection"] = graphql.NewObject(graphql.ObjectConfig{
			Name: gqltype.Name() + "Connection",
			Fields: graphql.Fields{
				"edges": &graphql.Field{
					Type: graphql.NewList(edge),
				},
				"pageInfo": &graphql.Field{
					Type: PageInfo,
				},
				"pageCursor": &graphql.Field{
					Type: PageCursors,
				},
				"totalCount": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return connTypes[gqltype.Name()+"Connection"]
}
