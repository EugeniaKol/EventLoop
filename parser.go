package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parse parses file into commands
func Parse(file string) []Command {
	raw, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return []Command{}
	}
	defer raw.Close()

	scanner := bufio.NewScanner(raw)
	var commands []Command

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		command := split[0]

		switch command {
		case "print":
			phrase := strings.Join(split[1:], " ")

			fmt.Println("COMMAND: ", command)
			fmt.Println("ARGS: ", phrase)
			commands = append(commands, &printCmd{msg: phrase})
		case "add":
			if len(split) != 3 {
				commands = append(commands, &printCmd{msg: "Wrong number of args for add: " + strconv.Itoa(len(split) - 1)})
				break
			}

			fmt.Println("COMMAND: ", command)
			fmt.Println("ARGS: ", split[1], split[2])

			a, err := strconv.Atoi(split[1])
			b, err := strconv.Atoi(split[2])

			if err != nil {
				commands = append(commands, &printCmd{msg: err.Error()})
				break
			}
			commands = append(commands, &addCmd{a: a, b: b})
		default:
			commands = append(commands, &printCmd{msg: "Sorry, i cant " + command + " yet :'("})
		}
	}
	return commands
}

