package main

type Parser interface {
	Parse(b []byte) (Router, error)
}
