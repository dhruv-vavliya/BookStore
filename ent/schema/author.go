package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("name").Default("unknown"),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		// Author -> Books [one-to-many]
		edge.To("books", Book.Type).
		Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
