package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const CHUNKSIZE = 10

var folderPath string = "/home/user/Desktop/golang/FTP/server"

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Waiting for connections ")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		folderPath = "/home/user/Desktop/golang/FTP/server"
		go handleConnection(conn)
	}
}

func handleRead(conn net.Conn) {
	var fileName string = getData(conn)
	filePath := (folderPath + "/" + fileName)
	fl, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err, fl)
		conn.Write([]byte(err.Error()))
		return
	}
	fileSize := fl.Size()
	conn.Write([]byte(strconv.FormatInt(fileSize, 10)))
	time.Sleep(time.Millisecond * 10)
	sendFile(conn, fileSize, filePath)
	fmt.Print("The file has been read ")
}

func handleWrite(conn net.Conn) {
	var fileName string = getData(conn)
	var data string = getData(conn)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	conn.Write([]byte("Written successfully"))

	fmt.Println(fileName, data)
}
func handleList(conn net.Conn) {
	fmt.Println(folderPath)
	dir, err := os.Open(folderPath)
	if err != nil {
		fmt.Println(err)
		conn.Write([]byte(err.Error()))
		return
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading folder contents:", err)
		conn.Write([]byte(err.Error()))
		return
	}

	var folderData string
	for _, file := range files {
		folderData += file.Name() + "\n"
	}
	conn.Write([]byte(folderData))

}
func handleChangeDirectory(conn net.Conn) {
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	folderName := string(buff[:n])
	if folderName == ".." {
		index := strings.LastIndex(folderPath, "/")
		folderPath = folderPath[:index]
		if len(folderPath) == 0 {
			folderPath = "/"
		}
		conn.Write([]byte("done"))
		return
	}
	fileExists := directoryVerifier(folderName)
	if fileExists {
		folderPath += "/" + folderName
		fmt.Println(folderName)
		conn.Write([]byte("done"))
	} else {
		conn.Write([]byte("No directory found"))
	}

}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buff := make([]byte, 32768)
		n, err := conn.Read(buff)
		if err != nil {
			return
		}
		operation := string(buff[:n])
		fmt.Println(operation)
		switch operation {
		case "Read":
			{
				handleRead(conn)
				break
			}
		case "Write":
			{
				handleWrite(conn)
				break
			}
		case "List":
			{
				handleList(conn)
				break
			}
		case "ChangeDirectory":
			{
				handleChangeDirectory(conn)
				break
			}
		case "Stop":
			{
				conn.Close()
				return
			}
		default:
			{
				fmt.Println("Wrong state was passed")
				return
			}
		}

	}
}

func min(a int64, b int64) int64 {
	if a > b {
		fmt.Println(a, b, b)
		return b
	}
	fmt.Println(a, b, a)

	return a
}

func directoryVerifier(directoryName string) bool {
	dir, err := os.Open(folderPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading folder contents:", err)
		return false
	}
	for _, file := range files {
		if file.Name() == directoryName {
			return true
		}
	}
	return false
}
