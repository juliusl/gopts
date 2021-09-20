package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/juliusl/gopts/pkg/opts"
)

var (
	a, b string
)

func init() {
	err := opts.Parse(Options, Usage)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}

func main() {

	fmt.Println(a, b)
}

func Options(option, value string) error {
	switch option {
	case "a", "all":
		a = value
	case "b", "ball":
		if value == "test2" {
			return errors.New("test2 is not an allowed value")
		}
		b = value
	}

	return nil
}

func Usage(option string) string {
	switch option {
	case "a", "all":
		return "usage: -a|--all <any value>"
	case "b", "ball":
		return "usage: -b|--ball <any value except test2>"
	default:
		return "unknown option"
	}
}
