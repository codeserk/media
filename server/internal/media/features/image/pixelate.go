package image

import (
	"image"
	"image/color"
	"math"
	"media/internal/media"
)

func (s *service) Pixelate(img image.Image) (image.Image, error) {
	imgBounds := img.Bounds()
	imgWidth := imgBounds.Max.X
	imgHeight := imgBounds.Max.Y
	widthPixels := 6
	size := int(math.Ceil(float64(imgWidth) / float64(widthPixels)))

	pixelatedImg := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for x := 0; x < imgWidth; x += size {
		for y := 0; y < imgHeight; y += size {
			rect := image.Rect(x, y, x+size, y+size)

			if rect.Max.X > imgWidth {
				rect.Max.X = imgWidth
			}
			if rect.Max.Y > imgHeight {
				rect.Max.Y = imgHeight
			}

			r, g, b := calculateMeanAverageColorWithRect(img, rect, true)
			col := color.RGBA{r, g, b, 255}

			for x2 := rect.Min.X; x2 < rect.Max.X; x2++ {
				for y2 := rect.Min.Y; y2 < rect.Max.Y; y2++ {
					pixelatedImg.Set(x2, y2, col)
				}
			}
		}
	}

	return s.Resize(pixelatedImg, media.ResizeParams{Width: widthPixels, Height: 0, KeepAspectRatio: true, Filter: media.Nearest})
}

func calculateMeanAverageColorWithRect(
	img image.Image,
	rect image.Rectangle,
	useSquaredAverage bool,
) (red, green, blue uint8) {
	var rSum, gSum, bSum, count int

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rSum += int(r)
			gSum += int(g)
			bSum += int(b)
			count++
		}
	}

	if count == 0 {
		return 0, 0, 0
	}

	if useSquaredAverage {
		rSum = int(math.Sqrt(float64(rSum / count)))
		gSum = int(math.Sqrt(float64(gSum / count)))
		bSum = int(math.Sqrt(float64(bSum / count)))
	} else {
		rSum /= count
		gSum /= count
		bSum /= count
	}

	return uint8(rSum), uint8(gSum), uint8(bSum)
}
