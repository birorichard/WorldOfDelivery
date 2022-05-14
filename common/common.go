package common

import (
	"fmt"
	"testing"
)

func HandleError(err error, message string) {
	if err != nil {
		fmt.Printf("Error: \"%s\" | Message: \"%s\"\n", message, err)
	}
}

func HandleErrorForTesting(t *testing.T, err error, message string) {
	if err != nil {
		t.Fatalf("Error: \"%s\" | Message: \"%s\"\n", message, err)
	}
}
