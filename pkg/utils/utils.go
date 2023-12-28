package utils

import (
	"encoding/json"
	"io" // no longer use io/ioutil
	"net/http"
)

// to parse body
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
