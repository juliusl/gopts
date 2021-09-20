package opts

import (
	"os"
	"strings"
)

// Usage should return and example usage of an option. Writes to standard error
type Usage = func(option string) string

// Configure should return an error if the option cannot be configured with the value
type Configure = func(option, value string) error

// Parse parses os arguments and passes them to Configure and Usage
func Parse(configure Configure, usage Usage) error {
	if configure == nil {
		configure = Echo
	}

	args := os.Args
	for i := 0; len(args) > 1 && i+2 < len(os.Args); i++ {
		args = os.Args[i+1 : i+3]
		option := strings.TrimLeft(args[0], "-")
		value := args[1]

		err := configure(option, value)
		if err != nil {
			if usage != nil {
				os.Stderr.WriteString(usage(option))
				os.Stderr.WriteString("\n")
			}
			return err
		}
	}

	return nil
}

// Echo is the default config handler. Each line is an option/value pair
func Echo(option, value string) error {
	os.Stdout.WriteString(option)
	os.Stdout.WriteString(" ")
	os.Stdout.WriteString(value)
	os.Stdout.WriteString("\n")
	return nil
}
