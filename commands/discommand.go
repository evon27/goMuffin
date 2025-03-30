package commands

import (
	// "fmt"

	"sync"

	"github.com/bwmarrin/discordgo"
)

type messageRun func(ctx *MsgContext)
type chatInputRun func(ctx *InterContext)

type Category string

type DetailedDescription struct {
	Usage    string
	Examples []string
}

type Command struct {
	*discordgo.ApplicationCommand
	Aliases             []string
	DetailedDescription *DetailedDescription
	Category            Category
	MessageRun          messageRun
	ChatInputRun        chatInputRun
}

type DiscommandStruct struct {
	Commands map[string]*Command
	Aliases  map[string]string
}

type MsgContext struct {
	Session *discordgo.Session
	Msg     *discordgo.MessageCreate
	Args    []string
	Command *Command
}

type InterContext struct {
	Session *discordgo.Session
	Inter   *discordgo.InteractionCreate
	Command *Command
}

const (
	Chattings Category = "채팅"
	Generals  Category = "일반"
)

var mutex *sync.Mutex = &sync.Mutex{}

func new() *DiscommandStruct {
	discommand := DiscommandStruct{
		Commands: map[string]*Command{},
		Aliases:  map[string]string{},
	}
	return &discommand
}

func (d *DiscommandStruct) LoadCommand(c *Command) {
	mutex.Lock()
	d.Commands[c.Name] = c
	d.Aliases[c.Name] = c.Name

	for _, alias := range c.Aliases {
		d.Aliases[alias] = c.Name
	}
	mutex.Unlock()
}

// func (d *DiscommandStruct) addMessageRun(name string, run messageRun) {
// 	d.messageRuns[name] = run
// }

// func (d *DiscommandStruct) addChatInputRun(name string, run chatInputRun) {
// 	d.chatInputRuns[name] = run
// }

func (d *DiscommandStruct) MessageRun(name string, s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	command := d.Commands[name]
	if command == nil {
		return
	}
	command.MessageRun(&MsgContext{s, m, args, command})
}

func (d *DiscommandStruct) ChatInputRun(name string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := d.Commands[name]
	if command == nil {
		return
	}
	command.ChatInputRun(&InterContext{s, i, command})
}

var Discommand *DiscommandStruct = new()
