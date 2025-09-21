package theme

import (
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/util"
)

const (
	white, black = "#ffffff", "#000000"
	threshold    = (65535 * 3) / 2
)

var Default = func() *Theme {
	nbl := "#b3aac5"
	if o := util.GetEnv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#2b2439"
	if o := util.GetEnv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key:  "default",
		Base: "#beb1d7",
		Light: &Colors{
			Border: "1px solid #cccccc", LinkDecoration: "none",
			Foreground: "#000000", ForegroundMuted: "#3f3a48",
			Background: "#ffffff", BackgroundMuted: "#f0edf3",
			LinkForeground: "#241e2f", LinkVisitedForeground: "#241e2f",
			NavForeground: "#2a2a2a", NavBackground: nbl,
			MenuForeground: "#000000", MenuSelectedForeground: "#000000",
			MenuBackground: "#e0dce7", MenuSelectedBackground: "#c2bad0",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #444444", LinkDecoration: "none",
			Foreground: "#dddddd", ForegroundMuted: "#a5a0b0",
			Background: "#121212", BackgroundMuted: "#1d1925",
			LinkForeground: "#e0dce7", LinkVisitedForeground: "#b3aac5",
			NavForeground: "#f8f9fa", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuSelectedForeground: "#dddddd",
			MenuBackground: "#241e2f", MenuSelectedBackground: "#6e657e",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()

func TextColorFor(clr string) string {
	c, err := ParseHexColor(clr)
	if err != nil {
		return white
	}
	r, g, b, _ := c.RGBA()
	total := r + g + b
	if total < threshold {
		return white
	}
	return black
}

func ParseHexColor(s string) (color.RGBA, error) {
	ret := color.RGBA{A: 0xff}
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &ret.R, &ret.G, &ret.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &ret.R, &ret.G, &ret.B)
		// Double the hex digits:
		ret.R *= 17
		ret.G *= 17
		ret.B *= 17
	default:
		err = errors.Errorf("invalid length [%d], must be 7 or 4", len(s))
	}
	return ret, err
}
