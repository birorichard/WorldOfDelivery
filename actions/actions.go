package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/service"
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
	go service.RegisterShipMovement(&dto)

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
	go service.EndShipTracking(&dto)

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
	go service.StartShipTracking(&dto)

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

	shipRoutes := service.GetShipRoutes(r.URL.Query().Get("fromCache") == "true")
	json.NewEncoder(w).Encode(shipRoutes)

}

func handleOkRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func handleBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
