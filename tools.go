//go:build tools
// +build tools

// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	_ "github.com/a-h/templ/cmd/templ"
	_ "golang.org/x/tools/cmd/stringer"
)
