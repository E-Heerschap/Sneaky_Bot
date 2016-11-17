package commands

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
	"github.com/kingpulse/sneaky_bot/utils"
	"time"
	"bytes"
	"strconv"
	"math/rand"
)

/*
Author: Edwin Heerschap
This handles the commands given to sneaky_bot
 */

type CommandsManager struct {
	commandsMap map[string]func(params []string, m discordgo.MessageCreate)
	BotID       string
	discordPtr *discordgo.Session
	ball8Answers [20]string
}

func NewManager(dg *discordgo.Session) (c CommandsManager){

	c.discordPtr = dg

	//Setting up command map
	c.commandsMap = make(map[string]func(params []string, m discordgo.MessageCreate))
	c.commandsMap["ping"] = c.ping
	c.commandsMap["dice"] = c.dice
	c.commandsMap["addrss"] = addRss
	c.commandsMap["removerss"] = removeRss
	c.commandsMap["8ball"] = c.ball8

	bot, err := dg.User("@me")


	c.BotID = bot.ID

	if err != nil {
		fmt.Print("Error: failed to get Bot ID. Official error: ", err)
	}

	//List of 8 ball answers found from: https://en.wikipedia.org/wiki/Magic_8-Ball
	ball8Answers := [20]string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes, definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",


	}

	c.ball8Answers = ball8Answers

	return c
}

func (c *CommandsManager) OnCommandCall(discordSession *discordgo.Session, messageCreate *discordgo.MessageCreate){

	//Not considering if its the bots message.
	if messageCreate.Author.ID == c.BotID {
		return
	}

	if strings.HasPrefix(messageCreate.Content, "//") {

		//Formatting string for use
		msg := strings.ToLower(messageCreate.Content)

		tokens := strings.Split(msg[2:], " ")

		//This calls the appropriate command if one is found.
		cmd, ok := c.commandsMap[tokens[0]]
		if ok {
			go cmd(tokens, *messageCreate)
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
func (c * CommandsManager) ping(params []string, m discordgo.MessageCreate) {

	t, err := time.Parse(time.RFC3339, m.Timestamp)
	if err != nil {
		fmt.Println("Error: Failed to parse time from discord ping request.")
	}

	ping := time.Since(t)

	str := bytes.NewBufferString("``Ping to bot: ")
	str.WriteString(strconv.FormatFloat(ping.Seconds() * 100, 'f', 0, 64))
	str.WriteString("ms``")

	go utils.MessageCreate(c.discordPtr, str.String(), m.ChannelID)
}

//Dice allows the user to roll a virtual dice of a given size.
func (c *CommandsManager) dice(params []string, m discordgo.MessageCreate){

	if len(params) < 2 {
		//User has not entered a number for the dice roll.
		go utils.MessageCreate(c.discordPtr, "Enter a number for the roll. Example: //Dice 6", m.ChannelID)
	}

	//Generating random number based on given number
	diceNo, err := strconv.Atoi(params[1])
	if err != nil && diceNo > 0 {
		go utils.MessageCreate(c.discordPtr, "Sorry you did not enter a valid number.", m.ChannelID)
		return
	}
	num := rand.Intn(diceNo)

	//Sending message back to client
	str := bytes.NewBufferString("``You rolled a: ")
	str.WriteString(strconv.Itoa(num))
	str.WriteString("!``")
	go utils.MessageCreate(c.discordPtr, str.String(), m.ChannelID)
}


//addRss allows the guild to add a rss feed to their rss notification channel
func addRss(params []string, m discordgo.MessageCreate){

}

//removeRss allows the guild to remove a rss feed from their rss notification channel
func removeRss(params []string, m discordgo.MessageCreate){

}

//ball8 Sends an answer back to the sender with an answer from an 8 ball.
//The replied answer has nothing to do with the passed question (params[1:]) but it
//but the question is required.
func (c *CommandsManager) ball8(params []string, m discordgo.MessageCreate){
	if len(params) < 2 {
		go utils.MessageCreate(c.discordPtr, "``You need to ask a question as well!``", m.ChannelID)
	}else{
		message := bytes.NewBufferString("``")
		message.WriteString(c.ball8Answers[rand.Intn(19)])
		message.WriteString("``")

		go utils.MessageCreate(c.discordPtr, message.String(), m.ChannelID)
	}
}


