/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/jeffotoni/gmongocrud/conf"
	"github.com/jeffotoni/gmongocrud/handler"
	"github.com/jeffotoni/gmongocrud/lib/cache"
	"github.com/jeffotoni/gmongocrud/lib/context"
	"github.com/jeffotoni/gmongocrud/lib/cors"
	"github.com/jeffotoni/gmongocrud/lib/gjwt"
	"github.com/jeffotoni/gmongocrud/lib/template"
	"gopkg.in/macaron.v1"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	/*
		Redis Cache
		Add this lib to import session: _ "github.com/go-macaron/cache/redis"
		Later replaces the cache in memory instructions for the lines below
		optCache := mcache.Options{
				Adapter:       conf.Cfg.Section("").Key("cache_adapter").Value(),
				AdapterConfig: conf.Cfg.Section("").Key("cache_adapter_config").Value(),
			}
		app.Use(mcache.Cacher(optCache))
	*/
	app.Use(session.Sessioner())
	app.Use(context.Contexter())
	app.Use(cors.Cors())
}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {

	app.Get("", func() string {
		return "Hello, Mercurius Works!"
	})

	// Group de rotas
	app.Group("/v1", func() {

		// app.Group("/oauth2", func() {
		// 	app.Get("/token", handler.GetAccessToken)
		// 	app.Post("/generatecredentials", handler.GetUserCredentials)
		// })

		app.Group("/user", func() {

			app.Post("/login", handler.Login)

			// app.Post("/login", func() string {
			// 	return "login ok."
			// })
		})

		app.Group("/public", func() {

			app.Post("/ping", handler.Ping)
		})

		// // group
		// // mongoDb
		app.Group("/curriculum", func() {

			// add on base
			app.Post("/", handler.HandlerFuncAuth(gjwt.AuthBasicJwt, handler.CurriculumCreate))

			// delete database
			//app.Delete("/:id", handler.CurriculumDelete)

			// update database
			app.Put("/:id", handler.CurriculumUpdate)

			// search database
			app.Get("/:id", handler.CurriculumFind)

			// buscando na base de dados todos registros
			// app.Get("/", handler.CurriculumFindAll)
		})

	})
}
