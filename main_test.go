package main

import (
	"fmt"
	"testing"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repositories"
	"github.com/birorichard/WorldOfDelivery/service"
)

func TestShipLeavePort(t *testing.T) {

	repositories.OpenDB()
	repositories.CreateScheme()

	var shipIds = []string{
		"01CY1J41CYM2CDEPYKJRNQMGP9",
		"1133C91092E64F9EA9738A2B01C",
		"1444B3A1E0864AFDB2D28B8BFE",
	}

	var shipLeavePorts = []model.ShipLeavePortDTO{
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[0], X: 0, Y: 0,
			},
			PortId:          2,
			DestinationPort: 91},
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[1], X: 256, Y: 34,
			},
			PortId:          82,
			DestinationPort: 97},
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[2], X: 1022, Y: 987,
			},
			PortId:          7,
			DestinationPort: 125},
	}

	var movements = []model.ShipPositionDTO{
		{ShipId: shipIds[2], X: 1023, Y: 987},
		{ShipId: shipIds[0], X: 1, Y: 1},
		{ShipId: shipIds[1], X: 257, Y: 34},
		{ShipId: shipIds[0], X: 2, Y: 2},
		{ShipId: shipIds[2], X: 1022, Y: 987},
		{ShipId: shipIds[1], X: 258, Y: 34},
		{ShipId: shipIds[0], X: 2, Y: 3},
		{ShipId: shipIds[0], X: 3, Y: 3},
		{ShipId: shipIds[2], X: 1022, Y: 988},
		{ShipId: shipIds[0], X: 3, Y: 4},
		{ShipId: shipIds[1], X: 258, Y: 35},
		{ShipId: shipIds[0], X: 4, Y: 4},
	}

	var shipReachedDestinations = []model.ShipReachedDestinationDTO{
		{ShipPositionDTO: model.ShipPositionDTO{
			ShipId: shipIds[0], X: 4, Y: 6,
		},
			PortId: 91, Payment: 420000},
	}

	for _, element := range shipLeavePorts {
		service.StartShipTracking(&element)
	}

	for _, element := range movements {
		service.RegisterShipMovement(&element)
	}

	for _, element := range shipReachedDestinations {
		service.EndShipTracking(&element)
	}

	if len(service.RouteCache) != 3 {
		t.Errorf("Not enough ships")
	}
	fmt.Println(repositories.GetRoute(2, 91))

}

func UntrackedShipsShouldBeIgnored(t *testing.T) {

}

func RoutesAreInProperOrder(t *testing.T) {

}
