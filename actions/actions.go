package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/services"
)

func ShipMovement(w http.ResponseWriter, r *http.Request) {
	var dto model.ShipPositionDTO
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleBadRequest(w)
	}
	err = json.Unmarshal(bodyBytes, &dto)
	if err != nil {
		handleBadRequest(w)
	}
	go services.RegisterShipMovement(&dto)

	handleOkRequest(w)
}

func ShipReachedDestination(w http.ResponseWriter, r *http.Request) {
	var dto model.ShipReachedDestinationDTO
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleBadRequest(w)
	}
	err = json.Unmarshal(bodyBytes, &dto)
	if err != nil {
		handleBadRequest(w)
	}
	go services.EndShipTracking(&dto)

	handleOkRequest(w)

}

func ShipLeavePort(w http.ResponseWriter, r *http.Request) {

	var dto model.ShipLeavePortDTO
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleBadRequest(w)
	}
	err = json.Unmarshal(bodyBytes, &dto)
	if err != nil {
		handleBadRequest(w)
	}
	go services.StartShipTracking(&dto)

	handleOkRequest(w)

}

func ShipUnderAttack(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func DiveComplete(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func ReachedLand(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func Explosion(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func GetShipRoutes(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)

	fromCache := false

	if r.URL.Query().Get("fromCache") == "true" {
		fromCache = true
	}
	shipRoutes := services.GetShipRoutes(fromCache)
	json.NewEncoder(w).Encode(shipRoutes)

}

func handleOkRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func handleBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
