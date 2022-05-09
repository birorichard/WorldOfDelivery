package main

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
