package embed

import (
	"bytes"
	"github.com/approvers/go-vcnotify/types/guild"
	"github.com/bwmarrin/discordgo"
	"text/template"
)


func CreateEmbed(embedContent guild.CustomEmbedContent, formatContent FormatContent) (embed discordgo.MessageEmbed, err error){
	var description string

	description, err = makeDescription(embedContent, formatContent)
	if err != nil {
		return embed, err
	}

	embed.Description = description
	embed.Title = embedContent.Title
	embed.Color = 0x1e63e9

	return embed, nil
}


func makeDescription(embedContent guild.CustomEmbedContent, formatContent FormatContent) (result string, err error) {
	var msgResult bytes.Buffer

	msgTpl, err := template.New("customGuildEmbedTemplate").Parse(embedContent.Description)
	err = msgTpl.Execute(&msgResult, formatContent)
	if err != nil {
		return result, err
	}

	return msgResult.String(), nil
}
