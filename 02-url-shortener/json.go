package main

import (
	"encoding/json"
	"fmt"
)

func parseJson(data []byte) (Router, error) {
	var entries []pathURL

	err := json.Unmarshal(data, &entries)
	if err != nil {
		return Router{}, fmt.Errorf("failed to parse json: %w", err)
	}
	return buildRouter(entries), nil
}
