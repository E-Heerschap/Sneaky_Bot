package rss

import (
	"github.com/kingpulse/sneaky_bot/utils"
	"github.com/mmcdole/gofeed"
	"github.com/bwmarrin/discordgo"
)

type RssChannel struct{
	feeds []string
	postedMessages []string
}

var(
	rssChannels map[string]RssChannel;
)


func addPostedMessage(rssChannel *RssChannel, message string){

}

func checkRssFeeds(s *discordgo.Session, channelID string){
	fp := gofeed.NewParser();


	foundFlag := false;


	feed, _ := fp.ParseURL("http://feeds.thescoreesports.com/csgo.rss");

	var rssChannel *RssChannel;
	rssChannel = &RssChannels[channelID];

	for i := 0; i < len(feed.Items); i++{
		for n:= 0; n < len(*rssChannel.postedMessages); n++{
			if(*rssChannel.postedMessages[n] == feed.Items[i].Link){
				foundFlag = true;
			}
		}
		if(!foundFlag){
			utils.MessageCreate(s, feed.Items[i], channelID);
		}
	}

	feed, _ = fp.ParseURL("http://feeds.thescoreesports.com/lol.rss");
	for i := 0; i < len(feed.Items); i++{
		for n := 0; n < len(messages); n++{
			if(messages[n].Content == feed.Items[i].Link){

				foundFlag = true;
			}
		}
		if(!foundFlag){
			utils.MessageCreate(discordPtr, feed.Items[i].Link,"3");
		}
		foundFlag = false;
	}
}
