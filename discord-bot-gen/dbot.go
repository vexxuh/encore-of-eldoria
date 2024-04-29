// I apologize for the over commentting. I am prefering to explain the code here for learning purposes.

package discordbotgen

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func RunBot(token string) {
	bot, err := discordgo.New("Bot " + token) //HTTP Header Creation
	fmt.Println("Generating Bot with token: " + token)

	if err != nil {
		log.Fatal("error creating bot with token: " + token)
	}

	//EventHandlers for the Bot to listen to
	bot.AddHandler(newMessage)

	err = bot.Open()

	if err != nil {
		log.Fatal("error opening bot with token: " + token)
	}

	defer bot.Close() // Defer runs this line when the code is shutdown, via error or otherwise

	//Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1) //New Go Channel of type os.Signal with buffer length of 1

	// Spins up a thread that waits for a ^C or any type of user input that interrupts
	// puts that interrupt input intp the variable "c"
	signal.Notify(c, os.Interrupt)
	<-c //Hold program here until something is in the "c" variable
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// return if message found is from our own bot
	if message.Author.ID == discord.State.User.ID {
		fmt.Println("not responding to myself")
		return
	}

	fmt.Println("Looking for response for message: " + message.Content + "In ChannelID: " + message.ChannelID)

	// respond to messages
	switch {
	case strings.Contains(message.Content, "hi"):
		discord.ChannelMessageSend(message.ChannelID, "Hello!")
		fmt.Println("wow")
	case strings.Contains(message.Content, "travel"):
		discord.ChannelMessageSend(message.ChannelID, "Where would you like to go? \nPlains\nCave\nTown\nDungeon")
	case strings.Contains(message.Content, "location"):
		output := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeImage,
				Title:       "Current Location",
				Description: "Location for " + message.Author.Username,
				Image: &discordgo.MessageEmbedImage{
					URL:    "https://c8.alamy.com/comp/2MFKMD0/grass-hill-plains-field-or-pasture-with-deep-blue-sky-and-clouds-generic-plain-minimalist-background-image-land-and-sky-spring-summer-uk-2MFKMD0.jpg",
					Width:  633,
					Height: 465,
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Description",
						Value:  "You are in a grassy field, sometimes containing slimes and wolves",
						Inline: false,
					},
				},
			}},
		}
		discord.ChannelMessageSendComplex(message.ChannelID, output)
	case strings.Contains(message.Content, "random image"):
		output := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeImage,
				Title:       "A Random Image",
				Description: message.Author.Username + ", here is a random image for you.",
				Image: &discordgo.MessageEmbedImage{
					URL:    "https://source.unsplash.com/random/200x200?sig=" + strconv.Itoa(rand.Intn(1000000000)),
					Width:  600,
					Height: 600,
				},
			}},
		}
		discord.ChannelMessageSendComplex(message.ChannelID, output)
	default:
		fmt.Println("did not enter swtich case")
	}
}

/*
	message.ChannelID is the id of the channel that the message was sent in
*/
