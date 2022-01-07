package theme

import (
	"os"
)

var ThemeDefault = func() *Theme {
	nbl := "#beb1d7"
	if o := os.Getenv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#3f384e"
	if o := os.Getenv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key: "default",
		Light: &Colors{
			Border: "1px solid #dddddd", LinkDecoration: "none",
			Foreground: "#000000", ForegroundMuted: "#999999",
			Background: "#ffffff", BackgroundMuted: "#eeeeee",
			LinkForeground: "#2d414e", LinkVisitedForeground: "#406379",
			NavForeground: "#000000", NavBackground: nbl,
			MenuForeground: "#000000", MenuBackground: "#e1e1ff", MenuSelectedBackground: "#cbbffa", MenuSelectedForeground: "#000000",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#FF0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#ffffff", ForegroundMuted: "#999999",
			Background: "#121212", BackgroundMuted: "#333333",
			LinkForeground: "#674f92", LinkVisitedForeground: "#695779",
			NavForeground: "#ffffff", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuBackground: "#171f24", MenuSelectedBackground: "#333333", MenuSelectedForeground: "#ffffff",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#FF0000",
		},
	}
}()
