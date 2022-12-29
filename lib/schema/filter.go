package schema

import (
	"github.com/graphql-go/graphql"
)

var filterTypes = make(map[string]*graphql.InputObject)

func NewWhereInput(name string) *graphql.InputObject {
	if _, ok := filterTypes[name]; !ok {
		whereInput := graphql.NewInputObject(graphql.InputObjectConfig{
			Name:   name,
			Fields: graphql.InputObjectConfigFieldMap{},
		})
		whereInput.AddFieldConfig("not", &graphql.InputObjectFieldConfig{
			Type: whereInput,
		})
		whereInput.AddFieldConfig("and", &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(whereInput)),
		})
		whereInput.AddFieldConfig("or", &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(whereInput)),
		})
		filterTypes[name] = whereInput
	}
	return filterTypes[name]
}

func NewWhereInputForOutput(name string, output graphql.Output) *graphql.InputObject {
	if _, ok := filterTypes[name]; !ok {
		wi := NewWhereInput(name)
		fields := graphql.FieldDefinitionMap{}
		switch t := output.(type) {
		case *graphql.Object:
			fields = t.Fields()
		case *graphql.Interface:
			fields = t.Fields()
		}
		for _, field := range fields {
			t := field.Type
			if nnt, ok := field.Type.(*graphql.NonNull); ok {
				t = nnt.OfType
			}
			if list, ok := field.Type.(*graphql.List); ok {
				wi.AddFieldConfig("has"+field.Name, &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				})
				wi.AddFieldConfig("has"+field.Name+"With", &graphql.InputObjectFieldConfig{
					Type: NewWhereInputForOutput(field.Name, list.OfType),
				})
				continue
			}
			addEQ(wi, field.Name, t)
			if field.Type.Name() == "String" {
				addStringCompare(wi, field.Name, field.Type)
			}
		}
		filterTypes[name] = wi
	}
	return filterTypes[name]
}
func addEQ(whereInput *graphql.InputObject, name string, gqlType graphql.Type) {
	whereInput.AddFieldConfig(name, &graphql.InputObjectFieldConfig{
		Type: gqlType,
	})
	whereInput.AddFieldConfig(name+"NEQ", &graphql.InputObjectFieldConfig{
		Type: gqlType,
	})
	listType := graphql.NewList(graphql.NewNonNull(gqlType))
	whereInput.AddFieldConfig(name+"In", &graphql.InputObjectFieldConfig{
		Type: listType,
	})
	whereInput.AddFieldConfig(name+"NotIn", &graphql.InputObjectFieldConfig{
		Type: listType,
	})
}

func addStringCompare(whereInput *graphql.InputObject, name string, gqlType graphql.Type) {
	whereInput.AddFieldConfig(name+"Contains", &graphql.InputObjectFieldConfig{
		Type: gqlType,
	})
	whereInput.AddFieldConfig(name+"HasPrefix", &graphql.InputObjectFieldConfig{
		Type: gqlType,
	})
	whereInput.AddFieldConfig(name+"HasSuffix", &graphql.InputObjectFieldConfig{
		Type: gqlType,
	})
}
