package image

import (
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	image2 "image"
	"image/png"
	"os"
)

type image struct {
	path string
}

func NewImage() *image {

	return &image{}
}

func (i *image) LoadImage(path string) (*resource, error) {

	sourceFile, err := os.Open(path)

	if err != nil {

		return nil, errors.WithStack(err)
	}

	i.path = path

	defer sourceFile.Close()

	sourceImage, err := png.Decode(sourceFile)

	if err != nil {

		return nil, errors.WithStack(err)
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

func (r *resource) Resize(percent int) *resource {

	if !(percent >= 1 && percent <= 100) {

		r.error = errors.New("图片缩放要大于1并且小于100")

		return r
	}

	sourceWidth := cast.ToFloat64(r.sourceImageResource.Bounds().Size().X)

	resizeWidth := cast.ToInt(sourceWidth * (cast.ToFloat64(percent) / 100.0))

	r.dealImageResource = resize.Resize(uint(resizeWidth), 0, r.dealImageResource, resize.Lanczos3)

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
