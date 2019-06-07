package video

import (
	"context"
	"testing"
	"time"
)

func TestLoopCancel(t *testing.T) {
	cfg := Config{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	service := CreateVideoService(cfg)

	service.Start(ctx)

	time.Sleep(time.Second * 3)

	cancel()
}
