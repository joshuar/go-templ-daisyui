// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package dropdown

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/display/icon"
	"github.com/joshuar/go-templ-daisyui/modifiers/state"
)

type HelperOption components.Option2[*HelperProps]

type HelperProps struct {
	icon *icon.Props
}

// BuildInfoHelper creates a new Dropdown Info Helper without rendering it. The properties can then
// be modified before finally rendering by calling the Show() method.
func BuildHelper(status state.State, options ...HelperOption) *HelperProps {
	helper := &HelperProps{}

	switch status {
	case state.Info:
		helper.icon = icon.BuildInfoIcon()
	case state.Success:
		helper.icon = icon.BuildSuccessIcon()
	case state.Warning:
		helper.icon = icon.BuildWarningIcon()
	case state.Error:
		helper.icon = icon.BuildErrorIcon()
	}

	for _, option := range options {
		option(helper)
	}

	return helper
}
