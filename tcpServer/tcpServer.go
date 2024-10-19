package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании слушателя ", err)
		os.Exit(1)
	}
	defer listener.Close()

	//graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("\nЗавершение работы сервера")
		listener.Close()
		wg.Wait()
		fmt.Println("Сервер завершил работу")
		os.Exit(0)
	}()

	fmt.Println("Сервер запущен на порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка подключения ", err)
			break
		}

		wg.Add(1)
		go func(conn net.Conn) {
			defer wg.Done()
			handleConnection(conn)
		}(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message, _ := bufio.NewReader(conn).ReadString('\n')
	//fmt.Println("Новое подключение ", conn.RemoteAddr())
	fmt.Println("Сообщение: ", message)
	conn.Write([]byte("Получено сообщение"))
}
