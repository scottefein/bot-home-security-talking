package main

import (
	"botexample/client"
	"botexample/commands"
	"botexample/conf"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"os/exec"
)

// Globally accessible variables
var currentUserId int64
var currentHost string

var whatTimeIsItRegex = regexp.MustCompile(`(?i)what time is it`)
var botathonRegex = regexp.MustCompile(`(?i)#botathon`)

func main() {
	currentHost, _ = os.Hostname()
	fmt.Println("hello " + currentHost)

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("Cannot start Woodhouse without telling him which environment you'd like to use. Your command line needs to look like \"go run Hello.go resources/nexus.json\"")
	}
	configurationLoader := conf.ConfigurationLoader{ConfigurationFileName: argsWithoutProg[0]}
	config := configurationLoader.Load(configurationLoader.ConfigurationFileName)
	botClient := client.BotClient{
		AgentUrl:          config.AgentUrl,
		SessionAuthUrl:    config.SessionAuthUrl,
		KeyManagerAuthUrl: config.KeyManagerAuthUrl,
		PodUrl:            config.PodUrl,
		CertFilePath:      config.CertFilePath,
		KeyFilePath:       config.KeyFilePath,
	}

	sum := 1
	for sum < 1000 {
		ret := botClient.Authenticate()

		if ret > 0 {
			fmt.Printf("Could not auth, will retry in 2.5 seconds...")
			time.Sleep(2500 * time.Millisecond)
		} else {
			break
		}
	}

	currentUserId = botClient.GetCurrentUserId()

	messageHandlers := registerMessageHandlers(botClient, config)

	var buffer bytes.Buffer
	buffer.WriteString("<messageML>")
	buffer.WriteString("I was just started, which seems I was down for a while")
	buffer.WriteString("</messageML>")
	botClient.SendMessageMLMessage(config.StreamId, buffer.String())

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Symphony Home Security Bot")

	go func() {
		defer wg.Done()

		channel := botClient.StartStreaming(wg)
		for {
			message := <-channel
			switch message := message.(type) {
			case client.V2Message:
				if message.FromUserId == currentUserId {
					break
				}

				var found int

				found = 0
				//			replyIfScotch(message.Message, message.StreamId, botClient)
				//			replyIfBotathon(message.Message, message.StreamId, botClient)
				for _, commandHandler := range messageHandlers {
					loc := commandHandler.MatchRegex().FindStringIndex(message.Message)
					if loc != nil && loc[0] == 11 {
						// first location after messageML tag
						commandHandler.OnMessage(message, botClient, messageHandlers)
						found = 1
					}
				}

				if found == 0 {
					realMessage := message.Message[11 : len(message.Message)-12]
					commandToRun := "festival --tts < \"" + realMessage + "\""

					fmt.Printf(commandToRun)

					cmd := exec.Command("festival", "--tts")
					cmd.Stdin = strings.NewReader(realMessage)
					var out bytes.Buffer
					cmd.Stdout = &out
					err := cmd.Run()
					if err != nil {
						log.Fatal(err)
					}
				}
				break
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			if len(strings.TrimSpace(text)) > 0 {
				processInput(text, botClient, config)
			}
		}
	}()

	wg.Wait()
}

func registerMessageHandlers(botclient client.BotClient, config conf.Configuration) []commands.CommandHandler {
	messageHandlers := make([]commands.CommandHandler, 0)

	messageHandlers = append(messageHandlers, commands.HelpCommand{})
	messageHandlers = append(messageHandlers, commands.ContributeCommand{})
	messageHandlers = append(messageHandlers, commands.AboutCommand{})

	return messageHandlers
}

func processInput(input string, botclient client.BotClient, config conf.Configuration) {
	fmt.Println("input received: " + input)

	var buffer bytes.Buffer
	buffer.WriteString("<messageML>")
	buffer.WriteString(input)
	buffer.WriteString("</messageML>")
	botclient.SendMessageMLMessage(config.StreamId, buffer.String())
}
