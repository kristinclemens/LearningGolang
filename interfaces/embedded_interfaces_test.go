package interfaces

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFoobar_FooPanicsWithoutImplementation(t *testing.T) {
	foobar := Foobar{}
	// This panics with an invalid memory address or nil pointer dereference error because Foo is not implemented
	// We're allowed to call Foo because the Foobarable interface is embedded in Foobar
	assert.Panics(t, func() {
		foobar.Foo()
	})
}

func TestFoobar_ImplementsFoobarable(t *testing.T) {
	typeOfFoobar := reflect.TypeOf((*Foobar)(nil))
	typeOfFoobarable := reflect.TypeOf((*Foobarable)(nil)).Elem()
	assert.True(t, typeOfFoobar.Implements(typeOfFoobarable))
}

func TestFoobar_ImplementsFooable(t *testing.T) {
	typeOfFoobar := reflect.TypeOf((*Foobar)(nil))
	typeOfFooable := reflect.TypeOf((*Fooable)(nil)).Elem()
	assert.True(t, typeOfFoobar.Implements(typeOfFooable))
}

func TestFoobar_ImplementsBarable(t *testing.T) {
	typeOfFoobar := reflect.TypeOf((*Foobar)(nil))
	typeOfBarable := reflect.TypeOf((*Barable)(nil)).Elem()
	assert.True(t, typeOfFoobar.Implements(typeOfBarable))
}
