package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Float("price").Positive(),
		field.String("preview_url").NotEmpty().Default(""),
		field.Strings("categories"),
		field.Strings("images"),
		field.String("imagesStorage"),
		field.String("description").NotEmpty(),
		field.Int("stock"),
		field.Int("stock_reserved"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("seller", User.Type).Ref("products").Unique(),
	}
}
