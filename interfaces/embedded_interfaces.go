package interfaces

// Fooable can Foo
type Fooable interface {
	Foo() string
}

// Barable can Bar
type Barable interface {
	Bar() string
}

// Foobarable is the union of Fooable and Barable.
type Foobarable interface {
	Fooable
	Barable
}

// Foobar implements Foobarable
type Foobar struct {
	Foobarable
}

// What happens without implemented methods?
// Calling Foo or Bar on a Foobar object causes a panic with an invalid memory address or nil pointer dereference error
