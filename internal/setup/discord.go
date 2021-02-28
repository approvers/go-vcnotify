package setup

import (
	"github.com/approvers/go-vcnotify/config"
	"github.com/bwmarrin/discordgo"
	"log"

	"github.com/approvers/go-vcnotify/handlers"
)

var (
	discordToken = config.GetEnvironmentVariable(config.DiscordTokenEnvironmentName)
)


func InitializeSession() *discordgo.Session {
	discord, err := discordgo.New()
	if err != nil {
		log.Panicf("Error: Unknown error has occured while creating discord session.\n" +
			"Detail:%s\n", err)
	}

	addHandlers(discord)

	discord.Token = "Bot " + discordToken

	openSession(discord)

	return discord
}


func addHandlers(discord *discordgo.Session) {
	discord.AddHandler(handlers.OnMessageCreate)
	discord.AddHandler(handlers.OnVoiceStateChanged)
}


func openSession(discord *discordgo.Session) {
	if err := discord.Open(); err != nil {
		log.Panicf("Error: Unknown error has occured while opening discord session.\n" +
			"Detail:%s\n", err)
	}
}
