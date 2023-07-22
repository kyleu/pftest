// Content managed by Project Forge, see [projectforge.md] for details.
//go:build js

package main

import "syscall/js"

func Audit(msg string, code bool, args ...any) {
	js.Global().Call("audit", append([]any{msg, code}, args...)...)
}
