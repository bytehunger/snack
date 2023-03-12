package main

import (
	"errors"
	"flag"
)

// Entrypoint of program
func main() {
	adapterPtr := flag.String("a", "file", "The adapter (default: file)")
	flag.Parse()

	adapter, err := NewAdapter(*adapterPtr, flag.Args())

	if err != nil {
		panic(err)
	}

	generator := NewGenerator(adapter)
	err = generator.Generate()

	if err != nil {
		panic(err)
	}
}

// NewAdapter returns the right adapter for given argument.
func NewAdapter(name string, args []string) (a Adapter, err error) {
	if len(args) != 1 {
		err = errors.New("wrong number of arguments")
		return
	}

	switch name {
	case "file":
		a = &FileAdapter{Source: args[0]}
	case "json":
		a = &JSONAdapter{Source: args[0]}
	case "yaml":
		a = &YAMLAdapter{Source: args[0]}
	default:
		err = errors.New("invalid adapter: " + name)
	}

	return
}

// NewGenerator returns the generator for the given page.
func NewGenerator(a Adapter) Generator {
	return Generator{Adapter: a}
}
