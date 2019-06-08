package capture

import (
	"context"
	"godashpi/config"
	"testing"
	"time"
)

func TestServiceStart(t *testing.T) {
	cfg := config.Capture{}
	cmdRun := false
	cmd := mockCommand{
		func(ctx context.Context, cfg config.Capture) error {
			cmdRun = true

			return nil
		},
	}

	service := CreateCaptureService(cfg, &cmd)

	err := service.Start()
	if err != nil {
		t.Error("unexpected error", err)
	}

	time.Sleep(time.Second * 3)

	err = service.Stop()
	if err != nil {
		t.Error("unexpected error", err)
	}

	if !cmdRun {
		t.Error("command was not run")
	}
}
