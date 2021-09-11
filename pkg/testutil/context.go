package testutil

import (
	"context"

	"go.uber.org/zap"
)

func NewTestContext() context.Context {
	return context.WithValue(context.Background(), "key", zap.NewExample())
}
