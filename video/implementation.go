package video

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type serviceImplementation struct {
	cfg Config
	ctx context.Context
}

// CreateVideoService constructor for serviceImplementation
func CreateVideoService(cfg Config) Service {
	service := serviceImplementation{
		cfg: cfg,
	}

	return &service
}

func (s *serviceImplementation) Start(ctx context.Context) error {
	s.ctx = ctx

	cmd := exec.CommandContext(ctx, "raspivid", s.startArgs()...)
	err := cmd.Start()
	if err != nil {
		return err
	}

	go s.loop()

	return nil
}

func (s *serviceImplementation) outputPath() string {
	return "/tmp/godashvids"
}

func (s *serviceImplementation) loop() {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-s.ctx.Done():
			fmt.Printf("quiting...\n")
			return
		case <-tick:
			fmt.Printf("ticking: %s\n", time.Now())
		}
	}
}
