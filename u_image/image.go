package u_image

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
)

func SaveImage(data []byte, fileName string) error {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	out, _ := os.Create(fmt.Sprintf("./%s", fileName))
	defer out.Close()

	switch path.Ext(fileName) {
	case ".png":
		err = png.Encode(out, img)
		if err != nil {
			return err
		}
	case ".jpg", ".jpeg":
		var opts jpeg.Options
		opts.Quality = 1

		err = jpeg.Encode(out, img, &opts)
		if err != nil {
			return err
		}
	default:
		return errors.New("only support `png,jpg,jpeg` format")
	}

	return nil
}
