package repository

import (
	"fmt"
	"testing"

	"github.com/birorichard/WorldOfDelivery/common"
	"github.com/birorichard/WorldOfDelivery/model"
)

var testShipIds = []string{"01CY1J41CYM2CDEPYKJRNQMGP9", "1133C91092E64F9EA9738A2B01", "1444B3A1E0864AFDB2D28B8BFE"}

func wipeDb() {
	queryString := []string{
		`DROP TABLE Route;`,
	}
	var err error
	for _, query := range queryString {
		_, err = Database.Exec(query)
		common.HandleError(err, "Drop scheme failed")
	}
}

func TestSchemeSuccesfullyCreated(t *testing.T) {
	OpenDB()
	CreateScheme()

	queryString := "INSERT INTO Route (SourcePortId, DestinationPortId, PosX, PosY, StepOrder) VALUES (1, 1, 1, 1, 1)"
	_, err := Database.Exec(queryString)
	common.HandleErrorForTesting(t, err, fmt.Sprintf("sql.Exec: Query: %s", queryString))
	wipeDb()
}

func TestAddRouteAddsRoute(t *testing.T) {

	OpenDB()
	CreateScheme()

	entity := model.ShipRouteCache{
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
		ShipId:                   testShipIds[0],
	}

	AddRoute(&entity)

	queryString := `SELECT * FROM Route;`

	rows, err := Database.Query(queryString)
	common.HandleErrorForTesting(t, err, fmt.Sprintf("sql.Exec: Query: %s", queryString))

	for rows.Next() {
		var portId int
		var destinationPortId int
		var posX int
		var posY int
		var stepOrder int

		err := rows.Scan(&portId, &destinationPortId, &posX, &posY, &stepOrder)
		common.HandleErrorForTesting(t, err, fmt.Sprintf("sql.Exec: Query: %s", queryString))

		if portId != entity.TableData.SourcePortId ||
			destinationPortId != entity.TableData.DestinationPortId ||
			posX != entity.TableData.Steps[stepOrder].X ||
			posY != entity.TableData.Steps[stepOrder].Y {
			t.Fatal("Route wasn't added properly.")
		}

	}
	wipeDb()
}

func TestGetRouteReturnsTheProperRoute(t *testing.T) {

	OpenDB()
	CreateScheme()

	queryString := `INSERT INTO Route (SourcePortId, DestinationPortId, PosX, PosY, StepOrder) VALUES 
	(1, 1, 1, 1, 0), (1, 1, 1, 2, 1), (1, 1, 1, 3, 2), (2, 5, 11, 1, 0), (2, 5, 11, 2, 1), (2, 5, 11, 3, 2)`

	_, err := Database.Exec(queryString)
	common.HandleErrorForTesting(t, err, fmt.Sprintf("sql.Exec: Query: %s", queryString))

	routeDto := GetRoute(2, 5)

	if routeDto.SourcePortId != 2 || routeDto.DestinationPortId != 5 {
		t.Fatal("GetRoute didn't return the proper route object")
	}
	wipeDb()
}
