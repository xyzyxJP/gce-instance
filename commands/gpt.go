package commands

import (
	"strings"

	"github.com/slack-go/slack"
	"github.com/xyzyxJP/gce-instance/scrape"
)

type SubHandlerGPT struct {
	c *slack.Client
}

func NewSubHandlerGPT(c *slack.Client) *SubHandlerGPT {
	return &SubHandlerGPT{c}
}

func (o *SubHandlerGPT) Name() string {
	return "gpt"
}

func (o *SubHandlerGPT) Handle(slashCmd *slack.SlashCommand) error {
	q := strings.Join(strings.Fields(slashCmd.Text)[1:], "")

	res, err := scrape.Scrape(q)
	if err != nil {
		return err
	}

	var blocks []slack.Block

	blocks = append(blocks, slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", res, false, false), nil, nil),
	)

	_, _, _, err = o.c.SendMessage(slashCmd.ChannelID, slack.MsgOptionBlocks(blocks[:]...), slack.MsgOptionDisableLinkUnfurl(), slack.MsgOptionDisableMediaUnfurl())

	return err
}
