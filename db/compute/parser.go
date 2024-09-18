package compute

import (
	"errors"
	"regexp"
	"strings"
)

type Command struct {
	Action string
	Args   []string
}

var (
	argumentRegex = regexp.MustCompile(`^(\w+)$`)
)

func Parse(input string) (*Command, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return nil, errors.New("empty command")
	}

	action := tokens[0]

	if action != "GET" && action != "SET" && action != "DEL" {
		return nil, errors.New("invalid command")
	}

	args := tokens[1:]
	for _, arg := range args {
		if !argumentRegex.MatchString(arg) {
			return nil, errors.New("invalid argument: " + arg)
		}
	}

	switch action {
	case "SET":
		if len(args) != 2 {
			return nil, errors.New("invalid number of arguments")
		}
	case "GET", "DEL":
		if len(args) != 1 {
			return nil, errors.New("invalid number of arguments")
		}
	}

	return &Command{
		Action: action,
		Args:   args,
	}, nil
}
