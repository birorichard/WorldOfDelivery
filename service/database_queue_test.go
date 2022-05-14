package service

import (
	"testing"

	"github.com/birorichard/WorldOfDelivery/common"
	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
	"github.com/birorichard/WorldOfDelivery/test_data"
)

func wipeDb() {
	queryString := []string{
		`DROP TABLE Route;`,
	}
	var err error
	for _, query := range queryString {
		_, err = repository.Database.Exec(query)
		common.HandleError(err, "Drop scheme failed")
	}
}

func TestSetupMakesTheProperSizeBuffer(t *testing.T) {
	capacity := 5
	queue := DatabaseQueue{}
	queue.Setup(capacity)
	if cap(queue.cache) != capacity {
		t.Fatalf("The DatabaseQueue wasn't setup with the proper capacity!")
	}

}

func TestAddShouldAddTheRouteToTheQueue(t *testing.T) {
	queue := DatabaseQueue{}
	queue.Setup(5)

	routeCache := model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      4,
			DestinationPortId: 23,
			Steps: []model.Position{
				{X: 3, Y: 3, StepOrder: 0},
				{X: 3, Y: 4, StepOrder: 1},
				{X: 3, Y: 5, StepOrder: 2},
				{X: 3, Y: 6, StepOrder: 3},
			},
			Committed: false,
		},
		PlannedDestinationPortId: 23,
		Discovered:               false,
		ShipId:                   test_data.TestShipIds[0],
	}

	queue.Add(&routeCache)

	if queue.GetSize() != 1 {
		t.Fatal("The RouteCache wasn't added to the queue.")
	}
}

func TestInsertIntoTheDBFromQueueShouldMarkRouteAsCommittedInCache(t *testing.T) {
	repository.OpenDB()
	repository.CreateScheme()

	queue := DatabaseQueue{}
	queue.Setup(5)

	routeCache := model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      4,
			DestinationPortId: 23,
			Steps: []model.Position{
				{X: 3, Y: 3, StepOrder: 0},
				{X: 3, Y: 4, StepOrder: 1},
				{X: 3, Y: 5, StepOrder: 2},
				{X: 3, Y: 6, StepOrder: 3},
			},
			Committed: false,
		},
		PlannedDestinationPortId: 23,
		Discovered:               false,
		ShipId:                   test_data.TestShipIds[0],
	}

	queue.addRoute(&routeCache)

	element := RouteCache[test_data.TestShipIds[0]]

	if !element.TableData.Committed {
		t.Fatal("The RouteCache wasn't marked as committed.")
	}

	wipeDb()
	repository.Database.Close()

}
