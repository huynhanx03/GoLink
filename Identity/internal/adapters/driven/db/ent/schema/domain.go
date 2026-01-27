package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	e "go-link/common/pkg/database/ent"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Mixin of the Domain.
func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.TimeMixin{},
		e.SoftDeleteMixin{},
	}
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain").
			Unique().
			NotEmpty(),
		field.Int("tenant_id"),
		field.Bool("is_verified").
			Default(false),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("domains").
			Field("tenant_id").
			Unique().
			Required(),
	}
}
