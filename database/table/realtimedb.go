package table

import (
	"errors"
	"firebase.google.com/go/db"
	"github.com/approvers/go-vcnotify/types/guild"
	"github.com/approvers/go-vcnotify/database/operator"
	"log"
)


type FirebaseDataTable struct {
	guilds        map[string]guild.Guild
	operator      operator.DBOperatorInterface
	isInitialized bool
}


func NewRealtimeDBTable(client *db.Client) (table FirebaseDataTable) {
	op := operator.NewRealtimeDBOperator(client)
	table.operator = operator.DBOperatorInterface(op)

	return table
}


func (table *FirebaseDataTable) Initialize() {
	err := table.fetchAll()
	if err != nil {
		log.Panicf("Error: Failed to initialize datatable. Please check the credential and the URL are correct." +
			"Details: %s", err)
	}

	table.isInitialized = true
}


func (table FirebaseDataTable) GetAll() (err error, guilds map[string]guild.Guild) {
	err = initializedChecker(table)
	if err != nil {
		return err, nil
	}

	return nil, table.guilds
}


func (table FirebaseDataTable) Get(guildID string) (err error, guild guild.Guild) {
	err = initializedChecker(table)
	if err != nil {
		return err, guild
	}

	var ok bool
	guild, ok = table.guilds[guildID]
	if ok {
		return nil, guild
	}

	err, guild = table.operator.Get(guildID)

	return err, guild
}


func (table *FirebaseDataTable) Set(guild guild.Guild) (err error) {
	err = initializedChecker(*table)
	if err != nil {
		return err
	}

	err = table.operator.Set(guild)
	if err != nil {
		return err
	}

	err = table.fetchAll()

	return nil
}


func (table *FirebaseDataTable) Delete(guildID string) (err error) {
	err = initializedChecker(*table)
	if err != nil {
		return err
	}

	err = table.operator.Delete(guildID)
	if err != nil {
		return err
	}

	err = table.fetchAll()

	return err
}


func initializedChecker(table FirebaseDataTable) (err error) {
	if !table.isInitialized {
		err = errors.New("error: You tried to use this method before the table is initialized")
		return err
	}

	return nil
}


func (table *FirebaseDataTable) fetchAll() (err error) {
	err, table.guilds = table.operator.GetAll()
	if err != nil {
		return err
	}

	return nil
}
