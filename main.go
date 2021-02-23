package main

import (
	"fmt"
	"github.com/approvers/go-vcnotify/on_close"
	"github.com/approvers/go-vcnotify/setup"
	"github.com/approvers/go-vcnotify/utils"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	discord := setup.GetDiscordSession()
	defer discord.Close()

	dbClient := setup.GetDBClient()

	fmt.Printf("Info: Succeed to boot at %s\n", utils.GetCurrentTimeOfJST())


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
	on_close.PreClosingProcess()
	return
}
