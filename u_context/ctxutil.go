package u_context

import "context"

func Get[Value any](ctx context.Context, key string) (Value, bool) {
	v, ok := ctx.Value(key).(Value)
	return v, ok
}

func Set[Value any](ctx context.Context, key string, value Value) context.Context {
	return context.WithValue(ctx, key, value)
}
