package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// BuildSplitView creates a split view that remembers the user-adjusted ratio.
func BuildSplitView(left, right fyne.CanvasObject, splitRatio *float64) fyne.CanvasObject {
	split := container.NewHSplit(
		container.NewPadded(left),
		container.NewPadded(right),
	)
	split.SetOffset(*splitRatio)

	// Continuously track user-adjusted ratio
	go func() {
		prev := split.Offset
		for {
			current := split.Offset
			if current != prev {
				*splitRatio = current
				prev = current
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return split
}
