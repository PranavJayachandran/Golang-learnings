package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const CHUNKSIZE = 10

func read(conn net.Conn, fileName string) {
	conn.Write([]byte("Read"))
	time.Sleep(time.Millisecond * 10)
	conn.Write([]byte(fileName))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	reply := string(buffer[:n])
	size, err := strconv.Atoi(reply)
	if err != nil {
		errorMessage(reply)
		return
	}
	fmt.Println("Incoming file size is ", size)
	getFile(conn)
}
func write(conn net.Conn, fileName string, dataToWrite []string) {
	conn.Write([]byte("Write"))
	time.Sleep(time.Millisecond * 10)
	conn.Write([]byte(fileName))
	time.Sleep(time.Millisecond * 10)
	conn.Write([]byte(strings.Join(dataToWrite, " ")))

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Print(string(buffer[:n]), "\n\n")
}
func list(conn net.Conn) {
	conn.Write([]byte("List"))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Print(string(buffer[:n]), "\n\n")
}
func changeDirectory(conn net.Conn, directoryName string) {
	conn.Write([]byte("ChangeDirectory"))
	time.Sleep(time.Millisecond * 10)
	conn.Write([]byte(directoryName))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	reply := string(buffer[:n])
	if reply != "done" {
		errorMessage(reply)
	}
}
func main() {
	var commands []string = []string{
		"read file_name\tTo read a file",
		"write file_name data_to_write\tTo write a file",
		"list\tTo list all the files and folders in the current directory in the remote folder",
		"cd folder_name\tEnter into a folder",
	}
	cmdTextStyle := getCmdFontColor()
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		var breakloop bool = false
		var input string
		cmdTextStyle.Print("Type the commands  ")
		input = getData("")
		userInput := strings.Split(input, " ")
		switch userInput[0] {
		case "read":
			{
				if len(userInput) != 2 {
					errorMessage("Should follow the format read file_name")
					break
				}
				read(conn, userInput[1])
				break
			}
		case "write":
			{
				if len(userInput) < 3 {
					errorMessage("Should be of format write file_name data_to_write")
					break
				}
				write(conn, userInput[1], userInput[2:])
				break
			}
		case "list":
			{
				list(conn)
				break
			}
		case "help":
			{
				for _, command := range commands {
					fmt.Println(command)
				}
				break
			}
		case "cd":
			{
				if len(userInput) != 2 {
					errorMessage("Should follow the format cd folder_name or cd ..")
					break
				}
				changeDirectory(conn, userInput[1])

			}
		default:
			{
				errorMessage("Wrong command use help for more info")
				break
			}
		}
		if breakloop {
			break
		}

	}
	conn.Close()
}
