package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/{{cookiecutter.githubname}}/{{cookiecutter.botname}}/util"
)

type PingCommand struct {
}

func (command PingCommand) Init() {

}

func (command PingCommand) Match(token string, isUser bool) bool {
	return util.SimplePublicCommand("ping", token, isUser)
}

func (command PingCommand) Run(Client *discordgo.Session, Message *discordgo.MessageCreate) {
	Client.ChannelMessageSend(Message.ChannelID, "Pong!")
}
