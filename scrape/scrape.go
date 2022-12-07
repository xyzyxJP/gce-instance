package scrape

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	fmt.Println(strings.NewReader(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("authority", "chat.openai.com")
	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("accept-language", "ja")
	req.Header.Add("authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJzYWhhc2hpMjAwMkBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZ2VvaXBfY291bnRyeSI6IlVTIn0sImh0dHBzOi8vYXBpLm9wZW5haS5jb20vYXV0aCI6eyJ1c2VyX2lkIjoidXNlci1vN0N1Q2ZGcmxrSDNqTGlFb0hQdnVWbkEifSwiaXNzIjoiaHR0cHM6Ly9hdXRoMC5vcGVuYWkuY29tLyIsInN1YiI6Imdvb2dsZS1vYXV0aDJ8MTAzNjUzNzU3Njg1NDMwNTY5MDU4IiwiYXVkIjpbImh0dHBzOi8vYXBpLm9wZW5haS5jb20vdjEiLCJodHRwczovL29wZW5haS5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjcwMzc5ODI5LCJleHAiOjE2NzA0NjYyMjksImF6cCI6IlRkSkljYmUxNldvVEh0Tjk1bnl5d2g1RTR5T282SXRHIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSBtb2RlbC5yZWFkIG1vZGVsLnJlcXVlc3Qgb3JnYW5pemF0aW9uLnJlYWQgb2ZmbGluZV9hY2Nlc3MifQ.hXomVp4edytQjcnA0ktSN9cHqzgXhe5vGmAA_4YKimNzykXXjWaYqb9Ako0e1kHSVfOGKQzYAK6WQwz3hNKpfHfVW2Cd4EHi-2Rnb9s3d_yZUK1kHVeRxRe32wgZFLPjJ2-WsvK1JfpskjVk6OVgQ4dZ1n2_R7Jv5jnLbf8bN7a1iJegIRxgAmXfQ3uzCSvOo-63SOGTst18TLxOFPAez-1HX25P8WiGg-LOLg3PWBAZmTpl9bwhH7W5fb-KiTatNs7IStF-dTLxAsnHto_dUD6z3sW2yt9A0ZxmZBcxGuafqgOX5KIMtpu90MeAFEV_zVLImEkuGoY9HqkWSHKwbQ")
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
	req.Header.Add("Cookie", "__Host-next-auth.csrf-token=88159c616daa99599db842026b1c90c016f3f3fe9a93ba73f3c9139bf94eb231%7C6123aa998e5cd3d15c97dc40309731a245d5e2e5e6a7a0b1d333918c7d3fd602; __Secure-next-auth.callback-url=https%3A%2F%2Fchat.openai.com; __Secure-next-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..F_nuh2JeWgTn9iZV.sjxmhK2NC9IhHLl9uWq081u96l5Phjki_lA8s6uwbbAr6IOPriYLAtmBvZDGtpJOe8mpp1wiRMdjUaasDvLy6Og4ytKUqgKNNNmm67Vlti05RVD6KsMwi4-allmG2rw2GvFAB_bC7KuLFbGGF5kv8QFNBMBf199Pk4rHgjQIciiLU7IpYv72lC3DjN2n21d8uQA3FoHPgfqtD2GCrws7XNvhFJKzQQF2DIIwfgY0gyAe6vjRNhjFlMXg1_fOMmAP3J7BR9kx6Mz0RsckmbnAuDGu7vWcPYKCvav8nz2O1kzZmmIaPF44EEdjW368rhAf9CBWCkLn_UzPGQQc3Zzwd_VjuHkSzliRraVJE0Xi-5QQgkiMGo2ZrKVZNv7J0Mb-jn3ggBGli_SyL3CFIRD_b_IcrGQ1K7nLzUX6DYpBJpiijawvc6EeQhGrlwwhh2fZLWbsl_vbcr5fZpaZPBeDSxhrqWE2jb9m37eezoa3mJq5QbOBvhTOIaXLjiOopcs9qPNR6haqMli5rANPP5nqmqCBg2xGgq1wLY_XDOZQIcTJlrSSYZj-RnWuwxrqK16HXRsBlMrzQtDaVsy8KXtpGLEumYKO63e3XZSLRKufkDJbNkDvCJGB4L4r93jFhqZLUxOm27WgsMule3hX3eAT6WEeOo5TvuNQr18r_KgdqphhSjK9K_sr3nYhjHJW_H6s3McBDYVbenzwSboR7MSkQnfEUNfxaZhj8GlqWwC6SbA5X_4y6zz8uLSYIF6gqUyogzhYn2MJ3Q6jZSumTWyfB6Lkzr2N4mn4ik9KMM8PxbVN1P1KcBz1Zhsew1kfIJi5lmHlD5YEr_gmOFobt1mFQXGOuj7erBo-QEVKTNIxJyDBwJijrL8zvCOs_d4anXeAZlN7hFx90dhPqwaps7TjKrJDVvedBzPsi6TY7WtoV67XoYpaCwPZyh3kXbJUZ3iJVc7lu4iyj5om1QPCm-fynQht5yYkhlKH9G0H6hlgznQfudo7VIfDkixfY82a6Pt3lFiIhMpKvlwDoEZfyUlVFE4oeSJvbPfPpfeSCDxWTslT89OFQhwWN_IKePXrHtHm4KPqtGQGNY6s3cPHcImghUd7-k_qk0gXYk14TgQMfGaxeBbxxEVsQgULueC0b_yhU8hj6T6IgP5iTWfVBIJc-OzW2lJIGtbWN5Wuw6Ca5efvCsh04Bq8V1ygdmGCVTzy38LSWQfFbTNu8MbF4aaRtNHzICLEoqLZwznrLlX5eFGh5fqcmlr6II2hIFcQfvwcBDNm5gQU4zD9mqDc6bUw_vU_ztZHarVmgFryYWbHcnND9qpq8louqXSAogSiHgVjpXAWV61wQBAjRQO3IJzHEI6QvV8n70u2wMtTs158YLuh_B9Mci7ObTn68J4vdIYlJxSC6GvrkniKOh3uvqrwT3wzf_pmknRwj9zZijd5ExmZ4iRPN9lUwFEe7ZR6nWXxVtOL7Ypa6x-5svx4kHMInTLA5JbB5CoqH5BoDA8DpRCV9JN2BWLNiJmEnw-IeG1Cm2wnR3xqCs2bSD3WSIxw4GceYlXNEZ8MJ4gtfGhq3HwzS4_zZen1VepQ48nyj2wZUrxnWHVKrlYLJLYbSscNSW-aiJD9hz0j13kfZpn3iqTXJI-2n6h59m7lBd5SVv5oJTC40B-9mXJlDZ_D-RXhFxnvaMxPcmFRnGrx2sieRjPZ-qJL1XSi4FYhAEjJwFKRs0VtxUBK4Gw-towFFL8DLqOCInRjnLmtT62a-WtL7gOA5Eqzk4-76qLcda9drZtnBW9ATXumqLF1zXT256j6qRMBuGHw7fFAM7PoUNuKjWiBw4A7-AhNx2uuIPlyocdc-5OD865XE0WEpCgdv6VkqrgqOCMT-mhSuNoTmB5wDSUrbMgdHs0zVdFfCfQfF37UAdAxpgW6_rTyNRqK04nV43c4YC9DK9skd-KbPt-rAO7Z8wL4BRp4cqFG86h9IQcYsVOgbvArNBRxj9sUgiT9AHlHJP0Gxo6k2Fad-ZfdSBzcLFUj_fcEakj-8zJcBLC3XSVBqC8hin1VjBfLsuUEoj1Ens6QBEoBmrSITQFDCuO8kWGwq9YYmCkp9btLE9Nd6Cg_ZpDXlI3GFa6RPndYFjk--1Sw8fla4goV64Vx05GE1TkL5KVveugRVk5xd-7TVzaKT7T6KGtdAS0XUxpopPFNS39DwNlHNBToE37SMddVTwr5cfYLioHIQq2cNWkYZJg2yLLHMVc0Luua6JYb.IfLRFMsCmGUyJJM3YsqYpw")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
