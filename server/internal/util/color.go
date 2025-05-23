package util

import (
	"fmt"
	"image/color"
)

func ColorToString(color color.Color) string {
	r, g, b, _ := color.RGBA()
	return fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
}
