# bmpwriter

This little program is just there to experiment with writing files byte by byte using go. It writes the header and data of a simple bitmap file. As a reference [wikipedia](https://en.wikipedia.org/wiki/BMP_file_format) was used.

The header struct looks like this:

```go
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
```

It is written by the `binary` package directly to disk:

```go
binary.Write(f, binary.LittleEndian, header)
binary.Write(f, binary.LittleEndian, data)
```

Please note that I tried to avoid the use of the `unsafe` package, so no `SizeOf` was available to me. Therefore lines like this exist:

```go
header.BitmapOffset = 40 + 14
```

The numbers are based on the types included in the header and have been counted by hand. This is not such a big problem in this case as the header spec for the bitmap format won't change. Even the standard library does it like this. See `writer.go` in the `image` package: [writer.go](https://github.com/golang/image/blob/master/bmp/writer.go)

The program will write a `test.bmp`file with the result into the current working directory.

