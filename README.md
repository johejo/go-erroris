# go-erroris

[![ci](https://github.com/johejo/go-erroris/workflows/ci/badge.svg?branch=main)](https://github.com/johejo/go-erroris/actions?query=workflow%3Aci)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/johejo/go-erroris)](https://pkg.go.dev/github.com/johejo/go-erroris)
[![codecov](https://codecov.io/gh/johejo/go-erroris/branch/main/graph/badge.svg)](https://codecov.io/gh/johejo/go-erroris)
[![Go Report Card](https://goreportcard.com/badge/github.com/johejo/go-erroris)](https://goreportcard.com/report/github.com/johejo/go-erroris)
[![License](https://img.shields.io/github/license/johejo/go-erroris)](https://github.com/johejo/go-erroris/blob/main/LICENSE)

Package erroris provides error matching functions using errors.Is.

## Example

```go
package erroris_test

import (
	"errors"
	"fmt"

	"github.com/johejo/go-erroris"
)

var (
	fooErr  = errors.New("foo")
	barErr  = errors.New("bar")
	buzzErr = errors.New("buzz")
)

func Example() {
	fn := func() error {
		return fooErr
	}
	err := fn()
	if erroris.OneOf(err, fooErr, barErr) {
		fmt.Println("match")
	}
	if !erroris.OneOf(err, barErr, buzzErr) {
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
```

## License

MIT

## Author

Mitsuo Heijo (@johejo)
