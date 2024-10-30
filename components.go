// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

type Option[T any] func(T) T

type customisableClass interface {
	AddClasses(class ...string)
	Class() string
}
