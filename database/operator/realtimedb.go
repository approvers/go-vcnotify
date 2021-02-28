package operator

import (
	"context"
	"firebase.google.com/go/db"
	"github.com/approvers/go-vcnotify/config"
	"github.com/approvers/go-vcnotify/types/guild"
)

const (
	guildDir = config.RealtimeDBGuildDirectoryName
)

type RealtimeDBOperator struct {
	client *db.Client
}


func NewRealtimeDBOperator(client *db.Client) RealtimeDBOperator {
	return RealtimeDBOperator{client: client}
}


func (operator RealtimeDBOperator) GetAll() (err error, guilds map[string]guild.Guild) {
	ref := operator.client.NewRef(guildDir)

	err = ref.Get(context.Background(), &guilds)
	if err != nil {
		return err, nil
	}

	return nil, guilds
}


func (operator RealtimeDBOperator) Get(entry string) (err error, guild guild.Guild) {
	return nil, guild
}


func (operator RealtimeDBOperator) Set(entry guild.Guild) (err error) {
	ctx := context.Background()
	refToGuilds := operator.client.NewRef(guildDir).Child(entry.GuildID)

	err = refToGuilds.Set(ctx, entry)
	if err != nil {
		return err
	}

	return nil
}


func (operator RealtimeDBOperator) Delete(guildID string) (err error) {
	ctx := context.Background()
	ref := operator.client.NewRef(guildDir).Child(guildID)

	err = ref.Set(ctx, nil)

	return err
}
