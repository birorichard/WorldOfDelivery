package service

import (
	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
)

type DatabaseQueue struct {
	cache chan model.ShipRouteCache
	run   bool
}

func (d *DatabaseQueue) Setup(capacity int) {
	d.cache = make(chan model.ShipRouteCache, capacity)
}

func (d *DatabaseQueue) Add(routeCache *model.ShipRouteCache) {
	d.cache <- *routeCache
}

func (d *DatabaseQueue) IsFull() bool {
	return len(d.cache) == cap(d.cache)
}

func (d *DatabaseQueue) GetSize() int {
	return len(d.cache)
}

// Inserts routes to the database when the added ShipRouteCache's count reached a given capacity
func (d *DatabaseQueue) Start() {
	d.run = true
	for d.run {
		if d.IsFull() {
			for i := 0; i < cap(d.cache); i++ {
				route := <-d.cache
				d.addRoute(&route)
			}
		}
	}
}

func (d *DatabaseQueue) addRoute(route *model.ShipRouteCache) {
	repository.AddRoute(route)
	lock.Lock()
	cachedRoute := RouteCache[route.ShipId]
	cachedRoute.TableData.Commited = true
	RouteCache[route.ShipId] = cachedRoute
	lock.Unlock()
}

func (d *DatabaseQueue) Close() {
	close(d.cache)
	d.run = false
}

func (d *DatabaseQueue) Reset() {
	for len(d.cache) > 0 {
		<-d.cache
	}
}
