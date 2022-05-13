package service

import (
	"math"
	"sync"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
)

var RouteCache = map[string]model.ShipRouteCache{}

var DbQueue = DatabaseQueue{}

var lock = sync.RWMutex{}

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
		Discovered: false,
		ShipId:     dto.ShipId,
	}

	RouteCache[dto.ShipId] = route
}

func RegisterShipMovement(dto *model.ShipPositionDTO) {
	lock.Lock()
	if route, ok := RouteCache[dto.ShipId]; ok {
		if route.Discovered || route.Canceled {
			lock.Unlock()
			return
		}
		if len(route.TableData.Steps) > 0 {
			lastStep := route.TableData.Steps[len(route.TableData.Steps)-1]
			if !isShipMovementValid(dto.ShipId, dto.X, dto.Y, lastStep.X, lastStep.Y) {
				// CancelShipRoute(dto.ShipId)
				// route.Canceled = true

				// RouteCache[dto.ShipId] = route
				lock.Unlock()
				return
			}
		}
		route.TableData.Steps = append(route.TableData.Steps, model.Position{X: dto.X, Y: dto.Y, StepOrder: len(route.TableData.Steps)})
		RouteCache[dto.ShipId] = route

	}
	lock.Unlock()
}

func CancelShipRoute(shipId string) {
	lock.Lock()
	if route, ok := RouteCache[shipId]; ok {
		route.Canceled = true
		RouteCache[shipId] = route

	}
	lock.Unlock()
}

func isShipMovementValid(shipId string, PosX int, PosY int, PrevPosX int, PrevPosY int) bool {
	if PrevPosX == PosX && math.Abs(float64(PrevPosY-PosY)) == 1 {
		return true

	}
	if PrevPosY == PosY && math.Abs(float64(PrevPosX-PosX)) == 1 {
		return true
	}

	if math.Abs(float64(PrevPosX-PosX)) == 1 && math.Abs(float64(PrevPosY-PosY)) == 1 {
		return true
	}

	// route := RouteCache[shipId]
	// route.Canceled = true

	// RouteCache[shipId] = route
	return false
}

func EndShipTracking(dto *model.ShipReachedDestinationDTO) {
	lock.Lock()

	if route, ok := RouteCache[dto.ShipId]; ok {
		// if len(route.TableData.Steps) > 0 {
		// 	lastStep := route.TableData.Steps[len(route.TableData.Steps)-1]
		// 	fmt.Println(dto.ShipPositionDTO)
		// 	fmt.Println(lastStep.X, ", ", lastStep.Y)
		// 	fmt.Println("\n\n")
		// 	if !isShipMovementValid(dto.ShipId, dto.ShipPositionDTO.X, dto.ShipPositionDTO.Y, lastStep.X, lastStep.Y) {
		// 		// CancelShipRoute(dto.ShipId)
		// 		// route.Canceled = true

		// 		// RouteCache[dto.ShipId] = route
		// 		lock.Unlock()
		// 		return
		// 	}
		// }
		route.Discovered = true
		RouteCache[dto.ShipId] = route
		lock.Unlock()
		DbQueue.Add(&route)
	} else {
		lock.Unlock()
	}
}

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

func GetShipRoutes(fromCache bool) []model.ShipRouteDTO {
	if fromCache {
		return getRouteDtosFromCache()
	} else {
		return repository.GetAllRoutes()
	}

}

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
