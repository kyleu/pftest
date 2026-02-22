package registry

import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/tui/screens"
	"github.com/kyleu/pftest/app/controller/tui/screens/settings"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/util"
)

func Bootstrap(st *app.State, logger util.Logger) *screens.Registry {
	reg := screens.NewRegistry()

	reg.AddScreen(screens.NewMainMenuScreen(reg))

	docsScreenItem := &menu.Item{Key: screens.KeyDocs, Title: "Documentation", Description: "Browse markdown documentation", Icon: "book", Route: screens.KeyDocs}
	reg.Register(docsScreenItem, screens.NewDocumentationScreen())
	aboutScreenItem := &menu.Item{Key: screens.KeyAbout, Title: "About", Description: "Information about " + util.AppName, Icon: "info", Route: screens.KeyAbout}
	reg.Register(aboutScreenItem, screens.NewAboutScreen())

	// reg.AddScreen(SomeNewScreen())

	settings.Register(reg)

	return reg
}
