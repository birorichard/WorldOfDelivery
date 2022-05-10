package repositories

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
	dbCreateQueries := []string{
		`CREATE TABLE Routes (SourcePortId INT, DestinationPortId INT, PosX INT, PosY INT, StepOrder INT);`,
		// `CREATE TABLE Routes (SourcePortId INT, DestinationPortId INT, PosX INT, PosY INT, StepOrder INT, PRIMARY KEY(SourcePortId, DestinationPortId));`,
	}

	for _, query := range dbCreateQueries {
		_, DbError = Database.Exec(query)
		if DbError != nil {
			fmt.Printf("SQL DB Create: Error: %s\n", DbError)
		}
	}
}

func AddRoute(dto *model.ShipRoute) {
	sourceId := dto.SourcePortId
	destinationId := dto.DestinationPortId
	queryString := "INSERT INTO Routes (SourcePortId, DestinationPortId, PosX, PosY, StepOrder) VALUES "
	values := []string{}
	for index, element := range dto.Steps {
		values = append(values, (fmt.Sprintf("(%d, %d, %d, %d, %d)", sourceId, destinationId, element.X, element.Y, index)))
	}

	finalQuery := queryString + strings.Join(values[:], ", ")
	_, DbError := Database.Exec(finalQuery)
	if DbError != nil {
		fmt.Printf("SQL DB AddRoute method: Error: %s\n", DbError)
	}
}

func GetRoutes(portId int) {
	query := `SELECT * FROM Routes WHERE SourcePortId = $1`

	rows, error := Database.Query(query, portId)
	if error != nil {
		fmt.Printf("sql.Exec: Error: %s\n", error)
	}

	for rows.Next() {
		var sourceId int
		var destinationId int
		var posX int
		var posY int
		var stepOrder int

		if err := rows.Scan(&sourceId, &destinationId, &posX, &posY, &stepOrder); err != nil {
			fmt.Printf("sql.Exec: Error: %s\n", err)
		}
		fmt.Printf("%d, %d, %d, %d, %d\n", sourceId, destinationId, posX, posY, stepOrder)
	}
}
