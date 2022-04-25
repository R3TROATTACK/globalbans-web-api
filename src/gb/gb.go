package gb

import (
	"context"
	"os"

	"github.com/asaskevich/EventBus"
	"insanitygaming.net/bans/src/gb/services/addons"
	"insanitygaming.net/bans/src/gb/services/database"
	"insanitygaming.net/bans/src/gb/services/router"
)

type GB struct {
	eventBus EventBus.Bus
	database *database.Database
	router   *router.Router
	addons   *addons.Registry
}

var app *GB

func New() *GB {
	if app == nil {
		app = &GB{}
		app.Init()
	}

	return app
}

func App() *GB {
	return app
}

func (gb *GB) Init() {
	gb.eventBus = EventBus.New()
	gb.database = database.New()
	gb.router = router.New()
	gb.addons = addons.New()
}

func (gb *GB) Setup() {
	gb.database.RunSetup(gb.Context())
}

func (gb *GB) Run() {
	gb.router.Run(os.Getenv("APP_ADDR") + ":" + os.Getenv("APP_PORT"))
}

func (gb *GB) Router() *router.Router {
	return gb.router
}

func (gb *GB) Database() *database.Database {
	return gb.database
}

func (gb *GB) Addons() *addons.Registry {
	return gb.addons
}

func (gb *GB) EventBus() EventBus.Bus {
	return gb.eventBus
}

func (gb *GB) Context() context.Context {
	return context.Background()
}
