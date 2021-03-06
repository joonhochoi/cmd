package cmd

import "context"

// This file provides utilities for request ID.
// Request ID is passed as context values.

type contextKey string

const (
	// RequestIDContextKey is a context key for request ID.
	RequestIDContextKey contextKey = "request_id"
)

func (k contextKey) String() string {
	return "cmd: context key: " + string(k)
}

// WithRequestID returns a new context with a request ID as a value.
func WithRequestID(ctx context.Context, reqid string) context.Context {
	return context.WithValue(ctx, RequestIDContextKey, reqid)
}
