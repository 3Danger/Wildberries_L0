package main

import (
	"awesomeProject/srcs/Backend"
	"awesomeProject/srcs/Backend/Utils"
	"awesomeProject/srcs/Frontend"
)

func main() {
	config := Utils.ParseArgs()
	backend := Backend.BackEnd(config)
	defer backend.Close()
	Frontend.Handler(&backend.JModelSlice)
}
