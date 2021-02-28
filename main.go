package main

import (
	"github.com/approvers/go-vcnotify/internal/exit"
	"github.com/approvers/go-vcnotify/internal/setup"
	"log"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	discord := setup.InitializeSession()
	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	log.Println("Info: Succeed to boot!")

	<-sc
	err := exit.PreExitProcess()
	if err != nil {
		log.Printf("Error: Error occured while exiting. Detail: %s", err)
	}
	return
}
