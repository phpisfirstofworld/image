package image

import (
	"github.com/nfnt/resize"
	"github.com/phpisfirstofworld/size"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	image2 "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type image struct {
	path string
	size *size.Size
}

func NewImage() *image {

	return &image{}
}

func (i *image) GetSize() *size.Size {

	return i.size
}

func (i *image) LoadImage(path string) (*resource, error) {

	sourceFile, err := os.Open(path)

	if err != nil {

		return nil, errors.WithStack(err)
	}

	i.path = path

	info, infoErr := sourceFile.Stat()

	if infoErr != nil {

		return nil, errors.WithStack(infoErr)
	}

	i.size = size.NewSize(info.Size())

	defer sourceFile.Close()

	b, readErr := ioutil.ReadAll(sourceFile)

	if readErr != nil {

		return nil, errors.WithStack(readErr)
	}

	types := http.DetectContentType(b)

	var sourceImage image2.Image

	switch types {

	case "image/jpeg":

		sourceImage, err = jpeg.Decode(strings.NewReader(string(b)))

		if err != nil {

			return nil, errors.WithStack(err)
		}

		break

	case "image/jpg":

		sourceImage, err = jpeg.Decode(strings.NewReader(string(b)))

		if err != nil {

			return nil, errors.WithStack(err)
		}

		break

	case "image/png":

		sourceImage, err = png.Decode(strings.NewReader(string(b)))

		if err != nil {

			return nil, errors.WithStack(err)
		}

		break

	case "image/gif":

		sourceImage, err = gif.Decode(strings.NewReader(string(b)))

		if err != nil {

			return nil, errors.WithStack(err)
		}

		break

	default:

		return nil, errors.WithStack(errors.New("暂不支持(" + types + ")此格式"))

	}

	return NewResource(sourceImage, sourceImage, i), nil

}

type resource struct {
	sourceImageResource image2.Image
	dealImageResource   image2.Image
	error               error
	image               *image
}

func NewResource(sourceImageResource image2.Image, dealImageResource image2.Image, image *image) *resource {

	return &resource{sourceImageResource: sourceImageResource, dealImageResource: dealImageResource, error: nil, image: image}
}

func (r *resource) GetSourceImageResource() image2.Image {

	return r.sourceImageResource
}

func (r *resource) ResizePercent(percent int) *resource {

	if !(percent >= 1 && percent <= 100) {

		r.error = errors.New("图片缩放要大于1并且小于100")

		return r
	}

	sourceWidth := cast.ToFloat64(r.sourceImageResource.Bounds().Size().X)

	resizeWidth := cast.ToInt(sourceWidth * (cast.ToFloat64(percent) / 100.0))

	r.dealImageResource = resize.Resize(uint(resizeWidth), 0, r.dealImageResource, resize.Lanczos3)

	return r
}

func (r *resource) ResizeWidth(width int) *resource {

	r.dealImageResource = resize.Resize(uint(width), 0, r.dealImageResource, resize.Lanczos3)

	return r

}

func (r *resource) ResizeHeight(height int) *resource {

	r.dealImageResource = resize.Resize(0, uint(height), r.dealImageResource, resize.Lanczos3)

	return r

}

func (r *resource) Save(path string) error {

	if r.error != nil {

		return errors.WithStack(r.error)
	}

	out, err := os.Create(path)

	if err != nil {

		return err
	}

	defer out.Close()

	return errors.WithStack(png.Encode(out, r.dealImageResource))

}

// OverSave 覆盖保存
func (r *resource) OverSave() error {

	if r.error != nil {

		return errors.WithStack(r.error)
	}

	out, err := os.Create(r.image.path + ".temp")

	if err != nil {

		return errors.WithStack(err)
	}

	err = errors.WithStack(png.Encode(out, r.dealImageResource))

	if err != nil {

		out.Close()

		os.Remove(r.image.path + ".temp")

		return err

	}

	out.Close()

	os.Remove(r.image.path)

	os.Rename(r.image.path+".temp", r.image.path)

	return err
}
