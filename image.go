package zzblog

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"os"
)

type ImageReader interface {
	io.Reader
	io.Seeker
}

const (
	ImagePng = iota
	ImageJpg
	ImageGif
	ImageBmp
	ImageUnknown
)

type Image struct {
	Id            string
	Width, Height uint
	Format        int
	Pathfile      string
}

func (img *Image) Write(w io.Writer, width, height uint) error {
	var bitmap image.Image
	f, err := os.Open(img.Pathfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var buf bytes.Buffer
	if img.Format == ImageGif {
		io.Copy(&buf, f)
		return nil
	}
	if bitmap, err = decodeImage(img.Format, f); err != nil {
		panic(err)
	}
	width, height = size(uint(img.Width), uint(img.Height), width, height, img.Format)
	resized := resize.Resize(width, height, bitmap, resize.NearestNeighbor)
	if err = encodeImage(w, img.Format, resized); err != nil {
		return err
	}
	return nil
}

func (img *Image) MimeType() string {
	return mime.TypeByExtension(getImageExt(img.Format))
}

func size(ow, oh, w, h uint, format int) (width, height uint) {
	if format == ImageGif {
		width = ow
		height = oh
		return
	}
	if w > ow {
		w = ow
	}
	if h > oh {
		h = oh
	}
	if w != 0 {
		if mh := w * oh / ow; h > mh {
			h = mh
		}
	}
	if h != 0 {
		if mw := h * ow / oh; w > mw {
			w = mw
		}
	}
	width = w
	height = h
	return
}

func decodeImage(format int, reader ImageReader) (img image.Image, err error) {
	reader.Seek(0, io.SeekStart)
	switch format {
	case ImageJpg:
		img, err = jpeg.Decode(reader)
	case ImagePng:
		img, err = png.Decode(reader)
	case ImageBmp:
		img, err = bmp.Decode(reader)
	case ImageGif:
		img, err = gif.Decode(reader)
	default:
		err = errors.New("unknown format")
	}
	return
}

func encodeImage(writer io.Writer, format int, img image.Image) error {
	switch format {
	case ImagePng:
		png.Encode(writer, img)
	case ImageGif:
		gif.Encode(writer, img, nil)
	case ImageJpg:
		jpeg.Encode(writer, img, nil)
	case ImageBmp:
		bmp.Encode(writer, img)
	default:
		return errors.New("unknown format")
	}
	return nil
}

func getImageExt(format int) string {
	switch format {
	case ImagePng:
		return ".png"
	case ImageJpg:
		return ".jpg"
	case ImageGif:
		return ".gif"
	case ImageBmp:
		return ".bmp"
	default:
		return ""
	}
}

func getImageId(r ImageReader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	md5Sumb := md5.Sum(buf.Bytes())
	return hex.EncodeToString(md5Sumb[:])
}

func getImageFormat(r ImageReader) int {
	bytes := make([]byte, 4)
	r.Seek(0, io.SeekStart)
	n, _ := r.Read(bytes)
	if n < 4 {
		return ImagePng
	}
	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return ImagePng
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return ImageJpg
	}
	if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
		return ImageGif
	}
	if bytes[0] == 0x42 && bytes[1] == 0x4D {
		return ImageBmp
	}
	return ImageUnknown
}

func getImageConfig(format int, r ImageReader) (conf image.Config, err error) {
	r.Seek(0, io.SeekStart)
	switch format {
	case ImageJpg:
		conf, err = jpeg.DecodeConfig(r)
	case ImagePng:
		conf, err = png.DecodeConfig(r)
	case ImageBmp:
		conf, err = bmp.DecodeConfig(r)
	case ImageGif:
		conf, err = gif.DecodeConfig(r)
	default:
		err = errors.New("unknown format")
	}
	return
}

func getImageFromFile(f ImageReader) (img *Image, err error) {
	img = new(Image)
	img.Id = getImageId(f)
	img.Format = getImageFormat(f)
	if img.Format == ImageUnknown {
		err = errors.New("unknown format")
		return
	}
	conf, e := getImageConfig(img.Format, f)
	if e != nil {
		err = e
		return
	}
	img.Height = uint(conf.Height)
	img.Width = uint(conf.Width)
	return
}

func getImage(file string) (img *Image, err error) {
	f, e := os.Open(file)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	img, err = getImageFromFile(f)
	if err != nil {
		img.Pathfile = file
	}
	return
}
