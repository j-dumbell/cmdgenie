package llm

import (
	"context"
)

type Model interface {
	Ask(ctx context.Context, prompt string) (string, error)
}
