package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
)

type CommandExecutor interface {
	Handle(slashCmd *slack.SlashCommand) error
	Name() string
}

type SlashHandler struct {
	client    *slack.Client
	executors []CommandExecutor
}

func NewSlashHandler(client *slack.Client) *SlashHandler {
	return &SlashHandler{client, make([]CommandExecutor, 0, 100)}
}

func (s *SlashHandler) RegisterSubHandlers(executors ...CommandExecutor) {
	s.executors = append(s.executors, executors...)
}

func (s *SlashHandler) Handle(rw http.ResponseWriter, slashCmd *slack.SlashCommand) {
	tokens := strings.Fields(slashCmd.Text)
	// TODO: show help if the length of tokens is 0

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	// 打ったコマンドを表示させる
	msg := &slack.Msg{ResponseType: slack.ResponseTypeInChannel}
	_ = json.NewEncoder(rw).Encode(msg)

	go func() {
		subcommand := tokens[0]
		for _, executor := range s.executors {
			if executor.Name() == subcommand {
				if err := executor.Handle(slashCmd); err != nil {
					log.Println(err)
					return
				}
				break
			}
		}
	}()
}
