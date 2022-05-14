package main

import (
	"net/http"

	"github.com/birorichard/WorldOfDelivery/handlers"
	"github.com/birorichard/WorldOfDelivery/logging"
	"github.com/birorichard/WorldOfDelivery/repository"
	"github.com/birorichard/WorldOfDelivery/service"
	"github.com/go-chi/chi/v5"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	router := chi.NewRouter()

	router.Route("/ship", handlers.ShipHandler)
	router.Route("/radio", handlers.RadioHandler)
	router.Route("/missile", handlers.MissileHandler)

	router.Route("/route", handlers.ShipRouteHandler)

	defer repository.Database.Close()

	repository.OpenDB()
	repository.CreateScheme()

	go logging.ConfigureLogging()

	service.DbQueue.Setup(10)
	go service.DbQueue.Start()

	http.ListenAndServe(":8080", router)

}
