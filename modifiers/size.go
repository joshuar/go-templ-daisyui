// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package modifiers

import "github.com/a-h/templ"

type Size string

func (s Size) String() string {
	return string(s)
}

func (s Size) Valid() bool {
	return s != ""
}

func (s Size) Class() templ.KeyValue[string, bool] {
	return templ.KV(s.String(), s.Valid())
}
