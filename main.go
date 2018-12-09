package main

import (
	"github.com/gempir/go-twitch-irc"
	"fmt"
	"strings"
	"log"
	"os"
	"time"
)

func main() {
	current_time := time.Now().UTC()
	file, erro := os.Create(current_time.Format("2006-01-02.txt"))
	if erro != nil {
		log.Fatal("cannot create file", erro)
	}
	defer file.Close()
	client := twitch.NewClient("justinfan123123", "oauth:123123123")

	client.OnNewClearchatMessage(func(channel string, user twitch.User, message twitch.Message) {
		fmt.Fprintf(file, message.Text)
		fmt.Fprintf(file, "\n")
		fmt.Println(message.Text)
	})

	client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
		if strings.Contains(message.Text, "!commands") {
			fmt.Fprintf(file, user.Username)
			fmt.Fprintf(file, ": ")
			fmt.Fprintf(file, message.Text)
			fmt.Fprintf(file, "\n")
			fmt.Println(message.Text)
		}
		if strings.Contains(message.Text, "!addcom") {
			fmt.Fprintf(file, user.Username)
			fmt.Fprintf(file, ": ")
			fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
			fmt.Println(message.Text)
		}
		if strings.Contains(message.Text, "!delcom") {
			fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
			fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
			fmt.Println(message.Text)
		}
		if strings.Contains(message.Text, "!editcom") {
			fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
			fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
			fmt.Println(message.Text)
		}
		if strings.Contains(message.Text, "!command add") {
			fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
                        fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
                        fmt.Println(message.Text)
		}
		if strings.Contains(message.Text, "!command edit") {
                        fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
                        fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
                        fmt.Println(message.Text)
                }
		if strings.Contains(message.Text, "!command delete") {
                        fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
                        fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
                        fmt.Println(message.Text)
                }
	})

	client.Join("MrTobbzi")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
