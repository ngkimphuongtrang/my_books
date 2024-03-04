package handler

import (
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/store"
	"net/http"
)

type App struct {
	mux         *http.ServeMux
	httpHandler http.Handler
	middlewares []httpkit.Middleware

	config *config.AppConfig

	stores *store.DBStores
}

func NewApp(cfg *config.AppConfig, stores *store.DBStores) *App {
	app := &App{
		config: cfg,
		stores: stores,
	}
	app.setup()
	return app
}

func (app *App) setup() {
	app.mux = http.NewServeMux()
	app.httpHandler = app.mux
	app.middlewares = []httpkit.Middleware{&httpkit.RequestTimeMiddleware{}}

	app.setupRoutes()
}

func (app *App) setupRoutes() {
	routeHandlers := app.initRouteHandlers()
	app.addPublicRouteHandlers(routeHandlers...)
}

func (app *App) initRouteHandlers() []*httpkit.RouteHandler {
	return []*httpkit.RouteHandler{
		{
			Route:  RoutePing,
			Handle: app.handlePing,
		},
		{
			Route:  RouteCreateBook,
			Handle: app.handleCreateBook,
		},
	}
}

func (app *App) addPublicRouteHandlers(routeHandlers ...*httpkit.RouteHandler) {
	for _, routeHandler := range routeHandlers {
		var middlewareFn []func(ctx *httpkit.RequestContext)
		for _, r := range app.middlewares {
			middlewareFn = append(middlewareFn, r.Handle)
		}
		middlewareFn = append(middlewareFn, routeHandler.Handle)

		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			ctx := &httpkit.RequestContext{
				RouteHandler: routeHandler,
				Writer:       w,
				Response:     &httpkit.Response{},
				Request:      r,
				Params:       httpkit.Map{},
				Idx:          -1,
				Middlewares:  middlewareFn,
			}
			ctx.Next()
		}
		var handler http.Handler = http.HandlerFunc(handlerFunc)
		app.mux.Handle(routeHandler.Route.Path, handler)
	}
}

func (app *App) Start() {
	server := http.Server{
		Addr:              app.config.Server.Port,
		Handler:           app.httpHandler,
		ReadHeaderTimeout: app.config.Server.ReadHeaderTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
