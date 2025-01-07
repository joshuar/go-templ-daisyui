// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	AlertInfo    AlertType = iota // alert-info
	AlertSuccess                  // alert-success
	AlertWarning                  // alert-warning
	AlertError                    // alert-error
)

// AlertType represents the type of alert. These are the standard state values
// (info, success, warning and error).
type AlertType int
