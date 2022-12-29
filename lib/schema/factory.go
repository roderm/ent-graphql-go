package schema

import "entgo.io/ent/dialect/sql"

type SchemaTypeFactory struct{}

type Query interface {
	Where(...func(*sql.Selector))
}

func (f *SchemaTypeFactory) CreateFilter()                                              {}
func (f *SchemaTypeFactory) ApplyFilter(s *sql.Selector, filter map[string]interface{}) {}
func (f *SchemaTypeFactory) filterToPredicate(filter map[string]interface{}) *sql.Predicate {
	p := sql.P()
	return p
}
