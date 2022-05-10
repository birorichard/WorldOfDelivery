package service

import (
	"sync"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repositories"
)

var StoredShipRoutes = map[string]model.ShipRoute{}

var lock = sync.RWMutex{}

// var solvedPortIds []int

// var routes []ShipRoute

// var trackedShips []int

func StartTrackingShip(dto *model.ShipLeavePortDTO) {

	// megvan mar oldva az ut?
	if isPortSolved(&dto.PortId, &dto.DestinationPort) || isPortSolved(&dto.DestinationPort, &dto.PortId) {
		return
	}

	// el van mar kezdve az ut tarolasa?
	// itt meg nem nezzuk az uj hajokat, hatha azok hamarabb vegeznek, mint amit mar kovetunk
	// az adott uton
	if isPortRouteAlreadyTracking(&dto.PortId, &dto.DestinationPort) {
		return
	}

	lock.Lock()
	defer lock.Unlock()
	StoredShipRoutes[dto.ShipId] = model.ShipRoute{
		SourcePortId:      dto.PortId,
		DestinationPortId: dto.DestinationPort,
		Steps: []model.Position{
			{X: dto.X, Y: dto.Y},
		},
		Solved: false,
	}

	// fmt.Println(routes[dto.ShipId])

	// routes[dto.ShipId] = ShipRoute

	// var currentRoute ShipRoute
	// var currentShipId Ship
	// for key, element := range routes {
	// 	if key == dto.ShipId
	// 	{
	// 		currentShipId = dto.ShipId
	// 	}

	// 	if element[]{

	// 	}
	// }

	// routes = append(routes,
	// 	ShipRoute{dto.ShipId, dto.PortId, dto.DestinationPort, []Position{{dto.X, dto.Y}}})

	// ha meg nincs kovetve a hajo ES nincs kesz az indito kikoto
	// akkor kezdjuk el feljegyezni mind a harom adatot es hozza

	// ha mar kovetjuk a hajot es nincs befejezve, akkor adjuk hozza a mapban a poziciokhoz a mostanit
	// _, err := solvedPorts[dto.PortId]
}

func RegisterMove(dto *model.ShipPositionDTO) {
	// lock.RLock()
	lock.Lock()
	if route, ok := StoredShipRoutes[dto.ShipId]; ok {
		route.Steps = append(route.Steps, model.Position{X: dto.X, Y: dto.Y})
		StoredShipRoutes[dto.ShipId] = route

	}
	lock.Unlock()

	// route := StoredShipRoutes[dto.ShipId]
	// lock.RUnlock()

}

func CloseRoute(dto *model.ShipReachedDestinationDTO) {
	lock.Lock()

	// route := StoredShipRoutes[dto.ShipId]
	if route, ok := StoredShipRoutes[dto.ShipId]; ok {
		route.Steps = append(route.Steps, model.Position{X: dto.X, Y: dto.Y})
		route.Solved = true
		StoredShipRoutes[dto.ShipId] = route
		lock.Unlock()
		repositories.AddRoute(&route)

		repositories.GetRoutes(route.SourcePortId)

	} else {
		lock.Unlock()
	}

	// else {
	// 	lock.Unlock()
	// }

	// route.Steps = append(route.Steps, model.Position{X: dto.X, Y: dto.Y})

	// StoredShipRoutes[dto.ShipId] = route
	// fmt.Println(StoredShipRoutes[dto.ShipId])

}

func isPortSolved(sourcePortId *int, destinationPortId *int) bool {

	if len(StoredShipRoutes) == 0 {
		return false
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, element := range StoredShipRoutes {
		if element.SourcePortId == *sourcePortId && element.DestinationPortId == *destinationPortId && element.Solved {
			return true
		}
	}

	return false
}

func isShipTrackingInProgress(shipId *string) bool {
	lock.RLock()
	defer lock.RUnlock()
	for key := range StoredShipRoutes {
		if key == *shipId {
			return true
		}
	}

	return false
}

func isPortRouteAlreadyTracking(sourcePortId *int, destinationPortId *int) bool {
	lock.RLock()
	defer lock.RUnlock()

	for _, element := range StoredShipRoutes {
		if element.SourcePortId == *sourcePortId && element.DestinationPortId == *destinationPortId {
			return true
		}
	}

	return false
}
