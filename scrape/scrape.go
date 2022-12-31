package scrape

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/xyzyxJP/gce-instance/model"
)

func Auth() {
	url := "https://chat.openai.com/api/auth/session"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("cookie", "__Secure-next-auth.session-token="+os.Getenv("OPENAI_TOKEN"))
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Models() {
	url := "https://chat.openai.com/backend-api/models"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authorization", "Bearer "+os.Getenv("OPENAI_TOKEN"))
	req.Header.Add("cookie", "__Secure-next-auth.session-token="+os.Getenv("OPENAI_TOKEN"))
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Scrape(query string) (string, error) {
	url := "https://chat.openai.com/backend-api/conversation"
	method := "POST"

	// message_id := uuid.New()
	// conversation_id := uuid.New()
	// parent_message_id := uuid.New()
	message_id := "16bacb47-7f98-4c26-b29e-b71de80cc600"
	conversation_id := "df2b6d9f-fcf7-40c8-bef8-a3cb6e718e7e"
	parent_message_id := "37122ae7-5cd6-4c73-b865-76be05b40056"

	payload := fmt.Sprintf(`{
		"action": "next",
		"messages": [
			{
				"id": "%s",
				"role": "user",
				"content": {
					"content_type": "text",
					"parts": [
						"%s"
					]
				}
			}
		],
		"conversation_id": "%s",
		"parent_message_id": "%s",
		"model": "text-davinci-002-render"
	}`, message_id, query, conversation_id, parent_message_id)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		return "", err
	}
	req.Header.Add("authority", "chat.openai.com")
	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("accept-language", "ja")
	req.Header.Add("authorization", "Bearer "+os.Getenv("GPT_TOKEN"))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("dnt", "1")
	req.Header.Add("origin", "https://chat.openai.com")
	req.Header.Add("referer", "https://chat.openai.com/chat")
	req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	rows := strings.Split(string(body), "\n")
	lastRow := strings.Replace(rows[len(rows)-5], "data: ", "", 1)

	var message model.Response
	json.Unmarshal([]byte(lastRow), &message)

	return strings.Join(message.Message.Content.Parts, ""), nil
}
