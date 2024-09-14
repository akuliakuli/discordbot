package main

import (
	"fmt"

	"github.com/akuliakuli/discordbot/bot"
	"github.com/akuliakuli/discordbot/config"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}