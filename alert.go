// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate stringer -type=Alert -linecomment -output alert_generated.go
package main

const (
	AlertInfo    Alert = iota // alert-info
	AlertSuccess              // alert-success
	AlertWarning              // alert-warning
	AlertError                // alert-error
)

type Alert int
