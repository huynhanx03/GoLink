package metadata

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"

	"go-link/common/pkg/constraints"
)

type dataType int

const (
	typeString dataType = iota
	typeInt
	typeInt64
	typeBool
)

// mapping defines the configuration for context propagation
type mapping struct {
	header string
	key    interface{}
	dtype  dataType
}

// configuredMappings defines the mapping between gRPC metadata headers and context keys.
var configuredMappings = []mapping{
	{header: "x-user-id", key: constraints.ContextKeyUserID, dtype: typeInt},
	{header: "x-tenant-id", key: constraints.ContextKeyTenantID, dtype: typeInt},
	{header: "x-role", key: constraints.ContextKeyRole, dtype: typeString},
	{header: "x-role-level", key: constraints.ContextKeyRoleLevel, dtype: typeInt},
	{header: "x-tier-id", key: constraints.ContextKeyTierID, dtype: typeInt},
	{header: "x-is-admin", key: constraints.ContextKeyIsAdmin, dtype: typeBool},
}

// EnsureOutgoingContext injects mapped values from the context into the outgoing gRPC metadata.
func EnsureOutgoingContext(ctx context.Context) context.Context {
	md := metadata.New(nil)

	for _, m := range configuredMappings {
		val := ctx.Value(m.key)
		if val == nil {
			continue
		}

		var strVal string

		switch m.dtype {
		case typeString:
			if v, ok := val.(string); ok {
				strVal = v
			}
		case typeInt:
			if v, ok := val.(int); ok {
				strVal = strconv.Itoa(v)
			}
		case typeInt64:
			if v, ok := val.(int64); ok {
				strVal = strconv.FormatInt(v, 10)
			}
		case typeBool:
			if v, ok := val.(bool); ok {
				strVal = strconv.FormatBool(v)
			}
		}

		if strVal != "" {
			md.Set(m.header, strVal)
		}
	}

	if existingMD, ok := metadata.FromOutgoingContext(ctx); ok {
		md = metadata.Join(existingMD, md)
	}

	return metadata.NewOutgoingContext(ctx, md)
}

// ExtractIncomingContext extracts mapped values from incoming gRPC metadata and injects them into the context.
func ExtractIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	for _, m := range configuredMappings {
		vals := md.Get(m.header)
		if len(vals) == 0 || vals[0] == "" {
			continue
		}

		valStr := vals[0]

		switch m.dtype {
		case typeString:
			ctx = context.WithValue(ctx, m.key, valStr)
		case typeInt:
			if v, err := strconv.Atoi(valStr); err == nil {
				ctx = context.WithValue(ctx, m.key, v)
			}
		case typeInt64:
			if v, err := strconv.ParseInt(valStr, 10, 64); err == nil {
				ctx = context.WithValue(ctx, m.key, v)
			}
		case typeBool:
			if v, err := strconv.ParseBool(valStr); err == nil {
				ctx = context.WithValue(ctx, m.key, v)
			}
		}
	}

	return ctx
}
