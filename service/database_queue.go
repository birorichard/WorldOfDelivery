package service

import (
	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
)

type DatabaseQueue struct {
	cache chan model.ShipRouteCache
}

func (d *DatabaseQueue) Add(routeCache *model.ShipRouteCache) {
	d.cache <- *routeCache
}

// func (d *DatabaseQueue) IsEmpty() bool {
// 	return len(d.cache) == 0
// }

func (d *DatabaseQueue) IsFull() bool {
	return len(d.cache) == cap(d.cache)
}

// Inserts routes to the database when the added ShipRouteCache's count reached a given capacity
func (d *DatabaseQueue) Start(capacity int) {
	d.cache = make(chan model.ShipRouteCache, capacity)
	for {
		if d.IsFull() {
			for i := 0; i < capacity; i++ {
				route := <-d.cache
				repository.AddRoute(&route)
				lock.Lock()
				cachedRoute := RouteCache[route.ShipId]
				cachedRoute.TableData.Commited = true
				RouteCache[route.ShipId] = cachedRoute

				lock.Unlock()
			}
		}
	}
}
