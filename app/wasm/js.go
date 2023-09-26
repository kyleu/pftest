//go:build js

// Package wasm - Content managed by Project Forge, see [projectforge.md] for details.
package main

import "syscall/js"

func Audit(msg string, codes ...any) {
	js.Global().Call("audit", append([]any{msg}, codes...)...)
}
