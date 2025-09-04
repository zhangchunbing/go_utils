//go:build debug

package debug

import (
	"fmt"
)

func Fprint(w io.Writer, a ...any) {
	fmt.Fprint(w, a...)
}

func Print(a ...any) { 
	fmt.Print(a)
}

func Println(a ...any) {
    fmt.Println(a)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}