// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package modifiers

import (
	"slices"
	"strings"

	"github.com/a-h/templ"
)

var validColorNames = []Color{"primary", "secondary", "accent", "neutral", "info", "succes", "warning", "error"}

// Color represents a DaisyUI color name.
//
// https://daisyui.com/docs/colors/#list-of-all-daisyui-color-names
type Color string

func (c Color) String() string {
	return string(c)
}

// Valid returns a boolean indicating whether the given Color is derived from a valid color name.
func (c Color) Valid() bool {
	return c != "" &&
		slices.ContainsFunc(validColorNames,
			func(validColor Color) bool { return strings.Contains(c.String(), string(validColor)) })
}

func (c Color) Class() templ.KeyValue[string, bool] {
	return templ.KV(c.String(), c.Valid())
}
