package hdns

import (
	"fmt"
	"net/http"
)

func Error(acceptedCodes []int, responseCode int) error {
	for _, c := range acceptedCodes {
		if c == responseCode {
			return fmt.Errorf("error from API: %d %s", responseCode, http.StatusText(responseCode))
		}
	}
	return fmt.Errorf("undocumented status code in the response: %d", responseCode)
}
