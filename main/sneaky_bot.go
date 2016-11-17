package main;
import (
	"os"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"

	"io/ioutil"
	"github.com/kingpulse/sneaky_bot/commands"
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
	fmt.Println("Creating session...");
	var buffer bytes.Buffer;
	buffer.WriteString("Bot ");
	buffer.WriteString(sbSettings.BotToken);
	discord, err := discordgo.New(buffer.String());
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

	c := commands.NewManager(discordPtr)

	//Setting up discord message listener
	discord.AddHandler(c.OnCommandCall)

	//Opening discord connection.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error: Failed opening connection. Official error: ", err)
		return
	}


	fmt.Println("Sneaky Bot has started successfully.")

	<-make(chan struct{})
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



