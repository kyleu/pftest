package foo

var AllFoos = Foos{FooA, FooB, FooC, FooD}

type (
	Foo  string
	Foos []Foo
)

const (
	FooA Foo = "a"
	FooB Foo = "b"
	FooC Foo = "c"
	FooD Foo = "d"
)
