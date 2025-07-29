package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type noScrollBarTheme struct {
	fyne.Theme
}

func (n noScrollBarTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameScrollBar {
		return color.NRGBA{0, 0, 0, 0}
	}
	return n.Theme.Color(name, variant)
}
