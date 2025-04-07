package db

import (
	appErr "github.com/MatheusTimmers/backend-test/internal/errors"

	"github.com/MatheusTimmers/backend-test/pkg/models"
	"github.com/google/uuid"
)

func CreateUser(req models.RegisterRequest) (*models.User, error) {
  tx := DB.Begin()
  defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
  }()

  inviter, err := getInviter(req.InviteCode)
	if err != nil {
		tx.Rollback()
	  return nil, err
	}

	var invitedBy *uint
	if inviter != nil {
		if err := tx.Save(inviter).Error; err != nil {
			tx.Rollback()
			return nil, appErr.Internal("failed to update inviter points")
		}
		invitedBy = &inviter.ID
	}

  newCode := uuid.New().String()[:8]

	user := models.User{
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		InviteCode: newCode,
		InvitedBy:  invitedBy,
		Points:     1,
	}
  
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, appErr.DBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, appErr.Internal("failed to commit to DB ")
	}

	return &user, nil
}

func getInviter(inviteCode string) (*models.User, error) {
  if inviteCode == "" {
    return nil, nil 
  }

  var inviter models.User
	err := DB.Where("invite_code = ?", inviteCode).First(&inviter).Error
	if err != nil {
		return nil, appErr.BadRequest("invalid invite code")
	}

	inviter.Points += 1

	return &inviter, nil
}
