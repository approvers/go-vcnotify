package guild

var defaultGuild = Guild{
	GuildID: "",
	Config: Config{
		VoiceDiffChID: "",
		CustomEmbed: CustomEmbedSet{
			In: CustomEmbedContent{
				Title: "{{.MemberName}}が{{.ChannelName}}に入りました",
				Description: "何かが始まる予感がする。",
			},
			Out: CustomEmbedContent{
				Title: "{{.MemberName}}が{{.ChannelName}}から抜けました",
				Description: "あいつは良い奴だったよ...",
			},
		},
	},
}
