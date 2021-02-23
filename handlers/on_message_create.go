package handlers

import (
	"github.com/bwmarrin/discordgo"
)


func OnMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}
}
