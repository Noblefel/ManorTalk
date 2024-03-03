package img

import (
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

var ErrTooLarge = errors.New("too large")
var ErrType = errors.New("invalid type")
var max = 2 * 1024 * 1024

// Upload will verify the file's size and type, then store it locally
func Upload(r io.ReadSeeker, path string) error {
	buff := make([]byte, max+1)
	n, err := r.Read(buff)
	if err != nil {
		return err
	}

	if n > max {
		return ErrTooLarge
	}

	_, err = r.Seek(0, 0)
	if err != nil {
		return err
	}

	fileType, err := checkType(buff)
	if err != nil {
		return err
	}

	if err := saveToLocal(r, fileType, path); err != nil {
		return err
	}

	return nil
}

// checkType will verify if the file type is PNG or JPG
func checkType(buff []byte) (string, error) {
	fileType := http.DetectContentType(buff)
	switch fileType {
	case "image/png", "image/jpg", "image/jpeg":
		return fileType, nil
	default:
		return "", ErrType
	}
}

// saveToLocal will store the image locally with the given path
func saveToLocal(r io.Reader, fileType, name string) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	if fileType == "image/png" {
		img, err := png.Decode(r)
		if err != nil {
			return err
		}

		err = png.Encode(out, img)
		if err != nil {
			return err
		}
	} else {
		img, err := jpeg.Decode(r)
		if err != nil {
			return err
		}

		err = jpeg.Encode(out, img, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
