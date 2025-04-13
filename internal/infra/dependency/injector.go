package dependency

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MatheusTimmers/backend-test/internal/interface/http/routes"
	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	"github.com/MatheusTimmers/backend-test/internal/application/usecase/user"
	"github.com/MatheusTimmers/backend-test/internal/infra/db"
	"github.com/MatheusTimmers/backend-test/internal/infra/email"
)

type injector struct {
	Db *gorm.DB
	App *fiber.App
	Mailer usecase.Mailer
	Users  user.UserService
}

var injectorInit sync.Once
var instance *injector

func Injector() *injector {
	if instance == nil {
		injectorInit.Do(
			func() {
				instance = &injector{}
			},
		)
	}
	return instance
}

func (i *injector) Inject() {
	if i.Db == nil {
		i.Db = db.Db()
	}

	if i.Mailer == nil {
		i.Mailer = email.NewMailer()
	}

	if i.Users == nil {
		i.Users = user.NewUserService(db.NewUserRepository(i.Db), i.Mailer)
	}

	if i.App == nil {
		i.App = fiber.New()
		routes.RegisterRoutes(i.App, i.Users)
	}
}
