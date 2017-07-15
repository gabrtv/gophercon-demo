package logbomb

import "fmt"

type errUnrecognizedLogWriterType struct {
	logWriterType string
}

func newErrUnrecognizedLogWriterType(logWriterType string) errUnrecognizedLogWriterType {
	return errUnrecognizedLogWriterType{logWriterType: logWriterType}
}

func (e errUnrecognizedLogWriterType) Error() string {
	return fmt.Sprintf("Unrecognized LogWriter type: %s", e.logWriterType)
}
