package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/gookit/goutil/dump"
)

func Listen(port int) {
	p := strconv.Itoa(port)
	adress := "127.0.0.1:" + p
	ln, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("TCP 服务启动")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleMsg(conn)
	}
}

func handleMsg(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生了未处理的异常", err)
		}
	}()

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Bytes()

		var message Message
		err := json.Unmarshal(line, &message)
		if err != nil {
			fmt.Println("解析JSON失败:", err)
			return
		}

		// 可以使用 xml 读取消息
		dump.Print(message)
		// log.Printf("%+v", message)
	}

	if err := scanner.Err(); err != nil {
		log.Println("错误：", err)
	}
}
