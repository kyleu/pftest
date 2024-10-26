//go:build js
// +build js

package main

import "github.com/kyleu/pftest/app/wasm"

func main() {
	w, err := wasm.NewWASM()
	if err != nil {
		panic(err)
	}
	err = w.Run()
	if err != nil {
		panic(err)
	}
}
