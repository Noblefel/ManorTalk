package user

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/img"
	"github.com/google/uuid"
)

func (s *userService) UpdateProfile(payload models.UpdateProfileInput, username string, authId int) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", ErrNoUser
		}

		return "", fmt.Errorf("getting user by slug: %w", err)
	}

	if authId != user.Id {
		return "", ErrUnauthorized
	}

	user.Name = payload.Name
	user.Username = payload.Username
	user.Bio = payload.Bio

	files, ok := payload.Files["avatar"]
	if ok {
		f, err := files[0].Open()
		if err != nil {
			return "", fmt.Errorf("opening file: %w", err)
		}
		defer f.Close()

		user.Avatar = fmt.Sprintf(
			"%d%d-%s%s",
			time.Now().UnixNano(),
			user.Id,
			uuid.New(),
			filepath.Ext(files[0].Filename),
		)

		err = img.Upload(f, "images/avatar/"+user.Avatar)
		if err != nil {
			switch err {
			case img.ErrTooLarge:
				return "", ErrAvatarTooLarge
			case img.ErrType:
				return "", ErrAvatarInvalid
			default:
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
