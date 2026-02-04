package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	e "go-link/common/pkg/database/ent"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Mixin of the Tenant.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.TimeMixin{},
		e.SoftDeleteMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("plan_id").
			Default(0),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant_members", TenantMember.Type),
		edge.To("domains", Domain.Type),
	}
}
