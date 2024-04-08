// Package foo - Content managed by Project Forge, see [projectforge.md] for details.
package foo

var (
	FooA = Foo{Key: "a"}
	FooB = Foo{Key: "b"}
	FooC = Foo{Key: "c"}
	FooD = Foo{Key: "d"}

	AllFoos = Foos{FooA, FooB, FooC, FooD}
)

type Foo string

const (
	FooA Foo = "a"
	FooB Foo = "b"
	FooC Foo = "c"
	FooD Foo = "d"
)
