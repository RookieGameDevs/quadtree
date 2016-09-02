package bmp

import "image"

//go:generate stringer -type=Color
type Color byte

const (
	Black Color = iota
	White
	// only for pixel regions, indicates the region contains both Black and
	// White pixels
	Gray
)

// Bitmap represents a rectangular image with a color depth of 1 bit.
//
// Though every pixel could be contained in 1bit of information, for simplicity
// Bits is a slice of Color (1 byte integer that can take only two values:
// Black and White.
type Bitmap struct {
	Width, Height int     // bitmap dimensions
	Bits          []Color // rectangular color array, mapped to 1D
}

// NewBitmapFromImage creates a Bitmap from an image.
//
// The source image should be black and white, i.e pixels component (RGB)
// should all have the same values. That's why we only check the Red component
// of the source image in order to determine if the pixel is Black or White
func NewBitmapFromImage(img image.Image) *Bitmap {
	minx := img.Bounds().Min.X
	miny := img.Bounds().Min.Y
	maxx := img.Bounds().Max.X
	maxy := img.Bounds().Max.Y

	w := maxx - minx
	h := maxy - miny

	bmp := Bitmap{
		Width:  w,
		Height: h,
		Bits:   make([]Color, w*h),
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r != 0 {
				bmp.Bits[x+w*y] = White
			}
		}
	}

	return &bmp
}
