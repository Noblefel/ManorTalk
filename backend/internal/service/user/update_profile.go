package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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

	var oldImage string

	if payload.Avatar != nil {
		name := fmt.Sprintf("%s-%d", uuid.New(), authId)
		ext, err := img.Verify(payload.Avatar, name)
		if err != nil {
			switch err {
			case img.ErrTooLarge:
				return "", ErrAvatarTooLarge
			case img.ErrType:
				return "", ErrAvatarInvalid
			default:
				return "", fmt.Errorf("verifying image: %w", err)
			}
		}

		oldImage, user.Avatar = user.Avatar, name+ext

		err = img.Save(payload.Avatar, "images/avatar/", user.Avatar)
		if err != nil {
			return "", fmt.Errorf("saving image: %w", err)
		}
	}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return "", ErrDuplicateUsername
		}

		return "", fmt.Errorf("updating user: %w", err)
	}

	if oldImage != "" {
		if err := os.Remove("images/avatar/" + oldImage); err != nil {
			log.Println("unable to delete image: ", err)
		}
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
