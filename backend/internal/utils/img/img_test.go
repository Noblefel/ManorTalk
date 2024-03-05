package img

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestVerify(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.png")
		defer os.Remove(f.Name())
		png.Encode(f, image.Rect(0, 0, 1, 1))
		f.Seek(0, 0)

		s, err := Verify(f, "/test.png")
		if err != nil {
			t.Errorf("expecting no error, got %v", err)
		}

		if s != ".png" {
			t.Errorf("want .png, got %s", s)
		}
	})

	t.Run("fail reading", func(t *testing.T) {
		var b bytes.Reader
		_, err := Verify(&b, "")
		if err == nil {
			t.Error("expecting error")
		}
	})

	t.Run("fail too large", func(t *testing.T) {
		b := make([]byte, max+1)
		r := bytes.NewReader(b)

		_, err := Verify(r, "")
		if err == nil {
			t.Error("expecting error")
		}
	})

	t.Run("fail invalid type", func(t *testing.T) {
		b := make([]byte, 1)
		r := bytes.NewReader(b)

		_, err := Verify(r, "")
		if err == nil {
			t.Error("expecting error")
		}
	})
}

func TestCheckType(t *testing.T) {
	t.Run("return image png", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.png")
		defer os.Remove(f.Name())
		png.Encode(f, image.Rect(0, 0, 1, 1))
		f.Seek(0, 0)

		buff := make([]byte, max+1)
		f.Read(buff)

		fileType, err := checkType(buff)
		if err != nil {
			t.Errorf("expecting no error, got %v", err)
		}

		if fileType != "image/png" {
			t.Errorf("got %q, want image/png", fileType)
		}
	})

	t.Run("return image jpeg", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.jpeg")
		defer os.Remove(f.Name())
		jpeg.Encode(f, image.Rect(0, 0, 1, 1), nil)
		f.Seek(0, 0)

		buff := make([]byte, max+1)
		f.Read(buff)

		fileType, err := checkType(buff)
		if err != nil {
			t.Errorf("expecting no error, got %v", err)
		}

		if fileType != "image/jpeg" {
			t.Errorf("got %q, want image/jpeg", fileType)
		}
	})

	t.Run("return invalid type", func(t *testing.T) {
		_, err := checkType(nil)
		if err == nil {
			t.Error("expecting error")
		}
	})
}

func TestSave(t *testing.T) {
	t.Run("fail no path name", func(t *testing.T) {
		err := Save(nil, "", "")
		if err == nil {
			t.Error("expecting error")
		}
	})

	t.Run("saving png", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.png")
		defer os.Remove(f.Name())
		png.Encode(f, image.Rect(0, 0, 1, 1))
		f.Seek(0, 0)

		err := Save(f, "", "/test.png")
		if err != nil {
			t.Errorf("expecting no error, got %v", err)
		}
	})

	t.Run("fail decoding png", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.jpeg")
		defer os.Remove(f.Name())
		jpeg.Encode(f, image.Rect(0, 0, 1, 1), nil)

		err := Save(f, "", "/test.png")
		if err == nil {
			t.Errorf("expecting error")
		}
	})

	t.Run("saving jpeg", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.jpeg")
		defer os.Remove(f.Name())
		jpeg.Encode(f, image.Rect(0, 0, 1, 1), nil)
		f.Seek(0, 0)

		err := Save(f, "", "/test.jpeg")
		if err != nil {
			t.Errorf("expecting no error, got %v", err)
		}
	})

	t.Run("fail decoding jpeg", func(t *testing.T) {
		f, _ := os.CreateTemp("", "test.png")
		defer os.Remove(f.Name())
		png.Encode(f, image.Rect(0, 0, 1, 1))

		err := Save(f, "", "/test.jpeg")
		if err == nil {
			t.Errorf("expecting error")
		}
	})
}
