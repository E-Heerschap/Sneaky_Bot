package main

type Settings struct {
	BotToken string `json:"Bot Token"`
	RssList []string `json:"RSS Feeds"`
	RssRefresh int `json:"RSS Timer"`
}
