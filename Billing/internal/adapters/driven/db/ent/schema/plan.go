package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	e "go-link/common/pkg/database/ent"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Mixin of the Plan.
func (Plan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.TimeMixin{},
		e.SoftDeleteMixin{},
	}
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(100).
			NotEmpty().
			Comment("Display name"),
		field.String("description").
			MaxLen(200).
			Optional().
			Comment("Plan description"),
		field.Float("base_price").
			Positive().
			Comment("Base price in USD"),
		field.String("period").
			MaxLen(20).
			NotEmpty().
			Comment("Billing period: 'month' or 'year'"),
		field.JSON("features", map[string]interface{}{}).
			Optional().
			Comment("Feature limits (e.g., {\"max_links\": 1000})"),
		field.Bool("is_active").
			Default(true).
			Comment("Soft delete flag"),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriptions", Subscription.Type),
	}
}
