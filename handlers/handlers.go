package handlers

import (
	"github.com/birorichard/WorldOfDelivery/actions"
	"github.com/birorichard/WorldOfDelivery/middleware"
	"github.com/go-chi/chi/v5"
)

func ShipHandler(r chi.Router) {

	r.Use(middleware.CreateRequestCounterMiddleware)
	r.Post("/DiveComplete", actions.DiveComplete)
	r.Post("/ReachedLand", actions.ReachedLand)

}

func RadioHandler(r chi.Router) {
	r.Use(middleware.CreateRequestCounterMiddleware)
	r.Post("/ShipReachedDestination", actions.ShipReachedDestination)
	r.Post("/ShipMovement", actions.ShipMovement)
	r.Post("/ShipLeavePort", actions.ShipLeavePort)
	r.Post("/ShipUnderAttack", actions.ShipUnderAttack)

}

func MissileHandler(r chi.Router) {
	r.Use(middleware.CreateRequestCounterMiddleware)
	r.Post("/Explosion", actions.Explosion)
}

func ShipRouteHandler(r chi.Router) {
	r.Get("/", actions.GetShipRoutes)
}
