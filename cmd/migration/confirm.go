package migration

import (
	"fmt"
	"strings"

	"github.com/mdmoshiur/example-go/internal/logger"
)

func askForConfirmation(operation string) bool {
	var response string

	operation = strings.ToUpper(operation)
	message := fmt.Sprintf("%s: Please type yes or no for confirmation and then press enter: ", operation)
	fmt.Print(message)

	_, err := fmt.Scanln(&response)
	if err != nil {
		logger.Fatal(err)
	}

	okayResponses := []string{"yes", "Yes", "YES"}
	nokayResponses := []string{"no", "No", "NO"}

	switch {
	case containsString(okayResponses, response):
		return true
	case containsString(nokayResponses, response):
		return false
	default:
		return askForConfirmation(operation)
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
