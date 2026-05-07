package buttrbase

import "fmt"

// ButtrbaseError is returned for non-2xx HTTP responses from the API.
type ButtrbaseError struct {
	StatusCode int
	Detail     string
	Body       []byte
}

func (e *ButtrbaseError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("buttrbase: HTTP %d: %s", e.StatusCode, e.Detail)
	}
	return fmt.Sprintf("buttrbase: HTTP %d", e.StatusCode)
}
