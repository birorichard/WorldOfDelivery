package services

import (
	"sync"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repositories"
)

var ShipRouteCache = map[string]model.ShipRouteCache{}

var lock = sync.RWMutex{}

func StartShipTracking(dto *model.ShipLeavePortDTO) {
	if isTheRouteAlreadyDiscovered(&dto.PortId, &dto.DestinationPort) || isTheRouteBeingFollowed(&dto.PortId, &dto.DestinationPort) {
		return
	}

	lock.Lock()
	defer lock.Unlock()
	ShipRouteCache[dto.ShipId] = model.ShipRouteCache{
		TableData: model.ShipRoute{
			SourcePortId:      dto.PortId,
			DestinationPortId: dto.DestinationPort,
			Steps:             []model.Position{},
		},
		Discovered: false,
	}
}

func RegisterShipMovement(dto *model.ShipPositionDTO) {
	lock.Lock()
	if route, ok := ShipRouteCache[dto.ShipId]; ok {
		if route.Discovered {
			lock.Unlock()
			return
		}
		route.TableData.Steps = append(route.TableData.Steps, model.Position{X: dto.X, Y: dto.Y, StepOrder: len(route.TableData.Steps)})
		ShipRouteCache[dto.ShipId] = route

	}
	lock.Unlock()
}

func EndShipTracking(dto *model.ShipReachedDestinationDTO) {
	lock.Lock()

	if route, ok := ShipRouteCache[dto.ShipId]; ok {
		route.Discovered = true
		ShipRouteCache[dto.ShipId] = route
		lock.Unlock()
		repositories.AddRoute(&route)
	} else {
		lock.Unlock()
	}
}

func GetFoundRoutesCount() int {
	var count int
	lock.RLock()
	defer lock.RUnlock()
	for _, v := range ShipRouteCache {
		if v.Discovered {
			count++
		}
	}

	return count
}

func GetShipRoutes(fromCache bool) []model.ShipRoute {
	if fromCache {
		return getRouteDtosFromCache()
	} else {
		return repositories.GetAllRoutesFrom()
	}

}

func getRouteDtosFromCache() []model.ShipRoute {
	var routeDtos []model.ShipRoute = make([]model.ShipRoute, 0)

	for _, route := range ShipRouteCache {
		if route.Discovered {
			routeDtos = append(routeDtos, model.ShipRoute{SourcePortId: route.TableData.SourcePortId, DestinationPortId: route.TableData.DestinationPortId, Steps: route.TableData.Steps})

			// for stepIndex, step := range route.TableData.Steps {
			// 	routeDtos = append(routeDtos, model.ShipRoute{SourcePortId: route.TableData.SourcePortId, DestinationPortId: route.TableData.DestinationPortId, PositionX: step.X, PositionY: step.Y, StepOrder: stepIndex})

			// }
		}
	}

	return routeDtos
}

func isTheRouteAlreadyDiscovered(sourcePortId *int, destinationPortId *int) bool {

	if len(ShipRouteCache) == 0 {
		return false
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, element := range ShipRouteCache {
		if element.TableData.SourcePortId == *sourcePortId && element.TableData.DestinationPortId == *destinationPortId && element.Discovered {
			return true
		}
	}

	return false
}

func isTheRouteBeingFollowed(sourcePortId *int, destinationPortId *int) bool {
	lock.RLock()
	defer lock.RUnlock()

	for _, element := range ShipRouteCache {
		if element.TableData.SourcePortId == *sourcePortId && element.TableData.DestinationPortId == *destinationPortId {
			return true
		}
	}

	return false
}
