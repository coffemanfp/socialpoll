package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type tweet struct {
	Text string
}

func readFromTwitter(votes chan<- string) {
	options, err := loadOptions()
	if err != nil {
		log.Println("failed to load options: ", err)
		return
	}

	u, err := url.Parse("https://stream.twitter.com/1.1/statuses/filter.json")
	if err != nil {
		log.Println("creating filter request failed: ", err)
		return
	}

	query := make(url.Values)
	query.Set("track", strings.Join(options, ","))

	req, err := http.NewRequest(
		"POST",
		u.String(),
		strings.NewReader(
			query.Encode(),
		),
	)
	if err != nil {
		log.Println("creating filter request failed: ", err)
		return
	}

	resp, err := makeRequest(req, query)
	if err != nil {
		log.Println("making request failed:", err)
		return
	}
	reader := resp.Body
	decoder := json.NewDecoder(reader)

	for {
		var tweet tweet
		if err := decoder.Decode(&tweet); err != nil {
			break
		}

		for _, option := range options {
			contains := strings.Contains(
				strings.ToLower(tweet.Text),
				strings.ToLower(option),
			)

			if contains {
				log.Println("vote:", option)
				votes <- option
			}
		}
	}
}

func startTwitterStream(stopchan <-chan struct{}, votes chan<- string) <-chan struct{} {
	stoppedChan := make(chan struct{}, 1)
	go func() {
		defer func() {
			stoppedChan <- struct{}{}
		}()

		for {
			select {
			case <-stopchan:
				log.Println("stopping Twitter...")
				return
			default:
				log.Println("Querying Twitter...")
				readFromTwitter(votes)
				log.Println(" (waiting)")
				time.Sleep(10 * time.Second) // wait before reconnecting
			}
		}
	}()

	return stoppedChan
}

var (
	authSetupOnce sync.Once
	httpClient    *http.Client
)

func makeRequest(req *http.Request, params url.Values) (res *http.Response, err error) {
	authSetupOnce.Do(func() {
		authSetupOnce.Do(func() {
			setupTwitterAuth()
			httpClient = &http.Client{
				Transport: &http.Transport{
					Dial: dial,
				},
			}
		})
	})

	formEnc := params.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))
	req.Header.Set(
		"Authorization",
		authClient.AuthorizationHeader(creds, "POST", req.URL, params),
	)

	res, err = httpClient.Do(req)
	return
}
