package application

import (
	"github.com/MatheusTimmers/backend-test/internal/infra/config"
	"github.com/MatheusTimmers/backend-test/internal/infra/dependency"
	"github.com/MatheusTimmers/backend-test/pkg/logger"
)

func Start() {
	injector := dependency.Injector()
	injector.Inject()

	logger.Log().Panic(injector.App.Listen(":" + config.Config().Application.ServerPort))
}
