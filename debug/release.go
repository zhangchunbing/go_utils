//go:build !debug

package debug

func Fprint(w io.Writer, a ...any) {}
func Print(a ...any) {}
func Println(a ...any) {} // 生产环境：空函数（无操作）
func Printf(format string, a ...any) { }