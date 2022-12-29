package ent_graphql_go

import "github.com/graphql-go/graphql"

type (
	// Annotation annotates fields and edges with metadata for templates.
	Annotation struct {
		Type graphql.Type

		SkipOutput bool
		SkipCreate bool
		SkipUpdate bool

		IsInterface bool
	}
)

func (Annotation) Name() string {
	return "EntGraphGophers"
}

func Type(t graphql.Type) Annotation {
	return Annotation{
		Type: t,
	}
}

func IsInterface() Annotation {
	return Annotation{
		IsInterface: true,
	}
}

func SkipOutput() Annotation {
	return Annotation{
		SkipOutput: true,
	}
}

func SkipCreate() Annotation {
	return Annotation{
		SkipCreate: true,
	}
}

func SkipUpdate() Annotation {
	return Annotation{
		SkipUpdate: true,
	}
}

func SkipAll() Annotation {
	return Annotation{
		SkipOutput: true,
		SkipCreate: true,
		SkipUpdate: true,
	}
}
