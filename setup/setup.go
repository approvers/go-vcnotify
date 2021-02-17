package setup

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"

	"github.com/approvers/go-vcnotify/handlers"
	"github.com/approvers/go-vcnotify/utils"
)


func Setup() *discordgo.Session {
	discord, err := discordgo.New()

	if err != nil {
		displayError := fmt.Sprintf("Error: Unknown error has occured while creating discord session.\n" +
			"Detail:%s\n", err)
		panic(displayError)
	}

	discord.Token = utils.GetToken()

	discord.AddHandler(handlers.OnMessageCreate)
	discord.AddHandler(handlers.OnVoiceStateChanged)

	if err = discord.Open(); err != nil {
		displayError := fmt.Sprintf("Error: Unknown error has occured while opening discord session.\n" +
			"Detail:%s\n", err)
		log.Panicf(displayError)
	}

	return discord
}
