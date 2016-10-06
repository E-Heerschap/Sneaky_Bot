package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//Sends the passed message to the specified channel.
func MessageCreate(s *discordgo.Session, message string, channelId string) {
	d, err := s.ChannelMessageSend(channelId, message);
	if(d == nil){
		fmt.Println("Error printing message!");
	}
	if(err != nil){
		fmt.Println("Error printing message!");
	}
}
