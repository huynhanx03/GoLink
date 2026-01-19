package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	e "go-link/common/pkg/database/ent"
)

// TenantMember holds the schema definition for the TenantMember entity.
type TenantMember struct {
	ent.Schema
}

// Mixin of the TenantMember.
func (TenantMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.BaseMixin{},
	}
}

// Fields of the TenantMember.
func (TenantMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id"),
		field.Int("user_id"),
		field.Int("role_id"),
	}
}

// Edges of the TenantMember.
func (TenantMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("tenant_members").
			Field("tenant_id").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("tenant_members").
			Field("user_id").
			Unique().
			Required(),
		edge.From("role", Role.Type).
			Ref("tenant_members").
			Field("role_id").
			Unique().
			Required(),
	}
}

// Indexes of the TenantMember.
func (TenantMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id", "user_id").
			Unique(),
	}
}
