// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package validation

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Validation struct {
	required bool
	validate bool
	hint     string
	*attributes.Attributes
}

func (v *Validation) SetRequired(value bool) {
	v.validate = true
	v.required = value
}

func (v *Validation) SetValidation(value bool) {
	v.validate = value
}

func (v *Validation) SetHint(value string) {
	v.hint = value
}

func (v *Validation) SetPattern(value string) {
	v.SetAttribute("pattern", value)
}

func (v *Validation) SetMinLength(minLength int) {
	v.SetAttribute("minlength", strconv.Itoa(minLength))
}

func (v *Validation) SetMaxLength(maxLength int) {
	v.SetAttribute("maxlength", strconv.Itoa(maxLength))
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
	return v.GetAttribute("pattern")
}

func (v *Validation) ValidationOptions() templ.Attributes {
	return v.ShowAttributes()
}

func New() *Validation {
	return &Validation{
		required:   false,
		validate:   false,
		Attributes: attributes.New(),
	}
}
