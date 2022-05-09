package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/birorichard/WorldOfDelivery/model"
)

func ShipMovement(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func ShipReachedDestination(w http.ResponseWriter, r *http.Request) {
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
	// fmt.Println(dto.ShipPositionDTO, dto.DestinationPort)
	handleOkRequest(w)
	w.Write(bodyBytes)
}

func ShipUnderAttack(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func DiveComplete(w http.ResponseWriter, r *http.Request) {
	var dto model.DiveCompleteDTO
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleBadRequest(w)
	}
	err = json.Unmarshal(bodyBytes, &dto)
	if err != nil {
		handleBadRequest(w)
	}
	handleOkRequest(w)

	w.Write(bodyBytes)

}

func ReachedLand(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func Explosion(w http.ResponseWriter, r *http.Request) {
	handleOkRequest(w)
}

func handleOkRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func handleBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
