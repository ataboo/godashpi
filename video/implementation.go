package video

import "context"

type serviceImplementation struct {
	cfg Config
	ctx context.Context
}

// CreateVideoService constructor for serviceImplementation
func CreateVideoService(ctx context.Context, cfg Config) Service {
	service := serviceImplementation{
		cfg: cfg,
		ctx: ctx,
	}

	return &service
}

func (s *serviceImplementation) Start() error {
	return nil
}

func (s *serviceImplementation) Stop() error {
	return nil
}
