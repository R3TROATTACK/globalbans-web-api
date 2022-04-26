package gb

import (
	"context"
	"encoding/json"
	"go/importer"
	"os"
	"reflect"

	"github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/services/addons"
	"insanitygaming.net/bans/src/gb/services/auth"
	"insanitygaming.net/bans/src/gb/services/database"
	"insanitygaming.net/bans/src/gb/services/logger"
	"insanitygaming.net/bans/src/gb/services/router"
)

type GB struct {
	eventBus EventBus.Bus
	database *database.Database
	router   *router.Router
	addons   *addons.Registry
}

var app *GB

func New(setup bool) *GB {
	if app == nil {
		app = &GB{}
		app.Init(setup)
	}

	return app
}

func App() *GB {
	return app
}

func (gb *GB) Init(setup bool) {
	gb.eventBus = EventBus.New()
	gb.database = database.New()
	gb.router = router.New()
	gb.addons = addons.New()

	gb.router.AddMiddleware(func(c *gin.Context) {
		c.Set("app", gb)
		c.Next()
	})

	gb.router.AddMiddleware(func(c *gin.Context) {
		c.Set("eventbus", gb.eventBus)
		c.Next()
	})

	type Route struct {
		Path     string              `json:"path"`
		Package  string              `json:"package"`
		Callback string              `json:"callback"`
		Methods  []router.HTTPMethod `json:"methods"`
	}

	type RouteGroup struct {
		Path   string  `json:"path"`
		Routes []Route `json:"routes"`
		Auth   bool    `json:"auth"`
	}

	routes := make([]Route, 0)
	groups := make([]RouteGroup, 0)

	f, err := os.Open("routes.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&routes)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&groups)
	if err != nil {
		panic(err)
	}

	imp := importer.Default()
	for _, route := range routes {
		pkg, err := imp.Import(route.Package)
		if err != nil {
			logger.Logger().Error(err)
			continue
		}
		r := reflect.ValueOf(pkg)
		if r.IsNil() {
			logger.Logger().Error("Package not found")
			continue
		}

		v := r.MethodByName(route.Callback).Interface().(func(*gin.Context))

		gb.router.AddRoute(route.Path, route.Methods, v)
	}

	for _, group := range groups {
		r := gb.router.Group(group.Path, func(c *gin.Context) {

		})
		for _, route := range group.Routes {
			pkg, err := imp.Import(route.Package)
			if err != nil {
				logger.Logger().Error(err)
				continue
			}
			ref := reflect.ValueOf(pkg)
			if ref.IsNil() {
				logger.Logger().Error("Package not found")
				continue
			}
			v := ref.MethodByName(route.Callback).Interface().(func(*gin.Context))
			for _, method := range route.Methods {
				r.Handle(string(method), route.Path, v)
			}
		}
		if group.Auth {
			r.Use(auth.Middleware)
		}
	}

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
