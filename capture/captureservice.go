package capture

import (
	"context"
	"fmt"
	"godashpi/config"
	"time"
)

type captureService struct {
	cfg    config.Capture
	ctx    context.Context
	cancel context.CancelFunc
	cmd    Command
}

// CreateCaptureService constructor for captureService
func CreateCaptureService(cfg config.Capture, command Command) Service {
	service := captureService{
		cfg: cfg,
	}

	return &service
}

func (s *captureService) Start() error {
	if s.ctx != nil || s.cancel != nil {
		return fmt.Errorf("service already started")
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())

	if err := s.cmd.Start(s.ctx, s.cfg); err != nil {
		return err
	}

	go s.loop()

	return nil
}

func (s *captureService) Stop() error {
	if s.ctx == nil || s.cancel == nil {
		return fmt.Errorf("service already stopped")
	}

	s.cancel()
	s.ctx = nil
	s.cancel = nil

	return nil
}

func (s *captureService) outputPath() string {
	return "/tmp/godashvids"
}

func (s *captureService) loop() {
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
