package table

import "github.com/approvers/go-vcnotify/types/guild"

type DBTableInterface interface {
	Initialize()
	GetAll() (err error, guilds map[string]guild.Guild)
	Get(guildID string) (err error, guild guild.Guild)
	Set(guild guild.Guild) (err error)
	Delete(guildID string) (err error)
}
