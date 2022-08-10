package u_context

import "context"

func Get[K any](ctx context.Context, key string) K {
	return ctx.Value(key).(K)
}

func Set[K any](ctx context.Context, key string, value K) context.Context {
	return context.WithValue(ctx, key, value)
}