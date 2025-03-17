// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package validation

type Validation struct {
	required bool
	validate bool
	hint     string
	pattern  string
}

func (v *Validation) SetRequired(value bool) {
	v.required = value
}

func (v *Validation) SetValidation(value bool) {
	v.validate = value
}

func (v *Validation) SetHint(value string) {
	v.hint = value
}

func (v *Validation) SetPattern(value string) {
	v.pattern = value
}

// Validate returns whether validation is enabled.
func (v *Validation) Validate() bool {
	return v.validate
}

// Required returns whether required is enabled.
func (v *Validation) Required() bool {
	return v.validate
}

// Hint returns the validation hint.
func (v *Validation) Hint() string {
	return v.hint
}

// Pattern returns the validation pattern.
func (v *Validation) Pattern() string {
	return v.pattern
}

func New() *Validation {
	return &Validation{
		required: true,
		validate: true,
	}
}
