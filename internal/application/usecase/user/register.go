package user

import (
	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	"github.com/MatheusTimmers/backend-test/pkg/models"

	appErr "github.com/MatheusTimmers/backend-test/pkg/errors"
	"github.com/google/uuid"
)

func (s *userService) Register(input models.RegisterRequest) (*models.User, error) {
	invitedBy, err := s.getInviter(input.InviteCode)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:       input.Name,
		Email:      input.Email,
		Phone:      input.Phone,
		InvitedBy:  invitedBy,
		InviteCode: uuid.New().String()[:8],
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	err = s.mailer.SendEmail(user.Name, user.Email, user.InviteCode, usecase.NewUser)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) getInviter(inviteCode string) (*uint, error) {
	if inviteCode != "" {
		inviter, err := s.repo.FindInviterByInviteCode(inviteCode)
		if err != nil {
			return nil, appErr.DBError(err, "failed to find inviter")
		}

		inviter.Points++
		err = s.repo.SaveInviter(inviter)
		if err != nil {
			return nil, appErr.DBError(err, "failed to update inviter points")
		}

		err = s.mailer.SendEmail(inviter.Name, inviter.Email, inviter.InviteCode, usecase.NewPoint)
		if err != nil {
			return nil, err
		}

		return &inviter.ID, nil
	}

	return nil, nil
}
