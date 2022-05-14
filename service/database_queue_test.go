package service

import (
	"testing"

	"github.com/birorichard/WorldOfDelivery/common"
	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repository"
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

var testShipIds = []string{"01CY1J41CYM2CDEPYKJRNQMGP9", "1133C91092E64F9EA9738A2B01", "1444B3A1E0864AFDB2D28B8BFE"}

func Test(t *testing.T) {
	t.Error()
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
			Commited: false,
		},
		PlannedDestinationPortId: 23,
		Discovered:               false,
		ShipId:                   testShipIds[0],
	}

	queue.Add(&routeCache)

	if queue.GetSize() != 1 {
		t.Fatal("The RouteCache wasn't added to the queue.")
	}
}

func TestInsertIntoTheDBFroMQueueShouldMarkRouteAsCommitedInCache(t *testing.T) {
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
			Commited: false,
		},
		PlannedDestinationPortId: 23,
		Discovered:               false,
		ShipId:                   testShipIds[0],
	}

	queue.addRoute(&routeCache)

	element := RouteCache[testShipIds[0]]

	if !element.TableData.Commited {
		t.Fatal("The RouteCache wasn't marked as commited.")
	}

	wipeDb()
	repository.Database.Close()

}
