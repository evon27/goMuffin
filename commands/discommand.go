package commands

import (
	// "fmt"

	"github.com/bwmarrin/discordgo"
)

type messageRun func(s *discordgo.Session, m *discordgo.MessageCreate)
type chatInputRun func(s *discordgo.Session, m *discordgo.InteractionCreate)

type DetailedDescription struct {
	Usage    string
	Examples []string
}

type Command struct {
	*discordgo.ApplicationCommand
	Aliases             []string
	DetailedDescription *DetailedDescription
	discommand          *DiscommandStruct
}

type DiscommandStruct struct {
	Commands      map[string]*Command
	Aliases       map[string]string
	messageRuns   map[string]messageRun
	chatInputRuns map[string]chatInputRun
}

func new() *DiscommandStruct {
	discommand := DiscommandStruct{
		Commands:      map[string]*Command{},
		Aliases:       map[string]string{},
		messageRuns:   map[string]messageRun{},
		chatInputRuns: map[string]chatInputRun{},
	}

	go discommand.loadCommands(HelpCommand)
	go discommand.loadCommands(DataLengthCommand)
	go discommand.loadCommands(LearnCommand)
	go discommand.loadCommands(LearnedDataListCommand)
	go discommand.loadCommands(InformationCommand)

	go discommand.addMessageRun(HelpCommand.Name, HelpCommand.helpMessageRun)
	go discommand.addMessageRun(DataLengthCommand.Name, DataLengthCommand.dataLengthMessageRun)
	go discommand.addMessageRun(LearnCommand.Name, LearnCommand.learnMessageRun)
	go discommand.addMessageRun(LearnedDataListCommand.Name, LearnedDataListCommand.learnedDataListMessageRun)
	go discommand.addMessageRun(InformationCommand.Name, InformationCommand.informationMessageRun)

	go discommand.addChatInputRun(DataLengthCommand.Name, DataLengthCommand.dataLenghChatInputRun)
	go discommand.addChatInputRun(LearnCommand.Name, LearnCommand.learnChatInputRun)
	go discommand.addChatInputRun(LearnedDataListCommand.Name, LearnedDataListCommand.learnedDataListChatInputRun)
	go discommand.addChatInputRun(InformationCommand.Name, DataLengthCommand.informationChatInputRun)
	return &discommand
}

func (d *DiscommandStruct) loadCommands(command *Command) {
	d.Commands[command.Name] = command
	d.Aliases[command.Name] = command.Name
	command.discommand = d

	for _, alias := range command.Aliases {
		d.Aliases[alias] = command.Name
	}
}

func (d *DiscommandStruct) addMessageRun(name string, run messageRun) {
	d.messageRuns[name] = run
}

func (d *DiscommandStruct) addChatInputRun(name string, run chatInputRun) {
	d.chatInputRuns[name] = run
}

func (d *DiscommandStruct) MessageRun(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// 더욱 나아진
	d.messageRuns[command](s, m)
}

func (d *DiscommandStruct) ChatInputRun(command string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	d.chatInputRuns[command](s, i)
}

var Discommand *DiscommandStruct = new()
