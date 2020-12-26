package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Parse parses file
func Parse(file string) []Command {
	raw, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return []Command{}
	}
	defer raw.Close()

	scanner := bufio.NewScanner(raw)
	//scanner.Split(bufio.ScanWords)
	//fmt.Println("Распарсило:")

	//команда add принимает 2 аргумента
	//var argsAdd []int
	var commands []Command
	//var add Command

	// //проверка, что первое слово - add
	// scanner.Scan()
	// item := scanner.Text()

	// if item != "add" {
	// 	fmt.Println("Invalid command: ", item)
	// 	return
	// }

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
		//fmt.Println(item)

		// item, err := strconv.Atoi(item)
		// if err != nil {
		// 	fmt.Println("Invalid argument type: ", item, " ", reflect.TypeOf(item))
		// 	return
		// }

		// args = append(args, item)
	}
	// fmt.Println("ARGS: ", args)
	// //fmt.Printf("%+v\n", "add command ")

	// if len(args) != 2 {
	// 	fmt.Println("Wrong number of arguments: ", len(args))
	// 	return
	// }

	//TODO: return две команды - add и print
	return commands
}

// func main() {
// 	Parse("testfile")
// }
