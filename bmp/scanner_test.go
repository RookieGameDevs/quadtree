package bmp

import (
	"image"
	"testing"
)

func TestNaiveScannerIsWhite(t *testing.T) {
	ss := []string{
		"000",
		"100",
		"011",
	}

	var testTbl = []struct {
		minx, miny, maxx, maxy int
		expected               bool
	}{
		{0, 0, 2, 2, false},
		{1, 1, 2, 2, false},
		{0, 1, 0, 1, true},
		{0, 0, 0, 0, false},
		{1, 0, 1, 0, false},
		{1, 0, 2, 1, false},
		{1, 2, 2, 2, true},
	}

	bmp := NewBitmapFromStrings(ss)
	var scanner NaiveScanner

	for _, tt := range testTbl {
		actual := scanner.IsWhite(bmp, image.Point{tt.minx, tt.miny}, image.Point{tt.maxx, tt.maxy})
		if actual != tt.expected {
			t.Errorf("NaiveScanner.IsWhite() (%d,%d|%d,%d): expected %v, actual %v", tt.minx, tt.miny, tt.maxx, tt.maxy, tt.expected, actual)
		}
	}
}

func TestNaiveScannerIsBlack(t *testing.T) {
	ss := []string{
		"111",
		"011",
		"100",
	}

	var testTbl = []struct {
		minx, miny, maxx, maxy int
		expected               bool
	}{
		{0, 0, 2, 2, false},
		{1, 1, 2, 2, false},
		{0, 1, 0, 1, true},
		{0, 0, 0, 0, false},
		{1, 0, 1, 0, false},
		{1, 0, 2, 1, false},
		{1, 2, 2, 2, true},
		{2, 2, 2, 2, true},
	}

	bmp := NewBitmapFromStrings(ss)
	var scanner NaiveScanner

	for _, tt := range testTbl {
		actual := scanner.IsBlack(bmp, image.Point{tt.minx, tt.miny}, image.Point{tt.maxx, tt.maxy})
		if actual != tt.expected {
			t.Errorf("NaiveScanner.IsBlack() (%d,%d|%d,%d): expected %v, actual %v", tt.minx, tt.miny, tt.maxx, tt.maxy, tt.expected, actual)
		}
	}
}

func TestNaiveScannerFilled(t *testing.T) {
	ss := []string{
		"111",
		"011",
		"100",
	}
	var testTbl = []struct {
		minx, miny, maxx, maxy int
		expected               Color
	}{
		{0, 0, 2, 2, Gray},
		{1, 1, 2, 2, Gray},
		{0, 1, 0, 1, Black},
		{0, 0, 0, 0, White},
		{1, 0, 1, 0, White},
		{1, 0, 2, 1, White},
	}

	bmp := NewBitmapFromStrings(ss)
	var scanner NaiveScanner

	for _, tt := range testTbl {
		actual := scanner.IsFilled(bmp, image.Point{tt.minx, tt.miny}, image.Point{tt.maxx, tt.maxy})
		if actual != tt.expected {
			t.Errorf("NaiveScanner.IsFilled() (%d,%d|%d,%d): expected %v, actual %v", tt.minx, tt.miny, tt.maxx, tt.maxy, tt.expected, actual)
		}
	}
}
