// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "strings"

type Table struct {
	class    []string
	Zebra    bool
	PinRows  bool
	PinCols  bool
	RowHover bool
}

func (t Table) Class() string {
	return strings.Join(t.class, " ")
}

type TableOption func(Table) Table

// For <table> to show zebra stripe rows.
func WithZebra() TableOption {
	return func(t Table) Table {
		t.Zebra = true
		return t
	}
}

// For <table> to make all the rows inside <thead> and <tfoot> sticky.
func WithPinRows() TableOption {
	return func(t Table) Table {
		t.PinRows = true
		return t
	}
}

// For <table> to make all the <th> columns sticky.
func WithPinCols() TableOption {
	return func(t Table) Table {
		t.PinCols = true
		return t
	}
}

// For <table> with hover on row effect.
func WithRowHover() TableOption {
	return func(t Table) Table {
		t.RowHover = true
		return t
	}
}
