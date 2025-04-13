package user

import "github.com/MatheusTimmers/backend-test/internal/application/usecase"

type userService struct {
	repo   usecase.UserRepository
	mailer usecase.Mailer
}

func NewUserService(repo usecase.UserRepository, mailer usecase.Mailer) UserService {
	return &userService{repo: repo, mailer: mailer}
}
