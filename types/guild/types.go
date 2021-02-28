package guild

type Guild struct {
	GuildID string `json:"guild_id"`
	Config  Config `json:"config"`
}

type Config struct {
	VoiceDiffChID string         `json:"voice_diff_ch"`
	CustomEmbed   CustomEmbedSet `json:"custom_embed"`
}

type CustomEmbedSet struct {
	In  CustomEmbedContent `json:"in"`
	Out CustomEmbedContent `json:"out"`
}

type CustomEmbedContent struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

/*

root/
  └ "guilds"
    ├ <guildID : string>
      ├ <guildID : string>
	  └ "config"
        ├ "voiceDiffChID" : string
        └ "customEmbed"
          ├ "in"
            ├ "title" : string
            └ "description": string
          └ "out"
            ├ "title" : string
            └ "description": string
    ├ <guildID : string>
    ︙

 */