package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка подключения ", err)
		os.Exit(1)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите сообщения ")
	message, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, message)
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Ответ от сервера ", response)
}
