package gosrcfmt

import (
	"fmt"
)

type wrapError struct {
	msg string
	err error
}

func (w *wrapError) Error() string {
	return fmt.Sprintf("%s: %s", w.msg, w.err)
}

func (w *wrapError) Unwrap() error {
	return w.err
}
