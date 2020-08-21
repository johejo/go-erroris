// Package erroris provides error matching functions using errors.Is.
package erroris

import "errors"

// OneOf reports whether one of target errors matches the given err.
func OneOf(err error, targets ...error) bool {
	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

// AllOf reports whether all of target errors matches the given err.
func AllOf(err error, targets ...error) bool {
	for _, t := range targets {
		if !errors.Is(err, t) {
			return false
		}
	}
	return true
}
