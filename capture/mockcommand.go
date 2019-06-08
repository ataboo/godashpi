package capture

import (
	"context"
	"godashpi/config"
)

// MockCommand mock running the video capture command
type mockCommand struct {
	StartClosure func(context.Context, config.Capture) error
}

// Start start the command
func (c *mockCommand) Start(ctx context.Context, cfg config.Capture) error {
	if c.StartClosure == nil {
		return nil
	}

	return c.StartClosure(ctx, cfg)
}
