package blah

func Foo() int {
	return 42
}

func Bar() int {
	return 432
}

var G int

func Baz() int {
	println(Foo(), Bar())
//line p.go:98989
	if G == 101 {
		return 3
	}
	println(Foo() + Bar())
	return 432
}

//line bar.go 101
