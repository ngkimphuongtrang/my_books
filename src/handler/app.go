package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/store"
)

type App struct {
	mux         *http.ServeMux
	httpHandler http.Handler
	middlewares []httpkit.Middleware

	config *config.AppConfig

	stores   *store.DBStores
	handlers *Handler
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

	validation := NewValidation()
	app.handlers = NewHandler(app.stores, validation)
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
			Handle: app.handlers.BookHandler.handleCreateBook,
		},
		{
			Route:  RouteCreateRead,
			Handle: app.handlers.ReadHandler.handleCreateRead,
		},
		{
			Route:  RouteListBooks,
			Handle: app.handlers.BookHandler.handleListBooks,
		},
		{
			Route:  RouteListReads,
			Handle: app.handlers.ReadHandler.handleListReads,
		},
	}
}

func (app *App) addPublicRouteHandlers(routeHandlers ...*httpkit.RouteHandler) {
	router := mux.NewRouter()
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
				Params:       container.Map{},
				Idx:          -1,
				Middlewares:  middlewareFn,
			}
			ctx.Next()
		}

		router.HandleFunc(routeHandler.Route.Path, handlerFunc).Methods(routeHandler.Route.Method)
	}
	app.mux.Handle("/", router) // Or you can directly assign app.mux = router
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
