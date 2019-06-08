package capture

// Service manages capture of video files and GC of files
type Service interface {
	Start() error
	Stop() error
}