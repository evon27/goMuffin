package commands

import (
	"github.com/bwmarrin/discordgo"
)

type messageRun func(s *discordgo.Session, m *discordgo.MessageCreate)

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
	Commands    map[string]Command
	Aliases     map[string]string
	messageRuns map[string]interface{}
}

func new() *DiscommandStruct {
	discommand := DiscommandStruct{
		Commands:    map[string]Command{},
		Aliases:     map[string]string{},
		messageRuns: map[string]interface{}{},
	}

	discommand.loadCommands(HelpCommand)

	discommand.addMessageRun(HelpCommand.Name, HelpCommand.helpMessageRun)
	return &discommand
}

func (d *DiscommandStruct) loadCommands(command Command) {
	d.Commands[command.Name] = command
	d.Aliases[command.Name] = command.Name

	for _, alias := range command.Aliases {
		d.Aliases[alias] = command.Name
	}
}

func (d *DiscommandStruct) addMessageRun(name string, run messageRun) {
	d.messageRuns[name] = run
}

func (d *DiscommandStruct) MessageRun(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// 더욱 나아진
	d.messageRuns[command].(messageRun)(s, m)
}

var Discommand *DiscommandStruct = new()
