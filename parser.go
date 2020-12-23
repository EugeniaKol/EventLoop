package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//TODO: функция Parse должна возвращать []Command
//Раскомментируйте main и запустите сам parser.go чтобы проверить

func Parse(file string) {
	raw, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer raw.Close()

	scanner := bufio.NewScanner(raw)
	scanner.Split(bufio.ScanWords)
	fmt.Println("Распарсило:")

	//команда add принимает 2 аргумента
	var args []int

	//проверка, что первое слово - add
	scanner.Scan()
	item := scanner.Text()
	fmt.Println(item)

	if item != "add" {
		fmt.Println("Invalid command: ", item)
		return
	}

	for scanner.Scan() {
		item = scanner.Text()
		//fmt.Println(item)

		item, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Invalid argument type: ", item, " ", reflect.TypeOf(item))
			return
		}

		args = append(args, item)
	}
	fmt.Println(args)

	if len(args) != 2 {
		fmt.Println("Wrong number of arguments: ", len(args))
		return
	}

	//TODO: return две команды - add и print
	return
}

//func main(){
//	Parse("testfile")
//}