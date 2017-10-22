package main

import (
	"encoding/binary"
	"math/rand"
	"os"
)

// BitmapHeader represents the header fields needed for the bitmap format
type BitmapHeader struct {
	FileType                [2]byte
	FileSize                uint32
	Reserved1               uint16
	Reserved2               uint16
	BitmapOffset            uint32
	SizeOfBitmapHeader      uint32
	PixelWidth              uint32
	PixelHeight             uint32
	Planes                  uint16
	BitsPerPixel            uint16
	Compression             uint32
	ImageSize               uint32
	HorizontalResolution    uint32
	VerticalResolution      uint32
	NumberOfColorsInPalette uint32
	ImportantColors         uint32
}

func main() {
	var ImageWidth uint32 = 1024
	var ImageHeight uint32 = 768

	var x, y uint32

	var numberOfPixels = ImageHeight * ImageWidth
	var data = make([]uint32, numberOfPixels)
	var c = 0
	var pixel, r, g, b uint32
	for y = 0; y < ImageHeight; y++ {
		for x = 0; x < ImageWidth; x++ {
			r = uint32(rand.Int() % 256)
			g = uint32(rand.Int()%256) << 8
			b = uint32(rand.Int()%256) << 16
			pixel = 0xFF000000
			pixel += r
			pixel += g
			pixel += b
			data[c] = pixel
			c++
		}
	}

	header := BitmapHeader{}

	var SizeOfData = uint32(4 * numberOfPixels)

	header.FileType = [2]byte{'B', 'M'}
	header.FileSize = 40 + 14 + SizeOfData
	header.BitmapOffset = 40 + 14
	header.SizeOfBitmapHeader = 40
	header.PixelWidth = ImageWidth
	header.PixelHeight = ImageHeight
	header.Compression = 0
	header.BitsPerPixel = 32
	header.Planes = 1
	header.ImageSize = SizeOfData
	header.NumberOfColorsInPalette = 0
	header.ImportantColors = 0

	var cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(cwd + "/test.bmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	binary.Write(f, binary.LittleEndian, header)
	binary.Write(f, binary.LittleEndian, data)
}
