package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("price").Default(0),
		field.Time("date").Default(time.Now()),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
		Ref("books").
		Unique(),
	}
}
