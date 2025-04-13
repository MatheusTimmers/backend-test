package usecase

import (
	"github.com/MatheusTimmers/backend-test/pkg/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	SaveInviter(inviter *models.User) error
	FindInviterByInviteCode(inviteCode string) (*models.User, error)
	GetTopUsers() ([]models.User, int, error)
}
