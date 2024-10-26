//go:build js
// +build js

package main

import "github.com/kyleu/pftest/app/wasm"

func main() {
	w, err := wasm.NewWASM()
	if err != nil {
		panic(err)
	}
	<-w.CloseCh
}
