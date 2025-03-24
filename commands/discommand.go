package commands

import (
	"github.com/bwmarrin/discordgo"
)

type DetailedDescription struct {
	Usage    string
	Examples []string
}

type Command struct {
	Name                string
	Aliases             []string
	Description         string
	DetailedDescription DetailedDescription
}

type DiscommandStruct struct {
	Commands map[string]Command
	Aliases  map[string]string
}

func new() *DiscommandStruct {
	discommand := DiscommandStruct{
		Commands: map[string]Command{},
		Aliases:  map[string]string{},
	}

	discommand.loadCommands(HelpCommand)
	return &discommand
}

func (d *DiscommandStruct) loadCommands(command Command) {
	d.Commands[command.Name] = command
	d.Aliases[command.Name] = command.Name

	for _, alias := range command.Aliases {
		d.Aliases[alias] = command.Name
	}
}

func (d *DiscommandStruct) MessageRun(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// 극한의 하드코딩 으아악
	switch command {
	case "도움말":
		HelpCommand.MessageRun(s, m)
	}
}

var Discommand *DiscommandStruct = new()
