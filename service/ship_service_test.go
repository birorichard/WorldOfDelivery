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
		t.Fatal("The route has been added with wrong ship id.")
	}
}

func TestStartShipTrackingShouldAddAsUncommitted(t *testing.T) {
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

	if cachedRoute.TableData.Committed {
		t.Fatal("The route has been added as committed to the cache.")
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

func TestRegisterShipMovementShouldIgnoreFromShipWhoseRouteIsAlreadyDiscovered(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          23,
		DestinationPort: 95,
	}

	shipPositionDtos := []model.ShipPositionDTO{
		{
			ShipId: test_data.TestShipIds[0],
			X:      5,
			Y:      9,
		},
		{
			ShipId: test_data.TestShipIds[0],
			X:      6,
			Y:      9,
		},
	}

	StartShipTracking(&leavePortDto)
	for _, position := range shipPositionDtos {
		RegisterShipMovement(&position)
	}

	route := RouteCache[test_data.TestShipIds[0]]
	route.Discovered = true
	RouteCache[test_data.TestShipIds[0]] = route

	RegisterShipMovement(&model.ShipPositionDTO{
		ShipId: test_data.TestShipIds[0],
		X:      6,
		Y:      10,
	})

	route = RouteCache[test_data.TestShipIds[0]]
	if len(route.TableData.Steps) > 2 {
		t.Fatal("Discovered route's step wasn't ignored.")
	}
}

func TestRegisterShipMovementShouldUpdateTheTableSize(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	tableSize.X = 300
	tableSize.Y = 500

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          23,
		DestinationPort: 95,
	}

	shipPositionDto := model.ShipPositionDTO{
		ShipId: test_data.TestShipIds[0],
		X:      670,
		Y:      512,
	}

	StartShipTracking(&leavePortDto)
	RegisterShipMovement(&shipPositionDto)

	if tableSize.X != 670 || tableSize.Y != 512 {
		t.Fatal("Table size wasn't updated.")
	}
}

func TestRegisterEndShipTrackingShouldSetTheRouteAsDiscovered(t *testing.T) {
	DbQueue = DatabaseQueue{}
	DbQueue.Setup(50)

	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          23,
		DestinationPort: 95,
	}

	shipPositionDto := model.ShipPositionDTO{
		ShipId: test_data.TestShipIds[0],
		X:      670,
		Y:      512,
	}

	shipDestinationReachedDto := model.ShipReachedDestinationDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
			X:      670,
			Y:      512,
		},
		PortId:  95,
		Payment: 10,
	}

	StartShipTracking(&leavePortDto)
	RegisterShipMovement(&shipPositionDto)
	EndShipTracking(&shipDestinationReachedDto)

	route := RouteCache[test_data.TestShipIds[0]]

	if !route.Discovered {
		t.Fatal("Finished route wasn't marked as discovered.")
	}
}

func TestRegisterEndShipTrackingShouldAddTheRouteToTheQueue(t *testing.T) {
	DbQueue = DatabaseQueue{}
	DbQueue.Setup(50)

	RouteCache = map[string]model.ShipRouteCache{}

	leavePortDto := model.ShipLeavePortDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
		},
		PortId:          23,
		DestinationPort: 95,
	}

	shipPositionDto := model.ShipPositionDTO{
		ShipId: test_data.TestShipIds[0],
		X:      670,
		Y:      512,
	}

	shipDestinationReachedDto := model.ShipReachedDestinationDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
			X:      670,
			Y:      512,
		},
		PortId:  95,
		Payment: 10,
	}

	StartShipTracking(&leavePortDto)
	RegisterShipMovement(&shipPositionDto)
	EndShipTracking(&shipDestinationReachedDto)

	if DbQueue.GetSize() != 1 {
		t.Fatal("Finished route wasn't added to the Db queue.")
	}
}

func TestGetFoundRoutesCountShouldCountOnlyDiscoveredRoutes(t *testing.T) {
	RouteCache = map[string]model.ShipRouteCache{}

	RouteCache[test_data.TestShipIds[0]] = model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      6,
			DestinationPortId: 90,
			Steps: []model.Position{
				{X: 3, Y: 3, StepOrder: 0},
				{X: 3, Y: 4, StepOrder: 1},
				{X: 3, Y: 5, StepOrder: 2},
				{X: 3, Y: 6, StepOrder: 3},
			},
			Committed: false,
		},
		PlannedDestinationPortId: 90,
		Discovered:               true,
		ShipId:                   test_data.TestShipIds[0],
	}

	RouteCache[test_data.TestShipIds[1]] = model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      3,
			DestinationPortId: 11,
			Steps: []model.Position{
				{X: 23, Y: 3, StepOrder: 0},
				{X: 23, Y: 4, StepOrder: 1},
				{X: 23, Y: 5, StepOrder: 2},
				{X: 23, Y: 6, StepOrder: 3},
			},
			Committed: false,
		},
		PlannedDestinationPortId: 11,
		Discovered:               true,
		ShipId:                   test_data.TestShipIds[1],
	}

	RouteCache[test_data.TestShipIds[2]] = model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      4,
			DestinationPortId: 23,
			Steps: []model.Position{
				{X: 13, Y: 3, StepOrder: 0},
				{X: 13, Y: 4, StepOrder: 1},
				{X: 13, Y: 5, StepOrder: 2},
				{X: 13, Y: 6, StepOrder: 3},
			},
			Committed: false,
		},
		PlannedDestinationPortId: 23,
		Discovered:               false,
		ShipId:                   test_data.TestShipIds[2],
	}

	discoveredRouteCount := GetFoundRoutesCount()

	if discoveredRouteCount != 2 {
		t.Fatal("Wrong count of discovered routes.")
	}
}
