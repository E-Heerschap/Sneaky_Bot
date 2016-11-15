package commands

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
	"github.com/kingpulse/sneaky_bot/utils"
)

/*
Author: Edwin Heerschap
This handles the commands given to sneaky_bot
 */

var(
	commandsMap map[string]func(params []string, m discordgo.MessageCreate)
	BotID string
)

func initializeCommands(dg *discordgo.Session){
	commandsMap["ping"] = ping
	commandsMap["dice"] = dice
	commandsMap["addrss"] = addRss
	var err string
	BotID, err = dg.User("@me")

	if err != nil {
		fmt.Print("Error: failed to get Bot ID. Official error: ", err)
	}
}

func OnCommandCall(discordSession *discordgo.Session, messageCreate *discordgo.MessageCreate){

	//Not considering if its the bots message.
	if messageCreate.Author.ID == BotID {
		return
	}

	if strings.HasPrefix(messageCreate.Content, "//") {

		//Formatting string for use
		msg := strings.ToLower(messageCreate.Content)

		tokens := strings.Split(msg[2:], " ")

		//This calls the appropriate command if one is found.
		cmd, ok := commandsMap[tokens[0]]
		if ok {
			go cmd(tokens)
		}else {
			go utils.MessageCreate(discordSession, "Sorry that command doesn't exist.", messageCreate.ChannelID)
		}

	}

}

/*
--------------------------------
Functions for commands are below
---------------------------------
 */

//Ping sends a message back displaying the ping to the bot
func ping(params []string, m discordgo.MessageCreate) {

}

//Dice allows the user to roll a virtual dice of a given size.
func dice(params []string, m discordgo.MessageCreate){

}

//addRss allows the guild to add a rss feed to their rss notification channel
func addRss(params []string, m discordgo.MessageCreate){

}

//removeRss allows the guild to remove a rss feed from their rss notification channel
func removeRss(params []string, m discordgo.MessageCreate){

}



