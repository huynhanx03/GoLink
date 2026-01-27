package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	e "go-link/common/pkg/database/ent"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Mixin of the Subscription.
func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.TimeMixin{},
		e.SoftDeleteMixin{},
	}
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").
			Positive().
			Comment("Reference to Identity Service Tenant"),
		field.Int("plan_id").
			Comment("Reference to Plan"),
		field.Enum("status").
			Values("ACTIVE", "PENDING", "PAID", "PAST_DUE", "CANCELED").
			Default("PENDING").
			Comment("Subscription status"),
		field.Time("current_period_start").
			Comment("Start of the current billing cycle"),
		field.Time("current_period_end").
			Comment("End of the current billing cycle (Expiry Date)"),
		field.Bool("cancel_at_period_end").
			Default(false).
			Comment("Auto-renew flag"),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).
			Ref("subscriptions").
			Unique().
			Field("plan_id").
			Required(),
		edge.To("invoices", Invoice.Type),
	}
}

// Indexes of the Subscription.
func (Subscription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
		index.Fields("status"),
	}
}
