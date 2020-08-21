package erroris_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/johejo/go-erroris"
)

func TestOneOf(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err3 := errors.New("err3")
	err4 := errors.New("err4")
	errs := []error{err1, err2, err3}
	for _, err := range errs {
		if !erroris.OneOf(err, errs...) {
			t.Error("err shoule be one of errs")
		}
		if erroris.OneOf(err4, errs...) {
			t.Error("err4 is not one of errs")
		}
	}
}

func TestAllOf(t *testing.T) {
	parent := errors.New("parent")
	child := fmt.Errorf("child wraps parent: %w", parent)
	grandChild := fmt.Errorf("grandChild wraps child: %w", child)
	other := errors.New("other")

	if !erroris.AllOf(grandChild, child, parent) {
		t.Error("grandChild matches child and parent")
	}
	if erroris.AllOf(child, parent, other) {
		t.Error("child should not matches child and other")
	}
	if erroris.AllOf(other, child, parent) {
		t.Error("other should not match child and parent")
	}
}
