package screens

import (
	"github.com/kyleu/pftest/app/controller/tui/mvc"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/util"
)

func Bootstrap(_ *mvc.State) *Registry {
	reg := NewRegistry()

	reg.AddScreen(NewMainMenuScreen(reg))
	docsScreenItem := &menu.Item{Key: KeyDocs, Title: "Documentation", Description: "Browse embedded markdown documentation", Icon: "book", Route: KeyDocs}
	reg.Register(docsScreenItem, NewDocumentationScreen())
	settingsScreenItem := &menu.Item{Key: KeySettings, Title: "Settings", Description: "Runtime and diagnostics", Icon: "settings", Route: KeySettings}
	reg.Register(settingsScreenItem, NewSettingsScreen())
	aboutScreenItem := &menu.Item{Key: KeyAbout, Title: "About", Description: "Information about " + util.AppName, Icon: "info", Route: KeyAbout}
	reg.Register(aboutScreenItem, NewAboutScreen())

	// reg.AddScreen(SomeNewScreen())

	return reg
}
