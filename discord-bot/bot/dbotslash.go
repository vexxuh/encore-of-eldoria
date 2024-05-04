package discordbot

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = getBotParams("GUILD_ID")
	BotToken       = getBotParams("BOT_TOKEN")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	fmt.Println(BotToken)
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMinValue = 1.0
	// dmPermission                   = false
	// defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			// Command: Create - Create player character.
			// Player States: Normal, Combat, Dead, Nil.
			// Requires: character name.
			// Returns: Flavor text.
			Name:        "create",
			Description: "Creates a new character with the given name.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "name",
					Description: "The name of your new character.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			// Command: Inventory - Show the player Inventory including items, stats, weapons, armor, name, xp.
			// Player States: Normal, Combat, Dead.
			// Requires: none.
			// Optional: Action - Equip, Use
			// Optional: Item
			// Returns: Static text.
			Name:        "inventory",
			Description: "Can show your inventory.\nCan be used to equip items and gear.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "action",
					Description: "Equip or Use, Leave empty to inspect items or open inventory.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "equip",
							Value: "equip",
						},
						{
							Name:  "use",
							Value: "use",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "item",
					Description: "The name of an item. Leave empty to unequip gear or open inventory.",
					Required:    false,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "apple",
							Value: "apple",
						},
						{
							Name:  "potion",
							Value: "potion",
						},
						{
							Name:  "potionPlus",
							Value: "potionPlus",
						},
					},
				},
			},
		},
		{
			// Command: Travel - Take the player to the location they specify.
			// Player States: Normal.
			// Requires: A choice from the discord bot's list of places.
			// Returns: Flavor text.
			Name:        "travel",
			Description: "Moves you to your desired location.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "location",
					Description: "The place you want to go.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "plains",
							Value: "plains",
						},
						{
							Name:  "cave",
							Value: "cave",
						},
						{
							Name:  "forest",
							Value: "forest",
						},
						{
							Name:  "town",
							Value: "town",
						},
						{
							Name:  "dungeon",
							Value: "dungeon",
						},
					},
				},
			},
		},
		{
			// Command: Search - Puts player into a fight with a static enemy depending on current location.
			// Player State: Normal.
			// Requires: none.
			// Returns: Flavor text.
			Name:        "search",
			Description: "Searches the area for enemies to fight and loot to find.",
		},
		{
			// Command: Attack: - Enacts one cycle in the current fight if the player is in is one.
			// Player State: Combat.
			// Requires: Attack-type.
			// Returns: flavor text.
			Name:        "attack",
			Description: "Attacks the enemy.",
		},
		{
			// Command: Store - Opens the store if the player is in the town.
			// Optional: Buy and Sell.
			// Optional: Item.
			// Optional: number of items to buy/sell.
			// Player State: Normal.
			// Returns: flavor text.
			Name:        "store",
			Description: "Browses the store when in town.\nCan be used to buy and sell items.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "transaction-type",
					Description: "Buy or sell an item. Leave emtpy to see available items.",
					Type:        discordgo.ApplicationCommandOptionString,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "buy",
							Value: "buy",
						},
						{
							Name:  "sell",
							Value: "sell",
						},
					},
				},
				{
					Name:        "item",
					Description: "Buy or sell specified item. Leave emtpy to see available items.",
					Type:        discordgo.ApplicationCommandOptionString,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "apple",
							Value: "apple",
						},
						{
							Name:  "potion",
							Value: "potion",
						},
						{
							Name:  "potionPlus",
							Value: "potionPlus",
						},
					},
				},
				{
					Name:        "quantity",
					Description: "Ammount of items being bought or sold.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					MinValue:    &integerOptionMinValue,
					MaxValue:    999999,
				},
			},
		},
		{
			// Command: Cheatmode - Adds weapons armor potions and potions+ to the players inv for testing
			// Requires: None.
			// Returns: text.
			// Player State: Normal, Combat.
			Name:        "cheatmode",
			Description: "Gives the player many items for testing purposes.",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"create": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID
			options := i.ApplicationCommandData().Options
			name := options[0].StringValue()

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.CreateCharacter", userId, name)
			fmt.Printf("%+v\n", resp)

			// return and format either error message or creation data
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						`Player successfully created!
						Character Name: %s
						Chracter Level: %d 
						Current Health: %d
						`,
						resp.P.User,
						resp.P.C_level,
						resp.P.C_health,
					),
				},
			})
		},
		"inventory": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID
			options := i.ApplicationCommandData().Options
			eoeMessage := ""
			if len(options) == 2 {
				eoeMessage = options[0].StringValue() + " " + options[1].StringValue()
			} else if len(options) == 1 {
				if options[0].StringValue() == "equip" || options[0].StringValue() == "use" {
					eoeMessage = options[0].StringValue() + " _"
				} else {
					eoeMessage = "_ " + options[0].StringValue()
				}
			}

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Inventory", userId, eoeMessage)
			fmt.Printf("%+v\n", resp)

			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}

			// Determine and format type of response
			var c []string
			if len(options) == 0 {
				c = []string{"You have opened your inventory:",
					"\nName: ", resp.P.Username,
					"\nCurrent State: ", resp.P.P_state,
					"\nCurrent Area: ", resp.P.C_area,
					"\nLevel: ", strconv.Itoa(resp.P.C_level),
					"\nHealth: ", strconv.Itoa(resp.P.C_health+resp.P.B_health) + "/" + strconv.Itoa(resp.P.M_health),
					"\nStrength: ", strconv.Itoa(resp.P.S_strength),
					"\nAgility: ", strconv.Itoa(resp.P.S_agility),
					"\nConstitution: ", strconv.Itoa(resp.P.S_constitution),
					"\nIntelligence: ", strconv.Itoa(resp.P.S_intelligence),
					"\nWisdom: ", strconv.Itoa(resp.P.S_wisdom),
					"\nSword Proficiency Level: ", strconv.Itoa(resp.P.W_s_sword),
					"\nAxe Proficiency Level: ", strconv.Itoa(resp.P.W_s_axe),
					"\nSpear Proficiency Level: ", strconv.Itoa(resp.P.W_s_spear),
					"\nApples: ", strconv.Itoa(resp.P.Inventory.I_apple),
					"\nPotions: ", strconv.Itoa(resp.P.Inventory.I_potion),
					"\nPotions Plus: ", strconv.Itoa(resp.P.Inventory.I_potionPlus),
					"\nGold: ", strconv.Itoa(resp.P.Inventory.C_gold)}
			} else {
				c = []string{resp.TextGen}
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: strings.Join(c, ""),
				},
			})
		},
		"travel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID
			options := i.ApplicationCommandData().Options
			location := options[0].StringValue()

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Travel", userId, location)
			fmt.Printf("%+v\n", resp)

			// return and format either error message or creation data
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: resp.TextGen,
				},
			})
		},
		"search": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Search", userId, "")
			fmt.Printf("%+v\n", resp)

			// return and format either error message or creation data
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: resp.TextGen,
				},
			})
		},
		"attack": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Attack", userId, "")
			fmt.Printf("%+v\n", resp)

			// return and format either error message or creation data
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: resp.TextGen,
				},
			})
		},
		"store": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID
			options := i.ApplicationCommandData().Options
			eoeMessage := ""
			if len(options) == 3 {
				eoeMessage = options[0].StringValue() + " " + options[1].StringValue() + " " + options[2].StringValue()
			} else if len(options) == 2 {
				eoeMessage = options[0].StringValue() + " " + options[1].StringValue() + " " + strconv.Itoa(1)
			} else if len(options) == 1 {
				if options[0].StringValue() == "buy" || options[0].StringValue() == "sell" {
					eoeMessage = options[0].StringValue() + " _" + " _"
				} else {
					eoeMessage = "_ " + "_ " + options[0].StringValue()
				}
			}

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Store", userId, eoeMessage)
			fmt.Printf("%+v\n", resp)

			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}

			// Determine and format type of response
			var c []string
			if len(options) == 0 {
				c = []string{"Here's what we have in stock: ",
					"\nApples | Cost: 1 gold",
					"\nPotions | Cost: 5 gold",
					"\nPotions Plus | Cost: 10 gold"}
			} else {
				c = []string{resp.TextGen}
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: strings.Join(c, ""),
				},
			})
		},
		"cheatmode": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Grab user input from bot interaction
			userId := i.Member.User.ID

			// Create encore API request + error handling
			resp, err := backendGameProcessorRequest("eoe_api.Cheatmode", userId, "")
			fmt.Printf("%+v\n", resp)

			// return and format either error message or creation data
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There has been an error with your request",
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: resp.TextGen,
				},
			})
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func RunBotSlash() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
