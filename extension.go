package ent_graphql_go

import (
	"embed"
	"fmt"
	"reflect"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
	"github.com/graphql-go/graphql"
)

//go:embed templates/*
var templates embed.FS

type EntGraphers struct {
	entc.DefaultExtension
}

func NewExtension() entc.Extension {
	return &EntGraphers{}
}

var (
	isInterface = func(node *gen.Type) bool {
		for _, a := range node.Annotations {
			graphers, ok := a.(Annotation)
			if ok && graphers.IsInterface {
				return true
			}
		}
		return false
	}
	ouputField = func(field *gen.Field) bool {
		for _, a := range field.Annotations {
			graphers, ok := a.(Annotation)
			if ok {
				return !graphers.SkipOutput && hasType(field)
			}
		}
		return hasType(field)
	}
	createField = func(field *gen.Field) bool {
		for _, a := range field.Annotations {
			graphers, ok := a.(Annotation)
			if ok {
				return !graphers.SkipCreate && hasType(field)
			}
		}
		return hasType(field)
	}
	updateField = func(field *gen.Field) bool {
		for _, a := range field.Annotations {
			graphers, ok := a.(Annotation)
			if ok {
				return !graphers.SkipUpdate && hasType(field)
			}
		}
		return hasType(field)
	}
	hasType = func(field *gen.Field) bool {
		return gqlType(field) != nil
	}
	gqlType = func(field *gen.Field) graphql.Type {
		for _, a := range field.Annotations {
			graphers, ok := a.(Annotation)
			if ok && graphers.Type != nil {
				return graphers.Type
			}
		}
		var t graphql.Type
		switch true {
		case field.IsString():
			t = graphql.String
		case field.IsInt() || field.IsInt64():
			t = graphql.Int
		case field.IsBool():
			t = graphql.Boolean
		case field.IsTime():
			t = graphql.DateTime
		}
		if t == nil {
			return nil
		}
		if field.Optional {
			t = graphql.NewNonNull(t)
		}
		return t
	}
	gqlTypeString = func(field *gen.Field) string {
		gt := gqlType(field)
		if gt == nil {
			return "graphql.String"
		}
		rt := reflect.TypeOf(gt)
		fmt.Println(rt.Name())
		var t string
		switch true {
		case field.IsString():
			t = "graphql.String"
		case field.IsInt() || field.IsInt64():
			t = "graphql.Int"
		case field.IsBool():
			t = "graphql.Boolean"
		case field.IsTime():
			t = "graphql.DateTime"
		default:
			return "graphql.String"
		}
		if field.Optional {
			t = fmt.Sprintf("graphql.NewNonNull(%s)", t)
		}
		return t
	}
	TemplateFuncs = template.FuncMap{
		"HasType":       hasType,
		"GqlType":       gqlType,
		"GqlTypeString": gqlTypeString,
		"IsOutput":      ouputField,
		"IsCreate":      createField,
		"IsUpdate":      updateField,
		"IsInterface":   isInterface,
	}
)

func (EntGraphers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{},
	}
}

var GraphFilter = loadLocalTpl("templates/graphers_filter.tmpl")
var GraphQuery = loadLocalTpl("templates/graphers_query.tmpl")
var GraphTypes = loadLocalTpl("templates/graphers_types.tmpl")

func loadLocalTpl(path string) *gen.Template {
	content, err := templates.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return gen.MustParse(
		gen.NewTemplate(path).Funcs(TemplateFuncs).Parse(string(content)),
	)
}

func (EntGraphers) Templates() []*gen.Template {
	return []*gen.Template{
		GraphFilter,
		GraphQuery,
		GraphTypes,
	}
}
