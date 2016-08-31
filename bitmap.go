package quadtree

import "image"

// Bitmap is an rectangular bit array
type Bitmap struct {
	Width, Height int    // bitmap dimensions
	Bits          []byte // bit array
}

// NewBitmapFromImage creates a Bitmap from an image.
//
// Only the Red component is checked on the source image, as sources images
// should be black and white (1 bit)
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
		Bits:   make([]byte, w*h),
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r != 0 {
				bmp.Bits[x+w*y] = 1
			}
		}
	}

	return &bmp
}

// allZeroes returns true if all bits are 0 in a given rectangle
func (bmp *Bitmap) allZeroes(rect image.Rectangle) bool {

	// naive implementation
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if bmp.Bits[x+bmp.Width] != 0 {
				// immediately returns at the first 1 found
				return false
			}
		}
	}
	return true
}
