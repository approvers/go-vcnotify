package setup

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	handler "../handlers"
	utils "../utils"
)


func Setup() {
	discord, err := discordgo.New()

	if err != nil {
		displayError := fmt.Sprintf("Error: Unknown error has occured while creating discord session.\n" +
			"Detail:%s\n", err)
		panic(displayError)
	}

	discord.Token = utils.GetToken()

	discord.AddHandler(handler.OnMessageCreate)
	discord.AddHandler(handler.OnVoiceStateChanged)

	if err = discord.Open(); err != nil {
		displayError := fmt.Sprintf("Error: Unknown error has occured while opening discord session.\n" +
			"Detail:%s\n", err)
		panic(displayError)
	}

	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	fmt.Printf("Info: Succeed to boot at %s\n", utils.GetCurrentTimeOfJST())

	<-sc
	return
}
