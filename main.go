package main

import (
	"fmt"

	"gitlab.com/DeveloperDurp/durpot/handlers"
)

func main() {

	err := handlers.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
	<-make(chan struct{})
}
