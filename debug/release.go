//go:build !debug

package debug

import "io"

func Fprint(w io.Writer, a ...any)   {}
func Fprintf(w io.Writer, a ...any)  {}
func Fprintln(w io.Writer, a ...any) {}
func Print(a ...any)                 {}
func Printf(format string, a ...any) {}
func Println(a ...any)               {} // 生产环境：空函数（无操作）
func Errorf(format string, a ...any) {}
