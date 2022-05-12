package main

import (
	"net/http"

	"github.com/birorichard/WorldOfDelivery/handlers"
	"github.com/birorichard/WorldOfDelivery/logging"
	"github.com/birorichard/WorldOfDelivery/middleware"
	"github.com/birorichard/WorldOfDelivery/repositories"
	"github.com/go-chi/chi/v5"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.CreateRequestCounterMiddleware)
	router.Route("/ship", handlers.ShipHandler)
	router.Route("/radio", handlers.RadioHandler)
	router.Route("/missile", handlers.MissileHandler)
	router.Route("/route", handlers.ShipRouteHandler)

	defer repositories.Database.Close()

	repositories.OpenDB()
	repositories.CreateScheme()

	go logging.ConfigureLogging()

	http.ListenAndServe(":8080", router)

}
