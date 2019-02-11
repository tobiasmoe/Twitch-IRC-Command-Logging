package main

import (
        "github.com/gempir/go-twitch-irc"
        "fmt"
        "strings"
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
                fmt.Fprintf(filOBe, "\n")
                fmt.Println(message.Text)
        })
*/
        client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
		if (len(user.UserType) != 0 || user.Username == "mrtobbzi") {
			if strings.Contains(message.Text, "!command") {
			        current_time := time.Now().UTC()
				if fileExists(current_time.Format("2006-01-02.txt")) {
					writeFile(current_time.Format("2006-01-02.txt"), user.Username, message.Text)
				} else {
					createFile(current_time.Format("2006-01-02.txt"))
					writeFile(current_time.Format("2006-01-02.txt"), user.Username, message.Text)
				}
				fmt.Println(message.Text)

	                }
		}
        })

        client.Join("mrtobbzi")

        err := client.Connect()
        if err != nil {
                panic(err)
        }
}


func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeFile(filename string, name string, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("%s, %s\n", name, message))
	return err
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	return err
}
