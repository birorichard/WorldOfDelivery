package service

import "github.com/birorichard/WorldOfDelivery/model"

type Position struct {
	X int
	Y int
}

type ShipRoute struct {
	ShipId          int
	PortId          int
	DestinationPort int
	Route           []Position
}

var solvedPorts []int

var routes []ShipRoute

var trackedShips []int

func StartTrackingShip(dto *model.ShipLeavePortDTO) {
	if isPortSolved(&dto.PortId) || isPortSolved(&dto.DestinationPort) {
		return
	}

	// ha meg nincs kovetve a hajo ES nincs kesz az indito kikoto
	// akkor kezdjuk el feljegyezni mind a harom adatot es hozza
	// _, err := solvedPorts[dto.PortId]
}

func isPortSolved(portId *int) bool {
	i := 0
	stop := false
	for i < 5 && !stop {
		if solvedPorts[i] == *portId {
			return true
		}
	}

	return false

}

func isShipTrackingInProgress(shipId *int) bool {
	i := 0
	stop := false
	for i < 5 && !stop {
		if trackedShips[i] == *shipId {
			return true
		}
	}

	return false
}
