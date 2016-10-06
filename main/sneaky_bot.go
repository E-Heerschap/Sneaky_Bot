package main;
import (
	"github.com/kingpulse/sneaky_bot/utils"
	"os"
	"bytes"
	"encoding/json"
	"github.com/mmcdole/gofeed"
	"fmt"
	"github.com/bwmarrin/discordgo"

	"io/ioutil"
);

/*
Author: Edwin Heerschap
This file handles the startup of the bot.
 */

var(
	discordPtr *discordgo.Session;
)

func main(){

	//Loading settings.
	sbSettings := Settings{};
	if (!loadSettings(&sbSettings)){
		fmt.Println("Shutting down Sneaky Bot for settings modification...");
		return;
	}

	//Creating new Discord session.
	//TODO Change bot token to be included in the sbSettings.JSON
	fmt.Println("Creating session...");
	var buffer bytes.Buffer;
	buffer.WriteString("Bot ");
	buffer.WriteString(sbSettings.BotToken);
	discord, err := discordgo.New("Bot ");
	discordPtr = discord;

	//Checking for errors creating Discord session
	if(err != nil){
		fmt.Println("Error: Discord Session failed to start... Official error: ", err);
		return;
	}

	//Checking to ensure Discord session was intialized.
	if(discord == nil){
		fmt.Println("Error: Discord Session is nil! Shutting down bot.");
		return;
	}

	//Opening discord connection.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error: Failed opening connection. Official error: ", err)
		return
	}

	fmt.Println("Sneaky Bot has started successfully.")

}

//Handles the loading of settings from Settings.JSON contained in the same directory as the executable.
func loadSettings(settings *Settings)(settingsExist bool){

	//Getting current directory
	dir, err := os.Getwd();
	if(err != nil){
		//TODO Print error to error log with log id.
		fmt.Println("Error: Failed to get current directory.");
	}

	//Checking if settings file exists
	var buffer bytes.Buffer;
	buffer.WriteString(dir);
	buffer.WriteString("\\SBSettings.JSON");
	_, err = os.Stat(buffer.String());

	if( err != nil) {
		//Settings file doesn't exist. Creating a default settings file.
		fmt.Println("Could not find settings... Creating default settings file.")
		file, err := os.Create(buffer.String());

		if(err != nil){
			fmt.Println("Failed to create settings file!");
			return false;
		}

		s := &Settings{
			RssList: []string{""},
			RssRefresh: 60,
			BotToken: "",
		};
		str, _ := json.MarshalIndent(s, "", "	");
		file.Write(str);
		err = file.Close();
		if(err != nil){
			fmt.Println("Failed to close settings file after reading");
		}
		return false;
	}else{
		//Settings file exists. Loading settings.
		file, err := ioutil.ReadFile(buffer.String());
		if(err != nil){
			fmt.Println("Failed to load settings. Shutting down.");
			return false;
		}
		err = json.Unmarshal(file, settings);
		if(err != nil){
			fmt.Println("Failed to read settings. Invalid JSON: ");
			fmt.Println(err);
			return false;
		}

		return true;
	}

}



func printNews(s *discordgo.Session){
	fp := gofeed.NewParser();

	messages, err := s.ChannelMessages("231397628565258240", 100 ,"9223372036854775807", "0");

	if (err != nil){
		fmt.Println("Could not get channel messages", err);
	}
	fmt.Println(len(messages));
	foundFlag := false;


	feed, _ := fp.ParseURL("http://feeds.thescoreesports.com/csgo.rss");
	for i := 0; i < len(feed.Items); i++{
		for n := 0; n < len(messages); n++{
			if(messages[n].Content == feed.Items[i].Link){

				foundFlag = true;
			}
		}
		if(!foundFlag){
			utils.MessageCreate(discordPtr, feed.Items[i].Link, "3");
		}
		foundFlag = false;
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