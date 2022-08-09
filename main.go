package main

import (
	"Miyoushe_Genshin_AutoSigner/client"
	"Miyoushe_Genshin_AutoSigner/config"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalln("config loading err,err:", err.Error())
	}
	client.Run() // run task
}
