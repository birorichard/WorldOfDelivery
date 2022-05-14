package logging

import (
	"fmt"
	"time"

	"github.com/birorichard/WorldOfDelivery/counter"
	"github.com/birorichard/WorldOfDelivery/service"
)

func ConfigureLogging() {

	for range time.Tick(time.Second) {
		fmt.Println(getLogMessage())
	}
}

func getLogMessage() string {
	message := fmt.Sprintf(
		"Elapsed time in seconds: %d | Request per second: %d | Routes found: %d",
		*counter.ElapsedTimeInSecondsCounter.GetCurrentValue(),
		*counter.RequestCounter.GetCurrentValue(),
		service.GetFoundRoutesCount(),
	)
	counter.ElapsedTimeInSecondsCounter.Increment()
	counter.RequestCounter.Reset()
	return message
}
