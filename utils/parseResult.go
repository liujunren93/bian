package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseResult(res *http.Response, v any) error {
	if res.StatusCode == 200 {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, v)
	}

	return fmt.Errorf("%v", res)
}
