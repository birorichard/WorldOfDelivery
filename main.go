package main

import (
	"net/http"

	"github.com/birorichard/WorldOfDelivery/handlers"
	"github.com/birorichard/WorldOfDelivery/logging"
	"github.com/birorichard/WorldOfDelivery/middleware"
	"github.com/birorichard/WorldOfDelivery/repositories"
	"github.com/go-chi/chi/v5"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.CreateRequestCounterMiddleware)
	router.Route("/ship", handlers.ShipHandler)
	router.Route("/radio", handlers.RadioHandler)
	router.Route("/missile", handlers.MissileHandler)

	defer repositories.Database.Close()

	repositories.OpenDB()
	repositories.CreateScheme()

	go logging.StartRequestCountLogging()

	// query := `SELECT address.street_number, address.street FROM address
	// 						JOIN user_addresses ON address.id=user_addresses.address_id
	// 						WHERE user_addresses.user_id = $1;`

	// rows, error := db.Query(query, 1)
	// if error != nil {
	// 	fmt.Printf("sql.Exec: Error: %s\n", error)
	// }

	// var addresses []string
	// for rows.Next() {
	// 	var number int
	// 	var street string
	// 	if err := rows.Scan(&number, &street); err != nil {
	// 		fmt.Printf("sql.Exec: Error: %s\n", err)
	// 	}
	// 	fmt.Printf("%d %s", number, street)
	// 	addresses = append(addresses, fmt.Sprintf("%d %s", number, street))
	// }

	http.ListenAndServe(":8080", router)

}
