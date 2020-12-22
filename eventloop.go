package main

import "io"

//print
//add 2 5

type Command interface {
	Execute(h Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Loop struct {

}

func(l Loop) Start(){

}

func (l Loop) Post(cmd Command) {

}

func (l Loop) AwaitFinish() {

}

func parse(input io.Reader) []Command{
	//TODO: parse input into array of commands
}

func main(){
	//use parser, start loop, post parsed commands in a cycle and await finish
}