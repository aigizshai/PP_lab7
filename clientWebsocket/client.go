package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

type Message struct {
	UserName string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	ws, err := websocket.Dial("ws://localhost:8081/ws", "", "http://localhost/")
	if err != nil {
		log.Fatal("Ошибка подключения ", err)
	}
	defer ws.Close()

	fmt.Println("Подключение к серверу успешно. Введите имя")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	go func() {
		for {
			var msg Message
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				log.Println("Ошибка получения сообщения ", err)
			}
			fmt.Printf("%s : %s\n", msg.UserName, msg.Message)
		}
	}()

	for {
		//fmt.Print("сообщение: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "/q" {
			ws.Close()
			os.Exit(0)
			break
		}
		msg := Message{
			UserName: username,
			Message:  text,
		}

		err := websocket.JSON.Send(ws, msg)
		if err != nil {
			log.Println("Ошибка отправки ", err)
			return
		}
	}
}
