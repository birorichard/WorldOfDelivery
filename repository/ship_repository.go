package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/birorichard/WorldOfDelivery/common"
	"github.com/birorichard/WorldOfDelivery/model"
	_ "github.com/proullon/ramsql/driver"
)

var Database *sql.DB

func OpenDB() {
	var err error
	Database, err = sql.Open("ramsql", "WorldOfDeliveryDB")
	common.HandleError(err, "SQL Open error")
}

func CreateScheme() {
	queryString := []string{
		`CREATE TABLE Route (SourcePortId INT, DestinationPortId INT, PosX INT, PosY INT, StepOrder INT);`,
	}
	var err error
	for _, query := range queryString {
		_, err = Database.Exec(query)
		common.HandleError(err, "Create scheme failed")
	}
}

// Adds one ShipRouteCache instance to the database
func AddRoute(dto *model.ShipRouteCache) {
	sourceId := dto.TableData.SourcePortId
	destinationId := dto.TableData.DestinationPortId
	queryString := "INSERT INTO Route (SourcePortId, DestinationPortId, PosX, PosY, StepOrder) VALUES "
	values := []string{}
	for index, element := range dto.TableData.Steps {
		values = append(values, (fmt.Sprintf("(%d, %d, %d, %d, %d)", sourceId, destinationId, element.X, element.Y, index)))
	}

	finalQuery := queryString + strings.Join(values[:], ", ")
	_, err := Database.Exec(finalQuery)
	common.HandleError(err, fmt.Sprintf("sql.Exec: Query: %s", finalQuery))

}

func GetRoute(portId int, destinationPortId int) model.ShipRouteDTO {
	queryString := `SELECT PosX, PosY FROM Route WHERE SourcePortId = $1 AND DestinationPortId = $2 ORDER BY StepOrder ASC`

	rows, err := Database.Query(queryString, portId, destinationPortId)
	common.HandleError(err, fmt.Sprintf("sql.Exec: Query: %s", queryString))

	var steps []model.Position = make([]model.Position, 0)
	for rows.Next() {
		var posX int
		var posY int
		if DbError := rows.Scan(&posX, &posY); DbError != nil {
			fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", queryString, DbError)
		}

		steps = append(steps, model.Position{X: posX, Y: posY})
	}

	return model.ShipRouteDTO{SourcePortId: portId, DestinationPortId: destinationPortId, Steps: steps}

}

func GetAllRoutes() []model.ShipRouteDTO {
	var routeDtos []model.ShipRouteDTO = make([]model.ShipRouteDTO, 0)
	queryString := `SELECT SourcePortId, DestinationPortId, PosX, PosY, StepOrder FROM Route ORDER BY SourcePortId, DestinationPortId, StepOrder ASC`

	rows, err := Database.Query(queryString)
	common.HandleError(err, fmt.Sprintf("sql.Exec: Query: %s", queryString))

	previousPortId := -1
	previousDestinationPortId := -1
	steps := make([]model.Position, 0)
	for rows.Next() {
		var nextPortId int
		var nextDestinationPortId int
		var posX int
		var posY int
		var stepOrder int

		if DbError := rows.Scan(&nextPortId, &nextDestinationPortId, &posX, &posY, &stepOrder); DbError != nil {
			fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", queryString, DbError)
		}

		if (previousPortId == -1 && previousDestinationPortId == -1) || previousPortId == nextPortId && previousDestinationPortId == nextDestinationPortId {
			steps = append(steps, model.Position{X: posX, Y: posY, StepOrder: stepOrder})
		} else {
			routeDtos = append(routeDtos, model.ShipRouteDTO{SourcePortId: previousPortId, DestinationPortId: previousDestinationPortId, Steps: steps, Commited: true})
			steps = []model.Position{{X: posX, Y: posY, StepOrder: stepOrder}}
		}

		previousPortId = nextPortId
		previousDestinationPortId = nextDestinationPortId
	}

	return routeDtos

}
