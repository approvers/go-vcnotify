package handlers

import (
	"github.com/bwmarrin/discordgo"
)


func OnVoiceStateChanged(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	user, _ := session.User(event.UserID)
	if user.Bot {
		return
	}
}
