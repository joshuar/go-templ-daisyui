// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T) T
