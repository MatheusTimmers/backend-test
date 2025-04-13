package application

import (
	"log"

	"github.com/MatheusTimmers/backend-test/internal/infra/config"
	"github.com/MatheusTimmers/backend-test/internal/infra/dependency"
)

func Start() {
	injector := dependency.Injector()
	injector.Inject()

	log.Fatal(injector.App.Listen(":" + config.Config().Application.ServerPort))
}
