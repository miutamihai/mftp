package driver

type Driver interface {
	Write() error
	Initialize()
	IsInitialized() bool
	Log(values ...any)
}
