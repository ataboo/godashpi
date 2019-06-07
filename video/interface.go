package video

import "context"

type Service interface {
	Start(context.Context) error
}
