package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

func getFileName(message string) string {
	data := getData(message)
	fileName := strings.Split(data, " ")
	return fileName[0]
}
func getData(messsage string) string {
	// var fileName string
	if len(messsage) != 0 {
		getMessageColor().Print(messsage, "\t")
	}
	stdreader := bufio.NewReader(os.Stdin)
	cmd, _ := stdreader.ReadString('\n')
	cmd = strings.Trim(cmd, "\n")
	return cmd
}
func getFile(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, _ := conn.Read(buffer)
		if string(buffer[:n]) == "\\tOver" {
			break
		}
		fmt.Print(string(buffer[:n]))
		if n < CHUNKSIZE {
			break
		}
	}
	fmt.Print("\n\n")
}

func errorMessage(err string) {
	getErrorColor().Println(err)
}

func getCmdFontColor() *color.Color {
	customColor := color.New(color.FgGreen).Add(color.Bold)
	return customColor
}
func getMessageColor() *color.Color {
	customColor := color.New(color.FgBlue)
	return customColor
}
func getErrorColor() *color.Color {
	customColor := color.New(color.FgRed).Add(color.Bold)
	return customColor
}
