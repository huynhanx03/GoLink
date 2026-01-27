package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	e "go-link/common/pkg/database/ent"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Mixin of the Invoice.
func (Invoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		e.TimeMixin{},
		e.SoftDeleteMixin{},
	}
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("subscription_id").
			Comment("Reference to Subscription"),
		field.Int("tenant_id").
			Positive().
			Comment("Context info for tenant"),
		field.Float("amount").
			Comment("Invoice amount"),
		field.String("currency").
			MaxLen(3).
			NotEmpty().
			Comment("Currency code"),
		field.Enum("status").
			Values("PENDING", "PAID", "FAILED", "REFUNDED").
			Default("PENDING").
			Comment("Invoice status"),
		field.String("payment_id").
			Optional().
			Comment("Stored Payment Gateway Transaction ID"),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subscription", Subscription.Type).
			Ref("invoices").
			Unique().
			Field("subscription_id").
			Required(),
	}
}

// Indexes of the Invoice.
func (Invoice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
		index.Fields("subscription_id"),
	}
}
