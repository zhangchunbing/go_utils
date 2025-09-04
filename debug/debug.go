//go:build debug

package debug

import (
	"fmt"
	"io"
)

func Fprint(w io.Writer, a ...any) {
	fmt.Fprint(w, a...)
}

func Fprintf(w io.Writer, a ...any) {
	fmt.Fprintf(w, a...)
}

func Fprintln(w io.Writer, a ...any) {
	fmt.Fprintln(w, a...)
}

func Print(a ...any) {
	fmt.Print(a)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Println(a ...any) {
	fmt.Println(a)
}

func Errorf(format string, a ...any) {
	fmt.Errorf(string, a...)
}
