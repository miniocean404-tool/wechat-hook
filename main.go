package main

import (
	"log"
	"wechat_hook/src/tcp"
)

// 用于监听 hook 中的消息的服务
func main() {
	// 	Ldate         = 1 << iota     // 日期：2009/01/23
	// 	Ltime                         // 时间：01:23:23
	// 	Lmicroseconds                 // 微妙级别时间：01:23:23.123123 （用于增强 Ltime 位）
	// 	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	// 	Lshortfile                    // 文件名+行号：d.go:23 （会覆盖 Llongfile）
	// 	LUTC                          // 使用 UTC 时间
	// 	Lmsgprefix                    // 在每个日志信息前面添加信息前缀
	// 	LstdFlags     = Ldate | Ltime // 标准 logger 的初始值
	log.SetFlags(log.Lmsgprefix | log.Ldate | log.Ltime | log.Lshortfile)

	tcp.Listen(19099)
}
