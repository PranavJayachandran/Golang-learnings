package main

import (
	"io"
	"log"
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
		log.Fatal(err)
		return
	}
	for {
		buffer := make([]byte, min(fileSize, CHUNKSIZE))
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			return
		}
		conn.Write(buffer[:n])
		time.Sleep(time.Millisecond * 10)
		fileSize -= int64(n)
		// To know when the complete file has been send
		if fileSize <= 0 {
			break
		}
	}
}
