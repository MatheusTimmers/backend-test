package handlers

import (
	appErr "github.com/MatheusTimmers/backend-test/internal/errors"
	"github.com/MatheusTimmers/backend-test/pkg/models"

	"github.com/MatheusTimmers/backend-test/internal/db"
	"github.com/gofiber/fiber/v2"
)

func Ranking(c *fiber.Ctx) error {
  users, count, err := db.GetTopUsers()
  if err != nil {
    return appErr.DBError(err)
  }

  var ranking []models.RankingItem
  for i, u := range(users) {
    ranking = append(ranking, models.RankingItem{
			Index:  uint(i + 1),
			Name:   u.Name,
			Email:  u.Email,
			Phone:  u.Phone,
			Points: u.Points,
		})
  }

  return c.JSON(models.RankingListResponse{Total: count, Ranking: ranking})
}
