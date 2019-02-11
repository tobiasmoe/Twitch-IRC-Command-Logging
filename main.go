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
        client := twitch.NewClient("MrTobbzi", "123")
/*
        client.OnNewClearchatMessage(func(channel string, user twitch.User, message twitch.Message) {
		current_time := time.Now().UTC()
        	file, erro := os.Create(current_time.Format("2006-01-02.txt"))
        	if erro != nil {
               		log.Fatal("cannot create file", erro)
        	}
        	defer file.Close()
		fmt.Fprintf(file, message.Text)
                fmt.Fprintf(file, "\n")
                fmt.Println(message.Text)
        })
*/
        client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
	if (len(user.UserType) != 0 || user.Username == "mrtobbzi") {
		if strings.Contains(message.Text, "!command") {
                        current_time := time.Now().UTC()
			if _ := os.Stat(current_time.Format("2006-01-02.txt")) {
				file, erro := os.Create(current_time.Format("2006-01-02.txt"))
				if erro != nil {
				log.Fatal("cannot create file", erro)
				}
				defer file.Close()
			}
			fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ", ")
                        fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
                        fmt.Println(message.Text)
                }
	}
		/*
                if strings.Contains(message.Text, "!command add") {
                        fmt.Fprintf(file, user.Username)
                        fmt.Fprintf(file, ": ")
                        fmt.Fprintf(file, message.Text)
                        fmt.Fprintf(file, "\n")
                        fmt.Println(message.Text)
                }
                if strings.Contains(message.Text, "!command remove") {
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
		*/
        })

        client.Join("mrtobbzi")

        err := client.Connect()
        if err != nil {
                panic(err)
        }
}



