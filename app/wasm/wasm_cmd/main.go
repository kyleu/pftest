//go:build js
// +build js

package main

import "github.com/kyleu/pftest/app/wasm"

func main() {
	w := wasm.NewWASM()
	<-w.CloseCh
}
