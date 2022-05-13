package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/birorichard/WorldOfDelivery/model"
)

var Database *sql.DB
var DbError error

func OpenDB() {
	Database, DbError = sql.Open("ramsql", "InMemoryDb")

	if DbError != nil {
		fmt.Printf("SQL OPEN : Error : %s\n", DbError)
	}

}

func CreateScheme() {
	queryString := []string{
		`CREATE TABLE Routes (SourcePortId INT, DestinationPortId INT, PosX INT, PosY INT, StepOrder INT);`,
	}

	for _, query := range queryString {
		_, DbError = Database.Exec(query)
		if DbError != nil {
			fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", query, DbError)
		}
	}
}

// Adds one ShipRouteCache instance to the database
func AddRoute(dto *model.ShipRouteCache) {
	sourceId := dto.TableData.SourcePortId
	destinationId := dto.TableData.DestinationPortId
	queryString := "INSERT INTO Routes (SourcePortId, DestinationPortId, PosX, PosY, StepOrder) VALUES "
	values := []string{}
	for index, element := range dto.TableData.Steps {
		values = append(values, (fmt.Sprintf("(%d, %d, %d, %d, %d)", sourceId, destinationId, element.X, element.Y, index)))
	}

	finalQuery := queryString + strings.Join(values[:], ", ")
	_, DbError := Database.Exec(finalQuery)
	if DbError != nil {
		fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", finalQuery, DbError)
	}
}

func GetRoute(portId int, destinationPortId int) model.ShipRouteDTO {
	queryString := `SELECT PosX, PosY FROM Routes WHERE SourcePortId = $1 AND DestinationPortId = $2 ORDER BY StepOrder ASC`

	rows, DbError := Database.Query(queryString, portId, destinationPortId)
	if DbError != nil {
		fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", queryString, DbError)
	}
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
	queryString := `SELECT SourcePortId, DestinationPortId, PosX, PosY, StepOrder FROM Routes ORDER BY SourcePortId, DestinationPortId, StepOrder ASC`

	rows, DbError := Database.Query(queryString)
	if DbError != nil {
		fmt.Printf("sql.Exec: Query: %s\nError: %s\n\n", queryString, DbError)
	}

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
