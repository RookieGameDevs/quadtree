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
// should be black (all color components at 0) or white (all color components
// are 1)
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
func (bmp *Bitmap) allZeroes(topLeft, bottomRight image.Point) bool {

	// naive implementation
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			if bmp.Bits[x+bmp.Width*y] != 0 {
				// immediately returns at the first 1 found
				return false
			}
		}
	}
	return true
}
