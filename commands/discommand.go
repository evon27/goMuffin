package commands

import (
	// "fmt"

	"github.com/bwmarrin/discordgo"
)

type messageRun func(ctx *MsgContext)
type chatInputRun func(s *InterContext)

type Category string

type DetailedDescription struct {
	Usage    string
	Examples []string
}

type Command struct {
	*discordgo.ApplicationCommand
	Aliases             []string
	DetailedDescription *DetailedDescription
	discommand          *DiscommandStruct
	Category            Category
}

type DiscommandStruct struct {
	Commands      map[string]*Command
	Aliases       map[string]string
	messageRuns   map[string]messageRun
	chatInputRuns map[string]chatInputRun
}

type MsgContext struct {
	Session *discordgo.Session
	Msg     *discordgo.MessageCreate
	Args    []string
}

type InterContext struct {
	Session *discordgo.Session
	Inter   *discordgo.InteractionCreate
}

const (
	Chattings Category = "채팅"
	Generals  Category = "일반"
)

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

	go discommand.addChatInputRun(HelpCommand.Name, HelpCommand.helpChatInputRun)
	go discommand.addChatInputRun(DataLengthCommand.Name, DataLengthCommand.dataLenghChatInputRun)
	go discommand.addChatInputRun(LearnCommand.Name, LearnCommand.learnChatInputRun)
	go discommand.addChatInputRun(LearnedDataListCommand.Name, LearnedDataListCommand.learnedDataListChatInputRun)
	go discommand.addChatInputRun(InformationCommand.Name, DataLengthCommand.informationChatInputRun)
	return &discommand
}

func (d *DiscommandStruct) loadCommands(c *Command) {
	d.Commands[c.Name] = c
	d.Aliases[c.Name] = c.Name
	c.discommand = d

	for _, alias := range c.Aliases {
		d.Aliases[alias] = c.Name
	}
}

func (d *DiscommandStruct) addMessageRun(name string, run messageRun) {
	d.messageRuns[name] = run
}

func (d *DiscommandStruct) addChatInputRun(name string, run chatInputRun) {
	d.chatInputRuns[name] = run
}

func (d *DiscommandStruct) MessageRun(command string, s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 더욱 나아진
	d.messageRuns[command](&MsgContext{s, m, args})
}

func (d *DiscommandStruct) ChatInputRun(command string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	d.chatInputRuns[command](&InterContext{s, i})
}

var Discommand *DiscommandStruct = new()
