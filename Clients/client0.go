package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]

	l, err := net.Dial("tcp4", "127.0.0.1"+PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	var sendMessage string = arguments[2] + "\n"
	l.Write([]byte(sendMessage))

	netData, err := bufio.NewReader(l).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(netData)

	l.Close()
}
