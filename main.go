package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/xyzyxJP/gce-instance/commands"
)

func main() {
	client := slack.New(os.Getenv("SLACK_TOKEN"))
	slashHandler := NewSlashHandler(client)

	slashHandler.RegisterSubHandlers(
		commands.NewSubHandlerGPT(client),
	)

	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SLACK_SECRET"))
		if err != nil {
			log.Println("failed to verify secrets:", err)
			http.Error(w, "failed to verify secrets", http.StatusInternalServerError)
			return
		}

		r.Body = io.NopCloser(io.TeeReader(r.Body, &verifier))
		slashCmd, err := slack.SlashCommandParse(r)
		if err != nil {
			log.Println("failed to parse slash command:", err)
			http.Error(w, "failed to parse slash command", http.StatusInternalServerError)
			return
		}

		if err := verifier.Ensure(); err != nil {
			log.Println("failed to ensure compares the signature:", err)
			http.Error(w, "failed to ensure compares the signature", http.StatusUnauthorized)
			return
		}

		slashHandler.Handle(w, &slashCmd)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("SLACK_PORT")), nil); err != nil {
		log.Fatal(err)
	}
}
