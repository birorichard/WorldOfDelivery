package logging

import (
	"fmt"
	"time"

	"github.com/birorichard/WorldOfDelivery/counter"
)

func StartRequestCountLogging() {

	for range time.Tick(time.Second) {
		go logRequestCountToStdOut()
	}
}

func logRequestCountToStdOut() {
	fmt.Println(counter.GetRequestCount(), " / second")
	counter.Reset()
}
