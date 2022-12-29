package schema

import "github.com/graphql-go/graphql"

// []func(*entdb.EventVirtual) error

var ExtensionStore = []Extension{}

type ExtensionField struct {
	Name  string
	Field *graphql.InputObjectFieldConfig
}
type Extension struct {
	Inputs  []string
	Fields  []ExtensionField
	Resolve func(graphql.ResolveParams, any) error
}

func AfterSave[T any](params graphql.ResolveParams) []func(T) error {
	funcs := []func(T) error{}
	for _, ext := range ExtensionStore {
		funcs = append(funcs, func(t T) error {
			return ext.Resolve(params, t)
		})
	}
	return funcs
}

func GetInput(input *graphql.InputObject) *graphql.InputObject {
	for _, ext := range ExtensionStore {
		for _, inputName := range ext.Inputs {
			if inputName == input.Name() {
				for _, field := range ext.Fields {
					input.AddFieldConfig(field.Name, field.Field)
				}
			}
		}
	}
	return input
}
