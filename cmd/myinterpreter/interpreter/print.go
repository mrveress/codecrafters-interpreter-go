package interpreter

import (
	"fmt"
	"io"
	"os"
)

func Fprintf(w io.Writer, format string, a ...interface{}) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		/* Skip */
	}
}

func Fprintln(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		/* Skip */
	}
}

func Errorf(code int, format string, a ...interface{}) {
	Fprintf(os.Stderr, format, a...)
	os.Exit(code)
}

func Errorln(code int, a ...interface{}) {
	Fprintln(os.Stderr, a...)
	os.Exit(code)
}
