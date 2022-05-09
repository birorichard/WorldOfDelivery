package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/birorichard/WorldOfDelivery/handlers"
	"github.com/birorichard/WorldOfDelivery/logging"
	"github.com/birorichard/WorldOfDelivery/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	db, err := sql.Open("ramsql", "InMemoryDb")
	if err != nil {
		fmt.Printf("sql.Open : Error : %s\n", err)
	}

	defer db.Close()

	batch := []string{
		`CREATE TABLE address (id BIGSERIAL PRIMARY KEY, street TEXT, street_number INT);`,
		`CREATE TABLE user_addresses (address_id INT, user_id INT);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Victor Hugo', 32);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la République', 23);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Charles Martel', 5);`,
		`INSERT INTO address (street, street_number) VALUES ('chemin du bout du monde ', 323);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la liberté', 2);`,
		`INSERT INTO address (street, street_number) VALUES ('avenue des champs', 12);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 2);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 3);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 4);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 5);`,
	}

	for _, b := range batch {
		_, err = db.Exec(b)
		if err != nil {
			fmt.Printf("sql.Exec: Error: %s\n", err)
		}
	}

	router := chi.NewRouter()

	router.Use(middleware.CreateRequestCounterMiddleware)
	router.Route("/ship", handlers.ShipHandler)
	router.Route("/radio", handlers.RadioHandler)
	router.Route("/missile", handlers.MissileHandler)

	go logging.StartRequestCountLogging()

	query := `SELECT address.street_number, address.street FROM address 
							JOIN user_addresses ON address.id=user_addresses.address_id 
							WHERE user_addresses.user_id = $1;`

	rows, err := db.Query(query, 1)
	if err != nil {
		fmt.Printf("sql.Exec: Error: %s\n", err)
	}

	var addresses []string
	for rows.Next() {
		var number int
		var street string
		if err := rows.Scan(&number, &street); err != nil {
			fmt.Printf("sql.Exec: Error: %s\n", err)
		}
		fmt.Printf("%d %s", number, street)
		addresses = append(addresses, fmt.Sprintf("%d %s", number, street))
	}

	http.ListenAndServe(":8080", router)

}
