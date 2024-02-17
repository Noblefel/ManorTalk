package user

import (
	"database/sql"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/google/uuid"
)

func (s *userService) UpdateProfile(payload models.UpdateProfileInput, username string, authId int) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", ErrNoUser
		}

		return "", fmt.Errorf("error getting user by slug: %w", err)
	}

	if authId != user.Id {
		return "", ErrUnauthorized
	}

	user.Name = payload.Name
	user.Username = payload.Username
	user.Bio = payload.Bio

	file, ok := payload.Files["avatar"]
	if ok {
		f, err := file[0].Open()
		if err != nil {
			return "", err
		}
		defer f.Close()

		if file[0].Size > 2*1024*1024 {
			return "", ErrAvatarTooLarge
		}

		buff := make([]byte, 512)
		_, err = f.Read(buff)
		if err != nil {
			return "", err
		}

		fileType := http.DetectContentType(buff)
		switch fileType {
		case "image/png", "image/jpg", "image/jpeg":
		default:
			return "", ErrAvatarInvalid
		}

		_, err = f.Seek(0, 0)
		if err != nil {
			return "", err
		}

		user.Avatar = fmt.Sprintf(
			"%d%d-%s%s",
			time.Now().UnixNano(),
			user.Id,
			uuid.New(),
			filepath.Ext(file[0].Filename),
		)

		out, err := os.Create("images/avatar/" + user.Avatar)
		if err != nil {
			return "", err
		}
		defer out.Close()

		if fileType == "image/png" {
			img, err := png.Decode(f)
			if err != nil {
				return "", err
			}

			err = png.Encode(out, img)
			if err != nil {
				return "", err
			}
		} else {
			img, err := jpeg.Decode(f)
			if err != nil {
				return "", err
			}

			err = jpeg.Encode(out, img, nil)
			if err != nil {
				return "", err
			}
		}
	}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return "", ErrDuplicateUsername
		}

		return "", fmt.Errorf("error updating user: %w", err)
	}

	return user.Avatar, nil
}

func (s *mockUserService) UpdateProfile(payload models.UpdateProfileInput, username string, authId int) (string, error) {
	switch username {
	case ErrNoUser.Error():
		return "", ErrNoUser
	case ErrUnauthorized.Error():
		return "", ErrUnauthorized
	case ErrAvatarTooLarge.Error():
		return "", ErrAvatarTooLarge
	case ErrAvatarInvalid.Error():
		return "", ErrAvatarInvalid
	case ErrDuplicateUsername.Error():
		return "", ErrDuplicateUsername
	case "unexpected error":
		return "", errors.New("unexpected error")
	default:
		return "", nil
	}
}
