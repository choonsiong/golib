package golib

// Foo is a function.
// Check out [http://example.com]
func Foo() {
	// I am foo
}

// Bar is a function following the specification
// defined in [PROTO 000]
//
// [PROTO 000]: http://example.com
func Bar() {
	// I am bar
}

// Person is a person.
// To instantiate, refer to [golib.PersonFactory]
type Person struct {
	Name string
}

// PersonFactory will instantiate a new Person object.
func PersonFactory() Person {
	return Person{"Foo"}
}
