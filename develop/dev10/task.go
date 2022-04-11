package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeoutPtr := flag.Int("timeout", 10, "таймаут на подключение к серверу.")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Println("usage: [flags] host port")
		flag.PrintDefaults()
		os.Exit(-1)
	}
	address := fmt.Sprintf("%s:%s", args[0], args[1])

	// Подключаемся к сокету
	d := net.Dialer{Timeout: time.Duration(*timeoutPtr) * time.Second}
	conn, err := d.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	var text string

	for {
		fmt.Print("Text to send: ")
		_, err := fmt.Scanf("%s\r\n", &text)
		if err == io.EOF {
			break
		}
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			fmt.Println("Connection closed")
			break
		}

		if err != nil {
			fmt.Printf("Error while read: %s\n", err)
			break
		}

		fmt.Println("Message from server: " + message)
	}

}
