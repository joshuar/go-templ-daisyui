// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"strings"

	"github.com/a-h/templ"
)

type Tabs struct {
	class  []string
	panels []TabPanel
}

// TabPanel represents a panel for a tab. It has a label displayed on its tab
// and some content displayed in the panel.
type TabPanel struct {
	Content templ.Component
	Label   string
}

func (t Tabs) Class() string {
	return strings.Join(t.class, " ")
}

type TabsOption Option[Tabs]

// Add borders to tabs.
func WithTabBorders() TabsOption {
	return func(t Tabs) Tabs {
		t.class = append(t.class, "tabs-bordered")
		return t
	}
}

// Add lift to tabs.
func WithTabLift() TabsOption {
	return func(t Tabs) Tabs {
		t.class = append(t.class, "tabs-lifted")
		return t
	}
}

// Add boxes to tabs.
func WithTabBox() TabsOption {
	return func(t Tabs) Tabs {
		t.class = append(t.class, "tabs-boxed")
		return t
	}
}

// Add boxes to tabs.
func WithPanels(panels ...TabPanel) TabsOption {
	return func(t Tabs) Tabs {
		t.panels = append(t.panels, panels...)
		return t
	}
}

func NewTabs(options ...TabsOption) Tabs {
	tabs := Tabs{
		class: []string{"tabs"},
	}

	for _, option := range options {
		tabs = option(tabs)
	}

	return tabs
}

func NewTabPanel(label string, content templ.Component) TabPanel {
	return TabPanel{
		Label:   label,
		Content: content,
	}
}
