package capture

import (
	"context"
	"godashpi/config"
	"os/exec"
)

type raspividCommand struct {
	//
}

// Start the command
func (r *raspividCommand) Start(ctx context.Context, cfg config.Capture) error {
	cmd := exec.CommandContext(ctx, "raspivid", cfg.ToArgs()...)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
