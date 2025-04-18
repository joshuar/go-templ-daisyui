// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package modifiers

import "github.com/a-h/templ"

type Style string

func (s Style) String() string {
	return string(s)
}

func (s Style) Valid() bool {
	return s != ""
}

func (s Style) Class() templ.KeyValue[string, bool] {
	return templ.KV(s.String(), s.Valid())
}
