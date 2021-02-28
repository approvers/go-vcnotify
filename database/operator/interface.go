package operator

import "github.com/approvers/go-vcnotify/types/guild"

type DBOperatorInterface interface {
	GetAll() (err error, guilds map[string]guild.Guild)
	Get(entry string) (err error, guild guild.Guild)
	Set(entry guild.Guild) (err error)
	Delete(guildID string) (err error)
}
