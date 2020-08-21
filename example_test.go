package erroris_test

import (
	"errors"
	"fmt"

	"github.com/johejo/go-erroris"
)

var (
	errFoo  = errors.New("foo")
	errBar  = errors.New("bar")
	errBuzz = errors.New("buzz")
)

func Example() {
	fn := func() error {
		return errFoo
	}
	err := fn()
	if erroris.OneOf(err, errFoo, errBar) {
		fmt.Println("match")
	}
	if !erroris.OneOf(err, errBar, errBuzz) {
		fmt.Println("does not match")
	}

	parent := errors.New("parent")
	child := fmt.Errorf("child wraps parent: %w", parent)
	grandChild := fmt.Errorf("grandChild wraps child: %w", child)
	other := errors.New("other")

	if !erroris.AllOf(grandChild, child, parent) {
		panic("grandChild matches child and parent")
	}
	if erroris.AllOf(child, parent, other) {
		panic("child should not matches child and other")
	}
	if erroris.AllOf(other, child, parent) {
		panic("other should not match child and parent")
	}

	// Output:
	// match
	// does not match
}
