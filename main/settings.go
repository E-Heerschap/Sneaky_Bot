package main

type Settings struct {
	//Token of bot to login
	BotToken string `json:"Bot Token"`
	RssRefresh int `json:"RSS Timer"`
}