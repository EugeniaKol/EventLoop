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
			fmt.Println("COMMAND: ", command)
			fmt.Println("ARGS: ", split[1])
			commands = append(commands, &printCmd{msg: split[1]})
		case "add":
			fmt.Println("COMMAND: ", command)
			fmt.Println("ARGS: ", split[1], split[2])

			a, err := strconv.Atoi(split[1])
			b, err := strconv.Atoi(split[2])
			commands = append(commands, &addCmd{a: a, b: b})

			if err != nil {
				fmt.Println(err)
			}

		}
	}
	return commands
}

// func main() {
// 	Parse("testfile")
// }
