package service

import (
	"testing"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/test_data"
)

func TestStartShipTrackingShouldAddNewRouteToTheCache(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)

	if len(RouteCache) == 0 {
		t.Fatal("The ShipLeavePortDTO wasn't added to the cache!")
	}

}

func TestStartShipTrackingShouldAddRouteAsUndiscovered(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	cachedRoute := RouteCache[test_data.TestShipIds[0]]

	if cachedRoute.Discovered {
		t.Fatal("The route has been added as discovered to the cache.")
	}
}

func TestStartShipTrackingShouldAddRouteWithProperShipId(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	cachedRoute := RouteCache[test_data.TestShipIds[0]]

	if cachedRoute.ShipId != test_data.TestShipIds[0] {
		t.Fatal("The route has been added with wron ship id.")
	}
}

func TestStartShipTrackingShouldAddAsUncommited(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	cachedRoute := RouteCache[test_data.TestShipIds[0]]

	if cachedRoute.TableData.Commited {
		t.Fatal("The route has been added as commited to the cache.")
	}
}

func TestStartShipTrackingShouldNotBeAddPosition(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	cachedRoute := RouteCache[test_data.TestShipIds[0]]

	if len(cachedRoute.TableData.Steps) > 0 {
		t.Fatal("The route has been added with position data.")
	}
}

func TestStartShipTrackingShouldIgnoreIfTheRouteisAlreadyTracked(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          22,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	leavePortDto.ShipId = test_data.TestShipIds[1]

	StartShipTracking(&leavePortDto)

	if len(RouteCache) > 1 {
		t.Fatal("An already tracked route has been added to the cache.")
	}
}

func TestStartShipTrackingShouldIgnoreIfTheShipIsAlreadyTracked(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          23,
		DestinationPort: 95,
	}

	StartShipTracking(&leavePortDto)
	leavePortDto.PortId = 26
	StartShipTracking(&leavePortDto)

	if len(RouteCache) > 1 {
		t.Fatal("An already tracked ship has been added to the cache.")
	}
}

// func TestRegisterShipMovementShouldIgnoreDiscoveredRoutes
