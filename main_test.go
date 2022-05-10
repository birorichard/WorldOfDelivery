package main

import (
	"testing"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/repositories"
	"github.com/birorichard/WorldOfDelivery/service"
)

func TestShipLeavePort(t *testing.T) {

	repositories.OpenDB()
	repositories.CreateScheme()

	var shipIds = []string{
		"01CY1J41CYM2CDEPYKJRNQMGP9",
		"1133C91092E64F9EA9738A2B01C",
		"1444B3A1E0864AFDB2D28B8BFE",
	}

	var shipLeavePorts = []model.ShipLeavePortDTO{
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[0], X: 0, Y: 0,
			},
			PortId:          0,
			DestinationPort: 91},
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[1], X: 256, Y: 34,
			},
			PortId:          82,
			DestinationPort: 97},
		{
			ShipPositionDTO: model.ShipPositionDTO{
				ShipId: shipIds[2], X: 1022, Y: 987,
			},
			PortId:          7,
			DestinationPort: 125},
	}

	var movements = []model.ShipPositionDTO{
		{ShipId: shipIds[2], X: 1023, Y: 987},
		{ShipId: shipIds[0], X: 1, Y: 1},
		{ShipId: shipIds[1], X: 257, Y: 34},
		{ShipId: shipIds[0], X: 2, Y: 2},
		{ShipId: shipIds[2], X: 1022, Y: 987},
		{ShipId: shipIds[1], X: 258, Y: 34},
		{ShipId: shipIds[0], X: 2, Y: 3},
		{ShipId: shipIds[0], X: 3, Y: 3},
		{ShipId: shipIds[2], X: 1022, Y: 988},
		{ShipId: shipIds[0], X: 3, Y: 4},
		{ShipId: shipIds[1], X: 258, Y: 35},
		{ShipId: shipIds[0], X: 4, Y: 4},
	}

	var shipReachedDestinations = []model.ShipReachedDestinationDTO{
		{ShipPositionDTO: model.ShipPositionDTO{
			ShipId: shipIds[0], X: 4, Y: 6,
		},
			PortId: 91, Payment: 420000},
	}

	for _, element := range shipLeavePorts {
		service.StartTrackingShip(&element)
	}

	for _, element := range movements {
		service.RegisterMove(&element)
	}

	for _, element := range shipReachedDestinations {
		service.CloseRoute(&element)
	}

	if len(service.StoredShipRoutes) != 3 {
		t.Errorf("Not enough ships")
	}

	repositories.GetRoutes()

	// if route, ok := service.StoredShipRoutes[shipIds[0]]; ok {

	// }
	// else
	// {
	// 	t.Errorf("Main ship not found")
	// }

}

// func TestShip(t *testing.T) {
// 	total := 5
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
// 	}
// }

// type Client struct {
// 	url string
// }

// func NewClient(url string) Client {
// 	return Client{url}
// }

// func (c Client) ShipLeavePort(dto model.ShipLeavePortDTO) (string, error) {
// 	json_data, err := json.Marshal(dto)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	res, err := http.Post("radio/ShipLeavePort", "application/json", bytes.NewBuffer(json_data))
// 	if err != nil {
// 		return "", errors.Wrap(err, "unable to complete Get request")
// 	}
// 	defer res.Body.Close()
// 	out, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return "", errors.Wrap(err, "unable to read response data")
// 	}

// 	return string(out), nil
// }

// func TestClientUpperCase(t *testing.T) {
// 	expected := "dummy data"
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, expected)
// 	}))
// 	defer svr.Close()
// 	c := NewClient(svr.URL)
// 	res, err := c.UpperCase("anything")
// 	if err != nil {
// 		t.Errorf("expected err to be nil got %v", err)
// 	}
// 	// res: expected\r\n
// 	// due to the http protocol cleanup response
// 	res = strings.TrimSpace(res)
// 	if res != expected {
// 		t.Errorf("expected res to be %s got %s", expected, res)
// 	}
// }
