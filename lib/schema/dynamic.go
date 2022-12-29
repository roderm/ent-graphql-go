package schema

type Config struct {
	Fields map[string]Field
}

type Field struct {
	DB          any                     `yaml:"db"`
	Fieldconfig any                     `yaml:"fieldconfig"`
	GraphQL     map[string]GraphQLField `yml:"graphQL"`
}

type GraphQLField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// func (c *Config) ApplyToSchema(s graphql.Schema) error {
// 	errors := []error{}
// }
