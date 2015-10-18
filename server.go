package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var IP = getIP()
var PORT = getPort()

func main() {
	listen, err := net.Listen("tcp", IP+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Listening on " + IP + ":" + PORT)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	connReader := bufio.NewReader(conn)
	message, _ := connReader.ReadString('\n')
	message = strings.TrimSuffix(message, "\r\n")
	response := "\n"
	if strings.HasPrefix(message, "HELO ") {
		text := strings.TrimPrefix(message, "HELO")
		text = strings.TrimSuffix(text, "\\n")
		response = HELO(text)
	} else if message == "KILL_SERVICE\\n" {
		killService()
	} else {
		response = otherMessage(message)
	}
	fmt.Fprintf(conn, response)
}

func killService() {
	os.Exit(0)
}

func otherMessage(message string) string {
	return message + "\n"
}

func HELO(text string) string {
	return "HELO " + text + "\\nIP:" + IP + "\\nPort:" + PORT + "\\nStudentID:ea5f6b94d6a8a8f1e7890f6a64883cdc2b6125821e20ddd36a33b773bd46b727\\n\n"
}

func getIP() string {
	e := os.Getenv("CS4032_LAB_2_IP")
	if len(e) > 0 {
		return e
	}
	return "127.0.0.1"
}

func getPort() string {
	e := os.Getenv("CS4032_LAB_2_PORT")
	if len(e) > 0 {
		return e
	}
	return "8080"
}
