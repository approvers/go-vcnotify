package guild


func CreateFromIDs(guildID string, voiceDiffID string) (guild Guild) {
	guild = defaultGuild
	guild.GuildID = guildID
	guild.Config.VoiceDiffChID = voiceDiffID

	return guild
}
