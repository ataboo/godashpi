package video

type Service interface {
	Start() error
	Stop() error
}
