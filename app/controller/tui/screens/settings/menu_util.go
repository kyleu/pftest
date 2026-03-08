package settings

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/kyleu/pftest/app/controller/tui/components"
	"github.com/kyleu/pftest/app/controller/tui/layout"
	"github.com/kyleu/pftest/app/controller/tui/screens"
	"github.com/kyleu/pftest/app/controller/tui/style"
	"github.com/kyleu/pftest/app/lib/menu"
)

func menuClamp(cursor int, count int) int {
	return screens.MenuClampCursor(cursor, count)
}

func menuDelta(msg tea.Msg) (int, bool) {
	return screens.MenuMoveDelta(msg)
}

func menuWindow(items menu.Items, cursor int, height int) (menu.Items, int) {
	if len(items) == 0 || height <= 0 || len(items) <= height {
		return items, menuClamp(cursor, len(items))
	}
	cursor = menuClamp(cursor, len(items))
	start := max(cursor-(height/2), 0)
	maxStart := len(items) - height
	if start > maxStart {
		start = maxStart
	}
	end := start + height
	return items[start:end], cursor - start
}

func renderMenuBody(items menu.Items, cursor int, st style.Styles, rects layout.Rects) string {
	w, h, _ := panelDimensions(st.Panel, rects)
	winItems, winCursor := menuWindow(items, cursor, h)
	return components.RenderMenuList(winItems, winCursor, st, w)
}
