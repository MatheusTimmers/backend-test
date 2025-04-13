package user

import (
	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	"github.com/MatheusTimmers/backend-test/pkg/models"

	appErr "github.com/MatheusTimmers/backend-test/pkg/errors"
)


func (s *userService) Ranking() ([]models.RankingItem, *int, error) {
	users, count, err := s.repo.GetTopUsers()
	if err != nil {
		return nil, nil, appErr.DBError(err, "failed to get top users")
	}

	var ranking []models.RankingItem
	for i, u := range users {
		ranking = append(ranking, models.RankingItem{
			Index:  uint(i + 1),
			Name:   u.Name,
			Email:  u.Email,
			Phone:  u.Phone,
			Points: u.Points,
		})
	}

	return ranking, &count, err
}

func (s *userService) NotifyWinners() error {
	users, _, err := s.repo.GetTopUsers()
	if err != nil {
		return appErr.DBError(err, "failed to get top users")
	}

	for _, u := range users {
		if err := s.mailer.SendEmail(u.Name, u.Email, u.InviteCode, usecase.Winner); err != nil {
			return appErr.Internal("failed to notify winner: " + u.Email)
		}
	}

	return nil
}
