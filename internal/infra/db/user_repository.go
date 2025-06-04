package db

import (
	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	appErr "github.com/MatheusTimmers/backend-test/pkg/errors"

	"github.com/MatheusTimmers/backend-test/pkg/models"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUser(user *models.User) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return appErr.DBError(err, "failed to save user to DB")
	}

	if err := tx.Commit().Error; err != nil {
		return appErr.DBError(err, "failed to commit to DB ")
	}

	return nil
}

func (r *gormUserRepository) SaveInviter(inviter *models.User) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(&inviter).Error; err != nil {
		tx.Rollback()
		return appErr.DBError(err, "failed to save user to DB")
	}

	if err := tx.Commit().Error; err != nil {
		return appErr.Internal("failed to commit to DB ")
	}

	return nil
}

func (r *gormUserRepository) FindInviterByInviteCode(inviteCode string) (*models.User, error) {
	if inviteCode == "" {
		return nil, nil
	}

	var inviter models.User
	err := r.db.Where("invite_code = ?", inviteCode).First(&inviter).Error
	if err != nil {
		return nil, appErr.BadRequest("invalid invite code")
	}

	inviter.Points += 1

	return &inviter, nil
}

func (r *gormUserRepository) GetTopUsers() ([]models.User, int, error) {
	var users []models.User
	err := r.db.Select("name", "email", "phone", "points").Order("points DESC").Limit(10).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	var count int64
	err = r.db.Table("users").Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
}
