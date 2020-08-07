package appcontext

import (
	"context"

	"github.com/sirupsen/logrus"
)

type indexContext int

const (
	loggerKey indexContext = iota
)

// WithLogger receives a logger and stores it as the base logger for the package.
func WithLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
