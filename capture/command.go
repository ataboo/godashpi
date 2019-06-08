package capture

import (
	"context"
	"godashpi/config"
)

// Command wraps a shell command that will capture the video
type Command interface {
	Start(ctx context.Context, cfg config.Capture) error
}
