package service

import (
	"math"
	"sync"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
	"github.com/birorichard/WorldOfDelivery/validation"
)

var RouteCache = map[string]model.ShipRouteCache{}

var DbQueue = DatabaseQueue{}

var lock = sync.RWMutex{}

var tableSize = model.TableSize{X: 0, Y: 0}

func StartShipTracking(dto *model.ShipLeavePortDTO) {
	if isTheRouteAlreadyDiscovered(&dto.PortId, &dto.DestinationPort) || isTheRouteBeingFollowed(&dto.PortId, &dto.DestinationPort) {
		return
	}

	lock.Lock()
	defer lock.Unlock()
	route := model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      dto.PortId,
			DestinationPortId: dto.DestinationPort,
			Steps:             []model.Position{},
			Commited:          false,
		},
		Discovered:               false,
		ShipId:                   dto.ShipId,
		PlannedDestinationPortId: dto.DestinationPort,
	}

	RouteCache[dto.ShipId] = route
}

func RegisterShipMovement(dto *model.ShipPositionDTO) {
	lock.Lock()
	if route, ok := RouteCache[dto.ShipId]; ok {
		if route.Discovered {
			lock.Unlock()
			return
		}
		if len(route.TableData.Steps) > 0 {
			lastStep := route.TableData.Steps[len(route.TableData.Steps)-1]
			if !validation.IsShipMovementValid(dto.X, dto.Y, lastStep.X, lastStep.Y) {
				lock.Unlock()
				return
			}
		}
		route.TableData.Steps = append(route.TableData.Steps, model.Position{X: dto.X, Y: dto.Y, StepOrder: len(route.TableData.Steps)})
		RouteCache[dto.ShipId] = route
		setTableSize(dto.X, dto.Y)
	}
	lock.Unlock()
}

func EndShipTracking(dto *model.ShipReachedDestinationDTO) {
	lock.Lock()

	if route, ok := RouteCache[dto.ShipId]; ok {
		if !validation.IsReachingDestinationValid(&route, dto) {
			lock.Unlock()
			return
		}

		route.Discovered = true
		RouteCache[dto.ShipId] = route
		lock.Unlock()
		DbQueue.Add(&route)
	} else {
		lock.Unlock()
	}
}

// Returns the count of the routes that already found
func GetFoundRoutesCount() int {
	var count int
	lock.RLock()
	defer lock.RUnlock()
	for _, v := range RouteCache {
		if v.Discovered {
			count++
		}
	}

	return count
}

// Returns all the found routes
func GetShipRoutes(fromCache bool) model.GetAllRoutesResponseDTO {
	var routes []model.ShipRouteDTO

	if fromCache {
		routes = getRouteDtosFromCache()
	} else {
		routes = repository.GetAllRoutes()
	}
	return model.GetAllRoutesResponseDTO{TableSize: model.TableSize{X: int(math.Ceil(float64(tableSize.X) * 1.02)), Y: int(math.Ceil(float64(tableSize.Y) * 1.02))}, Routes: routes}
}

// Maps the routes from type that used for cache to API response
func getRouteDtosFromCache() []model.ShipRouteDTO {
	var routeDtos []model.ShipRouteDTO = make([]model.ShipRouteDTO, 0)

	for _, route := range RouteCache {
		if route.Discovered {
			routeDtos = append(routeDtos, model.ShipRouteDTO{SourcePortId: route.TableData.SourcePortId, DestinationPortId: route.TableData.DestinationPortId, Steps: route.TableData.Steps, Commited: route.TableData.Commited})
		}
	}

	return routeDtos
}

func isTheRouteAlreadyDiscovered(sourcePortId *int, destinationPortId *int) bool {

	if len(RouteCache) == 0 {
		return false
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, element := range RouteCache {
		if element.TableData.SourcePortId == *sourcePortId && element.TableData.DestinationPortId == *destinationPortId && element.Discovered {
			return true
		}
	}

	return false
}

func isTheRouteBeingFollowed(sourcePortId *int, destinationPortId *int) bool {
	lock.RLock()
	defer lock.RUnlock()

	for _, element := range RouteCache {
		if element.TableData.SourcePortId == *sourcePortId && element.TableData.DestinationPortId == *destinationPortId {
			return true
		}
	}

	return false
}

func setTableSize(posX int, posY int) {
	if posX > tableSize.X {
		tableSize.X = posX
	}

	if posY > tableSize.Y {
		tableSize.Y = posY
	}
}
