package scrape

import (
	"fmt"
	"log"
	"testing"
)

func TestScrape(t *testing.T) {
	res, err := Scrape("こんにちは")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// &{{
//     "action": "next",
//     "messages": [
//         {
//             "id": "16bacb47-7f98-4c26-b29e-b71de80cc600",
//             "role": "user",
//             "content": {
//                 "content_type": "text",
//                 "parts": [
//                     "こんばんは"
//                 ]
//             }
//         }
//     ],
//     "conversation_id": "df2b6d9f-fcf7-40c8-bef8-a3cb6e718e7e",
//     "parent_message_id": "37122ae7-5cd6-4c73-b865-76be05b40056",
//     "model": "text-davinci-002-render"
// } 0 -1}