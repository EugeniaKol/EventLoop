package main

import (
	"fmt"
	"strconv"
)

//container for print function
type printCmd struct {
	msg string
}

//container for add function
type addCmd struct {
	a, b int
}

//execute printing
func (pCmd *printCmd) Execute(h Handler) {
	fmt.Println(pCmd.msg)
}

//execute addition
func (addCmd *addCmd) Execute(h Handler) {
	res := addCmd.a + addCmd.b
	h.Post(&printCmd{
		msg: strconv.Itoa(res)})
}
