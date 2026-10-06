package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func parseYML(data []byte) (Router, error) {
	var entries []pathURL
	err := yaml.Unmarshal(data, &entries)
	if err != nil {
		return Router{}, fmt.Errorf("Could not parse the file  %w", err)
	}
	return buildRouter(entries), nil
}
