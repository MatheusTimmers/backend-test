package user

import "github.com/MatheusTimmers/backend-test/pkg/models"

type UserService interface {
	Register(input models.RegisterRequest) (*models.User, error)
	Ranking() ([]models.RankingItem, *int, error)
	NotifyWinners() error
}
