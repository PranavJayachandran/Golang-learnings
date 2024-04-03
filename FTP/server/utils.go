package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func getData(conn net.Conn) string {
	data := make([]byte, 1024)
	n, _ := conn.Read(data)
	return string(data[:n])
}

func sendFile(conn net.Conn, fileSize int64, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}
	for {
		buffer := make([]byte, min(fileSize, CHUNKSIZE))
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			conn.Write([]byte(err.Error()))
			time.Sleep(time.Millisecond * 10)
			conn.Write([]byte("\\tOver"))
			return
		}
		if err == io.EOF {
			return
		}
		conn.Write(buffer[:n])
		time.Sleep(time.Millisecond * 10)
		fileSize -= int64(n)
		// To know when the complete file has been send
		if fileSize <= 0 {
			conn.Write([]byte("\\tOver"))
		}
	}
}
