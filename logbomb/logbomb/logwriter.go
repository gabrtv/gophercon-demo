package logbomb

type logWriter interface {
	write(message string) error
}
