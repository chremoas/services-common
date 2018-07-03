package args

import (
	"bytes"
	"fmt"
	proto "github.com/chremoas/chremoas/proto"
	discordsrv "github.com/chremoas/discord-gateway/proto"
	"github.com/chremoas/services-common/discord"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type Args struct {
	argMap  map[string]*Command
	argList []string

	serviceName    string
	serviceType    string
	serviceVersion string

	userId    string
	channelId string
}

type Command struct {
	Funcptr func(ctx context.Context, request *proto.ExecRequest) string
	Help    string
}

func NewArg(serviceName, serviceType, serviceVersion string, discordClient discordsrv.DiscordGatewayService) *Args {
	a := &Args{}
	a.argMap = make(map[string]*Command)

	a.serviceName = serviceName
	a.serviceType = serviceType
	a.serviceVersion = serviceVersion

	return a
}

func (a *Args) Add(name string, command *Command) {
	a.argList = append(a.argList, name)
	a.argMap[name] = command
}

func (a Args) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) (embed *discordsrv.SendMessageEmbed, err error) {
	var response string

	s := strings.Split(req.Sender, ":")
	a.userId = s[1]
	a.channelId = s[0]

	if len(req.Args) == 1 || req.Args[1] == "help" {
		response = a.help()
		embed = &discordsrv.SendMessageEmbed{
			ChannelID: a.channelId,
			Message:   a.embedHelp(),
		}
	} else {
		f, ok := a.argMap[req.Args[1]]
		if ok {
			response = f.Funcptr(ctx, req)
		} else {
			return nil, fmt.Errorf("not a valid subcommand: %s", req.Args[1])
		}
	}

	rsp.Result = []byte(response)
	return embed, nil
}

func (a Args) help() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Usage: !%s <subcommand> <arguments>\n", a.serviceName))
	buffer.WriteString("\nSubcommands:\n")

	for cmd := range a.argList {
		if a.argMap[a.argList[cmd]].Help != "" {
			buffer.WriteString(fmt.Sprintf("\t%s: %s\n",
				a.argList[cmd],
				a.argMap[a.argList[cmd]].Help,
			))
		}
	}

	return fmt.Sprintf("```%s```", buffer.String())
}

func (a Args) embedHelp() *discordsrv.MessageEmbed {
	var buffer bytes.Buffer
	s := fmt.Sprintf("%s-%s", a.serviceName, a.serviceType)

	for cmd := range a.argList {
		if a.argMap[a.argList[cmd]].Help != "" {
			buffer.WriteString(fmt.Sprintf("\t%s: %s\n",
				a.argList[cmd],
				a.argMap[a.argList[cmd]].Help,
			))
		}
	}

	return discord.NewEmbed().
		SetTitle("Do we need a title?").
		SetDescription("Do we need a description?").
		SetAuthor(&discordsrv.MessageEmbedAuthor{
			Name: fmt.Sprintf("Usage: !%s <subcommand> <arguments>", a.serviceName),
			URL:  fmt.Sprintf("https://github.com/chremoas/%s/tree/%s", s, a.serviceVersion),
		}).
		AddField("Subcommands", buffer.String()).
		SetFooter(&discordsrv.MessageEmbedFooter{
			Text:    fmt.Sprintf("Chremoas Chat Bot | %s (%s)", s, a.serviceVersion),
			IconURL: "https://avatars3.githubusercontent.com/u/33756515?s=400&u=af0c82e2ed951031a4c574f0e93a8b1db2598bb6&v=4",
		}).
		SetTimestamp(time.Now().UTC().String()).
		SetColor(0x00ff00).MessageEmbed
}
