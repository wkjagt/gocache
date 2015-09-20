package cache

import (
    "fmt"
		"strings"
)

var memoryMap = make(map[string]string)

type Command struct {
  name string
	params []string
}

func (c *Command) Handle() (string, error) {
	switch c.name {
	case "SET":
		memoryMap[c.params[0]] = c.params[1]
		return "OK", nil
	case "GET":
		return memoryMap[c.params[0]], nil
	case "DELETE":
		delete(memoryMap, c.params[0])
		return "OK", nil
	}
	return "", fmt.Errorf("Unknow command: '%s'", c.name)
}

func ExtractCommand(payLoad string) (Command, error) {
	commandParts := strings.Split(payLoad, " ")

	return Command{commandParts[0], commandParts[1:]}, nil
}
