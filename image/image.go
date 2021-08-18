package image

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/tikivn/tially/pkg/util/logaff"
)

func SaveImage(imgByte []byte, imgName string, imgFormat string) error {
	logger := logaff.GetNewLogger()

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		logger.Error(err)
		return err
	}

	out, _ := os.Create(fmt.Sprintf("./%s", imgName))
	defer out.Close()

	switch imgFormat {
	case "png":
		err = png.Encode(out, img)
		if err != nil {
			logger.Error(err)
			return err
		}
	case "jpg", "jpeg":
		var opts jpeg.Options
		opts.Quality = 1

		err = jpeg.Encode(out, img, &opts)
		if err != nil {
			logger.Error(err)
			return err
		}
	default:
		logger.Errorf("only support `png,jpg,jpeg` format")
		return errors.New("only support `png,jpg,jpeg` format")
	}

	return nil
}
