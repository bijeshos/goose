package error

import (
	"fmt"
	"time"
)

type GooseError struct {
	Time    time.Time
	Message string
}

func (e *GooseError) Error() string {
	return fmt.Sprintf("Error occurred at %v, with message '%s'",
		e.Time, e.Message)
}

func raiseError(time time.Time, message string) error {
	return &GooseError{time, message}
}
