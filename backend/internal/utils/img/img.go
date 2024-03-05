package img

import (
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

var ErrTooLarge = errors.New("too large")
var ErrType = errors.New("invalid type")
var max = 2 * 1024 * 1024

// Verify will check the file's size and return the extension
func Verify(r io.ReadSeeker, name string) (string, error) {
	buff := make([]byte, max+1)
	n, err := r.Read(buff)
	if err != nil {
		return "", err
	}

	if n > max {
		return "", ErrTooLarge
	}

	_, err = r.Seek(0, 0)
	if err != nil {
		return "", err
	}

	fileType, err := checkType(buff)
	if err != nil {
		return "", err
	}

	ext := "." + strings.Split(fileType, "/")[1]
	return ext, nil
}

// checkType will check if the file type is PNG or JPG
func checkType(buff []byte) (string, error) {
	fileType := http.DetectContentType(buff)
	switch fileType {
	case "image/png", "image/jpg", "image/jpeg":
		return fileType, nil
	default:
		return "", ErrType
	}
}

// Save will store the image locally
func Save(r io.ReadSeeker, dir, name string) error {
	fmt.Println(dir + name)
	out, err := os.Create(dir + name)
	if err != nil {
		return err
	}
	defer out.Close()

	if strings.HasSuffix(name, "png") {
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
