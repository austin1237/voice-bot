package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord session using the bot token.
	token, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		fmt.Println("Error: DISCORD_BOT_TOKEN environment variable not set.")
		return
	}
	
	dg, err := discordgo.New("Bot " + token)
	
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Add a message handler.
	dg.AddHandler(messageCreate)

	// Open a connection to Discord.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection: ", err)
		return
	}

	// Wait here until CTRL-C or other signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping", respond with "Pong!".
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
