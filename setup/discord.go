package setup

import (
	"github.com/bwmarrin/discordgo"
	"log"

	"github.com/approvers/go-vcnotify/handlers"
	"github.com/approvers/go-vcnotify/utils"
)


func GetDiscordSession() *discordgo.Session {
	discord, err := discordgo.New()

	if err != nil {
		log.Panicf("Error: Unknown error has occured while creating discord session.\n" +
			"Detail:%s\n", err)
	}

	discord.Token = utils.GetToken()

	discord.AddHandler(handlers.OnMessageCreate)
	discord.AddHandler(handlers.OnVoiceStateChanged)

	if err = discord.Open(); err != nil {
		log.Panicf("Error: Unknown error has occured while opening discord session.\n" +
			"Detail:%s\n", err)
	}

	return discord
}
