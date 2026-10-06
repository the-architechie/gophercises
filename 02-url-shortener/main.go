package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type parseFunc func([]byte) (Router, error)

var parsers = map[string]parseFunc{
	".yaml": parseYML,
	".yml":  parseYML,
	".json": parseJson,
}

func main() {

	file := flag.String("routes", "examples/paths.yaml", "routes file(.yaml, .yml or .json)")
	format := flag.String("format", "", "foce format : yaml | json(default from file extension)")
	addr := flag.String("addr", ":8080", "port address")
	flag.Parse()

	parse, err := parserFor(*file, *format)
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	r, err := parse(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting the server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, r.MapHandler()))
}

func parserFor(file, format string) (parseFunc, error) {
	key := filepath.Ext(file)
	if format != "" {
		key = "." + format
	}
	newParser, ok := parsers[key]
	if !ok {
		return nil, fmt.Errorf("unsupported format %q", key)
	}
	return newParser, nil
}
